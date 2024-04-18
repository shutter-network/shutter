package shcrypto

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
)

// EonSecretKeyShare represents a share of the eon secret key.
type EonSecretKeyShare big.Int

// EonPublicKeyShare represents a share of the eon public key.
type EonPublicKeyShare bls12381.PointG2

// EonPublicKey represents the combined eon public key.
type EonPublicKey bls12381.PointG2

// EpochID is the identifier of an epoch.
type EpochID bls12381.PointG1

// EpochSecretKeyShare represents a keyper's share of the epoch sk key.
type EpochSecretKeyShare bls12381.PointG1

// EpochSecretKey represents an epoch secret key.
type EpochSecretKey bls12381.PointG1

func (eonPublicKey *EonPublicKey) GobEncode() ([]byte, error) {
	return eonPublicKey.Marshal(), nil
}

func (eonPublicKey *EonPublicKey) GobDecode(data []byte) error {
	return eonPublicKey.Unmarshal(data)
}

func (eonPublicKey *EonPublicKey) Equal(pk2 *EonPublicKey) bool {
	g2 := bls12381.NewG2()
	return g2.Equal((*bls12381.PointG2)(eonPublicKey), (*bls12381.PointG2)(pk2))
}

func (eonPublicKeyShare *EonPublicKeyShare) GobEncode() ([]byte, error) {
	return eonPublicKeyShare.Marshal(), nil
}

func (eonPublicKeyShare *EonPublicKeyShare) GobDecode(data []byte) error {
	return eonPublicKeyShare.Unmarshal(data)
}

func (eonPublicKeyShare *EonPublicKeyShare) Equal(pk2 *EonPublicKeyShare) bool {
	g2 := bls12381.NewG2()
	return g2.Equal((*bls12381.PointG2)(eonPublicKeyShare), (*bls12381.PointG2)(pk2))
}

func (epochID *EpochID) GobEncode() ([]byte, error) {
	return epochID.Marshal(), nil
}

func (epochID *EpochID) GobDecode(data []byte) error {
	return epochID.Unmarshal(data)
}

func (epochID *EpochID) Equal(g2 *EpochID) bool {
	g1 := bls12381.NewG1()
	return g1.Equal((*bls12381.PointG1)(epochID), (*bls12381.PointG1)(g2))
}

func (epochSecretKeyShare *EpochSecretKeyShare) GobEncode() ([]byte, error) {
	return epochSecretKeyShare.Marshal(), nil
}

func (epochSecretKeyShare *EpochSecretKeyShare) GobDecode(data []byte) error {
	return epochSecretKeyShare.Unmarshal(data)
}

func (epochSecretKeyShare *EpochSecretKeyShare) Equal(g2 *EpochSecretKeyShare) bool {
	g1 := bls12381.NewG1()
	return g1.Equal((*bls12381.PointG1)(epochSecretKeyShare), (*bls12381.PointG1)(g2))
}

func (epochSecretKey *EpochSecretKey) GobEncode() ([]byte, error) {
	return epochSecretKey.Marshal(), nil
}

func (epochSecretKey *EpochSecretKey) GobDecode(data []byte) error {
	return epochSecretKey.Unmarshal(data)
}

func (epochSecretKey *EpochSecretKey) Equal(g2 *EpochSecretKey) bool {
	g1 := bls12381.NewG1()
	return g1.Equal((*bls12381.PointG1)(epochSecretKey), (*bls12381.PointG1)(g2))
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
		res.Mod(res, order)
	}
	share := EonSecretKeyShare(*res)
	return &share
}

// ComputeEonPublicKeyShare computes the eon public key share of the given keyper.
func ComputeEonPublicKeyShare(keyperIndex int, gammas []*Gammas) *EonPublicKeyShare {
	g2 := bls12381.NewG2()
	p := g2.Zero()
	keyperX := KeyperX(keyperIndex)
	for _, gs := range gammas {
		pi := gs.Pi(keyperX)
		g2.Add(p, p, pi)
	}
	epk := EonPublicKeyShare(*p)
	return &epk
}

// ComputeEonPublicKey computes the combined eon public key from the set of eon public key shares.
func ComputeEonPublicKey(gammas []*Gammas) *EonPublicKey {
	g2 := bls12381.NewG2()
	p := g2.Zero()
	for _, gs := range gammas {
		pi := gs.Pi(big.NewInt(0))
		g2.Add(p, p, pi)
	}
	epk := EonPublicKey(*p)
	return &epk
}

// ComputeEpochSecretKeyShare computes a keyper's epoch sk share.
func ComputeEpochSecretKeyShare(eonSecretKeyShare *EonSecretKeyShare, epochID *EpochID) *EpochSecretKeyShare {
	g1 := bls12381.NewG1()
	p := g1.MulScalar(new(bls12381.PointG1), (*bls12381.PointG1)(epochID), (*big.Int)(eonSecretKeyShare))
	epochSecretKeyShare := EpochSecretKeyShare(*p)
	return &epochSecretKeyShare
}

