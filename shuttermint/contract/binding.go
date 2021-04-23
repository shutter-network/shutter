// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Authorization is an auto generated low-level Go binding around an user-defined struct.
type Authorization struct {
	HalfStep      uint64
	BatchHash     [32]byte
	SignerIndices []uint64
	Signatures    [][]byte
}

// BatchConfig is an auto generated low-level Go binding around an user-defined struct.
type BatchConfig struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Keypers                []common.Address
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}

// CipherExecutionReceipt is an auto generated low-level Go binding around an user-defined struct.
type CipherExecutionReceipt struct {
	Executed        bool
	Executor        common.Address
	HalfStep        uint64
	CipherBatchHash [32]byte
	BatchHash       [32]byte
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a25df4f2fda4d58ac2869f48bcfa3eb96c2019c11f97cd2e834de58954eb427f64736f6c63430008040033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BatcherContractABI is the input ABI used to generate the binding from.
const BatcherContractABI = "[{\"inputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"configContractAddress\",\"type\":\"address\"},{\"internalType\":\"contractFeeBankContract\",\"name\":\"feeBankContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTransactionType\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"addTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchSizes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBankContract\",\"outputs\":[{\"internalType\":\"contractFeeBankContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newMinFee\",\"type\":\"uint64\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BatcherContractBin is the compiled bytecode used for deploying new contracts.
var BatcherContractBin = "0x608060405234801561001057600080fd5b5060405161104e38038061104e83398101604081905261002f916100a1565b600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100f2565b600080604083850312156100b3578182fd5b82516100be816100da565b60208401519092506100cf816100da565b809150509250929050565b6001600160a01b03811681146100ef57600080fd5b50565b610f4d806101016000396000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b14610155578063bf66a18214610173578063bfd260ca14610193578063c87afa8a146101c9578063f2fde38b1461020f57600080fd5b8063246673dc1461009657806324ec7590146100ab57806336e1290d146100e857806348fd5acc14610120578063715018a614610140575b600080fd5b6100a96100a4366004610d09565b61022f565b005b3480156100b757600080fd5b506005546100cb906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b3480156100f457600080fd5b50600254610108906001600160a01b031681565b6040516001600160a01b0390911681526020016100df565b34801561012c57600080fd5b506100a961013b366004610c9d565b610a7f565b34801561014c57600080fd5b506100a9610acc565b34801561016157600080fd5b506000546001600160a01b0316610108565b34801561017f57600080fd5b50600154610108906001600160a01b031681565b34801561019f57600080fd5b506100cb6101ae366004610c9d565b6003602052600090815260409020546001600160401b031681565b3480156101d557600080fd5b506102016101e4366004610cd5565b600460209081526000928352604080842090915290825290205481565b6040519081526020016100df565b34801561021b57600080fd5b506100a961022a366004610c3e565b610b40565b600154604051632e37abf360e21b81526001600160401b03861660048201526000916001600160a01b03169063b8deafcc9060240160206040518083038186803b15801561027c57600080fd5b505afa158015610290573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b49190610cb9565b60015460405163cab61fb960e01b81526001600160401b03831660048201529192506000916001600160a01b039091169063cab61fb99060240160206040518083038186803b15801561030657600080fd5b505afa15801561031a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033e9190610cb9565b60015460405163026fee7d60e61b81526001600160401b03851660048201529192506000916001600160a01b0390911690639bfb9f409060240160206040518083038186803b15801561039057600080fd5b505afa1580156103a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c89190610cb9565b6001546040516321be792760e11b81526001600160401b03861660048201529192506000916001600160a01b039091169063437cf24e9060240160206040518083038186803b15801561041a57600080fd5b505afa15801561042e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104529190610c61565b6001546040516359d6accd60e01b81526001600160401b03871660048201529192506001600160a01b0316906359d6accd9060240160206040518083038186803b15801561049f57600080fd5b505afa1580156104b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d79190610c7d565b6105325760405162461bcd60e51b815260206004820152602160248201527f42617463686572436f6e74726163743a206261746368206e6f742061637469766044820152606560f81b60648201526084015b60405180910390fd5b6001546040516368080ea760e01b81526001600160401b0380871660048301528a16602482015260009182916001600160a01b03909116906368080ea79060440160606040518083038186803b15801561058b57600080fd5b505afa15801561059f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c39190610d93565b5091509150816001600160401b031643101561062d5760405162461bcd60e51b815260206004820152602360248201527f42617463686572436f6e74726163743a206261746368206e6f74206f70656e206044820152621e595d60ea1b6064820152608401610529565b806001600160401b031643106106935760405162461bcd60e51b815260206004820152602560248201527f42617463686572436f6e74726163743a20626174636820616c726561647920636044820152641b1bdcd95960da1b6064820152608401610529565b866106ee5760405162461bcd60e51b815260206004820152602560248201527f42617463686572436f6e74726163743a207472616e73616374696f6e20697320604482015264656d70747960d81b6064820152608401610529565b6001600160401b0385168711156107535760405162461bcd60e51b8152602060048201526024808201527f42617463686572436f6e74726163743a207472616e73616374696f6e20746f6f6044820152632062696760e01b6064820152608401610529565b6001600160401b038a81166000908152600360205260409020548186169161077d918a9116610e91565b11156107d75760405162461bcd60e51b815260206004820152602360248201527f42617463686572436f6e74726163743a20626174636820616c726561647920666044820152621d5b1b60ea1b6064820152608401610529565b6005546001600160401b03163410156108325760405162461bcd60e51b815260206004820152601e60248201527f42617463686572436f6e74726163743a2066656520746f6f20736d616c6c00006044820152606401610529565b6001600160401b038a16600090815260046020526040812089908990838d600181111561086f57634e487b7160e01b600052602160045260246000fd5b600181111561088e57634e487b7160e01b600052602160045260246000fd5b8152602001908152602001600020546040516020016108af93929190610ddf565b604051602081830303815290604052905060008180519060200120905080600460008e6001600160401b03166001600160401b0316815260200190815260200160002060008d600181111561091457634e487b7160e01b600052602160045260246000fd5b600181111561093357634e487b7160e01b600052602160045260246000fd5b81526020019081526020016000208190555089899050600360008e6001600160401b03166001600160401b0316815260200190815260200160002060008282829054906101000a90046001600160401b031661098f9190610ea9565b92506101000a8154816001600160401b0302191690836001600160401b031602179055506000341180156109cb57506001600160a01b03851615155b15610a325760025460405163f340fa0160e01b81526001600160a01b0387811660048301529091169063f340fa019034906024016000604051808303818588803b158015610a1857600080fd5b505af1158015610a2c573d6000803e3d6000fd5b50505050505b7ffc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c988c8c8c8c85604051610a69959493929190610e26565b60405180910390a1505050505050505050505050565b6000546001600160a01b03163314610aa95760405162461bcd60e51b815260040161052990610df1565b6005805467ffffffffffffffff19166001600160401b0392909216919091179055565b6000546001600160a01b03163314610af65760405162461bcd60e51b815260040161052990610df1565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314610b6a5760405162461bcd60e51b815260040161052990610df1565b6001600160a01b038116610bcf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610529565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b803560028110610c3957600080fd5b919050565b600060208284031215610c4f578081fd5b8135610c5a81610eea565b9392505050565b600060208284031215610c72578081fd5b8151610c5a81610eea565b600060208284031215610c8e578081fd5b81518015158114610c5a578182fd5b600060208284031215610cae578081fd5b8135610c5a81610f02565b600060208284031215610cca578081fd5b8151610c5a81610f02565b60008060408385031215610ce7578081fd5b8235610cf281610f02565b9150610d0060208401610c2a565b90509250929050565b60008060008060608587031215610d1e578182fd5b8435610d2981610f02565b9350610d3760208601610c2a565b925060408501356001600160401b0380821115610d52578384fd5b818701915087601f830112610d65578384fd5b813581811115610d73578485fd5b886020828501011115610d84578485fd5b95989497505060200194505050565b600080600060608486031215610da7578283fd5b8351610db281610f02565b6020850151909350610dc381610f02565b6040850151909250610dd481610f02565b809150509250925092565b82848237909101908152602001919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6001600160401b0386168152600060028610610e5057634e487b7160e01b81526021600452602481fd5b85602083015260806040830152836080830152838560a08401378060a0858401015260a0601f19601f86011683010190508260608301529695505050505050565b60008219821115610ea457610ea4610ed4565b500190565b60006001600160401b03808316818516808303821115610ecb57610ecb610ed4565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610eff57600080fd5b50565b6001600160401b0381168114610eff57600080fdfea2646970667358221220c38f922f9698bbf58fba681acdde18369b7fe78306d9200dd8d0468a943f799a64736f6c63430008040033"

// DeployBatcherContract deploys a new Ethereum contract, binding an instance of BatcherContract to it.
func DeployBatcherContract(auth *bind.TransactOpts, backend bind.ContractBackend, configContractAddress common.Address, feeBankContractAddress common.Address) (common.Address, *types.Transaction, *BatcherContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatcherContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatcherContractBin), backend, configContractAddress, feeBankContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatcherContract{BatcherContractCaller: BatcherContractCaller{contract: contract}, BatcherContractTransactor: BatcherContractTransactor{contract: contract}, BatcherContractFilterer: BatcherContractFilterer{contract: contract}}, nil
}

// BatcherContract is an auto generated Go binding around an Ethereum contract.
type BatcherContract struct {
	BatcherContractCaller     // Read-only binding to the contract
	BatcherContractTransactor // Write-only binding to the contract
	BatcherContractFilterer   // Log filterer for contract events
}

// BatcherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatcherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatcherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatcherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatcherContractSession struct {
	Contract     *BatcherContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatcherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatcherContractCallerSession struct {
	Contract *BatcherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BatcherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatcherContractTransactorSession struct {
	Contract     *BatcherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BatcherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatcherContractRaw struct {
	Contract *BatcherContract // Generic contract binding to access the raw methods on
}

// BatcherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatcherContractCallerRaw struct {
	Contract *BatcherContractCaller // Generic read-only contract binding to access the raw methods on
}

// BatcherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatcherContractTransactorRaw struct {
	Contract *BatcherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatcherContract creates a new instance of BatcherContract, bound to a specific deployed contract.
func NewBatcherContract(address common.Address, backend bind.ContractBackend) (*BatcherContract, error) {
	contract, err := bindBatcherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatcherContract{BatcherContractCaller: BatcherContractCaller{contract: contract}, BatcherContractTransactor: BatcherContractTransactor{contract: contract}, BatcherContractFilterer: BatcherContractFilterer{contract: contract}}, nil
}

// NewBatcherContractCaller creates a new read-only instance of BatcherContract, bound to a specific deployed contract.
func NewBatcherContractCaller(address common.Address, caller bind.ContractCaller) (*BatcherContractCaller, error) {
	contract, err := bindBatcherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatcherContractCaller{contract: contract}, nil
}

// NewBatcherContractTransactor creates a new write-only instance of BatcherContract, bound to a specific deployed contract.
func NewBatcherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BatcherContractTransactor, error) {
	contract, err := bindBatcherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatcherContractTransactor{contract: contract}, nil
}

// NewBatcherContractFilterer creates a new log filterer instance of BatcherContract, bound to a specific deployed contract.
func NewBatcherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BatcherContractFilterer, error) {
	contract, err := bindBatcherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatcherContractFilterer{contract: contract}, nil
}

// bindBatcherContract binds a generic wrapper to an already deployed contract.
func bindBatcherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatcherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatcherContract *BatcherContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatcherContract.Contract.BatcherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatcherContract *BatcherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatcherContract.Contract.BatcherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatcherContract *BatcherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatcherContract.Contract.BatcherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatcherContract *BatcherContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatcherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatcherContract *BatcherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatcherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatcherContract *BatcherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatcherContract.Contract.contract.Transact(opts, method, params...)
}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_BatcherContract *BatcherContractCaller) BatchHashes(opts *bind.CallOpts, arg0 uint64, arg1 uint8) ([32]byte, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "batchHashes", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_BatcherContract *BatcherContractSession) BatchHashes(arg0 uint64, arg1 uint8) ([32]byte, error) {
	return _BatcherContract.Contract.BatchHashes(&_BatcherContract.CallOpts, arg0, arg1)
}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_BatcherContract *BatcherContractCallerSession) BatchHashes(arg0 uint64, arg1 uint8) ([32]byte, error) {
	return _BatcherContract.Contract.BatchHashes(&_BatcherContract.CallOpts, arg0, arg1)
}

// BatchSizes is a free data retrieval call binding the contract method 0xbfd260ca.
//
// Solidity: function batchSizes(uint64 ) view returns(uint64)
func (_BatcherContract *BatcherContractCaller) BatchSizes(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "batchSizes", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BatchSizes is a free data retrieval call binding the contract method 0xbfd260ca.
//
// Solidity: function batchSizes(uint64 ) view returns(uint64)
func (_BatcherContract *BatcherContractSession) BatchSizes(arg0 uint64) (uint64, error) {
	return _BatcherContract.Contract.BatchSizes(&_BatcherContract.CallOpts, arg0)
}

// BatchSizes is a free data retrieval call binding the contract method 0xbfd260ca.
//
// Solidity: function batchSizes(uint64 ) view returns(uint64)
func (_BatcherContract *BatcherContractCallerSession) BatchSizes(arg0 uint64) (uint64, error) {
	return _BatcherContract.Contract.BatchSizes(&_BatcherContract.CallOpts, arg0)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_BatcherContract *BatcherContractCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_BatcherContract *BatcherContractSession) ConfigContract() (common.Address, error) {
	return _BatcherContract.Contract.ConfigContract(&_BatcherContract.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_BatcherContract *BatcherContractCallerSession) ConfigContract() (common.Address, error) {
	return _BatcherContract.Contract.ConfigContract(&_BatcherContract.CallOpts)
}

// FeeBankContract is a free data retrieval call binding the contract method 0x36e1290d.
//
// Solidity: function feeBankContract() view returns(address)
func (_BatcherContract *BatcherContractCaller) FeeBankContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "feeBankContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeBankContract is a free data retrieval call binding the contract method 0x36e1290d.
//
// Solidity: function feeBankContract() view returns(address)
func (_BatcherContract *BatcherContractSession) FeeBankContract() (common.Address, error) {
	return _BatcherContract.Contract.FeeBankContract(&_BatcherContract.CallOpts)
}

// FeeBankContract is a free data retrieval call binding the contract method 0x36e1290d.
//
// Solidity: function feeBankContract() view returns(address)
func (_BatcherContract *BatcherContractCallerSession) FeeBankContract() (common.Address, error) {
	return _BatcherContract.Contract.FeeBankContract(&_BatcherContract.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() view returns(uint64)
func (_BatcherContract *BatcherContractCaller) MinFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "minFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() view returns(uint64)
func (_BatcherContract *BatcherContractSession) MinFee() (uint64, error) {
	return _BatcherContract.Contract.MinFee(&_BatcherContract.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() view returns(uint64)
func (_BatcherContract *BatcherContractCallerSession) MinFee() (uint64, error) {
	return _BatcherContract.Contract.MinFee(&_BatcherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatcherContract *BatcherContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatcherContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatcherContract *BatcherContractSession) Owner() (common.Address, error) {
	return _BatcherContract.Contract.Owner(&_BatcherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatcherContract *BatcherContractCallerSession) Owner() (common.Address, error) {
	return _BatcherContract.Contract.Owner(&_BatcherContract.CallOpts)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x246673dc.
//
// Solidity: function addTransaction(uint64 batchIndex, uint8 transactionType, bytes transaction) payable returns()
func (_BatcherContract *BatcherContractTransactor) AddTransaction(opts *bind.TransactOpts, batchIndex uint64, transactionType uint8, transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "addTransaction", batchIndex, transactionType, transaction)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x246673dc.
//
// Solidity: function addTransaction(uint64 batchIndex, uint8 transactionType, bytes transaction) payable returns()
func (_BatcherContract *BatcherContractSession) AddTransaction(batchIndex uint64, transactionType uint8, transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.Contract.AddTransaction(&_BatcherContract.TransactOpts, batchIndex, transactionType, transaction)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x246673dc.
//
// Solidity: function addTransaction(uint64 batchIndex, uint8 transactionType, bytes transaction) payable returns()
func (_BatcherContract *BatcherContractTransactorSession) AddTransaction(batchIndex uint64, transactionType uint8, transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.Contract.AddTransaction(&_BatcherContract.TransactOpts, batchIndex, transactionType, transaction)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatcherContract *BatcherContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatcherContract *BatcherContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatcherContract.Contract.RenounceOwnership(&_BatcherContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatcherContract *BatcherContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatcherContract.Contract.RenounceOwnership(&_BatcherContract.TransactOpts)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x48fd5acc.
//
// Solidity: function setMinFee(uint64 newMinFee) returns()
func (_BatcherContract *BatcherContractTransactor) SetMinFee(opts *bind.TransactOpts, newMinFee uint64) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "setMinFee", newMinFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x48fd5acc.
//
// Solidity: function setMinFee(uint64 newMinFee) returns()
func (_BatcherContract *BatcherContractSession) SetMinFee(newMinFee uint64) (*types.Transaction, error) {
	return _BatcherContract.Contract.SetMinFee(&_BatcherContract.TransactOpts, newMinFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x48fd5acc.
//
// Solidity: function setMinFee(uint64 newMinFee) returns()
func (_BatcherContract *BatcherContractTransactorSession) SetMinFee(newMinFee uint64) (*types.Transaction, error) {
	return _BatcherContract.Contract.SetMinFee(&_BatcherContract.TransactOpts, newMinFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatcherContract *BatcherContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatcherContract *BatcherContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatcherContract.Contract.TransferOwnership(&_BatcherContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatcherContract *BatcherContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatcherContract.Contract.TransferOwnership(&_BatcherContract.TransactOpts, newOwner)
}

// BatcherContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatcherContract contract.
type BatcherContractOwnershipTransferredIterator struct {
	Event *BatcherContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatcherContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatcherContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatcherContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherContractOwnershipTransferred represents a OwnershipTransferred event raised by the BatcherContract contract.
type BatcherContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatcherContract *BatcherContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatcherContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatcherContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatcherContractOwnershipTransferredIterator{contract: _BatcherContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatcherContract *BatcherContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatcherContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatcherContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherContractOwnershipTransferred)
				if err := _BatcherContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatcherContract *BatcherContractFilterer) ParseOwnershipTransferred(log types.Log) (*BatcherContractOwnershipTransferred, error) {
	event := new(BatcherContractOwnershipTransferred)
	if err := _BatcherContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatcherContractTransactionAddedIterator is returned from FilterTransactionAdded and is used to iterate over the raw logs and unpacked data for TransactionAdded events raised by the BatcherContract contract.
type BatcherContractTransactionAddedIterator struct {
	Event *BatcherContractTransactionAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatcherContractTransactionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherContractTransactionAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatcherContractTransactionAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatcherContractTransactionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherContractTransactionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherContractTransactionAdded represents a TransactionAdded event raised by the BatcherContract contract.
type BatcherContractTransactionAdded struct {
	BatchIndex      uint64
	TransactionType uint8
	Transaction     []byte
	BatchHash       [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionAdded is a free log retrieval operation binding the contract event 0xfc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c98.
//
// Solidity: event TransactionAdded(uint64 batchIndex, uint8 transactionType, bytes transaction, bytes32 batchHash)
func (_BatcherContract *BatcherContractFilterer) FilterTransactionAdded(opts *bind.FilterOpts) (*BatcherContractTransactionAddedIterator, error) {

	logs, sub, err := _BatcherContract.contract.FilterLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return &BatcherContractTransactionAddedIterator{contract: _BatcherContract.contract, event: "TransactionAdded", logs: logs, sub: sub}, nil
}

// WatchTransactionAdded is a free log subscription operation binding the contract event 0xfc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c98.
//
// Solidity: event TransactionAdded(uint64 batchIndex, uint8 transactionType, bytes transaction, bytes32 batchHash)
func (_BatcherContract *BatcherContractFilterer) WatchTransactionAdded(opts *bind.WatchOpts, sink chan<- *BatcherContractTransactionAdded) (event.Subscription, error) {

	logs, sub, err := _BatcherContract.contract.WatchLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherContractTransactionAdded)
				if err := _BatcherContract.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionAdded is a log parse operation binding the contract event 0xfc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c98.
//
// Solidity: event TransactionAdded(uint64 batchIndex, uint8 transactionType, bytes transaction, bytes32 batchHash)
func (_BatcherContract *BatcherContractFilterer) ParseTransactionAdded(log types.Log) (*BatcherContractTransactionAdded, error) {
	event := new(BatcherContractTransactionAdded)
	if err := _BatcherContract.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigContractABI is the input ABI used to generate the binding from.
const ConfigContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"headsUp\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigUnscheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"batchBoundaryBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"batchingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configBatchSizeLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configBatchSpan\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configChangeHeadsUpBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configExecutionTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"configForBatchIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"keypers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structBatchConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configForConfigIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"keypers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structBatchConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"configIndexForBatchIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"keyperIndex\",\"type\":\"uint64\"}],\"name\":\"configKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configStartBatchIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configStartBlockNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configTargetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configTargetFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configTransactionGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"configIndex\",\"type\":\"uint64\"}],\"name\":\"configTransactionSizeLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newKeypers\",\"type\":\"address[]\"}],\"name\":\"nextConfigAddKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigBatchSizeLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigBatchSpan\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigExecutionTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"keyperIndex\",\"type\":\"uint64\"}],\"name\":\"nextConfigKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nextConfigRemoveKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSpan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetExecutionTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"name\":\"nextConfigSetFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBatchIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"}],\"name\":\"nextConfigSetTargetAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"}],\"name\":\"nextConfigSetTargetFunctionSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigStartBatchIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigStartBlockNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigTargetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigTargetFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigTransactionGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigTransactionSizeLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfigs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduleNextConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromStartBlockNumber\",\"type\":\"uint64\"}],\"name\":\"unscheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigContractBin is the compiled bytecode used for deploying new contracts.
var ConfigContractBin = "0x60a06040523480156200001157600080fd5b5060405162002e6338038062002e63833981016040819052620000349162000335565b600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060016200013d604080516101808082018352600080835260208084018290526060848601819052808501839052608080860184905260a080870185905260c080880186905260e0808901879052610100808a01889052610120808b01899052610140808c018a90526101609b8c018a90528c519a8b018d52898b528a89018a90528c518a81529889018d529b8a019790975294880187905292870186905290860185905285018490528401839052830182905282018190529281018390529081019190915290565b815460018181018455600093845260209384902083516005909302018054858501516001600160401b0390811668010000000000000000026001600160801b03199092169416939093179290921782556040830151805193949293620001ac93928501929190910190620002b4565b50606082015160028201805460808086015160a087015160c0808901516001600160401b039788166001600160801b031990961695909517680100000000000000009388168402176001600160801b0316600160801b928816929092026001600160c01b0390811692909217600160c01b95881686021790955560e0808901516003890180546101008c0151928a166001600160e01b0319909116176001600160a01b0392831690950294909417909355610120890151600490980180546101408b0151610160909b0151999094166001600160c01b031994851617600160a01b9a90921c99909902171695909416909102939093179093559290911b909116905262000365565b8280548282559060005260206000209081019282156200030c579160200282015b828111156200030c57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620002d5565b506200031a9291506200031e565b5090565b5b808211156200031a57600081556001016200031f565b60006020828403121562000347578081fd5b81516001600160401b03811681146200035e578182fd5b9392505050565b60805160c01c612ad1620003926000396000818161074101528181610bea0152611f6a0152612ad16000f3fe608060405234801561001057600080fd5b50600436106103265760003560e01c80638da5cb5b116101b8578063cd21aee711610104578063ede71604116100a2578063f7ff384a1161007c578063f7ff384a14610834578063fa84ea0214610846578063fad07bfb14610859578063fda17adf1461086c57600080fd5b8063ede71604146107dd578063f2fde38b14610809578063f464ba631461081c57600080fd5b8063d9a58f24116100de578063d9a58f2414610789578063db293e511461079c578063e0b70ed1146107b4578063eb8b5209146107c557600080fd5b8063cd21aee71461073c578063ce9919b814610763578063d1ac2e521461077657600080fd5b8063b8deafcc11610171578063c25e1f241161014b578063c25e1f24146106f0578063c7c6e9f414610703578063c9515c5814610716578063cab61fb91461072957600080fd5b8063b8deafcc146106b9578063b99fd57c146106cc578063bcf67268146106dd57600080fd5b80638da5cb5b1461064b5780639bfb9f401461065c5780639c7cf8e31461066f5780639d63753e14610680578063ab38348e14610693578063b5a43cef146106a657600080fd5b80635df1683e116102775780636a53e51e11610230578063719f2e171161020a578063719f2e17146105ff57806373ed43db1461061257806381e905a31461062557806386404b7f1461063857600080fd5b80636a53e51e146105cc5780636c52838b146105e4578063715018a6146105f757600080fd5b80635df1683e146104c6578063606df514146104e657806362fced0e146104f957806364e9f6711461050c578063660744dc1461057c57806368080ea71461058f57600080fd5b8063287447c4116102e4578063437cf24e116102be578063437cf24e14610452578063564093fc1461047d57806359d6accd146104905780635dc6fdb8146104b357600080fd5b8063287447c41461041f5780632b2cc6c4146104275780632b6cfe681461043a57600080fd5b806298fa221461032b5780630415b00c146103be5780630b884fa8146103e95780630f0aae6f146103fc578063175cbda91461040457806318b5e83014610415575b600080fd5b61033e610339366004612781565b610884565b604080516001600160401b039c8d1681529a8c1660208c0152988b16988a01989098529589166060890152938816608088015291871660a0870152861660c08601526001600160a01b0390811660e0860152166101008401526001600160e01b031916610120830152909116610140820152610160015b60405180910390f35b6103d16103cc366004612799565b610914565b6040516001600160401b0390911681526020016103b5565b6103d16103f7366004612799565b610967565b6001546103d1565b6005546001600160401b03166103d1565b61041d6109bd565b005b6003546103d1565b61041d6104353660046126bc565b6111ea565b600654600160c01b90046001600160401b03166103d1565b610465610460366004612799565b611245565b6040516001600160a01b0390911681526020016103b5565b61041d61048b366004612799565b61129b565b6104a361049e366004612799565b6112e8565b60405190151581526020016103b5565b61041d6104c1366004612799565b611340565b6104d96104d4366004612799565b611397565b6040516103b5919061285d565b61041d6104f4366004612799565b61150c565b61041d6105073660046126ea565b61155e565b60025460045460055460065461033e936001600160401b0380821694600160401b9283900482169481831694848304841694600160801b8404851694600160c01b94859004811694848216946001600160a01b0393900483169392831692600160a01b810460e01b92919004168b565b61046561058a366004612799565b6116a0565b6105a261059d3660046127b3565b6116ea565b604080516001600160401b03948516815292841660208401529216918101919091526060016103b5565b600454600160401b90046001600160401b03166103d1565b6104d96105f2366004612799565b6118f4565b61041d611a5b565b61041d61060d366004612799565b611acf565b61041d610620366004612799565b611b21565b61041d610633366004612799565b611b6e565b6103d1610646366004612799565b611bc5565b6000546001600160a01b0316610465565b6103d161066a366004612799565b611c1b565b6004546001600160401b03166103d1565b61041d61068e366004612799565b611c71565b6103d16106a1366004612799565b611d2e565b6103d16106b4366004612799565b611d7a565b6103d16106c7366004612799565b611dc9565b6006546001600160a01b0316610465565b61041d6106eb3660046126bc565b611e41565b6104656106fe366004612799565b611e8d565b61041d610711366004612799565b611edc565b61041d610724366004612799565b611f33565b6103d1610737366004612799565b6121a6565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b61041d610771366004612799565b6121fc565b61041d610784366004612759565b612249565b6103d1610797366004612799565b612297565b600554600160401b90046001600160a01b0316610465565b6002546001600160401b03166103d1565b600454600160c01b90046001600160401b03166103d1565b6107f06107eb366004612799565b6122dd565b6040516001600160e01b031990911681526020016103b5565b61041d6108173660046126bc565b61232d565b600454600160801b90046001600160401b03166103d1565b600654600160a01b900460e01b6107f0565b6104656108543660046127b3565b612417565b6103d1610867366004612799565b612498565b600254600160401b90046001600160401b03166103d1565b6001818154811061089457600080fd5b600091825260209091206005909102018054600282015460038301546004909301546001600160401b038084169550600160401b9384900481169483821694808504831694600160801b8104841694600160c01b91829004851694848116946001600160a01b03949004841693821692600160a01b830460e01b9204168b565b60006001826001600160401b03168154811061094057634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020154600160401b90046001600160401b031692915050565b60006001826001600160401b03168154811061099357634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160020154600160401b90046001600160401b031692915050565b6000546001600160a01b031633146109f05760405162461bcd60e51b81526004016109e790612828565b60405180910390fd5b610a0260016001600160401b03612a1f565b6001600160401b031660018054905010610a775760405162461bcd60e51b815260206004820152603060248201527f436f6e666967436f6e74726163743a206e756d626572206f6620636f6e66696760448201526f1cc8195e18d959591cc81d5a5b9d0d8d60821b60648201526084016109e7565b6001805460009190610a8a908290612a08565b81548110610aa857634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805161018081018252600590930290910180546001600160401b038082168552600160401b90910416838501526001810180548351818702810187018552818152949592949386019392830182828015610b3857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b1a575b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b928390048216608080860191909152600386015480841660a08701526001600160a01b03929004821660c086015260049095015490811660e0808601919091526001600160e01b0319600160a01b830490911b1661010085015291909104811661012090920191909152908201519192507f00000000000000000000000000000000000000000000000000000000000000009182821691161115610c1c575060808101515b6000610c2882436129ae565b610c339060016129ae565b6002549091506001600160401b03808316600160401b909204161015610ca95760405162461bcd60e51b815260206004820152602560248201527f436f6e666967436f6e74726163743a20737461727420626c6f636b20746f6f206044820152646561726c7960d81b60648201526084016109e7565b60808301516001600160401b031615610df25782516002546001600160401b03918216911611610d2f5760405162461bcd60e51b815260206004820152602b60248201527f436f6e666967436f6e74726163743a20737461727420626174636820696e646560448201526a1e081d1bdbc81cdb585b1b60aa1b60648201526084016109e7565b8251600254600091610d49916001600160401b0316612a1f565b6002546080860151919250600160401b90046001600160401b031690610d709083906129d9565b8560200151610d7f91906129ae565b6001600160401b031614610dec5760405162461bcd60e51b815260206004820152602e60248201527f436f6e666967436f6e74726163743a20636f6e666967207472616e736974696f60448201526d6e206e6f74207365616d6c65737360901b60648201526084016109e7565b50610e86565b82516002546001600160401b03908116911614610e865760405162461bcd60e51b815260206004820152604660248201527f436f6e666967436f6e74726163743a207472616e736974696f6e2066726f6d2060448201527f696e61637469766520636f6e66696720776974682077726f6e67207374617274606482015265040d2dcc8caf60d31b608482015260a4016109e7565b600180548082018255600091909152600280547fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6600590930292830180546001600160401b0392831667ffffffffffffffff1982168117835584546001600160801b031990921617600160401b91829004909316029190911781556003805492939192610f36927fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7019190612560565b506002828101805491830180546001600160401b0393841667ffffffffffffffff198083168217845584546001600160801b0319909316909117600160401b928390048616830217808455845467ffffffffffffffff60801b198216600160801b918290048816909102908117855594546001600160801b039091166001600160c01b0395861617600160c01b9182900487168202179093556003808801805491880180549288169383168417815590546001600160e01b0319909216909217908390046001600160a01b0390811690930217905560049586018054969095018054969091166001600160a01b031987168117825585546001600160c01b031990971617600160a01b9687900463ffffffff1690960295909517808655935493909116928190049091160217905561106c6124e7565b8051600280546020808501516001600160401b03908116600160401b026001600160801b03199093169416939093171781556040830151805191926110b792600392909101906125b0565b506060820151600282018054608085015160a086015160c08701516001600160401b039586166001600160801b031990941693909317600160401b9286168302176001600160801b0316600160801b918616919091026001600160c01b0390811691909117600160c01b93861684021790935560e0808701516003870180546101008a01519288166001600160e01b0319909116176001600160a01b03928316909402939093179092556101208701516004909601805461014089015161016090990151979093166001600160c01b031990931692909217600160a01b9790911c96909602959095179091169282160291909117909155600154604051911681527f38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae69060200160405180910390a1505050565b6000546001600160a01b031633146112145760405162461bcd60e51b81526004016109e790612828565b600580546001600160a01b03909216600160401b0268010000000000000000600160e01b0319909216919091179055565b60006001826001600160401b03168154811061127157634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160030154600160401b90046001600160a01b031692915050565b6000546001600160a01b031633146112c55760405162461bcd60e51b81526004016109e790612828565b6005805467ffffffffffffffff19166001600160401b0392909216919091179055565b6000806001836001600160401b03168154811061131557634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160020154600160401b90046001600160401b03161192915050565b6000546001600160a01b0316331461136a5760405162461bcd60e51b81526004016109e790612828565b600480546001600160401b03909216600160401b0267ffffffffffffffff60401b19909216919091179055565b61139f612605565b60006113aa83611dc9565b90506001816001600160401b0316815481106113d657634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805161018081018252600590930290910180546001600160401b038082168552600160401b9091041683850152600181018054835181870281018701855281815294959294938601939283018282801561146657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611448575b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9283900482166080850152600385015480831660a0860152046001600160a01b0390811660c085015260049094015493841660e080850191909152600160a01b8504901b6001600160e01b0319166101008401529204909116610120909101529392505050565b6000546001600160a01b031633146115365760405162461bcd60e51b81526004016109e790612828565b600480546001600160401b03909216600160c01b026001600160c01b03909216919091179055565b6000546001600160a01b031633146115885760405162461bcd60e51b81526004016109e790612828565b611599816001600160401b03612a08565b60035411156116035760405162461bcd60e51b815260206004820152603060248201527f436f6e666967436f6e74726163743a206e756d626572206f66206b657970657260448201526f1cc8195e18d959591cc81d5a5b9d0d8d60821b60648201526084016109e7565b60005b6001600160401b03811682111561169b57600383836001600160401b03841681811061164257634e487b7160e01b600052603260045260246000fd5b905060200201602081019061165791906126bc565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b039092169190911790558061169381612a5e565b915050611606565b505050565b60006002600101826001600160401b0316815481106116cf57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031692915050565b6000806000806001866001600160401b03168154811061171a57634e487b7160e01b600052603260045260246000fd5b60009182526020822060059091020154600180546001600160401b03600160401b909304831694509091891690811061176357634e487b7160e01b600052603260045260246000fd5b60009182526020822060059091020154600180546001600160401b03928316945090918a169081106117a557634e487b7160e01b600052603260045260246000fd5b906000526020600020906005020160020160089054906101000a90046001600160401b0316905060006001896001600160401b0316815481106117f857634e487b7160e01b600052603260045260246000fd5b906000526020600020906005020160040160189054906101000a90046001600160401b031690506000826001600160401b03161161184657634e487b7160e01b600052600160045260246000fd5b826001600160401b0316886001600160401b0316101561187657634e487b7160e01b600052600160045260246000fd5b6000611882848a612a1f565b905060008361189181846129d9565b61189b90886129ae565b6118a591906129ae565b905060006118b38583612a1f565b90506001836001600160401b0316106118d3576118d08582612a1f565b90505b60006118df85846129ae565b919a5091985096505050505050509250925092565b6118fc612605565b6001826001600160401b03168154811061192657634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805161018081018252600590930290910180546001600160401b038082168552600160401b909104168385015260018101805483518187028101870185528181529495929493860193928301828280156119b657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611998575b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9283900482166080850152600385015480831660a0860152046001600160a01b0390811660c085015260049094015493841660e080850191909152600160a01b8504901b6001600160e01b03191661010084015292049091166101209091015292915050565b6000546001600160a01b03163314611a855760405162461bcd60e51b81526004016109e790612828565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314611af95760405162461bcd60e51b81526004016109e790612828565b600680546001600160401b03909216600160c01b026001600160c01b03909216919091179055565b6000546001600160a01b03163314611b4b5760405162461bcd60e51b81526004016109e790612828565b6004805467ffffffffffffffff19166001600160401b0392909216919091179055565b6000546001600160a01b03163314611b985760405162461bcd60e51b81526004016109e790612828565b600280546001600160401b03909216600160401b0267ffffffffffffffff60401b19909216919091179055565b60006001826001600160401b031681548110611bf157634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160040154600160c01b90046001600160401b031692915050565b60006001826001600160401b031681548110611c4757634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160020154600160801b90046001600160401b031692915050565b6000546001600160a01b03163314611c9b5760405162461bcd60e51b81526004016109e790612828565b6003546001600160401b0382168110611d1e5760005b826001600160401b0316816001600160401b0316101561169b576003805480611cea57634e487b7160e01b600052603160045260246000fd5b600082815260209020810160001990810180546001600160a01b031916905501905580611d1681612a5e565b915050611cb1565b611d2a6003600061266a565b5050565b60006001826001600160401b031681548110611d5a57634e487b7160e01b600052603260045260246000fd5b60009182526020909120600590910201546001600160401b031692915050565b60006001826001600160401b031681548110611da657634e487b7160e01b600052603260045260246000fd5b60009182526020909120600360059092020101546001600160401b031692915050565b600180546000918291611ddc9190612a08565b90505b826001600160401b031660018281548110611e0a57634e487b7160e01b600052603260045260246000fd5b60009182526020909120600590910201546001600160401b031611611e2f5792915050565b80611e3981612a47565b915050611ddf565b6000546001600160a01b03163314611e6b5760405162461bcd60e51b81526004016109e790612828565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b60006001826001600160401b031681548110611eb957634e487b7160e01b600052603260045260246000fd5b60009182526020909120600460059092020101546001600160a01b031692915050565b6000546001600160a01b03163314611f065760405162461bcd60e51b81526004016109e790612828565b600480546001600160401b03909216600160801b0267ffffffffffffffff60801b19909216919091179055565b6000546001600160a01b03163314611f5d5760405162461bcd60e51b81526004016109e790612828565b611f906001600160401b037f00000000000000000000000000000000000000000000000000000000000000001643612996565b816001600160401b031611611ffa5760405162461bcd60e51b815260206004820152602a60248201527f436f6e666967436f6e74726163743a2066726f6d20737461727420626c6f636b60448201526920746f6f206561726c7960b01b60648201526084016109e7565b600180549060009061200c9083612a08565b90505b80156120f85760006001828154811061203857634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020180549091506001600160401b03808616600160401b90920416106120df57600180548061208557634e487b7160e01b600052603160045260246000fd5b60008281526020812060056000199093019283020180546001600160801b0319168155906120b6600183018261266a565b506000600282018190556003820180546001600160e01b031916905560049091015590556120e5565b506120f8565b50806120f081612a47565b91505061200f565b506001546001600160401b038216116121625760405162461bcd60e51b815260206004820152602660248201527f436f6e666967436f6e74726163743a206e6f20636f6e6669677320756e73636860448201526519591d5b195960d21b60648201526084016109e7565b6001546040516001600160401b0390911681527f202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf9060200160405180910390a15050565b60006001826001600160401b0316815481106121d257634e487b7160e01b600052603260045260246000fd5b6000918252602090912060059091020160020154600160c01b90046001600160401b031692915050565b6000546001600160a01b031633146122265760405162461bcd60e51b81526004016109e790612828565b6002805467ffffffffffffffff19166001600160401b0392909216919091179055565b6000546001600160a01b031633146122735760405162461bcd60e51b81526004016109e790612828565b6006805460e09290921c600160a01b0263ffffffff60a01b19909216919091179055565b60006001826001600160401b0316815481106122c357634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600590920201015492915050565b60006001826001600160401b03168154811061230957634e487b7160e01b600052603260045260246000fd5b6000918252602090912060046005909202010154600160a01b900460e01b92915050565b6000546001600160a01b031633146123575760405162461bcd60e51b81526004016109e790612828565b6001600160a01b0381166123bc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109e7565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006001836001600160401b03168154811061244357634e487b7160e01b600052603260045260246000fd5b9060005260206000209060050201600101826001600160401b03168154811061247c57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03169392505050565b60006001826001600160401b0316815481106124c457634e487b7160e01b600052603260045260246000fd5b60009182526020909120600260059092020101546001600160401b031692915050565b6124ef612605565b506040805161018081018252600080825260208083018290528351828152908101845292820192909252606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b8280548282559060005260206000209081019282156125a05760005260206000209182015b828111156125a0578254825591600101919060010190612585565b506125ac92915061268b565b5090565b8280548282559060005260206000209081019282156125a0579160200282015b828111156125a057825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906125d0565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b5080546000825590600052602060002090810190612688919061268b565b50565b5b808211156125ac576000815560010161268c565b80356001600160401b03811681146126b757600080fd5b919050565b6000602082840312156126cd578081fd5b81356001600160a01b03811681146126e3578182fd5b9392505050565b600080602083850312156126fc578081fd5b82356001600160401b0380821115612712578283fd5b818501915085601f830112612725578283fd5b813581811115612733578384fd5b8660208260051b8501011115612747578384fd5b60209290920196919550909350505050565b60006020828403121561276a578081fd5b81356001600160e01b0319811681146126e3578182fd5b600060208284031215612792578081fd5b5035919050565b6000602082840312156127aa578081fd5b6126e3826126a0565b600080604083850312156127c5578182fd5b6127ce836126a0565b91506127dc602084016126a0565b90509250929050565b6000815180845260208085019450808401835b8381101561281d5781516001600160a01b0316875295820195908201906001016127f8565b509495945050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b602081526128776020820183516001600160401b03169052565b6000602083015161289360408401826001600160401b03169052565b5060408301516101808060608501526128b06101a08501836127e5565b915060608501516128cc60808601826001600160401b03169052565b5060808501516001600160401b03811660a08601525060a08501516001600160401b03811660c08601525060c08501516001600160401b03811660e08601525060e0850151610100612928818701836001600160401b03169052565b8601519050610120612944868201836001600160a01b03169052565b8601519050610140612960868201836001600160a01b03169052565b860151905061016061297d868201836001600160e01b0319169052565b909501516001600160401b031693019290925250919050565b600082198211156129a9576129a9612a85565b500190565b60006001600160401b038083168185168083038211156129d0576129d0612a85565b01949350505050565b60006001600160401b03808316818516818304811182151516156129ff576129ff612a85565b02949350505050565b600082821015612a1a57612a1a612a85565b500390565b60006001600160401b0383811690831681811015612a3f57612a3f612a85565b039392505050565b600081612a5657612a56612a85565b506000190190565b60006001600160401b0380831681811415612a7b57612a7b612a85565b6001019392505050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220f00cf8b72b1c93dbb93943215e115472ee0dbca29f4d1cb0b5e3f38cb1b6208264736f6c63430008040033"

// DeployConfigContract deploys a new Ethereum contract, binding an instance of ConfigContract to it.
func DeployConfigContract(auth *bind.TransactOpts, backend bind.ContractBackend, headsUp uint64) (common.Address, *types.Transaction, *ConfigContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConfigContractBin), backend, headsUp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConfigContract{ConfigContractCaller: ConfigContractCaller{contract: contract}, ConfigContractTransactor: ConfigContractTransactor{contract: contract}, ConfigContractFilterer: ConfigContractFilterer{contract: contract}}, nil
}

// ConfigContract is an auto generated Go binding around an Ethereum contract.
type ConfigContract struct {
	ConfigContractCaller     // Read-only binding to the contract
	ConfigContractTransactor // Write-only binding to the contract
	ConfigContractFilterer   // Log filterer for contract events
}

// ConfigContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigContractSession struct {
	Contract     *ConfigContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigContractCallerSession struct {
	Contract *ConfigContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ConfigContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigContractTransactorSession struct {
	Contract     *ConfigContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ConfigContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigContractRaw struct {
	Contract *ConfigContract // Generic contract binding to access the raw methods on
}

// ConfigContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigContractCallerRaw struct {
	Contract *ConfigContractCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigContractTransactorRaw struct {
	Contract *ConfigContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfigContract creates a new instance of ConfigContract, bound to a specific deployed contract.
func NewConfigContract(address common.Address, backend bind.ContractBackend) (*ConfigContract, error) {
	contract, err := bindConfigContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConfigContract{ConfigContractCaller: ConfigContractCaller{contract: contract}, ConfigContractTransactor: ConfigContractTransactor{contract: contract}, ConfigContractFilterer: ConfigContractFilterer{contract: contract}}, nil
}

// NewConfigContractCaller creates a new read-only instance of ConfigContract, bound to a specific deployed contract.
func NewConfigContractCaller(address common.Address, caller bind.ContractCaller) (*ConfigContractCaller, error) {
	contract, err := bindConfigContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigContractCaller{contract: contract}, nil
}

// NewConfigContractTransactor creates a new write-only instance of ConfigContract, bound to a specific deployed contract.
func NewConfigContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigContractTransactor, error) {
	contract, err := bindConfigContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigContractTransactor{contract: contract}, nil
}

// NewConfigContractFilterer creates a new log filterer instance of ConfigContract, bound to a specific deployed contract.
func NewConfigContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigContractFilterer, error) {
	contract, err := bindConfigContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigContractFilterer{contract: contract}, nil
}

// bindConfigContract binds a generic wrapper to an already deployed contract.
func bindConfigContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConfigContract *ConfigContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfigContract.Contract.ConfigContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConfigContract *ConfigContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfigContract.Contract.ConfigContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConfigContract *ConfigContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfigContract.Contract.ConfigContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConfigContract *ConfigContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfigContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConfigContract *ConfigContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfigContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConfigContract *ConfigContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfigContract.Contract.contract.Transact(opts, method, params...)
}

// BatchBoundaryBlocks is a free data retrieval call binding the contract method 0x68080ea7.
//
// Solidity: function batchBoundaryBlocks(uint64 configIndex, uint64 batchIndex) view returns(uint64, uint64, uint64)
func (_ConfigContract *ConfigContractCaller) BatchBoundaryBlocks(opts *bind.CallOpts, configIndex uint64, batchIndex uint64) (uint64, uint64, uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "batchBoundaryBlocks", configIndex, batchIndex)

	if err != nil {
		return *new(uint64), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// BatchBoundaryBlocks is a free data retrieval call binding the contract method 0x68080ea7.
//
// Solidity: function batchBoundaryBlocks(uint64 configIndex, uint64 batchIndex) view returns(uint64, uint64, uint64)
func (_ConfigContract *ConfigContractSession) BatchBoundaryBlocks(configIndex uint64, batchIndex uint64) (uint64, uint64, uint64, error) {
	return _ConfigContract.Contract.BatchBoundaryBlocks(&_ConfigContract.CallOpts, configIndex, batchIndex)
}

// BatchBoundaryBlocks is a free data retrieval call binding the contract method 0x68080ea7.
//
// Solidity: function batchBoundaryBlocks(uint64 configIndex, uint64 batchIndex) view returns(uint64, uint64, uint64)
func (_ConfigContract *ConfigContractCallerSession) BatchBoundaryBlocks(configIndex uint64, batchIndex uint64) (uint64, uint64, uint64, error) {
	return _ConfigContract.Contract.BatchBoundaryBlocks(&_ConfigContract.CallOpts, configIndex, batchIndex)
}

// BatchingActive is a free data retrieval call binding the contract method 0x59d6accd.
//
// Solidity: function batchingActive(uint64 configIndex) view returns(bool)
func (_ConfigContract *ConfigContractCaller) BatchingActive(opts *bind.CallOpts, configIndex uint64) (bool, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "batchingActive", configIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchingActive is a free data retrieval call binding the contract method 0x59d6accd.
//
// Solidity: function batchingActive(uint64 configIndex) view returns(bool)
func (_ConfigContract *ConfigContractSession) BatchingActive(configIndex uint64) (bool, error) {
	return _ConfigContract.Contract.BatchingActive(&_ConfigContract.CallOpts, configIndex)
}

// BatchingActive is a free data retrieval call binding the contract method 0x59d6accd.
//
// Solidity: function batchingActive(uint64 configIndex) view returns(bool)
func (_ConfigContract *ConfigContractCallerSession) BatchingActive(configIndex uint64) (bool, error) {
	return _ConfigContract.Contract.BatchingActive(&_ConfigContract.CallOpts, configIndex)
}

// ConfigBatchSizeLimit is a free data retrieval call binding the contract method 0x9bfb9f40.
//
// Solidity: function configBatchSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigBatchSizeLimit(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configBatchSizeLimit", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigBatchSizeLimit is a free data retrieval call binding the contract method 0x9bfb9f40.
//
// Solidity: function configBatchSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigBatchSizeLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigBatchSizeLimit(&_ConfigContract.CallOpts, configIndex)
}

// ConfigBatchSizeLimit is a free data retrieval call binding the contract method 0x9bfb9f40.
//
// Solidity: function configBatchSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigBatchSizeLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigBatchSizeLimit(&_ConfigContract.CallOpts, configIndex)
}

// ConfigBatchSpan is a free data retrieval call binding the contract method 0x0b884fa8.
//
// Solidity: function configBatchSpan(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigBatchSpan(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configBatchSpan", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigBatchSpan is a free data retrieval call binding the contract method 0x0b884fa8.
//
// Solidity: function configBatchSpan(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigBatchSpan(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigBatchSpan(&_ConfigContract.CallOpts, configIndex)
}

// ConfigBatchSpan is a free data retrieval call binding the contract method 0x0b884fa8.
//
// Solidity: function configBatchSpan(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigBatchSpan(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigBatchSpan(&_ConfigContract.CallOpts, configIndex)
}

// ConfigChangeHeadsUpBlocks is a free data retrieval call binding the contract method 0xcd21aee7.
//
// Solidity: function configChangeHeadsUpBlocks() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigChangeHeadsUpBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configChangeHeadsUpBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigChangeHeadsUpBlocks is a free data retrieval call binding the contract method 0xcd21aee7.
//
// Solidity: function configChangeHeadsUpBlocks() view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigChangeHeadsUpBlocks() (uint64, error) {
	return _ConfigContract.Contract.ConfigChangeHeadsUpBlocks(&_ConfigContract.CallOpts)
}

// ConfigChangeHeadsUpBlocks is a free data retrieval call binding the contract method 0xcd21aee7.
//
// Solidity: function configChangeHeadsUpBlocks() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigChangeHeadsUpBlocks() (uint64, error) {
	return _ConfigContract.Contract.ConfigChangeHeadsUpBlocks(&_ConfigContract.CallOpts)
}

// ConfigExecutionTimeout is a free data retrieval call binding the contract method 0x86404b7f.
//
// Solidity: function configExecutionTimeout(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigExecutionTimeout(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configExecutionTimeout", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigExecutionTimeout is a free data retrieval call binding the contract method 0x86404b7f.
//
// Solidity: function configExecutionTimeout(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigExecutionTimeout(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigExecutionTimeout(&_ConfigContract.CallOpts, configIndex)
}

// ConfigExecutionTimeout is a free data retrieval call binding the contract method 0x86404b7f.
//
// Solidity: function configExecutionTimeout(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigExecutionTimeout(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigExecutionTimeout(&_ConfigContract.CallOpts, configIndex)
}

// ConfigFeeReceiver is a free data retrieval call binding the contract method 0x437cf24e.
//
// Solidity: function configFeeReceiver(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) ConfigFeeReceiver(opts *bind.CallOpts, configIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configFeeReceiver", configIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigFeeReceiver is a free data retrieval call binding the contract method 0x437cf24e.
//
// Solidity: function configFeeReceiver(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) ConfigFeeReceiver(configIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigFeeReceiver(&_ConfigContract.CallOpts, configIndex)
}

// ConfigFeeReceiver is a free data retrieval call binding the contract method 0x437cf24e.
//
// Solidity: function configFeeReceiver(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) ConfigFeeReceiver(configIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigFeeReceiver(&_ConfigContract.CallOpts, configIndex)
}

// ConfigForBatchIndex is a free data retrieval call binding the contract method 0x5df1683e.
//
// Solidity: function configForBatchIndex(uint64 batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCaller) ConfigForBatchIndex(opts *bind.CallOpts, batchIndex uint64) (BatchConfig, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configForBatchIndex", batchIndex)

	if err != nil {
		return *new(BatchConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BatchConfig)).(*BatchConfig)

	return out0, err

}

// ConfigForBatchIndex is a free data retrieval call binding the contract method 0x5df1683e.
//
// Solidity: function configForBatchIndex(uint64 batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractSession) ConfigForBatchIndex(batchIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.ConfigForBatchIndex(&_ConfigContract.CallOpts, batchIndex)
}

// ConfigForBatchIndex is a free data retrieval call binding the contract method 0x5df1683e.
//
// Solidity: function configForBatchIndex(uint64 batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCallerSession) ConfigForBatchIndex(batchIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.ConfigForBatchIndex(&_ConfigContract.CallOpts, batchIndex)
}

// ConfigForConfigIndex is a free data retrieval call binding the contract method 0x6c52838b.
//
// Solidity: function configForConfigIndex(uint64 configIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCaller) ConfigForConfigIndex(opts *bind.CallOpts, configIndex uint64) (BatchConfig, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configForConfigIndex", configIndex)

	if err != nil {
		return *new(BatchConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BatchConfig)).(*BatchConfig)

	return out0, err

}

// ConfigForConfigIndex is a free data retrieval call binding the contract method 0x6c52838b.
//
// Solidity: function configForConfigIndex(uint64 configIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractSession) ConfigForConfigIndex(configIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.ConfigForConfigIndex(&_ConfigContract.CallOpts, configIndex)
}

// ConfigForConfigIndex is a free data retrieval call binding the contract method 0x6c52838b.
//
// Solidity: function configForConfigIndex(uint64 configIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCallerSession) ConfigForConfigIndex(configIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.ConfigForConfigIndex(&_ConfigContract.CallOpts, configIndex)
}

// ConfigIndexForBatchIndex is a free data retrieval call binding the contract method 0xb8deafcc.
//
// Solidity: function configIndexForBatchIndex(uint64 batchIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigIndexForBatchIndex(opts *bind.CallOpts, batchIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configIndexForBatchIndex", batchIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigIndexForBatchIndex is a free data retrieval call binding the contract method 0xb8deafcc.
//
// Solidity: function configIndexForBatchIndex(uint64 batchIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigIndexForBatchIndex(batchIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigIndexForBatchIndex(&_ConfigContract.CallOpts, batchIndex)
}

// ConfigIndexForBatchIndex is a free data retrieval call binding the contract method 0xb8deafcc.
//
// Solidity: function configIndexForBatchIndex(uint64 batchIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigIndexForBatchIndex(batchIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigIndexForBatchIndex(&_ConfigContract.CallOpts, batchIndex)
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 configIndex, uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) ConfigKeypers(opts *bind.CallOpts, configIndex uint64, keyperIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configKeypers", configIndex, keyperIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 configIndex, uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) ConfigKeypers(configIndex uint64, keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, configIndex, keyperIndex)
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 configIndex, uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) ConfigKeypers(configIndex uint64, keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, configIndex, keyperIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigNumKeypers(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configNumKeypers", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigNumKeypers(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, configIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigNumKeypers(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, configIndex)
}

// ConfigStartBatchIndex is a free data retrieval call binding the contract method 0xab38348e.
//
// Solidity: function configStartBatchIndex(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigStartBatchIndex(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configStartBatchIndex", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigStartBatchIndex is a free data retrieval call binding the contract method 0xab38348e.
//
// Solidity: function configStartBatchIndex(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigStartBatchIndex(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigStartBatchIndex(&_ConfigContract.CallOpts, configIndex)
}

