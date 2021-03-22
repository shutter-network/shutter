// Sandbox just contains example code how to use certain libraries
package sandbox

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

// TestSigning shows how to sign messages and check signatures.
func TestSigning(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("fatal: %s", err)
	}

	msg := []byte("message to be signed")

	hash := sha3.New256()
	_, err = hash.Write(msg)
	require.Nil(t, err)
	h := hash.Sum(nil)

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	t.Logf("generated key for address %s", address.Hex())
	signature, err := crypto.Sign(h, privateKey)
	if err != nil {
		t.Fatalf("error %s %s", signature, err)
	} else {
		t.Logf("signature: %s", hexutil.Encode(signature))
	}

	// Now check the signature

	pubkey, err := crypto.SigToPub(h, signature)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	signerAddress := crypto.PubkeyToAddress(*pubkey)
	t.Logf("signer %s", signerAddress.Hex())
	if signerAddress != address {
		t.Fatalf("addresses to not match")
	}
}
