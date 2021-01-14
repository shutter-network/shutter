package shutterevents

import (
	"github.com/ethereum/go-ethereum/common"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/shcrypto"
)

type _evattr struct{}

var evattr = _evattr{}

func (_evattr) Sender(value common.Address) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Sender"),
		Value: encodeAddress(value),
		Index: true,
	}
}

func (_evattr) Eon(eon uint64) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Eon"),
		Value: encodeUint64(eon),
		Index: true,
	}
}

func (_evattr) Epoch(epoch uint64) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Epoch"),
		Value: encodeUint64(epoch),
		Index: true,
	}
}

func (_evattr) BatchIndex(batchIndex uint64) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("BatchIndex"),
		Value: encodeUint64(batchIndex),
		Index: true,
	}
}

func (_evattr) Accused(accused []common.Address) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Accused"),
		Value: encodeAddresses(accused),
	}
}

func (_evattr) Accusers(accusers []common.Address) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Accusers"),
		Value: encodeAddresses(accusers),
	}
}

func (_evattr) Receivers(receivers []common.Address) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte("Receivers"),
		Value: encodeAddresses(receivers),
	}
}

//
// Encoding/decoding helpers
//

func newByteSequencePair(key string, value [][]byte) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte(key),
		Value: encodeByteSequence(value),
	}
}

func newGammas(key string, gammas *shcrypto.Gammas) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte(key),
		Value: encodeGammas(gammas),
	}
}

func newEpochSecretKeyShare(key string, share *shcrypto.EpochSecretKeyShare) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte(key),
		Value: encodeEpochSecretKeyShare(share),
	}
}
