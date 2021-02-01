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
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122048b504f555ec6cb71f2b66830f4b4b88146f0b934c1eb112f38cf23d45d3f06164736f6c63430007010033"

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
const BatcherContractABI = "[{\"inputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"contractFeeBankContract\",\"name\":\"_feeBankContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTransactionType\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_transaction\",\"type\":\"bytes\"}],\"name\":\"addTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchSizes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBankContract\",\"outputs\":[{\"internalType\":\"contractFeeBankContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_minFee\",\"type\":\"uint64\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BatcherContractBin is the compiled bytecode used for deploying new contracts.
var BatcherContractBin = "0x608060405234801561001057600080fd5b50604051610f9b380380610f9b83398101604081905261002f916100b8565b60006100396100b4565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b039384166001600160a01b03199182161790915560028054929093169116179055610109565b3390565b600080604083850312156100ca578182fd5b82516100d5816100f1565b60208401519092506100e6816100f1565b809150509250929050565b6001600160a01b038116811461010657600080fd5b50565b610e83806101186000396000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b1461012d578063bf66a18214610142578063bfd260ca14610157578063c87afa8a14610177578063f2fde38b146101a457610091565b8063246673dc1461009657806324ec7590146100ab57806336e1290d146100d657806348fd5acc146100f8578063715018a614610118575b600080fd5b6100a96100a4366004610a87565b6101c4565b005b3480156100b757600080fd5b506100c061059c565b6040516100cd9190610d8d565b60405180910390f35b3480156100e257600080fd5b506100eb6105ab565b6040516100cd9190610b27565b34801561010457600080fd5b506100a9610113366004610a36565b6105ba565b34801561012457600080fd5b506100a9610612565b34801561013957600080fd5b506100eb610691565b34801561014e57600080fd5b506100eb6106a0565b34801561016357600080fd5b506100c0610172366004610a36565b6106af565b34801561018357600080fd5b50610197610192366004610a52565b6106ca565b6040516100cd9190610b3b565b3480156101b057600080fd5b506100a96101bf3660046108d1565b6106e7565b6101cc6107a1565b60015460405163700465b160e11b81526001600160a01b039091169063e008cb62906101fc908890600401610d8d565b60006040518083038186803b15801561021457600080fd5b505afa158015610228573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261025091908101906108f4565b9050600081608001516001600160401b0316116102885760405162461bcd60e51b815260040161027f90610bcf565b60405180910390fd5b80600001516001600160401b0316856001600160401b031610156102a857fe5b80516080820151602083015191870391818302019081016001600160401b0382164310156102e85760405162461bcd60e51b815260040161027f90610d03565b806001600160401b031643106103105760405162461bcd60e51b815260040161027f90610d49565b8461032d5760405162461bcd60e51b815260040161027f90610b8a565b60c08401516001600160401b031685111561035a5760405162461bcd60e51b815260040161027f90610c10565b60a08401516001600160401b0389811660009081526003602052604090205491811691168601111561039e5760405162461bcd60e51b815260040161027f90610cc0565b6005546001600160401b03163410156103c95760405162461bcd60e51b815260040161027f90610c89565b6001600160401b0388166000908152600460205260408120606091889188918b60018111156103f457fe5b60018111156103ff57fe5b81526020019081526020016000205460405160200161042093929190610b12565b604051602081830303815290604052905060008180519060200120905080600460008c6001600160401b03166001600160401b0316815260200190815260200160002060008b600181111561047157fe5b600181111561047c57fe5b815260208082019290925260409081016000908120939093556001600160401b038d811684526003909252909120805467ffffffffffffffff1981169083168a0190921691909117905534158015906104e257506101008601516001600160a01b031615155b156105515760025461010087015160405163f340fa0160e01b81526001600160a01b039092169163f340fa0191349161051e9190600401610b27565b6000604051808303818588803b15801561053757600080fd5b505af115801561054b573d6000803e3d6000fd5b50505050505b7ffc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c988a8a8a8a85604051610588959493929190610da1565b60405180910390a150505050505050505050565b6005546001600160401b031681565b6002546001600160a01b031681565b6105c261079d565b6000546001600160a01b039081169116146105ef5760405162461bcd60e51b815260040161027f90610c54565b6005805467ffffffffffffffff19166001600160401b0392909216919091179055565b61061a61079d565b6000546001600160a01b039081169116146106475760405162461bcd60e51b815260040161027f90610c54565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6001546001600160a01b031681565b6003602052600090815260409020546001600160401b031681565b600460209081526000928352604080842090915290825290205481565b6106ef61079d565b6000546001600160a01b0390811691161461071c5760405162461bcd60e51b815260040161027f90610c54565b6001600160a01b0381166107425760405162461bcd60e51b815260040161027f90610b44565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b805161081181610e20565b92915050565b600082601f830112610827578081fd5b81516001600160401b0381111561083c578182fd5b602080820261084c828201610dfa565b8381529350818401858301828701840188101561086857600080fd5b600092505b8483101561089457805161088081610e20565b82526001929092019190830190830161086d565b505050505092915050565b80516001600160e01b03198116811461081157600080fd5b80356002811061081157600080fd5b805161081181610e38565b6000602082840312156108e2578081fd5b81356108ed81610e20565b9392505050565b600060208284031215610905578081fd5b81516001600160401b038082111561091b578283fd5b8184019150610180808387031215610931578384fd5b61093a81610dfa565b905061094686846108c6565b815261095586602085016108c6565b602082015260408301518281111561096b578485fd5b61097787828601610817565b60408301525061098a86606085016108c6565b606082015261099c86608085016108c6565b60808201526109ae8660a085016108c6565b60a08201526109c08660c085016108c6565b60c08201526109d28660e085016108c6565b60e082015261010091506109e886838501610806565b8282015261012091506109fd86838501610806565b828201526101409150610a128683850161089f565b828201526101609150610a27868385016108c6565b91810191909152949350505050565b600060208284031215610a47578081fd5b81356108ed81610e38565b60008060408385031215610a64578081fd5b8235610a6f81610e38565b9150610a7e84602085016108b7565b90509250929050565b60008060008060608587031215610a9c578182fd5b8435610aa781610e38565b9350610ab686602087016108b7565b925060408501356001600160401b0380821115610ad1578384fd5b818701915087601f830112610ae4578384fd5b813581811115610af2578485fd5b886020828501011115610b03578485fd5b95989497505060200194505050565b60008385833750909101908152602001919050565b6001600160a01b0391909116815260200190565b90815260200190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b60208082526025908201527f42617463686572436f6e74726163743a207472616e73616374696f6e20697320604082015264656d70747960d81b606082015260800190565b60208082526021908201527f42617463686572436f6e74726163743a206261746368206e6f742061637469766040820152606560f81b606082015260800190565b60208082526024908201527f42617463686572436f6e74726163743a207472616e73616374696f6e20746f6f6040820152632062696760e01b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601e908201527f42617463686572436f6e74726163743a2066656520746f6f20736d616c6c0000604082015260600190565b60208082526023908201527f42617463686572436f6e74726163743a20626174636820616c726561647920666040820152621d5b1b60ea1b606082015260800190565b60208082526026908201527f42617463686572436f6e74726163743a206261746368206e6f742073746172746040820152651959081e595d60d21b606082015260800190565b60208082526024908201527f42617463686572436f6e74726163743a20626174636820616c726561647920656040820152631b99195960e21b606082015260800190565b6001600160401b0391909116815260200190565b60006001600160401b038716825260028610610db957fe5b85602083015260806040830152836080830152838560a08401378060a0858401015260a0601f19601f86011683010190508260608301529695505050505050565b6040518181016001600160401b0381118282101715610e1857600080fd5b604052919050565b6001600160a01b0381168114610e3557600080fd5b50565b6001600160401b0381168114610e3557600080fdfea264697066735822122039fe239ef0b764cc7da3f2c6f790cd075088452fc7bff555275609da04f4695364736f6c63430007010033"

