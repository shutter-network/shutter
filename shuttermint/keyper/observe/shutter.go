package observe

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	pkgErrors "github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client"

	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

var errEonNotFound = errors.New("eon not found")

func init() {
	gob.Register(ethcrypto.S256()) // Allow gob to serialize ecsda.PrivateKey
}

// EncryptionPublicKey is a gob serializable version of ecies.PublicKey
type EncryptionPublicKey ecies.PublicKey

func (epk *EncryptionPublicKey) GobEncode() ([]byte, error) {
	return ethcrypto.FromECDSAPub((*ecies.PublicKey)(epk).ExportECDSA()), nil
}

func (epk *EncryptionPublicKey) GobDecode(data []byte) error {
	pubkey, err := ethcrypto.UnmarshalPubkey(data)
	if err != nil {
		return err
	}
	*epk = *(*EncryptionPublicKey)(ecies.ImportECDSAPublic(pubkey))
	return nil
}

// Encrypt the given message m
func (epk *EncryptionPublicKey) Encrypt(rand io.Reader, m []byte) ([]byte, error) {
	return ecies.Encrypt(rand, (*ecies.PublicKey)(epk), m, nil, nil)
}

// Shutter let's a keyper fetch all necessary information from a shuttermint node. The only source
// for the data stored in this struct should be the shutter node.  The SyncToHead method can be
// used to update the data. All other accesses should be read-only.
type Shutter struct {
	CurrentBlock         int64
	KeyperEncryptionKeys map[common.Address]*EncryptionPublicKey
	BatchConfigs         []shutterevents.BatchConfig
	Batches              map[uint64]*BatchData
	Eons                 []Eon
}

// NewShutter creates an empty Shutter struct
func NewShutter() *Shutter {
	return &Shutter{
		CurrentBlock:         -1,
		KeyperEncryptionKeys: make(map[common.Address]*EncryptionPublicKey),
		Batches:              make(map[uint64]*BatchData),
	}
}

type Eon struct {
	Eon                  uint64
	StartHeight          int64
	StartEvent           shutterevents.EonStarted
	Commitments          []shutterevents.PolyCommitment
	PolyEvals            []shutterevents.PolyEval
	Accusations          []shutterevents.Accusation
	Apologies            []shutterevents.Apology
	EpochSecretKeyShares []shutterevents.EpochSecretKeyShare
}

type BatchData struct {
	BatchIndex           uint64
	DecryptionSignatures []shutterevents.DecryptionSignature
}

func (shutter *Shutter) applyTxEvents(height int64, events []abcitypes.Event) {
	for _, ev := range events {
		x, err := shutterevents.MakeEvent(ev, height)
		if err != nil {
			log.Printf("Error: malformed event: %+v ev=%+v", err, ev)
		} else {
			shutter.applyEvent(x)
		}
	}
}

func (shutter *Shutter) getBatchData(batchIndex uint64) *BatchData {
	b, ok := shutter.Batches[batchIndex]
	if !ok {
		b = &BatchData{BatchIndex: batchIndex}
		shutter.Batches[batchIndex] = b
	}
	return b
}

func (shutter *Shutter) searchEon(eon uint64) int {
	return sort.Search(
		len(shutter.Eons),
		func(i int) bool {
			return eon <= shutter.Eons[i].Eon
		},
	)
}

func (shutter *Shutter) FindEonByBatchIndex(batchIndex uint64) (*Eon, error) {
	for i := len(shutter.Eons) - 1; i >= 0; i-- {
		if shutter.Eons[i].StartEvent.BatchIndex <= batchIndex {
			return &shutter.Eons[i], nil
		}
	}
	return nil, pkgErrors.WithStack(errEonNotFound)
}

func (shutter *Shutter) FindEon(eon uint64) (*Eon, error) {
	idx := shutter.searchEon(eon)
	if idx == len(shutter.Eons) || eon < shutter.Eons[idx].Eon {
		return nil, pkgErrors.WithStack(errEonNotFound)
	}
	return &shutter.Eons[idx], nil
}

