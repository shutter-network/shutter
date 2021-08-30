package epochkg

import (
	"reflect"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/puredkg"
	"github.com/shutter-network/shutter/shlib/shtest"
)

func Results(t *testing.T) []*puredkg.Result {
	t.Helper()
	eon := uint64(5)
	numKeypers := uint64(3)
	threshold := uint64(2)

	dkgs := []*puredkg.PureDKG{}
	for i := uint64(0); i < numKeypers; i++ {
		dkg := puredkg.NewPureDKG(eon, numKeypers, threshold, i)
		dkgs = append(dkgs, &dkg)
	}

	// dealing phase
	for _, dkg := range dkgs {
		polyCommitmentMsg, polyEvalMsgs, err := dkg.StartPhase1Dealing()
		assert.NilError(t, err)

		for _, receiverDKG := range dkgs {
			err := receiverDKG.HandlePolyCommitmentMsg(polyCommitmentMsg)
			assert.NilError(t, err)
		}

		assert.Equal(t, int(numKeypers), 1+len(polyEvalMsgs))
		for _, msg := range polyEvalMsgs {
			err := dkgs[msg.Receiver].HandlePolyEvalMsg(msg)
			assert.NilError(t, err)
		}
	}

	// accusation phase
	for _, dkg := range dkgs {
		accusations := dkg.StartPhase2Accusing()
		assert.Assert(t, len(accusations) == 0)
	}

	// apology phase
	for _, dkg := range dkgs {
		apologies := dkg.StartPhase3Apologizing()
		assert.Assert(t, len(apologies) == 0)
	}

	// finalize
	for _, dkg := range dkgs {
		dkg.Finalize()
	}

	var results []*puredkg.Result

	for _, dkg := range dkgs {
		result, err := dkg.ComputeResult()
		assert.NilError(t, err)
		results = append(results, &result)
	}
	for _, r := range results {
		assert.Assert(t, reflect.DeepEqual(r.PublicKey, results[0].PublicKey))
	}
	return results
}

// TestEpochKG tests the happy case: everyone online and honest.
func TestEpochKG(t *testing.T) {
	results := Results(t)
	var kgs []*EpochKG
	for _, r := range results {
		kgs = append(kgs, NewEpochKG(r))
	}

	epoch := uint64(50)
	for sender, kg := range kgs {
		share := EpochSecretKeyShare{
			Eon:    kg.Eon,
			Epoch:  epoch,
			Sender: uint64(sender),
			Share:  kg.ComputeEpochSecretKeyShare(epoch),
		}
		for _, k := range kgs {
			err := k.HandleEpochSecretKeyShare(&share)
			assert.NilError(t, err)
		}
	}

	// every EpochKG should end up with the same key
	for _, kg := range kgs {
		_, ok := kg.SecretShares[epoch]
		assert.Assert(t, !ok)
		key, ok := kg.SecretKeys[epoch]
		assert.Assert(t, ok)
		assert.Assert(t, key != nil)
		assert.DeepEqual(t, kgs[0].SecretKeys[epoch], key)
	}

	shtest.EnsureGobable(t, kgs[0], new(EpochKG))
}