// DeployBatcherContract deploys a new Ethereum contract, binding an instance of BatcherContract to it.
func DeployBatcherContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configContract common.Address, _feeBankContract common.Address) (common.Address, *types.Transaction, *BatcherContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatcherContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatcherContractBin), backend, _configContract, _feeBankContract)
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
// Solidity: function addTransaction(uint64 _batchIndex, uint8 _type, bytes _transaction) payable returns()
func (_BatcherContract *BatcherContractTransactor) AddTransaction(opts *bind.TransactOpts, _batchIndex uint64, _type uint8, _transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "addTransaction", _batchIndex, _type, _transaction)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x246673dc.
//
// Solidity: function addTransaction(uint64 _batchIndex, uint8 _type, bytes _transaction) payable returns()
func (_BatcherContract *BatcherContractSession) AddTransaction(_batchIndex uint64, _type uint8, _transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.Contract.AddTransaction(&_BatcherContract.TransactOpts, _batchIndex, _type, _transaction)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x246673dc.
//
// Solidity: function addTransaction(uint64 _batchIndex, uint8 _type, bytes _transaction) payable returns()
func (_BatcherContract *BatcherContractTransactorSession) AddTransaction(_batchIndex uint64, _type uint8, _transaction []byte) (*types.Transaction, error) {
	return _BatcherContract.Contract.AddTransaction(&_BatcherContract.TransactOpts, _batchIndex, _type, _transaction)
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
// Solidity: function setMinFee(uint64 _minFee) returns()
func (_BatcherContract *BatcherContractTransactor) SetMinFee(opts *bind.TransactOpts, _minFee uint64) (*types.Transaction, error) {
	return _BatcherContract.contract.Transact(opts, "setMinFee", _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x48fd5acc.
//
// Solidity: function setMinFee(uint64 _minFee) returns()
func (_BatcherContract *BatcherContractSession) SetMinFee(_minFee uint64) (*types.Transaction, error) {
	return _BatcherContract.Contract.SetMinFee(&_BatcherContract.TransactOpts, _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x48fd5acc.
//
// Solidity: function setMinFee(uint64 _minFee) returns()
func (_BatcherContract *BatcherContractTransactorSession) SetMinFee(_minFee uint64) (*types.Transaction, error) {
	return _BatcherContract.Contract.SetMinFee(&_BatcherContract.TransactOpts, _minFee)
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
const ConfigContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configChangeHeadsUpBlocks\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigUnscheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"configChangeHeadsUpBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"}],\"name\":\"configKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configIndex\",\"type\":\"uint64\"}],\"name\":\"configNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"keypers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structBatchConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newKeypers\",\"type\":\"address[]\"}],\"name\":\"nextConfigAddKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"nextConfigKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nextConfigRemoveKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSpan\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSpan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_executionTimeout\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetExecutionTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"nextConfigSetFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBatchIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startBlockNumber\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_targetAddress\",\"type\":\"address\"}],\"name\":\"nextConfigSetTargetAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_targetFunctionSelector\",\"type\":\"bytes4\"}],\"name\":\"nextConfigSetTargetFunctionSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_threshold\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_transactionGasLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_transactionSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfigs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduleNextConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_fromStartBlockNumber\",\"type\":\"uint64\"}],\"name\":\"unscheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigContractBin is the compiled bytecode used for deploying new contracts.
var ConfigContractBin = "0x60a06040523480156200001157600080fd5b5060405162002300380380620023008339810160408190526200003491620003b8565b60006200004062000249565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001620000966200024d565b815460018181018455600093845260209384902083516005909302018054858501516001600160401b039081166801000000000000000002600160401b600160801b0319919095166001600160401b0319909216919091171692909217825560408301518051939492936200011493928501929190910190620002c8565b50606082015160028201805460808086015160a087015160c0808901516001600160401b03199586166001600160401b0398891617600160401b600160801b03191668010000000000000000948916850217600160801b600160c01b031916600160801b93891693909302929092176001600160c01b03908116600160c01b93891684021790965560e0808a015160038a0180546101008d01519816918a1691909117600160401b600160e01b0319166001600160a01b0397881690950294909417909355610120890151600490980180546101408b0151610160909b01516001600160a01b0319909116999096169890981763ffffffff60a01b1916600160a01b9990931c989098029190911790931691909316909102179091556001600160c01b03199290911b919091169052620003e8565b3390565b6200025762000332565b506040805161018081018252600080825260208083018290528351828152908101845292820192909252606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b82805482825590600052602060002090810192821562000320579160200282015b828111156200032057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620002e9565b506200032e92915062000397565b5090565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b5b808211156200032e5780546001600160a01b031916815560010162000398565b600060208284031215620003ca578081fd5b81516001600160401b0381168114620003e1578182fd5b9392505050565b60805160c01c611eeb62000415600039806105fa5280610b7e528061111e52806112a35250611eeb6000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c806373ed43db116100f9578063cd21aee711610097578063d9a58f2411610071578063d9a58f2414610351578063e008cb6214610364578063f2fde38b14610384578063fa84ea0214610397576101a8565b8063cd21aee714610323578063ce9919b81461032b578063d1ac2e521461033e576101a8565b80639d63753e116100d35780639d63753e146102d7578063bcf67268146102ea578063c7c6e9f4146102fd578063c9515c5814610310576101a8565b806373ed43db146102a957806381e905a3146102bc5780638da5cb5b146102cf576101a8565b80635dc6fdb81161016657806364e9f6711161014057806364e9f67114610266578063660744dc1461026e578063715018a61461028e578063719f2e1714610296576101a8565b80635dc6fdb81461022d578063606df5141461024057806362fced0e14610253576101a8565b806298fa22146101ad5780630f0aae6f146101e057806318b5e830146101f5578063287447c4146101ff5780632b2cc6c414610207578063564093fc1461021a575b600080fd5b6101c06101bb3660046118ee565b6103aa565b6040516101d79b9a99989796959493929190611e2d565b60405180910390f35b6101e8610437565b6040516101d79190611e19565b6101fd61043d565b005b6101e8610a88565b6101fd61021536600461182a565b610a8e565b6101fd610228366004611906565b610aef565b6101fd61023b366004611906565b610b47565b6101fd61024e366004611906565b610bfa565b6101fd610261366004611858565b610c57565b6101c0610d35565b61028161027c366004611906565b610da1565b6040516101d791906119c0565b6101fd610dd9565b6101fd6102a4366004611906565b610e58565b6101fd6102b7366004611906565b610eb5565b6101fd6102ca366004611906565b610f0d565b610281610f6f565b6101fd6102e5366004611906565b610f7e565b6101fd6102f836600461182a565b61102e565b6101fd61030b366004611906565b611085565b6101fd61031e366004611906565b6110e7565b6101e86112a1565b6101fd610339366004611906565b6112c5565b6101fd61034c3660046118c6565b61131d565b6101e861035f366004611906565b611376565b610377610372366004611906565b6113a8565b6040516101d79190611d1d565b6101fd61039236600461182a565b611526565b6102816103a5366004611921565b6115dc565b600181815481106103b757fe5b600091825260209091206005909102018054600282015460038301546004909301546001600160401b038084169550600160401b9384900481169483821694808504831694600160801b8104841694600160c01b91829004851694848116946001600160a01b03949004841693821692600160a01b830460e01b9204168b565b60015490565b610445611637565b6000546001600160a01b0390811691161461047b5760405162461bcd60e51b815260040161047290611c4f565b60405180910390fd5b60015467fffffffffffffffe116104a45760405162461bcd60e51b815260040161047290611a6e565b6104ac6116b4565b6001805460001981019081106104be57fe5b60009182526020918290206040805161018081018252600590930290910180546001600160401b038082168552600160401b9091041683850152600181018054835181870281018701855281815294959294938601939283018282801561054e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610530575b50505091835250506002828101546001600160401b038082166020850152600160401b80830482166040860152600160801b830482166060860152600160c01b9283900482166080860152600386015480831660a08701528190046001600160a01b0390811660c087015260049096015495861660e080870191909152600160a01b8704901b6001600160e01b03191661010086015291909404841661012090930192909252549293507f0000000000000000000000000000000000000000000000000000000000000000821643019204161161063d5760405162461bcd60e51b815260040161047290611b7a565b60808101516001600160401b0316156106d05780516002546001600160401b0391821691161161067f5760405162461bcd60e51b815260040161047290611cd2565b8051600254608083015160208401516001600160401b038084169490940393918402018116600160401b90920416146106ca5760405162461bcd60e51b815260040161047290611c84565b506106ff565b80516002546001600160401b039081169116146106ff5760405162461bcd60e51b815260040161047290611abe565b6001805480820182556000919091526002805460059092027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68101805467ffffffffffffffff19166001600160401b03948516178082558354600160401b9081900490951690940267ffffffffffffffff60401b19909416939093178355600380549293926107b1927fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7019190611719565b5060028281018054918301805467ffffffffffffffff199081166001600160401b03948516178083558354600160401b908190048616810267ffffffffffffffff60401b19909216919091178084558454600160801b9081900487160267ffffffffffffffff60801b19909116178084559354600160c01b90819004861681026001600160c01b03958616179093556003808801805491880180549094169187169190911780845590546001600160a01b03908390048116909202600160401b600160e01b031990911617909155600495860180549690950180546001600160a01b0319169690911695909517808655845463ffffffff600160a01b91829004160263ffffffff60a01b1990911617808655935481900490921690910291161790556108db61163b565b80516002805460208085015167ffffffffffffffff199092166001600160401b039485161767ffffffffffffffff60401b1916600160401b94909216939093021781556040830151805191926109379260039290910190611769565b506060820151600282018054608085015160a086015160c087015167ffffffffffffffff199384166001600160401b039687161767ffffffffffffffff60401b1916600160401b93871684021767ffffffffffffffff60801b1916600160801b92871692909202919091176001600160c01b03908116600160c01b92871683021790945560e0808801516003880180546101008b0151961691881691909117600160401b600160e01b0319166001600160a01b039586169094029390931790925561012087015160049096018054610140890151610160909901516001600160a01b0319909116979094169690961763ffffffff60a01b1916600160a01b9790921c969096021790911691169092029190911790556001546040517f38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae691610a7d91611e19565b60405180910390a150565b60035490565b610a96611637565b6000546001600160a01b03908116911614610ac35760405162461bcd60e51b815260040161047290611c4f565b600580546001600160a01b03909216600160401b02600160401b600160e01b0319909216919091179055565b610af7611637565b6000546001600160a01b03908116911614610b245760405162461bcd60e51b815260040161047290611c4f565b6005805467ffffffffffffffff19166001600160401b0392909216919091179055565b610b4f611637565b6000546001600160a01b03908116911614610b7c5760405162461bcd60e51b815260040161047290611c4f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316816001600160401b031610610bcd5760405162461bcd60e51b8152600401610472906119d4565b600480546001600160401b03909216600160401b0267ffffffffffffffff60401b19909216919091179055565b610c02611637565b6000546001600160a01b03908116911614610c2f5760405162461bcd60e51b815260040161047290611c4f565b600480546001600160401b03909216600160c01b026001600160c01b03909216919091179055565b610c5f611637565b6000546001600160a01b03908116911614610c8c5760405162461bcd60e51b815260040161047290611c4f565b6003546001600160401b038290031015610cb85760405162461bcd60e51b815260040161047290611b2a565b60005b6001600160401b038116821115610d3057600383836001600160401b038416818110610ce357fe5b9050602002016020810190610cf8919061182a565b815460018082018455600093845260209093200180546001600160a01b0319166001600160a01b039290921691909117905501610cbb565b505050565b6002546004546005546006546001600160401b0380851694600160401b9081900482169482811694828204841694600160801b8304851694600160c01b938490048116948184169493046001600160a01b039081169390831692600160a01b810460e01b92919004168b565b60006002600101826001600160401b031681548110610dbc57fe5b6000918252602090912001546001600160a01b031690505b919050565b610de1611637565b6000546001600160a01b03908116911614610e0e5760405162461bcd60e51b815260040161047290611c4f565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b610e60611637565b6000546001600160a01b03908116911614610e8d5760405162461bcd60e51b815260040161047290611c4f565b600680546001600160401b03909216600160c01b026001600160c01b03909216919091179055565b610ebd611637565b6000546001600160a01b03908116911614610eea5760405162461bcd60e51b815260040161047290611c4f565b6004805467ffffffffffffffff19166001600160401b0392909216919091179055565b610f15611637565b6000546001600160a01b03908116911614610f425760405162461bcd60e51b815260040161047290611c4f565b600280546001600160401b03909216600160401b0267ffffffffffffffff60401b19909216919091179055565b6000546001600160a01b031690565b610f86611637565b6000546001600160a01b03908116911614610fb35760405162461bcd60e51b815260040161047290611c4f565b6003546001600160401b038216811061101e5760005b826001600160401b0316816001600160401b03161015611018576003805480610fee57fe5b600082815260209020810160001990810180546001600160a01b0319169055019055600101610fc9565b5061102a565b61102a600360006117be565b5050565b611036611637565b6000546001600160a01b039081169116146110635760405162461bcd60e51b815260040161047290611c4f565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b61108d611637565b6000546001600160a01b039081169116146110ba5760405162461bcd60e51b815260040161047290611c4f565b600480546001600160401b03909216600160801b0267ffffffffffffffff60801b19909216919091179055565b6110ef611637565b6000546001600160a01b0390811691161461111c5760405162461bcd60e51b815260040161047290611c4f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160401b03164301816001600160401b03161161116f5760405162461bcd60e51b815260040161047290611c05565b60015460001981015b80156112395760006001828154811061118d57fe5b6000918252602090912060059091020180549091506001600160401b03808616600160401b90920416106112295760018054806111c657fe5b60008281526020812060056000199093019283020180546fffffffffffffffffffffffffffffffff191681559061120060018301826117be565b506000600282018190556003820180546001600160e01b0319169055600490910155905561122f565b50611239565b5060001901611178565b506001546001600160401b038216116112645760405162461bcd60e51b815260040161047290611bbf565b6001546040517f202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf9161129591611e19565b60405180910390a15050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6112cd611637565b6000546001600160a01b039081169116146112fa5760405162461bcd60e51b815260040161047290611c4f565b6002805467ffffffffffffffff19166001600160401b0392909216919091179055565b611325611637565b6000546001600160a01b039081169116146113525760405162461bcd60e51b815260040161047290611c4f565b6006805460e09290921c600160a01b0263ffffffff60a01b19909216919091179055565b60006001826001600160401b03168154811061138e57fe5b600091825260209091206001600590920201015492915050565b6113b06116b4565b600154600019015b6000600182815481106113c757fe5b6000918252602090912060059091020180549091506001600160401b0380861691161161151c57604080516101808101825282546001600160401b038082168352600160401b90910416602080830191909152600184018054845181840281018401865281815293948694908601939092919083018282801561147357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611455575b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9283900482166080850152600385015480831660a0860152046001600160a01b0390811660c085015260049094015493841660e080850191909152600160a01b8504901b6001600160e01b0319166101008401529204909116610120909101529250610dd4915050565b50600019016113b8565b61152e611637565b6000546001600160a01b0390811691161461155b5760405162461bcd60e51b815260040161047290611c4f565b6001600160a01b0381166115815760405162461bcd60e51b815260040161047290611a28565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006001836001600160401b0316815481106115f457fe5b9060005260206000209060050201600101826001600160401b03168154811061161957fe5b6000918252602090912001546001600160a01b031690505b92915050565b3390565b6116436116b4565b506040805161018081018252600080825260208083018290528351828152908101845292820192909252606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b8280548282559060005260206000209081019282156117595760005260206000209182015b8281111561175957825482559160010191906001019061173e565b506117659291506117df565b5090565b828054828255906000526020600020908101928215611759579160200282015b8281111561175957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611789565b50805460008255906000526020600020908101906117dc91906117fe565b50565b5b808211156117655780546001600160a01b03191681556001016117e0565b5b8082111561176557600081556001016117ff565b80356001600160401b038116811461163157600080fd5b60006020828403121561183b578081fd5b81356001600160a01b0381168114611851578182fd5b9392505050565b6000806020838503121561186a578081fd5b82356001600160401b0380821115611880578283fd5b818501915085601f830112611893578283fd5b8135818111156118a1578384fd5b86602080830285010111156118b4578384fd5b60209290920196919550909350505050565b6000602082840312156118d7578081fd5b81356001600160e01b031981168114611851578182fd5b6000602082840312156118ff578081fd5b5035919050565b600060208284031215611917578081fd5b6118518383611813565b60008060408385031215611933578182fd5b61193d8484611813565b915061194c8460208501611813565b90509250929050565b6001600160a01b03169052565b6000815180845260208085019450808401835b8381101561199a5781516001600160a01b031687529582019590820190600101611975565b509495945050505050565b6001600160e01b0319169052565b6001600160401b03169052565b6001600160a01b0391909116815260200190565b60208082526034908201527f436f6e666967436f6e74726163743a206261746368207370616e206e6f7420736040820152730686f72746572207468616e2068656164732075760641b606082015260800190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b60208082526030908201527f436f6e666967436f6e74726163743a206e756d626572206f6620636f6e66696760408201526f1cc8195e18d959591cc81d5a5b9d0d8d60821b606082015260800190565b60208082526046908201527f436f6e666967436f6e74726163743a207472616e736974696f6e2066726f6d2060408201527f696e61637469766520636f6e66696720776974682077726f6e67207374617274606082015265040d2dcc8caf60d31b608082015260a00190565b60208082526030908201527f436f6e666967436f6e74726163743a206e756d626572206f66206b657970657260408201526f1cc8195e18d959591cc81d5a5b9d0d8d60821b606082015260800190565b60208082526025908201527f436f6e666967436f6e74726163743a20737461727420626c6f636b20746f6f206040820152646561726c7960d81b606082015260800190565b60208082526026908201527f436f6e666967436f6e74726163743a206e6f20636f6e6669677320756e73636860408201526519591d5b195960d21b606082015260800190565b6020808252602a908201527f436f6e666967436f6e74726163743a2066726f6d20737461727420626c6f636b60408201526920746f6f206561726c7960b01b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602e908201527f436f6e666967436f6e74726163743a20636f6e666967207472616e736974696f60408201526d6e206e6f74207365616d6c65737360901b606082015260800190565b6020808252602b908201527f436f6e666967436f6e74726163743a20737461727420626174636820696e646560408201526a1e081d1bdbc81cdb585b1b60aa1b606082015260800190565b600060208252611d316020830184516119b3565b6020830151611d4360408401826119b3565b506040830151610180806060850152611d606101a0850183611962565b91506060850151611d7460808601826119b3565b506080850151611d8760a08601826119b3565b5060a0850151611d9a60c08601826119b3565b5060c0850151611dad60e08601826119b3565b5060e0850151610100611dc2818701836119b3565b8601519050610120611dd686820183611955565b8601519050610140611dea86820183611955565b8601519050610160611dfe868201836119a5565b8601519050611e0f858301826119b3565b5090949350505050565b6001600160401b0391909116815260200190565b6001600160401b038c811682528b811660208301528a811660408301528981166060830152888116608083015287811660a0830152861660c08201526001600160a01b0385811660e083015284166101008201526001600160e01b031983166101208201526101608101611ea56101408301846119b3565b9c9b50505050505050505050505056fea26469706673582212205ebe8e3fe50972abc277dda1a9c27c8dedf3742763396a4cfbfff031da62af8e64736f6c63430007010033"

// DeployConfigContract deploys a new Ethereum contract, binding an instance of ConfigContract to it.
func DeployConfigContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configChangeHeadsUpBlocks uint64) (common.Address, *types.Transaction, *ConfigContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConfigContractBin), backend, _configChangeHeadsUpBlocks)
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

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 _configIndex, uint64 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) ConfigKeypers(opts *bind.CallOpts, _configIndex uint64, _keyperIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configKeypers", _configIndex, _keyperIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 _configIndex, uint64 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) ConfigKeypers(_configIndex uint64, _keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, _configIndex, _keyperIndex)
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xfa84ea02.
//
// Solidity: function configKeypers(uint64 _configIndex, uint64 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) ConfigKeypers(_configIndex uint64, _keyperIndex uint64) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, _configIndex, _keyperIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 _configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCaller) ConfigNumKeypers(opts *bind.CallOpts, _configIndex uint64) (uint64, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "configNumKeypers", _configIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 _configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractSession) ConfigNumKeypers(_configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, _configIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0xd9a58f24.
//
// Solidity: function configNumKeypers(uint64 _configIndex) view returns(uint64)
func (_ConfigContract *ConfigContractCallerSession) ConfigNumKeypers(_configIndex uint64) (uint64, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, _configIndex)
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

	outstruct.StartBatchIndex = out[0].(uint64)
	outstruct.StartBlockNumber = out[1].(uint64)
	outstruct.Threshold = out[2].(uint64)
	outstruct.BatchSpan = out[3].(uint64)
	outstruct.BatchSizeLimit = out[4].(uint64)
	outstruct.TransactionSizeLimit = out[5].(uint64)
	outstruct.TransactionGasLimit = out[6].(uint64)
	outstruct.FeeReceiver = out[7].(common.Address)
	outstruct.TargetAddress = out[8].(common.Address)
	outstruct.TargetFunctionSelector = out[9].([4]byte)
	outstruct.ExecutionTimeout = out[10].(uint64)

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

// GetConfig is a free data retrieval call binding the contract method 0xe008cb62.
//
// Solidity: function getConfig(uint64 _batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCaller) GetConfig(opts *bind.CallOpts, _batchIndex uint64) (BatchConfig, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "getConfig", _batchIndex)

	if err != nil {
		return *new(BatchConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BatchConfig)).(*BatchConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xe008cb62.
//
// Solidity: function getConfig(uint64 _batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractSession) GetConfig(_batchIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.GetConfig(&_ConfigContract.CallOpts, _batchIndex)
}

// GetConfig is a free data retrieval call binding the contract method 0xe008cb62.
//
// Solidity: function getConfig(uint64 _batchIndex) view returns((uint64,uint64,address[],uint64,uint64,uint64,uint64,uint64,address,address,bytes4,uint64))
func (_ConfigContract *ConfigContractCallerSession) GetConfig(_batchIndex uint64) (BatchConfig, error) {
	return _ConfigContract.Contract.GetConfig(&_ConfigContract.CallOpts, _batchIndex)
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

	outstruct.StartBatchIndex = out[0].(uint64)
	outstruct.StartBlockNumber = out[1].(uint64)
	outstruct.Threshold = out[2].(uint64)
	outstruct.BatchSpan = out[3].(uint64)
	outstruct.BatchSizeLimit = out[4].(uint64)
	outstruct.TransactionSizeLimit = out[5].(uint64)
	outstruct.TransactionGasLimit = out[6].(uint64)
	outstruct.FeeReceiver = out[7].(common.Address)
	outstruct.TargetAddress = out[8].(common.Address)
	outstruct.TargetFunctionSelector = out[9].([4]byte)
	outstruct.ExecutionTimeout = out[10].(uint64)

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

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 _index) view returns(address)
func (_ConfigContract *ConfigContractCaller) NextConfigKeypers(opts *bind.CallOpts, _index uint64) (common.Address, error) {
	var out []interface{}
	err := _ConfigContract.contract.Call(opts, &out, "nextConfigKeypers", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 _index) view returns(address)
func (_ConfigContract *ConfigContractSession) NextConfigKeypers(_index uint64) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, _index)
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0x660744dc.
//
// Solidity: function nextConfigKeypers(uint64 _index) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) NextConfigKeypers(_index uint64) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, _index)
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
// Solidity: function nextConfigAddKeypers(address[] _newKeypers) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigAddKeypers(opts *bind.TransactOpts, _newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigAddKeypers", _newKeypers)
}

// NextConfigAddKeypers is a paid mutator transaction binding the contract method 0x62fced0e.
//
// Solidity: function nextConfigAddKeypers(address[] _newKeypers) returns()
func (_ConfigContract *ConfigContractSession) NextConfigAddKeypers(_newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigAddKeypers(&_ConfigContract.TransactOpts, _newKeypers)
}

// NextConfigAddKeypers is a paid mutator transaction binding the contract method 0x62fced0e.
//
// Solidity: function nextConfigAddKeypers(address[] _newKeypers) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigAddKeypers(_newKeypers []common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigAddKeypers(&_ConfigContract.TransactOpts, _newKeypers)
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
// Solidity: function nextConfigSetBatchSizeLimit(uint64 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSizeLimit(opts *bind.TransactOpts, _batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSizeLimit", _batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7c6e9f4.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint64 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSizeLimit(_batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, _batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7c6e9f4.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint64 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSizeLimit(_batchSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, _batchSizeLimit)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 _batchSpan) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSpan(opts *bind.TransactOpts, _batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSpan", _batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 _batchSpan) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSpan(_batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, _batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x5dc6fdb8.
//
// Solidity: function nextConfigSetBatchSpan(uint64 _batchSpan) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSpan(_batchSpan uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, _batchSpan)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 _executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetExecutionTimeout(opts *bind.TransactOpts, _executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetExecutionTimeout", _executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 _executionTimeout) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetExecutionTimeout(_executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetExecutionTimeout(&_ConfigContract.TransactOpts, _executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0x719f2e17.
//
// Solidity: function nextConfigSetExecutionTimeout(uint64 _executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetExecutionTimeout(_executionTimeout uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetExecutionTimeout(&_ConfigContract.TransactOpts, _executionTimeout)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address _feeReceiver) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetFeeReceiver", _feeReceiver)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address _feeReceiver) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetFeeReceiver(&_ConfigContract.TransactOpts, _feeReceiver)
}

// NextConfigSetFeeReceiver is a paid mutator transaction binding the contract method 0x2b2cc6c4.
//
// Solidity: function nextConfigSetFeeReceiver(address _feeReceiver) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetFeeReceiver(&_ConfigContract.TransactOpts, _feeReceiver)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBatchIndex(opts *bind.TransactOpts, _startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBatchIndex", _startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBatchIndex(_startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, _startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0xce9919b8.
//
// Solidity: function nextConfigSetStartBatchIndex(uint64 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBatchIndex(_startBatchIndex uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, _startBatchIndex)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBlockNumber(opts *bind.TransactOpts, _startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBlockNumber", _startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBlockNumber(_startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBlockNumber(&_ConfigContract.TransactOpts, _startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x81e905a3.
//
// Solidity: function nextConfigSetStartBlockNumber(uint64 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBlockNumber(_startBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBlockNumber(&_ConfigContract.TransactOpts, _startBlockNumber)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address _targetAddress) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTargetAddress(opts *bind.TransactOpts, _targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTargetAddress", _targetAddress)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address _targetAddress) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTargetAddress(_targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetAddress(&_ConfigContract.TransactOpts, _targetAddress)
}

