package app

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

// MaxTxsPerBlock is the maximum number of txs by a single sender per block.
const MaxTxsPerBlock = 10

// NewCheckTxState returns a new check tx state.
func NewCheckTxState() *CheckTxState {
	s := CheckTxState{}
	s.Reset()
	return &s
}

// Reset should be called at every block commit so that all nodes get the same value no matter
// their view on the network.
func (s *CheckTxState) Reset() {
	s.TxCounts = make(map[common.Address]int)
	s.NonceTracker = NewNonceTracker()
}

// SetMembers sets the member set allowed to send txs. Duplicate addresses are ignored.
func (s *CheckTxState) SetMembers(members []common.Address) {
	s.Members = make(map[common.Address]bool)
	for _, address := range members {
		s.Members[address] = true
	}
}

// AddTx checks if a tx can be added and updates the internal state accordingly.
// Returns true if the sender is a member (or the member set is empty) and has not exceeded their
// tx limit yet.
func (s *CheckTxState) AddTx(sender common.Address, msg *shmsg.MessageWithNonce) bool {
	if len(s.Members) > 0 && !s.Members[sender] {
		return false
	}
	if s.TxCounts[sender] >= MaxTxsPerBlock {
		return false
	}
	if !s.NonceTracker.Check(sender, msg.RandomNonce) {
		return false
	}

	s.TxCounts[sender]++
	s.NonceTracker.Add(sender, msg.RandomNonce)
	return true
}
