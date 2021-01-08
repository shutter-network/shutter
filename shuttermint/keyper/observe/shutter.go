package observe

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client"

	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
)

var errEonNotFound = errors.New("eon not found")

// Shutter let's a keyper fetch all necessary information from a shuttermint node. The only source
// for the data stored in this struct should be the shutter node.  The SyncToHead method can be
// used to update the data. All other accesses should be read-only.
type Shutter struct {
	CurrentBlock         int64
	KeyperEncryptionKeys map[common.Address]*ecies.PublicKey
	BatchConfigs         []shutterevents.BatchConfig
	Batches              map[uint64]*Batch
	Eons                 []Eon
}

// NewShutter creates an empty Shutter struct
func NewShutter() *Shutter {
	return &Shutter{
		CurrentBlock:         -1,
		KeyperEncryptionKeys: make(map[common.Address]*ecies.PublicKey),
		Batches:              make(map[uint64]*Batch),
	}
}

type Eon struct {
	Eon         uint64
	StartHeight int64
	StartEvent  shutterevents.EonStarted
	Commitments []shutterevents.PolyCommitment
	PolyEvals   []shutterevents.PolyEval
	Accusations []shutterevents.Accusation
	Apologies   []shutterevents.Apology
}

type Batch struct {
	BatchIndex           uint64
	DecryptionSignatures []shutterevents.DecryptionSignature
}

func (shutter *Shutter) applyTxEvents(height int64, events []abcitypes.Event) {
	for _, ev := range events {
		x, err := shutterevents.MakeEvent(ev, height)
		if err != nil {
			log.Printf("Error: malformed event: %s ev=%+v", err, ev)
		} else {
			shutter.applyEvent(height, x)
		}
	}
}

func (shutter *Shutter) getBatch(batchIndex uint64) *Batch {
	b, ok := shutter.Batches[batchIndex]
	if !ok {
		b = &Batch{BatchIndex: batchIndex}
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

func (shutter *Shutter) FindEon(eon uint64) (*Eon, error) {
	idx := shutter.searchEon(eon)
	if idx == len(shutter.Eons) || eon < shutter.Eons[idx].Eon {
		return nil, errEonNotFound
	}
	return &shutter.Eons[idx], nil
}

func (shutter *Shutter) applyEvent(height int64, ev shutterevents.IEvent) {
	warn := func() {
		fmt.Printf("XXX observing event not yet implemented: %s%+v\n", reflect.TypeOf(ev), ev)
	}
	switch e := ev.(type) {
	case shutterevents.CheckIn:
		shutter.KeyperEncryptionKeys[e.Sender] = e.EncryptionPublicKey
	case shutterevents.BatchConfig:
		shutter.BatchConfigs = append(shutter.BatchConfigs, e)
	case shutterevents.DecryptionSignature:
		b := shutter.getBatch(e.BatchIndex)
		b.DecryptionSignatures = append(b.DecryptionSignatures, e)
	case shutterevents.EonStarted:
		idx := shutter.searchEon(e.Eon)
		if idx < len(shutter.Eons) {
			panic("eons should increase")
		}
		shutter.Eons = append(shutter.Eons, Eon{Eon: e.Eon, StartEvent: e, StartHeight: height})
	case shutterevents.PolyCommitment:
		eon, err := shutter.FindEon(e.Eon)
		if err != nil {
			panic(err) // XXX we should remove that later
		}
		eon.Commitments = append(eon.Commitments, e)
	case shutterevents.PolyEval:
		eon, err := shutter.FindEon(e.Eon)
		if err != nil {
			panic(err) // XXX we should remove that later
		}
		eon.PolyEvals = append(eon.PolyEvals, e)
	case shutterevents.Accusation:
		eon, err := shutter.FindEon(e.Eon)
		if err != nil {
			panic(err) // XXX we should remove that later
		}
		eon.Accusations = append(eon.Accusations, e)
	case shutterevents.Apology:
		eon, err := shutter.FindEon(e.Eon)
		if err != nil {
			panic(err) // XXX we should remove that later
		}
		eon.Apologies = append(eon.Apologies, e)

	default:
		warn()
		panic("applyEvent: unknown event. giving up")
	}
}

func (shutter *Shutter) fetchAndApplyEvents(ctx context.Context, shmcl client.Client, targetHeight int64) error {
	if targetHeight < shutter.CurrentBlock {
		panic("internal error: fetchAndApplyEvents bad arguments")
	}
	query := fmt.Sprintf("tx.height >= %d and tx.height<=%d", shutter.CurrentBlock+1, targetHeight)

	page := 1
	perPage := 200
	for {
		res, err := shmcl.TxSearch(ctx, query, false, &page, &perPage, "")
		if err != nil {
			return err
		}
		for _, tx := range res.Txs {
			events := tx.TxResult.GetEvents()
			shutter.applyTxEvents(tx.Height, events)
		}
		if page*perPage > res.TotalCount {
			break
		}
		page++
	}
	return nil
}

// IsCheckedIn checks if the given address sent it's checkin message
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

func (shutter *Shutter) FindBatchConfigByBatchIndex(batchIndex uint64) shutterevents.BatchConfig {
	for i := len(shutter.BatchConfigs); i > 0; i++ {
		if shutter.BatchConfigs[i-1].StartBatchIndex <= batchIndex {
			return shutter.BatchConfigs[i-1]
		}
	}
	return shutterevents.BatchConfig{}
}

// SyncToHead syncs the state with the remote state. It fetches events from new blocks since the
// last sync and updates the state by calling applyEvent for each event.
// XXX this mutates the object in place. we may want to control mutation of the Shutter struct.
func (shutter *Shutter) SyncToHead(ctx context.Context, shmcl client.Client) error {
	latestBlock, err := shmcl.Block(ctx, nil)
	if err != nil {
		return err
	}

	err = shutter.fetchAndApplyEvents(ctx, shmcl, latestBlock.Block.Header.Height)
	if err != nil {
		return err
	}
	shutter.CurrentBlock = latestBlock.Block.Header.Height
	return nil
}
