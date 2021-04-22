package keyper

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config contains validated configuration parameters for the keyper client.
type Config struct {
	ChainID                     string
	ShuttermintURL              string
	EthereumURL                 string
	DBDir                       string
	SigningKey                  *ecdsa.PrivateKey
	ValidatorKey                ed25519.PrivateKey `mapstructure:"ValidatorSeed"`
	EncryptionKey               *ecies.PrivateKey
	ConfigContractAddress       common.Address `mapstructure:"ConfigContract"`
	BatcherContractAddress      common.Address `mapstructure:"BatcherContract"`
	KeyBroadcastContractAddress common.Address `mapstructure:"KeyBroadcastContract"`
	ExecutorContractAddress     common.Address `mapstructure:"ExecutorContract"`
	DepositContractAddress      common.Address `mapstructure:"DepositContract"`
	KeyperSlasherAddress        common.Address `mapstructure:"KeyperSlasher"`
	MainChainFollowDistance     uint64         // in main chain blocks
	ExecutionStaggering         uint64         // in main chain blocks
	DKGPhaseLength              uint64         // in shuttermint blocks
}

func stringToEd25519PrivateKey(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String || t != reflect.TypeOf(ed25519.PrivateKey{}) {
		return data, nil
	}
	seed, err := hex.DecodeString(data.(string))
	if err != nil {
		return nil, err
	}
	if len(seed) != ed25519.SeedSize {
		return nil, errors.Errorf("invalid seed length %d (must be %d)", len(seed), ed25519.SeedSize)
	}
	return ed25519.NewKeyFromSeed(seed), nil
}

func stringToEcdsaPrivateKey(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String || t != reflect.TypeOf(&ecdsa.PrivateKey{}) {
		return data, nil
	}
	return crypto.HexToECDSA(data.(string))
}

func stringToEciesPrivateKey(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String || t != reflect.TypeOf(&ecies.PrivateKey{}) {
		return data, nil
	}
	encryptionKeyECDSA, err := crypto.HexToECDSA(data.(string))
	if err != nil {
		return nil, err
	}

	return ecies.ImportECDSA(encryptionKeyECDSA), nil
}

func stringToAddress(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String || t != reflect.TypeOf(common.Address{}) {
		return data, nil
	}

	ds := data.(string)
	addr := common.HexToAddress(ds)
	if addr.Hex() != ds {
		return nil, fmt.Errorf("not a checksummed address: %s", ds)
	}
	return addr, nil
}

// Unmarshal unmarshals a keyper Config from the the given Viper object.
func (config *Config) Unmarshal(v *viper.Viper) error {
	return v.Unmarshal(
		config,
		viper.DecodeHook(
			mapstructure.ComposeDecodeHookFunc(
				stringToEd25519PrivateKey,
				stringToEcdsaPrivateKey,
				stringToEciesPrivateKey,
				stringToAddress,
				mapstructure.StringToTimeDurationHookFunc(),
				mapstructure.StringToSliceHookFunc(","),
			),
		),
	)
}

// Address returns the keyper's Ethereum address.
func (config *Config) Address() common.Address {
	return crypto.PubkeyToAddress(config.SigningKey.PublicKey)
}
