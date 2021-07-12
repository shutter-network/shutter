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
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	pkgErrors "github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/rpc/client"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/shutter-network/shutter/shuttermint/keyper/shutterevents"
	"github.com/shutter-network/shutter/shuttermint/medley"
)

const (
	shuttermintTimeout       = 10 * time.Second
	shutterReconnectInterval = 5 * time.Second
)

var errEonNotFound = errors.New("eon not found")

func init() {
	// Allow gob to serialize ecsda.PrivateKey and ed25519.PubKey
	gob.Register(ethcrypto.S256())
	gob.Register(ed25519.GenPrivKeyFromSecret([]byte{}).PubKey())
}

// EncryptionPublicKey is a gob serializable version of ecies.PublicKey.
type EncryptionPublicKey ecies.PublicKey

func (epk *EncryptionPublicKey) GobEncode() ([]byte, error) {
	return ethcrypto.FromECDSAPub((*ecies.PublicKey)(epk).ExportECDSA()), nil
}

func (epk *EncryptionPublicKey) GobDecode(data []byte) error {
	pubkey, err := ethcrypto.UnmarshalPubkey(data)
	if err != nil {
		return pkgErrors.Wrap(err, "failed to unmarshal encryption public key")
	}
	*epk = *(*EncryptionPublicKey)(ecies.ImportECDSAPublic(pubkey))
	return nil
}

// Encrypt the given message m.
func (epk *EncryptionPublicKey) Encrypt(rand io.Reader, m []byte) ([]byte, error) {
	return ecies.Encrypt(rand, (*ecies.PublicKey)(epk), m, nil, nil)
}

// ShutterFilter is used to filter the shutter state we do build. Filtering is done in
// Shutter.ApplyFilter.
type ShutterFilter struct {
	SyncHeight int64
	BatchIndex uint64
}

func (filter ShutterFilter) NeedsUpdate(newFilter ShutterFilter) bool {
	return newFilter.SyncHeight > filter.SyncHeight || newFilter.BatchIndex > filter.BatchIndex
}

// Shutter let's a keyper fetch all necessary information from a shuttermint node. The only source
// for the data stored in this struct should be the shutter node.  The SyncToHead method can be
// used to update the data. All other accesses should be read-only.
type Shutter struct {
	CurrentBlock         int64
	LastCommittedHeight  int64
	NodeStatus           *rpctypes.ResultStatus
	KeyperEncryptionKeys map[common.Address]*EncryptionPublicKey
	BatchConfigs         []shutterevents.BatchConfig
	Batches              map[uint64]*BatchData
	Eons                 []Eon
	Filter               ShutterFilter
}

// NewShutter creates an empty Shutter struct.
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

func (eon *Eon) ApplyFilter(syncHeight int64) *Eon {
	clone := Eon{
		Eon:         eon.Eon,
		StartHeight: eon.StartHeight,
		StartEvent:  eon.StartEvent,
	}
	clone.Commitments = append(clone.Commitments, eon.GetPolyCommitments(syncHeight)...)
	clone.PolyEvals = append(clone.PolyEvals, eon.GetPolyEvals(syncHeight)...)
	clone.Accusations = append(clone.Accusations, eon.GetAccusations(syncHeight)...)
	clone.Apologies = append(clone.Apologies, eon.GetApologies(syncHeight)...)
	return &clone
}

func (eon *Eon) GetPolyCommitments(syncHeight int64) []shutterevents.PolyCommitment {
	slice := eon.Commitments
	idx := sort.Search(len(slice),
		func(i int) bool {
			return slice[i].Height >= syncHeight
		})
	if idx == len(slice) {
		return nil
	}
	return slice[idx:]
}

func (eon *Eon) GetPolyEvals(syncHeight int64) []shutterevents.PolyEval {
	slice := eon.PolyEvals
	idx := sort.Search(len(slice),
		func(i int) bool {
			return slice[i].Height >= syncHeight
		})
	if idx == len(slice) {
		return nil
	}
	return slice[idx:]
}

