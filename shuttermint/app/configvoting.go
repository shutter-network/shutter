package app

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// NewConfigVoting creates a ConfigVoting struct
func NewConfigVoting() ConfigVoting {
	return ConfigVoting{
		Votes: make(map[common.Address]int),
	}
}

// AddVote adds a vote from the given sender for the given batchConfig.  It returns an error if the
// sender tries to vote twice.
func (cb *ConfigVoting) AddVote(sender common.Address, batchConfig BatchConfig) error {
	_, ok := cb.Votes[sender]
	if ok {
		return fmt.Errorf("sender %s already voted", sender.Hex())
	}

	for i, bc := range cb.Candidates {
		if reflect.DeepEqual(bc, batchConfig) {
			cb.Votes[sender] = i
			return nil
		}
	}

	cb.Candidates = append(cb.Candidates, batchConfig)
	cb.Votes[sender] = len(cb.Candidates) - 1
	return nil
}

// Outcome checks if one of the candidates has more than numRequiredVotes.
func (cb *ConfigVoting) Outcome(numRequiredVotes int) (BatchConfig, bool) {
	var numVotes []int = make([]int, len(cb.Candidates))

	for _, vote := range cb.Votes {
		numVotes[vote]++
	}
	for i, v := range numVotes {
		if v >= numRequiredVotes {
			return cb.Candidates[i], true
		}
	}
	return BatchConfig{}, false
}