// ConfigStartBatchIndex is a free data retrieval call binding the contract method 0xab38348e.
//
// Solidity: function configStartBatchIndex(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigStartBatchIndex(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigStartBatchIndex(&_ConfigContract.CallOpts, configIndex)
}

// ConfigStartBlockNumber is a free data retrieval call binding the contract method 0x0415b00c.
//
// Solidity: function configStartBlockNumber(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigStartBlockNumber(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configStartBlockNumber", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigStartBlockNumber is a free data retrieval call binding the contract method 0x0415b00c.
//
// Solidity: function configStartBlockNumber(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigStartBlockNumber(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigStartBlockNumber(&_ConfigContract.CallOpts, configIndex)
}

// ConfigStartBlockNumber is a free data retrieval call binding the contract method 0x0415b00c.
//
// Solidity: function configStartBlockNumber(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigStartBlockNumber(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigStartBlockNumber(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTargetAddress is a free data retrieval call binding the contract method 0xc25e1f24.
//
// Solidity: function configTargetAddress(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) ConfigTargetAddress(opts *bind.CallOpts, configIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configTargetAddress", configIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigTargetAddress is a free data retrieval call binding the contract method 0xc25e1f24.
//
// Solidity: function configTargetAddress(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) ConfigTargetAddress(configIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigTargetAddress(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTargetAddress is a free data retrieval call binding the contract method 0xc25e1f24.
//
// Solidity: function configTargetAddress(uint64 configIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) ConfigTargetAddress(configIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigTargetAddress(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xede71604.
//
// Solidity: function configTargetFunctionSelector(uint64 configIndex) view returns(bytes4)
func (_ConfigContract *ConfigContractCaller) ConfigTargetFunctionSelector(opts *bind.CallOpts, configIndex uint64) ([4]byte, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configTargetFunctionSelector", configIndex)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xede71604.
//
// Solidity: function configTargetFunctionSelector(uint64 configIndex) view returns(bytes4)
func (_ConfigContract *ConfigContractSession) ConfigTargetFunctionSelector(configIndex uint64) ([4]byte, error) {
	return _ConfigContract.Contract.ConfigTargetFunctionSelector(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xede71604.
//
// Solidity: function configTargetFunctionSelector(uint64 configIndex) view returns(bytes4)
func (_ConfigContract *ConfigContractCallerSession) ConfigTargetFunctionSelector(configIndex uint64) ([4]byte, error) {
	return _ConfigContract.Contract.ConfigTargetFunctionSelector(&_ConfigContract.CallOpts, configIndex)
}

// ConfigThreshold is a free data retrieval call binding the contract method 0xfad07bfb.
//
// Solidity: function configThreshold(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigThreshold(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configThreshold", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigThreshold is a free data retrieval call binding the contract method 0xfad07bfb.
//
// Solidity: function configThreshold(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigThreshold(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigThreshold(&_ConfigContract.CallOpts, configIndex)
}

// ConfigThreshold is a free data retrieval call binding the contract method 0xfad07bfb.
//
// Solidity: function configThreshold(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigThreshold(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigThreshold(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTransactionGasLimit is a free data retrieval call binding the contract method 0xb5a43cef.
//
// Solidity: function configTransactionGasLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigTransactionGasLimit(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configTransactionGasLimit", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigTransactionGasLimit is a free data retrieval call binding the contract method 0xb5a43cef.
//
// Solidity: function configTransactionGasLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigTransactionGasLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigTransactionGasLimit(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTransactionGasLimit is a free data retrieval call binding the contract method 0xb5a43cef.
//
// Solidity: function configTransactionGasLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigTransactionGasLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigTransactionGasLimit(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xcab61fb9.
//
// Solidity: function configTransactionSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigTransactionSizeLimit(opts *bind.CallOpts, configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configTransactionSizeLimit", configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xcab61fb9.
//
// Solidity: function configTransactionSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigTransactionSizeLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigTransactionSizeLimit(&_ConfigContract.CallOpts, configIndex)
}

// ConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xcab61fb9.
//
// Solidity: function configTransactionSizeLimit(uint64 configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigTransactionSizeLimit(configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigTransactionSizeLimit(&_ConfigContract.CallOpts, configIndex)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractCaller) Configs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configs", arg0)

	outstruct := new(struct {
		StartBatchIndex        uint64
		StartBlockNumber       uint64
		Threshold              uint64
		BatchSpan              uint64
		BatchSizeLimit         uint64
		TransactionSizeLimit   uint64
		TransactionGasLimit    uint64
		FeeReceiver            common.Address
		TargetAddress          common.Address
		TargetFunctionSelector [4]byte
		ExecutionTimeout       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartBlockNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.BatchSpan = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.BatchSizeLimit = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TransactionSizeLimit = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.TransactionGasLimit = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.FeeReceiver = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.TargetAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.TargetFunctionSelector = *abi.ConvertType(out[9], new([4]byte)).(*[4]byte)
	outstruct.ExecutionTimeout = *abi.ConvertType(out[10], new(uint64)).(*uint64)

	return *outstruct, err

}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractSession) Configs(arg0 *big.Int) (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	return _ConfigContract.Contract.Configs(&_ConfigContract.CallOpts, arg0)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractCallerSession) Configs(arg0 *big.Int) (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	return _ConfigContract.Contract.Configs(&_ConfigContract.CallOpts, arg0)
}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractCaller) NextConfig(opts *bind.CallOpts) (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfig")

	outstruct := new(struct {
		StartBatchIndex        uint64
		StartBlockNumber       uint64
		Threshold              uint64
		BatchSpan              uint64
		BatchSizeLimit         uint64
		TransactionSizeLimit   uint64
		TransactionGasLimit    uint64
		FeeReceiver            common.Address
		TargetAddress          common.Address
		TargetFunctionSelector [4]byte
		ExecutionTimeout       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartBlockNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.BatchSpan = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.BatchSizeLimit = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TransactionSizeLimit = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.TransactionGasLimit = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.FeeReceiver = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.TargetAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.TargetFunctionSelector = *abi.ConvertType(out[9], new([4]byte)).(*[4]byte)
	outstruct.ExecutionTimeout = *abi.ConvertType(out[10], new(uint64)).(*uint64)

	return *outstruct, err

}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractSession) NextConfig() (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	return _ConfigContract.Contract.NextConfig(&_ConfigContract.CallOpts)
}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint64 startBatchIndex, uint64 startBlockNumber, uint64 threshold, uint64 batchSpan, uint64 batchSizeLimit, uint64 transactionSizeLimit, uint64 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint64 executionTimeout)
func (_ConfigContract *ConfigContractCallerSession) NextConfig() (struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       uint64
}, error) {
	return _ConfigContract.Contract.NextConfig(&_ConfigContract.CallOpts)
}

// NextConfigBatchSizeLimit is a free data retrieval call binding the contract method 0xf464ba63.
//
// Solidity: function nextConfigBatchSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigBatchSizeLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigBatchSizeLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigBatchSizeLimit is a free data retrieval call binding the contract method 0xf464ba63.
//
// Solidity: function nextConfigBatchSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigBatchSizeLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigBatchSizeLimit(&_ConfigContract.CallOpts)
}

// NextConfigBatchSizeLimit is a free data retrieval call binding the contract method 0xf464ba63.
//
// Solidity: function nextConfigBatchSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigBatchSizeLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigBatchSizeLimit(&_ConfigContract.CallOpts)
}

// NextConfigBatchSpan is a free data retrieval call binding the contract method 0x6a53e51e.
//
// Solidity: function nextConfigBatchSpan() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigBatchSpan(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigBatchSpan")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigBatchSpan is a free data retrieval call binding the contract method 0x6a53e51e.
//
// Solidity: function nextConfigBatchSpan() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigBatchSpan() (uint64, error) {
	return _ConfigContract.Contract.NextConfigBatchSpan(&_ConfigContract.CallOpts)
}

// NextConfigBatchSpan is a free data retrieval call binding the contract method 0x6a53e51e.
//
// Solidity: function nextConfigBatchSpan() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigBatchSpan() (uint64, error) {
	return _ConfigContract.Contract.NextConfigBatchSpan(&_ConfigContract.CallOpts)
}

// NextConfigExecutionTimeout is a free data retrieval call binding the contract method 0x2b6cfe68.
//
// Solidity: function nextConfigExecutionTimeout() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigExecutionTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigExecutionTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigExecutionTimeout is a free data retrieval call binding the contract method 0x2b6cfe68.
//
// Solidity: function nextConfigExecutionTimeout() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigExecutionTimeout() (uint64, error) {
	return _ConfigContract.Contract.NextConfigExecutionTimeout(&_ConfigContract.CallOpts)
}

// NextConfigExecutionTimeout is a free data retrieval call binding the contract method 0x2b6cfe68.
//
// Solidity: function nextConfigExecutionTimeout() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigExecutionTimeout() (uint64, error) {
	return _ConfigContract.Contract.NextConfigExecutionTimeout(&_ConfigContract.CallOpts)
}

// NextConfigFeeReceiver is a free data retrieval call binding the contract method 0xdb293e51.
//
// Solidity: function nextConfigFeeReceiver() view returns(address)
func (_ConfigContract *ConfigContractCaller) NextConfigFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextConfigFeeReceiver is a free data retrieval call binding the contract method 0xdb293e51.
//
// Solidity: function nextConfigFeeReceiver() view returns(address)
func (_ConfigContract *ConfigContractSession) NextConfigFeeReceiver() (common.Address, error) {
	return _ConfigContract.Contract.NextConfigFeeReceiver(&_ConfigContract.CallOpts)
}

// NextConfigFeeReceiver is a free data retrieval call binding the contract method 0xdb293e51.
//
// Solidity: function nextConfigFeeReceiver() view returns(address)
func (_ConfigContract *ConfigContractCallerSession) NextConfigFeeReceiver() (common.Address, error) {
	return _ConfigContract.Contract.NextConfigFeeReceiver(&_ConfigContract.CallOpts)
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) NextConfigKeypers(opts *bind.CallOpts, keyperIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigKeypers", keyperIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) NextConfigKeypers(keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, keyperIndex)
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) NextConfigKeypers(keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, keyperIndex)
}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigNumKeypers(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigNumKeypers")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigNumKeypers() (uint64, error) {
	return _ConfigContract.Contract.NextConfigNumKeypers(&_ConfigContract.CallOpts)
}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigNumKeypers() (uint64, error) {
	return _ConfigContract.Contract.NextConfigNumKeypers(&_ConfigContract.CallOpts)
}

// NextConfigStartBatchIndex is a free data retrieval call binding the contract method 0xe0b70ed1.
//
// Solidity: function nextConfigStartBatchIndex() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigStartBatchIndex(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigStartBatchIndex")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigStartBatchIndex is a free data retrieval call binding the contract method 0xe0b70ed1.
//
// Solidity: function nextConfigStartBatchIndex() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigStartBatchIndex() (uint64, error) {
	return _ConfigContract.Contract.NextConfigStartBatchIndex(&_ConfigContract.CallOpts)
}

// NextConfigStartBatchIndex is a free data retrieval call binding the contract method 0xe0b70ed1.
//
// Solidity: function nextConfigStartBatchIndex() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigStartBatchIndex() (uint64, error) {
	return _ConfigContract.Contract.NextConfigStartBatchIndex(&_ConfigContract.CallOpts)
}

// NextConfigStartBlockNumber is a free data retrieval call binding the contract method 0xfda17adf.
//
// Solidity: function nextConfigStartBlockNumber() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigStartBlockNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigStartBlockNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigStartBlockNumber is a free data retrieval call binding the contract method 0xfda17adf.
//
// Solidity: function nextConfigStartBlockNumber() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigStartBlockNumber() (uint64, error) {
	return _ConfigContract.Contract.NextConfigStartBlockNumber(&_ConfigContract.CallOpts)
}

// NextConfigStartBlockNumber is a free data retrieval call binding the contract method 0xfda17adf.
//
// Solidity: function nextConfigStartBlockNumber() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigStartBlockNumber() (uint64, error) {
	return _ConfigContract.Contract.NextConfigStartBlockNumber(&_ConfigContract.CallOpts)
}

// NextConfigTargetAddress is a free data retrieval call binding the contract method 0xb99fd57c.
//
// Solidity: function nextConfigTargetAddress() view returns(address)
func (_ConfigContract *ConfigContractCaller) NextConfigTargetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigTargetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextConfigTargetAddress is a free data retrieval call binding the contract method 0xb99fd57c.
//
// Solidity: function nextConfigTargetAddress() view returns(address)
func (_ConfigContract *ConfigContractSession) NextConfigTargetAddress() (common.Address, error) {
	return _ConfigContract.Contract.NextConfigTargetAddress(&_ConfigContract.CallOpts)
}

// NextConfigTargetAddress is a free data retrieval call binding the contract method 0xb99fd57c.
//
// Solidity: function nextConfigTargetAddress() view returns(address)
func (_ConfigContract *ConfigContractCallerSession) NextConfigTargetAddress() (common.Address, error) {
	return _ConfigContract.Contract.NextConfigTargetAddress(&_ConfigContract.CallOpts)
}

// NextConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xf7ff384a.
//
// Solidity: function nextConfigTargetFunctionSelector() view returns(bytes4)
func (_ConfigContract *ConfigContractCaller) NextConfigTargetFunctionSelector(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigTargetFunctionSelector")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// NextConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xf7ff384a.
//
// Solidity: function nextConfigTargetFunctionSelector() view returns(bytes4)
func (_ConfigContract *ConfigContractSession) NextConfigTargetFunctionSelector() ([4]byte, error) {
	return _ConfigContract.Contract.NextConfigTargetFunctionSelector(&_ConfigContract.CallOpts)
}

// NextConfigTargetFunctionSelector is a free data retrieval call binding the contract method 0xf7ff384a.
//
// Solidity: function nextConfigTargetFunctionSelector() view returns(bytes4)
func (_ConfigContract *ConfigContractCallerSession) NextConfigTargetFunctionSelector() ([4]byte, error) {
	return _ConfigContract.Contract.NextConfigTargetFunctionSelector(&_ConfigContract.CallOpts)
}

// NextConfigThreshold is a free data retrieval call binding the contract method 0x9c7cf8e3.
//
// Solidity: function nextConfigThreshold() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigThreshold is a free data retrieval call binding the contract method 0x9c7cf8e3.
//
// Solidity: function nextConfigThreshold() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigThreshold() (uint64, error) {
	return _ConfigContract.Contract.NextConfigThreshold(&_ConfigContract.CallOpts)
}

// NextConfigThreshold is a free data retrieval call binding the contract method 0x9c7cf8e3.
//
// Solidity: function nextConfigThreshold() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigThreshold() (uint64, error) {
	return _ConfigContract.Contract.NextConfigThreshold(&_ConfigContract.CallOpts)
}

// NextConfigTransactionGasLimit is a free data retrieval call binding the contract method 0x175cbda9.
//
// Solidity: function nextConfigTransactionGasLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigTransactionGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigTransactionGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigTransactionGasLimit is a free data retrieval call binding the contract method 0x175cbda9.
//
// Solidity: function nextConfigTransactionGasLimit() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigTransactionGasLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigTransactionGasLimit(&_ConfigContract.CallOpts)
}

// NextConfigTransactionGasLimit is a free data retrieval call binding the contract method 0x175cbda9.
//
// Solidity: function nextConfigTransactionGasLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigTransactionGasLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigTransactionGasLimit(&_ConfigContract.CallOpts)
}

// NextConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xeb8b5209.
//
// Solidity: function nextConfigTransactionSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NextConfigTransactionSizeLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigTransactionSizeLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xeb8b5209.
//
// Solidity: function nextConfigTransactionSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NextConfigTransactionSizeLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigTransactionSizeLimit(&_ConfigContract.CallOpts)
}

