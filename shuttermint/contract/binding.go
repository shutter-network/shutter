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

// BatcherContractABI is the input ABI used to generate the binding from.
const BatcherContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_configContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeBankAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTransactionType\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_transaction\",\"type\":\"bytes\"}],\"name\":\"addTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumTransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchSizes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBankContract\",\"outputs\":[{\"internalType\":\"contractFeeBankContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_minFee\",\"type\":\"uint64\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BatcherContractBin is the compiled bytecode used for deploying new contracts.
var BatcherContractBin = "0x60806040523480156200001157600080fd5b506040516200172638038062001726833981810160405281019062000037919062000190565b6000620000496200017160201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200021f565b600033905090565b6000815190506200018a8162000205565b92915050565b60008060408385031215620001a457600080fd5b6000620001b48582860162000179565b9250506020620001c78582860162000179565b9150509250929050565b6000620001de82620001e5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200021081620001d1565b81146200021c57600080fd5b50565b6114f7806200022f6000396000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b14610148578063bf66a18214610173578063bfd260ca1461019e578063c87afa8a146101db578063f2fde38b1461021857610091565b8063246673dc1461009657806324ec7590146100b257806336e1290d146100dd57806348fd5acc14610108578063715018a614610131575b600080fd5b6100b060048036038101906100ab9190610fb6565b610241565b005b3480156100be57600080fd5b506100c76106b1565b6040516100d49190611261565b60405180910390f35b3480156100e957600080fd5b506100f26106cb565b6040516100ff9190611206565b60405180910390f35b34801561011457600080fd5b5061012f600480360381019061012a9190610f51565b6106f1565b005b34801561013d57600080fd5b506101466107b2565b005b34801561015457600080fd5b5061015d610905565b60405161016a91906111b5565b60405180910390f35b34801561017f57600080fd5b5061018861092e565b60405161019591906111eb565b60405180910390f35b3480156101aa57600080fd5b506101c560048036038101906101c09190610f51565b610954565b6040516101d29190611261565b60405180910390f35b3480156101e757600080fd5b5061020260048036038101906101fd9190610f7a565b61097b565b60405161020f91906111d0565b60405180910390f35b34801561022457600080fd5b5061023f600480360381019061023a9190610ee7565b6109a0565b005b610249610b6a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62866040518263ffffffff1660e01b81526004016102a49190611261565b60006040518083038186803b1580156102bc57600080fd5b505afa1580156102d0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102f99190610f10565b90506000816080015167ffffffffffffffff161161031657600080fd5b806000015167ffffffffffffffff168567ffffffffffffffff16101561033857fe5b6000816000015186039050600082608001518202836020015101905060008360800151820190508167ffffffffffffffff1643101561037657600080fd5b8067ffffffffffffffff16431061038c57600080fd5b8585905060001061039c57600080fd5b8360c0015167ffffffffffffffff168686905011156103ba57600080fd5b8360a0015167ffffffffffffffff1686869050600360008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1601111561041f57600080fd5b600560009054906101000a900467ffffffffffffffff1667ffffffffffffffff1634101561044c57600080fd5b60608686600460008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008b600181111561048557fe5b600181111561049057fe5b8152602001908152602001600020546040516020016104b19392919061118b565b604051602081830303815290604052905060008180519060200120905080600460008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008b600181111561050457fe5b600181111561050f57fe5b81526020019081526020016000208190555087879050600360008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff160192506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000341180156105cd5750600073ffffffffffffffffffffffffffffffffffffffff1686610100015173ffffffffffffffffffffffffffffffffffffffff1614155b1561066657600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f340fa01348861010001516040518363ffffffff1660e01b815260040161063391906111b5565b6000604051808303818588803b15801561064c57600080fd5b505af1158015610660573d6000803e3d6000fd5b50505050505b7ffc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c988a8a8a8a8560405161069d95949392919061127c565b60405180910390a150505050505050505050565b600560009054906101000a900467ffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106f9610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d90611241565b60405180910390fd5b80600560006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6107ba610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610847576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083e90611241565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6004602052816000526040600020602052806000526040600020600091509150505481565b6109a8610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2c90611241565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611221565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b600081359050610c758161146c565b92915050565b600081519050610c8a8161146c565b92915050565b600082601f830112610ca157600080fd5b8151610cb4610caf826112f7565b6112ca565b91508181835260208401935060208101905083856020840282011115610cd957600080fd5b60005b83811015610d095781610cef8882610c7b565b845260208401935060208301925050600181019050610cdc565b5050505092915050565b600081519050610d2281611483565b92915050565b60008083601f840112610d3a57600080fd5b8235905067ffffffffffffffff811115610d5357600080fd5b602083019150836001820283011115610d6b57600080fd5b9250929050565b600081359050610d818161149a565b92915050565b60006101808284031215610d9a57600080fd5b610da56101806112ca565b90506000610db584828501610ed2565b6000830152506020610dc984828501610ed2565b602083015250604082015167ffffffffffffffff811115610de957600080fd5b610df584828501610c90565b6040830152506060610e0984828501610ed2565b6060830152506080610e1d84828501610ed2565b60808301525060a0610e3184828501610ed2565b60a08301525060c0610e4584828501610ed2565b60c08301525060e0610e5984828501610ed2565b60e083015250610100610e6e84828501610c7b565b61010083015250610120610e8484828501610c7b565b61012083015250610140610e9a84828501610d13565b61014083015250610160610eb084828501610ed2565b6101608301525092915050565b600081359050610ecc816114aa565b92915050565b600081519050610ee1816114aa565b92915050565b600060208284031215610ef957600080fd5b6000610f0784828501610c66565b91505092915050565b600060208284031215610f2257600080fd5b600082015167ffffffffffffffff811115610f3c57600080fd5b610f4884828501610d87565b91505092915050565b600060208284031215610f6357600080fd5b6000610f7184828501610ebd565b91505092915050565b60008060408385031215610f8d57600080fd5b6000610f9b85828601610ebd565b9250506020610fac85828601610d72565b9150509250929050565b60008060008060608587031215610fcc57600080fd5b6000610fda87828801610ebd565b9450506020610feb87828801610d72565b935050604085013567ffffffffffffffff81111561100857600080fd5b61101487828801610d28565b925092505092959194509250565b61102b8161134c565b82525050565b61103a8161135e565b82525050565b61105161104c8261135e565b611444565b82525050565b6000611063838561131f565b9350611070838584611435565b6110798361144e565b840190509392505050565b60006110908385611330565b935061109d838584611435565b82840190509392505050565b6110b2816113db565b82525050565b6110c1816113ff565b82525050565b6110d081611423565b82525050565b60006110e360268361133b565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061114960208361133b565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b611185816113c7565b82525050565b6000611198828587611084565b91506111a48284611040565b602082019150819050949350505050565b60006020820190506111ca6000830184611022565b92915050565b60006020820190506111e56000830184611031565b92915050565b600060208201905061120060008301846110a9565b92915050565b600060208201905061121b60008301846110b8565b92915050565b6000602082019050818103600083015261123a816110d6565b9050919050565b6000602082019050818103600083015261125a8161113c565b9050919050565b6000602082019050611276600083018461117c565b92915050565b6000608082019050611291600083018861117c565b61129e60208301876110c7565b81810360408301526112b1818587611057565b90506112c06060830184611031565b9695505050505050565b6000604051905081810181811067ffffffffffffffff821117156112ed57600080fd5b8060405250919050565b600067ffffffffffffffff82111561130e57600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611357826113a7565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b60008190506113a28261145f565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b60006113e6826113ed565b9050919050565b60006113f8826113a7565b9050919050565b600061140a82611411565b9050919050565b600061141c826113a7565b9050919050565b600061142e82611394565b9050919050565b82818337600083830152505050565b6000819050919050565b6000601f19601f8301169050919050565b6002811061146957fe5b50565b6114758161134c565b811461148057600080fd5b50565b61148c81611368565b811461149757600080fd5b50565b600281106114a757600080fd5b50565b6114b3816113c7565b81146114be57600080fd5b5056fea2646970667358221220c891094267257ece8f96ac52347219ccd49e3012958eff4ec5f8956783f2caf664736f6c63430007010033"

