package app

import (
	"github.com/ethereum/go-ethereum/common"
)

// NewVoting creates a new Voting struct.
func NewVoting() Voting {
	return Voting{
		Votes: make(map[common.Address]int),
	}
}

// AddVoteForIndex adds a vote from the given sender for the given index.
func (v *Voting) AddVoteForIndex(sender common.Address, index int) {
	v.Votes[sender] = index
}

// OutcomeIndex checks if one of the candidate indices has more than numRequiredVotes.
func (v *Voting) OutcomeIndex(numRequiredVotes int) (int, bool) {
	numVotes := make(map[int]int)

	for _, vote := range v.Votes {
		numVotes[vote]++
	}
	for index, votes := range numVotes {
		if votes >= numRequiredVotes {
			return index, true
		}
	}
	return -1, false
}
