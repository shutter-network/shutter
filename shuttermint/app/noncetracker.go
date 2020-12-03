package app

import "github.com/ethereum/go-ethereum/common"

// NewNonceTracker creates a new NonceTracker.
func NewNonceTracker() *NonceTracker {
	return &NonceTracker{
		RandomNonces: make(map[common.Address]map[uint64]bool),
	}
}

// Check returns true if the given nonce is free and false if it has been added already.
func (t *NonceTracker) Check(sender common.Address, randomNonce uint64) bool {
	m, ok := t.RandomNonces[sender]
	if !ok {
		return true
	}
	return !m[randomNonce]
}

// Add adds the given nonce if it hasn't already.
func (t *NonceTracker) Add(sender common.Address, randomNonce uint64) {
	m, ok := t.RandomNonces[sender]
	if !ok {
		m = make(map[uint64]bool)
		t.RandomNonces[sender] = m
	}
	m[randomNonce] = true
}