// DeployBatcherContract deploys a new Ethereum contract, binding an instance of BatcherContract to it.
func DeployBatcherContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configContractAddress common.Address, _feeBankAddress common.Address) (common.Address, *types.Transaction, *BatcherContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatcherContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatcherContractBin), backend, _configContractAddress, _feeBankAddress)
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
func (_BatcherContract *BatcherContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_BatcherContract *BatcherContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "batchHashes", arg0, arg1)
	return *ret0, err
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "batchSizes", arg0)
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "configContract")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "feeBankContract")
	return *ret0, err
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "minFee")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BatcherContract.contract.Call(opts, out, "owner")
	return *ret0, err
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
	return event, nil
}

// ConfigContractABI is the input ABI used to generate the binding from.
const ConfigContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configChangeHeadsUpBlocks\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numConfigs\",\"type\":\"uint64\"}],\"name\":\"ConfigUnscheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"configChangeHeadsUpBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"}],\"name\":\"configKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_configIndex\",\"type\":\"uint64\"}],\"name\":\"configNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"keypers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structBatchConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSpan\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionSizeLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"executionTimeout\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newKeypers\",\"type\":\"address[]\"}],\"name\":\"nextConfigAddKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"nextConfigKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigNumKeypers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nextConfigRemoveKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSpan\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetBatchSpan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_executionTimeout\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetExecutionTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"nextConfigSetFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startBatchIndex\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBatchIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startBlockNumber\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetStartBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_targetAddress\",\"type\":\"address\"}],\"name\":\"nextConfigSetTargetAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_targetFunctionSelector\",\"type\":\"bytes4\"}],\"name\":\"nextConfigSetTargetFunctionSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_threshold\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_transactionGasLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_transactionSizeLimit\",\"type\":\"uint64\"}],\"name\":\"nextConfigSetTransactionSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfigs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduleNextConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_fromStartBlockNumber\",\"type\":\"uint64\"}],\"name\":\"unscheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigContractBin is the compiled bytecode used for deploying new contracts.
var ConfigContractBin = "0x60a06040523480156200001157600080fd5b5060405162003ba538038062003ba58339818101604052810190620000379190620006d1565b6000620000496200039760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506001620000f96200039f60201b60201c565b908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010190805190602001906200019e929190620004f2565b5060608201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101208201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101408201518160040160146101000a81548163ffffffff021916908360e01c02179055506101608201518160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050508067ffffffffffffffff1660808167ffffffffffffffff1660c01b81525050506200072b565b600033905090565b620003a962000581565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff81118015620003f057600080fd5b50604051908082528060200260200182016040528015620004205781602001602082028036833780820191505090505b508152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff16815250905090565b8280548282559060005260206000209081019282156200056e579160200282015b828111156200056d5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000513565b5b5090506200057d91906200067d565b5090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b5b80821115620006b657600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016200067e565b5090565b600081519050620006cb8162000711565b92915050565b600060208284031215620006e457600080fd5b6000620006f484828501620006ba565b91505092915050565b600067ffffffffffffffff82169050919050565b6200071c81620006fd565b81146200072857600080fd5b50565b60805160c01c61344d6200075860003980610a7352806114085280611f685280612203525061344d6000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c806373ed43db116100f9578063cd21aee711610097578063d9a58f2411610071578063d9a58f2414610453578063e008cb6214610483578063f2fde38b146104b3578063fa84ea02146104cf576101a8565b8063cd21aee7146103fd578063ce9919b81461041b578063d1ac2e5214610437576101a8565b80639d63753e116100d35780639d63753e1461038d578063bcf67268146103a9578063c7c6e9f4146103c5578063c9515c58146103e1576101a8565b806373ed43db1461033757806381e905a3146103535780638da5cb5b1461036f576101a8565b80635dc6fdb81161016657806364e9f6711161014057806364e9f671146102b9578063660744dc146102e1578063715018a614610311578063719f2e171461031b576101a8565b80635dc6fdb814610265578063606df5141461028157806362fced0e1461029d576101a8565b806298fa22146101ad5780630f0aae6f146101e757806318b5e83014610205578063287447c41461020f5780632b2cc6c41461022d578063564093fc14610249575b600080fd5b6101c760048036038101906101c29190612eaa565b6104ff565b6040516101de9b9a9998979695949392919061324a565b60405180910390f35b6101ef610653565b6040516101fc919061322f565b60405180910390f35b61020d610660565b005b6102176111c1565b604051610224919061322f565b60405180910390f35b61024760048036038101906102429190612e13565b6111d1565b005b610263600480360381019061025e9190612ed3565b6112ad565b005b61027f600480360381019061027a9190612ed3565b611371565b005b61029b60048036038101906102969190612ed3565b611474565b005b6102b760048036038101906102b29190612e3c565b611537565b005b6102c16116aa565b6040516102d89b9a9998979695949392919061324a565b60405180910390f35b6102fb60048036038101906102f69190612ed3565b6117df565b60405161030891906131b2565b60405180910390f35b61031961182d565b005b61033560048036038101906103309190612ed3565b611980565b005b610351600480360381019061034c9190612ed3565b611a44565b005b61036d60048036038101906103689190612ed3565b611b07565b005b610377611bcb565b60405161038491906131b2565b60405180910390f35b6103a760048036038101906103a29190612ed3565b611bf4565b005b6103c360048036038101906103be9190612e13565b611d32565b005b6103df60048036038101906103da9190612ed3565b611e0e565b005b6103fb60048036038101906103f69190612ed3565b611ed1565b005b610405612201565b604051610412919061322f565b60405180910390f35b61043560048036038101906104309190612ed3565b612225565b005b610451600480360381019061044c9190612e81565b6122e9565b005b61046d60048036038101906104689190612ed3565b6123a2565b60405161047a919061322f565b60405180910390f35b61049d60048036038101906104989190612ed3565b6123d7565b6040516104aa919061320d565b60405180910390f35b6104cd60048036038101906104c89190612e13565b612799565b005b6104e960048036038101906104e49190612efc565b61295b565b6040516104f691906131b2565b60405180910390f35b6001818154811061050c57fe5b90600052602060002090600502016000915090508060000160009054906101000a900467ffffffffffffffff16908060000160089054906101000a900467ffffffffffffffff16908060020160009054906101000a900467ffffffffffffffff16908060020160089054906101000a900467ffffffffffffffff16908060020160109054906101000a900467ffffffffffffffff16908060020160189054906101000a900467ffffffffffffffff16908060030160009054906101000a900467ffffffffffffffff16908060030160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460e01b908060040160189054906101000a900467ffffffffffffffff1690508b565b6000600180549050905090565b6106686129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ec906131ed565b60405180910390fd5b600167ffffffffffffffff0367ffffffffffffffff166001805490501061071b57600080fd5b610723612b24565b6001808080549050038154811061073657fe5b9060005260206000209060050201604051806101800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801561083957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116107ef575b505050505081526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020016004820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff164301600260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610acc57600080fd5b6000816080015167ffffffffffffffff161115610b8f57806000015167ffffffffffffffff16600260000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610b2057600080fd5b60008160000151600260000160009054906101000a900467ffffffffffffffff16039050600260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168183608001510283602001510167ffffffffffffffff1614610b8957600080fd5b50610bcd565b806000015167ffffffffffffffff16600260000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610bcc57600080fd5b5b6001600290806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900467ffffffffffffffff168160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160089054906101000a900467ffffffffffffffff168160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820181600101908054610c96929190612c20565b506002820160009054906101000a900467ffffffffffffffff168160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160089054906101000a900467ffffffffffffffff168160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160109054906101000a900467ffffffffffffffff168160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160189054906101000a900467ffffffffffffffff168160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506003820160009054906101000a900467ffffffffffffffff168160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004820160149054906101000a900460e01b8160040160146101000a81548163ffffffff021916908360e01c02179055506004820160189054906101000a900467ffffffffffffffff168160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050610f326129d5565b600260008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151816001019080519060200190610fae929190612c72565b5060608201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101208201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101408201518160040160146101000a81548163ffffffff021916908360e01c02179055506101608201518160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050507f38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae66001805490506040516111b6919061322f565b60405180910390a150565b6000600260010180549050905090565b6111d96129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125d906131ed565b60405180910390fd5b80600260030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6112b56129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611342576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611339906131ed565b60405180910390fd5b80600260030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6113796129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fd906131ed565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168167ffffffffffffffff161061144657600080fd5b806002800160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b61147c6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611509576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611500906131ed565b60405180910390fd5b806002800160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b61153f6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146115cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c3906131ed565b60405180910390fd5b8181905067ffffffffffffffff80160360026001018054905011156115f057600080fd5b60005b828290508167ffffffffffffffff1610156116a557600260010183838367ffffffffffffffff1681811061162357fe5b90506020020160208101906116389190612e13565b9080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506115f3565b505050565b60028060000160009054906101000a900467ffffffffffffffff16908060000160089054906101000a900467ffffffffffffffff16908060020160009054906101000a900467ffffffffffffffff16908060020160089054906101000a900467ffffffffffffffff16908060020160109054906101000a900467ffffffffffffffff16908060020160189054906101000a900467ffffffffffffffff16908060030160009054906101000a900467ffffffffffffffff16908060030160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460e01b908060040160189054906101000a900467ffffffffffffffff1690508b565b600060026001018267ffffffffffffffff16815481106117fb57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6118356129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118b9906131ed565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6119886129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611a15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a0c906131ed565b60405180910390fd5b80600260040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611a4c6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611ad9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad0906131ed565b60405180910390fd5b806002800160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611b0f6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611b9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b93906131ed565b60405180910390fd5b80600260000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611bfc6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611c89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c80906131ed565b60405180910390fd5b60006002600101805490509050808267ffffffffffffffff1611611d1c5760005b8267ffffffffffffffff168167ffffffffffffffff161015611d16576002600101805480611cd457fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590558080600101915050611caa565b50611d2e565b60026001016000611d2d9190612cfc565b5b5050565b611d3a6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dbe906131ed565b60405180910390fd5b80600260040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611e166129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611ea3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e9a906131ed565b60405180910390fd5b806002800160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611ed96129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611f66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5d906131ed565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1643018167ffffffffffffffff1611611fa857600080fd5b600060018054905090506000600180805490500390505b60008111156121a557600060018281548110611fd757fe5b906000526020600020906005020190508367ffffffffffffffff168160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161061219057600180548061202557fe5b6001900381819060005260206000209060050201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556001820160006120809190612cfc565b6002820160006101000a81549067ffffffffffffffff02191690556002820160086101000a81549067ffffffffffffffff02191690556002820160106101000a81549067ffffffffffffffff02191690556002820160186101000a81549067ffffffffffffffff02191690556003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160146101000a81549063ffffffff02191690556004820160186101000a81549067ffffffffffffffff021916905550509055612196565b506121a5565b50808060019003915050611fbf565b508067ffffffffffffffff16600180549050106121c157600080fd5b7f202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf6001805490506040516121f5919061322f565b60405180910390a15050565b7f000000000000000000000000000000000000000000000000000000000000000081565b61222d6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146122ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b1906131ed565b60405180910390fd5b80600260000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6122f16129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461237e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612375906131ed565b60405180910390fd5b80600260040160146101000a81548163ffffffff021916908360e01c021790555050565b600060018267ffffffffffffffff16815481106123bb57fe5b9060005260206000209060050201600101805490509050919050565b6123df612b24565b6000600180805490500390505b6000811061278a5760006001828154811061240357fe5b906000526020600020906005020190508367ffffffffffffffff168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff161161277b5780604051806101800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801561253c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116124f2575b505050505081526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020016004820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505092505050612794565b508080600190039150506123ec565b50600061279357fe5b5b919050565b6127a16129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461282e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612825906131ed565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561289e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612895906131cd565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060018367ffffffffffffffff168154811061297457fe5b90600052602060002090600502016001018267ffffffffffffffff168154811061299a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b600033905090565b6129dd612b24565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff81118015612a2357600080fd5b50604051908082528060200260200182016040528015612a525781602001602082028036833780820191505090505b508152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff16815250905090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b828054828255906000526020600020908101928215612c615760005260206000209182015b82811115612c60578254825591600101919060010190612c45565b5b509050612c6e9190612d1d565b5090565b828054828255906000526020600020908101928215612ceb579160200282015b82811115612cea5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612c92565b5b509050612cf89190612d1d565b5090565b5080546000825590600052602060002090810190612d1a9190612d58565b50565b5b80821115612d5457600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612d1e565b5090565b5b80821115612d71576000816000905550600101612d59565b5090565b600081359050612d84816133bb565b92915050565b60008083601f840112612d9c57600080fd5b8235905067ffffffffffffffff811115612db557600080fd5b602083019150836020820283011115612dcd57600080fd5b9250929050565b600081359050612de3816133d2565b92915050565b600081359050612df8816133e9565b92915050565b600081359050612e0d81613400565b92915050565b600060208284031215612e2557600080fd5b6000612e3384828501612d75565b91505092915050565b60008060208385031215612e4f57600080fd5b600083013567ffffffffffffffff811115612e6957600080fd5b612e7585828601612d8a565b92509250509250929050565b600060208284031215612e9357600080fd5b6000612ea184828501612dd4565b91505092915050565b600060208284031215612ebc57600080fd5b6000612eca84828501612de9565b91505092915050565b600060208284031215612ee557600080fd5b6000612ef384828501612dfe565b91505092915050565b60008060408385031215612f0f57600080fd5b6000612f1d85828601612dfe565b9250506020612f2e85828601612dfe565b9150509250929050565b6000612f448383612f50565b60208301905092915050565b612f598161333f565b82525050565b612f688161333f565b82525050565b6000612f7982613305565b612f83818561331d565b9350612f8e836132f5565b8060005b83811015612fbf578151612fa68882612f38565b9750612fb183613310565b925050600181019050612f92565b5085935050505092915050565b612fd581613351565b82525050565b612fe481613351565b82525050565b6000612ff760268361332e565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061305d60208361332e565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000610180830160008301516130a96000860182613194565b5060208301516130bc6020860182613194565b50604083015184820360408601526130d48282612f6e565b91505060608301516130e96060860182613194565b5060808301516130fc6080860182613194565b5060a083015161310f60a0860182613194565b5060c083015161312260c0860182613194565b5060e083015161313560e0860182613194565b5061010083015161314a610100860182612f50565b5061012083015161315f610120860182612f50565b50610140830151613174610140860182612fcc565b50610160830151613189610160860182613194565b508091505092915050565b61319d816133a7565b82525050565b6131ac816133a7565b82525050565b60006020820190506131c76000830184612f5f565b92915050565b600060208201905081810360008301526131e681612fea565b9050919050565b6000602082019050818103600083015261320681613050565b9050919050565b600060208201905081810360008301526132278184613090565b905092915050565b600060208201905061324460008301846131a3565b92915050565b600061016082019050613260600083018e6131a3565b61326d602083018d6131a3565b61327a604083018c6131a3565b613287606083018b6131a3565b613294608083018a6131a3565b6132a160a08301896131a3565b6132ae60c08301886131a3565b6132bb60e0830187612f5f565b6132c9610100830186612f5f565b6132d7610120830185612fdb565b6132e56101408301846131a3565b9c9b505050505050505050505050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061334a8261337d565b9050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6133c48161333f565b81146133cf57600080fd5b50565b6133db81613351565b81146133e657600080fd5b50565b6133f28161339d565b81146133fd57600080fd5b50565b613409816133a7565b811461341457600080fd5b5056fea26469706673582212207d62ef11aeeb9777b34eaee476c98d93416591df5d1a0c6fc378fc81efc6fa3564736f6c63430007010033"

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
func (_ConfigContract *ConfigContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ConfigContract *ConfigContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configChangeHeadsUpBlocks")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configKeypers", _configIndex, _keyperIndex)
	return *ret0, err
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configNumKeypers", _configIndex)
	return *ret0, err
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
	ret := new(struct {
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
	out := ret
	err := _ConfigContract.contract.Call(opts, out, "configs", arg0)
	return *ret, err
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
	var (
		ret0 = new(BatchConfig)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "getConfig", _batchIndex)
	return *ret0, err
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
	ret := new(struct {
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
	out := ret
	err := _ConfigContract.contract.Call(opts, out, "nextConfig")
	return *ret, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "nextConfigKeypers", _index)
	return *ret0, err
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "nextConfigNumKeypers")
	return *ret0, err
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "numConfigs")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "owner")
	return *ret0, err
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
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201c49c6a8fe51f7322b89863b4fc8ccf5f10a624fc21c185ec5410bb01d83f35864736f6c63430007010033"

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
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// ExecutorContractABI is the input ABI used to generate the binding from.
const ExecutorContractABI = "[{\"inputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"contractBatcherContract\",\"name\":\"_batcherContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numExecutionHalfSteps\",\"type\":\"uint64\"}],\"name\":\"CipherExecutionSkipped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"batcherContract\",\"outputs\":[{\"internalType\":\"contractBatcherContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cipherBatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_decryptionKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"_signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"executeCipherBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_transactions\",\"type\":\"bytes[]\"}],\"name\":\"executePlainBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numExecutionHalfSteps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipCipherExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutorContractBin is the compiled bytecode used for deploying new contracts.
var ExecutorContractBin = "0x60806040523480156200001157600080fd5b5060405162001d8338038062001d838339818101604052810190620000379190620000ee565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001bf565b600081519050620000d1816200018b565b92915050565b600081519050620000e881620001a5565b92915050565b600080604083850312156200010257600080fd5b60006200011285828601620000d7565b92505060206200012585828601620000c0565b9150509250929050565b60006200013c826200016b565b9050919050565b600062000150826200012f565b9050919050565b600062000164826200012f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001968162000143565b8114620001a257600080fd5b50565b620001b08162000157565b8114620001bc57600080fd5b50565b611bb480620001cf6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806372654848146100675780638f6dccfb14610083578063beb3b50e1461008d578063bf66a182146100ab578063d57a29d0146100c9578063fa6385f4146100e5575b600080fd5b610081600480360381019061007c9190611305565b610103565b005b61008b610600565b005b610095610808565b6040516100a2919061178a565b60405180910390f35b6100b361082e565b6040516100c091906117a5565b60405180910390f35b6100e360048036038101906100de9190611297565b610852565b005b6100ed610b2a565b6040516100fa9190611840565b60405180910390f35b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161012f57fe5b0667ffffffffffffffff161461014457600080fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161017057fe5b04905061017b610ed0565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b81526004016101d49190611840565b60006040518083038186803b1580156101ec57600080fd5b505afa158015610200573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061022991906113d1565b90506000816080015167ffffffffffffffff161161024657600080fd5b6001820181608001510281602001510167ffffffffffffffff1643101561026c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c87afa8a8360006040518363ffffffff1660e01b81526004016102ca929190611884565b60206040518083038186803b1580156102e257600080fd5b505afa1580156102f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031a91906112dc565b8a1461032557600080fd5b60006103428261012001518361014001518460e001518d8d610b44565b9050816060015167ffffffffffffffff1685859050101561036257600080fd5b86869050858590501461037457600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c8a846040516020016103af9493929190611692565b60405160208183030381529060405280519060200120905060005b868690508167ffffffffffffffff16101561055b5736600088888467ffffffffffffffff168181106103f857fe5b905060200281019061040a91906118ad565b9150915060008b8b8567ffffffffffffffff1681811061042657fe5b905060200201602081019061043b9190611412565b905060008467ffffffffffffffff16148061049657508b8b6001860367ffffffffffffffff1681811061046a57fe5b905060200201602081019061047f9190611412565b67ffffffffffffffff168167ffffffffffffffff16115b61049f57600080fd5b60006104ef8685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610cdf565b905087604001518267ffffffffffffffff168151811061050b57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461054a57600080fd5b5050505080806001019150506103ca565b506001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a600160149054906101000a900467ffffffffffffffff16836040516105ea92919061185b565b60405180910390a1505050505050505050505050565b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161062c57fe5b0667ffffffffffffffff161461064157600080fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161066d57fe5b049050610678610ed0565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b81526004016106d19190611840565b60006040518083038186803b1580156106e957600080fd5b505afa1580156106fd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061072691906113d1565b90506000816080015167ffffffffffffffff161161074357600080fd5b806101600151600183018260800151028260200151010167ffffffffffffffff1643101561077057600080fd5b6001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507fa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be600160149054906101000a900467ffffffffffffffff166040516107fc9190611840565b60405180910390a15050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161087e57fe5b0667ffffffffffffffff161461089357600080fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16816108bf57fe5b0490506108ca610ed0565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b81526004016109239190611840565b60006040518083038186803b15801561093b57600080fd5b505afa15801561094f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061097891906113d1565b90506000816080015167ffffffffffffffff161161099257fe5b6001820181608001510281602001510167ffffffffffffffff164310156109b557fe5b60006109d28261012001518361014001518460e001518888610b44565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c87afa8a8460016040518363ffffffff1660e01b8152600401610a32929190611884565b60206040518083038186803b158015610a4a57600080fd5b505afa158015610a5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8291906112dc565b8114610a8d57600080fd5b6001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a600160149054906101000a900467ffffffffffffffff1682604051610b1b92919061185b565b60405180910390a15050505050565b600160149054906101000a900467ffffffffffffffff1681565b60008060005b848490508167ffffffffffffffff161015610cd15760608786868467ffffffffffffffff16818110610b7857fe5b9050602002810190610b8a91906118ad565b604051602401610b9b929190611766565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090508873ffffffffffffffffffffffffffffffffffffffff168767ffffffffffffffff1682604051610c2a919061170a565b60006040518083038160008787f1925050503d8060008114610c68576040519150601f19603f3d011682016040523d82523d6000602084013e610c6d565b606091505b50505085858367ffffffffffffffff16818110610c8657fe5b9050602002810190610c9891906118ad565b84604051602001610cab939291906116e0565b604051602081830303815290604052805190602001209250508080600101915050610b4a565b508091505095945050505050565b60006041825114610d25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1c906117e0565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610da8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9f90611800565b60405180910390fd5b601b8160ff1614158015610dc05750601c8160ff1614155b15610e00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df790611820565b60405180910390fd5b600060018783868660405160008152602001604052604051610e259493929190611721565b6020604051602081039080840390855afa158015610e47573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ec3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eba906117c0565b60405180910390fd5b8094505050505092915050565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b600081519050610fdb81611b22565b92915050565b600082601f830112610ff257600080fd5b815161100561100082611931565b611904565b9150818183526020840193506020810190508385602084028201111561102a57600080fd5b60005b8381101561105a57816110408882610fcc565b84526020840193506020830192505060018101905061102d565b5050505092915050565b60008083601f84011261107657600080fd5b8235905067ffffffffffffffff81111561108f57600080fd5b6020830191508360208202830111156110a757600080fd5b9250929050565b60008083601f8401126110c057600080fd5b8235905067ffffffffffffffff8111156110d957600080fd5b6020830191508360208202830111156110f157600080fd5b9250929050565b60008135905061110781611b39565b92915050565b60008151905061111c81611b39565b92915050565b60008151905061113181611b50565b92915050565b6000610180828403121561114a57600080fd5b611155610180611904565b9050600061116584828501611282565b600083015250602061117984828501611282565b602083015250604082015167ffffffffffffffff81111561119957600080fd5b6111a584828501610fe1565b60408301525060606111b984828501611282565b60608301525060806111cd84828501611282565b60808301525060a06111e184828501611282565b60a08301525060c06111f584828501611282565b60c08301525060e061120984828501611282565b60e08301525061010061121e84828501610fcc565b6101008301525061012061123484828501610fcc565b6101208301525061014061124a84828501611122565b6101408301525061016061126084828501611282565b6101608301525092915050565b60008135905061127c81611b67565b92915050565b60008151905061129181611b67565b92915050565b600080602083850312156112aa57600080fd5b600083013567ffffffffffffffff8111156112c457600080fd5b6112d085828601611064565b92509250509250929050565b6000602082840312156112ee57600080fd5b60006112fc8482850161110d565b91505092915050565b60008060008060008060008060a0898b03121561132157600080fd5b600061132f8b828c016110f8565b985050602089013567ffffffffffffffff81111561134c57600080fd5b6113588b828c01611064565b9750975050604061136b8b828c016110f8565b955050606089013567ffffffffffffffff81111561138857600080fd5b6113948b828c016110ae565b9450945050608089013567ffffffffffffffff8111156113b357600080fd5b6113bf8b828c01611064565b92509250509295985092959890939650565b6000602082840312156113e357600080fd5b600082015167ffffffffffffffff8111156113fd57600080fd5b61140984828501611137565b91505092915050565b60006020828403121561142457600080fd5b60006114328482850161126d565b91505092915050565b61144c61144782611991565b611ac9565b82525050565b61145b816119a3565b82525050565b61147261146d826119a3565b611adb565b82525050565b60006114848385611964565b9350611491838584611a87565b61149a83611af7565b840190509392505050565b60006114b18385611975565b93506114be838584611a87565b82840190509392505050565b60006114d582611959565b6114df8185611975565b93506114ef818560208601611a96565b80840191505092915050565b61150481611a2d565b82525050565b61151381611a51565b82525050565b61152281611a75565b82525050565b6000611535601883611980565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b6000611575601f83611980565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b60006115b5602283611980565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061161b602283611980565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b61167d81611a0c565b82525050565b61168c81611a20565b82525050565b600061169e828761143b565b6014820191506116ae8286611461565b6020820191506116be8285611461565b6020820191506116ce8284611461565b60208201915081905095945050505050565b60006116ed8285876114a5565b91506116f98284611461565b602082019150819050949350505050565b600061171682846114ca565b915081905092915050565b60006080820190506117366000830187611452565b6117436020830186611683565b6117506040830185611452565b61175d6060830184611452565b95945050505050565b60006020820190508181036000830152611781818486611478565b90509392505050565b600060208201905061179f60008301846114fb565b92915050565b60006020820190506117ba600083018461150a565b92915050565b600060208201905081810360008301526117d981611528565b9050919050565b600060208201905081810360008301526117f981611568565b9050919050565b60006020820190508181036000830152611819816115a8565b9050919050565b600060208201905081810360008301526118398161160e565b9050919050565b60006020820190506118556000830184611674565b92915050565b60006040820190506118706000830185611674565b61187d6020830184611452565b9392505050565b60006040820190506118996000830185611674565b6118a66020830184611519565b9392505050565b600080833560016020038436030381126118c657600080fd5b80840192508235915067ffffffffffffffff8211156118e457600080fd5b6020830192506001820236038313156118fc57600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff8211171561192757600080fd5b8060405250919050565b600067ffffffffffffffff82111561194857600080fd5b602082029050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061199c826119ec565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b60008190506119e782611b15565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b6000611a3882611a3f565b9050919050565b6000611a4a826119ec565b9050919050565b6000611a5c82611a63565b9050919050565b6000611a6e826119ec565b9050919050565b6000611a80826119d9565b9050919050565b82818337600083830152505050565b60005b83811015611ab4578082015181840152602081019050611a99565b83811115611ac3576000848401525b50505050565b6000611ad482611ae5565b9050919050565b6000819050919050565b6000611af082611b08565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b60028110611b1f57fe5b50565b611b2b81611991565b8114611b3657600080fd5b50565b611b42816119a3565b8114611b4d57600080fd5b50565b611b59816119ad565b8114611b6457600080fd5b50565b611b7081611a0c565b8114611b7b57600080fd5b5056fea2646970667358221220e31b91c54e5bd1e80599bee1c9809a97466dfc7858db27009c74f968eaa592b264736f6c63430007010033"

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
func (_ExecutorContract *ExecutorContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ExecutorContract *ExecutorContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExecutorContract.contract.Call(opts, out, "batcherContract")
	return *ret0, err
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

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_ExecutorContract *ExecutorContractCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExecutorContract.contract.Call(opts, out, "configContract")
	return *ret0, err
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

// NumExecutionHalfSteps is a free data retrieval call binding the contract method 0xfa6385f4.
//
// Solidity: function numExecutionHalfSteps() view returns(uint64)
func (_ExecutorContract *ExecutorContractCaller) NumExecutionHalfSteps(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ExecutorContract.contract.Call(opts, out, "numExecutionHalfSteps")
	return *ret0, err
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

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x72654848.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, bytes32 _decryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_ExecutorContract *ExecutorContractTransactor) ExecuteCipherBatch(opts *bind.TransactOpts, _cipherBatchHash [32]byte, _transactions [][]byte, _decryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.contract.Transact(opts, "executeCipherBatch", _cipherBatchHash, _transactions, _decryptionKey, _signerIndices, _signatures)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x72654848.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, bytes32 _decryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_ExecutorContract *ExecutorContractSession) ExecuteCipherBatch(_cipherBatchHash [32]byte, _transactions [][]byte, _decryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, _cipherBatchHash, _transactions, _decryptionKey, _signerIndices, _signatures)
}

