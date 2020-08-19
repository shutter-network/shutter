package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var addresses [10]common.Address

func init() {
	for i := 0; i < 10; i++ {
		addresses[i] = common.BigToAddress(big.NewInt(int64(i)))
	}
}

// TestAddpublickeycommitment tests the basic functionality of AddPublicKeyCommitment
func TestAddPublicKeyCommitment(t *testing.T) {
	batchKeys := BatchKeys{Config: &BatchConfig{Keypers: addresses[:5]}}
	t.Logf("batch: %+v", batchKeys)
	err := batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0]})
	if err != nil {
		t.Fatalf("could not add public key commitment: %s", err)
	}

	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}

	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[0]})
	if err == nil {
		t.Fatalf("no error")
	}
	t.Logf("received expected error: %s", err)
	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}

	err = batchKeys.AddPublicKeyCommitment(PublicKeyCommitment{Sender: addresses[6]})
	if err == nil {
		t.Fatalf("no error")
	}

	t.Logf("received expected error: %s", err)
	if len(batchKeys.Commitments) != 1 {
		t.Fatalf("wrong number of commitments: %s", batchKeys.Commitments)
	}
}
