package puredkg

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

func TestDealing(t *testing.T) {
	eon := 1
	numKeypers := 5
	threshold := uint64(3)

	pure := NewPureDKG(uint64(eon), uint64(numKeypers), threshold, 1)
	polyCommitment, polyEval, err := pure.StartPhase1Dealing()
	require.Nil(t, err)
	require.Equal(t, numKeypers, len(polyEval))

	for i := 0; i < numKeypers; i++ {
		require.NotNil(t, polyEval[i].Eval)
		require.NotNil(t, polyCommitment.Gammas)
		pretty.Println("Degree", polyCommitment.Gammas.Degree())
		require.Truef(t,
			crypto.VerifyPolyEval(i, polyEval[i].Eval, polyCommitment.Gammas, threshold+1),
			// XXX the paper uses a different threshold off by one
			"Commitment does not verify i=%d", i,
		)
	}
}