// NextConfigSetTargetAddress is a paid mutator transaction binding the contract method 0xbcf67268.
//
// Solidity: function nextConfigSetTargetAddress(address _targetAddress) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTargetAddress(_targetAddress common.Address) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetAddress(&_ConfigContract.TransactOpts, _targetAddress)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 _targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTargetFunctionSelector(opts *bind.TransactOpts, _targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTargetFunctionSelector", _targetFunctionSelector)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 _targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTargetFunctionSelector(_targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetFunctionSelector(&_ConfigContract.TransactOpts, _targetFunctionSelector)
}

// NextConfigSetTargetFunctionSelector is a paid mutator transaction binding the contract method 0xd1ac2e52.
//
// Solidity: function nextConfigSetTargetFunctionSelector(bytes4 _targetFunctionSelector) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTargetFunctionSelector(_targetFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTargetFunctionSelector(&_ConfigContract.TransactOpts, _targetFunctionSelector)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 _threshold) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetThreshold(opts *bind.TransactOpts, _threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetThreshold", _threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 _threshold) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, _threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0x73ed43db.
//
// Solidity: function nextConfigSetThreshold(uint64 _threshold) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, _threshold)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionGasLimit(opts *bind.TransactOpts, _transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionGasLimit", _transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionGasLimit(_transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, _transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0x564093fc.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint64 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionGasLimit(_transactionGasLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, _transactionGasLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionSizeLimit(opts *bind.TransactOpts, _transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionSizeLimit", _transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionSizeLimit(_transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionSizeLimit(&_ConfigContract.TransactOpts, _transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x606df514.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint64 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionSizeLimit(_transactionSizeLimit uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionSizeLimit(&_ConfigContract.TransactOpts, _transactionSizeLimit)
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
// Solidity: function unscheduleConfigs(uint64 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) UnscheduleConfigs(opts *bind.TransactOpts, _fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "unscheduleConfigs", _fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xc9515c58.
//
// Solidity: function unscheduleConfigs(uint64 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) UnscheduleConfigs(_fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.UnscheduleConfigs(&_ConfigContract.TransactOpts, _fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xc9515c58.
//
// Solidity: function unscheduleConfigs(uint64 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) UnscheduleConfigs(_fromStartBlockNumber uint64) (*types.Transaction, error) {
	return _ConfigContract.Contract.UnscheduleConfigs(&_ConfigContract.TransactOpts, _fromStartBlockNumber)
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
const DepositContractABI = "[{\"inputs\":[{\"internalType\":\"contractIERC777\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalDelayBlocks\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalRequestedBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"name\":\"DepositChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWithdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequestedBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slasherAddress\",\"type\":\"address\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositContractBin is the compiled bytecode used for deploying new contracts.
var DepositContractBin = "0x6080604052600080546001600160a01b031916731820a4b7618bde71dce8cdc73aab6c95905fad2417905534801561003657600080fd5b50604051610df0380380610df0833981016040819052610055916100f8565b600180546001600160a01b0319166001600160a01b03838116919091179091556000546040516329965a1d60e01b81529116906329965a1d906100c09030907fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b908290600401610126565b600060405180830381600087803b1580156100da57600080fd5b505af11580156100ee573d6000803e3d6000fd5b5050505050610149565b600060208284031215610109578081fd5b81516001600160a01b038116811461011f578182fd5b9392505050565b6001600160a01b0393841681526020810192909252909116604082015260600190565b610c98806101586000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c8063b611d8d911610066578063b611d8d9146100fb578063b799036c1461010e578063b8ba16fd1461012e578063c96be4cb1461014e578063dbaf21451461016157610092565b806223de291461009757806351cff8d9146100ac5780639b4fed88146100bf578063aabc2496146100e8575b600080fd5b6100aa6100a536600461086b565b610169565b005b6100aa6100ba366004610849565b6101c2565b6100d26100cd366004610849565b61037f565b6040516100df9190610c2e565b60405180910390f35b6100aa6100f6366004610849565b6103a6565b6100d2610109366004610849565b6103f1565b61012161011c366004610849565b61041f565b6040516100df9190610964565b61014161013c366004610849565b610447565b6040516100df9190610c25565b6100aa61015c366004610849565b610462565b6100aa61054a565b6001546001600160a01b0316331461019c5760405162461bcd60e51b815260040161019390610bde565b60405180910390fd5b60006101aa84860186610915565b90506101b7888783610672565b505050505050505050565b6101ca6107bf565b5033600090815260036020908152604091829020825160808101845281548082526001909201546001600160401b0380821694830194909452600160401b810490931693810193909352600160801b90910460ff16151560608301526102425760405162461bcd60e51b815260040161019390610a3d565b80606001511561024e57fe5b600081604001516001600160401b03161161027b5760405162461bcd60e51b8152600401610193906109a0565b80602001518160400151016001600160401b03164310156102ae5760405162461bcd60e51b8152600401610193906109ed565b336000908152600360205260408082209182556001918201805470ffffffffffffffffffffffffffffffffff19169055905482519151634decdde360e11b81526001600160a01b0390911691639bd9bbc69161030e91869160040161093c565b600060405180830381600087803b15801561032857600080fd5b505af115801561033c573d6000803e3d6000fd5b50505050336001600160a01b0316600080516020610c4383398151915260008060006001600060405161037395949392919061096f565b60405180910390a25050565b6001600160a01b03166000908152600360205260409020600101546001600160401b031690565b6002546001600160a01b0316156103cf5760405162461bcd60e51b815260040161019390610b92565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0316600090815260036020526040902060010154600160401b90046001600160401b031690565b6001600160a01b0316600090815260036020526040902060010154600160801b900460ff1690565b6001600160a01b031660009081526003602052604090205490565b6002546001600160a01b0316331461047957600080fd5b6104816107bf565b506001600160a01b03811660008181526003602081815260408084208151608081018352600180830180548884528387018981528487018a8152606086018581528c8c5299909852845190955593519551965167ffffffffffffffff199094166001600160401b039687161767ffffffffffffffff60401b1916600160401b96909716959095029590951760ff60801b1916600160801b921515929092029190911790925551909392600080516020610c4383398151915292610373928291829182919061096f565b6105526107bf565b5033600090815260036020908152604091829020825160808101845281548082526001909201546001600160401b0380821694830194909452600160401b810490931693810193909352600160801b90910460ff16151560608301526105ca5760405162461bcd60e51b815260040161019390610a3d565b8060600151156105d657fe5b60408101516001600160401b0316156106015760405162461bcd60e51b815260040161019390610b45565b33600081815260036020908152604080832060010180546001600160401b0343908116600160401b0267ffffffffffffffff60401b19909216919091179091558551928601519151600080516020610c438339815191529461066794939291819061096f565b60405180910390a250565b61067a6107bf565b506001600160a01b0383166000908152600360209081526040918290208251608081018452815481526001909101546001600160401b03808216938301849052600160401b8204811694830194909452600160801b900460ff161515606082015291831610156106fc5760405162461bcd60e51b815260040161019390610a74565b60408101516001600160401b0316156107275760405162461bcd60e51b815260040161019390610ac9565b8060600151156107495760405162461bcd60e51b815260040161019390610b10565b80516001600160a01b03851660008181526003602052604080822093870184556001909301805467ffffffffffffffff19166001600160401b038716179055835192519192600080516020610c43833981519152926107b1929188019187918190819061096f565b60405180910390a250505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b80356001600160a01b03811681146107fd57600080fd5b92915050565b60008083601f840112610814578182fd5b5081356001600160401b0381111561082a578182fd5b60208301915083602082850101111561084257600080fd5b9250929050565b60006020828403121561085a578081fd5b61086483836107e6565b9392505050565b60008060008060008060008060c0898b031215610886578384fd5b6108908a8a6107e6565b975061089f8a60208b016107e6565b96506108ae8a60408b016107e6565b95506060890135945060808901356001600160401b03808211156108d0578586fd5b6108dc8c838d01610803565b909650945060a08b01359150808211156108f4578384fd5b506109018b828c01610803565b999c989b5096995094979396929594505050565b600060208284031215610926578081fd5b81356001600160401b0381168114610864578182fd5b6001600160a01b03929092168252602082015260606040820181905260009082015260800190565b901515815260200190565b9485526001600160401b03938416602086015291909216604084015290151560608301521515608082015260a00190565b6020808252602d908201527f4465706f736974436f6e74726163743a207769746864726177616c206e6f742060408201526c1c995c5d595cdd1959081e595d609a1b606082015260800190565b60208082526030908201527f4465706f736974436f6e74726163743a207769746864726177616c2064656c6160408201526f1e481b9bdd081c185cdcd959081e595d60821b606082015260800190565b6020808252601b908201527f4465706f736974436f6e74726163743a206e6f206465706f7369740000000000604082015260600190565b60208082526035908201527f4465706f736974436f6e74726163743a207769746864726177616c2064656c616040820152741e4818d85b9b9bdd08189948191958dc99585cd959605a1b606082015260800190565b60208082526027908201527f4465706f736974436f6e74726163743a207769746864726177616c20696e2070604082015266726f677265737360c81b606082015260800190565b6020808252818101527f4465706f736974436f6e74726163743a206163636f756e7420736c6173686564604082015260600190565b6020808252602d908201527f4465706f736974436f6e74726163743a207769746864726177616c20616c726560408201526c18591e481c995c5d595cdd1959609a1b606082015260800190565b6020808252602c908201527f4465706f736974436f6e74726163743a20736c6173686572206164647265737360408201526b08185b1c9958591e481cd95d60a21b606082015260800190565b60208082526027908201527f4465706f736974436f6e74726163743a20726563656976656420696e76616c6960408201526632103a37b5b2b760c91b606082015260800190565b90815260200190565b6001600160401b039190911681526020019056fe04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97fa264697066735822122027ab48a54f4c59d39562a7433ea3ce5e18f2a3624d95ef900aea7c227d6990fd64736f6c63430007010033"