func (shutter *Shutter) applyCheckIn(e shutterevents.CheckIn) error { //nolint:unparam
	shutter.KeyperEncryptionKeys[e.Sender] = (*EncryptionPublicKey)(e.EncryptionPublicKey)
	return nil
}

func (shutter *Shutter) applyBatchConfig(e shutterevents.BatchConfig) error { //nolint:unparam
	shutter.BatchConfigs = append(shutter.BatchConfigs, e)
	return nil
}

func (shutter *Shutter) applyDecryptionSignature(e shutterevents.DecryptionSignature) error { //nolint:unparam
	b := shutter.getBatchData(e.BatchIndex)
	b.DecryptionSignatures = append(b.DecryptionSignatures, e)
	return nil
}

func (shutter *Shutter) applyEonStarted(e shutterevents.EonStarted) error {
	idx := shutter.searchEon(e.Eon)
	if idx < len(shutter.Eons) {
		return pkgErrors.Errorf("eons should increase")
	}
	shutter.Eons = append(shutter.Eons, Eon{Eon: e.Eon, StartEvent: e, StartHeight: e.Height})
	return nil
}

func (shutter *Shutter) applyPolyCommitment(e shutterevents.PolyCommitment) error {
	eon, err := shutter.FindEon(e.Eon)
	if err != nil {
		return err
	}
	eon.Commitments = append(eon.Commitments, e)
	return nil
}

func (shutter *Shutter) applyPolyEval(e shutterevents.PolyEval) error {
	eon, err := shutter.FindEon(e.Eon)
	if err != nil {
		return err
	}
	eon.PolyEvals = append(eon.PolyEvals, e)
	return nil
}

func (shutter *Shutter) applyAccusation(e shutterevents.Accusation) error {
	eon, err := shutter.FindEon(e.Eon)
	if err != nil {
		return err
	}
	eon.Accusations = append(eon.Accusations, e)
	return nil
}

func (shutter *Shutter) applyApology(e shutterevents.Apology) error {
	eon, err := shutter.FindEon(e.Eon)
	if err != nil {
		return err
	}
	eon.Apologies = append(eon.Apologies, e)
	return nil
}

func (shutter *Shutter) applyEpochSecretKeyShare(e shutterevents.EpochSecretKeyShare) error {
	eon, err := shutter.FindEon(e.Eon)
	if err != nil {
		return err
	}
	eon.EpochSecretKeyShares = append(eon.EpochSecretKeyShares, e)
	return nil
}

func (shutter *Shutter) applyEvent(ev shutterevents.IEvent) {
	var err error
	switch e := ev.(type) {
	case *shutterevents.CheckIn:
		err = shutter.applyCheckIn(*e)
	case *shutterevents.BatchConfig:
		err = shutter.applyBatchConfig(*e)
	case *shutterevents.DecryptionSignature:
		err = shutter.applyDecryptionSignature(*e)
	case *shutterevents.EonStarted:
		err = shutter.applyEonStarted(*e)
	case *shutterevents.PolyCommitment:
		err = shutter.applyPolyCommitment(*e)
	case *shutterevents.PolyEval:
		err = shutter.applyPolyEval(*e)
	case *shutterevents.Accusation:
		err = shutter.applyAccusation(*e)
	case *shutterevents.Apology:
		err = shutter.applyApology(*e)
	case *shutterevents.EpochSecretKeyShare:
		err = shutter.applyEpochSecretKeyShare(*e)
	default:
		err = pkgErrors.Errorf("not yet implemented for %s", reflect.TypeOf(ev))
	}
	if err != nil {
		log.Printf("Error in apply event: %+v, event: %+v", err, ev)
	}
}

