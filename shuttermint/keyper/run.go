package keyper

import (
	"crypto/ecdsa"
	"log"
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

func Run(params BatchParams, ms MessageSender) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	// Wait for the start time
	SleepUntil(params.PublicKeyGenerationStartTime)
	log.Print("Starting key generation process", params)
	msg := NewPublicKeyCommitment(params.BatchIndex, key)
	log.Print("Generated pubkey", params)
	err = ms.SendMessage(msg)
	if err != nil {
		log.Print("Error while trying to send message:", err)
		return
	}

	SleepUntil(params.PrivateKeyGenerationStartTime)
	msg = NewSecretShare(params.BatchIndex, key)
	log.Print("Generated privkey", params)
	err = ms.SendMessage(msg)
	if err != nil {
		log.Println("Error while trying to send message:", err)
		return
	}
}
