package shmsg

import "testing"

func makeMessage() *Message {
	return &Message{
		Payload: &Message_PublicKeyCommitment{
			PublicKeyCommitment: &PublicKeyCommitment{
				BatchId:    "foo-1",
				Commitment: []byte("foobar"),
				Signature:  []byte("signature"),
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
