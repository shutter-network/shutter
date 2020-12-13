package keyper

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
)

type Keyper2 struct {
	Config    KeyperConfig
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain

	ContractCaller ContractCaller
	shmcl          client.Client
	MessageSender  MessageSender
	Interactive    bool
}

func NewKeyper2(kc KeyperConfig) Keyper2 {
	return Keyper2{
		Config:    kc,
		State:     &State{},
		Shutter:   observe.NewShutter(),
		MainChain: observe.NewMainChain(),

		// batchConfigs:          make(map[uint64]contract.BatchConfig),
		// batches:               make(map[uint64]*BatchState),
		// checkedIn:             false,
		// newHeaders:            make(chan *types.Header, newHeadersSize),
		// cipherExecutionParams: make(chan CipherExecutionParams),
		// keyperEncryptionKeys:  make(map[common.Address]*ecies.PublicKey),
		// dkg:                   make(map[uint64]*DKGInstance),
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

	return NewContractCaller(
		ethcl,
		config.SigningKey,
		configContract,
		keyBroadcastContract,
		batcherContract,
		executorContract,
	), nil
}

func (kpr *Keyper2) init() error {
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

	// executor := Executor{
	//	ctx:                   kpr.ctx,
	//	client:                kpr.ethcl,
	//	cc:                    &contractCaller,
	//	cipherExecutionParams: kpr.cipherExecutionParams,
	// }
	// kpr.executor = executor
}

func (kpr *Keyper2) syncMain(ctx context.Context) error {
	return kpr.MainChain.SyncToHead(ctx, kpr.ContractCaller.Ethclient, kpr.ContractCaller.ConfigContract)
}

func (kpr *Keyper2) syncShutter(ctx context.Context) error {
	return kpr.Shutter.SyncToHead(ctx, kpr.shmcl)
}

func (kpr *Keyper2) sync(ctx context.Context) error {
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

func (kpr *Keyper2) Run() error {
	err := kpr.init()
	if err != nil {
		return err
	}
	ctx := context.Background()

	for {
		err = kpr.sync(ctx)
		if err != nil {
			return err
		}

		fmt.Println("-----------------------------------------------------------------------")
		pretty.Println("SHUTTER   ==>", *kpr.Shutter)
		pretty.Println("MAINCHAIN ==>", *kpr.MainChain)
		kpr.runOneStep(ctx)
		pretty.Println("INTERNAL  ==>", *kpr.State)

		time.Sleep(10 * time.Second)
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

func (kpr *Keyper2) runOneStep(ctx context.Context) {
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
	log.Printf("Running %d actions", len(decider.Actions))

	for _, act := range decider.Actions {
		err := act.Run(ctx, kpr.MessageSender)
		// XXX at the moment we just let the whole program die. We need a better strategy
		// here. We could retry the actions or feed the errors back into our state
		if err != nil {
			panic(err)
		}
	}
}
