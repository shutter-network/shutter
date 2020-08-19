package shmsg

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func makeMessage() *Message {
	return &Message{
		Payload: &Message_PublicKeyCommitment{
			PublicKeyCommitment: &PublicKeyCommitment{
				BatchIndex: 1,
				Commitment: []byte("foobar"),
			},
		},
	}
}

func TestEncodeDecode(t *testing.T) {
	encoded, err := UrlEncodeMessage(makeMessage())
	if err != nil {
		t.Fatalf("Got error while encoding: %s", err)

	}
	t.Logf("Encoded: %s", encoded)
	msg, err := UrlDecodeMessage(encoded)
	if err != nil {
		t.Fatalf("Got error while decoding: %s", err)
	}
	t.Logf("decoded share=%+v", msg.GetPublicKeyShare())
	t.Logf("decoded commitment: %+v", msg.GetPublicKeyCommitment())

	if msg.GetPublicKeyCommitment() == nil {
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
	if msg.GetPublicKeyCommitment() == nil {
		t.Fatal("got no public key commitment")

	}
}