// DeployDepositContract deploys a new Ethereum contract, binding an instance of DepositContract to it.
func DeployDepositContract(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *DepositContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositContractBin), backend, token)
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
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_DepositContract *DepositContractTransactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "tokensReceived", operator, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_DepositContract *DepositContractSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _DepositContract.Contract.TokensReceived(&_DepositContract.TransactOpts, operator, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_DepositContract *DepositContractTransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _DepositContract.Contract.TokensReceived(&_DepositContract.TransactOpts, operator, from, to, amount, userData, operatorData)
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
var ECDSABin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fbff3f245b3eb3f42d3e3e129f50bd79f17083e7a5821421af6fb9a517ca073464736f6c63430007010033"

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
var ERC777Bin = "0x60806040523480156200001157600080fd5b506040516200237a3803806200237a833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084640100000000821115620001bc57600080fd5b908301906020820185811115620001d257600080fd5b8251866020820283011164010000000082111715620001f057600080fd5b82525081516020918201928201910280838360005b838110156200021f57818101518382015260200162000205565b5050505091909101604052505084516200024392506002915060208601906200040b565b508151620002599060039060208501906200040b565b5080516200026f90600490602084019062000490565b5060005b600454811015620002cf57600160056000600484815481106200029257fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560010162000273565b50604080516329965a1d60e01b815230600482018190527fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce2177054602483015260448201529051731820a4b7618bde71dce8cdc73aab6c95905fad24916329965a1d91606480830192600092919082900301818387803b1580156200035057600080fd5b505af115801562000365573d6000803e3d6000fd5b5050604080516329965a1d60e01b815230600482018190527faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a602483015260448201529051731820a4b7618bde71dce8cdc73aab6c95905fad2493506329965a1d9250606480830192600092919082900301818387803b158015620003e957600080fd5b505af1158015620003fe573d6000803e3d6000fd5b505050505050506200052e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200044e57805160ff19168380011785556200047e565b828001600101855582156200047e579182015b828111156200047e57825182559160200191906001019062000461565b506200048c929150620004f6565b5090565b828054828255906000526020600020908101928215620004e8579160200282015b82811115620004e857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004b1565b506200048c9291506200050d565b5b808211156200048c5760008155600101620004f7565b5b808211156200048c5780546001600160a01b03191681556001016200050e565b611e3c806200053e6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b63711461052a578063dd62ed3e14610558578063fad8b32a14610586578063fc673c4f146105ac578063fe9d9303146106ea57610116565b8063959b8c3f1461041757806395d89b411461043d5780639bd9bbc614610445578063a9059cbb146104fe57610116565b806323b872dd116100e957806323b872dd1461024a578063313ce56714610280578063556f0dc71461029e57806362ad1b83146102a657806370a08231146103f157610116565b806306e485381461011b57806306fdde0314610173578063095ea7b3146101f057806318160ddd14610230575b600080fd5b610123610795565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561015f578181015183820152602001610147565b505050509050019250505060405180910390f35b61017b6107f7565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b557818101518382015260200161019d565b50505050905090810190601f1680156101e25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61021c6004803603604081101561020657600080fd5b506001600160a01b038135169060200135610881565b604080519115158252519081900360200190f35b6102386108a3565b60408051918252519081900360200190f35b61021c6004803603606081101561026057600080fd5b506001600160a01b038135811691602081013590911690604001356108a9565b610288610a26565b6040805160ff9092168252519081900360200190f35b610238610a2b565b6103ef600480360360a08110156102bc57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156102f657600080fd5b82018360208201111561030857600080fd5b803590602001918460018302840111600160201b8311171561032957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561037b57600080fd5b82018360208201111561038d57600080fd5b803590602001918460018302840111600160201b831117156103ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a30945050505050565b005b6102386004803603602081101561040757600080fd5b50356001600160a01b0316610a92565b6103ef6004803603602081101561042d57600080fd5b50356001600160a01b0316610aad565b61017b610bf9565b6103ef6004803603606081101561045b57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561048a57600080fd5b82018360208201111561049c57600080fd5b803590602001918460018302840111600160201b831117156104bd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c5a945050505050565b61021c6004803603604081101561051457600080fd5b506001600160a01b038135169060200135610c84565b61021c6004803603604081101561054057600080fd5b506001600160a01b0381358116916020013516610d5d565b6102386004803603604081101561056e57600080fd5b506001600160a01b0381358116916020013516610dff565b6103ef6004803603602081101561059c57600080fd5b50356001600160a01b0316610e2a565b6103ef600480360360808110156105c257600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156105f157600080fd5b82018360208201111561060357600080fd5b803590602001918460018302840111600160201b8311171561062457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561067657600080fd5b82018360208201111561068857600080fd5b803590602001918460018302840111600160201b831117156106a957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f76945050505050565b6103ef6004803603604081101561070057600080fd5b81359190810190604081016020820135600160201b81111561072157600080fd5b82018360208201111561073357600080fd5b803590602001918460018302840111600160201b8311171561075457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610fd4945050505050565b606060048054806020026020016040519081016040528092919081815260200182805480156107ed57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107cf575b5050505050905090565b60028054604080516020601f60001961010060018716150201909416859004938401819004810282018101909252828152606093909290918301828280156107ed5780601f10610855576101008083540402835291602001916107ed565b820191906000526020600020905b81548152906001019060200180831161086357509395945050505050565b60008061088c610ffa565b9050610899818585610ffe565b5060019392505050565b60015490565b60006001600160a01b0383166108f05760405162461bcd60e51b8152600401808060200182810382526024815260200180611d226024913960400191505060405180910390fd5b6001600160a01b0384166109355760405162461bcd60e51b8152600401808060200182810382526026815260200180611d9b6026913960400191505060405180910390fd5b600061093f610ffa565b905061096d8186868660405180602001604052806000815250604051806020016040528060008152506110ea565b610999818686866040518060200160405280600081525060405180602001604052806000815250611317565b6109ed85826109e886604051806060016040528060298152602001611d72602991396001600160a01b03808c166000908152600860209081526040808320938b16835292905220549190611530565b610ffe565b610a1b81868686604051806020016040528060008152506040518060200160405280600081525060006115c7565b506001949350505050565b601290565b600190565b610a41610a3b610ffa565b86610d5d565b610a7c5760405162461bcd60e51b815260040180806020018281038252602c815260200180611d46602c913960400191505060405180910390fd5b610a8b8585858585600161184c565b5050505050565b6001600160a01b031660009081526020819052604090205490565b806001600160a01b0316610abf610ffa565b6001600160a01b03161415610b055760405162461bcd60e51b8152600401808060200182810382526024815260200180611c906024913960400191505060405180910390fd5b6001600160a01b03811660009081526005602052604090205460ff1615610b685760076000610b32610ffa565b6001600160a01b03908116825260208083019390935260409182016000908120918516815292529020805460ff19169055610baf565b600160066000610b76610ffa565b6001600160a01b03908116825260208083019390935260409182016000908120918616815292529020805460ff19169115159190911790555b610bb7610ffa565b6001600160a01b0316816001600160a01b03167ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f960405160405180910390a350565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156107ed5780601f10610855576101008083540402835291602001916107ed565b610c7f610c65610ffa565b84848460405180602001604052806000815250600161184c565b505050565b60006001600160a01b038316610ccb5760405162461bcd60e51b8152600401808060200182810382526024815260200180611d226024913960400191505060405180910390fd5b6000610cd5610ffa565b9050610d038182868660405180602001604052806000815250604051806020016040528060008152506110ea565b610d2f818286866040518060200160405280600081525060405180602001604052806000815250611317565b61089981828686604051806020016040528060008152506040518060200160405280600081525060006115c7565b6000816001600160a01b0316836001600160a01b03161480610dc857506001600160a01b03831660009081526005602052604090205460ff168015610dc857506001600160a01b0380831660009081526007602090815260408083209387168352929052205460ff16155b80610df857506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff165b9392505050565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b610e32610ffa565b6001600160a01b0316816001600160a01b03161415610e825760405162461bcd60e51b8152600401808060200182810382526021815260200180611cb46021913960400191505060405180910390fd5b6001600160a01b03811660009081526005602052604090205460ff1615610eee57600160076000610eb1610ffa565b6001600160a01b03908116825260208083019390935260409182016000908120918616815292529020805460ff1916911515919091179055610f2c565b60066000610efa610ffa565b6001600160a01b03908116825260208083019390935260409182016000908120918516815292529020805460ff191690555b610f34610ffa565b6001600160a01b0316816001600160a01b03167f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa160405160405180910390a350565b610f87610f81610ffa565b85610d5d565b610fc25760405162461bcd60e51b815260040180806020018281038252602c815260200180611d46602c913960400191505060405180910390fd5b610fce84848484611923565b50505050565b610ff6610fdf610ffa565b838360405180602001604052806000815250611923565b5050565b3390565b6001600160a01b0383166110435760405162461bcd60e51b8152600401808060200182810382526025815260200180611c006025913960400191505060405180910390fd5b6001600160a01b0382166110885760405162461bcd60e51b8152600401808060200182810382526023815260200180611de46023913960400191505060405180910390fd5b6001600160a01b03808416600081815260086020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6040805163555ddc6560e11b81526001600160a01b03871660048201527f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe89560248201529051600091731820a4b7618bde71dce8cdc73aab6c95905fad249163aabbb8ca91604480820192602092909190829003018186803b15801561116e57600080fd5b505afa158015611182573d6000803e3d6000fd5b505050506040513d602081101561119857600080fd5b505190506001600160a01b0381161561130e57806001600160a01b03166375ab97828888888888886040518763ffffffff1660e01b815260040180876001600160a01b03168152602001866001600160a01b03168152602001856001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561124357818101518382015260200161122b565b50505050905090810190601f1680156112705780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156112a357818101518382015260200161128b565b50505050905090810190601f1680156112d05780820380516001836020036101000a031916815260200191505b5098505050505050505050600060405180830381600087803b1580156112f557600080fd5b505af1158015611309573d6000803e3d6000fd5b505050505b50505050505050565b61132386868686610fce565b61136083604051806060016040528060278152602001611c47602791396001600160a01b0388166000908152602081905260409020549190611530565b6001600160a01b03808716600090815260208190526040808220939093559086168152205461138f9084611b5d565b600080866001600160a01b03166001600160a01b0316815260200190815260200160002081905550836001600160a01b0316856001600160a01b0316876001600160a01b03167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987868686604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611440578181015183820152602001611428565b50505050905090810190601f16801561146d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156114a0578181015183820152602001611488565b50505050905090810190601f1680156114cd5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a4836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3505050505050565b600081848411156115bf5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561158457818101518382015260200161156c565b50505050905090810190601f1680156115b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6040805163555ddc6560e11b81526001600160a01b03871660048201527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b60248201529051600091731820a4b7618bde71dce8cdc73aab6c95905fad249163aabbb8ca91604480820192602092909190829003018186803b15801561164b57600080fd5b505afa15801561165f573d6000803e3d6000fd5b505050506040513d602081101561167557600080fd5b505190506001600160a01b038116156117ee57806001600160a01b03166223de298989898989896040518763ffffffff1660e01b815260040180876001600160a01b03168152602001866001600160a01b03168152602001856001600160a01b031681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561171f578181015183820152602001611707565b50505050905090810190601f16801561174c5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561177f578181015183820152602001611767565b50505050905090810190601f1680156117ac5780820380516001836020036101000a031916815260200191505b5098505050505050505050600060405180830381600087803b1580156117d157600080fd5b505af11580156117e5573d6000803e3d6000fd5b50505050611842565b811561184257611806866001600160a01b0316611bb7565b156118425760405162461bcd60e51b815260040180806020018281038252604d815260200180611cd5604d913960600191505060405180910390fd5b5050505050505050565b6001600160a01b0386166118915760405162461bcd60e51b8152600401808060200182810382526022815260200180611c256022913960400191505060405180910390fd5b6001600160a01b0385166118ec576040805162461bcd60e51b815260206004820181905260248201527f4552433737373a2073656e6420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b60006118f6610ffa565b90506119068188888888886110ea565b611914818888888888611317565b61130e818888888888886115c7565b6001600160a01b0384166119685760405162461bcd60e51b8152600401808060200182810382526022815260200180611c6e6022913960400191505060405180910390fd5b6000611972610ffa565b90506119818186600087610fce565b611990818660008787876110ea565b6119cd84604051806060016040528060238152602001611dc1602391396001600160a01b0388166000908152602081905260409020549190611530565b6001600160a01b0386166000908152602081905260409020556001546119f39085611bbd565b600181905550846001600160a01b0316816001600160a01b03167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098868686604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611a78578181015183820152602001611a60565b50505050905090810190601f168015611aa55780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611ad8578181015183820152602001611ac0565b50505050905090810190601f168015611b055780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a36040805185815290516000916001600160a01b038816917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050505050565b600082820183811015610df8576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b3b151590565b6000610df883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061153056fe4552433737373a20617070726f76652066726f6d20746865207a65726f20616464726573734552433737373a2073656e642066726f6d20746865207a65726f20616464726573734552433737373a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433737373a206275726e2066726f6d20746865207a65726f20616464726573734552433737373a20617574686f72697a696e672073656c66206173206f70657261746f724552433737373a207265766f6b696e672073656c66206173206f70657261746f724552433737373a20746f6b656e20726563697069656e7420636f6e747261637420686173206e6f20696d706c656d656e74657220666f7220455243373737546f6b656e73526563697069656e744552433737373a207472616e7366657220746f20746865207a65726f20616464726573734552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f7220666f7220686f6c6465724552433737373a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654552433737373a207472616e736665722066726f6d20746865207a65726f20616464726573734552433737373a206275726e20616d6f756e7420657863656564732062616c616e63654552433737373a20617070726f766520746f20746865207a65726f2061646472657373a26469706673582212203bfe4c9d8aee0edda8a367c7500f2ff41c1db77fae1a6ef4552d0537da7673f664736f6c63430007010033"

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
const ExecutorContractABI = "[{\"inputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"contractBatcherContract\",\"name\":\"_batcherContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"}],\"name\":\"CipherExecutionSkipped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"batcherContract\",\"outputs\":[{\"internalType\":\"contractBatcherContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"cipherExecutionReceipts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"}],\"name\":\"executeCipherBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_transactions\",\"type\":\"bytes[]\"}],\"name\":\"executePlainBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_halfStep\",\"type\":\"uint64\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCipherExecutionReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numExecutionHalfSteps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipCipherExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutorContractBin is the compiled bytecode used for deploying new contracts.
var ExecutorContractBin = "0x608060405234801561001057600080fd5b5060405161144638038061144683398101604081905261002f91610060565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100b1565b60008060408385031215610072578182fd5b825161007d81610099565b602084015190925061008e81610099565b809150509250929050565b6001600160a01b03811681146100ae57600080fd5b50565b611386806100c06000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c87190b31161005b578063c87190b3146100dd578063ce0c159614610101578063d57a29d014610114578063fa6385f41461012757610088565b806325b36cbf1461008d5780638f6dccfb146100b6578063beb3b50e146100c0578063bf66a182146100d5575b600080fd5b6100a061009b366004610f3f565b61013c565b6040516100ad9190611231565b60405180910390f35b6100be6101b2565b005b6100c8610369565b6040516100ad9190611016565b6100c8610378565b6100f06100eb366004610f3f565b610387565b6040516100ad959493929190610fb0565b6100be61010f366004610da1565b6103cd565b6100be610122366004610d4a565b6107f5565b61012f610a34565b6040516100ad919061127a565b610144610ba8565b506001600160401b03908116600090815260026020818152604092839020835160a081018552815460ff81161515825261010081046001600160a01b031693820193909352600160a81b90920490941692810192909252600183015460608301529190910154608082015290565b60018054600160a01b900416156101e45760405162461bcd60e51b81526004016101db906110c3565b60405180910390fd5b6001546002600160a01b9091046001600160401b031604610203610bd6565b60005460405163700465b160e11b81526001600160a01b039091169063e008cb629061023390859060040161127a565b60006040518083038186803b15801561024b57600080fd5b505afa15801561025f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102879190810190610dfd565b9050600081608001516001600160401b0316116102b65760405162461bcd60e51b81526004016101db90611109565b80610160015182600101826080015102826020015101016001600160401b03164310156102f55760405162461bcd60e51b81526004016101db9061107a565b600180546001600160401b03600160a01b808304821684018216810267ffffffffffffffff60a01b1990931692909217928390556040517fa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be9361035d9390049091169061127a565b60405180910390a15050565b6001546001600160a01b031681565b6000546001600160a01b031681565b600260208190526000918252604090912080546001820154919092015460ff83169261010081046001600160a01b031692600160a81b9091046001600160401b03169185565b60018054600160a01b900416156103f65760405162461bcd60e51b81526004016101db906110c3565b6001546002600160a01b9091046001600160401b031604610415610bd6565b60005460405163700465b160e11b81526001600160a01b039091169063e008cb629061044590859060040161127a565b60006040518083038186803b15801561045d57600080fd5b505afa158015610471573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104999190810190610dfd565b9050600081608001516001600160401b0316116104c85760405162461bcd60e51b81526004016101db90611109565b816001018160800151028160200151016001600160401b03164310156105005760405162461bcd60e51b81526004016101db9061107a565b806040015151836001600160401b03161061052d5760405162461bcd60e51b81526004016101db906111e5565b8060400151836001600160401b03168151811061054657fe5b60200260200101516001600160a01b0316336001600160a01b03161461057e5760405162461bcd60e51b81526004016101db9061102a565b60015460405163643d7d4560e11b81526001600160a01b039091169063c87afa8a906105b19085906000906004016112a7565b60206040518083038186803b1580156105c957600080fd5b505afa1580156105dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106019190610d89565b861461061f5760405162461bcd60e51b81526004016101db9061114d565b600061063c8261012001518361014001518460e001518989610a4a565b90506040518060a00160405280600115158152602001336001600160a01b03168152602001600160149054906101000a90046001600160401b03166001600160401b031681526020018881526020018281525060026000600160149054906101000a90046001600160401b03166001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160010155608082015181600201559050506001601481819054906101000a90046001600160401b03168092919060010191906101000a8154816001600160401b0302191690836001600160401b03160217905550507f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a600160149054906101000a90046001600160401b0316826040516107e492919061128e565b60405180910390a150505050505050565b60018054600160a01b900481161461081f5760405162461bcd60e51b81526004016101db906110c3565b6001546002600160a01b9091046001600160401b03160461083e610bd6565b60005460405163700465b160e11b81526001600160a01b039091169063e008cb629061086e90859060040161127a565b60006040518083038186803b15801561088657600080fd5b505afa15801561089a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108c29190810190610dfd565b9050600081608001516001600160401b0316116108db57fe5b816001018160800151028160200151016001600160401b03164310156108fd57fe5b600061091a8261012001518361014001518460e001518888610a4a565b6001805460405163643d7d4560e11b81529293506001600160a01b03169163c87afa8a9161094d918791906004016112a7565b60206040518083038186803b15801561096557600080fd5b505afa158015610979573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099d9190610d89565b81146109bb5760405162461bcd60e51b81526004016101db9061119a565b600180546001600160401b03600160a01b808304821684018216810267ffffffffffffffff60a01b1990931692909217928390556040517f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a93610a2593900490911690849061128e565b60405180910390a15050505050565b600154600160a01b90046001600160401b031681565b60008060005b6001600160401b038116841115610b9d576060878686846001600160401b0316818110610a7957fe5b9050602002810190610a8b91906112ce565b604051602401610a9c929190610fe7565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b03838183161783525050505090506000896001600160a01b0316886001600160401b031683604051610af59190610f77565b60006040518083038160008787f1925050503d8060008114610b33576040519150601f19603f3d011682016040523d82523d6000602084013e610b38565b606091505b505090508686846001600160401b0316818110610b5157fe5b9050602002810190610b6391906112ce565b85604051602001610b7693929190610f62565b60405160208183030381529060405280519060200120935050508080600101915050610a50565b509695505050505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b80516001600160a01b0381168114610c5257600080fd5b92915050565b600082601f830112610c68578081fd5b81516001600160401b03811115610c7d578182fd5b6020808202610c8d828201611312565b83815293508184018583018287018401881015610ca957600080fd5b600092505b84831015610cd457610cc08882610c3b565b825260019290920191908301908301610cae565b505050505092915050565b60008083601f840112610cf0578081fd5b5081356001600160401b03811115610d06578182fd5b6020830191508360208083028501011115610d2057600080fd5b9250929050565b80516001600160e01b031981168114610c5257600080fd5b8051610c5281611338565b60008060208385031215610d5c578182fd5b82356001600160401b03811115610d71578283fd5b610d7d85828601610cdf565b90969095509350505050565b600060208284031215610d9a578081fd5b5051919050565b60008060008060608587031215610db6578182fd5b8435935060208501356001600160401b03811115610dd2578283fd5b610dde87828801610cdf565b9094509250506040850135610df281611338565b939692955090935050565b600060208284031215610e0e578081fd5b81516001600160401b0380821115610e24578283fd5b8184019150610180808387031215610e3a578384fd5b610e4381611312565b9050610e4f8684610d3f565b8152610e5e8660208501610d3f565b6020820152604083015182811115610e74578485fd5b610e8087828601610c58565b604083015250610e938660608501610d3f565b6060820152610ea58660808501610d3f565b6080820152610eb78660a08501610d3f565b60a0820152610ec98660c08501610d3f565b60c0820152610edb8660e08501610d3f565b60e08201526101009150610ef186838501610c3b565b828201526101209150610f0686838501610c3b565b828201526101409150610f1b86838501610d27565b828201526101609150610f3086838501610d3f565b91810191909152949350505050565b600060208284031215610f50578081fd5b8135610f5b81611338565b9392505050565b60008385833750909101908152602001919050565b60008251815b81811015610f975760208186018101518583015201610f7d565b81811115610fa55782828501525b509190910192915050565b94151585526001600160a01b039390931660208501526001600160401b039190911660408401526060830152608082015260a00190565b60006020825282602083015282846040840137818301604090810191909152601f909201601f19160101919050565b6001600160a01b0391909116815260200190565b60208082526030908201527f4578656375746f72436f6e74726163743a2073656e646572206973206e6f742060408201526f39b832b1b4b334b2b21035b2bcb832b960811b606082015260800190565b60208082526029908201527f4578656375746f72436f6e74726163743a206261746368206973206e6f7420636040820152681b1bdcd959081e595d60ba1b606082015260800190565b60208082526026908201527f4578656375746f72436f6e74726163743a20756e65787065637465642068616c60408201526506620737465760d41b606082015260800190565b60208082526024908201527f4578656375746f72436f6e74726163743a20636f6e66696720697320696e61636040820152637469766560e01b606082015260800190565b6020808252602d908201527f4578656375746f72436f6e74726163743a20696e636f7272656374206369706860408201526c0cae440c4c2e8c6d040d0c2e6d609b1b606082015260800190565b6020808252602b908201527f4578656375746f72436f6e74726163743a206261746368206861736820646f6560408201526a0e640dcdee840dac2e8c6d60ab1b606082015260800190565b6020808252602c908201527f4578656375746f72436f6e74726163743a206b657970657220696e646578206f60408201526b7574206f6620626f756e647360a01b606082015260800190565b8151151581526020808301516001600160a01b0316908201526040808301516001600160401b031690820152606080830151908201526080918201519181019190915260a00190565b6001600160401b0391909116815260200190565b6001600160401b03929092168252602082015260400190565b6001600160401b038316815260408101600283106112c157fe5b8260208301529392505050565b6000808335601e198436030181126112e4578283fd5b8301803591506001600160401b038211156112fd578283fd5b602001915036819003821315610d2057600080fd5b6040518181016001600160401b038111828210171561133057600080fd5b604052919050565b6001600160401b038116811461134d57600080fd5b5056fea2646970667358221220ea70db213ed6c328246febc488063d30001b1f5ae9879bdeae43161c63af94ae64736f6c63430007010033"

