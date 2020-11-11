package crypto

import (
	"fmt"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// EonSKShare represents a share of the eon secret key.
type EonSKShare big.Int

// EonPKShare represents a share of the eon public key.
type EonPKShare bn256.G2

// EonPK represents the combined eon public key.
type EonPK bn256.G2

// EpochID is the identifier of an epoch.
type EpochID bn256.G1

// EpochSKShare represents a keyper's share of the epoch sk key.
type EpochSKShare bn256.G1

// EpochSK represents an epoch secret key.
type EpochSK bn256.G1

// ComputeEonSKShare computes the a keyper's sk share from the set of poly evals received from the
// other keypers.
func ComputeEonSKShare(polyEvals []*big.Int) *EonSKShare {
	res := big.NewInt(0)
	for _, vi := range polyEvals {
		res.Add(res, vi)
		res.Mod(res, bn256.Order)
	}
	share := EonSKShare(*res)
	return &share
}

// ComputeEonPKShare computes the eon public key share of the given keyper.
func ComputeEonPKShare(keyperIndex int, gammas []*Gammas) *EonPKShare {
	g2 := new(bn256.G2).Set(zeroG2)
	keyperX := KeyperX(keyperIndex)
	for _, gs := range gammas {
		mu := gs.Mu(keyperX)
		g2 = new(bn256.G2).Add(g2, mu)
	}
	epk := EonPKShare(*g2)
	return &epk
}

// ComputeEonPK computes the combined eon public key from the set of eon public key shares.
func ComputeEonPK(pkShares []*EonPKShare) *EonPK {
	g2 := new(bn256.G2).Set(zeroG2)
	for _, share := range pkShares {
		tmp := new(bn256.G2)
		tmp.Add(g2, (*bn256.G2)(share))
		g2 = tmp
	}
	epk := EonPK(*g2)
	return &epk
}

// ComputeEpochSKShare computes a keyper's epoch sk share.
func ComputeEpochSKShare(eonSKShare *EonSKShare, epochID *EpochID) *EpochSKShare {
	g1 := new(bn256.G1).ScalarMult((*bn256.G1)(epochID), (*big.Int)(eonSKShare))
	epochSKShare := EpochSKShare(*g1)
	return &epochSKShare
}

// ComputeEpochID computes the id of the given epoch.
func ComputeEpochID(epochIndex uint64) *EpochID {
	epochIndexBig := new(big.Int).SetUint64(epochIndex)
	id := EpochID(*new(bn256.G1).ScalarBaseMult(epochIndexBig))
	return &id
}

// ComputeEpochSK computes the epoch secret key from a set of shares.
func ComputeEpochSK(keyperIndices []int, epochSKShares []*EpochSKShare, threshold uint64) (*EpochSK, error) {
	if len(keyperIndices) != len(epochSKShares) {
		return nil, fmt.Errorf("got %d keyper indices, but %d secret shares", len(keyperIndices), len(epochSKShares))
	}
	if uint64(len(keyperIndices)) != threshold {
		return nil, fmt.Errorf("got %d shares, but threshold is %d", len(keyperIndices), threshold)
	}

	skG1 := new(bn256.G1).Set(zeroG1)
	for i := 0; i < len(keyperIndices); i++ {
		keyperIndex := keyperIndices[i]
		share := epochSKShares[i]

		lambda := lagrangeCoefficient(keyperIndex, keyperIndices)
		qTimesLambda := new(bn256.G1).ScalarMult((*bn256.G1)(share), lambda)
		skG1 = new(bn256.G1).Add(skG1, qTimesLambda)
	}
	sk := EpochSK(*skG1)
	return &sk, nil
}

// VerifyEpochSKShare checks that an epoch sk share published by a keyper is correct.
func VerifyEpochSKShare(epochSKShare *EpochSKShare, eonPKShare *EonPKShare, epochID *EpochID) bool {
	g1s := []*bn256.G1{
		(*bn256.G1)(epochSKShare),
		new(bn256.G1).Neg((*bn256.G1)(epochID)),
	}
	g2s := []*bn256.G2{
		new(bn256.G2).ScalarBaseMult(big.NewInt(1)),
		(*bn256.G2)(eonPKShare),
	}
	return bn256.PairingCheck(g1s, g2s)
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