// ExecuteCipherBatch is a paid mutator transaction binding the contract method 0x72654848.
//
// Solidity: function executeCipherBatch(bytes32 _cipherBatchHash, bytes[] _transactions, bytes32 _decryptionKey, uint64[] _signerIndices, bytes[] _signatures) returns()
func (_ExecutorContract *ExecutorContractTransactorSession) ExecuteCipherBatch(_cipherBatchHash [32]byte, _transactions [][]byte, _decryptionKey [32]byte, _signerIndices []uint64, _signatures [][]byte) (*types.Transaction, error) {
	return _ExecutorContract.Contract.ExecuteCipherBatch(&_ExecutorContract.TransactOpts, _cipherBatchHash, _transactions, _decryptionKey, _signerIndices, _signatures)
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
	return event, nil
}

// FeeBankContractABI is the input ABI used to generate the binding from.
const FeeBankContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeBankContractBin is the compiled bytecode used for deploying new contracts.
var FeeBankContractBin = "0x608060405234801561001057600080fd5b506106ec806100206000396000f3fe60806040526004361061003f5760003560e01c80633ccfd60b14610044578063d6dad0601461005b578063f340fa01146100c0578063fc7e286d14610104575b600080fd5b34801561005057600080fd5b50610059610173565b005b34801561006757600080fd5b506100be6004803603604081101561007e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803567ffffffffffffffff1690602001909291905050506101d1565b005b610102600480360360208110156100d657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101df565b005b34801561011057600080fd5b506101536004803603602081101561012757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103fb565b604051808267ffffffffffffffff16815260200191505060405180910390f35b6101cf336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16610422565b565b6101db8282610422565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561021957600080fd5b6000341161022657600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff0367ffffffffffffffff1634111561029957600080fd5b346000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff160192506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d93382346000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16604051808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815260200194505050505060405180910390a150565b60006020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561045c57600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16905060008167ffffffffffffffff16116104ca57600080fd5b8067ffffffffffffffff168267ffffffffffffffff1611156104eb57600080fd5b8181036000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060008373ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff1660405180600001905060006040518083038185875af1925050503d80600081146105bc576040519150601f19603f3d011682016040523d82523d6000602084013e6105c1565b606091505b50509050806105cf57600080fd5b7f4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e3385856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16604051808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815260200194505050505060405180910390a15050505056fea26469706673582212200be728e82c3e5e5a6a0ad9bf305da4d20ecb24d233993e1c9372c48c53fcd87f64736f6c63430007010033"

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
func (_FeeBankContract *FeeBankContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_FeeBankContract *FeeBankContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _FeeBankContract.contract.Call(opts, out, "deposits", arg0)
	return *ret0, err
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
	return event, nil
}