// ComputeEpochID computes the id of the given epoch.
func ComputeEpochID(epochIDBytes []byte) *EpochID {
	return (*EpochID)(Hash1(epochIDBytes))
}

// LagrangeCoeffs stores the lagrange coefficients that are needed to compute an epoch secret key
// for a certain array of keypers. We use this to speedup epoch secret key generation.
type LagrangeCoeffs struct {
	lambdas []*big.Int
}

// NewLagrangeCoeffs computes the lagrange coefficients for the given array of keypers.
func NewLagrangeCoeffs(keyperIndices []int) *LagrangeCoeffs {
	lambdas := make([]*big.Int, len(keyperIndices))
	for i, keyperIndex := range keyperIndices {
		lambdas[i] = lagrangeCoefficient(keyperIndex, keyperIndices)
	}
	return &LagrangeCoeffs{
		lambdas: lambdas,
	}
}

// ComputeEpochSecretKey computes the epoch secret key given the secret key shares of the keypers.
// The caller has to ensure that the secret shares match the keyperIndices used during
// initialisation.
func (lc *LagrangeCoeffs) ComputeEpochSecretKey(epochSecretKeyShares []*EpochSecretKeyShare) (*EpochSecretKey, error) {
	if len(epochSecretKeyShares) != len(lc.lambdas) {
		return nil, fmt.Errorf("got %d shares, expected %d", len(epochSecretKeyShares), len(lc.lambdas))
	}
	g1 := bls12381.NewG1()
	skG1 := g1.Zero()
	qTimesLambda := new(bls12381.PointG1)
	for i, share := range epochSecretKeyShares {
		lambda := lc.lambdas[i]
		g1.MulScalar(qTimesLambda, (*bls12381.PointG1)(share), lambda)
		g1.Add(skG1, skG1, qTimesLambda)
	}
	return (*EpochSecretKey)(skG1), nil
}

// ComputeEpochSecretKey computes the epoch secret key from a set of shares.
func ComputeEpochSecretKey(keyperIndices []int, epochSecretKeyShares []*EpochSecretKeyShare, threshold uint64) (*EpochSecretKey, error) {
	if len(keyperIndices) != len(epochSecretKeyShares) {
		return nil, fmt.Errorf("got %d keyper indices, but %d secret shares", len(keyperIndices), len(epochSecretKeyShares))
	}
	if uint64(len(keyperIndices)) != threshold {
		return nil, fmt.Errorf("got %d shares, but threshold is %d", len(keyperIndices), threshold)
	}

	return NewLagrangeCoeffs(keyperIndices).ComputeEpochSecretKey(epochSecretKeyShares)
}

// VerifyEpochSecretKeyShare checks that an epoch sk share published by a keyper is correct.
func VerifyEpochSecretKeyShare(epochSecretKeyShare *EpochSecretKeyShare, eonPublicKeyShare *EonPublicKeyShare, epochID *EpochID) bool {
	g2 := bls12381.NewG2()
	pairingEngine := bls12381.NewPairingEngine()

	pairingEngine.AddPair((*bls12381.PointG1)(epochSecretKeyShare), g2.One())
	// pairingEngine.AddPairInv modifies the first argument, so we have to create a copy of epochID
	epochIDPoint := new(bls12381.PointG1).Set((*bls12381.PointG1)(epochID))
	pairingEngine.AddPairInv(epochIDPoint, (*bls12381.PointG2)(eonPublicKeyShare))
	return pairingEngine.Check()
}

// VerifyEpochSecretKey checks that an epoch secret key is the correct key for an epoch given the
// eon public key.
func VerifyEpochSecretKey(epochSecretKey *EpochSecretKey, eonPublicKey *EonPublicKey, epochIDBytes []byte) (bool, error) {
	sigma, err := RandomSigma(rand.Reader)
	if err != nil {
		return false, err
	}
	message := make([]byte, 32)
	_, err = rand.Read(message)
	if err != nil {
		return false, err
	}
	return VerifyEpochSecretKeyDeterministic(epochSecretKey, eonPublicKey, epochIDBytes, sigma, message)
}

// VerifyEpochSecretKeyDeterministic checks that an epoch secret key is the correct key for an
// epoch given the eon public key and random inputs for sigma and message.
func VerifyEpochSecretKeyDeterministic(epochSecretKey *EpochSecretKey, eonPublicKey *EonPublicKey, epochIDBytes []byte, sigma Block, message []byte) (bool, error) {
	epochID := ComputeEpochID(epochIDBytes)
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
	dx.Mod(dx, order)
	dxInv := invert(dx)
	lambdaK := new(big.Int).Mul(xk, dxInv)
	lambdaK.Mod(lambdaK, order)
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
		lambda.Mod(lambda, order)
	}
	return lambda
}

func invert(x *big.Int) *big.Int {
	orderMinus2 := new(big.Int).Sub(order, big.NewInt(2))
	return new(big.Int).Exp(x, orderMinus2, order)
}
