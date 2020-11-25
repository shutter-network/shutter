package shmsg

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func makeMessage() *Message {
	return &Message{
		Payload: &Message_CheckIn{
			CheckIn: &CheckIn{
				ValidatorPublicKey:  bytes.Repeat([]byte("x"), 32),
				EncryptionPublicKey: bytes.Repeat([]byte("y"), 33),
			},
		},
	}
}

func TestEncodeDecode(t *testing.T) {
	encoded, err := URLEncodeMessage(makeMessage())
	if err != nil {
		t.Fatalf("Got error while encoding: %s", err)
	}
	t.Logf("Encoded: %s", encoded)
	msg, err := URLDecodeMessage(encoded)
	if err != nil {
		t.Fatalf("Got error while decoding: %s", err)
	}
	t.Logf("decoded check in: %+v", msg.GetCheckIn())

	if msg.GetCheckIn() == nil {
		t.Fatal("got no public key commitment")
	}
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
	if msg.GetCheckIn() == nil {
		t.Fatal("got no check in")
	}
}
