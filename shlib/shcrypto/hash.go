package shcrypto

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"golang.org/x/crypto/sha3"
)

func keccak256(ds ...[]byte) []byte {
	state := sha3.NewLegacyKeccak256()
	for _, d := range ds {
		state.Write(d)
	}
	h := []byte{}
	return state.Sum(h)
}

func hashWithPrefix(p byte, b []byte) []byte {
	return keccak256(append([]byte{p}, b...))
}

// p. 7, H_1
func Hash1(b []byte) *bn256.G1 {
	h := hashWithPrefix(1, b)
	n := new(big.Int).SetBytes(h)
	p := new(bn256.G1).ScalarBaseMult(n)
	return p
}

// p. 7, H_2
func Hash2(gt *bn256.GT) Block {
	b := gt.Marshal()
	h := hashWithPrefix(2, b)
	var block Block
	copy(block[:], h)
	return block
}

// p. 7, H_3
func Hash3(b []byte) *big.Int {
	h := hashWithPrefix(3, b)
	i := new(big.Int).SetBytes(h)
	i = i.Mod(i, bn256.Order)
	return i
}

// p. 7, H_4
func Hash4(b []byte) Block {
	h := hashWithPrefix(4, b)
	var block Block
	copy(block[:], h)
	return block
}
