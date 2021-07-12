package deploy

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shuttermint/medley"
)

type ChecksumAddr = medley.ChecksumAddr

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
	TargetProxyContract   common.Address
	TargetContract        common.Address
}

type contractsJSON struct {
	ConfigContract        ChecksumAddr
	KeyBroadcastContract  ChecksumAddr
	FeeBankContract       ChecksumAddr
	BatcherContract       ChecksumAddr
	ExecutorContract      ChecksumAddr
	TokenContract         ChecksumAddr
	DepositContract       ChecksumAddr
	KeyperSlasherContract ChecksumAddr
	TargetProxyContract   ChecksumAddr
	TargetContract        ChecksumAddr
}

// MarshalJSON makes us output checksum addresses when marshaling as json.
func (c Contracts) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		contractsJSON{
			ChecksumAddr(c.ConfigContract),
			ChecksumAddr(c.KeyBroadcastContract),
			ChecksumAddr(c.FeeBankContract),
			ChecksumAddr(c.BatcherContract),
			ChecksumAddr(c.ExecutorContract),
			ChecksumAddr(c.TokenContract),
			ChecksumAddr(c.DepositContract),
			ChecksumAddr(c.KeyperSlasherContract),
			ChecksumAddr(c.TargetProxyContract),
			ChecksumAddr(c.TargetContract),
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
	c.TargetProxyContract = common.Address(tmp.TargetProxyContract)
	c.TargetContract = common.Address(tmp.TargetContract)
	return nil
}

// LoadContractsJSON loads and validates a contracts json file.
func LoadContractsJSON(path string) (*Contracts, error) {
	c := Contracts{}
	err := c.LoadJSON(path)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// LoadJSON loads a contracts.json file.
func (c *Contracts) LoadJSON(path string) error {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(d, c)
	if err != nil {
		return errors.Wrapf(err, "malformed contracts JSON file %s", path)
	}
	return nil
}

// SaveJSON saves a contracts.json file.
func (c *Contracts) SaveJSON(outputFile string) error {
	s, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outputFile, s, 0o644)
	if err != nil {
		return err
	}
	return nil
}
