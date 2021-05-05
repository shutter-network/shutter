package medley

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// ChecksumAddr is used internally to serialize addresses as checksum addresses when writing JSON
// files.
type ChecksumAddr common.Address

func (addr ChecksumAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(common.Address(addr).Hex())
}

func (addr *ChecksumAddr) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if !common.IsHexAddress(s) {
		return errors.Errorf("not a hex address: %s", s)
	}
	a := common.HexToAddress(s)
	if a.Hex() != s {
		return errors.Errorf("hex address with bad checksum: %s", s)
	}
	copy(addr[:], a[:])
	return nil
}
