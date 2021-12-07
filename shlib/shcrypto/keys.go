package shcrypto

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// EonSecretKeyShare represents a share of the eon secret key.
type EonSecretKeyShare big.Int

// EonPublicKeyShare represents a share of the eon public key.
type EonPublicKeyShare bn256.G2

// EonPublicKey represents the combined eon public key.
type EonPublicKey bn256.G2

// EpochID is the identifier of an epoch.
type EpochID bn256.G1

// EpochSecretKeyShare represents a keyper's share of the epoch sk key.
type EpochSecretKeyShare bn256.G1

// EpochSecretKey represents an epoch secret key.
type EpochSecretKey bn256.G1

func (eonPublicKey *EonPublicKey) GobEncode() ([]byte, error) {
	return eonPublicKey.Marshal(), nil
}

func (eonPublicKey *EonPublicKey) GobDecode(data []byte) error {
	return eonPublicKey.Unmarshal(data)
}

func (eonPublicKey *EonPublicKey) Equal(pk2 *EonPublicKey) bool {
	return EqualG2((*bn256.G2)(eonPublicKey), (*bn256.G2)(pk2))
}

func (eonPublicKeyShare *EonPublicKeyShare) GobEncode() ([]byte, error) {
	return eonPublicKeyShare.Marshal(), nil
}

func (eonPublicKeyShare *EonPublicKeyShare) GobDecode(data []byte) error {
	return eonPublicKeyShare.Unmarshal(data)
}

func (eonPublicKeyShare *EonPublicKeyShare) Equal(pk2 *EonPublicKeyShare) bool {
	return EqualG2((*bn256.G2)(eonPublicKeyShare), (*bn256.G2)(pk2))
}

func (epochID *EpochID) GobEncode() ([]byte, error) {
	return epochID.Marshal(), nil
}

func (epochID *EpochID) GobDecode(data []byte) error {
	return epochID.Unmarshal(data)
}

func (epochID *EpochID) Equal(g2 *EpochID) bool {
	return EqualG1((*bn256.G1)(epochID), (*bn256.G1)(g2))
}

func (epochSecretKeyShare *EpochSecretKeyShare) GobEncode() ([]byte, error) {
	return epochSecretKeyShare.Marshal(), nil
}

func (epochSecretKeyShare *EpochSecretKeyShare) GobDecode(data []byte) error {
	return epochSecretKeyShare.Unmarshal(data)
}

func (epochSecretKeyShare *EpochSecretKeyShare) Equal(g2 *EpochSecretKeyShare) bool {
	return EqualG1((*bn256.G1)(epochSecretKeyShare), (*bn256.G1)(g2))
}

func (epochSecretKey *EpochSecretKey) GobEncode() ([]byte, error) {
	return epochSecretKey.Marshal(), nil
}

func (epochSecretKey *EpochSecretKey) GobDecode(data []byte) error {
	return epochSecretKey.Unmarshal(data)
}

func (epochSecretKey *EpochSecretKey) Equal(g2 *EpochSecretKey) bool {
	return EqualG1((*bn256.G1)(epochSecretKey), (*bn256.G1)(g2))
}

func (eonSecretKeyShare *EonSecretKeyShare) GobEncode() ([]byte, error) {
	return (*big.Int)(eonSecretKeyShare).GobEncode()
}

func (eonSecretKeyShare *EonSecretKeyShare) GobDecode(data []byte) error {
	return (*big.Int)(eonSecretKeyShare).GobDecode(data)
}

func (eonSecretKeyShare *EonSecretKeyShare) Equal(e2 *EonSecretKeyShare) bool {
	return (*big.Int)(eonSecretKeyShare).Cmp((*big.Int)(e2)) == 0
}

// ComputeEonSecretKeyShare computes the keyper's secret key share from the set of poly evals
// received from the other keypers.
func ComputeEonSecretKeyShare(polyEvals []*big.Int) *EonSecretKeyShare {
	res := big.NewInt(0)
	for _, si := range polyEvals {
		res.Add(res, si)
		res.Mod(res, bn256.Order)
	}
	share := EonSecretKeyShare(*res)
	return &share
}

// ComputeEonPublicKeyShare computes the eon public key share of the given keyper.
func ComputeEonPublicKeyShare(keyperIndex int, gammas []*Gammas) *EonPublicKeyShare {
	g2 := new(bn256.G2).Set(zeroG2)
	keyperX := KeyperX(keyperIndex)
	for _, gs := range gammas {
		pi := gs.Pi(keyperX)
		g2 = new(bn256.G2).Add(g2, pi)
	}
	epk := EonPublicKeyShare(*g2)
	return &epk
}