// KeyBroadcastContractABI is the input ABI used to generate the binding from.
const KeyBroadcastContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_configContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encryptionKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"signerIndices\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"EncryptionKeyBroadcasted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_encryptionKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"_signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"broadcastEncryptionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KeyBroadcastContractBin is the compiled bytecode used for deploying new contracts.
var KeyBroadcastContractBin = "0x608060405234801561001057600080fd5b50604051610c83380380610c838339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610b758061010e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632712860b1461003b578063bf66a18214610057575b600080fd5b61005560048036038101906100509190610605565b610075565b005b61005f6101ee565b60405161006c9190610893565b60405180910390f35b61007d610212565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62886040518263ffffffff1660e01b81526004016100d691906108ae565b60006040518083038186803b1580156100ee57600080fd5b505afa158015610102573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061012b91906105c4565b90508060400151518867ffffffffffffffff161061014857600080fd5b80604001518867ffffffffffffffff168151811061016257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a157600080fd5b7f4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a338888888888886040516101dc979695949392919061082e565b60405180910390a15050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b60008151905061031d81610ae3565b92915050565b600082601f83011261033457600080fd5b8151610347610342826108f6565b6108c9565b9150818183526020840193506020810190508385602084028201111561036c57600080fd5b60005b8381101561039c5781610382888261030e565b84526020840193506020830192505060018101905061036f565b5050505092915050565b60008083601f8401126103b857600080fd5b8235905067ffffffffffffffff8111156103d157600080fd5b6020830191508360208202830111156103e957600080fd5b9250929050565b60008083601f84011261040257600080fd5b8235905067ffffffffffffffff81111561041b57600080fd5b60208301915083602082028301111561043357600080fd5b9250929050565b60008135905061044981610afa565b92915050565b60008151905061045e81610b11565b92915050565b6000610180828403121561047757600080fd5b6104826101806108c9565b90506000610492848285016105af565b60008301525060206104a6848285016105af565b602083015250604082015167ffffffffffffffff8111156104c657600080fd5b6104d284828501610323565b60408301525060606104e6848285016105af565b60608301525060806104fa848285016105af565b60808301525060a061050e848285016105af565b60a08301525060c0610522848285016105af565b60c08301525060e0610536848285016105af565b60e08301525061010061054b8482850161030e565b610100830152506101206105618482850161030e565b610120830152506101406105778482850161044f565b6101408301525061016061058d848285016105af565b6101608301525092915050565b6000813590506105a981610b28565b92915050565b6000815190506105be81610b28565b92915050565b6000602082840312156105d657600080fd5b600082015167ffffffffffffffff8111156105f057600080fd5b6105fc84828501610464565b91505092915050565b600080600080600080600060a0888a03121561062057600080fd5b600061062e8a828b0161059a565b975050602061063f8a828b0161059a565b96505060406106508a828b0161043a565b955050606088013567ffffffffffffffff81111561066d57600080fd5b6106798a828b016103f0565b9450945050608088013567ffffffffffffffff81111561069857600080fd5b6106a48a828b016103a6565b925092505092959891949750929550565b60006106c28484846107d4565b90509392505050565b60006106d78383610810565b60208301905092915050565b6106ec81610a69565b82525050565b60006106fe838561094c565b9350836020840285016107108461091e565b8060005b8781101561075657848403895261072b828461097f565b6107368682846106b5565b955061074184610932565b935060208b019a505050600181019050610714565b50829750879450505050509392505050565b6000610774838561095d565b935061077f82610928565b8060005b858110156107b85761079582846109d6565b61079f88826106cb565b97506107aa8361093f565b925050600181019050610783565b5085925050509392505050565b6107ce816109ff565b82525050565b60006107e0838561096e565b93506107ed838584610ac3565b6107f683610ad2565b840190509392505050565b61080a81610a7b565b82525050565b61081981610a55565b82525050565b61082881610a55565b82525050565b600060a082019050610843600083018a6106e3565b610850602083018961081f565b61085d60408301886107c5565b8181036060830152610870818688610768565b905081810360808301526108858184866106f2565b905098975050505050505050565b60006020820190506108a86000830184610801565b92915050565b60006020820190506108c3600083018461081f565b92915050565b6000604051905081810181811067ffffffffffffffff821117156108ec57600080fd5b8060405250919050565b600067ffffffffffffffff82111561090d57600080fd5b602082029050602081019050919050565b6000819050919050565b6000819050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000808335600160200384360303811261099857600080fd5b83810192508235915060208301925067ffffffffffffffff8211156109bc57600080fd5b6001820236038413156109ce57600080fd5b509250929050565b60006109e5602084018461059a565b905092915050565b60006109f882610a35565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000610a7482610a9f565b9050919050565b6000610a8682610a8d565b9050919050565b6000610a9882610a35565b9050919050565b6000610aaa82610ab1565b9050919050565b6000610abc82610a35565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b610aec816109ed565b8114610af757600080fd5b50565b610b03816109ff565b8114610b0e57600080fd5b50565b610b1a81610a09565b8114610b2557600080fd5b50565b610b3181610a55565b8114610b3c57600080fd5b5056fea26469706673582212204d7738c1825a6cc6d97e6d7e596085480641c61261aa57f960296cfc36f3b86464736f6c63430007010033"

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
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_KeyBroadcastContract *KeyBroadcastContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KeyBroadcastContract.contract.Call(opts, out, "configContract")
	return *ret0, err
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
	return event, nil
}