// NextConfigTransactionSizeLimit is a free data retrieval call binding the contract method 0xeb8b5209.
//
// Solidity: function nextConfigTransactionSizeLimit() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NextConfigTransactionSizeLimit() (uint64, error) {
	return _ConfigContract.Contract.NextConfigTransactionSizeLimit(&_ConfigContract.CallOpts)
}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint64)
func (_ConfigContract *ConfigContractCaller) NumConfigs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "numConfigs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint64)
func (_ConfigContract *ConfigContractSession) NumConfigs() (uint64, error) {
	return _ConfigContract.Contract.NumConfigs(&_ConfigContract.CallOpts)
}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) NumConfigs() (uint64, error) {
	return _ConfigContract.Contract.NumConfigs(&_ConfigContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConfigContract *ConfigContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConfigContract *ConfigContractSession) Owner() (common.Address, error) {
	return _ConfigContract.Contract.Owner(&_ConfigContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConfigContract *ConfigContractCallerSession) Owner() (common.Address, error) {
	return _ConfigContract.Contract.Owner(&_ConfigContract.CallOpts)
}

// NextConfigAddKeypers is a paid mutator transaction binding the contract method 0x62fced0e.
//
// Solidity: function nextConfigAddKeypers(address[] newKeypers) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigAddKeypers(opts *bind.TransactOpts, newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigAddKeypers", newKeypers)
}

// NextConfigAddKeypers is a paid mutator transaction binding the contract method 0x62fced0e.
//
// Solidity: function nextConfigAddKeypers(address[] newKeypers) returns()
func (_ConfigContract *ConfigContractSession) NextConfigAddKeypers(newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigAddKeypers(&_ConfigContract.TransactOpts, newKeypers)
}

// NextConfigAddKeypers is a paid mutator transaction binding the contract method 0x62fced0e.
//
// Solidity: function nextConfigAddKeypers(address[] newKeypers) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigAddKeypers(newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigAddKeypers(&_ConfigContract.TransactOpts, newKeypers)
}

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x9d63753e.
//
// Solidity: function nextConfigRemoveKeypers(uint64 n) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigRemoveKeypers(opts *bind.TransactOpts, n uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigRemoveKeypers", n)
}

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x9d63753e.
//
// Solidity: function nextConfigRemoveKeypers(uint64 n) returns()
func (_ConfigContract *ConfigContractSession) NextConfigRemoveKeypers(n uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigRemoveKeypers(&_ConfigContract.TransactOpts, n)
}

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x9d63753e.
//
// Solidity: function nextConfigRemoveKeypers(uint64 n) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigRemoveKeypers(n uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigRemoveKeypers(&_ConfigContract.TransactOpts, n)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7c6e9f4.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint64 batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSizeLimit(opts *bind.TransactOpts, batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSizeLimit", batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7c6e9f4.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint64 batchSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSizeLimit(batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7c6e9f4.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint64 batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSizeLimit(batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, batchSizeLimit)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 batchSpan) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSpan(opts *bind.TransactOpts, batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSpan", batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 batchSpan) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSpan(batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 batchSpan) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSpan(batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, batchSpan)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetExecutionTimeout(opts *bind.TransactOpts, executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetExecutionTimeout", executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 executionTimeout) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetExecutionTimeout(executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetExecutionTimeout(&_ConfigContract.TransactOpts, executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetExecutionTimeout(executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetExecutionTimeout(&_ConfigContract.TransactOpts, executionTimeout)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address feeReceiver) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetFeeReceiver(opts *bind.TransactOpts, feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetFeeReceiver", feeReceiver)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address feeReceiver) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetFeeReceiver(feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetFeeReceiver(&_ConfigContract.TransactOpts, feeReceiver)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address feeReceiver) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetFeeReceiver(feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetFeeReceiver(&_ConfigContract.TransactOpts, feeReceiver)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBatchIndex(opts *bind.TransactOpts, startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBatchIndex", startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 startBatchIndex) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBatchIndex(startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBatchIndex(startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, startBatchIndex)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBlockNumber(opts *bind.TransactOpts, startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBlockNumber", startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 startBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBlockNumber(startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBlockNumber(&_ConfigContract.TransactOpts, startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBlockNumber(startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBlockNumber(&_ConfigContract.TransactOpts, startBlockNumber)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address targetAddress) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTargetAddress(opts *bind.TransactOpts, targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTargetAddress", targetAddress)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address targetAddress) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTargetAddress(targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetAddress(&_ConfigContract.TransactOpts, targetAddress)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address targetAddress) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTargetAddress(targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetAddress(&_ConfigContract.TransactOpts, targetAddress)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTargetFunctionSelector(opts *bind.TransactOpts, targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTargetFunctionSelector", targetFunctionSelector)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTargetFunctionSelector(targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetFunctionSelector(&_ConfigContract.TransactOpts, targetFunctionSelector)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTargetFunctionSelector(targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetFunctionSelector(&_ConfigContract.TransactOpts, targetFunctionSelector)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 threshold) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetThreshold(opts *bind.TransactOpts, threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetThreshold", threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 threshold) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetThreshold(threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 threshold) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetThreshold(threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, threshold)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionGasLimit(opts *bind.TransactOpts, transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionGasLimit", transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 transactionGasLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionGasLimit(transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionGasLimit(transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, transactionGasLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionSizeLimit(opts *bind.TransactOpts, transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionSizeLimit", transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionSizeLimit(transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionSizeLimit(&_ConfigContract.TransactOpts, transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionSizeLimit(transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionSizeLimit(&_ConfigContract.TransactOpts, transactionSizeLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConfigContract *ConfigContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConfigContract *ConfigContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConfigContract.Contract.RenounceOwnership(&_ConfigContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConfigContract *ConfigContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConfigContract.Contract.RenounceOwnership(&_ConfigContract.TransactOpts)
}

// ScheduleNextConfig is a paid mutator transaction binding the contract method 0x18b5e830.
//
// Solidity: function scheduleNextConfig() returns()
func (_ConfigContract *ConfigContractTransactor) ScheduleNextConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "scheduleNextConfig")
}

// ScheduleNextConfig is a paid mutator transaction binding the contract method 0x18b5e830.
//
// Solidity: function scheduleNextConfig() returns()
func (_ConfigContract *ConfigContractSession) ScheduleNextConfig() (*types.Transaction, error) {
	return _ConfigContract.Contract.ScheduleNextConfig(&_ConfigContract.TransactOpts)
}

