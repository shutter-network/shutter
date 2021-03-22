package app

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// NewConfigVoting creates a ConfigVoting struct.
func NewConfigVoting() ConfigVoting {
	return ConfigVoting{
		Voting:     NewVoting(),
		Candidates: []BatchConfig{},
	}
}

// AddVote adds a vote from the given sender for the given batchConfig.  It returns an error if the
// sender tries to vote twice.
func (cfgv *ConfigVoting) AddVote(sender common.Address, batchConfig BatchConfig) error {
	_, ok := cfgv.Votes[sender]
	if ok {
		return errors.Errorf("sender %s already voted", sender.Hex())
	}

	for i, bc := range cfgv.Candidates {
		if reflect.DeepEqual(bc, batchConfig) {
			cfgv.AddVoteForIndex(sender, i)
			return nil
		}
	}

	cfgv.Candidates = append(cfgv.Candidates, batchConfig)
	cfgv.AddVoteForIndex(sender, len(cfgv.Candidates)-1)
	return nil
}

// Outcome checks if one of the candidates has more than numRequiredVotes.
func (cfgv *ConfigVoting) Outcome(numRequiredVotes int) (BatchConfig, bool) {
	outcomeIndex, success := cfgv.OutcomeIndex(numRequiredVotes)
	if !success {
		return BatchConfig{}, false
	}
	return cfgv.Candidates[outcomeIndex], true
}
