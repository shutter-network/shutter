package fx

import (
	"encoding/gob"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// ActionID identifies an action.
type ActionID uint64

// PendingActions contains information about the actions, which are currently running or are
// scheduled to be run. This struct is stored on disk with gob. We enumerate the actions the
// decider gives us.
type PendingActions struct {
	mux               sync.Mutex
	ActionMap         map[ActionID]IAction
	MainChainTXHashes map[ActionID]common.Hash
	CurrentID         ActionID
	path              string
}

// NewPendingActions creates a empty PendingActions struct.
func NewPendingActions(path string) *PendingActions {
	return &PendingActions{
		ActionMap:         make(map[ActionID]IAction),
		MainChainTXHashes: make(map[ActionID]common.Hash),
		CurrentID:         0,
		path:              path,
	}
}

// SortedIDs returns the sorted pending action ids.
func (pending *PendingActions) SortedIDs() []ActionID {
	pending.mux.Lock()
	defer pending.mux.Unlock()

	var ids []ActionID
	for id := range pending.ActionMap {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

// AddActions adds the given actions unless they have already been added. id is the ActionID of the
// first action. It returns a startID, endID tuple of actions to be scheduled.
func (pending *PendingActions) AddActions(id ActionID, actions []IAction) (ActionID, ActionID) {
	pending.mux.Lock()
	defer pending.mux.Unlock()
	if id > pending.CurrentID {
		panic("internal error: AddAction called with wrong id")
	}

	skip := pending.CurrentID - id

	actions = actions[skip:]

	startID := pending.CurrentID

	for i, act := range actions {
		pending.ActionMap[startID+ActionID(i)] = act
	}
	pending.CurrentID += ActionID(len(actions))
	pending.save()
	return startID, pending.CurrentID
}

// SetMainChainTXHash sets the transaction hash for the given main chain action.
func (pending *PendingActions) SetMainChainTXHash(id ActionID, hash common.Hash) {
	pending.mux.Lock()
	defer pending.mux.Unlock()
	pending.MainChainTXHashes[id] = hash
	pending.save()
}

// GetMainChainTXHash returns the transaction hash for the given main chain action.
func (pending *PendingActions) GetMainChainTXHash(id ActionID) common.Hash {
	pending.mux.Lock()
	defer pending.mux.Unlock()
	return pending.MainChainTXHashes[id]
}

// RemoveAction removes the action with the given id.
func (pending *PendingActions) RemoveAction(id ActionID) {
	pending.mux.Lock()
	defer pending.mux.Unlock()

	delete(pending.ActionMap, id)
	delete(pending.MainChainTXHashes, id)
	pending.save()
}

// GetAction returns the action with the given id.
func (pending *PendingActions) GetAction(id ActionID) IAction {
	pending.mux.Lock()
	defer pending.mux.Unlock()
	return pending.ActionMap[id]
}

// save saves the pending actions to disk. It panics if it cannot write the file to disk.
func (pending *PendingActions) save() {
	tmppath := pending.path + ".tmp"
	file, err := os.Create(tmppath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(pending)
	if err != nil {
		panic(err)
	}

	err = file.Sync()
	if err != nil {
		panic(err)
	}
	err = os.Rename(tmppath, pending.path)
	if err != nil {
		panic(err)
	}
}

// Load loads pending actions from disk.
func (pending *PendingActions) Load() error {
	pending.mux.Lock()
	defer pending.mux.Unlock()
	gobfile, err := os.Open(pending.path)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	defer gobfile.Close()
	dec := gob.NewDecoder(gobfile)
	err = dec.Decode(pending)
	if err != nil {
		return err
	}
	log.Printf("Loaded %d pending actions from %s", len(pending.ActionMap), pending.path)
	return nil
}
