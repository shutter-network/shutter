/*epochkg implements the epoch key generation given the result of a successful DKG generation with
  puredkg*/
package epochkg

import (
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shlib/puredkg"
	"github.com/shutter-network/shutter/shlib/shcrypto"
)

type (
	KeyperIndex = uint64
)

type EpochKG struct {
	Eon             uint64
	NumKeypers      uint64
	Threshold       uint64
	Keyper          KeyperIndex
	SecretKeyShare  *shcrypto.EonSecretKeyShare
	PublicKey       *shcrypto.EonPublicKey
	PublicKeyShares []*shcrypto.EonPublicKeyShare

	SecretShares map[uint64][]*EpochSecretKeyShare
	SecretKeys   map[uint64]*shcrypto.EpochSecretKey
}

type EpochSecretKeyShare struct {
	Eon    uint64
	Epoch  uint64
	Sender KeyperIndex
	Share  *shcrypto.EpochSecretKeyShare
}

func NewEpochKG(puredkgResult *puredkg.Result) *EpochKG {
	return &EpochKG{
		Eon:             puredkgResult.Eon,
		NumKeypers:      puredkgResult.NumKeypers,
		Threshold:       puredkgResult.Threshold,
		Keyper:          puredkgResult.Keyper,
		SecretKeyShare:  puredkgResult.SecretKeyShare,
		PublicKey:       puredkgResult.PublicKey,
		PublicKeyShares: puredkgResult.PublicKeyShares,

		SecretShares: make(map[uint64][]*EpochSecretKeyShare),
		SecretKeys:   make(map[uint64]*shcrypto.EpochSecretKey),
	}
}

func (epochkg *EpochKG) ComputeEpochSecretKeyShare(epoch uint64) *shcrypto.EpochSecretKeyShare {
	epochID := shcrypto.ComputeEpochID(epoch)
	return shcrypto.ComputeEpochSecretKeyShare(epochkg.SecretKeyShare, epochID)
}

func (epochkg *EpochKG) computeEpochSecretKey(shares []*EpochSecretKeyShare) (*shcrypto.EpochSecretKey, error) {
	var keyperIndices []int
	var epochSecretKeyShares []*shcrypto.EpochSecretKeyShare
	for _, s := range shares {
		keyperIndices = append(keyperIndices, int(s.Sender))
		epochSecretKeyShares = append(epochSecretKeyShares, s.Share)
	}
	return shcrypto.ComputeEpochSecretKey(keyperIndices, epochSecretKeyShares, epochkg.Threshold)
}

func (epochkg *EpochKG) addEpochSecretKeyShare(share *EpochSecretKeyShare) error {
	shares := epochkg.SecretShares[share.Epoch]
	for _, s := range shares {
		if s.Sender == share.Sender {
			return errors.Errorf(
				"already have EpochSecretKeyShare from sender %d for epoch %d",
				share.Sender,
				share.Epoch)
		}
	}
	shares = append(shares, share)
	if len(shares) != int(epochkg.Threshold) {
		epochkg.SecretShares[share.Epoch] = shares
		return nil
	}

	secretKey, err := epochkg.computeEpochSecretKey(shares)
	delete(epochkg.SecretShares, share.Epoch)
	epochkg.SecretKeys[share.Epoch] = secretKey // may be nil in the error case
	return err
}

func (epochkg *EpochKG) HandleEpochSecretKeyShare(share *EpochSecretKeyShare) error {
	if _, ok := epochkg.SecretKeys[share.Epoch]; ok {
		// We already have the key for this epoch
		return nil
	}
	epochID := shcrypto.ComputeEpochID(share.Epoch)
	if !shcrypto.VerifyEpochSecretKeyShare(
		share.Share,
		epochkg.PublicKeyShares[share.Sender],
		epochID,
	) {
		return errors.Errorf(
			"cannot verify epoch secret key share from sender %d for epoch %d",
			share.Sender,
			share.Epoch)
	}
	err := epochkg.addEpochSecretKeyShare(share)
	if err != nil {
		return err
	}

	return nil
}
