package sandbox

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ContractJSON stores the hex encoded addresses of all contracts.
type ContractsJSON struct {
	ConfigContract        string
	KeyBroadcastContract  string
	FeeBankContract       string
	BatcherContract       string
	ExecutorContract      string
	TokenContract         string
	DepositContract       string
	KeyperSlasherContract string
	TargetContract        string
}

func (c ContractsJSON) Validate() error {
	addresses := []string{
		c.ConfigContract,
		c.KeyBroadcastContract,
		c.FeeBankContract,
		c.BatcherContract,
		c.ExecutorContract,
		c.TokenContract,
		c.DepositContract,
		c.KeyperSlasherContract,
		c.TargetContract,
	}
	for _, address := range addresses {
		if err := validateAddress(address); err != nil {
			return err
		}
	}
	return nil
}

func validateAddress(address string) error {
	addressParsed := common.HexToAddress(address)
	if addressParsed.Hex() != address {
		return fmt.Errorf("invalid address")
	}
	return nil
}
