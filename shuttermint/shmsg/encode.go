package shmsg

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"google.golang.org/protobuf/proto"
)

// UrlEncodeMessage encodes Message as a string, which is safe to be used as part of an URL
func UrlEncodeMessage(msg *Message) (string, error) {
	out, err := proto.Marshal(msg)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(out), nil
}

// UrlDecodeMessage decodes a Message from the given string
func UrlDecodeMessage(encoded string) (*Message, error) {
	msg := Message{}
	out, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(out, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// Instead of relying on protocol buffers we simply send a signature, followed by the marshalled message

// eip191prefix is used to make sure we do not sign a valid ethereum transaction, see
// https://eips.ethereum.org/EIPS/eip-191
var eip191prefix = []byte{0x19, 's', 'h', 'm', 's', 'g'}

// SignMessage signs the given Message with the given private key
func SignMessage(msg *Message, privkey *ecdsa.PrivateKey) ([]byte, error) {
	marshalled, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	hash := sha3.New256()
	hash.Write(eip191prefix)
	hash.Write(marshalled)
	h := hash.Sum(nil)
	signature, err := crypto.Sign(h, privkey)
	if err != nil {
		return nil, err
	}

	return append(signature, marshalled...), nil
}

// GetSigner returns the signer address of a signed message
func GetSigner(signedMessage []byte) (common.Address, error) {
	var signer common.Address
	if len(signedMessage) < crypto.SignatureLength {
		return signer, errors.New("message too short")
	}
	hash := sha3.New256()
	hash.Write(eip191prefix)
	hash.Write(signedMessage[crypto.SignatureLength:])
	h := hash.Sum(nil)
	pubkey, err := crypto.SigToPub(h, signedMessage[:crypto.SignatureLength])
	if err != nil {
		return signer, err
	}
	signer = crypto.PubkeyToAddress(*pubkey)
	return signer, nil
}

// GetMessage returns the unmarshalled Message of a signed message
func GetMessage(signedMessage []byte) (*Message, error) {
	msg := Message{}
	if err := proto.Unmarshal(signedMessage[crypto.SignatureLength:], &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