// ScheduleNextConfig is a paid mutator transaction binding the contract method 0x18b5e830.
//
// Solidity: function scheduleNextConfig() returns()
func (_ConfigContract *ConfigContractTransactorSession) ScheduleNextConfig() (*types.Transaction, error) {
	return _ConfigContract.Contract.ScheduleNextConfig(&_ConfigContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConfigContract *ConfigContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConfigContract *ConfigContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.TransferOwnership(&_ConfigContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConfigContract *ConfigContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.TransferOwnership(&_ConfigContract.TransactOpts, newOwner)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xc9515c58.
//
// Solidity: function unscheduleConfigs(uint64 fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) UnscheduleConfigs(opts *bind.TransactOpts, fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "unscheduleConfigs", fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xc9515c58.
//
// Solidity: function unscheduleConfigs(uint64 fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) UnscheduleConfigs(fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.UnscheduleConfigs(&_ConfigContract.TransactOpts, fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xc9515c58.
//
// Solidity: function unscheduleConfigs(uint64 fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) UnscheduleConfigs(fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.UnscheduleConfigs(&_ConfigContract.TransactOpts, fromStartBlockNumber)
}

// ConfigContractConfigScheduledIterator is returned from FilterConfigScheduled and is used to iterate over the raw logs and unpacked data for ConfigScheduled events raised by the ConfigContract contract.
type ConfigContractConfigScheduledIterator struct {
	Event *ConfigContractConfigScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfigContractConfigScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigContractConfigScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfigContractConfigScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfigContractConfigScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigContractConfigScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigContractConfigScheduled represents a ConfigScheduled event raised by the ConfigContract contract.
type ConfigContractConfigScheduled struct {
	NumConfigs uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigScheduled is a free log retrieval operation binding the contract event 0x38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae6.
//
// Solidity: event ConfigScheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) FilterConfigScheduled(opts *bind.FilterOpts) (*ConfigContractConfigScheduledIterator, error) {

	logs, sub, err := _ConfigContract.contract.FilterLogs(opts, "ConfigScheduled")
	if err != nil {
		return nil, err
	}
	return &ConfigContractConfigScheduledIterator{contract: _ConfigContract.contract, event: "ConfigScheduled", logs: logs, sub: sub}, nil
}

// WatchConfigScheduled is a free log subscription operation binding the contract event 0x38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae6.
//
// Solidity: event ConfigScheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) WatchConfigScheduled(opts *bind.WatchOpts, sink chan<- *ConfigContractConfigScheduled) (event.Subscription, error) {

	logs, sub, err := _ConfigContract.contract.WatchLogs(opts, "ConfigScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigContractConfigScheduled)
				if err := _ConfigContract.contract.UnpackLog(event, "ConfigScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigScheduled is a log parse operation binding the contract event 0x38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae6.
//
// Solidity: event ConfigScheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) ParseConfigScheduled(log types.Log) (*ConfigContractConfigScheduled, error) {
	event := new(ConfigContractConfigScheduled)
	if err := _ConfigContract.contract.UnpackLog(event, "ConfigScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigContractConfigUnscheduledIterator is returned from FilterConfigUnscheduled and is used to iterate over the raw logs and unpacked data for ConfigUnscheduled events raised by the ConfigContract contract.
type ConfigContractConfigUnscheduledIterator struct {
	Event *ConfigContractConfigUnscheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfigContractConfigUnscheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigContractConfigUnscheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfigContractConfigUnscheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfigContractConfigUnscheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigContractConfigUnscheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigContractConfigUnscheduled represents a ConfigUnscheduled event raised by the ConfigContract contract.
type ConfigContractConfigUnscheduled struct {
	NumConfigs uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUnscheduled is a free log retrieval operation binding the contract event 0x202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf.
//
// Solidity: event ConfigUnscheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) FilterConfigUnscheduled(opts *bind.FilterOpts) (*ConfigContractConfigUnscheduledIterator, error) {

	logs, sub, err := _ConfigContract.contract.FilterLogs(opts, "ConfigUnscheduled")
	if err != nil {
		return nil, err
	}
	return &ConfigContractConfigUnscheduledIterator{contract: _ConfigContract.contract, event: "ConfigUnscheduled", logs: logs, sub: sub}, nil
}

// WatchConfigUnscheduled is a free log subscription operation binding the contract event 0x202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf.
//
// Solidity: event ConfigUnscheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) WatchConfigUnscheduled(opts *bind.WatchOpts, sink chan<- *ConfigContractConfigUnscheduled) (event.Subscription, error) {

	logs, sub, err := _ConfigContract.contract.WatchLogs(opts, "ConfigUnscheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigContractConfigUnscheduled)
				if err := _ConfigContract.contract.UnpackLog(event, "ConfigUnscheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigUnscheduled is a log parse operation binding the contract event 0x202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf.
//
// Solidity: event ConfigUnscheduled(uint64 numConfigs)
func (_ConfigContract *ConfigContractFilterer) ParseConfigUnscheduled(log types.Log) (*ConfigContractConfigUnscheduled, error) {
	event := new(ConfigContractConfigUnscheduled)
	if err := _ConfigContract.contract.UnpackLog(event, "ConfigUnscheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfigContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConfigContract contract.
type ConfigContractOwnershipTransferredIterator struct {
	Event *ConfigContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfigContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfigContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfigContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigContractOwnershipTransferred represents a OwnershipTransferred event raised by the ConfigContract contract.
type ConfigContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConfigContract *ConfigContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConfigContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConfigContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConfigContractOwnershipTransferredIterator{contract: _ConfigContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConfigContract *ConfigContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfigContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConfigContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigContractOwnershipTransferred)
				if err := _ConfigContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConfigContract *ConfigContractFilterer) ParseOwnershipTransferred(log types.Log) (*ConfigContractOwnershipTransferred, error) {
	event := new(ConfigContractOwnershipTransferred)
	if err := _ConfigContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DepositContractABI is the input ABI used to generate the binding from.
const DepositContractABI = "[{\"inputs\":[{\"internalType\":\"contractIERC777\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalDelayBlocks\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalRequestedBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"name\":\"DepositChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWithdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequestedBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slasherAddress\",\"type\":\"address\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC777\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositContractBin is the compiled bytecode used for deploying new contracts.
var DepositContractBin = "0x6080604052600080546001600160a01b031916731820a4b7618bde71dce8cdc73aab6c95905fad2417905534801561003657600080fd5b50604051610e16380380610e16833981016040819052610055916100fc565b600180546001600160a01b0319166001600160a01b03838116919091179091556000546040516329965a1d60e01b815230600482018190527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b602483015260448201529116906329965a1d90606401600060405180830381600087803b1580156100de57600080fd5b505af11580156100f2573d6000803e3d6000fd5b505050505061012a565b60006020828403121561010d578081fd5b81516001600160a01b0381168114610123578182fd5b9392505050565b610cdd806101396000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c8063b611d8d911610071578063b611d8d914610165578063b799036c146101a1578063b8ba16fd146101e7578063c96be4cb1461021e578063dbaf214514610231578063fc0c546a1461023957600080fd5b806223de29146100ad57806351cff8d9146100c25780639b4fed88146100d5578063aabc249614610127578063b13442711461013a575b600080fd5b6100c06100bb366004610b2f565b61024c565b005b6100c06100d0366004610b0e565b6102e1565b61010a6100e3366004610b0e565b6001600160a01b03166000908152600360205260409020600101546001600160401b031690565b6040516001600160401b0390911681526020015b60405180910390f35b6100c0610135366004610b0e565b610582565b60025461014d906001600160a01b031681565b6040516001600160a01b03909116815260200161011e565b61010a610173366004610b0e565b6001600160a01b0316600090815260036020526040902060010154600160401b90046001600160401b031690565b6101d76101af366004610b0e565b6001600160a01b0316600090815260036020526040902060010154600160801b900460ff1690565b604051901515815260200161011e565b6102106101f5366004610b0e565b6001600160a01b031660009081526003602052604090205490565b60405190815260200161011e565b6100c061022c366004610b0e565b610612565b6100c06106e9565b60015461014d906001600160a01b031681565b6001546001600160a01b031633146102bb5760405162461bcd60e51b815260206004820152602760248201527f4465706f736974436f6e74726163743a20726563656976656420696e76616c6960448201526632103a37b5b2b760c91b60648201526084015b60405180910390fd5b60006102c984860186610bd6565b90506102d6888783610897565b505050505050505050565b33600090815260036020908152604091829020825160808101845281548082526001909201546001600160401b0380821694830194909452600160401b810490931693810193909352600160801b90910460ff16151560608301526103885760405162461bcd60e51b815260206004820152601b60248201527f4465706f736974436f6e74726163743a206e6f206465706f736974000000000060448201526064016102b2565b8060600151156103a857634e487b7160e01b600052600160045260246000fd5b600081604001516001600160401b03161161041b5760405162461bcd60e51b815260206004820152602d60248201527f4465706f736974436f6e74726163743a207769746864726177616c206e6f742060448201526c1c995c5d595cdd1959081e595d609a1b60648201526084016102b2565b8060200151816040015161042f9190610c46565b6001600160401b03164310156104a05760405162461bcd60e51b815260206004820152603060248201527f4465706f736974436f6e74726163743a207769746864726177616c2064656c6160448201526f1e481b9bdd081c185cdcd959081e595d60821b60648201526084016102b2565b336000908152600360205260408082208281556001908101805470ffffffffffffffffffffffffffffffffff191690555483519151634decdde360e11b81526001600160a01b03868116600483015260248201939093526060604482015260648101939093521690639bd9bbc690608401600060405180830381600087803b15801561052b57600080fd5b505af115801561053f573d6000803e3d6000fd5b50505050336001600160a01b0316600080516020610c88833981519152600080600060016000604051610576959493929190610bfd565b60405180910390a25050565b6002546001600160a01b0316156105f05760405162461bcd60e51b815260206004820152602c60248201527f4465706f736974436f6e74726163743a20736c6173686572206164647265737360448201526b08185b1c9958591e481cd95d60a21b60648201526084016102b2565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001600160a01b0316331461062957600080fd5b6001600160a01b03811660008181526003602081815260408084208151608081018352600180830180548884528387018981528487018a8152606086018581528c8c529990985284519095559351955196511515600160801b0260ff60801b196001600160401b03988916600160401b026fffffffffffffffffffffffffffffffff199096169790981696909617939093179590951693909317905551909392600080516020610c88833981519152926105769282918291829190610bfd565b33600090815260036020908152604091829020825160808101845281548082526001909201546001600160401b0380821694830194909452600160401b810490931693810193909352600160801b90910460ff16151560608301526107905760405162461bcd60e51b815260206004820152601b60248201527f4465706f736974436f6e74726163743a206e6f206465706f736974000000000060448201526064016102b2565b8060600151156107b057634e487b7160e01b600052600160045260246000fd5b60408101516001600160401b0316156108215760405162461bcd60e51b815260206004820152602d60248201527f4465706f736974436f6e74726163743a207769746864726177616c20616c726560448201526c18591e481c995c5d595cdd1959609a1b60648201526084016102b2565b33600081815260036020908152604080832060010180546001600160401b0343908116600160401b026fffffffffffffffff000000000000000019909216919091179091558551928601519151600080516020610c888339815191529461088c949392918190610bfd565b60405180910390a250565b6001600160a01b0383166000908152600360209081526040918290208251608081018452815481526001909101546001600160401b03808216938301849052600160401b8204811694830194909452600160801b900460ff161515606082015291831610156109665760405162461bcd60e51b815260206004820152603560248201527f4465706f736974436f6e74726163743a207769746864726177616c2064656c616044820152741e4818d85b9b9bdd08189948191958dc99585cd959605a1b60648201526084016102b2565b60408101516001600160401b0316156109d15760405162461bcd60e51b815260206004820152602760248201527f4465706f736974436f6e74726163743a207769746864726177616c20696e2070604482015266726f677265737360c81b60648201526084016102b2565b806060015115610a235760405162461bcd60e51b815260206004820181905260248201527f4465706f736974436f6e74726163743a206163636f756e7420736c617368656460448201526064016102b2565b8051610a30908490610c2e565b6001600160a01b03851660008181526003602052604090209182556001909101805467ffffffffffffffff19166001600160401b0385161790558151600080516020610c8883398151915290610a87908690610c2e565b846000806000604051610a9e959493929190610bfd565b60405180910390a250505050565b80356001600160a01b0381168114610ac357600080fd5b919050565b60008083601f840112610ad9578182fd5b5081356001600160401b03811115610aef578182fd5b602083019150836020828501011115610b0757600080fd5b9250929050565b600060208284031215610b1f578081fd5b610b2882610aac565b9392505050565b60008060008060008060008060c0898b031215610b4a578384fd5b610b5389610aac565b9750610b6160208a01610aac565b9650610b6f60408a01610aac565b95506060890135945060808901356001600160401b0380821115610b91578586fd5b610b9d8c838d01610ac8565b909650945060a08b0135915080821115610bb5578384fd5b50610bc28b828c01610ac8565b999c989b5096995094979396929594505050565b600060208284031215610be7578081fd5b81356001600160401b0381168114610b28578182fd5b9485526001600160401b03938416602086015291909216604084015290151560608301521515608082015260a00190565b60008219821115610c4157610c41610c71565b500190565b60006001600160401b03808316818516808303821115610c6857610c68610c71565b01949350505050565b634e487b7160e01b600052601160045260246000fdfe04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97fa2646970667358221220689977c95f5f0d0cf9217d7086cf54558d16309112515e27e3291a0b9f0caca664736f6c63430008040033"

// DeployDepositContract deploys a new Ethereum contract, binding an instance of DepositContract to it.
func DeployDepositContract(auth *bind.TransactOpts, backend bind.ContractBackend, tokenContract common.Address) (common.Address, *types.Transaction, *DepositContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositContractBin), backend, tokenContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositContract{DepositContractCaller: DepositContractCaller{contract: contract}, DepositContractTransactor: DepositContractTransactor{contract: contract}, DepositContractFilterer: DepositContractFilterer{contract: contract}}, nil
}

// DepositContract is an auto generated Go binding around an Ethereum contract.
type DepositContract struct {
	DepositContractCaller     // Read-only binding to the contract
	DepositContractTransactor // Write-only binding to the contract
	DepositContractFilterer   // Log filterer for contract events
}

// DepositContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositContractSession struct {
	Contract     *DepositContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositContractCallerSession struct {
	Contract *DepositContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DepositContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositContractTransactorSession struct {
	Contract     *DepositContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DepositContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositContractRaw struct {
	Contract *DepositContract // Generic contract binding to access the raw methods on
}

// DepositContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositContractCallerRaw struct {
	Contract *DepositContractCaller // Generic read-only contract binding to access the raw methods on
}

// DepositContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositContractTransactorRaw struct {
	Contract *DepositContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositContract creates a new instance of DepositContract, bound to a specific deployed contract.
func NewDepositContract(address common.Address, backend bind.ContractBackend) (*DepositContract, error) {
	contract, err := bindDepositContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositContract{DepositContractCaller: DepositContractCaller{contract: contract}, DepositContractTransactor: DepositContractTransactor{contract: contract}, DepositContractFilterer: DepositContractFilterer{contract: contract}}, nil
}

// NewDepositContractCaller creates a new read-only instance of DepositContract, bound to a specific deployed contract.
func NewDepositContractCaller(address common.Address, caller bind.ContractCaller) (*DepositContractCaller, error) {
	contract, err := bindDepositContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositContractCaller{contract: contract}, nil
}

// NewDepositContractTransactor creates a new write-only instance of DepositContract, bound to a specific deployed contract.
func NewDepositContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositContractTransactor, error) {
	contract, err := bindDepositContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositContractTransactor{contract: contract}, nil
}

// NewDepositContractFilterer creates a new log filterer instance of DepositContract, bound to a specific deployed contract.
func NewDepositContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositContractFilterer, error) {
	contract, err := bindDepositContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositContractFilterer{contract: contract}, nil
}

// bindDepositContract binds a generic wrapper to an already deployed contract.
func bindDepositContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositContract *DepositContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositContract.Contract.DepositContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositContract *DepositContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositContract.Contract.DepositContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositContract *DepositContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositContract.Contract.DepositContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositContract *DepositContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositContract *DepositContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositContract *DepositContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositContract.Contract.contract.Transact(opts, method, params...)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address account) view returns(uint256)
func (_DepositContract *DepositContractCaller) GetDepositAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "getDepositAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address account) view returns(uint256)
func (_DepositContract *DepositContractSession) GetDepositAmount(account common.Address) (*big.Int, error) {
	return _DepositContract.Contract.GetDepositAmount(&_DepositContract.CallOpts, account)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address account) view returns(uint256)
func (_DepositContract *DepositContractCallerSession) GetDepositAmount(account common.Address) (*big.Int, error) {
	return _DepositContract.Contract.GetDepositAmount(&_DepositContract.CallOpts, account)
}

// GetWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x9b4fed88.
//
// Solidity: function getWithdrawalDelayBlocks(address account) view returns(uint64)
func (_DepositContract *DepositContractCaller) GetWithdrawalDelayBlocks(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "getWithdrawalDelayBlocks", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x9b4fed88.
//
// Solidity: function getWithdrawalDelayBlocks(address account) view returns(uint64)
func (_DepositContract *DepositContractSession) GetWithdrawalDelayBlocks(account common.Address) (uint64, error) {
	return _DepositContract.Contract.GetWithdrawalDelayBlocks(&_DepositContract.CallOpts, account)
}

// GetWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x9b4fed88.
//
// Solidity: function getWithdrawalDelayBlocks(address account) view returns(uint64)
func (_DepositContract *DepositContractCallerSession) GetWithdrawalDelayBlocks(account common.Address) (uint64, error) {
	return _DepositContract.Contract.GetWithdrawalDelayBlocks(&_DepositContract.CallOpts, account)
}

// GetWithdrawalRequestedBlock is a free data retrieval call binding the contract method 0xb611d8d9.
//
// Solidity: function getWithdrawalRequestedBlock(address account) view returns(uint64)
func (_DepositContract *DepositContractCaller) GetWithdrawalRequestedBlock(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "getWithdrawalRequestedBlock", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWithdrawalRequestedBlock is a free data retrieval call binding the contract method 0xb611d8d9.
//
// Solidity: function getWithdrawalRequestedBlock(address account) view returns(uint64)
func (_DepositContract *DepositContractSession) GetWithdrawalRequestedBlock(account common.Address) (uint64, error) {
	return _DepositContract.Contract.GetWithdrawalRequestedBlock(&_DepositContract.CallOpts, account)
}

// GetWithdrawalRequestedBlock is a free data retrieval call binding the contract method 0xb611d8d9.
//
// Solidity: function getWithdrawalRequestedBlock(address account) view returns(uint64)
func (_DepositContract *DepositContractCallerSession) GetWithdrawalRequestedBlock(account common.Address) (uint64, error) {
	return _DepositContract.Contract.GetWithdrawalRequestedBlock(&_DepositContract.CallOpts, account)
}

// IsSlashed is a free data retrieval call binding the contract method 0xb799036c.
//
// Solidity: function isSlashed(address account) view returns(bool)
func (_DepositContract *DepositContractCaller) IsSlashed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "isSlashed", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0xb799036c.
//
// Solidity: function isSlashed(address account) view returns(bool)
func (_DepositContract *DepositContractSession) IsSlashed(account common.Address) (bool, error) {
	return _DepositContract.Contract.IsSlashed(&_DepositContract.CallOpts, account)
}

// IsSlashed is a free data retrieval call binding the contract method 0xb799036c.
//
// Solidity: function isSlashed(address account) view returns(bool)
func (_DepositContract *DepositContractCallerSession) IsSlashed(account common.Address) (bool, error) {
	return _DepositContract.Contract.IsSlashed(&_DepositContract.CallOpts, account)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DepositContract *DepositContractCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DepositContract *DepositContractSession) Slasher() (common.Address, error) {
	return _DepositContract.Contract.Slasher(&_DepositContract.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DepositContract *DepositContractCallerSession) Slasher() (common.Address, error) {
	return _DepositContract.Contract.Slasher(&_DepositContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DepositContract *DepositContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DepositContract *DepositContractSession) Token() (common.Address, error) {
	return _DepositContract.Contract.Token(&_DepositContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DepositContract *DepositContractCallerSession) Token() (common.Address, error) {
	return _DepositContract.Contract.Token(&_DepositContract.CallOpts)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xdbaf2145.
//
// Solidity: function requestWithdrawal() returns()
func (_DepositContract *DepositContractTransactor) RequestWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "requestWithdrawal")
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xdbaf2145.
//
// Solidity: function requestWithdrawal() returns()
func (_DepositContract *DepositContractSession) RequestWithdrawal() (*types.Transaction, error) {
	return _DepositContract.Contract.RequestWithdrawal(&_DepositContract.TransactOpts)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xdbaf2145.
//
// Solidity: function requestWithdrawal() returns()
func (_DepositContract *DepositContractTransactorSession) RequestWithdrawal() (*types.Transaction, error) {
	return _DepositContract.Contract.RequestWithdrawal(&_DepositContract.TransactOpts)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasherAddress) returns()
func (_DepositContract *DepositContractTransactor) SetSlasher(opts *bind.TransactOpts, slasherAddress common.Address) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "setSlasher", slasherAddress)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasherAddress) returns()
func (_DepositContract *DepositContractSession) SetSlasher(slasherAddress common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.SetSlasher(&_DepositContract.TransactOpts, slasherAddress)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasherAddress) returns()
func (_DepositContract *DepositContractTransactorSession) SetSlasher(slasherAddress common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.SetSlasher(&_DepositContract.TransactOpts, slasherAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address account) returns()
func (_DepositContract *DepositContractTransactor) Slash(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "slash", account)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address account) returns()
func (_DepositContract *DepositContractSession) Slash(account common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.Slash(&_DepositContract.TransactOpts, account)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address account) returns()
func (_DepositContract *DepositContractTransactorSession) Slash(account common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.Slash(&_DepositContract.TransactOpts, account)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes userData, bytes ) returns()
func (_DepositContract *DepositContractTransactor) TokensReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, userData []byte, arg5 []byte) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "tokensReceived", arg0, from, arg2, amount, userData, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes userData, bytes ) returns()
func (_DepositContract *DepositContractSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, userData []byte, arg5 []byte) (*types.Transaction, error) {
	return _DepositContract.Contract.TokensReceived(&_DepositContract.TransactOpts, arg0, from, arg2, amount, userData, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes userData, bytes ) returns()
func (_DepositContract *DepositContractTransactorSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, userData []byte, arg5 []byte) (*types.Transaction, error) {
	return _DepositContract.Contract.TokensReceived(&_DepositContract.TransactOpts, arg0, from, arg2, amount, userData, arg5)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_DepositContract *DepositContractTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_DepositContract *DepositContractSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.Withdraw(&_DepositContract.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_DepositContract *DepositContractTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _DepositContract.Contract.Withdraw(&_DepositContract.TransactOpts, recipient)
}

// DepositContractDepositChangedIterator is returned from FilterDepositChanged and is used to iterate over the raw logs and unpacked data for DepositChanged events raised by the DepositContract contract.
type DepositContractDepositChangedIterator struct {
	Event *DepositContractDepositChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DepositContractDepositChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositContractDepositChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DepositContractDepositChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DepositContractDepositChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositContractDepositChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositContractDepositChanged represents a DepositChanged event raised by the DepositContract contract.
type DepositContractDepositChanged struct {
	Account                  common.Address
	Amount                   *big.Int
	WithdrawalDelayBlocks    uint64
	WithdrawalRequestedBlock uint64
	Withdrawn                bool
	Slashed                  bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDepositChanged is a free log retrieval operation binding the contract event 0x04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f.
//
// Solidity: event DepositChanged(address indexed account, uint256 amount, uint64 withdrawalDelayBlocks, uint64 withdrawalRequestedBlock, bool withdrawn, bool slashed)
func (_DepositContract *DepositContractFilterer) FilterDepositChanged(opts *bind.FilterOpts, account []common.Address) (*DepositContractDepositChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositContract.contract.FilterLogs(opts, "DepositChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &DepositContractDepositChangedIterator{contract: _DepositContract.contract, event: "DepositChanged", logs: logs, sub: sub}, nil
}

// WatchDepositChanged is a free log subscription operation binding the contract event 0x04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f.
//
// Solidity: event DepositChanged(address indexed account, uint256 amount, uint64 withdrawalDelayBlocks, uint64 withdrawalRequestedBlock, bool withdrawn, bool slashed)
func (_DepositContract *DepositContractFilterer) WatchDepositChanged(opts *bind.WatchOpts, sink chan<- *DepositContractDepositChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositContract.contract.WatchLogs(opts, "DepositChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositContractDepositChanged)
				if err := _DepositContract.contract.UnpackLog(event, "DepositChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositChanged is a log parse operation binding the contract event 0x04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f.
//
// Solidity: event DepositChanged(address indexed account, uint256 amount, uint64 withdrawalDelayBlocks, uint64 withdrawalRequestedBlock, bool withdrawn, bool slashed)
func (_DepositContract *DepositContractFilterer) ParseDepositChanged(log types.Log) (*DepositContractDepositChanged, error) {
	event := new(DepositContractDepositChanged)
	if err := _DepositContract.contract.UnpackLog(event, "DepositChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ba92e289188d2a25c4e0c3f2697ff3d2bd7c0df921f4094849ea79b9b099d56364736f6c63430008040033"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// ERC777ABI is the input ABI used to generate the binding from.
const ERC777ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"defaultOperators_\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC777Bin is the compiled bytecode used for deploying new contracts.
var ERC777Bin = "0x60806040523480156200001157600080fd5b5060405162001d4d38038062001d4d8339810160408190526200003491620003b0565b82516200004990600290602086019062000221565b5081516200005f90600390602085019062000221565b50805162000075906004906020840190620002b0565b5060005b8151811015620000f157600160056000848481518110620000aa57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580620000e88162000535565b91505062000079565b506040516329965a1d60e01b815230600482018190527fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce217705460248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad24906329965a1d90606401600060405180830381600087803b1580156200016c57600080fd5b505af115801562000181573d6000803e3d6000fd5b50506040516329965a1d60e01b815230600482018190527faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a60248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad2492506329965a1d9150606401600060405180830381600087803b158015620001ff57600080fd5b505af115801562000214573d6000803e3d6000fd5b5050505050505062000573565b8280546200022f90620004f8565b90600052602060002090601f0160209004810192826200025357600085556200029e565b82601f106200026e57805160ff19168380011785556200029e565b828001600101855582156200029e579182015b828111156200029e57825182559160200191906001019062000281565b50620002ac92915062000308565b5090565b8280548282559060005260206000209081019282156200029e579160200282015b828111156200029e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620002d1565b5b80821115620002ac576000815560010162000309565b600082601f83011262000330578081fd5b81516001600160401b038111156200034c576200034c6200055d565b602062000362601f8301601f19168201620004c5565b828152858284870101111562000376578384fd5b835b838110156200039557858101830151828201840152820162000378565b83811115620003a657848385840101525b5095945050505050565b600080600060608486031215620003c5578283fd5b83516001600160401b0380821115620003dc578485fd5b620003ea878388016200031f565b945060209150818601518181111562000401578485fd5b6200040f888289016200031f565b94505060408601518181111562000424578384fd5b8601601f8101881362000435578384fd5b8051828111156200044a576200044a6200055d565b8060051b92506200045d848401620004c5565b8181528481019083860185850187018c101562000478578788fd5b8795505b83861015620004b457805194506001600160a01b03851685146200049e578788fd5b848352600195909501949186019186016200047c565b508096505050505050509250925092565b604051601f8201601f191681016001600160401b0381118282101715620004f057620004f06200055d565b604052919050565b600181811c908216806200050d57607f821691505b602082108114156200052f57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156200055657634e487b7160e01b81526011600452602481fd5b5060010190565b634e487b7160e01b600052604160045260246000fd5b6117ca80620005836000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b63711461022b578063dd62ed3e1461023e578063fad8b32a14610277578063fc673c4f1461028a578063fe9d93031461029d57600080fd5b8063959b8c3f146101ea57806395d89b41146101fd5780639bd9bbc614610205578063a9059cbb1461021857600080fd5b806323b872dd116100e957806323b872dd14610183578063313ce56714610196578063556f0dc7146101a557806362ad1b83146101ac57806370a08231146101c157600080fd5b806306e485381461011b57806306fdde0314610139578063095ea7b31461014e57806318160ddd14610171575b600080fd5b6101236102b0565b60405161013091906115c1565b60405180910390f35b610141610312565b604051610130919061160e565b61016161015c3660046113d8565b61039b565b6040519015158152602001610130565b6001545b604051908152602001610130565b610161610191366004611308565b6103b3565b60405160128152602001610130565b6001610175565b6101bf6101ba366004611348565b61057c565b005b6101756101cf366004611298565b6001600160a01b031660009081526020819052604090205490565b6101bf6101f8366004611298565b6105b8565b6101416106d6565b6101bf610213366004611403565b6106e5565b6101616102263660046113d8565b610708565b6101616102393660046112d0565b6107bb565b61017561024c3660046112d0565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b6101bf610285366004611298565b61085d565b6101bf61029836600461145a565b610979565b6101bf6102ab3660046114d7565b6109b1565b6060600480548060200260200160405190810160405280929190818152602001828054801561030857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116102ea575b5050505050905090565b60606002805461032190611715565b80601f016020809104026020016040519081016040528092919081815260200182805461034d90611715565b80156103085780601f1061036f57610100808354040283529160200191610308565b820191906000526020600020905b81548152906001019060200180831161037d57509395945050505050565b6000336103a98185856109d0565b5060019392505050565b60006001600160a01b0383166103e45760405162461bcd60e51b81526004016103db90611621565b60405180910390fd5b6001600160a01b0384166104495760405162461bcd60e51b815260206004820152602660248201527f4552433737373a207472616e736665722066726f6d20746865207a65726f206160448201526564647265737360d01b60648201526084016103db565b600033905061047a818686866040518060200160405280600081525060405180602001604052806000815250610af7565b6104a6818686866040518060200160405280600081525060405180602001604052806000815250610c2e565b6001600160a01b038086166000908152600860209081526040808320938516835292905220548381101561052e5760405162461bcd60e51b815260206004820152602960248201527f4552433737373a207472616e7366657220616d6f756e74206578636565647320604482015268616c6c6f77616e636560b81b60648201526084016103db565b610542868361053d87856116fe565b6109d0565b6105708287878760405180602001604052806000815250604051806020016040528060008152506000610d9d565b50600195945050505050565b61058633866107bb565b6105a25760405162461bcd60e51b81526004016103db90611665565b6105b185858585856001610f71565b5050505050565b336001600160a01b038216141561061d5760405162461bcd60e51b8152602060048201526024808201527f4552433737373a20617574686f72697a696e672073656c66206173206f70657260448201526330ba37b960e11b60648201526084016103db565b6001600160a01b03811660009081526005602052604090205460ff161561066e573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff1916905561069d565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191660011790555b60405133906001600160a01b038316907ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f990600090a350565b60606003805461032190611715565b61070333848484604051806020016040528060008152506001610f71565b505050565b60006001600160a01b0383166107305760405162461bcd60e51b81526004016103db90611621565b6000339050610761818286866040518060200160405280600081525060405180602001604052806000815250610af7565b61078d818286866040518060200160405280600081525060405180602001604052806000815250610c2e565b6103a98182868660405180602001604052806000815250604051806020016040528060008152506000610d9d565b6000816001600160a01b0316836001600160a01b0316148061082657506001600160a01b03831660009081526005602052604090205460ff16801561082657506001600160a01b0380831660009081526007602090815260408083209387168352929052205460ff16155b8061085657506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff165b9392505050565b6001600160a01b0381163314156108c05760405162461bcd60e51b815260206004820152602160248201527f4552433737373a207265766f6b696e672073656c66206173206f70657261746f6044820152603960f91b60648201526084016103db565b6001600160a01b03811660009081526005602052604090205460ff1615610914573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610940565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191690555b60405133906001600160a01b038316907f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa190600090a350565b61098333856107bb565b61099f5760405162461bcd60e51b81526004016103db90611665565b6109ab84848484611054565b50505050565b6109cc33838360405180602001604052806000815250611054565b5050565b6001600160a01b038316610a345760405162461bcd60e51b815260206004820152602560248201527f4552433737373a20617070726f76652066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103db565b6001600160a01b038216610a965760405162461bcd60e51b815260206004820152602360248201527f4552433737373a20617070726f766520746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103db565b6001600160a01b0383811660008181526008602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe8956024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610b7357600080fd5b505afa158015610b87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bab91906112b4565b90506001600160a01b03811615610c2557604051633ad5cbc160e11b81526001600160a01b038216906375ab978290610bf2908a908a908a908a908a908a90600401611567565b600060405180830381600087803b158015610c0c57600080fd5b505af1158015610c20573d6000803e3d6000fd5b505050505b50505050505050565b6001600160a01b03851660009081526020819052604090205483811015610ca75760405162461bcd60e51b815260206004820152602760248201527f4552433737373a207472616e7366657220616d6f756e7420657863656564732060448201526662616c616e636560c81b60648201526084016103db565b610cb184826116fe565b6001600160a01b038088166000908152602081905260408082209390935590871681529081208054869290610ce79084906116e6565b92505081905550846001600160a01b0316866001600160a01b0316886001600160a01b03167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987878787604051610d3f939291906116b1565b60405180910390a4846001600160a01b0316866001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef86604051610d8c91815260200190565b60405180910390a350505050505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b6024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610e1957600080fd5b505afa158015610e2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5191906112b4565b90506001600160a01b03811615610ecd576040516223de2960e01b81526001600160a01b038216906223de2990610e96908b908b908b908b908b908b90600401611567565b600060405180830381600087803b158015610eb057600080fd5b505af1158015610ec4573d6000803e3d6000fd5b50505050610f67565b8115610f67576001600160a01b0386163b15610f675760405162461bcd60e51b815260206004820152604d60248201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460448201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60648201526c1ad95b9cd49958da5c1a595b9d609a1b608482015260a4016103db565b5050505050505050565b6001600160a01b038616610fd25760405162461bcd60e51b815260206004820152602260248201527f4552433737373a2073656e642066726f6d20746865207a65726f206164647265604482015261737360f01b60648201526084016103db565b6001600160a01b0385166110285760405162461bcd60e51b815260206004820181905260248201527f4552433737373a2073656e6420746f20746865207a65726f206164647265737360448201526064016103db565b33611037818888888888610af7565b611045818888888888610c2e565b610c2581888888888888610d9d565b6001600160a01b0384166110b55760405162461bcd60e51b815260206004820152602260248201527f4552433737373a206275726e2066726f6d20746865207a65726f206164647265604482015261737360f01b60648201526084016103db565b336110c581866000878787610af7565b6001600160a01b0385166000908152602081905260409020548481101561113a5760405162461bcd60e51b815260206004820152602360248201527f4552433737373a206275726e20616d6f756e7420657863656564732062616c616044820152626e636560e81b60648201526084016103db565b61114485826116fe565b6001600160a01b038716600090815260208190526040812091909155600180548792906111729084906116fe565b92505081905550856001600160a01b0316826001600160a01b03167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a40988787876040516111c0939291906116b1565b60405180910390a36040518581526000906001600160a01b038816907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050505050565b600082601f830112611221578081fd5b813567ffffffffffffffff8082111561123c5761123c611766565b604051601f8301601f19908116603f0116810190828211818310171561126457611264611766565b8160405283815286602085880101111561127c578485fd5b8360208701602083013792830160200193909352509392505050565b6000602082840312156112a9578081fd5b81356108568161177c565b6000602082840312156112c5578081fd5b81516108568161177c565b600080604083850312156112e2578081fd5b82356112ed8161177c565b915060208301356112fd8161177c565b809150509250929050565b60008060006060848603121561131c578081fd5b83356113278161177c565b925060208401356113378161177c565b929592945050506040919091013590565b600080600080600060a0868803121561135f578081fd5b853561136a8161177c565b9450602086013561137a8161177c565b935060408601359250606086013567ffffffffffffffff8082111561139d578283fd5b6113a989838a01611211565b935060808801359150808211156113be578283fd5b506113cb88828901611211565b9150509295509295909350565b600080604083850312156113ea578182fd5b82356113f58161177c565b946020939093013593505050565b600080600060608486031215611417578283fd5b83356114228161177c565b925060208401359150604084013567ffffffffffffffff811115611444578182fd5b61145086828701611211565b9150509250925092565b6000806000806080858703121561146f578384fd5b843561147a8161177c565b935060208501359250604085013567ffffffffffffffff8082111561149d578384fd5b6114a988838901611211565b935060608701359150808211156114be578283fd5b506114cb87828801611211565b91505092959194509250565b600080604083850312156114e9578182fd5b82359150602083013567ffffffffffffffff811115611506578182fd5b61151285828601611211565b9150509250929050565b60008151808452815b8181101561154157602081850181015186830182015201611525565b818111156115525782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c0608082018190526000906115a29083018561151c565b82810360a08401526115b4818561151c565b9998505050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156116025783516001600160a01b0316835292840192918401916001016115dd565b50909695505050505050565b602081526000610856602083018461151c565b60208082526024908201527f4552433737373a207472616e7366657220746f20746865207a65726f206164646040820152637265737360e01b606082015260800190565b6020808252602c908201527f4552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f60408201526b39103337b9103437b63232b960a11b606082015260800190565b8381526060602082015260006116ca606083018561151c565b82810360408401526116dc818561151c565b9695505050505050565b600082198211156116f9576116f9611750565b500190565b60008282101561171057611710611750565b500390565b600181811c9082168061172957607f821691505b6020821081141561174a57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461179157600080fd5b5056fea26469706673582212201846bdd21c79399e12f4e29ba2c3aa28b46234cc5885cba8bf960eede90e78d164736f6c63430008040033"

// DeployERC777 deploys a new Ethereum contract, binding an instance of ERC777 to it.
func DeployERC777(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, defaultOperators_ []common.Address) (common.Address, *types.Transaction, *ERC777, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC777ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC777Bin), backend, name_, symbol_, defaultOperators_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC777{ERC777Caller: ERC777Caller{contract: contract}, ERC777Transactor: ERC777Transactor{contract: contract}, ERC777Filterer: ERC777Filterer{contract: contract}}, nil
}

// ERC777 is an auto generated Go binding around an Ethereum contract.
type ERC777 struct {
	ERC777Caller     // Read-only binding to the contract
	ERC777Transactor // Write-only binding to the contract
	ERC777Filterer   // Log filterer for contract events
}

// ERC777Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC777Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC777Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC777Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC777Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC777Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC777Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC777Session struct {
	Contract     *ERC777           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC777CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC777CallerSession struct {
	Contract *ERC777Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC777TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC777TransactorSession struct {
	Contract     *ERC777Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC777Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC777Raw struct {
	Contract *ERC777 // Generic contract binding to access the raw methods on
}

// ERC777CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC777CallerRaw struct {
	Contract *ERC777Caller // Generic read-only contract binding to access the raw methods on
}

// ERC777TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC777TransactorRaw struct {
	Contract *ERC777Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC777 creates a new instance of ERC777, bound to a specific deployed contract.
func NewERC777(address common.Address, backend bind.ContractBackend) (*ERC777, error) {
	contract, err := bindERC777(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC777{ERC777Caller: ERC777Caller{contract: contract}, ERC777Transactor: ERC777Transactor{contract: contract}, ERC777Filterer: ERC777Filterer{contract: contract}}, nil
}

// NewERC777Caller creates a new read-only instance of ERC777, bound to a specific deployed contract.
func NewERC777Caller(address common.Address, caller bind.ContractCaller) (*ERC777Caller, error) {
	contract, err := bindERC777(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC777Caller{contract: contract}, nil
}

// NewERC777Transactor creates a new write-only instance of ERC777, bound to a specific deployed contract.
func NewERC777Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC777Transactor, error) {
	contract, err := bindERC777(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC777Transactor{contract: contract}, nil
}

// NewERC777Filterer creates a new log filterer instance of ERC777, bound to a specific deployed contract.
func NewERC777Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC777Filterer, error) {
	contract, err := bindERC777(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC777Filterer{contract: contract}, nil
}

// bindERC777 binds a generic wrapper to an already deployed contract.
func bindERC777(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC777ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC777 *ERC777Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC777.Contract.ERC777Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC777 *ERC777Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC777.Contract.ERC777Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC777 *ERC777Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC777.Contract.ERC777Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC777 *ERC777CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC777.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC777 *ERC777TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC777.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC777 *ERC777TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC777.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_ERC777 *ERC777Caller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_ERC777 *ERC777Session) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _ERC777.Contract.Allowance(&_ERC777.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_ERC777 *ERC777CallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _ERC777.Contract.Allowance(&_ERC777.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_ERC777 *ERC777Caller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_ERC777 *ERC777Session) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _ERC777.Contract.BalanceOf(&_ERC777.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_ERC777 *ERC777CallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _ERC777.Contract.BalanceOf(&_ERC777.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ERC777 *ERC777Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ERC777 *ERC777Session) Decimals() (uint8, error) {
	return _ERC777.Contract.Decimals(&_ERC777.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ERC777 *ERC777CallerSession) Decimals() (uint8, error) {
	return _ERC777.Contract.Decimals(&_ERC777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_ERC777 *ERC777Caller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_ERC777 *ERC777Session) DefaultOperators() ([]common.Address, error) {
	return _ERC777.Contract.DefaultOperators(&_ERC777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_ERC777 *ERC777CallerSession) DefaultOperators() ([]common.Address, error) {
	return _ERC777.Contract.DefaultOperators(&_ERC777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_ERC777 *ERC777Caller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_ERC777 *ERC777Session) Granularity() (*big.Int, error) {
	return _ERC777.Contract.Granularity(&_ERC777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_ERC777 *ERC777CallerSession) Granularity() (*big.Int, error) {
	return _ERC777.Contract.Granularity(&_ERC777.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_ERC777 *ERC777Caller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_ERC777 *ERC777Session) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _ERC777.Contract.IsOperatorFor(&_ERC777.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_ERC777 *ERC777CallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _ERC777.Contract.IsOperatorFor(&_ERC777.CallOpts, operator, tokenHolder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC777 *ERC777Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC777 *ERC777Session) Name() (string, error) {
	return _ERC777.Contract.Name(&_ERC777.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC777 *ERC777CallerSession) Name() (string, error) {
	return _ERC777.Contract.Name(&_ERC777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC777 *ERC777Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC777 *ERC777Session) Symbol() (string, error) {
	return _ERC777.Contract.Symbol(&_ERC777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC777 *ERC777CallerSession) Symbol() (string, error) {
	return _ERC777.Contract.Symbol(&_ERC777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC777 *ERC777Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC777.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC777 *ERC777Session) TotalSupply() (*big.Int, error) {
	return _ERC777.Contract.TotalSupply(&_ERC777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC777 *ERC777CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC777.Contract.TotalSupply(&_ERC777.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC777 *ERC777Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC777 *ERC777Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.Approve(&_ERC777.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC777 *ERC777TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.Approve(&_ERC777.TransactOpts, spender, value)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_ERC777 *ERC777Transactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_ERC777 *ERC777Session) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _ERC777.Contract.AuthorizeOperator(&_ERC777.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_ERC777 *ERC777TransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _ERC777.Contract.AuthorizeOperator(&_ERC777.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_ERC777 *ERC777Transactor) Burn(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "burn", amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_ERC777 *ERC777Session) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.Contract.Burn(&_ERC777.TransactOpts, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_ERC777 *ERC777TransactorSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.Contract.Burn(&_ERC777.TransactOpts, amount, data)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777Transactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777Session) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.Contract.OperatorBurn(&_ERC777.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777TransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.Contract.OperatorBurn(&_ERC777.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777Transactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777Session) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.Contract.OperatorSend(&_ERC777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_ERC777 *ERC777TransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _ERC777.Contract.OperatorSend(&_ERC777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_ERC777 *ERC777Transactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_ERC777 *ERC777Session) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _ERC777.Contract.RevokeOperator(&_ERC777.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_ERC777 *ERC777TransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _ERC777.Contract.RevokeOperator(&_ERC777.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_ERC777 *ERC777Transactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_ERC777 *ERC777Session) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.Contract.Send(&_ERC777.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_ERC777 *ERC777TransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC777.Contract.Send(&_ERC777.TransactOpts, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.Transfer(&_ERC777.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.Transfer(&_ERC777.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777Transactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777Session) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.TransferFrom(&_ERC777.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_ERC777 *ERC777TransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC777.Contract.TransferFrom(&_ERC777.TransactOpts, holder, recipient, amount)
}

// ERC777ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC777 contract.
type ERC777ApprovalIterator struct {
	Event *ERC777Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777Approval represents a Approval event raised by the ERC777 contract.
type ERC777Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC777 *ERC777Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC777ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC777ApprovalIterator{contract: _ERC777.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC777 *ERC777Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC777Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777Approval)
				if err := _ERC777.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC777 *ERC777Filterer) ParseApproval(log types.Log) (*ERC777Approval, error) {
	event := new(ERC777Approval)
	if err := _ERC777.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777AuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the ERC777 contract.
type ERC777AuthorizedOperatorIterator struct {
	Event *ERC777AuthorizedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777AuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777AuthorizedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777AuthorizedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777AuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777AuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777AuthorizedOperator represents a AuthorizedOperator event raised by the ERC777 contract.
type ERC777AuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*ERC777AuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &ERC777AuthorizedOperatorIterator{contract: _ERC777.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *ERC777AuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777AuthorizedOperator)
				if err := _ERC777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) ParseAuthorizedOperator(log types.Log) (*ERC777AuthorizedOperator, error) {
	event := new(ERC777AuthorizedOperator)
	if err := _ERC777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the ERC777 contract.
type ERC777BurnedIterator struct {
	Event *ERC777Burned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777Burned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777Burned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777Burned represents a Burned event raised by the ERC777 contract.
type ERC777Burned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*ERC777BurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC777BurnedIterator{contract: _ERC777.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *ERC777Burned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777Burned)
				if err := _ERC777.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) ParseBurned(log types.Log) (*ERC777Burned, error) {
	event := new(ERC777Burned)
	if err := _ERC777.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ERC777 contract.
type ERC777MintedIterator struct {
	Event *ERC777Minted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777Minted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777Minted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777Minted represents a Minted event raised by the ERC777 contract.
type ERC777Minted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*ERC777MintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC777MintedIterator{contract: _ERC777.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ERC777Minted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777Minted)
				if err := _ERC777.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) ParseMinted(log types.Log) (*ERC777Minted, error) {
	event := new(ERC777Minted)
	if err := _ERC777.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777RevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the ERC777 contract.
type ERC777RevokedOperatorIterator struct {
	Event *ERC777RevokedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777RevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777RevokedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777RevokedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777RevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777RevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777RevokedOperator represents a RevokedOperator event raised by the ERC777 contract.
type ERC777RevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*ERC777RevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &ERC777RevokedOperatorIterator{contract: _ERC777.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *ERC777RevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777RevokedOperator)
				if err := _ERC777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_ERC777 *ERC777Filterer) ParseRevokedOperator(log types.Log) (*ERC777RevokedOperator, error) {
	event := new(ERC777RevokedOperator)
	if err := _ERC777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777SentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the ERC777 contract.
type ERC777SentIterator struct {
	Event *ERC777Sent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777SentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777Sent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777Sent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777SentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777Sent represents a Sent event raised by the ERC777 contract.
type ERC777Sent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC777SentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC777SentIterator{contract: _ERC777.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) WatchSent(opts *bind.WatchOpts, sink chan<- *ERC777Sent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777Sent)
				if err := _ERC777.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_ERC777 *ERC777Filterer) ParseSent(log types.Log) (*ERC777Sent, error) {
	event := new(ERC777Sent)
	if err := _ERC777.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC777TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC777 contract.
type ERC777TransferIterator struct {
	Event *ERC777Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC777TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC777Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC777Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC777TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC777TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC777Transfer represents a Transfer event raised by the ERC777 contract.
type ERC777Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC777 *ERC777Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC777TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC777TransferIterator{contract: _ERC777.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC777 *ERC777Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC777Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC777.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC777Transfer)
				if err := _ERC777.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC777 *ERC777Filterer) ParseTransfer(log types.Log) (*ERC777Transfer, error) {
	event := new(ERC777Transfer)
	if err := _ERC777.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorContractABI is the input ABI used to generate the binding from.
const ExecutorContractABI = "[{\"inputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"configContractAddress\",\"type\":\"address\"},{\"internalType\":\"contractBatcherContract\",\"name\":\"batcherContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"}],\"name\":\"CipherExecutionSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"txIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"batcherContract\",\"outputs\":[{\"internalType\":\"contractBatcherContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"cipherExecutionReceipts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"keyperIndex\",\"type\":\"uint64\"}],\"name\":\"executeCipherBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"}],\"name\":\"executePlainBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCipherExecutionReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numExecutionHalfSteps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"skipCipherExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutorContractBin is the compiled bytecode used for deploying new contracts.
var ExecutorContractBin = "0x60806040523480156200001157600080fd5b5060405162001e1838038062001e18833981016040819052620000349162000066565b600080546001600160a01b039384166001600160a01b03199182161790915560018054929093169116179055620000bd565b6000806040838503121562000079578182fd5b82516200008681620000a4565b60208401519092506200009981620000a4565b809150509250929050565b6001600160a01b0381168114620000ba57600080fd5b50565b611d4b80620000cd6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063bf66a1821161005b578063bf66a182146101e1578063c87190b3146101f4578063d8e9a00114610288578063fa6385f41461029b57600080fd5b806325b36cbf1461008d5780634b2a026d1461018e57806387eb2378146101a3578063beb3b50e146101b6575b600080fd5b61013461009b3660046118e9565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160401b03908116600090815260026020818152604092839020835160a081018552815460ff8116151582526001600160a01b0361010082041693820193909352600160a81b90920490941692810192909252600183015460608301529190910154608082015290565b60405161018591908151151581526020808301516001600160a01b0316908201526040808301516001600160401b031690820152606080830151908201526080918201519181019190915260a00190565b60405180910390f35b6101a161019c366004611921565b6102cd565b005b6101a16101b1366004611973565b610866565b6001546101c9906001600160a01b031681565b6040516001600160a01b039091168152602001610185565b6000546101c9906001600160a01b031681565b6102486102023660046118e9565b600260208190526000918252604090912080546001820154919092015460ff83169261010081046001600160a01b031692600160a81b9091046001600160401b03169185565b6040805195151586526001600160a01b0390941660208601526001600160401b03909216928401929092526060830191909152608082015260a001610185565b6101a16102963660046118e9565b6112b5565b6001546102b590600160a01b90046001600160401b031681565b6040516001600160401b039091168152602001610185565b6001546001600160401b03808516916102f091600291600160a01b900416611c38565b6001600160401b03161461031f5760405162461bcd60e51b815260040161031690611b27565b60405180910390fd5b60015461033e90600290600160a01b90046001600160401b0316611cc1565b6001600160401b03166001146103665760405162461bcd60e51b815260040161031690611a9d565b60008054604051632e37abf360e21b81526001600160401b03861660048201526001600160a01b039091169063b8deafcc9060240160206040518083038186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103eb9190611905565b6000805460405163309787c960e21b81526001600160401b038416600482015292935090916001600160a01b039091169063c25e1f249060240160206040518083038186803b15801561043d57600080fd5b505afa158015610451573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610475919061185b565b60008054604051633b79c58160e21b81526001600160401b038616600482015292935090916001600160a01b039091169063ede716049060240160206040518083038186803b1580156104c757600080fd5b505afa1580156104db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ff91906118c1565b6000805460405163b5a43cef60e01b81526001600160401b038716600482015292935090916001600160a01b039091169063b5a43cef9060240160206040518083038186803b15801561055157600080fd5b505afa158015610565573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105899190611905565b6000546040516359d6accd60e01b81526001600160401b03871660048201529192506001600160a01b0316906359d6accd9060240160206040518083038186803b1580156105d657600080fd5b505afa1580156105ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060e9190611889565b61062857634e487b7160e01b600052600160045260246000fd5b600080546040516368080ea760e01b81526001600160401b0380881660048301528a1660248201526001600160a01b03909116906368080ea79060440160606040518083038186803b15801561067d57600080fd5b505afa158015610691573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b591906119e4565b50915050806001600160401b03164310156106e057634e487b7160e01b600052600160045260246000fd5b60006106ef8585858b8b6115f5565b6001805460405163643d7d4560e11b81529293506001600160a01b03169163c87afa8a91610722918d9190600401611bb9565b60206040518083038186803b15801561073a57600080fd5b505afa15801561074e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077291906118a9565b81146107d45760405162461bcd60e51b815260206004820152602b60248201527f4578656375746f72436f6e74726163743a206261746368206861736820646f6560448201526a0e640dcdee840dac2e8c6d60ab1b6064820152608401610316565b60018054600160a01b90046001600160401b03169060146107f483611c8e565b82546101009290920a6001600160401b0381810219909316918316021790915560015460408051600160a01b9092049092168152602081018490527f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a92500160405180910390a1505050505050505050565b6001546001600160401b038087169161088991600291600160a01b900416611c38565b6001600160401b0316146108af5760405162461bcd60e51b815260040161031690611b27565b6001546108ce90600290600160a01b90046001600160401b0316611cc1565b6001600160401b0316156108f45760405162461bcd60e51b815260040161031690611a9d565b60008054604051632e37abf360e21b81526001600160401b03881660048201526001600160a01b039091169063b8deafcc9060240160206040518083038186803b15801561094157600080fd5b505afa158015610955573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109799190611905565b6000805460405163309787c960e21b81526001600160401b038416600482015292935090916001600160a01b039091169063c25e1f249060240160206040518083038186803b1580156109cb57600080fd5b505afa1580156109df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a03919061185b565b60008054604051633b79c58160e21b81526001600160401b038616600482015292935090916001600160a01b039091169063ede716049060240160206040518083038186803b158015610a5557600080fd5b505afa158015610a69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8d91906118c1565b6000805460405163b5a43cef60e01b81526001600160401b038716600482015292935090916001600160a01b039091169063b5a43cef9060240160206040518083038186803b158015610adf57600080fd5b505afa158015610af3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b179190611905565b6000546040516359d6accd60e01b81526001600160401b03871660048201529192506001600160a01b0316906359d6accd9060240160206040518083038186803b158015610b6457600080fd5b505afa158015610b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9c9190611889565b610bb85760405162461bcd60e51b815260040161031690611ae3565b600080546040516368080ea760e01b81526001600160401b0380881660048301528c16602482015282916001600160a01b0316906368080ea79060440160606040518083038186803b158015610c0d57600080fd5b505afa158015610c21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4591906119e4565b9250925050806001600160401b03164310610ce65760018054600160a01b90046001600160401b0316906014610c7a83611c8e565b82546101009290920a6001600160401b03818102199093169183160217909155600154604051600160a01b90910490911681527fa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be915060200160405180910390a15050505050506112ae565b816001600160401b0316431015610d515760405162461bcd60e51b815260206004820152602960248201527f4578656375746f72436f6e74726163743a206261746368206973206e6f7420636044820152681b1bdcd959081e595d60ba1b6064820152608401610316565b6000805460405163366963c960e21b81526001600160401b03891660048201526001600160a01b039091169063d9a58f249060240160206040518083038186803b158015610d9e57600080fd5b505afa158015610db2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd69190611905565b9050806001600160401b0316886001600160401b031610610e4e5760405162461bcd60e51b815260206004820152602c60248201527f4578656375746f72436f6e74726163743a206b657970657220696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610316565b60008054604051637d42750160e11b81526001600160401b03808b1660048301528b1660248201526001600160a01b039091169063fa84ea029060440160206040518083038186803b158015610ea357600080fd5b505afa158015610eb7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610edb919061185b565b9050336001600160a01b03821614610f4e5760405162461bcd60e51b815260206004820152603060248201527f4578656375746f72436f6e74726163743a2073656e646572206973206e6f742060448201526f39b832b1b4b334b2b21035b2bcb832b960811b6064820152608401610316565b600160009054906101000a90046001600160a01b03166001600160a01b031663c87afa8a8e60006040518363ffffffff1660e01b8152600401610f92929190611bb9565b60206040518083038186803b158015610faa57600080fd5b505afa158015610fbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe291906118a9565b8c146110465760405162461bcd60e51b815260206004820152602d60248201527f4578656375746f72436f6e74726163743a20696e636f7272656374206369706860448201526c0cae440c4c2e8c6d040d0c2e6d609b1b6064820152608401610316565b8b158015611052575089155b8061106657508b1580159061106657508915155b6110eb5760405162461bcd60e51b815260206004820152604a60248201527f4578656375746f72436f6e74726163743a20636970686572426174636848617360448201527f682073686f756c64206265207a65726f20696666207472616e73616374696f6e6064820152697320697320656d70747960b01b608482015260a401610316565b60006110fa8888888f8f6115f5565b90506040518060a00160405280600115158152602001336001600160a01b03168152602001600160149054906101000a90046001600160401b03166001600160401b031681526020018e81526020018281525060026000600160149054906101000a90046001600160401b03166001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160010155608082015181600201559050506001601481819054906101000a90046001600160401b03168092919061123d90611c8e565b82546101009290920a6001600160401b0381810219909316918316021790915560015460408051600160a01b9092049092168152602081018490527f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a92500160405180910390a15050505050505050505b5050505050565b6001546001600160401b03808316916112d891600291600160a01b900416611c38565b6001600160401b0316146112fe5760405162461bcd60e51b815260040161031690611b27565b60015461131d90600290600160a01b90046001600160401b0316611cc1565b6001600160401b0316156113435760405162461bcd60e51b815260040161031690611a9d565b60008054604051632e37abf360e21b81526001600160401b03841660048201526001600160a01b039091169063b8deafcc9060240160206040518083038186803b15801561139057600080fd5b505afa1580156113a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c89190611905565b6000546040516359d6accd60e01b81526001600160401b03831660048201529192506001600160a01b0316906359d6accd9060240160206040518083038186803b15801561141557600080fd5b505afa158015611429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144d9190611889565b6114695760405162461bcd60e51b815260040161031690611ae3565b600080546040516368080ea760e01b81526001600160401b038085166004830152851660248201526001600160a01b03909116906368080ea79060440160606040518083038186803b1580156114be57600080fd5b505afa1580156114d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f691906119e4565b92505050806001600160401b031643101561156f5760405162461bcd60e51b815260206004820152603360248201527f4578656375746f72436f6e74726163743a20657865637574696f6e2074696d656044820152721bdd5d081b9bdd081c995858da1959081e595d606a1b6064820152608401610316565b60018054600160a01b90046001600160401b031690601461158f83611c8e565b82546101009290920a6001600160401b03818102199093169183160217909155600154604051600160a01b90910490911681527fa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be915060200160405180910390a1505050565b60008060005b6001600160401b038116841115611807576000878686846001600160401b031681811061163857634e487b7160e01b600052603260045260246000fd5b905060200281019061164a9190611bf4565b60405160240161165b929190611a6e565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b03838183161783525050505090506000808a6001600160a01b0316896001600160401b0316846040516116b59190611a52565b60006040518083038160008787f1925050503d80600081146116f3576040519150601f19603f3d011682016040523d82523d6000602084013e6116f8565b606091505b50915091508161178b577f6f580fd78fd3fc2e6db770c1da4b991b45906920b1fe2abd8cc7d760b8021049848989876001600160401b031681811061174d57634e487b7160e01b600052603260045260246000fd5b905060200281019061175f9190611bf4565b60405161176d929190611a30565b60405190819003812061178292918590611b6f565b60405180910390a15b8787856001600160401b03168181106117b457634e487b7160e01b600052603260045260246000fd5b90506020028101906117c69190611bf4565b866040516020016117d993929190611a40565b60405160208183030381529060405280519060200120945050505080806117ff90611c8e565b9150506115fb565b509695505050505050565b60008083601f840112611823578182fd5b5081356001600160401b03811115611839578182fd5b6020830191508360208260051b850101111561185457600080fd5b9250929050565b60006020828403121561186c578081fd5b81516001600160a01b0381168114611882578182fd5b9392505050565b60006020828403121561189a578081fd5b81518015158114611882578182fd5b6000602082840312156118ba578081fd5b5051919050565b6000602082840312156118d2578081fd5b81516001600160e01b031981168114611882578182fd5b6000602082840312156118fa578081fd5b813561188281611cfd565b600060208284031215611916578081fd5b815161188281611cfd565b600080600060408486031215611935578182fd5b833561194081611cfd565b925060208401356001600160401b0381111561195a578283fd5b61196686828701611812565b9497909650939450505050565b60008060008060006080868803121561198a578081fd5b853561199581611cfd565b94506020860135935060408601356001600160401b038111156119b6578182fd5b6119c288828901611812565b90945092505060608601356119d681611cfd565b809150509295509295909350565b6000806000606084860312156119f8578283fd5b8351611a0381611cfd565b6020850151909350611a1481611cfd565b6040850151909250611a2581611cfd565b809150509250925092565b8183823760009101908152919050565b82848237909101908152602001919050565b60008251611a64818460208701611c5e565b9190910192915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60208082526026908201527f4578656375746f72436f6e74726163743a20756e65787065637465642068616c60408201526506620737465760d41b606082015260800190565b60208082526024908201527f4578656375746f72436f6e74726163743a20636f6e66696720697320696e61636040820152637469766560e01b606082015260800190565b60208082526028908201527f4578656375746f72436f6e74726163743a20756e6578706563746564206261746040820152670c6d040d2dcc8caf60c31b606082015260800190565b6001600160401b03841681528260208201526060604082015260008251806060840152611ba3816080850160208701611c5e565b601f01601f191691909101608001949350505050565b6001600160401b03831681526040810160028310611be757634e487b7160e01b600052602160045260246000fd5b8260208301529392505050565b6000808335601e19843603018112611c0a578283fd5b8301803591506001600160401b03821115611c23578283fd5b60200191503681900382131561185457600080fd5b60006001600160401b0380841680611c5257611c52611ce7565b92169190910492915050565b60005b83811015611c79578181015183820152602001611c61565b83811115611c88576000848401525b50505050565b60006001600160401b0380831681811415611cb757634e487b7160e01b83526011600452602483fd5b6001019392505050565b60006001600160401b0380841680611cdb57611cdb611ce7565b92169190910692915050565b634e487b7160e01b600052601260045260246000fd5b6001600160401b0381168114611d1257600080fd5b5056fea2646970667358221220b27df4d1c313e50e1a22d64ddcce2cfa83980d980c9b123a03b6bb719c7de0a164736f6c63430008040033"

// DeployExecutorContract deploys a new Ethereum contract, binding an instance of ExecutorContract to it.
func DeployExecutorContract(auth *bind.TransactOpts, backend bind.ContractBackend, configContractAddress common.Address, batcherContractAddress common.Address) (common.Address, *types.Transaction, *ExecutorContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExecutorContractBin), backend, configContractAddress, batcherContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutorContract{ExecutorContractCaller: ExecutorContractCaller{contract: contract}, ExecutorContractTransactor: ExecutorContractTransactor{contract: contract}, ExecutorContractFilterer: ExecutorContractFilterer{contract: contract}}, nil
}

// ExecutorContract is an auto generated Go binding around an Ethereum contract.
type ExecutorContract struct {
	ExecutorContractCaller     // Read-only binding to the contract
	ExecutorContractTransactor // Write-only binding to the contract
	ExecutorContractFilterer   // Log filterer for contract events
}

// ExecutorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorContractSession struct {
	Contract     *ExecutorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorContractCallerSession struct {
	Contract *ExecutorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExecutorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorContractTransactorSession struct {
	Contract     *ExecutorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExecutorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorContractRaw struct {
	Contract *ExecutorContract // Generic contract binding to access the raw methods on
}

// ExecutorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorContractCallerRaw struct {
	Contract *ExecutorContractCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorContractTransactorRaw struct {
	Contract *ExecutorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutorContract creates a new instance of ExecutorContract, bound to a specific deployed contract.
func NewExecutorContract(address common.Address, backend bind.ContractBackend) (*ExecutorContract, error) {
	contract, err := bindExecutorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutorContract{ExecutorContractCaller: ExecutorContractCaller{contract: contract}, ExecutorContractTransactor: ExecutorContractTransactor{contract: contract}, ExecutorContractFilterer: ExecutorContractFilterer{contract: contract}}, nil
}

// NewExecutorContractCaller creates a new read-only instance of ExecutorContract, bound to a specific deployed contract.
func NewExecutorContractCaller(address common.Address, caller bind.ContractCaller) (*ExecutorContractCaller, error) {
	contract, err := bindExecutorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorContractCaller{contract: contract}, nil
}

// NewExecutorContractTransactor creates a new write-only instance of ExecutorContract, bound to a specific deployed contract.
func NewExecutorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorContractTransactor, error) {
	contract, err := bindExecutorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorContractTransactor{contract: contract}, nil
}

// NewExecutorContractFilterer creates a new log filterer instance of ExecutorContract, bound to a specific deployed contract.
func NewExecutorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorContractFilterer, error) {
	contract, err := bindExecutorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorContractFilterer{contract: contract}, nil
}

// bindExecutorContract binds a generic wrapper to an already deployed contract.
func bindExecutorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorContract *ExecutorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorContract.Contract.ExecutorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorContract *ExecutorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorContract *ExecutorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorContract *ExecutorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorContract *ExecutorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorContract *ExecutorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorContract.Contract.contract.Transact(opts, method, params...)
}

// BatcherContract is a free data retrieval call binding the contract method 0xbeb3b50e.
//
// Solidity: function batcherContract() view returns(address)
func (_ExecutorContract *ExecutorContractCaller) BatcherContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "batcherContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatcherContract is a free data retrieval call binding the contract method 0xbeb3b50e.
//
// Solidity: function batcherContract() view returns(address)
func (_ExecutorContract *ExecutorContractSession) BatcherContract() (common.Address, error) {
	return _ExecutorContract.Contract.BatcherContract(&_ExecutorContract.CallOpts)
}

// BatcherContract is a free data retrieval call binding the contract method 0xbeb3b50e.
//
// Solidity: function batcherContract() view returns(address)
func (_ExecutorContract *ExecutorContractCallerSession) BatcherContract() (common.Address, error) {
	return _ExecutorContract.Contract.BatcherContract(&_ExecutorContract.CallOpts)
}

// CipherExecutionReceipts is a free data retrieval call binding the contract method 0xc87190b3.
//
// Solidity: function cipherExecutionReceipts(uint64 ) view returns(bool executed, address executor, uint64 halfStep, bytes32 cipherBatchHash, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractCaller) CipherExecutionReceipts(opts *bind.CallOpts, arg0 uint64) (struct {
	Executed        bool
	Executor        common.Address
	HalfStep        uint64
	CipherBatchHash [32]byte
	BatchHash       [32]byte
}, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "cipherExecutionReceipts", arg0)

	outstruct := new(struct {
		Executed        bool
		Executor        common.Address
		HalfStep        uint64
		CipherBatchHash [32]byte
		BatchHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Executed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Executor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HalfStep = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.CipherBatchHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.BatchHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CipherExecutionReceipts is a free data retrieval call binding the contract method 0xc87190b3.
//
// Solidity: function cipherExecutionReceipts(uint64 ) view returns(bool executed, address executor, uint64 halfStep, bytes32 cipherBatchHash, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractSession) CipherExecutionReceipts(arg0 uint64) (struct {
	Executed        bool
	Executor        common.Address
	HalfStep        uint64
	CipherBatchHash [32]byte
	BatchHash       [32]byte
}, error) {
	return _ExecutorContract.Contract.CipherExecutionReceipts(&_ExecutorContract.CallOpts, arg0)
}

// CipherExecutionReceipts is a free data retrieval call binding the contract method 0xc87190b3.
//
// Solidity: function cipherExecutionReceipts(uint64 ) view returns(bool executed, address executor, uint64 halfStep, bytes32 cipherBatchHash, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractCallerSession) CipherExecutionReceipts(arg0 uint64) (struct {
	Executed        bool
	Executor        common.Address
	HalfStep        uint64
	CipherBatchHash [32]byte
	BatchHash       [32]byte
}, error) {
	return _ExecutorContract.Contract.CipherExecutionReceipts(&_ExecutorContract.CallOpts, arg0)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_ExecutorContract *ExecutorContractCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_ExecutorContract *ExecutorContractSession) ConfigContract() (common.Address, error) {
	return _ExecutorContract.Contract.ConfigContract(&_ExecutorContract.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_ExecutorContract *ExecutorContractCallerSession) ConfigContract() (common.Address, error) {
	return _ExecutorContract.Contract.ConfigContract(&_ExecutorContract.CallOpts)
}

// GetReceipt is a free data retrieval call binding the contract method 0x25b36cbf.
//
// Solidity: function getReceipt(uint64 halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractCaller) GetReceipt(opts *bind.CallOpts, halfStep uint64) (CipherExecutionReceipt, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "getReceipt", halfStep)

	if err != nil {
		return *new(CipherExecutionReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(CipherExecutionReceipt)).(*CipherExecutionReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0x25b36cbf.
//
// Solidity: function getReceipt(uint64 halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractSession) GetReceipt(halfStep uint64) (CipherExecutionReceipt, error) {
	return _ExecutorContract.Contract.GetReceipt(&_ExecutorContract.CallOpts, halfStep)
}

// GetReceipt is a free data retrieval call binding the contract method 0x25b36cbf.
//
// Solidity: function getReceipt(uint64 halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractCallerSession) GetReceipt(halfStep uint64) (CipherExecutionReceipt, error) {
	return _ExecutorContract.Contract.GetReceipt(&_ExecutorContract.CallOpts, halfStep)
}

// NumExecutionHalfSteps is a free data retrieval call binding the contract method 0xfa6385f4.
//
// Solidity: function numExecutionHalfSteps() view returns(uint64)
func (_ExecutorContract *ExecutorContractCaller) NumExecutionHalfSteps(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "numExecutionHalfSteps")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NumExecutionHalfSteps is a free data retrieval call binding the contract method 0xfa6385f4.
//
// Solidity: function numExecutionHalfSteps() view returns(uint64)
func (_ExecutorContract *ExecutorContractSession) NumExecutionHalfSteps() (uint64, error) {
	return _ExecutorContract.Contract.NumExecutionHalfSteps(&_ExecutorContract.CallOpts)
}

// NumExecutionHalfSteps is a free data retrieval call binding the contract method 0xfa6385f4.
//
// Solidity: function numExecutionHalfSteps() view returns(uint64)
func (_ExecutorContract *ExecutorContractCallerSession) NumExecutionHalfSteps() (uint64, error) {
	return _ExecutorContract.Contract.NumExecutionHalfSteps(&_ExecutorContract.CallOpts)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x87eb2378.
//
// Solidity: function executeCipherBatch(uint64 batchIndex, bytes32 cipherBatchHash, bytes[] transactions, uint64 keyperIndex) returns()
func (_ExecutorContract *ExecutorContractTransactor) ExecuteCipherBatch(opts *bind.TransactOpts, batchIndex uint64, cipherBatchHash [32]byte, transactions [][]byte, keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "executeCipherBatch", batchIndex, cipherBatchHash, transactions, keyperIndex)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x87eb2378.
//
// Solidity: function executeCipherBatch(uint64 batchIndex, bytes32 cipherBatchHash, bytes[] transactions, uint64 keyperIndex) returns()
func (_ExecutorContract *ExecutorContractSession) ExecuteCipherBatch(batchIndex uint64, cipherBatchHash [32]byte, transactions [][]byte, keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, batchIndex, cipherBatchHash, transactions, keyperIndex)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x87eb2378.
//
// Solidity: function executeCipherBatch(uint64 batchIndex, bytes32 cipherBatchHash, bytes[] transactions, uint64 keyperIndex) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) ExecuteCipherBatch(batchIndex uint64, cipherBatchHash [32]byte, transactions [][]byte, keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, batchIndex, cipherBatchHash, transactions, keyperIndex)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0x4b2a026d.
//
// Solidity: function executePlainBatch(uint64 batchIndex, bytes[] transactions) returns()
func (_ExecutorContract *ExecutorContractTransactor) ExecutePlainBatch(opts *bind.TransactOpts, batchIndex uint64, transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "executePlainBatch", batchIndex, transactions)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0x4b2a026d.
//
// Solidity: function executePlainBatch(uint64 batchIndex, bytes[] transactions) returns()
func (_ExecutorContract *ExecutorContractSession) ExecutePlainBatch(batchIndex uint64, transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutePlainBatch(&_ExecutorContract.TransactOpts, batchIndex, transactions)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0x4b2a026d.
//
// Solidity: function executePlainBatch(uint64 batchIndex, bytes[] transactions) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) ExecutePlainBatch(batchIndex uint64, transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutePlainBatch(&_ExecutorContract.TransactOpts, batchIndex, transactions)
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0xd8e9a001.
//
// Solidity: function skipCipherExecution(uint64 batchIndex) returns()
func (_ExecutorContract *ExecutorContractTransactor) SkipCipherExecution(opts *bind.TransactOpts, batchIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "skipCipherExecution", batchIndex)
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0xd8e9a001.
//
// Solidity: function skipCipherExecution(uint64 batchIndex) returns()
func (_ExecutorContract *ExecutorContractSession) SkipCipherExecution(batchIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.SkipCipherExecution(&_ExecutorContract.TransactOpts, batchIndex)
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0xd8e9a001.
//
// Solidity: function skipCipherExecution(uint64 batchIndex) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) SkipCipherExecution(batchIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.SkipCipherExecution(&_ExecutorContract.TransactOpts, batchIndex)
}

// ExecutorContractBatchExecutedIterator is returned from FilterBatchExecuted and is used to iterate over the raw logs and unpacked data for BatchExecuted events raised by the ExecutorContract contract.
type ExecutorContractBatchExecutedIterator struct {
	Event *ExecutorContractBatchExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExecutorContractBatchExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorContractBatchExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExecutorContractBatchExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExecutorContractBatchExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorContractBatchExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorContractBatchExecuted represents a BatchExecuted event raised by the ExecutorContract contract.
type ExecutorContractBatchExecuted struct {
	NumExecutionHalfSteps uint64
	BatchHash             [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBatchExecuted is a free log retrieval operation binding the contract event 0x3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a.
//
// Solidity: event BatchExecuted(uint64 numExecutionHalfSteps, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractFilterer) FilterBatchExecuted(opts *bind.FilterOpts) (*ExecutorContractBatchExecutedIterator, error) {

	logs, sub, err := _ExecutorContract.contract.FilterLogs(opts, "BatchExecuted")
	if err != nil {
		return nil, err
	}
	return &ExecutorContractBatchExecutedIterator{contract: _ExecutorContract.contract, event: "BatchExecuted", logs: logs, sub: sub}, nil
}

// WatchBatchExecuted is a free log subscription operation binding the contract event 0x3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a.
//
// Solidity: event BatchExecuted(uint64 numExecutionHalfSteps, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractFilterer) WatchBatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutorContractBatchExecuted) (event.Subscription, error) {

	logs, sub, err := _ExecutorContract.contract.WatchLogs(opts, "BatchExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorContractBatchExecuted)
				if err := _ExecutorContract.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchExecuted is a log parse operation binding the contract event 0x3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a.
//
// Solidity: event BatchExecuted(uint64 numExecutionHalfSteps, bytes32 batchHash)
func (_ExecutorContract *ExecutorContractFilterer) ParseBatchExecuted(log types.Log) (*ExecutorContractBatchExecuted, error) {
	event := new(ExecutorContractBatchExecuted)
	if err := _ExecutorContract.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorContractCipherExecutionSkippedIterator is returned from FilterCipherExecutionSkipped and is used to iterate over the raw logs and unpacked data for CipherExecutionSkipped events raised by the ExecutorContract contract.
type ExecutorContractCipherExecutionSkippedIterator struct {
	Event *ExecutorContractCipherExecutionSkipped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExecutorContractCipherExecutionSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorContractCipherExecutionSkipped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExecutorContractCipherExecutionSkipped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExecutorContractCipherExecutionSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorContractCipherExecutionSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorContractCipherExecutionSkipped represents a CipherExecutionSkipped event raised by the ExecutorContract contract.
type ExecutorContractCipherExecutionSkipped struct {
	NumExecutionHalfSteps uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCipherExecutionSkipped is a free log retrieval operation binding the contract event 0xa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be.
//
// Solidity: event CipherExecutionSkipped(uint64 numExecutionHalfSteps)
func (_ExecutorContract *ExecutorContractFilterer) FilterCipherExecutionSkipped(opts *bind.FilterOpts) (*ExecutorContractCipherExecutionSkippedIterator, error) {

	logs, sub, err := _ExecutorContract.contract.FilterLogs(opts, "CipherExecutionSkipped")
	if err != nil {
		return nil, err
	}
	return &ExecutorContractCipherExecutionSkippedIterator{contract: _ExecutorContract.contract, event: "CipherExecutionSkipped", logs: logs, sub: sub}, nil
}

// WatchCipherExecutionSkipped is a free log subscription operation binding the contract event 0xa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be.
//
// Solidity: event CipherExecutionSkipped(uint64 numExecutionHalfSteps)
func (_ExecutorContract *ExecutorContractFilterer) WatchCipherExecutionSkipped(opts *bind.WatchOpts, sink chan<- *ExecutorContractCipherExecutionSkipped) (event.Subscription, error) {

	logs, sub, err := _ExecutorContract.contract.WatchLogs(opts, "CipherExecutionSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorContractCipherExecutionSkipped)
				if err := _ExecutorContract.contract.UnpackLog(event, "CipherExecutionSkipped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCipherExecutionSkipped is a log parse operation binding the contract event 0xa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be.
//
// Solidity: event CipherExecutionSkipped(uint64 numExecutionHalfSteps)
func (_ExecutorContract *ExecutorContractFilterer) ParseCipherExecutionSkipped(log types.Log) (*ExecutorContractCipherExecutionSkipped, error) {
	event := new(ExecutorContractCipherExecutionSkipped)
	if err := _ExecutorContract.contract.UnpackLog(event, "CipherExecutionSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorContractTransactionFailedIterator is returned from FilterTransactionFailed and is used to iterate over the raw logs and unpacked data for TransactionFailed events raised by the ExecutorContract contract.
type ExecutorContractTransactionFailedIterator struct {
	Event *ExecutorContractTransactionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExecutorContractTransactionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorContractTransactionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExecutorContractTransactionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExecutorContractTransactionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorContractTransactionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorContractTransactionFailed represents a TransactionFailed event raised by the ExecutorContract contract.
type ExecutorContractTransactionFailed struct {
	TxIndex uint64
	TxHash  [32]byte
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionFailed is a free log retrieval operation binding the contract event 0x6f580fd78fd3fc2e6db770c1da4b991b45906920b1fe2abd8cc7d760b8021049.
//
// Solidity: event TransactionFailed(uint64 txIndex, bytes32 txHash, bytes data)
func (_ExecutorContract *ExecutorContractFilterer) FilterTransactionFailed(opts *bind.FilterOpts) (*ExecutorContractTransactionFailedIterator, error) {

	logs, sub, err := _ExecutorContract.contract.FilterLogs(opts, "TransactionFailed")
	if err != nil {
		return nil, err
	}
	return &ExecutorContractTransactionFailedIterator{contract: _ExecutorContract.contract, event: "TransactionFailed", logs: logs, sub: sub}, nil
}

// WatchTransactionFailed is a free log subscription operation binding the contract event 0x6f580fd78fd3fc2e6db770c1da4b991b45906920b1fe2abd8cc7d760b8021049.
//
// Solidity: event TransactionFailed(uint64 txIndex, bytes32 txHash, bytes data)
func (_ExecutorContract *ExecutorContractFilterer) WatchTransactionFailed(opts *bind.WatchOpts, sink chan<- *ExecutorContractTransactionFailed) (event.Subscription, error) {

	logs, sub, err := _ExecutorContract.contract.WatchLogs(opts, "TransactionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorContractTransactionFailed)
				if err := _ExecutorContract.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionFailed is a log parse operation binding the contract event 0x6f580fd78fd3fc2e6db770c1da4b991b45906920b1fe2abd8cc7d760b8021049.
//
// Solidity: event TransactionFailed(uint64 txIndex, bytes32 txHash, bytes data)
func (_ExecutorContract *ExecutorContractFilterer) ParseTransactionFailed(log types.Log) (*ExecutorContractTransactionFailed, error) {
	event := new(ExecutorContractTransactionFailed)
	if err := _ExecutorContract.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeBankContractABI is the input ABI used to generate the binding from.
const FeeBankContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeBankContractBin is the compiled bytecode used for deploying new contracts.
var FeeBankContractBin = "0x608060405234801561001057600080fd5b50610664806100206000396000f3fe60806040526004361061003f5760003560e01c80633ccfd60b14610044578063d6dad0601461005b578063f340fa011461007b578063fc7e286d1461008e575b600080fd5b34801561005057600080fd5b506100596100e0565b005b34801561006757600080fd5b50610059610076366004610543565b610105565b610059610089366004610522565b610113565b34801561009a57600080fd5b506100c46100a9366004610522565b6000602081905290815260409020546001600160401b031681565b6040516001600160401b03909116815260200160405180910390f35b3360008181526020819052604090205461010391906001600160401b03166102d5565b565b61010f82826102d5565b5050565b6001600160a01b0381166101425760405162461bcd60e51b815260040161013990610584565b60405180910390fd5b600034116101895760405162461bcd60e51b815260206004820152601460248201527346656542616e6b3a20666565206973207a65726f60601b6044820152606401610139565b6001600160a01b0381166000908152602081905260409020546101b7906001600160401b03908116906105f0565b6001600160401b031634111561021b5760405162461bcd60e51b8152602060048201526024808201527f46656542616e6b3a2062616c616e636520776f756c64206578636565642075696044820152631b9d0d8d60e21b6064820152608401610139565b6001600160a01b0381166000908152602081905260408120805434929061024c9084906001600160401b03166105c5565b82546101009290920a6001600160401b038181021990931691831602179091556001600160a01b038316600081815260208181526040918290205482513381529182019390935234841691810191909152911660608201527fc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d9915060800160405180910390a150565b6001600160a01b0382166102fb5760405162461bcd60e51b815260040161013990610584565b336000908152602081905260409020546001600160401b0316806103615760405162461bcd60e51b815260206004820152601960248201527f46656542616e6b3a206465706f73697420697320656d707479000000000000006044820152606401610139565b806001600160401b0316826001600160401b031611156103c35760405162461bcd60e51b815260206004820152601f60248201527f46656542616e6b3a20616d6f756e742065786365656473206465706f736974006044820152606401610139565b6103cd82826105f0565b33600090815260208190526040808220805467ffffffffffffffff19166001600160401b039485161790555190916001600160a01b03861691908516908381818185875af1925050503d8060008114610442576040519150601f19603f3d011682016040523d82523d6000602084013e610447565b606091505b50509050806104985760405162461bcd60e51b815260206004820152601f60248201527f46656542616e6b3a207769746864726177616c2063616c6c206661696c6564006044820152606401610139565b33600081815260208181526040918290205482519384526001600160a01b038816918401919091526001600160401b0386811684840152166060830152517f4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e9181900360800190a150505050565b80356001600160a01b038116811461051d57600080fd5b919050565b600060208284031215610533578081fd5b61053c82610506565b9392505050565b60008060408385031215610555578081fd5b61055e83610506565b915060208301356001600160401b0381168114610579578182fd5b809150509250929050565b60208082526021908201527f46656542616e6b3a207265636569766572206973207a65726f206164647265736040820152607360f81b606082015260800190565b60006001600160401b038083168185168083038211156105e7576105e7610618565b01949350505050565b60006001600160401b038381169083168181101561061057610610610618565b039392505050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220be915a649aba443074f55f42ec6b5f77340fefd8dd4716cb7ec0d0924650ae0964736f6c63430008040033"

// DeployFeeBankContract deploys a new Ethereum contract, binding an instance of FeeBankContract to it.
func DeployFeeBankContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeBankContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeBankContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeBankContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeBankContract{FeeBankContractCaller: FeeBankContractCaller{contract: contract}, FeeBankContractTransactor: FeeBankContractTransactor{contract: contract}, FeeBankContractFilterer: FeeBankContractFilterer{contract: contract}}, nil
}

// FeeBankContract is an auto generated Go binding around an Ethereum contract.
type FeeBankContract struct {
	FeeBankContractCaller     // Read-only binding to the contract
	FeeBankContractTransactor // Write-only binding to the contract
	FeeBankContractFilterer   // Log filterer for contract events
}

// FeeBankContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeBankContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBankContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeBankContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBankContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeBankContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBankContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeBankContractSession struct {
	Contract     *FeeBankContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeBankContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeBankContractCallerSession struct {
	Contract *FeeBankContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FeeBankContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeBankContractTransactorSession struct {
	Contract     *FeeBankContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FeeBankContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeBankContractRaw struct {
	Contract *FeeBankContract // Generic contract binding to access the raw methods on
}

// FeeBankContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeBankContractCallerRaw struct {
	Contract *FeeBankContractCaller // Generic read-only contract binding to access the raw methods on
}

// FeeBankContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeBankContractTransactorRaw struct {
	Contract *FeeBankContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeBankContract creates a new instance of FeeBankContract, bound to a specific deployed contract.
func NewFeeBankContract(address common.Address, backend bind.ContractBackend) (*FeeBankContract, error) {
	contract, err := bindFeeBankContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeBankContract{FeeBankContractCaller: FeeBankContractCaller{contract: contract}, FeeBankContractTransactor: FeeBankContractTransactor{contract: contract}, FeeBankContractFilterer: FeeBankContractFilterer{contract: contract}}, nil
}

// NewFeeBankContractCaller creates a new read-only instance of FeeBankContract, bound to a specific deployed contract.
func NewFeeBankContractCaller(address common.Address, caller bind.ContractCaller) (*FeeBankContractCaller, error) {
	contract, err := bindFeeBankContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeBankContractCaller{contract: contract}, nil
}

// NewFeeBankContractTransactor creates a new write-only instance of FeeBankContract, bound to a specific deployed contract.
func NewFeeBankContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeBankContractTransactor, error) {
	contract, err := bindFeeBankContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeBankContractTransactor{contract: contract}, nil
}

// NewFeeBankContractFilterer creates a new log filterer instance of FeeBankContract, bound to a specific deployed contract.
func NewFeeBankContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeBankContractFilterer, error) {
	contract, err := bindFeeBankContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeBankContractFilterer{contract: contract}, nil
}

// bindFeeBankContract binds a generic wrapper to an already deployed contract.
func bindFeeBankContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeBankContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeBankContract *FeeBankContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeBankContract.Contract.FeeBankContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeBankContract *FeeBankContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBankContract.Contract.FeeBankContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeBankContract *FeeBankContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeBankContract.Contract.FeeBankContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeBankContract *FeeBankContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeBankContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeBankContract *FeeBankContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBankContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeBankContract *FeeBankContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeBankContract.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint64)
func (_FeeBankContract *FeeBankContractCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _FeeBankContract.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint64)
func (_FeeBankContract *FeeBankContractSession) Deposits(arg0 common.Address) (uint64, error) {
	return _FeeBankContract.Contract.Deposits(&_FeeBankContract.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint64)
func (_FeeBankContract *FeeBankContractCallerSession) Deposits(arg0 common.Address) (uint64, error) {
	return _FeeBankContract.Contract.Deposits(&_FeeBankContract.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_FeeBankContract *FeeBankContractTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.contract.Transact(opts, "deposit", receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_FeeBankContract *FeeBankContractSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Deposit(&_FeeBankContract.TransactOpts, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_FeeBankContract *FeeBankContractTransactorSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Deposit(&_FeeBankContract.TransactOpts, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FeeBankContract *FeeBankContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBankContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FeeBankContract *FeeBankContractSession) Withdraw() (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw(&_FeeBankContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FeeBankContract *FeeBankContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw(&_FeeBankContract.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd6dad060.
//
// Solidity: function withdraw(address receiver, uint64 amount) returns()
func (_FeeBankContract *FeeBankContractTransactor) Withdraw0(opts *bind.TransactOpts, receiver common.Address, amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.contract.Transact(opts, "withdraw0", receiver, amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd6dad060.
//
// Solidity: function withdraw(address receiver, uint64 amount) returns()
func (_FeeBankContract *FeeBankContractSession) Withdraw0(receiver common.Address, amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw0(&_FeeBankContract.TransactOpts, receiver, amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd6dad060.
//
// Solidity: function withdraw(address receiver, uint64 amount) returns()
func (_FeeBankContract *FeeBankContractTransactorSession) Withdraw0(receiver common.Address, amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw0(&_FeeBankContract.TransactOpts, receiver, amount)
}

// FeeBankContractDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the FeeBankContract contract.
type FeeBankContractDepositEventIterator struct {
	Event *FeeBankContractDepositEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeBankContractDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBankContractDepositEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeBankContractDepositEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeBankContractDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBankContractDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBankContractDepositEvent represents a DepositEvent event raised by the FeeBankContract contract.
type FeeBankContractDepositEvent struct {
	Depositor   common.Address
	Receiver    common.Address
	Amount      uint64
	TotalAmount uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0xc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d9.
//
// Solidity: event DepositEvent(address depositor, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*FeeBankContractDepositEventIterator, error) {

	logs, sub, err := _FeeBankContract.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &FeeBankContractDepositEventIterator{contract: _FeeBankContract.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0xc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d9.
//
// Solidity: event DepositEvent(address depositor, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *FeeBankContractDepositEvent) (event.Subscription, error) {

	logs, sub, err := _FeeBankContract.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBankContractDepositEvent)
				if err := _FeeBankContract.contract.UnpackLog(event, "DepositEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositEvent is a log parse operation binding the contract event 0xc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d9.
//
// Solidity: event DepositEvent(address depositor, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) ParseDepositEvent(log types.Log) (*FeeBankContractDepositEvent, error) {
	event := new(FeeBankContractDepositEvent)
	if err := _FeeBankContract.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeBankContractWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the FeeBankContract contract.
type FeeBankContractWithdrawEventIterator struct {
	Event *FeeBankContractWithdrawEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeBankContractWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBankContractWithdrawEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeBankContractWithdrawEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeBankContractWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBankContractWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBankContractWithdrawEvent represents a WithdrawEvent event raised by the FeeBankContract contract.
type FeeBankContractWithdrawEvent struct {
	Sender      common.Address
	Receiver    common.Address
	Amount      uint64
	TotalAmount uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e.
//
// Solidity: event WithdrawEvent(address sender, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) FilterWithdrawEvent(opts *bind.FilterOpts) (*FeeBankContractWithdrawEventIterator, error) {

	logs, sub, err := _FeeBankContract.contract.FilterLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &FeeBankContractWithdrawEventIterator{contract: _FeeBankContract.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e.
//
// Solidity: event WithdrawEvent(address sender, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *FeeBankContractWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _FeeBankContract.contract.WatchLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBankContractWithdrawEvent)
				if err := _FeeBankContract.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawEvent is a log parse operation binding the contract event 0x4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e.
//
// Solidity: event WithdrawEvent(address sender, address receiver, uint64 amount, uint64 totalAmount)
func (_FeeBankContract *FeeBankContractFilterer) ParseWithdrawEvent(log types.Log) (*FeeBankContractWithdrawEvent, error) {
	event := new(FeeBankContractWithdrawEvent)
	if err := _FeeBankContract.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1820RegistryABI is the input ABI used to generate the binding from.
const IERC1820RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceImplementerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_interfaceHash\",\"type\":\"bytes32\"}],\"name\":\"getInterfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165Interface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165InterfaceNoCache\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"interfaceName\",\"type\":\"string\"}],\"name\":\"interfaceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_interfaceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"updateERC165Cache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC1820Registry is an auto generated Go binding around an Ethereum contract.
type IERC1820Registry struct {
	IERC1820RegistryCaller     // Read-only binding to the contract
	IERC1820RegistryTransactor // Write-only binding to the contract
	IERC1820RegistryFilterer   // Log filterer for contract events
}

// IERC1820RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1820RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1820RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1820RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1820RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1820RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1820RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1820RegistrySession struct {
	Contract     *IERC1820Registry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1820RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1820RegistryCallerSession struct {
	Contract *IERC1820RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IERC1820RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1820RegistryTransactorSession struct {
	Contract     *IERC1820RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IERC1820RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1820RegistryRaw struct {
	Contract *IERC1820Registry // Generic contract binding to access the raw methods on
}

// IERC1820RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1820RegistryCallerRaw struct {
	Contract *IERC1820RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1820RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1820RegistryTransactorRaw struct {
	Contract *IERC1820RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1820Registry creates a new instance of IERC1820Registry, bound to a specific deployed contract.
func NewIERC1820Registry(address common.Address, backend bind.ContractBackend) (*IERC1820Registry, error) {
	contract, err := bindIERC1820Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1820Registry{IERC1820RegistryCaller: IERC1820RegistryCaller{contract: contract}, IERC1820RegistryTransactor: IERC1820RegistryTransactor{contract: contract}, IERC1820RegistryFilterer: IERC1820RegistryFilterer{contract: contract}}, nil
}

// NewIERC1820RegistryCaller creates a new read-only instance of IERC1820Registry, bound to a specific deployed contract.
func NewIERC1820RegistryCaller(address common.Address, caller bind.ContractCaller) (*IERC1820RegistryCaller, error) {
	contract, err := bindIERC1820Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1820RegistryCaller{contract: contract}, nil
}

// NewIERC1820RegistryTransactor creates a new write-only instance of IERC1820Registry, bound to a specific deployed contract.
func NewIERC1820RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1820RegistryTransactor, error) {
	contract, err := bindIERC1820Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1820RegistryTransactor{contract: contract}, nil
}

// NewIERC1820RegistryFilterer creates a new log filterer instance of IERC1820Registry, bound to a specific deployed contract.
func NewIERC1820RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1820RegistryFilterer, error) {
	contract, err := bindIERC1820Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1820RegistryFilterer{contract: contract}, nil
}

// bindIERC1820Registry binds a generic wrapper to an already deployed contract.
func bindIERC1820Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1820RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1820Registry *IERC1820RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1820Registry.Contract.IERC1820RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1820Registry *IERC1820RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.IERC1820RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1820Registry *IERC1820RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.IERC1820RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1820Registry *IERC1820RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1820Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1820Registry *IERC1820RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1820Registry *IERC1820RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.contract.Transact(opts, method, params...)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 _interfaceHash) view returns(address)
func (_IERC1820Registry *IERC1820RegistryCaller) GetInterfaceImplementer(opts *bind.CallOpts, account common.Address, _interfaceHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IERC1820Registry.contract.Call(opts, &out, "getInterfaceImplementer", account, _interfaceHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 _interfaceHash) view returns(address)
func (_IERC1820Registry *IERC1820RegistrySession) GetInterfaceImplementer(account common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _IERC1820Registry.Contract.GetInterfaceImplementer(&_IERC1820Registry.CallOpts, account, _interfaceHash)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 _interfaceHash) view returns(address)
func (_IERC1820Registry *IERC1820RegistryCallerSession) GetInterfaceImplementer(account common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _IERC1820Registry.Contract.GetInterfaceImplementer(&_IERC1820Registry.CallOpts, account, _interfaceHash)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) view returns(address)
func (_IERC1820Registry *IERC1820RegistryCaller) GetManager(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _IERC1820Registry.contract.Call(opts, &out, "getManager", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) view returns(address)
func (_IERC1820Registry *IERC1820RegistrySession) GetManager(account common.Address) (common.Address, error) {
	return _IERC1820Registry.Contract.GetManager(&_IERC1820Registry.CallOpts, account)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) view returns(address)
func (_IERC1820Registry *IERC1820RegistryCallerSession) GetManager(account common.Address) (common.Address, error) {
	return _IERC1820Registry.Contract.GetManager(&_IERC1820Registry.CallOpts, account)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistryCaller) ImplementsERC165Interface(opts *bind.CallOpts, account common.Address, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1820Registry.contract.Call(opts, &out, "implementsERC165Interface", account, interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistrySession) ImplementsERC165Interface(account common.Address, interfaceId [4]byte) (bool, error) {
	return _IERC1820Registry.Contract.ImplementsERC165Interface(&_IERC1820Registry.CallOpts, account, interfaceId)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistryCallerSession) ImplementsERC165Interface(account common.Address, interfaceId [4]byte) (bool, error) {
	return _IERC1820Registry.Contract.ImplementsERC165Interface(&_IERC1820Registry.CallOpts, account, interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistryCaller) ImplementsERC165InterfaceNoCache(opts *bind.CallOpts, account common.Address, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1820Registry.contract.Call(opts, &out, "implementsERC165InterfaceNoCache", account, interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistrySession) ImplementsERC165InterfaceNoCache(account common.Address, interfaceId [4]byte) (bool, error) {
	return _IERC1820Registry.Contract.ImplementsERC165InterfaceNoCache(&_IERC1820Registry.CallOpts, account, interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address account, bytes4 interfaceId) view returns(bool)
func (_IERC1820Registry *IERC1820RegistryCallerSession) ImplementsERC165InterfaceNoCache(account common.Address, interfaceId [4]byte) (bool, error) {
	return _IERC1820Registry.Contract.ImplementsERC165InterfaceNoCache(&_IERC1820Registry.CallOpts, account, interfaceId)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string interfaceName) pure returns(bytes32)
func (_IERC1820Registry *IERC1820RegistryCaller) InterfaceHash(opts *bind.CallOpts, interfaceName string) ([32]byte, error) {
	var out []interface{}
	err := _IERC1820Registry.contract.Call(opts, &out, "interfaceHash", interfaceName)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string interfaceName) pure returns(bytes32)
func (_IERC1820Registry *IERC1820RegistrySession) InterfaceHash(interfaceName string) ([32]byte, error) {
	return _IERC1820Registry.Contract.InterfaceHash(&_IERC1820Registry.CallOpts, interfaceName)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string interfaceName) pure returns(bytes32)
func (_IERC1820Registry *IERC1820RegistryCallerSession) InterfaceHash(interfaceName string) ([32]byte, error) {
	return _IERC1820Registry.Contract.InterfaceHash(&_IERC1820Registry.CallOpts, interfaceName)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address account, bytes32 _interfaceHash, address implementer) returns()
func (_IERC1820Registry *IERC1820RegistryTransactor) SetInterfaceImplementer(opts *bind.TransactOpts, account common.Address, _interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.contract.Transact(opts, "setInterfaceImplementer", account, _interfaceHash, implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address account, bytes32 _interfaceHash, address implementer) returns()
func (_IERC1820Registry *IERC1820RegistrySession) SetInterfaceImplementer(account common.Address, _interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.SetInterfaceImplementer(&_IERC1820Registry.TransactOpts, account, _interfaceHash, implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address account, bytes32 _interfaceHash, address implementer) returns()
func (_IERC1820Registry *IERC1820RegistryTransactorSession) SetInterfaceImplementer(account common.Address, _interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.SetInterfaceImplementer(&_IERC1820Registry.TransactOpts, account, _interfaceHash, implementer)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address account, address newManager) returns()
func (_IERC1820Registry *IERC1820RegistryTransactor) SetManager(opts *bind.TransactOpts, account common.Address, newManager common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.contract.Transact(opts, "setManager", account, newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address account, address newManager) returns()
func (_IERC1820Registry *IERC1820RegistrySession) SetManager(account common.Address, newManager common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.SetManager(&_IERC1820Registry.TransactOpts, account, newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address account, address newManager) returns()
func (_IERC1820Registry *IERC1820RegistryTransactorSession) SetManager(account common.Address, newManager common.Address) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.SetManager(&_IERC1820Registry.TransactOpts, account, newManager)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address account, bytes4 interfaceId) returns()
func (_IERC1820Registry *IERC1820RegistryTransactor) UpdateERC165Cache(opts *bind.TransactOpts, account common.Address, interfaceId [4]byte) (*types.Transaction, error) {
	return _IERC1820Registry.contract.Transact(opts, "updateERC165Cache", account, interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address account, bytes4 interfaceId) returns()
func (_IERC1820Registry *IERC1820RegistrySession) UpdateERC165Cache(account common.Address, interfaceId [4]byte) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.UpdateERC165Cache(&_IERC1820Registry.TransactOpts, account, interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address account, bytes4 interfaceId) returns()
func (_IERC1820Registry *IERC1820RegistryTransactorSession) UpdateERC165Cache(account common.Address, interfaceId [4]byte) (*types.Transaction, error) {
	return _IERC1820Registry.Contract.UpdateERC165Cache(&_IERC1820Registry.TransactOpts, account, interfaceId)
}

// IERC1820RegistryInterfaceImplementerSetIterator is returned from FilterInterfaceImplementerSet and is used to iterate over the raw logs and unpacked data for InterfaceImplementerSet events raised by the IERC1820Registry contract.
type IERC1820RegistryInterfaceImplementerSetIterator struct {
	Event *IERC1820RegistryInterfaceImplementerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC1820RegistryInterfaceImplementerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1820RegistryInterfaceImplementerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC1820RegistryInterfaceImplementerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC1820RegistryInterfaceImplementerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1820RegistryInterfaceImplementerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1820RegistryInterfaceImplementerSet represents a InterfaceImplementerSet event raised by the IERC1820Registry contract.
type IERC1820RegistryInterfaceImplementerSet struct {
	Account       common.Address
	InterfaceHash [32]byte
	Implementer   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterfaceImplementerSet is a free log retrieval operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_IERC1820Registry *IERC1820RegistryFilterer) FilterInterfaceImplementerSet(opts *bind.FilterOpts, account []common.Address, interfaceHash [][32]byte, implementer []common.Address) (*IERC1820RegistryInterfaceImplementerSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _IERC1820Registry.contract.FilterLogs(opts, "InterfaceImplementerSet", accountRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return &IERC1820RegistryInterfaceImplementerSetIterator{contract: _IERC1820Registry.contract, event: "InterfaceImplementerSet", logs: logs, sub: sub}, nil
}

// WatchInterfaceImplementerSet is a free log subscription operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_IERC1820Registry *IERC1820RegistryFilterer) WatchInterfaceImplementerSet(opts *bind.WatchOpts, sink chan<- *IERC1820RegistryInterfaceImplementerSet, account []common.Address, interfaceHash [][32]byte, implementer []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _IERC1820Registry.contract.WatchLogs(opts, "InterfaceImplementerSet", accountRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1820RegistryInterfaceImplementerSet)
				if err := _IERC1820Registry.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterfaceImplementerSet is a log parse operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_IERC1820Registry *IERC1820RegistryFilterer) ParseInterfaceImplementerSet(log types.Log) (*IERC1820RegistryInterfaceImplementerSet, error) {
	event := new(IERC1820RegistryInterfaceImplementerSet)
	if err := _IERC1820Registry.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1820RegistryManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the IERC1820Registry contract.
type IERC1820RegistryManagerChangedIterator struct {
	Event *IERC1820RegistryManagerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC1820RegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1820RegistryManagerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC1820RegistryManagerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC1820RegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1820RegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1820RegistryManagerChanged represents a ManagerChanged event raised by the IERC1820Registry contract.
type IERC1820RegistryManagerChanged struct {
	Account    common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_IERC1820Registry *IERC1820RegistryFilterer) FilterManagerChanged(opts *bind.FilterOpts, account []common.Address, newManager []common.Address) (*IERC1820RegistryManagerChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _IERC1820Registry.contract.FilterLogs(opts, "ManagerChanged", accountRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &IERC1820RegistryManagerChangedIterator{contract: _IERC1820Registry.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_IERC1820Registry *IERC1820RegistryFilterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *IERC1820RegistryManagerChanged, account []common.Address, newManager []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _IERC1820Registry.contract.WatchLogs(opts, "ManagerChanged", accountRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1820RegistryManagerChanged)
				if err := _IERC1820Registry.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerChanged is a log parse operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_IERC1820Registry *IERC1820RegistryFilterer) ParseManagerChanged(log types.Log) (*IERC1820RegistryManagerChanged, error) {
	event := new(IERC1820RegistryManagerChanged)
	if err := _IERC1820Registry.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777ABI is the input ABI used to generate the binding from.
const IERC777ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC777 is an auto generated Go binding around an Ethereum contract.
type IERC777 struct {
	IERC777Caller     // Read-only binding to the contract
	IERC777Transactor // Write-only binding to the contract
	IERC777Filterer   // Log filterer for contract events
}

// IERC777Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC777Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC777Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC777Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC777Session struct {
	Contract     *IERC777          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC777CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC777CallerSession struct {
	Contract *IERC777Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC777TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC777TransactorSession struct {
	Contract     *IERC777Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC777Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC777Raw struct {
	Contract *IERC777 // Generic contract binding to access the raw methods on
}

// IERC777CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC777CallerRaw struct {
	Contract *IERC777Caller // Generic read-only contract binding to access the raw methods on
}

// IERC777TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC777TransactorRaw struct {
	Contract *IERC777Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC777 creates a new instance of IERC777, bound to a specific deployed contract.
func NewIERC777(address common.Address, backend bind.ContractBackend) (*IERC777, error) {
	contract, err := bindIERC777(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC777{IERC777Caller: IERC777Caller{contract: contract}, IERC777Transactor: IERC777Transactor{contract: contract}, IERC777Filterer: IERC777Filterer{contract: contract}}, nil
}

// NewIERC777Caller creates a new read-only instance of IERC777, bound to a specific deployed contract.
func NewIERC777Caller(address common.Address, caller bind.ContractCaller) (*IERC777Caller, error) {
	contract, err := bindIERC777(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777Caller{contract: contract}, nil
}

// NewIERC777Transactor creates a new write-only instance of IERC777, bound to a specific deployed contract.
func NewIERC777Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC777Transactor, error) {
	contract, err := bindIERC777(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777Transactor{contract: contract}, nil
}

// NewIERC777Filterer creates a new log filterer instance of IERC777, bound to a specific deployed contract.
func NewIERC777Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC777Filterer, error) {
	contract, err := bindIERC777(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC777Filterer{contract: contract}, nil
}

// bindIERC777 binds a generic wrapper to an already deployed contract.
func bindIERC777(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC777ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777 *IERC777Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777.Contract.IERC777Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777 *IERC777Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777.Contract.IERC777Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777 *IERC777Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777.Contract.IERC777Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777 *IERC777CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777 *IERC777TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777 *IERC777TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC777 *IERC777Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC777 *IERC777Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC777.Contract.BalanceOf(&_IERC777.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC777 *IERC777CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC777.Contract.BalanceOf(&_IERC777.CallOpts, owner)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_IERC777 *IERC777Caller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_IERC777 *IERC777Session) DefaultOperators() ([]common.Address, error) {
	return _IERC777.Contract.DefaultOperators(&_IERC777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_IERC777 *IERC777CallerSession) DefaultOperators() ([]common.Address, error) {
	return _IERC777.Contract.DefaultOperators(&_IERC777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_IERC777 *IERC777Caller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_IERC777 *IERC777Session) Granularity() (*big.Int, error) {
	return _IERC777.Contract.Granularity(&_IERC777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_IERC777 *IERC777CallerSession) Granularity() (*big.Int, error) {
	return _IERC777.Contract.Granularity(&_IERC777.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_IERC777 *IERC777Caller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_IERC777 *IERC777Session) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _IERC777.Contract.IsOperatorFor(&_IERC777.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_IERC777 *IERC777CallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _IERC777.Contract.IsOperatorFor(&_IERC777.CallOpts, operator, tokenHolder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC777 *IERC777Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC777 *IERC777Session) Name() (string, error) {
	return _IERC777.Contract.Name(&_IERC777.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC777 *IERC777CallerSession) Name() (string, error) {
	return _IERC777.Contract.Name(&_IERC777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC777 *IERC777Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC777 *IERC777Session) Symbol() (string, error) {
	return _IERC777.Contract.Symbol(&_IERC777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC777 *IERC777CallerSession) Symbol() (string, error) {
	return _IERC777.Contract.Symbol(&_IERC777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC777 *IERC777Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC777.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC777 *IERC777Session) TotalSupply() (*big.Int, error) {
	return _IERC777.Contract.TotalSupply(&_IERC777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC777 *IERC777CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC777.Contract.TotalSupply(&_IERC777.CallOpts)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_IERC777 *IERC777Transactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_IERC777 *IERC777Session) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _IERC777.Contract.AuthorizeOperator(&_IERC777.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_IERC777 *IERC777TransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _IERC777.Contract.AuthorizeOperator(&_IERC777.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_IERC777 *IERC777Transactor) Burn(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "burn", amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_IERC777 *IERC777Session) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.Contract.Burn(&_IERC777.TransactOpts, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_IERC777 *IERC777TransactorSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.Contract.Burn(&_IERC777.TransactOpts, amount, data)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777Transactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777Session) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.Contract.OperatorBurn(&_IERC777.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777TransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.Contract.OperatorBurn(&_IERC777.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777Transactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777Session) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.Contract.OperatorSend(&_IERC777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_IERC777 *IERC777TransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777.Contract.OperatorSend(&_IERC777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_IERC777 *IERC777Transactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_IERC777 *IERC777Session) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _IERC777.Contract.RevokeOperator(&_IERC777.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_IERC777 *IERC777TransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _IERC777.Contract.RevokeOperator(&_IERC777.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_IERC777 *IERC777Transactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_IERC777 *IERC777Session) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.Contract.Send(&_IERC777.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_IERC777 *IERC777TransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC777.Contract.Send(&_IERC777.TransactOpts, recipient, amount, data)
}

// IERC777AuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the IERC777 contract.
type IERC777AuthorizedOperatorIterator struct {
	Event *IERC777AuthorizedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC777AuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC777AuthorizedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC777AuthorizedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC777AuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC777AuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC777AuthorizedOperator represents a AuthorizedOperator event raised by the IERC777 contract.
type IERC777AuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*IERC777AuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _IERC777.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &IERC777AuthorizedOperatorIterator{contract: _IERC777.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *IERC777AuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _IERC777.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC777AuthorizedOperator)
				if err := _IERC777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) ParseAuthorizedOperator(log types.Log) (*IERC777AuthorizedOperator, error) {
	event := new(IERC777AuthorizedOperator)
	if err := _IERC777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the IERC777 contract.
type IERC777BurnedIterator struct {
	Event *IERC777Burned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC777BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC777Burned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC777Burned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC777BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC777BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC777Burned represents a Burned event raised by the IERC777 contract.
type IERC777Burned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*IERC777BurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IERC777.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IERC777BurnedIterator{contract: _IERC777.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *IERC777Burned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IERC777.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC777Burned)
				if err := _IERC777.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) ParseBurned(log types.Log) (*IERC777Burned, error) {
	event := new(IERC777Burned)
	if err := _IERC777.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the IERC777 contract.
type IERC777MintedIterator struct {
	Event *IERC777Minted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC777MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC777Minted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC777Minted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC777MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC777MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC777Minted represents a Minted event raised by the IERC777 contract.
type IERC777Minted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*IERC777MintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC777.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC777MintedIterator{contract: _IERC777.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *IERC777Minted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC777.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC777Minted)
				if err := _IERC777.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) ParseMinted(log types.Log) (*IERC777Minted, error) {
	event := new(IERC777Minted)
	if err := _IERC777.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777RevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the IERC777 contract.
type IERC777RevokedOperatorIterator struct {
	Event *IERC777RevokedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC777RevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC777RevokedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC777RevokedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC777RevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC777RevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC777RevokedOperator represents a RevokedOperator event raised by the IERC777 contract.
type IERC777RevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*IERC777RevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _IERC777.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &IERC777RevokedOperatorIterator{contract: _IERC777.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *IERC777RevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _IERC777.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC777RevokedOperator)
				if err := _IERC777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_IERC777 *IERC777Filterer) ParseRevokedOperator(log types.Log) (*IERC777RevokedOperator, error) {
	event := new(IERC777RevokedOperator)
	if err := _IERC777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777SentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the IERC777 contract.
type IERC777SentIterator struct {
	Event *IERC777Sent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC777SentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC777Sent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC777Sent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC777SentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC777SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC777Sent represents a Sent event raised by the IERC777 contract.
type IERC777Sent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC777SentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC777.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC777SentIterator{contract: _IERC777.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) WatchSent(opts *bind.WatchOpts, sink chan<- *IERC777Sent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC777.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC777Sent)
				if err := _IERC777.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_IERC777 *IERC777Filterer) ParseSent(log types.Log) (*IERC777Sent, error) {
	event := new(IERC777Sent)
	if err := _IERC777.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC777RecipientABI is the input ABI used to generate the binding from.
const IERC777RecipientABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC777Recipient is an auto generated Go binding around an Ethereum contract.
type IERC777Recipient struct {
	IERC777RecipientCaller     // Read-only binding to the contract
	IERC777RecipientTransactor // Write-only binding to the contract
	IERC777RecipientFilterer   // Log filterer for contract events
}

// IERC777RecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC777RecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777RecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC777RecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777RecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC777RecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777RecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC777RecipientSession struct {
	Contract     *IERC777Recipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC777RecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC777RecipientCallerSession struct {
	Contract *IERC777RecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IERC777RecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC777RecipientTransactorSession struct {
	Contract     *IERC777RecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IERC777RecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC777RecipientRaw struct {
	Contract *IERC777Recipient // Generic contract binding to access the raw methods on
}

// IERC777RecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC777RecipientCallerRaw struct {
	Contract *IERC777RecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IERC777RecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC777RecipientTransactorRaw struct {
	Contract *IERC777RecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC777Recipient creates a new instance of IERC777Recipient, bound to a specific deployed contract.
func NewIERC777Recipient(address common.Address, backend bind.ContractBackend) (*IERC777Recipient, error) {
	contract, err := bindIERC777Recipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC777Recipient{IERC777RecipientCaller: IERC777RecipientCaller{contract: contract}, IERC777RecipientTransactor: IERC777RecipientTransactor{contract: contract}, IERC777RecipientFilterer: IERC777RecipientFilterer{contract: contract}}, nil
}

// NewIERC777RecipientCaller creates a new read-only instance of IERC777Recipient, bound to a specific deployed contract.
func NewIERC777RecipientCaller(address common.Address, caller bind.ContractCaller) (*IERC777RecipientCaller, error) {
	contract, err := bindIERC777Recipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777RecipientCaller{contract: contract}, nil
}

// NewIERC777RecipientTransactor creates a new write-only instance of IERC777Recipient, bound to a specific deployed contract.
func NewIERC777RecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC777RecipientTransactor, error) {
	contract, err := bindIERC777Recipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777RecipientTransactor{contract: contract}, nil
}

// NewIERC777RecipientFilterer creates a new log filterer instance of IERC777Recipient, bound to a specific deployed contract.
func NewIERC777RecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC777RecipientFilterer, error) {
	contract, err := bindIERC777Recipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC777RecipientFilterer{contract: contract}, nil
}

// bindIERC777Recipient binds a generic wrapper to an already deployed contract.
func bindIERC777Recipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC777RecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777Recipient *IERC777RecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777Recipient.Contract.IERC777RecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777Recipient *IERC777RecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.IERC777RecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777Recipient *IERC777RecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.IERC777RecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777Recipient *IERC777RecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777Recipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777Recipient *IERC777RecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777Recipient *IERC777RecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.contract.Transact(opts, method, params...)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Recipient *IERC777RecipientTransactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Recipient.contract.Transact(opts, "tokensReceived", operator, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Recipient *IERC777RecipientSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.TokensReceived(&_IERC777Recipient.TransactOpts, operator, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Recipient *IERC777RecipientTransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Recipient.Contract.TokensReceived(&_IERC777Recipient.TransactOpts, operator, from, to, amount, userData, operatorData)
}

// IERC777SenderABI is the input ABI used to generate the binding from.
const IERC777SenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"tokensToSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC777Sender is an auto generated Go binding around an Ethereum contract.
type IERC777Sender struct {
	IERC777SenderCaller     // Read-only binding to the contract
	IERC777SenderTransactor // Write-only binding to the contract
	IERC777SenderFilterer   // Log filterer for contract events
}

// IERC777SenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC777SenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777SenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC777SenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777SenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC777SenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC777SenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC777SenderSession struct {
	Contract     *IERC777Sender    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC777SenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC777SenderCallerSession struct {
	Contract *IERC777SenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC777SenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC777SenderTransactorSession struct {
	Contract     *IERC777SenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC777SenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC777SenderRaw struct {
	Contract *IERC777Sender // Generic contract binding to access the raw methods on
}

// IERC777SenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC777SenderCallerRaw struct {
	Contract *IERC777SenderCaller // Generic read-only contract binding to access the raw methods on
}

// IERC777SenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC777SenderTransactorRaw struct {
	Contract *IERC777SenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC777Sender creates a new instance of IERC777Sender, bound to a specific deployed contract.
func NewIERC777Sender(address common.Address, backend bind.ContractBackend) (*IERC777Sender, error) {
	contract, err := bindIERC777Sender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC777Sender{IERC777SenderCaller: IERC777SenderCaller{contract: contract}, IERC777SenderTransactor: IERC777SenderTransactor{contract: contract}, IERC777SenderFilterer: IERC777SenderFilterer{contract: contract}}, nil
}

// NewIERC777SenderCaller creates a new read-only instance of IERC777Sender, bound to a specific deployed contract.
func NewIERC777SenderCaller(address common.Address, caller bind.ContractCaller) (*IERC777SenderCaller, error) {
	contract, err := bindIERC777Sender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777SenderCaller{contract: contract}, nil
}

// NewIERC777SenderTransactor creates a new write-only instance of IERC777Sender, bound to a specific deployed contract.
func NewIERC777SenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC777SenderTransactor, error) {
	contract, err := bindIERC777Sender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC777SenderTransactor{contract: contract}, nil
}

// NewIERC777SenderFilterer creates a new log filterer instance of IERC777Sender, bound to a specific deployed contract.
func NewIERC777SenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC777SenderFilterer, error) {
	contract, err := bindIERC777Sender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC777SenderFilterer{contract: contract}, nil
}

// bindIERC777Sender binds a generic wrapper to an already deployed contract.
func bindIERC777Sender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC777SenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777Sender *IERC777SenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777Sender.Contract.IERC777SenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777Sender *IERC777SenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777Sender.Contract.IERC777SenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777Sender *IERC777SenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777Sender.Contract.IERC777SenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC777Sender *IERC777SenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC777Sender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC777Sender *IERC777SenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC777Sender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC777Sender *IERC777SenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC777Sender.Contract.contract.Transact(opts, method, params...)
}

// TokensToSend is a paid mutator transaction binding the contract method 0x75ab9782.
//
// Solidity: function tokensToSend(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Sender *IERC777SenderTransactor) TokensToSend(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Sender.contract.Transact(opts, "tokensToSend", operator, from, to, amount, userData, operatorData)
}

// TokensToSend is a paid mutator transaction binding the contract method 0x75ab9782.
//
// Solidity: function tokensToSend(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Sender *IERC777SenderSession) TokensToSend(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Sender.Contract.TokensToSend(&_IERC777Sender.TransactOpts, operator, from, to, amount, userData, operatorData)
}

// TokensToSend is a paid mutator transaction binding the contract method 0x75ab9782.
//
// Solidity: function tokensToSend(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_IERC777Sender *IERC777SenderTransactorSession) TokensToSend(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _IERC777Sender.Contract.TokensToSend(&_IERC777Sender.TransactOpts, operator, from, to, amount, userData, operatorData)
}

// KeyBroadcastContractABI is the input ABI used to generate the binding from.
const KeyBroadcastContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keyper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numVotes\",\"type\":\"uint64\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"getBestKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"getBestKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"getBestKeyNumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getNumVotes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyper\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"keyperIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeyBroadcastContractBin is the compiled bytecode used for deploying new contracts.
var KeyBroadcastContractBin = "0x608060405234801561001057600080fd5b50604051610ca0380380610ca083398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b610c0f806100916000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636b32c7951161005b5780636b32c79514610137578063ab0a3ffe1461014c578063bf66a18214610175578063cf1d7a0a146101a057600080fd5b80631845bf5c146100825780632553600e146100dd57806341ecda09146100fd575b600080fd5b6100c06100903660046109e4565b6001600160401b039182166000908152600260209081526040808320845194830194909420835292905220541690565b6040516001600160401b0390911681526020015b60405180910390f35b6100f06100eb3660046109ac565b6101f1565b6040516100d49190610adb565b61012961010b3660046109ac565b6001600160401b039081166000908152600560205260409020541690565b6040519081526020016100d4565b61014a610145366004610a31565b6102aa565b005b61012961015a3660046109ac565b6001600160401b031660009081526004602052604090205490565b600054610188906001600160a01b031681565b6040516001600160a01b0390911681526020016100d4565b6101e16101ae366004610974565b6001600160401b031660009081526001602090815260408083206001600160a01b03949094168352929052205460ff1690565b60405190151581526020016100d4565b6001600160401b03811660009081526004602090815260408083205483526003909152902080546060919061022590610b5b565b80601f016020809104026020016040519081016040528092919081815260200182805461025190610b5b565b801561029e5780601f106102735761010080835404028352916020019161029e565b820191906000526020600020905b81548152906001019060200180831161028157829003601f168201915b50505050509050919050565b60008054604051632e37abf360e21b81526001600160401b03851660048201526001600160a01b039091169063b8deafcc9060240160206040518083038186803b1580156102f757600080fd5b505afa15801561030b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032f91906109c8565b6000805460405163017109f560e31b81526001600160401b038416600482015292935090916001600160a01b0390911690630b884fa89060240160206040518083038186803b15801561038157600080fd5b505afa158015610395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b991906109c8565b90506000816001600160401b03161161042a5760405162461bcd60e51b815260206004820152602860248201527f4b657942726f616463617374436f6e74726163743a20636f6e66696720697320604482015267696e61637469766560c01b60648201526084015b60405180910390fd5b60005460405163366963c960e21b81526001600160401b03841660048201526001600160a01b039091169063d9a58f249060240160206040518083038186803b15801561047657600080fd5b505afa15801561048a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ae91906109c8565b6001600160401b0316856001600160401b0316106105265760405162461bcd60e51b815260206004820152602f60248201527f4b657942726f616463617374436f6e74726163743a206b657970657220696e6460448201526e6578206f7574206f662072616e676560881b6064820152608401610421565b600054604051637d42750160e11b81526001600160401b038085166004830152871660248201526001600160a01b039091169063fa84ea029060440160206040518083038186803b15801561057a57600080fd5b505afa15801561058e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b29190610951565b6001600160a01b0316336001600160a01b0316146106255760405162461bcd60e51b815260206004820152602a60248201527f4b657942726f616463617374436f6e74726163743a2073656e646572206973206044820152693737ba1035b2bcb832b960b11b6064820152608401610421565b6001600160401b038416600090815260016020908152604080832033845290915290205460ff16156106b05760405162461bcd60e51b815260206004820152602e60248201527f4b657942726f616463617374436f6e74726163743a206b65797065722068617360448201526d08185b1c9958591e481d9bdd195960921b6064820152608401610421565b825160208085019190912060008181526003909252604090912080546106d590610b5b565b15905080156106e2575060015b15610708576000818152600360209081526040909120855161070692870190610832565b505b6001600160401b038086166000908152600260209081526040808320858452909152812054909161073b91166001610b24565b6001600160401b0387811660008181526001602081815260408084203385528252808420805460ff1916909317909255838352600281528183208884528152818320805467ffffffffffffffff191687871690811790915593835260059052902054929350911610156107e4576001600160401b038681166000908152600560209081526040808320805467ffffffffffffffff19169486169490941790935560049052208290555b336001600160a01b03167f305124b6ec831bb4150eb1ddbd4e8cc4b95687b9a6258b110fd0e9865914b0bf87878460405161082193929190610aee565b60405180910390a250505050505050565b82805461083e90610b5b565b90600052602060002090601f01602090048101928261086057600085556108a6565b82601f1061087957805160ff19168380011785556108a6565b828001600101855582156108a6579182015b828111156108a657825182559160200191906001019061088b565b506108b29291506108b6565b5090565b5b808211156108b257600081556001016108b7565b600082601f8301126108db578081fd5b81356001600160401b03808211156108f5576108f5610b96565b604051601f8301601f19908116603f0116810190828211818310171561091d5761091d610b96565b81604052838152866020858801011115610935578485fd5b8360208701602083013792830160200193909352509392505050565b600060208284031215610962578081fd5b815161096d81610bac565b9392505050565b60008060408385031215610986578081fd5b823561099181610bac565b915060208301356109a181610bc4565b809150509250929050565b6000602082840312156109bd578081fd5b813561096d81610bc4565b6000602082840312156109d9578081fd5b815161096d81610bc4565b600080604083850312156109f6578182fd5b8235610a0181610bc4565b915060208301356001600160401b03811115610a1b578182fd5b610a27858286016108cb565b9150509250929050565b600080600060608486031215610a45578081fd5b8335610a5081610bc4565b92506020840135610a6081610bc4565b915060408401356001600160401b03811115610a7a578182fd5b610a86868287016108cb565b9150509250925092565b60008151808452815b81811015610ab557602081850181015186830182015201610a99565b81811115610ac65782602083870101525b50601f01601f19169290920160200192915050565b60208152600061096d6020830184610a90565b60006001600160401b03808616835260606020840152610b116060840186610a90565b9150808416604084015250949350505050565b60006001600160401b03808316818516808303821115610b5257634e487b7160e01b84526011600452602484fd5b01949350505050565b600181811c90821680610b6f57607f821691505b60208210811415610b9057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610bc157600080fd5b50565b6001600160401b0381168114610bc157600080fdfea2646970667358221220bf7e2ff1347c1e47ff87d060abc9617968583d807e7492f62f8de85e8d52bee264736f6c63430008040033"

// DeployKeyBroadcastContract deploys a new Ethereum contract, binding an instance of KeyBroadcastContract to it.
func DeployKeyBroadcastContract(auth *bind.TransactOpts, backend bind.ContractBackend, configContractAddress common.Address) (common.Address, *types.Transaction, *KeyBroadcastContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyBroadcastContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyBroadcastContractBin), backend, configContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyBroadcastContract{KeyBroadcastContractCaller: KeyBroadcastContractCaller{contract: contract}, KeyBroadcastContractTransactor: KeyBroadcastContractTransactor{contract: contract}, KeyBroadcastContractFilterer: KeyBroadcastContractFilterer{contract: contract}}, nil
}

// KeyBroadcastContract is an auto generated Go binding around an Ethereum contract.
type KeyBroadcastContract struct {
	KeyBroadcastContractCaller     // Read-only binding to the contract
	KeyBroadcastContractTransactor // Write-only binding to the contract
	KeyBroadcastContractFilterer   // Log filterer for contract events
}

// KeyBroadcastContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyBroadcastContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyBroadcastContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyBroadcastContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyBroadcastContractSession struct {
	Contract     *KeyBroadcastContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeyBroadcastContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyBroadcastContractCallerSession struct {
	Contract *KeyBroadcastContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// KeyBroadcastContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyBroadcastContractTransactorSession struct {
	Contract     *KeyBroadcastContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// KeyBroadcastContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyBroadcastContractRaw struct {
	Contract *KeyBroadcastContract // Generic contract binding to access the raw methods on
}

// KeyBroadcastContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyBroadcastContractCallerRaw struct {
	Contract *KeyBroadcastContractCaller // Generic read-only contract binding to access the raw methods on
}

// KeyBroadcastContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyBroadcastContractTransactorRaw struct {
	Contract *KeyBroadcastContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyBroadcastContract creates a new instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContract(address common.Address, backend bind.ContractBackend) (*KeyBroadcastContract, error) {
	contract, err := bindKeyBroadcastContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContract{KeyBroadcastContractCaller: KeyBroadcastContractCaller{contract: contract}, KeyBroadcastContractTransactor: KeyBroadcastContractTransactor{contract: contract}, KeyBroadcastContractFilterer: KeyBroadcastContractFilterer{contract: contract}}, nil
}

// NewKeyBroadcastContractCaller creates a new read-only instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractCaller(address common.Address, caller bind.ContractCaller) (*KeyBroadcastContractCaller, error) {
	contract, err := bindKeyBroadcastContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractCaller{contract: contract}, nil
}

// NewKeyBroadcastContractTransactor creates a new write-only instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyBroadcastContractTransactor, error) {
	contract, err := bindKeyBroadcastContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractTransactor{contract: contract}, nil
}

// NewKeyBroadcastContractFilterer creates a new log filterer instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyBroadcastContractFilterer, error) {
	contract, err := bindKeyBroadcastContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractFilterer{contract: contract}, nil
}

// bindKeyBroadcastContract binds a generic wrapper to an already deployed contract.
func bindKeyBroadcastContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyBroadcastContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyBroadcastContract *KeyBroadcastContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyBroadcastContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyBroadcastContract *KeyBroadcastContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyBroadcastContract *KeyBroadcastContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.contract.Transact(opts, method, params...)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyBroadcastContract *KeyBroadcastContractSession) ConfigContract() (common.Address, error) {
	return _KeyBroadcastContract.Contract.ConfigContract(&_KeyBroadcastContract.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) ConfigContract() (common.Address, error) {
	return _KeyBroadcastContract.Contract.ConfigContract(&_KeyBroadcastContract.CallOpts)
}

// GetBestKey is a free data retrieval call binding the contract method 0x2553600e.
//
// Solidity: function getBestKey(uint64 startBatchIndex) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) GetBestKey(opts *bind.CallOpts, startBatchIndex uint64) ([]byte, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "getBestKey", startBatchIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBestKey is a free data retrieval call binding the contract method 0x2553600e.
//
// Solidity: function getBestKey(uint64 startBatchIndex) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractSession) GetBestKey(startBatchIndex uint64) ([]byte, error) {
	return _KeyBroadcastContract.Contract.GetBestKey(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetBestKey is a free data retrieval call binding the contract method 0x2553600e.
//
// Solidity: function getBestKey(uint64 startBatchIndex) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) GetBestKey(startBatchIndex uint64) ([]byte, error) {
	return _KeyBroadcastContract.Contract.GetBestKey(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetBestKeyHash is a free data retrieval call binding the contract method 0xab0a3ffe.
//
// Solidity: function getBestKeyHash(uint64 startBatchIndex) view returns(bytes32)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) GetBestKeyHash(opts *bind.CallOpts, startBatchIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "getBestKeyHash", startBatchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBestKeyHash is a free data retrieval call binding the contract method 0xab0a3ffe.
//
// Solidity: function getBestKeyHash(uint64 startBatchIndex) view returns(bytes32)
func (_KeyBroadcastContract *KeyBroadcastContractSession) GetBestKeyHash(startBatchIndex uint64) ([32]byte, error) {
	return _KeyBroadcastContract.Contract.GetBestKeyHash(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetBestKeyHash is a free data retrieval call binding the contract method 0xab0a3ffe.
//
// Solidity: function getBestKeyHash(uint64 startBatchIndex) view returns(bytes32)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) GetBestKeyHash(startBatchIndex uint64) ([32]byte, error) {
	return _KeyBroadcastContract.Contract.GetBestKeyHash(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetBestKeyNumVotes is a free data retrieval call binding the contract method 0x41ecda09.
//
// Solidity: function getBestKeyNumVotes(uint64 startBatchIndex) view returns(uint256)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) GetBestKeyNumVotes(opts *bind.CallOpts, startBatchIndex uint64) (*big.Int, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "getBestKeyNumVotes", startBatchIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBestKeyNumVotes is a free data retrieval call binding the contract method 0x41ecda09.
//
// Solidity: function getBestKeyNumVotes(uint64 startBatchIndex) view returns(uint256)
func (_KeyBroadcastContract *KeyBroadcastContractSession) GetBestKeyNumVotes(startBatchIndex uint64) (*big.Int, error) {
	return _KeyBroadcastContract.Contract.GetBestKeyNumVotes(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetBestKeyNumVotes is a free data retrieval call binding the contract method 0x41ecda09.
//
// Solidity: function getBestKeyNumVotes(uint64 startBatchIndex) view returns(uint256)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) GetBestKeyNumVotes(startBatchIndex uint64) (*big.Int, error) {
	return _KeyBroadcastContract.Contract.GetBestKeyNumVotes(&_KeyBroadcastContract.CallOpts, startBatchIndex)
}

// GetNumVotes is a free data retrieval call binding the contract method 0x1845bf5c.
//
// Solidity: function getNumVotes(uint64 startBatchIndex, bytes key) view returns(uint64)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) GetNumVotes(opts *bind.CallOpts, startBatchIndex uint64, key []byte) (uint64, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "getNumVotes", startBatchIndex, key)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumVotes is a free data retrieval call binding the contract method 0x1845bf5c.
//
// Solidity: function getNumVotes(uint64 startBatchIndex, bytes key) view returns(uint64)
func (_KeyBroadcastContract *KeyBroadcastContractSession) GetNumVotes(startBatchIndex uint64, key []byte) (uint64, error) {
	return _KeyBroadcastContract.Contract.GetNumVotes(&_KeyBroadcastContract.CallOpts, startBatchIndex, key)
}

// GetNumVotes is a free data retrieval call binding the contract method 0x1845bf5c.
//
// Solidity: function getNumVotes(uint64 startBatchIndex, bytes key) view returns(uint64)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) GetNumVotes(startBatchIndex uint64, key []byte) (uint64, error) {
	return _KeyBroadcastContract.Contract.GetNumVotes(&_KeyBroadcastContract.CallOpts, startBatchIndex, key)
}

// HasVoted is a free data retrieval call binding the contract method 0xcf1d7a0a.
//
// Solidity: function hasVoted(address keyper, uint64 startBatchIndex) view returns(bool)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) HasVoted(opts *bind.CallOpts, keyper common.Address, startBatchIndex uint64) (bool, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "hasVoted", keyper, startBatchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0xcf1d7a0a.
//
// Solidity: function hasVoted(address keyper, uint64 startBatchIndex) view returns(bool)
func (_KeyBroadcastContract *KeyBroadcastContractSession) HasVoted(keyper common.Address, startBatchIndex uint64) (bool, error) {
	return _KeyBroadcastContract.Contract.HasVoted(&_KeyBroadcastContract.CallOpts, keyper, startBatchIndex)
}

// HasVoted is a free data retrieval call binding the contract method 0xcf1d7a0a.
//
// Solidity: function hasVoted(address keyper, uint64 startBatchIndex) view returns(bool)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) HasVoted(keyper common.Address, startBatchIndex uint64) (bool, error) {
	return _KeyBroadcastContract.Contract.HasVoted(&_KeyBroadcastContract.CallOpts, keyper, startBatchIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x6b32c795.
//
// Solidity: function vote(uint64 keyperIndex, uint64 startBatchIndex, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactor) Vote(opts *bind.TransactOpts, keyperIndex uint64, startBatchIndex uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.contract.Transact(opts, "vote", keyperIndex, startBatchIndex, key)
}

// Vote is a paid mutator transaction binding the contract method 0x6b32c795.
//
// Solidity: function vote(uint64 keyperIndex, uint64 startBatchIndex, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractSession) Vote(keyperIndex uint64, startBatchIndex uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.Vote(&_KeyBroadcastContract.TransactOpts, keyperIndex, startBatchIndex, key)
}

// Vote is a paid mutator transaction binding the contract method 0x6b32c795.
//
// Solidity: function vote(uint64 keyperIndex, uint64 startBatchIndex, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactorSession) Vote(keyperIndex uint64, startBatchIndex uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.Vote(&_KeyBroadcastContract.TransactOpts, keyperIndex, startBatchIndex, key)
}

// KeyBroadcastContractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the KeyBroadcastContract contract.
type KeyBroadcastContractVotedIterator struct {
	Event *KeyBroadcastContractVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyBroadcastContractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyBroadcastContractVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyBroadcastContractVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyBroadcastContractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyBroadcastContractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyBroadcastContractVoted represents a Voted event raised by the KeyBroadcastContract contract.
type KeyBroadcastContractVoted struct {
	Keyper          common.Address
	StartBatchIndex uint64
	Key             []byte
	NumVotes        uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x305124b6ec831bb4150eb1ddbd4e8cc4b95687b9a6258b110fd0e9865914b0bf.
//
// Solidity: event Voted(address indexed keyper, uint64 startBatchIndex, bytes key, uint64 numVotes)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) FilterVoted(opts *bind.FilterOpts, keyper []common.Address) (*KeyBroadcastContractVotedIterator, error) {

	var keyperRule []interface{}
	for _, keyperItem := range keyper {
		keyperRule = append(keyperRule, keyperItem)
	}

	logs, sub, err := _KeyBroadcastContract.contract.FilterLogs(opts, "Voted", keyperRule)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractVotedIterator{contract: _KeyBroadcastContract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x305124b6ec831bb4150eb1ddbd4e8cc4b95687b9a6258b110fd0e9865914b0bf.
//
// Solidity: event Voted(address indexed keyper, uint64 startBatchIndex, bytes key, uint64 numVotes)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *KeyBroadcastContractVoted, keyper []common.Address) (event.Subscription, error) {

	var keyperRule []interface{}
	for _, keyperItem := range keyper {
		keyperRule = append(keyperRule, keyperItem)
	}

	logs, sub, err := _KeyBroadcastContract.contract.WatchLogs(opts, "Voted", keyperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyBroadcastContractVoted)
				if err := _KeyBroadcastContract.contract.UnpackLog(event, "Voted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoted is a log parse operation binding the contract event 0x305124b6ec831bb4150eb1ddbd4e8cc4b95687b9a6258b110fd0e9865914b0bf.
//
// Solidity: event Voted(address indexed keyper, uint64 startBatchIndex, bytes key, uint64 numVotes)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) ParseVoted(log types.Log) (*KeyBroadcastContractVoted, error) {
	event := new(KeyBroadcastContractVoted)
	if err := _KeyBroadcastContract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSlasherABI is the input ABI used to generate the binding from.
const KeyperSlasherABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appealPeriod\",\"type\":\"uint256\"},{\"internalType\":\"contractConfigContract\",\"name\":\"configContractAddress\",\"type\":\"address\"},{\"internalType\":\"contractExecutorContract\",\"name\":\"executorContractAddress\",\"type\":\"address\"},{\"internalType\":\"contractDepositContract\",\"name\":\"depositContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"}],\"name\":\"Accused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Appealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"accusations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"accused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"appealed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"keyperIndex\",\"type\":\"uint64\"}],\"name\":\"accuse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structAuthorization\",\"name\":\"authorization\",\"type\":\"tuple\"}],\"name\":\"appeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appealBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorContract\",\"outputs\":[{\"internalType\":\"contractExecutorContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeyperSlasherBin is the compiled bytecode used for deploying new contracts.
var KeyperSlasherBin = "0x60806040523480156200001157600080fd5b5060405162001cb138038062001cb18339810160408190526200003491620000ce565b6003849055600080546001600160a01b038581166001600160a01b0319928316179092556001805485841690831617905560028054928416929091168217905560405163555e124b60e11b815230600482015263aabc249690602401600060405180830381600087803b158015620000ab57600080fd5b505af1158015620000c0573d6000803e3d6000fd5b505050505050505062000142565b60008060008060808587031215620000e4578384fd5b845193506020850151620000f88162000129565b60408601519093506200010b8162000129565b60608601519092506200011e8162000129565b939692955090935050565b6001600160a01b03811681146200013f57600080fd5b50565b611b5f80620001526000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ab4dfa8a1161005b578063ab4dfa8a14610190578063bf66a182146101a3578063d02e74e6146101b6578063e94ad65b146101cd57600080fd5b80630be5fdf41461008d5780630e98ad4d146100bd57806331217be1146100d25780636864e7ee146100e5575b600080fd5b6001546100a0906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100d06100cb36600461178a565b6101e0565b005b6100d06100e036600461191c565b61049f565b6101496100f336600461191c565b6004602052600090815260409020805460019091015460ff808316926101008104821692620100008204909216916001600160a01b036301000000830416916001600160401b03600160b81b9091048116911686565b6040805196151587529415156020870152921515938501939093526001600160a01b031660608401526001600160401b0391821660808401521660a082015260c0016100b4565b6100d061019e366004611954565b610768565b6000546100a0906001600160a01b031681565b6101bf60035481565b6040519081526020016100b4565b6002546100a0906001600160a01b031681565b80516001600160401b03908116600090815260046020908152604091829020825160c081018452815460ff808216151580845261010083048216151595840195909552620100008204161515948201949094526001600160a01b0363010000008504166060820152600160b81b909304841660808401526001015490921660a0820152906102b55760405162461bcd60e51b815260206004820152601c60248201527f4b6579706572536c61736865723a206e6f2061636375736174696f6e0000000060448201526064015b60405180910390fd5b8060200151156103075760405162461bcd60e51b815260206004820152601f60248201527f4b6579706572536c61736865723a20616c72656164792061707065616c65640060448201526064016102ac565b60015482516040516325b36cbf60e01b81526001600160401b0390911660048201526000916001600160a01b0316906325b36cbf9060240160a06040518083038186803b15801561035757600080fd5b505afa15801561036b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038f9190611892565b905061039b8382610dd0565b6001602083810182815285516001600160401b039081166000908152600484526040808220885181549551838b015160608c015160808d015161ffff1990991693151561ff00191693909317610100921515929092029190911762010000600160b81b03191662010000911515919091026301000000600160b81b0319161763010000006001600160a01b03928316021767ffffffffffffffff60b81b1916600160b81b9686169690960295909517815560a08901519601805467ffffffffffffffff191696841696909617909555928501518751945192169316917f8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf91a3505050565b6001600160401b03808216600090815260046020908152604091829020825160c081018452815460ff808216151580845261010083048216151595840195909552620100008204161515948201949094526001600160a01b0363010000008504166060820152600160b81b909304841660808401526001015490921660a08201529061056d5760405162461bcd60e51b815260206004820152601c60248201527f4b6579706572536c61736865723a206e6f2061636375736174696f6e0000000060448201526064016102ac565b8060200151156105cb5760405162461bcd60e51b8152602060048201526024808201527f4b6579706572536c61736865723a207375636365737366756c6c792061707065604482015263185b195960e21b60648201526084016102ac565b80604001511561061d5760405162461bcd60e51b815260206004820152601e60248201527f4b6579706572536c61736865723a20616c726561647920736c6173686564000060448201526064016102ac565b6003548160a001516001600160401b03166106389190611a07565b4310156106995760405162461bcd60e51b815260206004820152602960248201527f4b6579706572536c61736865723a2061707065616c20706572696f64206e6f74604482015268081bdd995c881e595d60ba1b60648201526084016102ac565b600254606082015160405163c96be4cb60e01b81526001600160a01b03918216600482015291169063c96be4cb90602401600060405180830381600087803b1580156106e457600080fd5b505af11580156106f8573d6000803e3d6000fd5b5050506001600160401b03808416600090815260046020526040808220805462ff00001916620100001790556060850151608086015191516001600160a01b0390911694509216917fa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac59190a35050565b610773600283611a94565b6001600160401b0316156107d75760405162461bcd60e51b815260206004820152602560248201527f4b6579706572536c61736865723a206e6f742061206369706865722068616c66604482015264020737465760dc1b60648201526084016102ac565b6001600160401b03821660009081526004602052604090205460ff16156108405760405162461bcd60e51b815260206004820152601e60248201527f4b6579706572536c61736865723a20616c72656164792061636375736564000060448201526064016102ac565b600061084d600284611a1f565b60008054604051632e37abf360e21b81526001600160401b038416600482015292935090916001600160a01b039091169063b8deafcc9060240160206040518083038186803b15801561089f57600080fd5b505afa1580156108b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d79190611938565b60005460405163366963c960e21b81526001600160401b03831660048201529192506001600160a01b03169063d9a58f249060240160206040518083038186803b15801561092457600080fd5b505afa158015610938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095c9190611938565b6001600160401b0316836001600160401b0316106109cd5760405162461bcd60e51b815260206004820152602860248201527f4b6579706572536c61736865723a206b657970657220696e646578206f7574206044820152676f662072616e676560c01b60648201526084016102ac565b600054604051637d42750160e11b81526001600160401b038084166004830152851660248201526001600160a01b039091169063fa84ea029060440160206040518083038186803b158015610a2157600080fd5b505afa158015610a35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a599190611767565b6001600160a01b0316336001600160a01b031614610acd5760405162461bcd60e51b815260206004820152602b60248201527f4b6579706572536c61736865723a2073656e64657220646f6573206e6f74206d60448201526a30ba31b41035b2bcb832b960a91b60648201526084016102ac565b6001546040516325b36cbf60e01b81526001600160401b03861660048201526000916001600160a01b0316906325b36cbf9060240160a06040518083038186803b158015610b1a57600080fd5b505afa158015610b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b529190611892565b8051909150610bb55760405162461bcd60e51b815260206004820152602960248201527f4b6579706572536c61736865723a2068616c662073746570206e6f742079657460448201526808195e1958dd5d195960ba1b60648201526084016102ac565b6060810151610c175760405162461bcd60e51b815260206004820152602860248201527f4b6579706572536c61736865723a2063616e6e6f742061636375736520656d706044820152670e8f240c4c2e8c6d60c31b60648201526084016102ac565b6040518060c0016040528060011515815260200160001515815260200160001515815260200182602001516001600160a01b03168152602001866001600160401b03168152602001436001600160401b031681525060046000876001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160000160026101000a81548160ff02191690831515021790555060608201518160000160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060808201518160000160176101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160010160006101000a8154816001600160401b0302191690836001600160401b03160217905550905050336001600160a01b031681602001516001600160a01b0316866001600160401b03167f79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f1760405160405180910390a45050505050565b600060028260400151610de39190611a1f565b60008054604051632e37abf360e21b81526001600160401b038416600482015292935090916001600160a01b039091169063b8deafcc9060240160206040518083038186803b158015610e3557600080fd5b505afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d9190611938565b60005460405163fad07bfb60e01b81526001600160401b03831660048201529192506001600160a01b03169063fad07bfb9060240160206040518083038186803b158015610eba57600080fd5b505afa158015610ece573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef29190611938565b6001600160401b03168460600151511015610f5b5760405162461bcd60e51b8152602060048201526024808201527f4b6579706572536c61736865723a206e6f7420656e6f756768207369676e61746044820152637572657360e01b60648201526084016102ac565b83604001515184606001515114610fda5760405162461bcd60e51b815260206004820152603e60248201527f4b6579706572536c61736865723a206e756d626572206f66207369676e61747560448201527f72657320616e6420696e646963657320646f6573206e6f74206d61746368000060648201526084016102ac565b600061106b600160009054906101000a90046001600160a01b03166001600160a01b031663beb3b50e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561102d57600080fd5b505afa158015611041573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110659190611767565b856113d0565b80519060200120905060005b856060015151816001600160401b031610156113c85760008660600151826001600160401b0316815181106110bc57634e487b7160e01b600052603260045260246000fd5b6020026020010151905060008760400151836001600160401b0316815181106110f557634e487b7160e01b600052603260045260246000fd5b602090810291909101015160005460405163366963c960e21b81526001600160401b03881660048201529192506001600160a01b03169063d9a58f249060240160206040518083038186803b15801561114d57600080fd5b505afa158015611161573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111859190611938565b6001600160401b0316816001600160401b0316106111f65760405162461bcd60e51b815260206004820152602860248201527f4b6579706572536c61736865723a207369676e657220696e646578206f7574206044820152676f662072616e676560c01b60648201526084016102ac565b6001600160401b038316158061125b57506040880151611217600185611a45565b6001600160401b03168151811061123e57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160401b0316816001600160401b0316115b6112b95760405162461bcd60e51b815260206004820152602960248201527f4b6579706572536c61736865723a207369676e657220696e6469636573206e6f6044820152681d081bdc99195c995960ba1b60648201526084016102ac565b60006112c5858461145e565b600054604051637d42750160e11b81526001600160401b03808a166004830152851660248201529192506001600160a01b03169063fa84ea029060440160206040518083038186803b15801561131a57600080fd5b505afa15801561132e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113529190611767565b6001600160a01b0316816001600160a01b0316146113b25760405162461bcd60e51b815260206004820152601b60248201527f4b6579706572536c61736865723a2077726f6e67207369676e6572000000000060448201526064016102ac565b50505080806113c090611a6d565b915050611077565b505050505050565b606082600283604001516113e49190611a1f565b83606001518460800151604051602001611447949392919065032c8cac6e8f60d31b815260609490941b6bffffffffffffffffffffffff1916600685015260c09290921b6001600160c01b031916601a8401526022830152604282015260620190565b604051602081830303815290604052905092915050565b600081516041146114b15760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016102ac565b60208201516040830151606084015160001a6114cf868285856114d9565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156115565760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016102ac565b8360ff16601b148061156b57508360ff16601c145b6115c25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016102ac565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611616573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166116795760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016102ac565b95945050505050565b6000601f8381840112611693578182fd5b823560206116a86116a3836119e4565b6119b4565b80838252828201915082870188848660051b8a010111156116c7578687fd5b865b858110156117595781356001600160401b03808211156116e757898afd5b818b0191508b603f8301126116fa57898afd5b8682013560408282111561171057611710611ae6565b611721828c01601f19168a016119b4565b92508183528d81838601011115611736578b8cfd5b818185018a85013750810187018a905285525092840192908401906001016116c9565b509098975050505050505050565b600060208284031215611778578081fd5b815161178381611afc565b9392505050565b6000602080838503121561179c578182fd5b82356001600160401b03808211156117b2578384fd5b90840190608082870312156117c5578384fd5b6117cd61198c565b82356117d881611b14565b815282840135848201526040830135828111156117f3578586fd5b8301601f81018813611803578586fd5b80356118116116a3826119e4565b8082825287820191508784018b898560051b870101111561183057898afd5b8994505b8385101561185b57803561184781611b14565b835260019490940193918801918801611834565b5060408501525050506060830135935081841115611877578485fd5b61188387858501611682565b60608201529695505050505050565b600060a082840312156118a3578081fd5b60405160a081018181106001600160401b03821117156118c5576118c5611ae6565b604052825180151581146118d7578283fd5b815260208301516118e781611afc565b602082015260408301516118fa81611b14565b6040820152606083810151908201526080928301519281019290925250919050565b60006020828403121561192d578081fd5b813561178381611b14565b600060208284031215611949578081fd5b815161178381611b14565b60008060408385031215611966578081fd5b823561197181611b14565b9150602083013561198181611b14565b809150509250929050565b604051608081016001600160401b03811182821017156119ae576119ae611ae6565b60405290565b604051601f8201601f191681016001600160401b03811182821017156119dc576119dc611ae6565b604052919050565b60006001600160401b038211156119fd576119fd611ae6565b5060051b60200190565b60008219821115611a1a57611a1a611aba565b500190565b60006001600160401b0380841680611a3957611a39611ad0565b92169190910492915050565b60006001600160401b0383811690831681811015611a6557611a65611aba565b039392505050565b60006001600160401b0380831681811415611a8a57611a8a611aba565b6001019392505050565b60006001600160401b0380841680611aae57611aae611ad0565b92169190910692915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611b1157600080fd5b50565b6001600160401b0381168114611b1157600080fdfea26469706673582212204c7c858e328e5fb3178725f251bd1d077091a4abd9ea821af274227c24fce7bc64736f6c63430008040033"

// DeployKeyperSlasher deploys a new Ethereum contract, binding an instance of KeyperSlasher to it.
func DeployKeyperSlasher(auth *bind.TransactOpts, backend bind.ContractBackend, appealPeriod *big.Int, configContractAddress common.Address, executorContractAddress common.Address, depositContractAddress common.Address) (common.Address, *types.Transaction, *KeyperSlasher, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyperSlasherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyperSlasherBin), backend, appealPeriod, configContractAddress, executorContractAddress, depositContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyperSlasher{KeyperSlasherCaller: KeyperSlasherCaller{contract: contract}, KeyperSlasherTransactor: KeyperSlasherTransactor{contract: contract}, KeyperSlasherFilterer: KeyperSlasherFilterer{contract: contract}}, nil
}

// KeyperSlasher is an auto generated Go binding around an Ethereum contract.
type KeyperSlasher struct {
	KeyperSlasherCaller     // Read-only binding to the contract
	KeyperSlasherTransactor // Write-only binding to the contract
	KeyperSlasherFilterer   // Log filterer for contract events
}

// KeyperSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyperSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyperSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyperSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyperSlasherSession struct {
	Contract     *KeyperSlasher    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyperSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyperSlasherCallerSession struct {
	Contract *KeyperSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KeyperSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyperSlasherTransactorSession struct {
	Contract     *KeyperSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KeyperSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyperSlasherRaw struct {
	Contract *KeyperSlasher // Generic contract binding to access the raw methods on
}

// KeyperSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyperSlasherCallerRaw struct {
	Contract *KeyperSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// KeyperSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyperSlasherTransactorRaw struct {
	Contract *KeyperSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyperSlasher creates a new instance of KeyperSlasher, bound to a specific deployed contract.
func NewKeyperSlasher(address common.Address, backend bind.ContractBackend) (*KeyperSlasher, error) {
	contract, err := bindKeyperSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasher{KeyperSlasherCaller: KeyperSlasherCaller{contract: contract}, KeyperSlasherTransactor: KeyperSlasherTransactor{contract: contract}, KeyperSlasherFilterer: KeyperSlasherFilterer{contract: contract}}, nil
}

// NewKeyperSlasherCaller creates a new read-only instance of KeyperSlasher, bound to a specific deployed contract.
func NewKeyperSlasherCaller(address common.Address, caller bind.ContractCaller) (*KeyperSlasherCaller, error) {
	contract, err := bindKeyperSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherCaller{contract: contract}, nil
}

// NewKeyperSlasherTransactor creates a new write-only instance of KeyperSlasher, bound to a specific deployed contract.
func NewKeyperSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyperSlasherTransactor, error) {
	contract, err := bindKeyperSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherTransactor{contract: contract}, nil
}

// NewKeyperSlasherFilterer creates a new log filterer instance of KeyperSlasher, bound to a specific deployed contract.
func NewKeyperSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyperSlasherFilterer, error) {
	contract, err := bindKeyperSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherFilterer{contract: contract}, nil
}

// bindKeyperSlasher binds a generic wrapper to an already deployed contract.
func bindKeyperSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyperSlasherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSlasher *KeyperSlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSlasher.Contract.KeyperSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSlasher *KeyperSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.KeyperSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSlasher *KeyperSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.KeyperSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSlasher *KeyperSlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSlasher *KeyperSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSlasher *KeyperSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.contract.Transact(opts, method, params...)
}

// Accusations is a free data retrieval call binding the contract method 0x6864e7ee.
//
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, bool slashed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherCaller) Accusations(opts *bind.CallOpts, arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
	Slashed     bool
	Executor    common.Address
	HalfStep    uint64
	BlockNumber uint64
}, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "accusations", arg0)

	outstruct := new(struct {
		Accused     bool
		Appealed    bool
		Slashed     bool
		Executor    common.Address
		HalfStep    uint64
		BlockNumber uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Appealed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Slashed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Executor = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.HalfStep = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.BlockNumber = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// Accusations is a free data retrieval call binding the contract method 0x6864e7ee.
//
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, bool slashed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherSession) Accusations(arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
	Slashed     bool
	Executor    common.Address
	HalfStep    uint64
	BlockNumber uint64
}, error) {
	return _KeyperSlasher.Contract.Accusations(&_KeyperSlasher.CallOpts, arg0)
}

// Accusations is a free data retrieval call binding the contract method 0x6864e7ee.
//
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, bool slashed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherCallerSession) Accusations(arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
	Slashed     bool
	Executor    common.Address
	HalfStep    uint64
	BlockNumber uint64
}, error) {
	return _KeyperSlasher.Contract.Accusations(&_KeyperSlasher.CallOpts, arg0)
}

// AppealBlocks is a free data retrieval call binding the contract method 0xd02e74e6.
//
// Solidity: function appealBlocks() view returns(uint256)
func (_KeyperSlasher *KeyperSlasherCaller) AppealBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "appealBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppealBlocks is a free data retrieval call binding the contract method 0xd02e74e6.
//
// Solidity: function appealBlocks() view returns(uint256)
func (_KeyperSlasher *KeyperSlasherSession) AppealBlocks() (*big.Int, error) {
	return _KeyperSlasher.Contract.AppealBlocks(&_KeyperSlasher.CallOpts)
}

// AppealBlocks is a free data retrieval call binding the contract method 0xd02e74e6.
//
// Solidity: function appealBlocks() view returns(uint256)
func (_KeyperSlasher *KeyperSlasherCallerSession) AppealBlocks() (*big.Int, error) {
	return _KeyperSlasher.Contract.AppealBlocks(&_KeyperSlasher.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherSession) ConfigContract() (common.Address, error) {
	return _KeyperSlasher.Contract.ConfigContract(&_KeyperSlasher.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCallerSession) ConfigContract() (common.Address, error) {
	return _KeyperSlasher.Contract.ConfigContract(&_KeyperSlasher.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherSession) DepositContract() (common.Address, error) {
	return _KeyperSlasher.Contract.DepositContract(&_KeyperSlasher.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCallerSession) DepositContract() (common.Address, error) {
	return _KeyperSlasher.Contract.DepositContract(&_KeyperSlasher.CallOpts)
}

// ExecutorContract is a free data retrieval call binding the contract method 0x0be5fdf4.
//
// Solidity: function executorContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCaller) ExecutorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "executorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorContract is a free data retrieval call binding the contract method 0x0be5fdf4.
//
// Solidity: function executorContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherSession) ExecutorContract() (common.Address, error) {
	return _KeyperSlasher.Contract.ExecutorContract(&_KeyperSlasher.CallOpts)
}

// ExecutorContract is a free data retrieval call binding the contract method 0x0be5fdf4.
//
// Solidity: function executorContract() view returns(address)
func (_KeyperSlasher *KeyperSlasherCallerSession) ExecutorContract() (common.Address, error) {
	return _KeyperSlasher.Contract.ExecutorContract(&_KeyperSlasher.CallOpts)
}

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 halfStep, uint64 keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Accuse(opts *bind.TransactOpts, halfStep uint64, keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "accuse", halfStep, keyperIndex)
}

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 halfStep, uint64 keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherSession) Accuse(halfStep uint64, keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Accuse(&_KeyperSlasher.TransactOpts, halfStep, keyperIndex)
}

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 halfStep, uint64 keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Accuse(halfStep uint64, keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Accuse(&_KeyperSlasher.TransactOpts, halfStep, keyperIndex)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) authorization) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Appeal(opts *bind.TransactOpts, authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "appeal", authorization)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) authorization) returns()
func (_KeyperSlasher *KeyperSlasherSession) Appeal(authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Appeal(&_KeyperSlasher.TransactOpts, authorization)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) authorization) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Appeal(authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Appeal(&_KeyperSlasher.TransactOpts, authorization)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 halfStep) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Slash(opts *bind.TransactOpts, halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "slash", halfStep)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 halfStep) returns()
func (_KeyperSlasher *KeyperSlasherSession) Slash(halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Slash(&_KeyperSlasher.TransactOpts, halfStep)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 halfStep) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Slash(halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Slash(&_KeyperSlasher.TransactOpts, halfStep)
}

// KeyperSlasherAccusedIterator is returned from FilterAccused and is used to iterate over the raw logs and unpacked data for Accused events raised by the KeyperSlasher contract.
type KeyperSlasherAccusedIterator struct {
	Event *KeyperSlasherAccused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyperSlasherAccusedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSlasherAccused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyperSlasherAccused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyperSlasherAccusedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSlasherAccusedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSlasherAccused represents a Accused event raised by the KeyperSlasher contract.
type KeyperSlasherAccused struct {
	HalfStep uint64
	Executor common.Address
	Accuser  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccused is a free log retrieval operation binding the contract event 0x79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f17.
//
// Solidity: event Accused(uint64 indexed halfStep, address indexed executor, address indexed accuser)
func (_KeyperSlasher *KeyperSlasherFilterer) FilterAccused(opts *bind.FilterOpts, halfStep []uint64, executor []common.Address, accuser []common.Address) (*KeyperSlasherAccusedIterator, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _KeyperSlasher.contract.FilterLogs(opts, "Accused", halfStepRule, executorRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherAccusedIterator{contract: _KeyperSlasher.contract, event: "Accused", logs: logs, sub: sub}, nil
}

// WatchAccused is a free log subscription operation binding the contract event 0x79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f17.
//
// Solidity: event Accused(uint64 indexed halfStep, address indexed executor, address indexed accuser)
func (_KeyperSlasher *KeyperSlasherFilterer) WatchAccused(opts *bind.WatchOpts, sink chan<- *KeyperSlasherAccused, halfStep []uint64, executor []common.Address, accuser []common.Address) (event.Subscription, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _KeyperSlasher.contract.WatchLogs(opts, "Accused", halfStepRule, executorRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSlasherAccused)
				if err := _KeyperSlasher.contract.UnpackLog(event, "Accused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccused is a log parse operation binding the contract event 0x79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f17.
//
// Solidity: event Accused(uint64 indexed halfStep, address indexed executor, address indexed accuser)
func (_KeyperSlasher *KeyperSlasherFilterer) ParseAccused(log types.Log) (*KeyperSlasherAccused, error) {
	event := new(KeyperSlasherAccused)
	if err := _KeyperSlasher.contract.UnpackLog(event, "Accused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSlasherAppealedIterator is returned from FilterAppealed and is used to iterate over the raw logs and unpacked data for Appealed events raised by the KeyperSlasher contract.
type KeyperSlasherAppealedIterator struct {
	Event *KeyperSlasherAppealed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyperSlasherAppealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSlasherAppealed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyperSlasherAppealed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyperSlasherAppealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSlasherAppealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSlasherAppealed represents a Appealed event raised by the KeyperSlasher contract.
type KeyperSlasherAppealed struct {
	HalfStep uint64
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAppealed is a free log retrieval operation binding the contract event 0x8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf.
//
// Solidity: event Appealed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) FilterAppealed(opts *bind.FilterOpts, halfStep []uint64, executor []common.Address) (*KeyperSlasherAppealedIterator, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _KeyperSlasher.contract.FilterLogs(opts, "Appealed", halfStepRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherAppealedIterator{contract: _KeyperSlasher.contract, event: "Appealed", logs: logs, sub: sub}, nil
}

// WatchAppealed is a free log subscription operation binding the contract event 0x8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf.
//
// Solidity: event Appealed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) WatchAppealed(opts *bind.WatchOpts, sink chan<- *KeyperSlasherAppealed, halfStep []uint64, executor []common.Address) (event.Subscription, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _KeyperSlasher.contract.WatchLogs(opts, "Appealed", halfStepRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSlasherAppealed)
				if err := _KeyperSlasher.contract.UnpackLog(event, "Appealed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppealed is a log parse operation binding the contract event 0x8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf.
//
// Solidity: event Appealed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) ParseAppealed(log types.Log) (*KeyperSlasherAppealed, error) {
	event := new(KeyperSlasherAppealed)
	if err := _KeyperSlasher.contract.UnpackLog(event, "Appealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSlasherSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the KeyperSlasher contract.
type KeyperSlasherSlashedIterator struct {
	Event *KeyperSlasherSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyperSlasherSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSlasherSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyperSlasherSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyperSlasherSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSlasherSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSlasherSlashed represents a Slashed event raised by the KeyperSlasher contract.
type KeyperSlasherSlashed struct {
	HalfStep uint64
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0xa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac5.
//
// Solidity: event Slashed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) FilterSlashed(opts *bind.FilterOpts, halfStep []uint64, executor []common.Address) (*KeyperSlasherSlashedIterator, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _KeyperSlasher.contract.FilterLogs(opts, "Slashed", halfStepRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSlasherSlashedIterator{contract: _KeyperSlasher.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0xa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac5.
//
// Solidity: event Slashed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *KeyperSlasherSlashed, halfStep []uint64, executor []common.Address) (event.Subscription, error) {

	var halfStepRule []interface{}
	for _, halfStepItem := range halfStep {
		halfStepRule = append(halfStepRule, halfStepItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _KeyperSlasher.contract.WatchLogs(opts, "Slashed", halfStepRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSlasherSlashed)
				if err := _KeyperSlasher.contract.UnpackLog(event, "Slashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashed is a log parse operation binding the contract event 0xa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac5.
//
// Solidity: event Slashed(uint64 indexed halfStep, address indexed executor)
func (_KeyperSlasher *KeyperSlasherFilterer) ParseSlashed(log types.Log) (*KeyperSlasherSlashed, error) {
	event := new(KeyperSlasherSlashed)
	if err := _KeyperSlasher.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockBatcherContractABI is the input ABI used to generate the binding from.
const MockBatcherContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"setBatchHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockBatcherContractBin is the compiled bytecode used for deploying new contracts.
var MockBatcherContractBin = "0x608060405234801561001057600080fd5b506101c9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ad15b6c51461003b578063c87afa8a14610050575b600080fd5b61004e610049366004610158565b61008a565b005b61007861005e366004610126565b600060208181529281526040808220909352908152205481565b60405190815260200160405180910390f35b67ffffffffffffffff8316600090815260208190526040812082918460018111156100c557634e487b7160e01b600052602160045260246000fd5b60018111156100e457634e487b7160e01b600052602160045260246000fd5b8152602081019190915260400160002055505050565b80356002811061010957600080fd5b919050565b803567ffffffffffffffff8116811461010957600080fd5b60008060408385031215610138578182fd5b6101418361010e565b915061014f602084016100fa565b90509250929050565b60008060006060848603121561016c578081fd5b6101758461010e565b9250610183602085016100fa565b915060408401359050925092509256fea264697066735822122089effc56eee16483cc65e484e45eeab28a69e794b03f9e7eb808aac97928809964736f6c63430008040033"

// DeployMockBatcherContract deploys a new Ethereum contract, binding an instance of MockBatcherContract to it.
func DeployMockBatcherContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockBatcherContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockBatcherContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockBatcherContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockBatcherContract{MockBatcherContractCaller: MockBatcherContractCaller{contract: contract}, MockBatcherContractTransactor: MockBatcherContractTransactor{contract: contract}, MockBatcherContractFilterer: MockBatcherContractFilterer{contract: contract}}, nil
}

// MockBatcherContract is an auto generated Go binding around an Ethereum contract.
type MockBatcherContract struct {
	MockBatcherContractCaller     // Read-only binding to the contract
	MockBatcherContractTransactor // Write-only binding to the contract
	MockBatcherContractFilterer   // Log filterer for contract events
}

// MockBatcherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockBatcherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockBatcherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockBatcherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockBatcherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockBatcherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockBatcherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockBatcherContractSession struct {
	Contract     *MockBatcherContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MockBatcherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockBatcherContractCallerSession struct {
	Contract *MockBatcherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MockBatcherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockBatcherContractTransactorSession struct {
	Contract     *MockBatcherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MockBatcherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockBatcherContractRaw struct {
	Contract *MockBatcherContract // Generic contract binding to access the raw methods on
}

// MockBatcherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockBatcherContractCallerRaw struct {
	Contract *MockBatcherContractCaller // Generic read-only contract binding to access the raw methods on
}

// MockBatcherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockBatcherContractTransactorRaw struct {
	Contract *MockBatcherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockBatcherContract creates a new instance of MockBatcherContract, bound to a specific deployed contract.
func NewMockBatcherContract(address common.Address, backend bind.ContractBackend) (*MockBatcherContract, error) {
	contract, err := bindMockBatcherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockBatcherContract{MockBatcherContractCaller: MockBatcherContractCaller{contract: contract}, MockBatcherContractTransactor: MockBatcherContractTransactor{contract: contract}, MockBatcherContractFilterer: MockBatcherContractFilterer{contract: contract}}, nil
}

// NewMockBatcherContractCaller creates a new read-only instance of MockBatcherContract, bound to a specific deployed contract.
func NewMockBatcherContractCaller(address common.Address, caller bind.ContractCaller) (*MockBatcherContractCaller, error) {
	contract, err := bindMockBatcherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockBatcherContractCaller{contract: contract}, nil
}

// NewMockBatcherContractTransactor creates a new write-only instance of MockBatcherContract, bound to a specific deployed contract.
func NewMockBatcherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MockBatcherContractTransactor, error) {
	contract, err := bindMockBatcherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockBatcherContractTransactor{contract: contract}, nil
}

// NewMockBatcherContractFilterer creates a new log filterer instance of MockBatcherContract, bound to a specific deployed contract.
func NewMockBatcherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MockBatcherContractFilterer, error) {
	contract, err := bindMockBatcherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockBatcherContractFilterer{contract: contract}, nil
}

// bindMockBatcherContract binds a generic wrapper to an already deployed contract.
func bindMockBatcherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockBatcherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockBatcherContract *MockBatcherContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockBatcherContract.Contract.MockBatcherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockBatcherContract *MockBatcherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.MockBatcherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockBatcherContract *MockBatcherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.MockBatcherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockBatcherContract *MockBatcherContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockBatcherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockBatcherContract *MockBatcherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockBatcherContract *MockBatcherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.contract.Transact(opts, method, params...)
}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_MockBatcherContract *MockBatcherContractCaller) BatchHashes(opts *bind.CallOpts, arg0 uint64, arg1 uint8) ([32]byte, error) {
	var out []interface{}
	err := _MockBatcherContract.contract.Call(opts, &out, "batchHashes", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_MockBatcherContract *MockBatcherContractSession) BatchHashes(arg0 uint64, arg1 uint8) ([32]byte, error) {
	return _MockBatcherContract.Contract.BatchHashes(&_MockBatcherContract.CallOpts, arg0, arg1)
}

// BatchHashes is a free data retrieval call binding the contract method 0xc87afa8a.
//
// Solidity: function batchHashes(uint64 , uint8 ) view returns(bytes32)
func (_MockBatcherContract *MockBatcherContractCallerSession) BatchHashes(arg0 uint64, arg1 uint8) ([32]byte, error) {
	return _MockBatcherContract.Contract.BatchHashes(&_MockBatcherContract.CallOpts, arg0, arg1)
}

// SetBatchHash is a paid mutator transaction binding the contract method 0xad15b6c5.
//
// Solidity: function setBatchHash(uint64 batchIndex, uint8 transactionType, bytes32 batchHash) returns()
func (_MockBatcherContract *MockBatcherContractTransactor) SetBatchHash(opts *bind.TransactOpts, batchIndex uint64, transactionType uint8, batchHash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.contract.Transact(opts, "setBatchHash", batchIndex, transactionType, batchHash)
}

// SetBatchHash is a paid mutator transaction binding the contract method 0xad15b6c5.
//
// Solidity: function setBatchHash(uint64 batchIndex, uint8 transactionType, bytes32 batchHash) returns()
func (_MockBatcherContract *MockBatcherContractSession) SetBatchHash(batchIndex uint64, transactionType uint8, batchHash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.SetBatchHash(&_MockBatcherContract.TransactOpts, batchIndex, transactionType, batchHash)
}

// SetBatchHash is a paid mutator transaction binding the contract method 0xad15b6c5.
//
// Solidity: function setBatchHash(uint64 batchIndex, uint8 transactionType, bytes32 batchHash) returns()
func (_MockBatcherContract *MockBatcherContractTransactorSession) SetBatchHash(batchIndex uint64, transactionType uint8, batchHash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.SetBatchHash(&_MockBatcherContract.TransactOpts, batchIndex, transactionType, batchHash)
}

// MockTargetContractABI is the input ABI used to generate the binding from.
const MockTargetContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"Called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockTargetContractBin is the compiled bytecode used for deploying new contracts.
var MockTargetContractBin = "0x608060405234801561001057600080fd5b50610165806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635a6535fc14610030575b600080fd5b61004361003e36600461008a565b610045565b005b60005a90507fef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf9283838360405161007d939291906100f7565b60405180910390a1505050565b6000806020838503121561009c578182fd5b823567ffffffffffffffff808211156100b3578384fd5b818501915085601f8301126100c6578384fd5b8135818111156100d4578485fd5b8660208285010111156100e5578485fd5b60209290920196919550909350505050565b6040815282604082015282846060830137600080606085840101526060601f19601f860116830101905082602083015294935050505056fea26469706673582212208c484afc332cd5271509f74a4088404fa5a23c033d96e3ec6e453aeeadaeb22664736f6c63430008040033"

// DeployMockTargetContract deploys a new Ethereum contract, binding an instance of MockTargetContract to it.
func DeployMockTargetContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTargetContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockTargetContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockTargetContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTargetContract{MockTargetContractCaller: MockTargetContractCaller{contract: contract}, MockTargetContractTransactor: MockTargetContractTransactor{contract: contract}, MockTargetContractFilterer: MockTargetContractFilterer{contract: contract}}, nil
}

// MockTargetContract is an auto generated Go binding around an Ethereum contract.
type MockTargetContract struct {
	MockTargetContractCaller     // Read-only binding to the contract
	MockTargetContractTransactor // Write-only binding to the contract
	MockTargetContractFilterer   // Log filterer for contract events
}

// MockTargetContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockTargetContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTargetContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTargetContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTargetContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTargetContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTargetContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTargetContractSession struct {
	Contract     *MockTargetContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockTargetContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTargetContractCallerSession struct {
	Contract *MockTargetContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockTargetContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTargetContractTransactorSession struct {
	Contract     *MockTargetContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockTargetContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockTargetContractRaw struct {
	Contract *MockTargetContract // Generic contract binding to access the raw methods on
}

// MockTargetContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTargetContractCallerRaw struct {
	Contract *MockTargetContractCaller // Generic read-only contract binding to access the raw methods on
}

// MockTargetContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTargetContractTransactorRaw struct {
	Contract *MockTargetContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTargetContract creates a new instance of MockTargetContract, bound to a specific deployed contract.
func NewMockTargetContract(address common.Address, backend bind.ContractBackend) (*MockTargetContract, error) {
	contract, err := bindMockTargetContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTargetContract{MockTargetContractCaller: MockTargetContractCaller{contract: contract}, MockTargetContractTransactor: MockTargetContractTransactor{contract: contract}, MockTargetContractFilterer: MockTargetContractFilterer{contract: contract}}, nil
}

// NewMockTargetContractCaller creates a new read-only instance of MockTargetContract, bound to a specific deployed contract.
func NewMockTargetContractCaller(address common.Address, caller bind.ContractCaller) (*MockTargetContractCaller, error) {
	contract, err := bindMockTargetContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTargetContractCaller{contract: contract}, nil
}

// NewMockTargetContractTransactor creates a new write-only instance of MockTargetContract, bound to a specific deployed contract.
func NewMockTargetContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTargetContractTransactor, error) {
	contract, err := bindMockTargetContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTargetContractTransactor{contract: contract}, nil
}

// NewMockTargetContractFilterer creates a new log filterer instance of MockTargetContract, bound to a specific deployed contract.
func NewMockTargetContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MockTargetContractFilterer, error) {
	contract, err := bindMockTargetContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTargetContractFilterer{contract: contract}, nil
}

// bindMockTargetContract binds a generic wrapper to an already deployed contract.
func bindMockTargetContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockTargetContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTargetContract *MockTargetContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTargetContract.Contract.MockTargetContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTargetContract *MockTargetContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTargetContract.Contract.MockTargetContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTargetContract *MockTargetContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTargetContract.Contract.MockTargetContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTargetContract *MockTargetContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTargetContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTargetContract *MockTargetContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTargetContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTargetContract *MockTargetContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTargetContract.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x5a6535fc.
//
// Solidity: function call(bytes transaction) returns()
func (_MockTargetContract *MockTargetContractTransactor) Call(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _MockTargetContract.contract.Transact(opts, "call", transaction)
}

// Call is a paid mutator transaction binding the contract method 0x5a6535fc.
//
// Solidity: function call(bytes transaction) returns()
func (_MockTargetContract *MockTargetContractSession) Call(transaction []byte) (*types.Transaction, error) {
	return _MockTargetContract.Contract.Call(&_MockTargetContract.TransactOpts, transaction)
}

// Call is a paid mutator transaction binding the contract method 0x5a6535fc.
//
// Solidity: function call(bytes transaction) returns()
func (_MockTargetContract *MockTargetContractTransactorSession) Call(transaction []byte) (*types.Transaction, error) {
	return _MockTargetContract.Contract.Call(&_MockTargetContract.TransactOpts, transaction)
}

// MockTargetContractCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the MockTargetContract contract.
type MockTargetContractCalledIterator struct {
	Event *MockTargetContractCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockTargetContractCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTargetContractCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockTargetContractCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockTargetContractCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTargetContractCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTargetContractCalled represents a Called event raised by the MockTargetContract contract.
type MockTargetContractCalled struct {
	Transaction []byte
	Gas         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf92.
//
// Solidity: event Called(bytes transaction, uint256 gas)
func (_MockTargetContract *MockTargetContractFilterer) FilterCalled(opts *bind.FilterOpts) (*MockTargetContractCalledIterator, error) {

	logs, sub, err := _MockTargetContract.contract.FilterLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return &MockTargetContractCalledIterator{contract: _MockTargetContract.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf92.
//
// Solidity: event Called(bytes transaction, uint256 gas)
func (_MockTargetContract *MockTargetContractFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *MockTargetContractCalled) (event.Subscription, error) {

	logs, sub, err := _MockTargetContract.contract.WatchLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTargetContractCalled)
				if err := _MockTargetContract.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled is a log parse operation binding the contract event 0xef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf92.
//
// Solidity: event Called(bytes transaction, uint256 gas)
func (_MockTargetContract *MockTargetContractFilterer) ParseCalled(log types.Log) (*MockTargetContractCalled, error) {
	event := new(MockTargetContractCalled)
	if err := _MockTargetContract.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractABI is the input ABI used to generate the binding from.
const TestDepositTokenContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestDepositTokenContractBin is the compiled bytecode used for deploying new contracts.
var TestDepositTokenContractBin = "0x60806040523480156200001157600080fd5b5060408051808201825260038082526214d11560ea1b602080840182815285518087018752938452838201929092528451600081529081019094528251929391926200006091600291620005be565b50815162000076906003906020850190620005be565b5080516200008c9060049060208401906200064d565b5060005b81518110156200010857600160056000848481518110620000c157634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580620000ff8162000829565b91505062000090565b506040516329965a1d60e01b815230600482018190527fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce217705460248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad24906329965a1d90606401600060405180830381600087803b1580156200018357600080fd5b505af115801562000198573d6000803e3d6000fd5b50506040516329965a1d60e01b815230600482018190527faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a60248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad2492506329965a1d9150606401600060405180830381600087803b1580156200021657600080fd5b505af11580156200022b573d6000803e3d6000fd5b505050505050506200026733620f424060405180602001604052806000815250604051806020016040528060008152506200026d60201b60201c565b6200085d565b6001600160a01b038416620002c95760405162461bcd60e51b815260206004820181905260248201527f4552433737373a206d696e7420746f20746865207a65726f206164647265737360448201526064015b60405180910390fd5b60003390508360016000828254620002e29190620007d1565b90915550506001600160a01b0385166000908152602081905260408120805486929062000311908490620007d1565b909155506200032a9050816000878787876001620003c3565b846001600160a01b0316816001600160a01b03167f2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d868686604051620003739392919062000798565b60405180910390a36040518481526001600160a01b038616906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b6024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b1580156200044057600080fd5b505afa15801562000455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200047b9190620006bc565b90506001600160a01b03811615620004fd576040516223de2960e01b81526001600160a01b038216906223de2990620004c3908b908b908b908b908b908b906004016200073a565b600060405180830381600087803b158015620004de57600080fd5b505af1158015620004f3573d6000803e3d6000fd5b50505050620005ae565b8115620005ae5762000523866001600160a01b0316620005b860201b620009d01760201c565b15620005ae5760405162461bcd60e51b815260206004820152604d60248201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460448201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60648201526c1ad95b9cd49958da5c1a595b9d609a1b608482015260a401620002c0565b5050505050505050565b3b151590565b828054620005cc90620007ec565b90600052602060002090601f016020900481019282620005f057600085556200063b565b82601f106200060b57805160ff19168380011785556200063b565b828001600101855582156200063b579182015b828111156200063b5782518255916020019190600101906200061e565b5062000649929150620006a5565b5090565b8280548282559060005260206000209081019282156200063b579160200282015b828111156200063b57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200066e565b5b80821115620006495760008155600101620006a6565b600060208284031215620006ce578081fd5b81516001600160a01b0381168114620006e5578182fd5b9392505050565b60008151808452815b818110156200071357602081850181015186830182015201620006f5565b81811115620007255782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c0608082018190526000906200077790830185620006ec565b82810360a08401526200078b8185620006ec565b9998505050505050505050565b838152606060208201526000620007b36060830185620006ec565b8281036040840152620007c78185620006ec565b9695505050505050565b60008219821115620007e757620007e762000847565b500190565b600181811c908216806200080157607f821691505b602082108114156200082357634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141562000840576200084062000847565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6117d0806200086d6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b63711461022b578063dd62ed3e1461023e578063fad8b32a14610277578063fc673c4f1461028a578063fe9d93031461029d57600080fd5b8063959b8c3f146101ea57806395d89b41146101fd5780639bd9bbc614610205578063a9059cbb1461021857600080fd5b806323b872dd116100e957806323b872dd14610183578063313ce56714610196578063556f0dc7146101a557806362ad1b83146101ac57806370a08231146101c157600080fd5b806306e485381461011b57806306fdde0314610139578063095ea7b31461014e57806318160ddd14610171575b600080fd5b6101236102b0565b60405161013091906115c7565b60405180910390f35b610141610312565b6040516101309190611614565b61016161015c3660046113de565b61039b565b6040519015158152602001610130565b6001545b604051908152602001610130565b61016161019136600461130e565b6103b3565b60405160128152602001610130565b6001610175565b6101bf6101ba36600461134e565b61057c565b005b6101756101cf36600461129e565b6001600160a01b031660009081526020819052604090205490565b6101bf6101f836600461129e565b6105b8565b6101416106d6565b6101bf610213366004611409565b6106e5565b6101616102263660046113de565b610708565b6101616102393660046112d6565b6107bb565b61017561024c3660046112d6565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b6101bf61028536600461129e565b61085d565b6101bf610298366004611460565b610979565b6101bf6102ab3660046114dd565b6109b1565b6060600480548060200260200160405190810160405280929190818152602001828054801561030857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116102ea575b5050505050905090565b6060600280546103219061171b565b80601f016020809104026020016040519081016040528092919081815260200182805461034d9061171b565b80156103085780601f1061036f57610100808354040283529160200191610308565b820191906000526020600020905b81548152906001019060200180831161037d57509395945050505050565b6000336103a98185856109d6565b5060019392505050565b60006001600160a01b0383166103e45760405162461bcd60e51b81526004016103db90611627565b60405180910390fd5b6001600160a01b0384166104495760405162461bcd60e51b815260206004820152602660248201527f4552433737373a207472616e736665722066726f6d20746865207a65726f206160448201526564647265737360d01b60648201526084016103db565b600033905061047a818686866040518060200160405280600081525060405180602001604052806000815250610afd565b6104a6818686866040518060200160405280600081525060405180602001604052806000815250610c34565b6001600160a01b038086166000908152600860209081526040808320938516835292905220548381101561052e5760405162461bcd60e51b815260206004820152602960248201527f4552433737373a207472616e7366657220616d6f756e74206578636565647320604482015268616c6c6f77616e636560b81b60648201526084016103db565b610542868361053d8785611704565b6109d6565b6105708287878760405180602001604052806000815250604051806020016040528060008152506000610da3565b50600195945050505050565b61058633866107bb565b6105a25760405162461bcd60e51b81526004016103db9061166b565b6105b185858585856001610f77565b5050505050565b336001600160a01b038216141561061d5760405162461bcd60e51b8152602060048201526024808201527f4552433737373a20617574686f72697a696e672073656c66206173206f70657260448201526330ba37b960e11b60648201526084016103db565b6001600160a01b03811660009081526005602052604090205460ff161561066e573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff1916905561069d565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191660011790555b60405133906001600160a01b038316907ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f990600090a350565b6060600380546103219061171b565b61070333848484604051806020016040528060008152506001610f77565b505050565b60006001600160a01b0383166107305760405162461bcd60e51b81526004016103db90611627565b6000339050610761818286866040518060200160405280600081525060405180602001604052806000815250610afd565b61078d818286866040518060200160405280600081525060405180602001604052806000815250610c34565b6103a98182868660405180602001604052806000815250604051806020016040528060008152506000610da3565b6000816001600160a01b0316836001600160a01b0316148061082657506001600160a01b03831660009081526005602052604090205460ff16801561082657506001600160a01b0380831660009081526007602090815260408083209387168352929052205460ff16155b8061085657506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff165b9392505050565b6001600160a01b0381163314156108c05760405162461bcd60e51b815260206004820152602160248201527f4552433737373a207265766f6b696e672073656c66206173206f70657261746f6044820152603960f91b60648201526084016103db565b6001600160a01b03811660009081526005602052604090205460ff1615610914573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610940565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191690555b60405133906001600160a01b038316907f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa190600090a350565b61098333856107bb565b61099f5760405162461bcd60e51b81526004016103db9061166b565b6109ab8484848461105a565b50505050565b6109cc3383836040518060200160405280600081525061105a565b5050565b3b151590565b6001600160a01b038316610a3a5760405162461bcd60e51b815260206004820152602560248201527f4552433737373a20617070726f76652066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103db565b6001600160a01b038216610a9c5760405162461bcd60e51b815260206004820152602360248201527f4552433737373a20617070726f766520746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103db565b6001600160a01b0383811660008181526008602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe8956024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610b7957600080fd5b505afa158015610b8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb191906112ba565b90506001600160a01b03811615610c2b57604051633ad5cbc160e11b81526001600160a01b038216906375ab978290610bf8908a908a908a908a908a908a9060040161156d565b600060405180830381600087803b158015610c1257600080fd5b505af1158015610c26573d6000803e3d6000fd5b505050505b50505050505050565b6001600160a01b03851660009081526020819052604090205483811015610cad5760405162461bcd60e51b815260206004820152602760248201527f4552433737373a207472616e7366657220616d6f756e7420657863656564732060448201526662616c616e636560c81b60648201526084016103db565b610cb78482611704565b6001600160a01b038088166000908152602081905260408082209390935590871681529081208054869290610ced9084906116ec565b92505081905550846001600160a01b0316866001600160a01b0316886001600160a01b03167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987878787604051610d45939291906116b7565b60405180910390a4846001600160a01b0316866001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef86604051610d9291815260200190565b60405180910390a350505050505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b6024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610e1f57600080fd5b505afa158015610e33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5791906112ba565b90506001600160a01b03811615610ed3576040516223de2960e01b81526001600160a01b038216906223de2990610e9c908b908b908b908b908b908b9060040161156d565b600060405180830381600087803b158015610eb657600080fd5b505af1158015610eca573d6000803e3d6000fd5b50505050610f6d565b8115610f6d576001600160a01b0386163b15610f6d5760405162461bcd60e51b815260206004820152604d60248201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460448201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60648201526c1ad95b9cd49958da5c1a595b9d609a1b608482015260a4016103db565b5050505050505050565b6001600160a01b038616610fd85760405162461bcd60e51b815260206004820152602260248201527f4552433737373a2073656e642066726f6d20746865207a65726f206164647265604482015261737360f01b60648201526084016103db565b6001600160a01b03851661102e5760405162461bcd60e51b815260206004820181905260248201527f4552433737373a2073656e6420746f20746865207a65726f206164647265737360448201526064016103db565b3361103d818888888888610afd565b61104b818888888888610c34565b610c2b81888888888888610da3565b6001600160a01b0384166110bb5760405162461bcd60e51b815260206004820152602260248201527f4552433737373a206275726e2066726f6d20746865207a65726f206164647265604482015261737360f01b60648201526084016103db565b336110cb81866000878787610afd565b6001600160a01b038516600090815260208190526040902054848110156111405760405162461bcd60e51b815260206004820152602360248201527f4552433737373a206275726e20616d6f756e7420657863656564732062616c616044820152626e636560e81b60648201526084016103db565b61114a8582611704565b6001600160a01b03871660009081526020819052604081209190915560018054879290611178908490611704565b92505081905550856001600160a01b0316826001600160a01b03167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a40988787876040516111c6939291906116b7565b60405180910390a36040518581526000906001600160a01b038816907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050505050565b600082601f830112611227578081fd5b813567ffffffffffffffff808211156112425761124261176c565b604051601f8301601f19908116603f0116810190828211818310171561126a5761126a61176c565b81604052838152866020858801011115611282578485fd5b8360208701602083013792830160200193909352509392505050565b6000602082840312156112af578081fd5b813561085681611782565b6000602082840312156112cb578081fd5b815161085681611782565b600080604083850312156112e8578081fd5b82356112f381611782565b9150602083013561130381611782565b809150509250929050565b600080600060608486031215611322578081fd5b833561132d81611782565b9250602084013561133d81611782565b929592945050506040919091013590565b600080600080600060a08688031215611365578081fd5b853561137081611782565b9450602086013561138081611782565b935060408601359250606086013567ffffffffffffffff808211156113a3578283fd5b6113af89838a01611217565b935060808801359150808211156113c4578283fd5b506113d188828901611217565b9150509295509295909350565b600080604083850312156113f0578182fd5b82356113fb81611782565b946020939093013593505050565b60008060006060848603121561141d578283fd5b833561142881611782565b925060208401359150604084013567ffffffffffffffff81111561144a578182fd5b61145686828701611217565b9150509250925092565b60008060008060808587031215611475578384fd5b843561148081611782565b935060208501359250604085013567ffffffffffffffff808211156114a3578384fd5b6114af88838901611217565b935060608701359150808211156114c4578283fd5b506114d187828801611217565b91505092959194509250565b600080604083850312156114ef578182fd5b82359150602083013567ffffffffffffffff81111561150c578182fd5b61151885828601611217565b9150509250929050565b60008151808452815b818110156115475760208185018101518683018201520161152b565b818111156115585782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c0608082018190526000906115a890830185611522565b82810360a08401526115ba8185611522565b9998505050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156116085783516001600160a01b0316835292840192918401916001016115e3565b50909695505050505050565b6020815260006108566020830184611522565b60208082526024908201527f4552433737373a207472616e7366657220746f20746865207a65726f206164646040820152637265737360e01b606082015260800190565b6020808252602c908201527f4552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f60408201526b39103337b9103437b63232b960a11b606082015260800190565b8381526060602082015260006116d06060830185611522565b82810360408401526116e28185611522565b9695505050505050565b600082198211156116ff576116ff611756565b500190565b60008282101561171657611716611756565b500390565b600181811c9082168061172f57607f821691505b6020821081141561175057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461179757600080fd5b5056fea26469706673582212208b04f51b3ad66f1b39f3640d6975ae56227c7e47f3ed5c706ec4a1225300403c64736f6c63430008040033"

// DeployTestDepositTokenContract deploys a new Ethereum contract, binding an instance of TestDepositTokenContract to it.
func DeployTestDepositTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestDepositTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestDepositTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestDepositTokenContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestDepositTokenContract{TestDepositTokenContractCaller: TestDepositTokenContractCaller{contract: contract}, TestDepositTokenContractTransactor: TestDepositTokenContractTransactor{contract: contract}, TestDepositTokenContractFilterer: TestDepositTokenContractFilterer{contract: contract}}, nil
}

// TestDepositTokenContract is an auto generated Go binding around an Ethereum contract.
type TestDepositTokenContract struct {
	TestDepositTokenContractCaller     // Read-only binding to the contract
	TestDepositTokenContractTransactor // Write-only binding to the contract
	TestDepositTokenContractFilterer   // Log filterer for contract events
}

// TestDepositTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestDepositTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDepositTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestDepositTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDepositTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestDepositTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDepositTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestDepositTokenContractSession struct {
	Contract     *TestDepositTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TestDepositTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestDepositTokenContractCallerSession struct {
	Contract *TestDepositTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TestDepositTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestDepositTokenContractTransactorSession struct {
	Contract     *TestDepositTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TestDepositTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestDepositTokenContractRaw struct {
	Contract *TestDepositTokenContract // Generic contract binding to access the raw methods on
}

// TestDepositTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestDepositTokenContractCallerRaw struct {
	Contract *TestDepositTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestDepositTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestDepositTokenContractTransactorRaw struct {
	Contract *TestDepositTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestDepositTokenContract creates a new instance of TestDepositTokenContract, bound to a specific deployed contract.
func NewTestDepositTokenContract(address common.Address, backend bind.ContractBackend) (*TestDepositTokenContract, error) {
	contract, err := bindTestDepositTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContract{TestDepositTokenContractCaller: TestDepositTokenContractCaller{contract: contract}, TestDepositTokenContractTransactor: TestDepositTokenContractTransactor{contract: contract}, TestDepositTokenContractFilterer: TestDepositTokenContractFilterer{contract: contract}}, nil
}

// NewTestDepositTokenContractCaller creates a new read-only instance of TestDepositTokenContract, bound to a specific deployed contract.
func NewTestDepositTokenContractCaller(address common.Address, caller bind.ContractCaller) (*TestDepositTokenContractCaller, error) {
	contract, err := bindTestDepositTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractCaller{contract: contract}, nil
}

// NewTestDepositTokenContractTransactor creates a new write-only instance of TestDepositTokenContract, bound to a specific deployed contract.
func NewTestDepositTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestDepositTokenContractTransactor, error) {
	contract, err := bindTestDepositTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractTransactor{contract: contract}, nil
}

// NewTestDepositTokenContractFilterer creates a new log filterer instance of TestDepositTokenContract, bound to a specific deployed contract.
func NewTestDepositTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestDepositTokenContractFilterer, error) {
	contract, err := bindTestDepositTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractFilterer{contract: contract}, nil
}

// bindTestDepositTokenContract binds a generic wrapper to an already deployed contract.
func bindTestDepositTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestDepositTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDepositTokenContract *TestDepositTokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDepositTokenContract.Contract.TestDepositTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDepositTokenContract *TestDepositTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.TestDepositTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDepositTokenContract *TestDepositTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.TestDepositTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDepositTokenContract *TestDepositTokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDepositTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDepositTokenContract *TestDepositTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDepositTokenContract *TestDepositTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _TestDepositTokenContract.Contract.Allowance(&_TestDepositTokenContract.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _TestDepositTokenContract.Contract.Allowance(&_TestDepositTokenContract.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _TestDepositTokenContract.Contract.BalanceOf(&_TestDepositTokenContract.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _TestDepositTokenContract.Contract.BalanceOf(&_TestDepositTokenContract.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Decimals() (uint8, error) {
	return _TestDepositTokenContract.Contract.Decimals(&_TestDepositTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) Decimals() (uint8, error) {
	return _TestDepositTokenContract.Contract.Decimals(&_TestDepositTokenContract.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_TestDepositTokenContract *TestDepositTokenContractCaller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_TestDepositTokenContract *TestDepositTokenContractSession) DefaultOperators() ([]common.Address, error) {
	return _TestDepositTokenContract.Contract.DefaultOperators(&_TestDepositTokenContract.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) DefaultOperators() ([]common.Address, error) {
	return _TestDepositTokenContract.Contract.DefaultOperators(&_TestDepositTokenContract.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Granularity() (*big.Int, error) {
	return _TestDepositTokenContract.Contract.Granularity(&_TestDepositTokenContract.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) Granularity() (*big.Int, error) {
	return _TestDepositTokenContract.Contract.Granularity(&_TestDepositTokenContract.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _TestDepositTokenContract.Contract.IsOperatorFor(&_TestDepositTokenContract.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _TestDepositTokenContract.Contract.IsOperatorFor(&_TestDepositTokenContract.CallOpts, operator, tokenHolder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Name() (string, error) {
	return _TestDepositTokenContract.Contract.Name(&_TestDepositTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) Name() (string, error) {
	return _TestDepositTokenContract.Contract.Name(&_TestDepositTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Symbol() (string, error) {
	return _TestDepositTokenContract.Contract.Symbol(&_TestDepositTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) Symbol() (string, error) {
	return _TestDepositTokenContract.Contract.Symbol(&_TestDepositTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestDepositTokenContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractSession) TotalSupply() (*big.Int, error) {
	return _TestDepositTokenContract.Contract.TotalSupply(&_TestDepositTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestDepositTokenContract *TestDepositTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _TestDepositTokenContract.Contract.TotalSupply(&_TestDepositTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Approve(&_TestDepositTokenContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Approve(&_TestDepositTokenContract.TransactOpts, spender, value)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.AuthorizeOperator(&_TestDepositTokenContract.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.AuthorizeOperator(&_TestDepositTokenContract.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) Burn(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "burn", amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Burn(&_TestDepositTokenContract.TransactOpts, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Burn(&_TestDepositTokenContract.TransactOpts, amount, data)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.OperatorBurn(&_TestDepositTokenContract.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.OperatorBurn(&_TestDepositTokenContract.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.OperatorSend(&_TestDepositTokenContract.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.OperatorSend(&_TestDepositTokenContract.TransactOpts, sender, recipient, amount, data, operatorData)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.RevokeOperator(&_TestDepositTokenContract.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.RevokeOperator(&_TestDepositTokenContract.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Send(&_TestDepositTokenContract.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Send(&_TestDepositTokenContract.TransactOpts, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Transfer(&_TestDepositTokenContract.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.Transfer(&_TestDepositTokenContract.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.TransferFrom(&_TestDepositTokenContract.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_TestDepositTokenContract *TestDepositTokenContractTransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestDepositTokenContract.Contract.TransferFrom(&_TestDepositTokenContract.TransactOpts, holder, recipient, amount)
}

// TestDepositTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractApprovalIterator struct {
	Event *TestDepositTokenContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractApproval represents a Approval event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestDepositTokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractApprovalIterator{contract: _TestDepositTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractApproval)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseApproval(log types.Log) (*TestDepositTokenContractApproval, error) {
	event := new(TestDepositTokenContractApproval)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractAuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractAuthorizedOperatorIterator struct {
	Event *TestDepositTokenContractAuthorizedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractAuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractAuthorizedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractAuthorizedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractAuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractAuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractAuthorizedOperator represents a AuthorizedOperator event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractAuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*TestDepositTokenContractAuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractAuthorizedOperatorIterator{contract: _TestDepositTokenContract.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractAuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractAuthorizedOperator)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseAuthorizedOperator(log types.Log) (*TestDepositTokenContractAuthorizedOperator, error) {
	event := new(TestDepositTokenContractAuthorizedOperator)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractBurnedIterator struct {
	Event *TestDepositTokenContractBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractBurned represents a Burned event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractBurned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*TestDepositTokenContractBurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractBurnedIterator{contract: _TestDepositTokenContract.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractBurned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractBurned)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseBurned(log types.Log) (*TestDepositTokenContractBurned, error) {
	event := new(TestDepositTokenContractBurned)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractMintedIterator struct {
	Event *TestDepositTokenContractMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractMinted represents a Minted event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractMinted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*TestDepositTokenContractMintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractMintedIterator{contract: _TestDepositTokenContract.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractMinted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractMinted)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseMinted(log types.Log) (*TestDepositTokenContractMinted, error) {
	event := new(TestDepositTokenContractMinted)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractRevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractRevokedOperatorIterator struct {
	Event *TestDepositTokenContractRevokedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractRevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractRevokedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractRevokedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractRevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractRevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractRevokedOperator represents a RevokedOperator event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractRevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*TestDepositTokenContractRevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractRevokedOperatorIterator{contract: _TestDepositTokenContract.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractRevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractRevokedOperator)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseRevokedOperator(log types.Log) (*TestDepositTokenContractRevokedOperator, error) {
	event := new(TestDepositTokenContractRevokedOperator)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractSentIterator struct {
	Event *TestDepositTokenContractSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractSent represents a Sent event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractSent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TestDepositTokenContractSentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractSentIterator{contract: _TestDepositTokenContract.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractSent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractSent)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseSent(log types.Log) (*TestDepositTokenContractSent, error) {
	event := new(TestDepositTokenContractSent)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestDepositTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestDepositTokenContract contract.
type TestDepositTokenContractTransferIterator struct {
	Event *TestDepositTokenContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestDepositTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDepositTokenContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestDepositTokenContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestDepositTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDepositTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDepositTokenContractTransfer represents a Transfer event raised by the TestDepositTokenContract contract.
type TestDepositTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestDepositTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestDepositTokenContractTransferIterator{contract: _TestDepositTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestDepositTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestDepositTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDepositTokenContractTransfer)
				if err := _TestDepositTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestDepositTokenContract *TestDepositTokenContractFilterer) ParseTransfer(log types.Log) (*TestDepositTokenContractTransfer, error) {
	event := new(TestDepositTokenContractTransfer)
	if err := _TestDepositTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTargetContractABI is the input ABI used to generate the binding from.
const TestTargetContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"isNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestTargetContractBin is the compiled bytecode used for deploying new contracts.
var TestTargetContractBin = "0x608060405234801561001057600080fd5b506040516106ee3803806106ee83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b61065d806100916000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633a8a8beb14610046578063943d7209146100a0578063c34c08e5146100b5575b600080fd5b61008b6100543660046103b2565b6001600160a01b038216600090815260016020908152604080832067ffffffffffffffff8516845290915290205460ff1692915050565b60405190151581526020015b60405180910390f35b6100b36100ae3660046103f5565b6100e0565b005b6000546100c8906001600160a01b031681565b6040516001600160a01b039091168152602001610097565b6000546001600160a01b031633146101555760405162461bcd60e51b815260206004820152602d60248201527f54657374546172676574436f6e74726163743a206f6e6c79206578656375746f60448201526c722063616e206578656375746560981b60648201526084015b60405180910390fd5b6000806000808480602001905181019061016f9190610469565b93509350935093506000848051906020012090506000816040516020016101c291907f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b60408051601f198184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa15801561022d573d6000803e3d6000fd5b5050506020604051035190506000808880602001905181019061025091906104cc565b6001600160a01b038516600090815260016020908152604080832067ffffffffffffffff86168452909152902054919350915060ff16156102e25760405162461bcd60e51b815260206004820152602660248201527f54657374546172676574436f6e74726163743a206e6f6e636520616c726561646044820152651e481d5cd95960d21b606482015260840161014c565b6001600160a01b038316600090815260016020818152604080842067ffffffffffffffff8716855290915291829020805460ff19169091179055517f5932d06f5fe39d60725112950a65a5591f5cb869ce6bc5d90fd188fb255491d49061034e9085908490869061051b565b60405180910390a150505050505050505050565b600082601f830112610372578081fd5b8151610385610380826105a0565b61056f565b818152846020838601011115610399578283fd5b6103aa8260208301602087016105c8565b949350505050565b600080604083850312156103c4578182fd5b82356001600160a01b03811681146103da578283fd5b915060208301356103ea8161060e565b809150509250929050565b600060208284031215610406578081fd5b813567ffffffffffffffff81111561041c578182fd5b8201601f8101841361042c578182fd5b803561043a610380826105a0565b81815285602083850101111561044e578384fd5b81602084016020830137908101602001929092525092915050565b6000806000806080858703121561047e578182fd5b845167ffffffffffffffff811115610494578283fd5b6104a087828801610362565b945050602085015160ff811681146104b6578283fd5b6040860151606090960151949790965092505050565b600080604083850312156104de578182fd5b82516104e98161060e565b602084015190925067ffffffffffffffff811115610505578182fd5b61051185828601610362565b9150509250929050565b60018060a01b038416815260606020820152600083518060608401526105488160808501602088016105c8565b67ffffffffffffffff93909316604083015250601f91909101601f19160160800192915050565b604051601f8201601f1916810167ffffffffffffffff81118282101715610598576105986105f8565b604052919050565b600067ffffffffffffffff8211156105ba576105ba6105f8565b50601f01601f191660200190565b60005b838110156105e35781810151838201526020016105cb565b838111156105f2576000848401525b50505050565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff8116811461062457600080fd5b5056fea2646970667358221220f7c73698c640099ca005cb284d40fd4dbecb59d4df10d4af53195f586e7c45e964736f6c63430008040033"

// DeployTestTargetContract deploys a new Ethereum contract, binding an instance of TestTargetContract to it.
func DeployTestTargetContract(auth *bind.TransactOpts, backend bind.ContractBackend, executorAddress common.Address) (common.Address, *types.Transaction, *TestTargetContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTargetContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestTargetContractBin), backend, executorAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestTargetContract{TestTargetContractCaller: TestTargetContractCaller{contract: contract}, TestTargetContractTransactor: TestTargetContractTransactor{contract: contract}, TestTargetContractFilterer: TestTargetContractFilterer{contract: contract}}, nil
}

// TestTargetContract is an auto generated Go binding around an Ethereum contract.
type TestTargetContract struct {
	TestTargetContractCaller     // Read-only binding to the contract
	TestTargetContractTransactor // Write-only binding to the contract
	TestTargetContractFilterer   // Log filterer for contract events
}

// TestTargetContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTargetContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTargetContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTargetContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTargetContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTargetContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTargetContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTargetContractSession struct {
	Contract     *TestTargetContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestTargetContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTargetContractCallerSession struct {
	Contract *TestTargetContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TestTargetContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTargetContractTransactorSession struct {
	Contract     *TestTargetContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TestTargetContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTargetContractRaw struct {
	Contract *TestTargetContract // Generic contract binding to access the raw methods on
}

// TestTargetContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTargetContractCallerRaw struct {
	Contract *TestTargetContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestTargetContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTargetContractTransactorRaw struct {
	Contract *TestTargetContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestTargetContract creates a new instance of TestTargetContract, bound to a specific deployed contract.
func NewTestTargetContract(address common.Address, backend bind.ContractBackend) (*TestTargetContract, error) {
	contract, err := bindTestTargetContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestTargetContract{TestTargetContractCaller: TestTargetContractCaller{contract: contract}, TestTargetContractTransactor: TestTargetContractTransactor{contract: contract}, TestTargetContractFilterer: TestTargetContractFilterer{contract: contract}}, nil
}

// NewTestTargetContractCaller creates a new read-only instance of TestTargetContract, bound to a specific deployed contract.
func NewTestTargetContractCaller(address common.Address, caller bind.ContractCaller) (*TestTargetContractCaller, error) {
	contract, err := bindTestTargetContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTargetContractCaller{contract: contract}, nil
}

// NewTestTargetContractTransactor creates a new write-only instance of TestTargetContract, bound to a specific deployed contract.
func NewTestTargetContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTargetContractTransactor, error) {
	contract, err := bindTestTargetContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTargetContractTransactor{contract: contract}, nil
}

// NewTestTargetContractFilterer creates a new log filterer instance of TestTargetContract, bound to a specific deployed contract.
func NewTestTargetContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTargetContractFilterer, error) {
	contract, err := bindTestTargetContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTargetContractFilterer{contract: contract}, nil
}

// bindTestTargetContract binds a generic wrapper to an already deployed contract.
func bindTestTargetContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTargetContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTargetContract *TestTargetContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTargetContract.Contract.TestTargetContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTargetContract *TestTargetContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTargetContract.Contract.TestTargetContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTargetContract *TestTargetContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTargetContract.Contract.TestTargetContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTargetContract *TestTargetContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTargetContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTargetContract *TestTargetContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTargetContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTargetContract *TestTargetContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTargetContract.Contract.contract.Transact(opts, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_TestTargetContract *TestTargetContractCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestTargetContract.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_TestTargetContract *TestTargetContractSession) Executor() (common.Address, error) {
	return _TestTargetContract.Contract.Executor(&_TestTargetContract.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_TestTargetContract *TestTargetContractCallerSession) Executor() (common.Address, error) {
	return _TestTargetContract.Contract.Executor(&_TestTargetContract.CallOpts)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3a8a8beb.
//
// Solidity: function isNonceUsed(address account, uint64 nonce) view returns(bool)
func (_TestTargetContract *TestTargetContractCaller) IsNonceUsed(opts *bind.CallOpts, account common.Address, nonce uint64) (bool, error) {
	var out []interface{}
	err := _TestTargetContract.contract.Call(opts, &out, "isNonceUsed", account, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3a8a8beb.
//
// Solidity: function isNonceUsed(address account, uint64 nonce) view returns(bool)
func (_TestTargetContract *TestTargetContractSession) IsNonceUsed(account common.Address, nonce uint64) (bool, error) {
	return _TestTargetContract.Contract.IsNonceUsed(&_TestTargetContract.CallOpts, account, nonce)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3a8a8beb.
//
// Solidity: function isNonceUsed(address account, uint64 nonce) view returns(bool)
func (_TestTargetContract *TestTargetContractCallerSession) IsNonceUsed(account common.Address, nonce uint64) (bool, error) {
	return _TestTargetContract.Contract.IsNonceUsed(&_TestTargetContract.CallOpts, account, nonce)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x943d7209.
//
// Solidity: function executeTransaction(bytes txData) returns()
func (_TestTargetContract *TestTargetContractTransactor) ExecuteTransaction(opts *bind.TransactOpts, txData []byte) (*types.Transaction, error) {
	return _TestTargetContract.contract.Transact(opts, "executeTransaction", txData)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x943d7209.
//
// Solidity: function executeTransaction(bytes txData) returns()
func (_TestTargetContract *TestTargetContractSession) ExecuteTransaction(txData []byte) (*types.Transaction, error) {
	return _TestTargetContract.Contract.ExecuteTransaction(&_TestTargetContract.TransactOpts, txData)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x943d7209.
//
// Solidity: function executeTransaction(bytes txData) returns()
func (_TestTargetContract *TestTargetContractTransactorSession) ExecuteTransaction(txData []byte) (*types.Transaction, error) {
	return _TestTargetContract.Contract.ExecuteTransaction(&_TestTargetContract.TransactOpts, txData)
}

// TestTargetContractExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the TestTargetContract contract.
type TestTargetContractExecutedTransactionIterator struct {
	Event *TestTargetContractExecutedTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestTargetContractExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTargetContractExecutedTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestTargetContractExecutedTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestTargetContractExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTargetContractExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTargetContractExecutedTransaction represents a ExecutedTransaction event raised by the TestTargetContract contract.
type TestTargetContractExecutedTransaction struct {
	Sender common.Address
	Data   []byte
	Nonce  uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0x5932d06f5fe39d60725112950a65a5591f5cb869ce6bc5d90fd188fb255491d4.
//
// Solidity: event ExecutedTransaction(address sender, bytes data, uint64 nonce)
func (_TestTargetContract *TestTargetContractFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*TestTargetContractExecutedTransactionIterator, error) {

	logs, sub, err := _TestTargetContract.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &TestTargetContractExecutedTransactionIterator{contract: _TestTargetContract.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0x5932d06f5fe39d60725112950a65a5591f5cb869ce6bc5d90fd188fb255491d4.
//
// Solidity: event ExecutedTransaction(address sender, bytes data, uint64 nonce)
func (_TestTargetContract *TestTargetContractFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *TestTargetContractExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _TestTargetContract.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTargetContractExecutedTransaction)
				if err := _TestTargetContract.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutedTransaction is a log parse operation binding the contract event 0x5932d06f5fe39d60725112950a65a5591f5cb869ce6bc5d90fd188fb255491d4.
//
// Solidity: event ExecutedTransaction(address sender, bytes data, uint64 nonce)
func (_TestTargetContract *TestTargetContractFilterer) ParseExecutedTransaction(log types.Log) (*TestTargetContractExecutedTransaction, error) {
	event := new(TestTargetContractExecutedTransaction)
	if err := _TestTargetContract.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
