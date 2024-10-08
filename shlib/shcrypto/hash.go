package shcrypto

import (
	"math/big"

	blst "github.com/supranational/blst/bindings/go"
	"golang.org/x/crypto/sha3"
)

const HashToG1DST = "SHUTTER_V01_BLS12381G1_XMD:SHA-256_SSWU_RO_"

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

func Hash1(b []byte) *blst.P1Affine {
	bWithPrefix := append([]byte{1}, b...)
	p := blst.HashToG1(bWithPrefix, []byte(HashToG1DST))
	return p.ToAffine()
}

func Hash2(gt *blst.Fp12) Block {
	b := gt.ToBendian()
	h := hashWithPrefix(2, b)
	var block Block
	copy(block[:], h)
	return block
}

func Hash3(b []byte) *big.Int {
	h := hashWithPrefix(3, b)
	i := new(big.Int).SetBytes(h)
	i = i.Mod(i, order)
	return i
}

func Hash4(b []byte) Block {
	h := hashWithPrefix(4, b)
	var block Block
	copy(block[:], h)
	return block
}
