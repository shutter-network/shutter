package app

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	keys      [10]*ecdsa.PrivateKey
	addresses [10]common.Address
)

func init() {
	for i := 0; i < 10; i++ {
		var d [32]byte
		d[31] = byte(i + 1)
		k, err := crypto.ToECDSA(d[:])
		if err != nil {
			panic(err)
		}
		keys[i] = k
		addresses[i] = crypto.PubkeyToAddress(k.PublicKey)
	}
}

func TestAddEncryptionKeyAttestation(t *testing.T) {
	batchState := BatchState{
		Config:     &BatchConfig{Keypers: addresses[:5], Threshold: 3},
		BatchIndex: 5,
	}

	privateKey, err := crypto.GenerateKey()
	require.Nil(t, err)
	publicKey := crypto.FromECDSAPub(&privateKey.PublicKey)
	// for _, keyper := range addresses[:3] {
	//	commitment := PublicKeyCommitment{
	//		Sender: keyper,
	//		Pubkey: publicKey,
	//	}
	//	err = batchState.AddPublicKeyCommitment(commitment)
	//	require.Nil(t, err)
	// }
	// require.Equal(t, *batchState.PublicKey, privateKey.PublicKey)

	// don't accept messages from non-keypers
	nonKeyper := addresses[6]
	err = batchState.AddEncryptionKeyAttestation(EncryptionKeyAttestation{
		Sender: nonKeyper,
	})
	require.Error(t, err)

	// don't accept attestation for wrong key
	wrongKey := []byte("key")
	configContractAddress := common.HexToAddress("0x")
	preimage := EncryptionKeyPreimage(wrongKey, batchState.BatchIndex, configContractAddress)
	hash := crypto.Keccak256Hash(preimage)
	sig, err := crypto.Sign(hash.Bytes(), keys[0])
	require.Nil(t, err)
	att := EncryptionKeyAttestation{
		Sender:                addresses[0],
		EncryptionKey:         wrongKey,
		BatchIndex:            batchState.BatchIndex,
		ConfigContractAddress: configContractAddress,
		Signature:             sig,
	}
	require.True(t, att.VerifySignature())
	// err = batchState.AddEncryptionKeyAttestation(att)
	// require.Error(t, err)

	preimage = EncryptionKeyPreimage(publicKey, batchState.BatchIndex, configContractAddress)
	hash = crypto.Keccak256Hash(preimage)
	sig, err = crypto.Sign(hash.Bytes(), keys[0])
	require.Nil(t, err)

	// don't accept attestation with wrong config contract address
	err = batchState.AddEncryptionKeyAttestation(EncryptionKeyAttestation{
		Sender:                addresses[0],
		EncryptionKey:         publicKey,
		BatchIndex:            batchState.BatchIndex,
		ConfigContractAddress: addresses[6],
		Signature:             sig,
	})
	require.Error(t, err)

	// add attestation from keyper
	att = EncryptionKeyAttestation{
		Sender:                addresses[0],
		EncryptionKey:         publicKey,
		BatchIndex:            batchState.BatchIndex,
		ConfigContractAddress: configContractAddress,
		Signature:             sig,
	}
	require.True(t, att.VerifySignature())
	err = batchState.AddEncryptionKeyAttestation(att)
	require.Nil(t, err)

	// find attestation by sender
	require.True(t, len(batchState.EncryptionKeyAttestations) == 1)
	foundAtt, err := batchState.FindEncryptionKeyAttestation(addresses[0])
	require.Nil(t, err)
	require.Equal(t, foundAtt, att)

	// don't accept another attestation by keyper
	key2 := []byte("key2")
	preimage2 := EncryptionKeyPreimage(key2, batchState.BatchIndex, configContractAddress)
	hash2 := crypto.Keccak256Hash(preimage2)
	sig2, err := crypto.Sign(hash2.Bytes(), keys[0])
	require.Nil(t, err)
	att2 := EncryptionKeyAttestation{
		Sender:                addresses[0],
		EncryptionKey:         key2,
		BatchIndex:            batchState.BatchIndex,
		ConfigContractAddress: configContractAddress,
		Signature:             sig2,
	}
	err = batchState.AddEncryptionKeyAttestation(att2)
	require.Error(t, err)
}