// DeployExecutorContract deploys a new Ethereum contract, binding an instance of ExecutorContract to it.
func DeployExecutorContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configContract common.Address, _batcherContract common.Address) (common.Address, *types.Transaction, *ExecutorContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExecutorContractBin), backend, _configContract, _batcherContract)
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

	outstruct.Executed = out[0].(bool)
	outstruct.Executor = out[1].(common.Address)
	outstruct.HalfStep = out[2].(uint64)
	outstruct.CipherBatchHash = out[3].([32]byte)
	outstruct.BatchHash = out[4].([32]byte)

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
// Solidity: function getReceipt(uint64 _halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractCaller) GetReceipt(opts *bind.CallOpts, _halfStep uint64) (CipherExecutionReceipt, error) {
	var out []interface{}
	err := _ExecutorContract.contract.Call(opts, &out, "getReceipt", _halfStep)

	if err != nil {
		return *new(CipherExecutionReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(CipherExecutionReceipt)).(*CipherExecutionReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0x25b36cbf.
//
// Solidity: function getReceipt(uint64 _halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractSession) GetReceipt(_halfStep uint64) (CipherExecutionReceipt, error) {
	return _ExecutorContract.Contract.GetReceipt(&_ExecutorContract.CallOpts, _halfStep)
}

// GetReceipt is a free data retrieval call binding the contract method 0x25b36cbf.
//
// Solidity: function getReceipt(uint64 _halfStep) view returns((bool,address,uint64,bytes32,bytes32))
func (_ExecutorContract *ExecutorContractCallerSession) GetReceipt(_halfStep uint64) (CipherExecutionReceipt, error) {
	return _ExecutorContract.Contract.GetReceipt(&_ExecutorContract.CallOpts, _halfStep)
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

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0xce0c1596.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, uint64 _keyperIndex) returns()
func (_ExecutorContract *ExecutorContractTransactor) ExecuteCipherBatch(opts *bind.TransactOpts, _cipherBatchHash [32]byte, _transactions [][]byte, _keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "executeCipherBatch", _cipherBatchHash, _transactions, _keyperIndex)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0xce0c1596.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, uint64 _keyperIndex) returns()
func (_ExecutorContract *ExecutorContractSession) ExecuteCipherBatch(_cipherBatchHash [32]byte, _transactions [][]byte, _keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, _cipherBatchHash, _transactions, _keyperIndex)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0xce0c1596.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, uint64 _keyperIndex) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) ExecuteCipherBatch(_cipherBatchHash [32]byte, _transactions [][]byte, _keyperIndex uint64) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, _cipherBatchHash, _transactions, _keyperIndex)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0xd57a29d0.
//
// Solidity: function executePlainBatch(bytes[] _transactions) returns()
func (_ExecutorContract *ExecutorContractTransactor) ExecutePlainBatch(opts *bind.TransactOpts, _transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "executePlainBatch", _transactions)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0xd57a29d0.
//
// Solidity: function executePlainBatch(bytes[] _transactions) returns()
func (_ExecutorContract *ExecutorContractSession) ExecutePlainBatch(_transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutePlainBatch(&_ExecutorContract.TransactOpts, _transactions)
}

// ExecutePlainBatch is a paid mutator transaction binding the contract method 0xd57a29d0.
//
// Solidity: function executePlainBatch(bytes[] _transactions) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) ExecutePlainBatch(_transactions [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecutePlainBatch(&_ExecutorContract.TransactOpts, _transactions)
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0x8f6dccfb.
//
// Solidity: function skipCipherExecution() returns()
func (_ExecutorContract *ExecutorContractTransactor) SkipCipherExecution(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "skipCipherExecution")
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0x8f6dccfb.
//
// Solidity: function skipCipherExecution() returns()
func (_ExecutorContract *ExecutorContractSession) SkipCipherExecution() (*types.Transaction, error) {
	return _ExecutorContract.Contract.SkipCipherExecution(&_ExecutorContract.TransactOpts)
}

