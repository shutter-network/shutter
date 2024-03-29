package app

import (
	"fmt"
	"testing"

	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	"gotest.tools/v3/assert"
)

func makeKey(n int) []byte {
	return []byte(fmt.Sprintf("%32d", n))
}

func newpk(n int) ValidatorPubkey {
	res, err := NewValidatorPubkey(makeKey(n))
	if err != nil {
		panic(err)
	}
	return res
}

func TestMakePowermapEmpty(t *testing.T) {
	pm, err := MakePowermap([]abcitypes.ValidatorUpdate{})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(pm))
}

func TestMakePowermapBadType(t *testing.T) {
	_, err := MakePowermap([]abcitypes.ValidatorUpdate{
		{
			Power:  10,
			PubKey: tmcrypto.PublicKey{Sum: &tmcrypto.PublicKey_Ed25519{Ed25519: []byte("xxx")}},
		},
	})
	assert.Assert(t, err != nil)
}

func TestMakePowermap(t *testing.T) {
	var power int64 = 42
	pm, err := MakePowermap([]abcitypes.ValidatorUpdate{
		{Power: power, PubKey: tmcrypto.PublicKey{Sum: &tmcrypto.PublicKey_Ed25519{Ed25519: makeKey(1)}}},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(pm))
	assert.Equal(t, power, pm[newpk(1)])
}

func TestDiffPowermaps(t *testing.T) {
	oldpm := make(Powermap)
	for i := 0; i < 4; i++ {
		oldpm[newpk(i)] = int64(i)
	}

	newpm := make(Powermap)
	// drop 0,1
	newpm[newpk(2)] = oldpm[newpk(2)] // keep 2
	newpm[newpk(3)] = 15              // change 3
	newpm[newpk(4)] = 20              // add 4

	assert.Equal(t, 0, len(DiffPowermaps(oldpm, oldpm)))

	diff := DiffPowermaps(oldpm, newpm)

	expected := Powermap{
		newpk(0): 0,
		newpk(1): 0,
		newpk(3): 15,
		newpk(4): 20,
	}
	fmt.Println("DIFF", diff)

	assert.DeepEqual(t, expected, diff)
}
