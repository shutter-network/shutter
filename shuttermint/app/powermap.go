package app

import (
	"bytes"
	"sort"

	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

// MakePowermap creates a new Powermap with voting powers as specified in validators.
func MakePowermap(validators []abcitypes.ValidatorUpdate) (Powermap, error) {
	res := make(Powermap)
	for _, v := range validators {
		data := v.PubKey.GetEd25519()
		if data == nil {
			return res, errors.Errorf("cannot handle key %s", v.PubKey)
		}
		pk, err := NewValidatorPubkey(data)
		if err != nil {
			return res, err
		}
		res[pk] += v.Power
	}
	return res, nil
}

// SortValidators sorts a slice of ValidatorUpdates in a determistic way suitable for updating the
// validators in tendermint.
func SortValidators(validators []abcitypes.ValidatorUpdate) {
	sort.Slice(validators, func(i, j int) bool {
		return bytes.Compare(validators[i].PubKey.GetEd25519(), validators[j].PubKey.GetEd25519()) < 0
	})
}

// DiffPowermaps computes the diff to be applied by tendermint to change the old validators into
// the new validators.
func DiffPowermaps(oldpm, newpm Powermap) Powermap {
	res := make(Powermap)

	// Remove old keys
	for v := range oldpm {
		_, ok := newpm[v]
		if !ok {
			res[v] = 0
		}
	}

	// Update new keys
	for v, p := range newpm {
		if oldpm[v] != p {
			res[v] = p
		}
	}

	return res
}

// ValidatorUpdates computes a deterministic slice of ValidatorUpdate structs.
func (pm Powermap) ValidatorUpdates() []abcitypes.ValidatorUpdate {
	var res []abcitypes.ValidatorUpdate
	for k, p := range pm {
		res = append(res, abcitypes.ValidatorUpdate{
			Power:  p,
			PubKey: tmcrypto.PublicKey{Sum: &tmcrypto.PublicKey_Ed25519{Ed25519: []byte(k.Ed25519pubkey)}},
		})
	}
	SortValidators(res)
	return res
}
