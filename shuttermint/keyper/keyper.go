package keyper

import (
	"bufio"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

const (
	runSleepTime = 10 * time.Second

	// watchedTransactionsBufferSize is the maximum number of txs to watch. If this is too small,
	// actions will stall.
	watchedTransactionsBufferSize = 100
)

// IsWebsocketURL returns true iff the given URL is a websocket URL, i.e. if it starts with ws://
// or wss://. This is needed for the watchMainChainHeadBlock method.
func IsWebsocketURL(url string) bool {
	return strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")
}

// Address returns the keyper's Ethereum address.
func (c *KeyperConfig) Address() common.Address {
	return crypto.PubkeyToAddress(c.SigningKey.PublicKey)
}

type Keyper struct {
	Config    KeyperConfig
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain

	ContractCaller      ContractCaller
	shmcl               client.Client
	MessageSender       MessageSender
	WatchedTransactions chan *types.Transaction
	Interactive         bool
}

func NewKeyper(kc KeyperConfig) Keyper {
	return Keyper{
		Config:    kc,
		State:     &State{},
		Shutter:   observe.NewShutter(),
		MainChain: observe.NewMainChain(),

		WatchedTransactions: make(chan *types.Transaction, watchedTransactionsBufferSize),
	}
}

func NewContractCallerFromConfig(config KeyperConfig) (ContractCaller, error) {
	ethcl, err := ethclient.Dial(config.EthereumURL)
	if err != nil {
		return ContractCaller{}, err
	}
	configContract, err := contract.NewConfigContract(config.ConfigContractAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	keyBroadcastContract, err := contract.NewKeyBroadcastContract(config.KeyBroadcastContractAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	batcherContract, err := contract.NewBatcherContract(config.BatcherContractAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	executorContract, err := contract.NewExecutorContract(config.ExecutorContractAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	depositContract, err := contract.NewDepositContract(config.DepositContractAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	keyperSlasher, err := contract.NewKeyperSlasher(config.KeyperSlasherAddress, ethcl)
	if err != nil {
		return ContractCaller{}, err
	}

	return NewContractCaller(
		ethcl,
		config.SigningKey,
		configContract,
		keyBroadcastContract,
		batcherContract,
		executorContract,
		depositContract,
		keyperSlasher,
	), nil
}

func (kpr *Keyper) init() error {
	if kpr.shmcl != nil {
		panic("internal error: already initialized")
	}
	var err error
	kpr.shmcl, err = http.New(kpr.Config.ShuttermintURL, "/websocket")
	if err != nil {
		return errors.Wrapf(err, "create shuttermint client at %s", kpr.Config.ShuttermintURL)
	}
	ms := NewRPCMessageSender(kpr.shmcl, kpr.Config.SigningKey)
	kpr.MessageSender = &ms

	kpr.ContractCaller, err = NewContractCallerFromConfig(kpr.Config)
	return err
}

func (kpr *Keyper) syncMain(ctx context.Context) error {
	newMainChain, err := kpr.MainChain.SyncToHead(
		ctx,
		kpr.ContractCaller.Ethclient,
		kpr.ContractCaller.ConfigContract,
		kpr.ContractCaller.BatcherContract,
		kpr.ContractCaller.ExecutorContract,
		kpr.ContractCaller.DepositContract,
		kpr.ContractCaller.KeyperSlasher,
	)
	if err != nil {
		return err
	}
	kpr.MainChain = newMainChain
	return nil
}

func (kpr *Keyper) syncShutter(ctx context.Context) error {
	newShutter, err := kpr.Shutter.SyncToHead(ctx, kpr.shmcl)
	if err != nil {
		return err
	}
	kpr.Shutter = newShutter
	return nil
}

func (kpr *Keyper) sync(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return kpr.syncShutter(ctx)
	})
	group.Go(func() error {
		return kpr.syncMain(ctx)
	})
	err := group.Wait()
	return err
}

func (kpr *Keyper) ShortInfo() string {
	var dkgInfo []string
	for _, dkg := range kpr.State.DKGs {
		dkgInfo = append(dkgInfo, dkg.ShortInfo())
	}
	return fmt.Sprintf(
		"shutter block %d, main chain %d, last eon started %d, DKGs: %s",
		kpr.Shutter.CurrentBlock,
		kpr.MainChain.CurrentBlock,
		kpr.State.LastEonStarted,
		strings.Join(dkgInfo, " - "),
	)
}

func (kpr *Keyper) Run() error {
	err := kpr.init()
	if err != nil {
		return err
	}
	ctx := context.Background()

	go kpr.watchTransactions(ctx)

	for {
		err = kpr.sync(ctx)
		if err != nil {
			log.Printf("Error in sync: %s", err)
		} else {
			log.Println(kpr.ShortInfo())
			kpr.runOneStep(ctx)
		}

		time.Sleep(runSleepTime)
	}
}

func readline() {
	fmt.Printf("\n[press return to continue] > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf("\n")
}

type storedState struct {
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain
}

func (kpr *Keyper) gobpath() string {
	return filepath.Join(kpr.Config.DBDir, "state.gob")
}

func (kpr *Keyper) LoadState() error {
	gobpath := kpr.gobpath()

	gobfile, err := os.Open(gobpath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	log.Printf("Loading state from %s", gobpath)

	defer gobfile.Close()
	dec := gob.NewDecoder(gobfile)
	st := storedState{}
	err = dec.Decode(&st)
	if err != nil {
		return err
	}
	kpr.State = st.State
	kpr.Shutter = st.Shutter
	kpr.MainChain = st.MainChain
	return nil
}

func (kpr *Keyper) saveState() error {
	gobpath := kpr.gobpath()
	log.Printf("Saving state to %s", gobpath)
	tmppath := gobpath + ".tmp"
	file, err := os.Create(tmppath)
	if err != nil {
		return err
	}
	defer file.Close()
	st := storedState{
		State:     kpr.State,
		Shutter:   kpr.Shutter,
		MainChain: kpr.MainChain,
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(st)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}
	err = os.Rename(tmppath, gobpath)
	return err
}

func (kpr *Keyper) runOneStep(ctx context.Context) {
	decider := Decider{
		Config:    kpr.Config,
		State:     kpr.State,
		Shutter:   kpr.Shutter,
		MainChain: kpr.MainChain,
	}
	decider.Decide()
	if kpr.Interactive && len(decider.Actions) > 0 {
		log.Printf("Showing %d actions", len(decider.Actions))
		for _, act := range decider.Actions {
			fmt.Println(act)
		}
		readline()
	}
	err := kpr.saveState()
	if err != nil {
		panic(err)
	}
	log.Printf("Running %d actions", len(decider.Actions))

	runenv := RunEnv{
		MessageSender:       kpr.MessageSender,
		ContractCaller:      &kpr.ContractCaller,
		WatchedTransactions: kpr.WatchedTransactions,
	}
	for _, act := range decider.Actions {
		err := act.Run(ctx, runenv)
		// XXX at the moment we just let the whole program die. We need a better strategy
		// here. We could retry the actions or feed the errors back into our state
		if err != nil {
			panic(err)
		}
	}
}

func (kpr *Keyper) watchTransactions(ctx context.Context) {
	for {
		select {
		case tx := <-kpr.WatchedTransactions:
			receipt, err := medley.WaitMined(ctx, kpr.ContractCaller.Ethclient, tx.Hash())
			if err != nil {
				log.Printf("Error wating for transaction %s: %v", tx.Hash().Hex(), err)
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				log.Printf("Tx %s has failed and was reverted", tx.Hash().Hex())
			} else {
				log.Printf("Tx %s was successful", tx.Hash().Hex())
			}
		case <-ctx.Done():
			return
		}
	}
}
