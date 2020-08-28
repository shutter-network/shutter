package app

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var keys [10]*ecdsa.PrivateKey
var addresses [10]common.Address

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
	batchKeys := BatchKeys{Config: &BatchConfig{Keypers: addresses[:5]}}
	t.Logf("batch: %+v", batchKeys)
	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0], Pubkey: pubkey})
	if err != nil {
		t.Fatalf("could not add public key commitment: %s", err)
	}

	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}

	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0], Pubkey: pubkey})
	if err == nil {
		t.Fatalf("no error")
	}
	t.Logf("received expected error: %s", err)
	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}

	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[6], Pubkey: pubkey})
	if err == nil {
		t.Fatalf("no error")
	}

	t.Logf("received expected error: %s", err)
	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}
}

func TestAddSecretShare(t *testing.T) {
	key1, err := crypto.GenerateKey()
	require.Nil(t, err, "could not generate key")

	key2, err := crypto.GenerateKey()
	require.Nil(t, err, "could not generate key")

	batchKeys := BatchKeys{Config: &BatchConfig{Keypers: addresses[:5]}}
	t.Logf("batch: %+v", batchKeys)
	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{
		Sender: addresses[0],
		Pubkey: crypto.FromECDSAPub(&key1.PublicKey)})
	require.Nil(t, err)

	// this should fail because we didn't provide a public key
	err = batchKeys.AddSecretShare(SecretShare{Sender: addresses[1], Privkey: crypto.FromECDSA(key1)})
	require.NotNil(t, err, "added secret share without providing public key first")
	t.Logf("received expected error: %s", err)

	// this should fail because we use a non-matching private key
	err = batchKeys.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key2)})
	require.NotNil(t, err, "added secret share with non-matching key")
	t.Logf("received expected error: %s", err)

	// this should succeed
	err = batchKeys.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key1)})
	require.Nil(t, err, "could not add secret share: %s", err)

	err = batchKeys.AddSecretShare(SecretShare{Sender: addresses[0], Privkey: crypto.FromECDSA(key1)})
	require.NotNil(t, err, "providing a secret key a second time should fail")
}

func TestAddSecretShareSetsKeys(t *testing.T) {
	var keys [5]*ecdsa.PrivateKey
	for i := 0; i < len(keys); i++ {
		k, err := crypto.GenerateKey()
		require.Nil(t, err)
		keys[i] = k
	}
	batchKeys := BatchKeys{Config: &BatchConfig{Keypers: addresses[:5], Threshhold: 3}}
	for i, k := range keys {
		err := batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{
			Sender: addresses[i],
			Pubkey: crypto.FromECDSAPub(&k.PublicKey)})
		require.Nil(t, err)
		if i+1 < 3 {
			require.Nil(t, batchKeys.PublicKey, "should not have public key yet", i)
		} else {
			require.NotNil(t, batchKeys.PublicKey)
			require.Equal(t, batchKeys.PublicKey, &keys[2].PublicKey)
		}
	}

	for i, j := range []int{2, 1, 0, 3, 4} {
		k := keys[j]
		err := batchKeys.AddSecretShare(SecretShare{
			Sender:  addresses[j],
			Privkey: crypto.FromECDSA(k)})
		require.Nil(t, err)
		if i+1 < 3 {
			require.Nil(t, batchKeys.PrivateKey, "should not have public key yet", i)
		} else {
			require.NotNil(t, batchKeys.PrivateKey)
			require.Equal(t, batchKeys.PrivateKey, keys[2])
		}
	}
}
