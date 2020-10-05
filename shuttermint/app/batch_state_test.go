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

// TestAddpublickeycommitment tests the basic functionality of AddPublicKeyCommitment
func TestAddPublicKeyCommitment(t *testing.T) {
	privkey, err := crypto.GenerateKey()
	require.Nil(t, err, "key generation failed")
	pubkey := crypto.FromECDSAPub(&privkey.PublicKey)
	batchState := BatchState{Config: &BatchConfig{Keypers: addresses[:5]}}
	t.Logf("batch: %+v", batchState)
	err = batchState.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0], Pubkey: pubkey})
	if err != nil {
		t.Fatalf("could not add public key commitment: %s", err)
	}

	if len(batchState.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchState.Commitments)
	}

	err = batchState.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0], Pubkey: pubkey})
	if err == nil {
		t.Fatalf("no error")
	}
	t.Logf("received expected error: %s", err)
	if len(batchState.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchState.Commitments)
	}

	err = batchState.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[6], Pubkey: pubkey})
	if err == nil {
		t.Fatalf("no error")
	}

	t.Logf("received expected error: %s", err)
	if len(batchState.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchState.Commitments)
	}
}

func TestAddSecretShare(t *testing.T) {
	key1, err := crypto.GenerateKey()
	require.Nil(t, err, "could not generate key")

	key2, err := crypto.GenerateKey()
	require.Nil(t, err, "could not generate key")

	batchState := BatchState{Config: &BatchConfig{Keypers: addresses[:5]}}
	t.Logf("batch: %+v", batchState)
	err = batchState.AddPublicKeyCommitment(PublicKeyCommitment{
		Sender: addresses[0],
		Pubkey: crypto.FromECDSAPub(&key1.PublicKey),
	})
	require.Nil(t, err)

	// this should fail because we didn't provide a public key
	err = batchState.AddSecretShare(SecretShare{Sender: addresses[1], Privkey: crypto.FromECDSA(key1)})
	require.NotNil(t, err, "added secret share without providing public key first")
	t.Logf("received expected error: %s", err)

	// this should fail because we use a non-matching private key
	err = batchState.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key2)})
	require.NotNil(t, err, "added secret share with non-matching key")
	t.Logf("received expected error: %s", err)

	// this should succeed
	err = batchState.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key1)})
	require.Nil(t, err, "could not add secret share: %s", err)

	err = batchState.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key1)})
	require.NotNil(t, err, "providing a secret key a second time should fail")
}

func TestAddSecretShareSetsKeys(t *testing.T) {
	var keys [5]*ecdsa.PrivateKey
	for i := 0; i < len(keys); i++ {
		k, err := crypto.GenerateKey()
		require.Nil(t, err)
		keys[i] = k
	}
	batchState := BatchState{Config: &BatchConfig{Keypers: addresses[:5], Threshold: 3}}
	for i, k := range keys {
		err := batchState.AddPublicKeyCommitment(PublicKeyCommitment{
			Sender: addresses[i],
			Pubkey: crypto.FromECDSAPub(&k.PublicKey),
		})
		require.Nil(t, err)
		if i+1 < 3 {
			require.Nil(t, batchState.PublicKey, "should not have public key yet", i)
		} else {
			require.NotNil(t, batchState.PublicKey)
			require.Equal(t, batchState.PublicKey, &keys[2].PublicKey)
		}
	}

	for i, j := range []int{2, 1, 0, 3, 4} {
		k := keys[j]
		err := batchState.AddSecretShare(SecretShare{
			Sender:  addresses[j],
			Privkey: crypto.FromECDSA(k),
		})
		require.Nil(t, err)
		if i+1 < 3 {
			require.Nil(t, batchState.PrivateKey, "should not have public key yet", i)
		} else {
			require.NotNil(t, batchState.PrivateKey)
			require.Equal(t, batchState.PrivateKey, keys[2])
		}
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
	for _, keyper := range addresses[:3] {
		commitment := PublicKeyCommitment{
			Sender: keyper,
			Pubkey: publicKey,
		}
		err = batchState.AddPublicKeyCommitment(commitment)
		require.Nil(t, err)
	}
	require.Equal(t, *batchState.PublicKey, privateKey.PublicKey)

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
	err = batchState.AddEncryptionKeyAttestation(att)
	require.Error(t, err)

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
