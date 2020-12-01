package app

import "github.com/ethereum/go-ethereum/common"

func NewEonStartVoting() EonStartVoting {
	return EonStartVoting{
		Voting:     NewVoting(),
		Candidates: []uint64{},
	}
}

func (v *EonStartVoting) AddVote(sender common.Address, batchIndex uint64) {
	for i, b := range v.Candidates {
		if b == batchIndex {
			v.AddVoteForIndex(sender, i)
			return
		}
	}

	v.Candidates = append(v.Candidates, batchIndex)
	v.AddVoteForIndex(sender, len(v.Candidates)-1)
}

func (v *EonStartVoting) Outcome(numRequiredVotes int) (uint64, bool) {
	i, success := v.OutcomeIndex(numRequiredVotes)
	if !success {
		return 0, false
	}
	return v.Candidates[i], true
}