// ComputeEonPublicKey computes the combined eon public key from the set of eon public key shares.
func ComputeEonPublicKey(gammas []*Gammas) *EonPublicKey {
	g2 := new(bn256.G2).Set(zeroG2)
	for _, gs := range gammas {
		pi := gs.Pi(big.NewInt(0))
		g2 = new(bn256.G2).Add(g2, pi)
	}
	epk := EonPublicKey(*g2)
	return &epk
}

// ComputeEpochSecretKeyShare computes a keyper's epoch sk share.
func ComputeEpochSecretKeyShare(eonSecretKeyShare *EonSecretKeyShare, epochID *EpochID) *EpochSecretKeyShare {
	g1 := new(bn256.G1).ScalarMult((*bn256.G1)(epochID), (*big.Int)(eonSecretKeyShare))
	epochSecretKeyShare := EpochSecretKeyShare(*g1)
	return &epochSecretKeyShare
}

// ComputeEpochID computes the id of the given epoch.
func ComputeEpochID(epochIndex uint64) *EpochID {
	epochIndexBig := new(big.Int).SetUint64(epochIndex + 1)
	id := EpochID(*new(bn256.G1).ScalarBaseMult(epochIndexBig))
	return &id
}

// ComputeEpochSecretKey computes the epoch secret key from a set of shares.
func ComputeEpochSecretKey(keyperIndices []int, epochSecretKeyShares []*EpochSecretKeyShare, threshold uint64) (*EpochSecretKey, error) {
	if len(keyperIndices) != len(epochSecretKeyShares) {
		return nil, fmt.Errorf("got %d keyper indices, but %d secret shares", len(keyperIndices), len(epochSecretKeyShares))
	}
	if uint64(len(keyperIndices)) != threshold {
		return nil, fmt.Errorf("got %d shares, but threshold is %d", len(keyperIndices), threshold)
	}

	skG1 := new(bn256.G1).Set(zeroG1)
	for i := 0; i < len(keyperIndices); i++ {
		keyperIndex := keyperIndices[i]
		share := epochSecretKeyShares[i]

		lambda := lagrangeCoefficient(keyperIndex, keyperIndices)
		qTimesLambda := new(bn256.G1).ScalarMult((*bn256.G1)(share), lambda)
		skG1 = new(bn256.G1).Add(skG1, qTimesLambda)
	}
	sk := EpochSecretKey(*skG1)
	return &sk, nil
}

// VerifyEpochSecretKeyShare checks that an epoch sk share published by a keyper is correct.
func VerifyEpochSecretKeyShare(epochSecretKeyShare *EpochSecretKeyShare, eonPublicKeyShare *EonPublicKeyShare, epochID *EpochID) bool {
	g1s := []*bn256.G1{
		(*bn256.G1)(epochSecretKeyShare),
		new(bn256.G1).Neg((*bn256.G1)(epochID)),
	}
	g2s := []*bn256.G2{
		new(bn256.G2).ScalarBaseMult(big.NewInt(1)),
		(*bn256.G2)(eonPublicKeyShare),
	}
	return bn256.PairingCheck(g1s, g2s)
}

// VerifyEpochSecretKey checks that an epoch secret key is the correct key for an epoch given the
// eon public key.
func VerifyEpochSecretKey(epochSecretKey *EpochSecretKey, eonPublicKey *EonPublicKey, epochIndex uint64) (bool, error) {
	sigma, err := RandomSigma(rand.Reader)
	if err != nil {
		return false, err
	}
	message := make([]byte, 32)
	_, err = rand.Read(message)
	if err != nil {
		return false, err
	}
	epochID := ComputeEpochID(epochIndex)
	encryptedMessage := Encrypt(message, eonPublicKey, epochID, sigma)
	decryptedMessage, err := encryptedMessage.Decrypt(epochSecretKey)
	if err != nil {
		return false, nil
	}
	return bytes.Equal(decryptedMessage, message), nil
}

func lagrangeCoefficientFactor(k int, keyperIndex int) *big.Int {
	xj := KeyperX(keyperIndex)
	xk := KeyperX(k)
	dx := new(big.Int).Sub(xk, xj)
	dx.Mod(dx, bn256.Order)
	dxInv := invert(dx)
	lambdaK := new(big.Int).Mul(xk, dxInv)
	lambdaK.Mod(lambdaK, bn256.Order)
	return lambdaK
}

func lagrangeCoefficient(keyperIndex int, keyperIndices []int) *big.Int {
	lambda := big.NewInt(1)
	for _, k := range keyperIndices {
		if k == keyperIndex {
			continue
		}
		lambdaK := lagrangeCoefficientFactor(k, keyperIndex)
		lambda.Mul(lambda, lambdaK)
		lambda.Mod(lambda, bn256.Order)
	}
	return lambda
}

func invert(x *big.Int) *big.Int {
	orderMinus2 := new(big.Int).Sub(bn256.Order, big.NewInt(2))
	return new(big.Int).Exp(x, orderMinus2, bn256.Order)
}