func (eon *Eon) GetAccusations(syncHeight int64) []shutterevents.Accusation {
	slice := eon.Accusations
	idx := sort.Search(len(slice),
		func(i int) bool {
			return slice[i].Height >= syncHeight
		})
	if idx == len(slice) {
		return nil
	}
	return slice[idx:]
}

func (eon *Eon) GetApologies(syncHeight int64) []shutterevents.Apology {
	slice := eon.Apologies
	idx := sort.Search(len(slice),
		func(i int) bool {
			return slice[i].Height >= syncHeight
		})
	if idx == len(slice) {
		return nil
	}
	return slice[idx:]
}

func (eon *Eon) GetEpochSecretKeyShares(syncHeight int64) []shutterevents.EpochSecretKeyShare {
	slice := eon.EpochSecretKeyShares
	idx := sort.Search(len(slice),
		func(i int) bool {
			return slice[i].Height >= syncHeight
		})
	if idx == len(slice) {
		return nil
	}
	return slice[idx:]
}

type BatchData struct {
	BatchIndex           uint64
	DecryptionSignatures []shutterevents.DecryptionSignature
}

// filterSyncHeight removes events from shutter.Eons that were generated at a height below the
// Filter's SyncHeight.
func (shutter *Shutter) filterSyncHeight() {
	syncHeight := shutter.Filter.SyncHeight
	newEons := make([]Eon, len(shutter.Eons))
	for i := range shutter.Eons {
		newEons[i] = *shutter.Eons[i].ApplyFilter(syncHeight)
	}
	shutter.Eons = newEons
}

func (shutter *Shutter) filterBatchIndex() {
	batchIndex := shutter.Filter.BatchIndex
	newBatches := make(map[uint64]*BatchData)
	for b, bd := range shutter.Batches {
		if b >= batchIndex {
			newBatches[b] = bd
		}
	}
	shutter.Batches = newBatches
}

