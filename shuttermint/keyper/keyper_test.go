package keyper

import (
	"testing"
)

// TestShortInfo tests that Keyper.ShortInfo() does not panic, even though the Shutter and
// MainChain objects are empty. We shipped broken release with this problem once.
func TestShortInfo(_ *testing.T) {
	k := NewKeyper(Config{})
	k.ShortInfo()
}
