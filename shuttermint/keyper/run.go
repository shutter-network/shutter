package keyper

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/ethereum/go-ethereum/crypto"
)

func SleepUntil(t time.Time) {
	now := time.Now()
	time.Sleep(t.Sub(now))
}

func NewPublicKeyCommitment(batchIndex uint64, privkey *ecdsa.PrivateKey) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_PublicKeyCommitment{
			PublicKeyCommitment: &shmsg.PublicKeyCommitment{
				BatchIndex: batchIndex,
				Commitment: crypto.FromECDSAPub(&privkey.PublicKey),
			},
		},
	}
}

func NewSecretShare(batchIndex uint64, privkey *ecdsa.PrivateKey) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_SecretShare{
			SecretShare: &shmsg.SecretShare{
				BatchIndex: batchIndex,
				Privkey:    crypto.FromECDSA(privkey),
			},
		},
	}
}

func Run(params BatchParams) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	// Wait for the start time
	SleepUntil(params.PublicKeyGenerationStartTime)
	fmt.Println("Starting key generation process", params)
	msg := NewPublicKeyCommitment(params.BatchIndex, key)
	fmt.Println("Generated pubkey", params)
	// XXX Send message
	_ = msg

	SleepUntil(params.PrivateKeyGenerationStartTime)
	msg = NewSecretShare(params.BatchIndex, key)
	fmt.Println("Generated privkey", params)
	// XXX Send message
	_ = msg
}
