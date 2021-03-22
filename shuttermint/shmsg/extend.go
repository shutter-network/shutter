package shmsg

/* Extend the protocol buffer group types with easy marshaling/unmarshaling to the native group
   types
*/

import (
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// Set marshals the given value and stores the byte array.
func (g1 *G1) Set(v *bn256.G1) {
	g1.G1Bytes = v.Marshal()
}

// Get unmarshals the marshaled value.
func (g1 *G1) Get() (*bn256.G1, error) {
	v := new(bn256.G1)
	_, err := v.Unmarshal(g1.G1Bytes)
	return v, err
}

// Set marshals the given value and stores the byte array.
func (g2 *G2) Set(v *bn256.G2) {
	g2.G2Bytes = v.Marshal()
}

// Get unmarshals the marshaled value.
func (g2 *G2) Get() (*bn256.G2, error) {
	v := new(bn256.G2)
	_, err := v.Unmarshal(g2.G2Bytes)
	return v, err
}

// Set marshals the given value and stores the byte array.
func (gt *GT) Set(v *bn256.GT) {
	gt.Gtbytes = v.Marshal()
}

// Get unmarshals the marshaled value.
func (gt *GT) Get() (*bn256.GT, error) {
	v := new(bn256.GT)
	_, err := v.Unmarshal(gt.Gtbytes)
	return v, err
}