// MockBatcherContractABI is the input ABI used to generate the binding from.
const MockBatcherContractABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"batchHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumMockBatcherContract.TransactionType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"setBatchHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockBatcherContractBin is the compiled bytecode used for deploying new contracts.
var MockBatcherContractBin = "0x608060405234801561001057600080fd5b50610289806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ad15b6c51461003b578063c87afa8a14610057575b600080fd5b6100556004803603810190610050919061017e565b610087565b005b610071600480360381019061006c9190610142565b6100de565b60405161007e91906101dc565b60405180910390f35b806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008460018111156100bc57fe5b60018111156100c757fe5b815260200190815260200160002081905550505050565b6000602052816000526040600020602052806000526040600020600091509150505481565b60008135905061011281610215565b92915050565b6000813590506101278161022c565b92915050565b60008135905061013c8161023c565b92915050565b6000806040838503121561015557600080fd5b60006101638582860161012d565b925050602061017485828601610118565b9150509250929050565b60008060006060848603121561019357600080fd5b60006101a18682870161012d565b93505060206101b286828701610118565b92505060406101c386828701610103565b9150509250925092565b6101d6816101f7565b82525050565b60006020820190506101f160008301846101cd565b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b61021e816101f7565b811461022957600080fd5b50565b6002811061023957600080fd5b50565b61024581610201565b811461025057600080fd5b5056fea26469706673582212202787d968a71da6c5f032c70ae65fccfff0c3731bdcea0713ad13b5c4c68efc9364736f6c63430007010033"

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
func (_MockBatcherContract *MockBatcherContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_MockBatcherContract *MockBatcherContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MockBatcherContract.contract.Call(opts, out, "batchHashes", arg0, arg1)
	return *ret0, err
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
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
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
	return event, nil
}