// SkipCipherExecution is a paid mutator transaction binding the contract method 0x8f6dccfb.
//
// Solidity: function skipCipherExecution() returns()
func (_ExecutorContract *ExecutorContractTransactorSession) SkipCipherExecution() (*types.Transaction, error) {
	return _ExecutorContract.Contract.SkipCipherExecution(&_ExecutorContract.TransactOpts)
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

// FeeBankContractABI is the input ABI used to generate the binding from.
const FeeBankContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeBankContractBin is the compiled bytecode used for deploying new contracts.
var FeeBankContractBin = "0x608060405234801561001057600080fd5b506105c0806100206000396000f3fe60806040526004361061003f5760003560e01c80633ccfd60b14610044578063d6dad0601461005b578063f340fa011461009e578063fc7e286d146100c4575b600080fd5b34801561005057600080fd5b50610059610114565b005b34801561006757600080fd5b506100596004803603604081101561007e57600080fd5b5080356001600160a01b0316906020013567ffffffffffffffff1661013a565b610059600480360360208110156100b457600080fd5b50356001600160a01b0316610148565b3480156100d057600080fd5b506100f7600480360360208110156100e757600080fd5b50356001600160a01b03166102ca565b6040805167ffffffffffffffff9092168252519081900360200190f35b33600081815260208190526040902054610138919067ffffffffffffffff166102e6565b565b61014482826102e6565b5050565b6001600160a01b03811661018d5760405162461bcd60e51b81526004018080602001828103825260218152602001806105466021913960400191505060405180910390fd5b600034116101d9576040805162461bcd60e51b815260206004820152601460248201527346656542616e6b3a20666565206973207a65726f60601b604482015290519081900360640190fd5b6001600160a01b03811660009081526020819052604090205467ffffffffffffffff90811681031634111561023f5760405162461bcd60e51b81526004018080602001828103825260248152602001806105676024913960400191505060405180910390fd5b6001600160a01b03811660008181526020818152604091829020805467ffffffffffffffff80821634908101821667ffffffffffffffff1990931692909217928390558451338152938401959095528416828401529092166060830152517fc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d99181900360800190a150565b60006020819052908152604090205467ffffffffffffffff1681565b6001600160a01b03821661032b5760405162461bcd60e51b81526004018080602001828103825260218152602001806105466021913960400191505060405180910390fd5b3360009081526020819052604090205467ffffffffffffffff1680610397576040805162461bcd60e51b815260206004820152601960248201527f46656542616e6b3a206465706f73697420697320656d70747900000000000000604482015290519081900360640190fd5b8067ffffffffffffffff168267ffffffffffffffff161115610400576040805162461bcd60e51b815260206004820152601f60248201527f46656542616e6b3a20616d6f756e742065786365656473206465706f73697400604482015290519081900360640190fd5b33600090815260208190526040808220805467ffffffffffffffff191685850367ffffffffffffffff9081169190911790915590516001600160a01b038616918516908381818185875af1925050503d806000811461047b576040519150601f19603f3d011682016040523d82523d6000602084013e610480565b606091505b50509050806104d6576040805162461bcd60e51b815260206004820152601f60248201527f46656542616e6b3a207769746864726177616c2063616c6c206661696c656400604482015290519081900360640190fd5b33600081815260208181526040918290205482519384526001600160a01b0388169184019190915267ffffffffffffffff80871684840152166060830152517f4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e9181900360800190a15050505056fe46656542616e6b3a207265636569766572206973207a65726f206164647265737346656542616e6b3a2062616c616e636520776f756c64206578636565642075696e743634a26469706673582212204eaa06f169d245bc6c9dd1041d8203fbfb08409fd16fdec9d725ccb7bb884e1264736f6c63430007010033"

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
// Solidity: function deposit(address _receiver) payable returns()
func (_FeeBankContract *FeeBankContractTransactor) Deposit(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.contract.Transact(opts, "deposit", _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address _receiver) payable returns()
func (_FeeBankContract *FeeBankContractSession) Deposit(_receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Deposit(&_FeeBankContract.TransactOpts, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address _receiver) payable returns()
func (_FeeBankContract *FeeBankContractTransactorSession) Deposit(_receiver common.Address) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Deposit(&_FeeBankContract.TransactOpts, _receiver)
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
// Solidity: function withdraw(address _receiver, uint64 _amount) returns()
func (_FeeBankContract *FeeBankContractTransactor) Withdraw0(opts *bind.TransactOpts, _receiver common.Address, _amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.contract.Transact(opts, "withdraw0", _receiver, _amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd6dad060.
//
// Solidity: function withdraw(address _receiver, uint64 _amount) returns()
func (_FeeBankContract *FeeBankContractSession) Withdraw0(_receiver common.Address, _amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw0(&_FeeBankContract.TransactOpts, _receiver, _amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd6dad060.
//
// Solidity: function withdraw(address _receiver, uint64 _amount) returns()
func (_FeeBankContract *FeeBankContractTransactorSession) Withdraw0(_receiver common.Address, _amount uint64) (*types.Transaction, error) {
	return _FeeBankContract.Contract.Withdraw0(&_FeeBankContract.TransactOpts, _receiver, _amount)
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
const KeyBroadcastContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_configContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encryptionKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"signerIndices\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"EncryptionKeyBroadcasted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_encryptionKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"_signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"broadcastEncryptionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KeyBroadcastContractBin is the compiled bytecode used for deploying new contracts.
var KeyBroadcastContractBin = "0x608060405234801561001057600080fd5b5060405161082d38038061082d83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b61079c806100916000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632712860b1461003b578063bf66a18214610050575b600080fd5b61004e61004936600461049a565b61006e565b005b6100586101d2565b604051610065919061065c565b60405180910390f35b6100766101e1565b60005460405163700465b160e11b81526001600160a01b039091169063e008cb62906100a6908a90600401610711565b60006040518083038186803b1580156100be57600080fd5b505afa1580156100d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100fa9190810190610357565b90508060400151518867ffffffffffffffff16106101335760405162461bcd60e51b815260040161012a906106c2565b60405180910390fd5b80604001518867ffffffffffffffff168151811061014d57fe5b60200260200101516001600160a01b0316336001600160a01b0316146101855760405162461bcd60e51b815260040161012a90610670565b7f4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a338888888888886040516101c09796959493929190610560565b60405180910390a15050505050505050565b6000546001600160a01b031681565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b80516001600160a01b038116811461025d57600080fd5b92915050565b600082601f830112610273578081fd5b815167ffffffffffffffff811115610289578182fd5b6020808202610299828201610726565b838152935081840185830182870184018810156102b557600080fd5b600092505b848310156102e0576102cc8882610246565b8252600192909201919083019083016102ba565b505050505092915050565b60008083601f8401126102fc578081fd5b50813567ffffffffffffffff811115610313578182fd5b602083019150836020808302850101111561032d57600080fd5b9250929050565b80516001600160e01b03198116811461025d57600080fd5b805161025d8161074d565b600060208284031215610368578081fd5b815167ffffffffffffffff8082111561037f578283fd5b8184019150610180808387031215610395578384fd5b61039e81610726565b90506103aa868461034c565b81526103b9866020850161034c565b60208201526040830151828111156103cf578485fd5b6103db87828601610263565b6040830152506103ee866060850161034c565b6060820152610400866080850161034c565b60808201526104128660a0850161034c565b60a08201526104248660c0850161034c565b60c08201526104368660e0850161034c565b60e0820152610100915061044c86838501610246565b82820152610120915061046186838501610246565b82820152610140915061047686838501610334565b82820152610160915061048b8683850161034c565b91810191909152949350505050565b600080600080600080600060a0888a0312156104b4578283fd5b87356104bf8161074d565b965060208801356104cf8161074d565b955060408801359450606088013567ffffffffffffffff808211156104f2578485fd5b6104fe8b838c016102eb565b909650945060808a0135915080821115610516578384fd5b506105238a828b016102eb565b989b979a50959850939692959293505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b6001600160a01b038816815267ffffffffffffffff8781166020808401919091526040830188905260a060608401819052830186905260009187919060c08501845b898110156105c95784356105b58161074d565b8316825293830193908301906001016105a2565b50858103608087015286815282810193508287028101830188865b8981101561064857838303601f190187528135368c9003601e19018112610609578889fd5b8b0180358681111561061957898afd5b8036038d131561062757898afd5b61063485828a8501610536565b9888019894505050908501906001016105e4565b50909e9d5050505050505050505050505050565b6001600160a01b0391909116815260200190565b60208082526032908201527f4b657942726f616463617374436f6e74726163743a2073656e64657220646f6560408201527139903737ba1036b0ba31b41035b2bcb832b960711b606082015260800190565b6020808252602f908201527f4b657942726f616463617374436f6e74726163743a206b657970657220696e6460408201526e6578206f7574206f662072616e676560881b606082015260800190565b67ffffffffffffffff91909116815260200190565b60405181810167ffffffffffffffff8111828210171561074557600080fd5b604052919050565b67ffffffffffffffff8116811461076357600080fd5b5056fea26469706673582212207c70dc14a0e2ce4260084b5e40d284c135498ebcae45c4773fa30b11c2794a8564736f6c63430007010033"

// DeployKeyBroadcastContract deploys a new Ethereum contract, binding an instance of KeyBroadcastContract to it.
func DeployKeyBroadcastContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configContractAddress common.Address) (common.Address, *types.Transaction, *KeyBroadcastContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyBroadcastContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyBroadcastContractBin), backend, _configContractAddress)
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

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x2712860b.
//
// Solidity: function broadcastEncryptionKey(uint64 _keyperIndex, uint64 _batchIndex, bytes32 _encryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactor) BroadcastEncryptionKey(opts *bind.TransactOpts, _keyperIndex uint64, _batchIndex uint64, _encryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.contract.Transact(opts, "broadcastEncryptionKey", _keyperIndex, _batchIndex, _encryptionKey, _signerIndices, _signatures)
}

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x2712860b.
//
// Solidity: function broadcastEncryptionKey(uint64 _keyperIndex, uint64 _batchIndex, bytes32 _encryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractSession) BroadcastEncryptionKey(_keyperIndex uint64, _batchIndex uint64, _encryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.BroadcastEncryptionKey(&_KeyBroadcastContract.TransactOpts, _keyperIndex, _batchIndex, _encryptionKey, _signerIndices, _signatures)
}

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x2712860b.
//
// Solidity: function broadcastEncryptionKey(uint64 _keyperIndex, uint64 _batchIndex, bytes32 _encryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactorSession) BroadcastEncryptionKey(_keyperIndex uint64, _batchIndex uint64, _encryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.BroadcastEncryptionKey(&_KeyBroadcastContract.TransactOpts, _keyperIndex, _batchIndex, _encryptionKey, _signerIndices, _signatures)
}

// KeyBroadcastContractEncryptionKeyBroadcastedIterator is returned from FilterEncryptionKeyBroadcasted and is used to iterate over the raw logs and unpacked data for EncryptionKeyBroadcasted events raised by the KeyBroadcastContract contract.
type KeyBroadcastContractEncryptionKeyBroadcastedIterator struct {
	Event *KeyBroadcastContractEncryptionKeyBroadcasted // Event containing the contract specifics and raw log

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
func (it *KeyBroadcastContractEncryptionKeyBroadcastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyBroadcastContractEncryptionKeyBroadcasted)
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
		it.Event = new(KeyBroadcastContractEncryptionKeyBroadcasted)
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
func (it *KeyBroadcastContractEncryptionKeyBroadcastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyBroadcastContractEncryptionKeyBroadcastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyBroadcastContractEncryptionKeyBroadcasted represents a EncryptionKeyBroadcasted event raised by the KeyBroadcastContract contract.
type KeyBroadcastContractEncryptionKeyBroadcasted struct {
	Sender        common.Address
	BatchIndex    uint64
	EncryptionKey [32]byte
	SignerIndices []uint64
	Signatures    [][]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEncryptionKeyBroadcasted is a free log retrieval operation binding the contract event 0x4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint64 batchIndex, bytes32 encryptionKey, uint64[] signerIndices, bytes[] signatures)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) FilterEncryptionKeyBroadcasted(opts *bind.FilterOpts) (*KeyBroadcastContractEncryptionKeyBroadcastedIterator, error) {

	logs, sub, err := _KeyBroadcastContract.contract.FilterLogs(opts, "EncryptionKeyBroadcasted")
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractEncryptionKeyBroadcastedIterator{contract: _KeyBroadcastContract.contract, event: "EncryptionKeyBroadcasted", logs: logs, sub: sub}, nil
}

// WatchEncryptionKeyBroadcasted is a free log subscription operation binding the contract event 0x4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint64 batchIndex, bytes32 encryptionKey, uint64[] signerIndices, bytes[] signatures)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) WatchEncryptionKeyBroadcasted(opts *bind.WatchOpts, sink chan<- *KeyBroadcastContractEncryptionKeyBroadcasted) (event.Subscription, error) {

	logs, sub, err := _KeyBroadcastContract.contract.WatchLogs(opts, "EncryptionKeyBroadcasted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyBroadcastContractEncryptionKeyBroadcasted)
				if err := _KeyBroadcastContract.contract.UnpackLog(event, "EncryptionKeyBroadcasted", log); err != nil {
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

// ParseEncryptionKeyBroadcasted is a log parse operation binding the contract event 0x4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint64 batchIndex, bytes32 encryptionKey, uint64[] signerIndices, bytes[] signatures)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) ParseEncryptionKeyBroadcasted(log types.Log) (*KeyBroadcastContractEncryptionKeyBroadcasted, error) {
	event := new(KeyBroadcastContractEncryptionKeyBroadcasted)
	if err := _KeyBroadcastContract.contract.UnpackLog(event, "EncryptionKeyBroadcasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSlasherABI is the input ABI used to generate the binding from.
const KeyperSlasherABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_appealBlocks\",\"type\":\"uint256\"},{\"internalType\":\"contractConfigContract\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"contractExecutorContract\",\"name\":\"_executorContract\",\"type\":\"address\"},{\"internalType\":\"contractDepositContract\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"}],\"name\":\"Accused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Appealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"accusations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"accused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"appealed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"}],\"name\":\"accuse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structAuthorization\",\"name\":\"_authorization\",\"type\":\"tuple\"}],\"name\":\"appeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_halfStep\",\"type\":\"uint64\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeyperSlasherBin is the compiled bytecode used for deploying new contracts.
var KeyperSlasherBin = "0x60806040523480156200001157600080fd5b50604051620019f3380380620019f38339810160408190526200003491620000da565b6003849055600080546001600160a01b038086166001600160a01b03199283161790925560018054858416908316179055600280548484169216919091179081905560405163555e124b60e11b815291169063aabc2496906200009c90309060040162000135565b600060405180830381600087803b158015620000b757600080fd5b505af1158015620000cc573d6000803e3d6000fd5b505050505050505062000162565b60008060008060808587031215620000f0578384fd5b845193506020850151620001048162000149565b6040860151909350620001178162000149565b60608601519092506200012a8162000149565b939692955090935050565b6001600160a01b0391909116815260200190565b6001600160a01b03811681146200015f57600080fd5b50565b61188180620001726000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630e98ad4d1461005157806331217be1146100665780636864e7ee14610079578063ab4dfa8a146100a7575b600080fd5b61006461005f366004610f0a565b6100ba565b005b6100646100743660046111c2565b61031f565b61008c6100873660046111c2565b61050c565b60405161009e96959493929190611254565b60405180910390f35b6100646100b53660046111de565b610562565b6100c2610ce6565b5080516001600160401b03908116600090815260046020908152604091829020825160c081018452815460ff808216151580845261010083048216151595840195909552620100008204161515948201949094526001600160a01b0363010000008504166060820152600160b81b909304841660808401526001015490921660a08201529061016c5760405162461bcd60e51b815260040161016390611717565b60405180910390fd5b80602001511561018e5760405162461bcd60e51b81526004016101639061139f565b610196610d1b565b60015483516040516325b36cbf60e01b81526001600160a01b03909216916325b36cbf916101c691600401611796565b60a06040518083038186803b1580156101de57600080fd5b505afa1580156101f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102169190611153565b905061022283826108a6565b6001602083810182815285516001600160401b039081166000908152600484526040808220885181549551838b015160608c015160808d015160ff199099169315159390931761ff001916610100921515929092029190911762ff000019166201000091151591909102176301000000600160b81b03191663010000006001600160a01b03928316021767ffffffffffffffff60b81b1916600160b81b9686169690960295909517815560a08901519601805467ffffffffffffffff191696841696909617909555928501518751945192169316917f8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf91a3505050565b610327610ce6565b506001600160401b03808216600090815260046020908152604091829020825160c081018452815460ff808216151580845261010083048216151595840195909552620100008204161515948201949094526001600160a01b0363010000008504166060820152600160b81b909304841660808401526001015490921660a0820152906103c65760405162461bcd60e51b815260040161016390611717565b8060200151156103e85760405162461bcd60e51b81526004016101639061147b565b80604001511561040a5760405162461bcd60e51b8152600401610163906115c5565b6003548160a001516001600160401b03160143101561043b5760405162461bcd60e51b8152600401610163906112e8565b600254606082015160405163c96be4cb60e01b81526001600160a01b039092169163c96be4cb9161046e91600401611240565b600060405180830381600087803b15801561048857600080fd5b505af115801561049c573d6000803e3d6000fd5b5050506001600160401b03808416600090815260046020526040808220805462ff00001916620100001790556060850151608086015191516001600160a01b0390911694509216917fa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac59190a35050565b6004602052600090815260409020805460019091015460ff808316926101008104821692620100008204909216916001600160a01b036301000000830416916001600160401b03600160b81b9091048116911686565b60018216156105835760405162461bcd60e51b8152600401610163906116d2565b6001600160401b03821660009081526004602052604090205460ff16156105bc5760405162461bcd60e51b81526004016101639061158e565b6105c4610d49565b6000546001600160a01b031663e008cb6260026001600160401b038616046040518263ffffffff1660e01b81526004016105fe9190611796565b60006040518083038186803b15801561061657600080fd5b505afa15801561062a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106529190810190611011565b9050806040015151826001600160401b0316106106815760405162461bcd60e51b81526004016101639061174e565b8060400151826001600160401b03168151811061069a57fe5b60200260200101516001600160a01b0316336001600160a01b0316146106d25760405162461bcd60e51b815260040161016390611687565b6106da610d1b565b6001546040516325b36cbf60e01b81526001600160a01b03909116906325b36cbf9061070a908790600401611796565b60a06040518083038186803b15801561072257600080fd5b505afa158015610736573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075a9190611153565b805190915061077b5760405162461bcd60e51b81526004016101639061163e565b6040805160c08101825260018082526000602080840182815284860183815287830180516001600160a01b03908116606089019081526001600160401b038e811660808b0181815243831660a08d01908152828b5260049099528c8a209b518c54985197519451915160ff199099169015151761ff001916610100971515979097029690961762ff000019166201000093151593909302929092176301000000600160b81b0319166301000000958416959095029490941767ffffffffffffffff60b81b1916600160b81b95851695909502949094178855935196909501805467ffffffffffffffff1916969091169590951790945591519351339490921692917f79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f179190a450505050565b6108ae610d49565b60005460408301516001600160a01b039091169063e008cb62906002906001600160401b0316046040518263ffffffff1660e01b81526004016108f19190611796565b60006040518083038186803b15801561090957600080fd5b505afa15801561091d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109459190810190611011565b905080606001516001600160401b031683606001515110156109795760405162461bcd60e51b815260040161016390611508565b826040015151836060015151146109a25760405162461bcd60e51b81526004016101639061141e565b60015460408051635f59da8760e11b815290516000926001600160a01b03169163beb3b50e916004808301926020929190829003018186803b1580156109e757600080fd5b505afa1580156109fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1f9190610ee7565b83606001518460800151604051602001610a3b93929190611216565b60405160208183030381529060405280519060200120905060005b846060015151816001600160401b03161015610bb35760608560600151826001600160401b031681518110610a8757fe5b6020026020010151905060008660400151836001600160401b031681518110610aac57fe5b60200260200101519050846040015151816001600160401b031610610ae35760405162461bcd60e51b8152600401610163906113d6565b6001600160401b0383161580610b2d57508660400151600184036001600160401b031681518110610b1057fe5b60200260200101516001600160401b0316816001600160401b0316115b610b495760405162461bcd60e51b8152600401610163906114bf565b6000610b558584610bba565b90508560400151826001600160401b031681518110610b7057fe5b60200260200101516001600160a01b0316816001600160a01b031614610ba85760405162461bcd60e51b815260040161016390611331565b505050600101610a56565b5050505050565b60008151604114610bdd5760405162461bcd60e51b815260040161016390611368565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610c2f5760405162461bcd60e51b81526004016101639061154c565b8060ff16601b1480610c4457508060ff16601c145b610c605760405162461bcd60e51b8152600401610163906115fc565b600060018783868660405160008152602001604052604051610c859493929190611293565b6020604051602081039080840390855afa158015610ca7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610cda5760405162461bcd60e51b8152600401610163906112b1565b93505050505b92915050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080516101808101825260008082526020820181905260609282018390529181018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b8051610ce08161181e565b600082601f830112610dc9578081fd5b8151610ddc610dd7826117d0565b6117aa565b818152915060208083019084810181840286018201871015610dfd57600080fd5b60005b84811015610e25578151610e138161181e565b84529282019290820190600101610e00565b505050505092915050565b600082601f830112610e40578081fd5b8135610e4e610dd7826117d0565b818152915060208083019084810160005b84811015610e25578135870188603f820112610e7a57600080fd5b83810135610e8a610dd7826117ef565b81815260408b81848601011115610ea057600080fd5b610eaf83888401838701611812565b50865250509282019290820190600101610e5f565b80516001600160e01b031981168114610ce057600080fd5b8051610ce081611836565b600060208284031215610ef8578081fd5b8151610f038161181e565b9392505050565b60006020808385031215610f1c578182fd5b82356001600160401b0380821115610f32578384fd5b9084019060808287031215610f45578384fd5b610f4f60806117aa565b8235610f5a81611836565b81528284013584820152604083013582811115610f75578586fd5b8301601f81018813610f85578586fd5b8035610f93610dd7826117d0565b81815286810190838801888402850189018c1015610faf57898afd5b8994505b83851015610fda578035610fc681611836565b835260019490940193918801918801610fb3565b5060408501525050506060830135935081841115610ff6578485fd5b61100287858501610e30565b60608201529695505050505050565b600060208284031215611022578081fd5b81516001600160401b0380821115611038578283fd5b818401915061018080838703121561104e578384fd5b611057816117aa565b90506110638684610edc565b81526110728660208501610edc565b6020820152604083015182811115611088578485fd5b61109487828601610db9565b6040830152506110a78660608501610edc565b60608201526110b98660808501610edc565b60808201526110cb8660a08501610edc565b60a08201526110dd8660c08501610edc565b60c08201526110ef8660e08501610edc565b60e0820152610100915061110586838501610dae565b82820152610120915061111a86838501610dae565b82820152610140915061112f86838501610ec4565b82820152610160915061114486838501610edc565b91810191909152949350505050565b600060a08284031215611164578081fd5b61116e60a06117aa565b8251801515811461117d578283fd5b8152602083015161118d8161181e565b602082015260408301516111a081611836565b6040820152606083810151908201526080928301519281019290925250919050565b6000602082840312156111d3578081fd5b8135610f0381611836565b600080604083850312156111f0578081fd5b82356111fb81611836565b9150602083013561120b81611836565b809150509250929050565b60609390931b6bffffffffffffffffffffffff191683526014830191909152603482015260540190565b6001600160a01b0391909116815260200190565b9515158652931515602086015291151560408501526001600160a01b031660608401526001600160401b0390811660808401521660a082015260c00190565b93845260ff9290921660208401526040830152606082015260800190565b60208082526018908201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604082015260600190565b60208082526029908201527f4b6579706572536c61736865723a2061707065616c20706572696f64206e6f74604082015268081bdd995c881e595d60ba1b606082015260800190565b6020808252601b908201527f4b6579706572536c61736865723a2077726f6e67207369676e65720000000000604082015260600190565b6020808252601f908201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604082015260600190565b6020808252601f908201527f4b6579706572536c61736865723a20616c72656164792061707065616c656400604082015260600190565b60208082526028908201527f4b6579706572536c61736865723a207369676e657220696e646578206f7574206040820152676f662072616e676560c01b606082015260800190565b6020808252603e908201527f4b6579706572536c61736865723a206e756d626572206f66207369676e61747560408201527f72657320616e6420696e646963657320646f6573206e6f74206d617463680000606082015260800190565b60208082526024908201527f4b6579706572536c61736865723a207375636365737366756c6c792061707065604082015263185b195960e21b606082015260800190565b60208082526029908201527f4b6579706572536c61736865723a207369676e657220696e6469636573206e6f6040820152681d081bdc99195c995960ba1b606082015260800190565b60208082526024908201527f4b6579706572536c61736865723a206e6f7420656e6f756768207369676e61746040820152637572657360e01b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604082015261756560f01b606082015260800190565b6020808252601e908201527f4b6579706572536c61736865723a20616c726561647920616363757365640000604082015260600190565b6020808252601e908201527f4b6579706572536c61736865723a20616c726561647920736c61736865640000604082015260600190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604082015261756560f01b606082015260800190565b60208082526029908201527f4b6579706572536c61736865723a2068616c662073746570206e6f742079657460408201526808195e1958dd5d195960ba1b606082015260800190565b6020808252602b908201527f4b6579706572536c61736865723a2073656e64657220646f6573206e6f74206d60408201526a30ba31b41035b2bcb832b960a91b606082015260800190565b60208082526025908201527f4b6579706572536c61736865723a206e6f742061206369706865722068616c66604082015264020737465760dc1b606082015260800190565b6020808252601c908201527f4b6579706572536c61736865723a206e6f2061636375736174696f6e00000000604082015260600190565b60208082526028908201527f4b6579706572536c61736865723a206b657970657220696e646578206f7574206040820152676f662072616e676560c01b606082015260800190565b6001600160401b0391909116815260200190565b6040518181016001600160401b03811182821017156117c857600080fd5b604052919050565b60006001600160401b038211156117e5578081fd5b5060209081020190565b60006001600160401b03821115611804578081fd5b50601f01601f191660200190565b82818337506000910152565b6001600160a01b038116811461183357600080fd5b50565b6001600160401b038116811461183357600080fdfea2646970667358221220f9ede90b336c20e99470b5058e9af8add0b8f5e9be1107195cbce08d4f10fb9f64736f6c63430007010033"

// DeployKeyperSlasher deploys a new Ethereum contract, binding an instance of KeyperSlasher to it.
func DeployKeyperSlasher(auth *bind.TransactOpts, backend bind.ContractBackend, _appealBlocks *big.Int, _configContract common.Address, _executorContract common.Address, _depositContract common.Address) (common.Address, *types.Transaction, *KeyperSlasher, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyperSlasherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyperSlasherBin), backend, _appealBlocks, _configContract, _executorContract, _depositContract)
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

	outstruct.Accused = out[0].(bool)
	outstruct.Appealed = out[1].(bool)
	outstruct.Slashed = out[2].(bool)
	outstruct.Executor = out[3].(common.Address)
	outstruct.HalfStep = out[4].(uint64)
	outstruct.BlockNumber = out[5].(uint64)

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

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 _halfStep, uint64 _keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Accuse(opts *bind.TransactOpts, _halfStep uint64, _keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "accuse", _halfStep, _keyperIndex)
}

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 _halfStep, uint64 _keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherSession) Accuse(_halfStep uint64, _keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Accuse(&_KeyperSlasher.TransactOpts, _halfStep, _keyperIndex)
}

// Accuse is a paid mutator transaction binding the contract method 0xab4dfa8a.
//
// Solidity: function accuse(uint64 _halfStep, uint64 _keyperIndex) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Accuse(_halfStep uint64, _keyperIndex uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Accuse(&_KeyperSlasher.TransactOpts, _halfStep, _keyperIndex)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) _authorization) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Appeal(opts *bind.TransactOpts, _authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "appeal", _authorization)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) _authorization) returns()
func (_KeyperSlasher *KeyperSlasherSession) Appeal(_authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Appeal(&_KeyperSlasher.TransactOpts, _authorization)
}

// Appeal is a paid mutator transaction binding the contract method 0x0e98ad4d.
//
// Solidity: function appeal((uint64,bytes32,uint64[],bytes[]) _authorization) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Appeal(_authorization Authorization) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Appeal(&_KeyperSlasher.TransactOpts, _authorization)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 _halfStep) returns()
func (_KeyperSlasher *KeyperSlasherTransactor) Slash(opts *bind.TransactOpts, _halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.contract.Transact(opts, "slash", _halfStep)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 _halfStep) returns()
func (_KeyperSlasher *KeyperSlasherSession) Slash(_halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Slash(&_KeyperSlasher.TransactOpts, _halfStep)
}

// Slash is a paid mutator transaction binding the contract method 0x31217be1.
//
// Solidity: function slash(uint64 _halfStep) returns()
func (_KeyperSlasher *KeyperSlasherTransactorSession) Slash(_halfStep uint64) (*types.Transaction, error) {
	return _KeyperSlasher.Contract.Slash(&_KeyperSlasher.TransactOpts, _halfStep)
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
const MockBatcherContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"setBatchHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockBatcherContractBin is the compiled bytecode used for deploying new contracts.
var MockBatcherContractBin = "0x608060405234801561001057600080fd5b506101b8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ad15b6c51461003b578063c87afa8a14610050575b600080fd5b61004e61004936600461013c565b610079565b005b61006361005e366004610108565b6100c1565b6040516100709190610179565b60405180910390f35b67ffffffffffffffff8316600090815260208190526040812082918460018111156100a057fe5b60018111156100ab57fe5b8152602081019190915260400160002055505050565b600060208181529281526040808220909352908152205481565b8035600281106100ea57600080fd5b92915050565b803567ffffffffffffffff811681146100ea57600080fd5b6000806040838503121561011a578182fd5b61012484846100f0565b915061013384602085016100db565b90509250929050565b600080600060608486031215610150578081fd5b61015a85856100f0565b925061016985602086016100db565b9150604084013590509250925092565b9081526020019056fea2646970667358221220a503f3afe71890938e4a849999381ef29c71da487026bea76305175ad2f108fb64736f6c63430007010033"

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
// Solidity: function setBatchHash(uint64 _batchIndex, uint8 _type, bytes32 _hash) returns()
func (_MockBatcherContract *MockBatcherContractTransactor) SetBatchHash(opts *bind.TransactOpts, _batchIndex uint64, _type uint8, _hash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.contract.Transact(opts, "setBatchHash", _batchIndex, _type, _hash)
}

// SetBatchHash is a paid mutator transaction binding the contract method 0xad15b6c5.
//
// Solidity: function setBatchHash(uint64 _batchIndex, uint8 _type, bytes32 _hash) returns()
func (_MockBatcherContract *MockBatcherContractSession) SetBatchHash(_batchIndex uint64, _type uint8, _hash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.SetBatchHash(&_MockBatcherContract.TransactOpts, _batchIndex, _type, _hash)
}

// SetBatchHash is a paid mutator transaction binding the contract method 0xad15b6c5.
//
// Solidity: function setBatchHash(uint64 _batchIndex, uint8 _type, bytes32 _hash) returns()
func (_MockBatcherContract *MockBatcherContractTransactorSession) SetBatchHash(_batchIndex uint64, _type uint8, _hash [32]byte) (*types.Transaction, error) {
	return _MockBatcherContract.Contract.SetBatchHash(&_MockBatcherContract.TransactOpts, _batchIndex, _type, _hash)
}

// MockTargetContractABI is the input ABI used to generate the binding from.
const MockTargetContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"Called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockTargetContractBin is the compiled bytecode used for deploying new contracts.
var MockTargetContractBin = "0x608060405234801561001057600080fd5b5061014a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635a6535fc14610030575b600080fd5b6100a06004803603602081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b5090925090506100a2565b005b60005a90507fef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf9283838360405180806020018381526020018281038252858582818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a150505056fea264697066735822122000c0aa179637814bd455500e402bcb42b96b24bcb1fa4c22a2da410c1c59a34a64736f6c63430007010033"

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206bb2956d801c057be6442f61e44ca77632c3518f695383bf5c093f066f9ad03864736f6c63430007010033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TestDepositTokenContractABI is the input ABI used to generate the binding from.
const TestDepositTokenContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestDepositTokenContractBin is the compiled bytecode used for deploying new contracts.
var TestDepositTokenContractBin = "0x60806040523480156200001157600080fd5b5060408051808201825260038082526214d11560ea1b60208084018281528551808701875293845283820192909252845160008152908101909452825192939192620000609160029162000583565b5081516200007690600390602085019062000583565b5080516200008c90600490602084019062000608565b5060005b600454811015620000ec5760016005600060048481548110620000af57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560010162000090565b506040516329965a1d60e01b8152731820a4b7618bde71dce8cdc73aab6c95905fad24906329965a1d906200014a9030907fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce217705490829060040162000794565b600060405180830381600087803b1580156200016557600080fd5b505af11580156200017a573d6000803e3d6000fd5b50506040516329965a1d60e01b8152731820a4b7618bde71dce8cdc73aab6c95905fad2492506329965a1d9150620001db9030907faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a90829060040162000794565b600060405180830381600087803b158015620001f657600080fd5b505af11580156200020b573d6000803e3d6000fd5b505050505050506200024733620f424060405180602001604052806000815250604051806020016040528060008152506200024d60201b60201c565b620008d8565b6001600160a01b0384166200027f5760405162461bcd60e51b81526004016200027690620007ee565b60405180910390fd5b60006200028b620003c0565b90506200029c8160008787620003c4565b620002b884600154620003ca60201b62000a101790919060201c565b6001556001600160a01b03851660009081526020818152604090912054620002eb91869062000a10620003ca821b17901c565b6001600160a01b0386166000908152602081905260408120919091556200031a908290878787876001620003f9565b846001600160a01b0316816001600160a01b03167f2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d86868660405162000363939291906200089f565b60405180910390a3846001600160a01b031660006001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef86604051620003b1919062000896565b60405180910390a35050505050565b3390565b50505050565b600082820183811015620003f25760405162461bcd60e51b81526004016200027690620007b7565b9392505050565b60405163555ddc6560e11b8152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca90620004579089907fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b906004016200077b565b60206040518083038186803b1580156200047057600080fd5b505afa15801562000485573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004ab9190620006a6565b90506001600160a01b038116156200052d576040516223de2960e01b81526001600160a01b038216906223de2990620004f3908b908b908b908b908b908b906004016200071d565b600060405180830381600087803b1580156200050e57600080fd5b505af115801562000523573d6000803e3d6000fd5b5050505062000573565b8115620005735762000553866001600160a01b03166200057d60201b62000a351760201c565b15620005735760405162461bcd60e51b8152600401620002769062000823565b5050505050505050565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620005c657805160ff1916838001178555620005f6565b82800160010185558215620005f6579182015b82811115620005f6578251825591602001919060010190620005d9565b50620006049291506200066e565b5090565b82805482825590600052602060002090810192821562000660579160200282015b828111156200066057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000629565b506200060492915062000685565b5b808211156200060457600081556001016200066f565b5b80821115620006045780546001600160a01b031916815560010162000686565b600060208284031215620006b8578081fd5b81516001600160a01b0381168114620003f2578182fd5b60008151808452815b81811015620006f657602081850181015186830182015201620006d8565b81811115620007085782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c0608082018190526000906200075a90830185620006cf565b82810360a08401526200076e8185620006cf565b9998505050505050505050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0393841681526020810192909252909116604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252818101527f4552433737373a206d696e7420746f20746865207a65726f2061646472657373604082015260600190565b6020808252604d908201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460408201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60608201526c1ad95b9cd49958da5c1a595b9d609a1b608082015260a00190565b90815260200190565b600084825260606020830152620008ba6060830185620006cf565b8281036040840152620008ce8185620006cf565b9695505050505050565b61197d80620008e86000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b63711461021c578063dd62ed3e1461022f578063fad8b32a14610242578063fc673c4f14610255578063fe9d93031461026857610116565b8063959b8c3f146101db57806395d89b41146101ee5780639bd9bbc6146101f6578063a9059cbb1461020957610116565b806323b872dd116100e957806323b872dd14610183578063313ce56714610196578063556f0dc7146101ab57806362ad1b83146101b357806370a08231146101c857610116565b806306e485381461011b57806306fdde0314610139578063095ea7b31461014e57806318160ddd1461016e575b600080fd5b61012361027b565b60405161013091906114bf565b60405180910390f35b6101416102dd565b6040516101309190611517565b61016161015c3660046112bd565b610367565b604051610130919061150c565b610176610389565b6040516101309190611870565b6101616101913660046111ed565b61038f565b61019e6104d7565b60405161013091906118ae565b6101766104dc565b6101c66101c136600461122d565b6104e1565b005b6101766101d636600461117d565b610524565b6101c66101e936600461117d565b61053f565b61014161066c565b6101c66102043660046112e8565b6106cd565b6101616102173660046112bd565b6106f7565b61016161022a3660046111b5565b6107b1565b61017661023d3660046111b5565b610853565b6101c661025036600461117d565b61087e565b6101c661026336600461133f565b6109ab565b6101c66102763660046113bc565b6109ea565b606060048054806020026020016040519081016040528092919081815260200182805480156102d357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116102b5575b5050505050905090565b60028054604080516020601f60001961010060018716150201909416859004938401819004810282018101909252828152606093909290918301828280156102d35780601f1061033b576101008083540402835291602001916102d3565b820191906000526020600020905b81548152906001019060200180831161034957509395945050505050565b600080610372610a3b565b905061037f818585610a3f565b5060019392505050565b60015490565b60006001600160a01b0383166103c05760405162461bcd60e51b81526004016103b790611722565b60405180910390fd5b6001600160a01b0384166103e65760405162461bcd60e51b81526004016103b7906117b2565b60006103f0610a3b565b905061041e818686866040518060200160405280600081525060405180602001604052806000815250610af3565b61044a818686866040518060200160405280600081525060405180602001604052806000815250610c22565b61049e8582610499866040518060600160405280602981526020016118fc602991396001600160a01b03808c166000908152600860209081526040808320938b16835292905220549190610d4e565b610a3f565b6104cc8186868660405180602001604052806000815250604051806020016040528060008152506000610d7a565b506001949350505050565b601290565b600190565b6104f26104ec610a3b565b866107b1565b61050e5760405162461bcd60e51b81526004016103b790611766565b61051d85858585856001610ee1565b5050505050565b6001600160a01b031660009081526020819052604090205490565b806001600160a01b0316610551610a3b565b6001600160a01b031614156105785760405162461bcd60e51b81526004016103b79061162a565b6001600160a01b03811660009081526005602052604090205460ff16156105db57600760006105a5610a3b565b6001600160a01b03908116825260208083019390935260409182016000908120918516815292529020805460ff19169055610622565b6001600660006105e9610a3b565b6001600160a01b03908116825260208083019390935260409182016000908120918616815292529020805460ff19169115159190911790555b61062a610a3b565b6001600160a01b0316816001600160a01b03167ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f960405160405180910390a350565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102d35780601f1061033b576101008083540402835291602001916102d3565b6106f26106d8610a3b565b848484604051806020016040528060008152506001610ee1565b505050565b60006001600160a01b03831661071f5760405162461bcd60e51b81526004016103b790611722565b6000610729610a3b565b9050610757818286866040518060200160405280600081525060405180602001604052806000815250610af3565b610783818286866040518060200160405280600081525060405180602001604052806000815250610c22565b61037f8182868660405180602001604052806000815250604051806020016040528060008152506000610d7a565b6000816001600160a01b0316836001600160a01b0316148061081c57506001600160a01b03831660009081526005602052604090205460ff16801561081c57506001600160a01b0380831660009081526007602090815260408083209387168352929052205460ff16155b8061084c57506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff165b9392505050565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b610886610a3b565b6001600160a01b0316816001600160a01b031614156108b75760405162461bcd60e51b81526004016103b79061166e565b6001600160a01b03811660009081526005602052604090205460ff1615610923576001600760006108e6610a3b565b6001600160a01b03908116825260208083019390935260409182016000908120918616815292529020805460ff1916911515919091179055610961565b6006600061092f610a3b565b6001600160a01b03908116825260208083019390935260409182016000908120918516815292529020805460ff191690555b610969610a3b565b6001600160a01b0316816001600160a01b03167f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa160405160405180910390a350565b6109bc6109b6610a3b565b856107b1565b6109d85760405162461bcd60e51b81526004016103b790611766565b6109e484848484610f64565b50505050565b610a0c6109f5610a3b565b838360405180602001604052806000815250610f64565b5050565b60008282018381101561084c5760405162461bcd60e51b81526004016103b7906115b1565b3b151590565b3390565b6001600160a01b038316610a655760405162461bcd60e51b81526004016103b79061152a565b6001600160a01b038216610a8b5760405162461bcd60e51b81526004016103b79061182d565b6001600160a01b0380841660008181526008602090815260408083209487168084529490915290819020849055517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92590610ae6908590611870565b60405180910390a3505050565b60405163555ddc6560e11b8152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca90610b4f9089907f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe895906004016114a6565b60206040518083038186803b158015610b6757600080fd5b505afa158015610b7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9f9190611199565b90506001600160a01b03811615610c1957604051633ad5cbc160e11b81526001600160a01b038216906375ab978290610be6908a908a908a908a908a908a9060040161144c565b600060405180830381600087803b158015610c0057600080fd5b505af1158015610c14573d6000803e3d6000fd5b505050505b50505050505050565b610c2e868686866109e4565b610c6b836040518060600160405280602781526020016118d5602791396001600160a01b0388166000908152602081905260409020549190610d4e565b6001600160a01b038087166000908152602081905260408082209390935590861681522054610c9a9084610a10565b6001600160a01b0380861660008181526020819052604090819020939093559151878216918916907f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc8261467798790610cf390889088908890611879565b60405180910390a4836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610d3e9190611870565b60405180910390a3505050505050565b60008184841115610d725760405162461bcd60e51b81526004016103b79190611517565b505050900390565b60405163555ddc6560e11b8152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca90610dd69089907fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b906004016114a6565b60206040518083038186803b158015610dee57600080fd5b505afa158015610e02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e269190611199565b90506001600160a01b03811615610ea2576040516223de2960e01b81526001600160a01b038216906223de2990610e6b908b908b908b908b908b908b9060040161144c565b600060405180830381600087803b158015610e8557600080fd5b505af1158015610e99573d6000803e3d6000fd5b50505050610ed7565b8115610ed757610eba866001600160a01b0316610a35565b15610ed75760405162461bcd60e51b81526004016103b7906116af565b5050505050505050565b6001600160a01b038616610f075760405162461bcd60e51b81526004016103b79061156f565b6001600160a01b038516610f2d5760405162461bcd60e51b81526004016103b7906117f8565b6000610f37610a3b565b9050610f47818888888888610af3565b610f55818888888888610c22565b610c1981888888888888610d7a565b6001600160a01b038416610f8a5760405162461bcd60e51b81526004016103b7906115e8565b6000610f94610a3b565b9050610fa381866000876109e4565b610fb281866000878787610af3565b610fef84604051806060016040528060238152602001611925602391396001600160a01b0388166000908152602081905260409020549190610d4e565b6001600160a01b03861660009081526020819052604090205560015461101590856110bd565b600181905550846001600160a01b0316816001600160a01b03167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a409886868660405161106293929190611879565b60405180910390a360006001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040516110ae9190611870565b60405180910390a35050505050565b600061084c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610d4e565b600082601f83011261110f578081fd5b813567ffffffffffffffff80821115611126578283fd5b604051601f8301601f191681016020018281118282101715611146578485fd5b60405282815292508284830160200186101561116157600080fd5b8260208601602083013760006020848301015250505092915050565b60006020828403121561118e578081fd5b813561084c816118bc565b6000602082840312156111aa578081fd5b815161084c816118bc565b600080604083850312156111c7578081fd5b82356111d2816118bc565b915060208301356111e2816118bc565b809150509250929050565b600080600060608486031215611201578081fd5b833561120c816118bc565b9250602084013561121c816118bc565b929592945050506040919091013590565b600080600080600060a08688031215611244578081fd5b853561124f816118bc565b9450602086013561125f816118bc565b935060408601359250606086013567ffffffffffffffff80821115611282578283fd5b61128e89838a016110ff565b935060808801359150808211156112a3578283fd5b506112b0888289016110ff565b9150509295509295909350565b600080604083850312156112cf578182fd5b82356112da816118bc565b946020939093013593505050565b6000806000606084860312156112fc578283fd5b8335611307816118bc565b925060208401359150604084013567ffffffffffffffff811115611329578182fd5b611335868287016110ff565b9150509250925092565b60008060008060808587031215611354578384fd5b843561135f816118bc565b935060208501359250604085013567ffffffffffffffff80821115611382578384fd5b61138e888389016110ff565b935060608701359150808211156113a3578283fd5b506113b0878288016110ff565b91505092959194509250565b600080604083850312156113ce578182fd5b82359150602083013567ffffffffffffffff8111156113eb578182fd5b6113f7858286016110ff565b9150509250929050565b60008151808452815b818110156114265760208185018101518683018201520161140a565b818111156114375782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c06080820181905260009061148790830185611401565b82810360a08401526114998185611401565b9998505050505050505050565b6001600160a01b03929092168252602082015260400190565b6020808252825182820181905260009190848201906040850190845b818110156115005783516001600160a01b0316835292840192918401916001016114db565b50909695505050505050565b901515815260200190565b60006020825261084c6020830184611401565b60208082526025908201527f4552433737373a20617070726f76652066726f6d20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526022908201527f4552433737373a2073656e642066726f6d20746865207a65726f206164647265604082015261737360f01b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b60208082526022908201527f4552433737373a206275726e2066726f6d20746865207a65726f206164647265604082015261737360f01b606082015260800190565b60208082526024908201527f4552433737373a20617574686f72697a696e672073656c66206173206f70657260408201526330ba37b960e11b606082015260800190565b60208082526021908201527f4552433737373a207265766f6b696e672073656c66206173206f70657261746f6040820152603960f91b606082015260800190565b6020808252604d908201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460408201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60608201526c1ad95b9cd49958da5c1a595b9d609a1b608082015260a00190565b60208082526024908201527f4552433737373a207472616e7366657220746f20746865207a65726f206164646040820152637265737360e01b606082015260800190565b6020808252602c908201527f4552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f60408201526b39103337b9103437b63232b960a11b606082015260800190565b60208082526026908201527f4552433737373a207472616e736665722066726f6d20746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252818101527f4552433737373a2073656e6420746f20746865207a65726f2061646472657373604082015260600190565b60208082526023908201527f4552433737373a20617070726f766520746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b90815260200190565b6000848252606060208301526118926060830185611401565b82810360408401526118a48185611401565b9695505050505050565b60ff91909116815260200190565b6001600160a01b03811681146118d157600080fd5b5056fe4552433737373a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433737373a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654552433737373a206275726e20616d6f756e7420657863656564732062616c616e6365a26469706673582212206285c40a96abcd44729f9fd5f3a1047a292b8a4c1e8a703a37c8f6186588e19264736f6c63430007010033"

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
const TestTargetContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestTargetContractBin is the compiled bytecode used for deploying new contracts.
var TestTargetContractBin = "0x608060405234801561001057600080fd5b5060405161067138038061067183398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b6105e0806100916000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632d0335ab146100465780636c1032af1461006f578063943d720914610084575b600080fd5b6100596100543660046102b6565b610099565b604051610066919061051a565b60405180910390f35b6100776100be565b6040516100669190610410565b6100976100923660046102e4565b6100cd565b005b6001600160a01b031660009081526001602052604090205467ffffffffffffffff1690565b6000546001600160a01b031690565b6000546001600160a01b031633146101005760405162461bcd60e51b81526004016100f7906104cd565b60405180910390fd5b606060008060008480602001905181019061011b9190610358565b93509350935093506000848051906020012090506000600182868686604051600081526020016040526040516101549493929190610478565b6020604051602081039080840390855afa158015610176573d6000803e3d6000fd5b505050602060405103519050600060608780602001905181019061019a91906103bb565b6001600160a01b038516600090815260016020526040902054919350915067ffffffffffffffff8084169116146101e35760405162461bcd60e51b81526004016100f790610496565b6001600160a01b038316600090815260016020819052604091829020805467ffffffffffffffff191691850167ffffffffffffffff16919091179055517f5932d06f5fe39d60725112950a65a5591f5cb869ce6bc5d90fd188fb255491d49061025190859084908690610424565b60405180910390a1505050505050505050565b600082601f830112610274578081fd5b815161028761028282610556565b61052f565b915080825283602082850101111561029e57600080fd5b6102af81602084016020860161057a565b5092915050565b6000602082840312156102c7578081fd5b81356001600160a01b03811681146102dd578182fd5b9392505050565b6000602082840312156102f5578081fd5b813567ffffffffffffffff81111561030b578182fd5b8201601f8101841361031b578182fd5b803561032961028282610556565b81815285602083850101111561033d578384fd5b81602084016020830137908101602001929092525092915050565b6000806000806080858703121561036d578283fd5b845167ffffffffffffffff811115610383578384fd5b61038f87828801610264565b945050602085015160ff811681146103a5578384fd5b6040860151606090960151949790965092505050565b600080604083850312156103cd578182fd5b825167ffffffffffffffff80821682146103e5578384fd5b6020850151919350808211156103f9578283fd5b5061040685828601610264565b9150509250929050565b6001600160a01b0391909116815260200190565b600060018060a01b038516825260606020830152835180606084015261045181608085016020880161057a565b67ffffffffffffffff93909316604083015250601f91909101601f19160160800192915050565b93845260ff9290921660208401526040830152606082015260800190565b6020808252601f908201527f54657374546172676574436f6e74726163743a2077726f6e67206e6f6e636500604082015260600190565b6020808252602d908201527f54657374546172676574436f6e74726163743a206f6e6c79206578656375746f60408201526c722063616e206578656375746560981b606082015260800190565b67ffffffffffffffff91909116815260200190565b60405181810167ffffffffffffffff8111828210171561054e57600080fd5b604052919050565b600067ffffffffffffffff82111561056c578081fd5b50601f01601f191660200190565b60005b8381101561059557818101518382015260200161057d565b838111156105a4576000848401525b5050505056fea2646970667358221220507ff88611ec302b5680d8a72ac338aeefa125da0bcdb55cf6be9c9c5e0897f064736f6c63430007010033"

// DeployTestTargetContract deploys a new Ethereum contract, binding an instance of TestTargetContract to it.
func DeployTestTargetContract(auth *bind.TransactOpts, backend bind.ContractBackend, executor common.Address) (common.Address, *types.Transaction, *TestTargetContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTargetContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestTargetContractBin), backend, executor)
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

// GetExecutor is a free data retrieval call binding the contract method 0x6c1032af.
//
// Solidity: function getExecutor() view returns(address)
func (_TestTargetContract *TestTargetContractCaller) GetExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestTargetContract.contract.Call(opts, &out, "getExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutor is a free data retrieval call binding the contract method 0x6c1032af.
//
// Solidity: function getExecutor() view returns(address)
func (_TestTargetContract *TestTargetContractSession) GetExecutor() (common.Address, error) {
	return _TestTargetContract.Contract.GetExecutor(&_TestTargetContract.CallOpts)
}

// GetExecutor is a free data retrieval call binding the contract method 0x6c1032af.
//
// Solidity: function getExecutor() view returns(address)
func (_TestTargetContract *TestTargetContractCallerSession) GetExecutor() (common.Address, error) {
	return _TestTargetContract.Contract.GetExecutor(&_TestTargetContract.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64)
func (_TestTargetContract *TestTargetContractCaller) GetNonce(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _TestTargetContract.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64)
func (_TestTargetContract *TestTargetContractSession) GetNonce(account common.Address) (uint64, error) {
	return _TestTargetContract.Contract.GetNonce(&_TestTargetContract.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64)
func (_TestTargetContract *TestTargetContractCallerSession) GetNonce(account common.Address) (uint64, error) {
	return _TestTargetContract.Contract.GetNonce(&_TestTargetContract.CallOpts, account)
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
