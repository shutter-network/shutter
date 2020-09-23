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
func (cfgv *ConfigVoting) AddVote(sender common.Address, batchConfig BatchConfig) error {
	_, ok := cfgv.Votes[sender]
	if ok {
		return fmt.Errorf("sender %s already voted", sender.Hex())
	}

	for i, bc := range cfgv.Candidates {
		if reflect.DeepEqual(bc, batchConfig) {
			cfgv.Votes[sender] = i
			return nil
		}
	}

	cfgv.Candidates = append(cfgv.Candidates, batchConfig)
	cfgv.Votes[sender] = len(cfgv.Candidates) - 1
	return nil
}

// Outcome checks if one of the candidates has more than numRequiredVotes.
func (cfgv *ConfigVoting) Outcome(numRequiredVotes int) (BatchConfig, bool) {
	var numVotes []int = make([]int, len(cfgv.Candidates))

	for _, vote := range cfgv.Votes {
		numVotes[vote]++
	}
	for i, v := range numVotes {
		if v >= numRequiredVotes {
			return cfgv.Candidates[i], true
		}
	}
	return BatchConfig{}, false
}
