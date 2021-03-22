package shmsg

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"google.golang.org/protobuf/proto"
)

// Instead of relying on protocol buffers we simply send a signature, followed by the marshaled message

// Add a prefix to avoid accidentally signing data with special meaning in different context, in
// particular Ethereum transactions (c.f. EIP191 https://eips.ethereum.org/EIPS/eip-191).
var hashPrefix = []byte{0x19, 's', 'h', 'm', 's', 'g'}

// SignMessage signs the given Message with the given private key.
func SignMessage(msg proto.Message, privkey *ecdsa.PrivateKey) ([]byte, error) {
	marshaled, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	hash := sha3.New256()
	_, err = hash.Write(hashPrefix)
	if err != nil {
		return nil, err
	}
	_, err = hash.Write(marshaled)
	if err != nil {
		return nil, err
	}

	h := hash.Sum(nil)
	signature, err := crypto.Sign(h, privkey)
	if err != nil {
		return nil, err
	}

	return append(signature, marshaled...), nil
}

// GetSigner returns the signer address of a signed message.
func GetSigner(signedMessage []byte) (common.Address, error) {
	var signer common.Address
	if len(signedMessage) < crypto.SignatureLength {
		return signer, errors.New("message too short")
	}
	hash := sha3.New256()
	_, err := hash.Write(hashPrefix)
	if err != nil {
		return common.Address{}, err
	}

	_, err = hash.Write(signedMessage[crypto.SignatureLength:])
	if err != nil {
		return common.Address{}, err
	}
	h := hash.Sum(nil)
	pubkey, err := crypto.SigToPub(h, signedMessage[:crypto.SignatureLength])
	if err != nil {
		return signer, err
	}
	signer = crypto.PubkeyToAddress(*pubkey)
	return signer, nil
}

// GetMessage returns the unmarshalled Message of a signed message.
func GetMessage(signedMessage []byte) (*MessageWithNonce, error) {
	msg := MessageWithNonce{}
	if err := proto.Unmarshal(signedMessage[crypto.SignatureLength:], &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (m *Message) GobEncode() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Message) GobDecode(data []byte) error {
	return proto.Unmarshal(data, m)
}
