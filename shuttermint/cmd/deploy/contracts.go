package deploy

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// chksumAddr is used internally to serialize addresses as checksum addresses when writing JSON
// files.
type chksumAddr common.Address

func (addr chksumAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(common.Address(addr).Hex())
}

func (addr *chksumAddr) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	a := common.HexToAddress(s)
	if a.Hex() != s {
		return errors.Errorf("invalid address: %s", s)
	}
	copy(addr[:], a[:])
	return nil
}

// Contracts stores the addresses of all contracts.
type Contracts struct {
	ConfigContract        common.Address
	KeyBroadcastContract  common.Address
	FeeBankContract       common.Address
	BatcherContract       common.Address
	ExecutorContract      common.Address
	TokenContract         common.Address
	DepositContract       common.Address
	KeyperSlasherContract common.Address
	TargetContract        common.Address
}

type contractsJSON struct {
	ConfigContract        chksumAddr
	KeyBroadcastContract  chksumAddr
	FeeBankContract       chksumAddr
	BatcherContract       chksumAddr
	ExecutorContract      chksumAddr
	TokenContract         chksumAddr
	DepositContract       chksumAddr
	KeyperSlasherContract chksumAddr
	TargetContract        chksumAddr
}

// MarshalJSON makes us output checksum addresses when marshaling as json.
func (c Contracts) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		contractsJSON{
			chksumAddr(c.ConfigContract),
			chksumAddr(c.KeyBroadcastContract),
			chksumAddr(c.FeeBankContract),
			chksumAddr(c.BatcherContract),
			chksumAddr(c.ExecutorContract),
			chksumAddr(c.TokenContract),
			chksumAddr(c.DepositContract),
			chksumAddr(c.KeyperSlasherContract),
			chksumAddr(c.TargetContract),
		},
	)
}

func (c *Contracts) UnmarshalJSON(data []byte) error {
	tmp := contractsJSON{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	c.ConfigContract = common.Address(tmp.ConfigContract)
	c.KeyBroadcastContract = common.Address(tmp.KeyBroadcastContract)
	c.FeeBankContract = common.Address(tmp.FeeBankContract)
	c.BatcherContract = common.Address(tmp.BatcherContract)
	c.ExecutorContract = common.Address(tmp.ExecutorContract)
	c.TokenContract = common.Address(tmp.TokenContract)
	c.DepositContract = common.Address(tmp.DepositContract)
	c.KeyperSlasherContract = common.Address(tmp.KeyperSlasherContract)
	c.TargetContract = common.Address(tmp.TargetContract)
	return nil
}

// LoadContractsJSON loads and validates a contracts json file.
func LoadContractsJSON(path string) (*Contracts, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Contracts{}
	err = json.Unmarshal(d, c)
	if err != nil {
		return nil, errors.Wrap(err, "malformed contracts JSON")
	}

	return c, nil
}