func (shutter *Shutter) fetchAndApplyEvents(ctx context.Context, shmcl client.Client, targetHeight int64) error {
	if targetHeight < shutter.CurrentBlock {
		panic("internal error: fetchAndApplyEvents bad arguments")
	}
	query := fmt.Sprintf("tx.height >= %d and tx.height <= %d", shutter.CurrentBlock+1, targetHeight)

	// tendermint silently caps the perPage value at 100, make sure to stay below, otherwise
	// our exit condition is wrong and the log.Fatalf will trigger a panic below; see
	// https://github.com/brainbot-com/shutter/issues/50
	perPage := 100
	page := 1
	total := 0
	for {
		res, err := shmcl.TxSearch(ctx, query, false, &page, &perPage, "")
		if err != nil {
			return err
		}
		total += len(res.Txs)
		for _, tx := range res.Txs {
			events := tx.TxResult.GetEvents()
			shutter.applyTxEvents(tx.Height, events)
		}
		if page*perPage >= res.TotalCount {
			if total != res.TotalCount {
				log.Fatalf("internal error. got %d transactions, expected %d transactions from shuttermint for height %d..%d",
					total,
					res.TotalCount,
					shutter.CurrentBlock+1,
					targetHeight)
			}
			break
		}
		page++
	}
	return nil
}

// IsCheckedIn checks if the given address sent it's check-in message
func (shutter *Shutter) IsCheckedIn(addr common.Address) bool {
	_, ok := shutter.KeyperEncryptionKeys[addr]
	return ok
}

// IsKeyper checks if the given address is a keyper in any of the given configs
func (shutter *Shutter) IsKeyper(addr common.Address) bool {
	for _, cfg := range shutter.BatchConfigs {
		if cfg.IsKeyper(addr) {
			return true
		}
	}
	return false
}

func (shutter *Shutter) FindBatchConfigByConfigIndex(configIndex uint64) (shutterevents.BatchConfig, error) {
	for _, bc := range shutter.BatchConfigs {
		if bc.ConfigIndex == configIndex {
			return bc, nil
		}
	}
	return shutterevents.BatchConfig{}, pkgErrors.Errorf("cannot find BatchConfig with ConfigIndex==%d", configIndex)
}

func (shutter *Shutter) FindBatchConfigByBatchIndex(batchIndex uint64) shutterevents.BatchConfig {
	for i := len(shutter.BatchConfigs); i > 0; i++ {
		if shutter.BatchConfigs[i-1].StartBatchIndex <= batchIndex {
			return shutter.BatchConfigs[i-1]
		}
	}
	return shutterevents.BatchConfig{}
}

func (shutter *Shutter) Clone() *Shutter {
	clone := new(Shutter)
	medley.CloneWithGob(shutter, clone)
	return clone
}

func (shutter *Shutter) LastCommittedHeight(ctx context.Context, shmcl client.Client) (int64, error) {
	latestBlock, err := shmcl.Block(ctx, nil)
	if err != nil {
		return 0, err
	}
	if latestBlock.Block == nil || latestBlock.Block.LastCommit == nil {
		return 0, pkgErrors.Errorf("empty blockchain: %+v", latestBlock)
	}
	return latestBlock.Block.LastCommit.Height, nil
}

// SyncToHead syncs the state with the remote state. It fetches events from new blocks since the
// last sync and updates the state by calling applyEvent for each event. This method does not
// mutate the object in place, it rather returns a new object.
func (shutter *Shutter) SyncToHead(ctx context.Context, shmcl client.Client) (*Shutter, error) {
	height, err := shutter.LastCommittedHeight(ctx, shmcl)
	if err != nil {
		return nil, err
	}
	return shutter.SyncToHeight(ctx, shmcl, height)
}

// SyncToHeight syncs the state with the remote state until the given height
func (shutter *Shutter) SyncToHeight(ctx context.Context, shmcl client.Client, height int64) (*Shutter, error) {
	clone := shutter.Clone()
	err := clone.fetchAndApplyEvents(ctx, shmcl, height)
	if err != nil {
		return nil, err
	}
	clone.CurrentBlock = height
	return clone, nil
}