// ApplyFilter applies the given filter and returns a new shutter object with the filter applied.
func (shutter *Shutter) ApplyFilter(newFilter ShutterFilter) *Shutter {
	if !shutter.Filter.NeedsUpdate(newFilter) {
		return shutter
	}
	clone := *shutter
	clone.Filter = newFilter
	clone.filterSyncHeight()
	clone.filterBatchIndex()
	return &clone
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

func (shutter *Shutter) fetchAndApplyEvents(ctx context.Context, shmcl client.Client, targetHeight int64) (*Shutter, error) {
	if targetHeight < shutter.CurrentBlock {
		panic("internal error: fetchAndApplyEvents bad arguments")
	}
	cloned := false

	currentBlock := shutter.CurrentBlock
	const perQuery = 500
	logProgress := shutter.CurrentBlock+perQuery < targetHeight

	for {
		height := currentBlock + perQuery
		if height > targetHeight {
			height = targetHeight
		}

		query := fmt.Sprintf("tx.height >= %d and tx.height <= %d", currentBlock+1, height)
		if logProgress {
			log.Printf("fetchAndApplyEvents: query=%s targetHeight=%d", query, targetHeight)
		}

		// tendermint silently caps the perPage value at 100, make sure to stay below, otherwise
		// our exit condition is wrong and the log.Fatalf will trigger a panic below; see
		// https://github.com/shutter-network/shutter/issues/50
		perPage := 100
		page := 1
		total := 0
		for {
			res, err := shmcl.TxSearch(ctx, query, false, &page, &perPage, "")
			if err != nil {
				return nil, pkgErrors.Wrap(err, "failed to fetch shuttermint txs")
			}
			// Create a shallow or deep clone
			if !cloned {
				if res.TotalCount == 0 && height == targetHeight {
					return shutter.ShallowClone(), nil
				}
				shutter = shutter.Clone()
				cloned = true
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
		if height == targetHeight {
			break
		}
		currentBlock = height
	}

	return shutter, nil
}

// IsCheckedIn checks if the given address sent it's check-in message.
func (shutter *Shutter) IsCheckedIn(addr common.Address) bool {
	_, ok := shutter.KeyperEncryptionKeys[addr]
	return ok
}

// IsKeyper checks if the given address is a keyper in any of the given configs.
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
	for i := len(shutter.BatchConfigs) - 1; i >= 0; i-- {
		if shutter.BatchConfigs[i].StartBatchIndex <= batchIndex {
			return shutter.BatchConfigs[i]
		}
	}
	return shutterevents.BatchConfig{}
}

func (shutter *Shutter) ShallowClone() *Shutter {
	s := *shutter
	return &s
}

func (shutter *Shutter) Clone() *Shutter {
	clone := new(Shutter)
	medley.CloneWithGob(shutter, clone)
	return clone
}

func (shutter *Shutter) GetLastCommittedHeight(ctx context.Context, shmcl client.Client) (int64, error) {
	latestBlock, err := shmcl.Block(ctx, nil)
	if err != nil {
		return 0, pkgErrors.Wrap(err, "failed to get last committed height of shuttermint chain")
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
	height, err := shutter.GetLastCommittedHeight(ctx, shmcl)
	if err != nil {
		return nil, err
	}
	if height == shutter.CurrentBlock {
		return shutter, nil
	}
	return shutter.SyncToHeight(ctx, shmcl, height)
}

// SyncToHeight syncs the state with the remote state until the given height.
func (shutter *Shutter) SyncToHeight(ctx context.Context, shmcl client.Client, height int64) (*Shutter, error) {
	nodeStatus, err := shmcl.Status(ctx)
	if err != nil {
		return nil, pkgErrors.Wrap(err, "failed to get shuttermint status")
	}

	lastCommittedHeight, err := shutter.GetLastCommittedHeight(ctx, shmcl)
	if err != nil {
		return nil, err
	}

	clone, err := shutter.fetchAndApplyEvents(ctx, shmcl, height)
	if err != nil {
		return nil, err
	}

	clone.CurrentBlock = height
	clone.LastCommittedHeight = lastCommittedHeight
	clone.NodeStatus = nodeStatus

	return clone, nil
}

// IsSynced checks if the shuttermint node is synced with the network.
func (shutter *Shutter) IsSynced() bool {
	return shutter.NodeStatus == nil || !shutter.NodeStatus.SyncInfo.CatchingUp
}

// SyncShutter subscribes to new blocks and syncs the shutter object with the head block in a
// loop. It writes newly synced shutter objects to the shutters channel, as well as errors to the
// syncErrors channel.
func SyncShutter(ctx context.Context, shmcl client.Client, shutter *Shutter, shutters chan<- *Shutter, filter <-chan ShutterFilter) error {
	name := "keyper"
	query := "tm.event = 'NewBlock'"
	events, err := shmcl.Subscribe(ctx, name, query)
	if err != nil {
		return err
	}

	reconnect := func() {
		for {
			log.Println("Attempting reconnection to Shuttermint")

			ctx2, cancel2 := context.WithTimeout(ctx, shutterReconnectInterval)
			events, err = shmcl.Subscribe(ctx2, name, query)
			cancel2()

			if err != nil {
				// try again, unless context is canceled
				select {
				case <-ctx.Done():
					return
				default:
					continue
				}
			} else {
				log.Println("Shuttermint connection regained")
				return
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case f := <-filter:
			shutter = shutter.ApplyFilter(f)
		case <-events:
			newShutter, err := shutter.SyncToHead(ctx, shmcl)
			if err != nil {
				if err != context.Canceled {
					log.Printf("Error in Shutter.SyncToHead: %+v", err)
				}
			} else {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case shutters <- newShutter:
				}

				shutter = newShutter
			}
		case <-time.After(shuttermintTimeout):
			log.Println("No Shuttermint blocks received in a long time")
			reconnect()
		}
	}
}
