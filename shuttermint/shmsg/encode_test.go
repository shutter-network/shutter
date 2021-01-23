package shmsg

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func makeMessage() *MessageWithNonce {
	msg := &Message{
		Payload: &Message_CheckIn{
			CheckIn: &CheckIn{
				ValidatorPublicKey:  bytes.Repeat([]byte("x"), 32),
				EncryptionPublicKey: bytes.Repeat([]byte("y"), 33),
			},
		},
	}
	msgWithNonce := &MessageWithNonce{
		Msg:         msg,
		RandomNonce: 123,
	}
	return msgWithNonce
}

func TestSignMessage(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	if err != nil {
		t.Fatalf("fatal: %s", err)
	}
	signedMessage, err := SignMessage(makeMessage(), privateKey)
	require.Nil(t, err)
	t.Logf("signed message size %d", len(signedMessage))
	signer, err := GetSigner(signedMessage)
	if err != nil {
		t.Fatalf("could not get signer: %s", err)
	}
	if signer != address {
		t.Fatalf("wrong signer %s, expected %s", signer, address)
	}
	msg, err := GetMessage(signedMessage)
	if err != nil {
		t.Fatalf("could not get message: %s", err)
	}
	if msg.Msg.GetCheckIn() == nil {
		t.Fatal("got no check in")
	}
}
