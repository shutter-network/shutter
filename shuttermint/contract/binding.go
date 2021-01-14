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
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f28ef76f38333fbb6b43063cf7818b40469511362902ae449ff7465c7d70fb164736f6c63430007010033"

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
var BatcherContractBin = "0x60806040523480156200001157600080fd5b506040516200177f3803806200177f8339818101604052810190620000379190620001a7565b6000620000496200017160201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000278565b600033905090565b6000815190506200018a8162000244565b92915050565b600081519050620001a1816200025e565b92915050565b60008060408385031215620001bb57600080fd5b6000620001cb8582860162000179565b9250506020620001de8582860162000190565b9150509250929050565b6000620001f58262000224565b9050919050565b60006200020982620001e8565b9050919050565b60006200021d82620001e8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200024f81620001fc565b81146200025b57600080fd5b50565b620002698162000210565b81146200027557600080fd5b50565b6114f780620002886000396000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b14610148578063bf66a18214610173578063bfd260ca1461019e578063c87afa8a146101db578063f2fde38b1461021857610091565b8063246673dc1461009657806324ec7590146100b257806336e1290d146100dd57806348fd5acc14610108578063715018a614610131575b600080fd5b6100b060048036038101906100ab9190610fb6565b610241565b005b3480156100be57600080fd5b506100c76106b1565b6040516100d49190611261565b60405180910390f35b3480156100e957600080fd5b506100f26106cb565b6040516100ff9190611206565b60405180910390f35b34801561011457600080fd5b5061012f600480360381019061012a9190610f51565b6106f1565b005b34801561013d57600080fd5b506101466107b2565b005b34801561015457600080fd5b5061015d610905565b60405161016a91906111b5565b60405180910390f35b34801561017f57600080fd5b5061018861092e565b60405161019591906111eb565b60405180910390f35b3480156101aa57600080fd5b506101c560048036038101906101c09190610f51565b610954565b6040516101d29190611261565b60405180910390f35b3480156101e757600080fd5b5061020260048036038101906101fd9190610f7a565b61097b565b60405161020f91906111d0565b60405180910390f35b34801561022457600080fd5b5061023f600480360381019061023a9190610ee7565b6109a0565b005b610249610b6a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62866040518263ffffffff1660e01b81526004016102a49190611261565b60006040518083038186803b1580156102bc57600080fd5b505afa1580156102d0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102f99190610f10565b90506000816080015167ffffffffffffffff161161031657600080fd5b806000015167ffffffffffffffff168567ffffffffffffffff16101561033857fe5b6000816000015186039050600082608001518202836020015101905060008360800151820190508167ffffffffffffffff1643101561037657600080fd5b8067ffffffffffffffff16431061038c57600080fd5b8585905060001061039c57600080fd5b8360c0015167ffffffffffffffff168686905011156103ba57600080fd5b8360a0015167ffffffffffffffff1686869050600360008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1601111561041f57600080fd5b600560009054906101000a900467ffffffffffffffff1667ffffffffffffffff1634101561044c57600080fd5b60608686600460008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008b600181111561048557fe5b600181111561049057fe5b8152602001908152602001600020546040516020016104b19392919061118b565b604051602081830303815290604052905060008180519060200120905080600460008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008b600181111561050457fe5b600181111561050f57fe5b81526020019081526020016000208190555087879050600360008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff160192506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000341180156105cd5750600073ffffffffffffffffffffffffffffffffffffffff1686610100015173ffffffffffffffffffffffffffffffffffffffff1614155b1561066657600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f340fa01348861010001516040518363ffffffff1660e01b815260040161063391906111b5565b6000604051808303818588803b15801561064c57600080fd5b505af1158015610660573d6000803e3d6000fd5b50505050505b7ffc285e0b48a09e92ec4acb05226c557c0af1c3976d350d24b4fd4fa104f82c988a8a8a8a8560405161069d95949392919061127c565b60405180910390a150505050505050505050565b600560009054906101000a900467ffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106f9610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d90611241565b60405180910390fd5b80600560006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6107ba610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610847576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083e90611241565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6004602052816000526040600020602052806000526040600020600091509150505481565b6109a8610b62565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2c90611241565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611221565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b600081359050610c758161146c565b92915050565b600081519050610c8a8161146c565b92915050565b600082601f830112610ca157600080fd5b8151610cb4610caf826112f7565b6112ca565b91508181835260208401935060208101905083856020840282011115610cd957600080fd5b60005b83811015610d095781610cef8882610c7b565b845260208401935060208301925050600181019050610cdc565b5050505092915050565b600081519050610d2281611483565b92915050565b60008083601f840112610d3a57600080fd5b8235905067ffffffffffffffff811115610d5357600080fd5b602083019150836001820283011115610d6b57600080fd5b9250929050565b600081359050610d818161149a565b92915050565b60006101808284031215610d9a57600080fd5b610da56101806112ca565b90506000610db584828501610ed2565b6000830152506020610dc984828501610ed2565b602083015250604082015167ffffffffffffffff811115610de957600080fd5b610df584828501610c90565b6040830152506060610e0984828501610ed2565b6060830152506080610e1d84828501610ed2565b60808301525060a0610e3184828501610ed2565b60a08301525060c0610e4584828501610ed2565b60c08301525060e0610e5984828501610ed2565b60e083015250610100610e6e84828501610c7b565b61010083015250610120610e8484828501610c7b565b61012083015250610140610e9a84828501610d13565b61014083015250610160610eb084828501610ed2565b6101608301525092915050565b600081359050610ecc816114aa565b92915050565b600081519050610ee1816114aa565b92915050565b600060208284031215610ef957600080fd5b6000610f0784828501610c66565b91505092915050565b600060208284031215610f2257600080fd5b600082015167ffffffffffffffff811115610f3c57600080fd5b610f4884828501610d87565b91505092915050565b600060208284031215610f6357600080fd5b6000610f7184828501610ebd565b91505092915050565b60008060408385031215610f8d57600080fd5b6000610f9b85828601610ebd565b9250506020610fac85828601610d72565b9150509250929050565b60008060008060608587031215610fcc57600080fd5b6000610fda87828801610ebd565b9450506020610feb87828801610d72565b935050604085013567ffffffffffffffff81111561100857600080fd5b61101487828801610d28565b925092505092959194509250565b61102b8161134c565b82525050565b61103a8161135e565b82525050565b61105161104c8261135e565b611444565b82525050565b6000611063838561131f565b9350611070838584611435565b6110798361144e565b840190509392505050565b60006110908385611330565b935061109d838584611435565b82840190509392505050565b6110b2816113db565b82525050565b6110c1816113ff565b82525050565b6110d081611423565b82525050565b60006110e360268361133b565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061114960208361133b565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b611185816113c7565b82525050565b6000611198828587611084565b91506111a48284611040565b602082019150819050949350505050565b60006020820190506111ca6000830184611022565b92915050565b60006020820190506111e56000830184611031565b92915050565b600060208201905061120060008301846110a9565b92915050565b600060208201905061121b60008301846110b8565b92915050565b6000602082019050818103600083015261123a816110d6565b9050919050565b6000602082019050818103600083015261125a8161113c565b9050919050565b6000602082019050611276600083018461117c565b92915050565b6000608082019050611291600083018861117c565b61129e60208301876110c7565b81810360408301526112b1818587611057565b90506112c06060830184611031565b9695505050505050565b6000604051905081810181811067ffffffffffffffff821117156112ed57600080fd5b8060405250919050565b600067ffffffffffffffff82111561130e57600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611357826113a7565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b60008190506113a28261145f565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b60006113e6826113ed565b9050919050565b60006113f8826113a7565b9050919050565b600061140a82611411565b9050919050565b600061141c826113a7565b9050919050565b600061142e82611394565b9050919050565b82818337600083830152505050565b6000819050919050565b6000601f19601f8301169050919050565b6002811061146957fe5b50565b6114758161134c565b811461148057600080fd5b50565b61148c81611368565b811461149757600080fd5b50565b600281106114a757600080fd5b50565b6114b3816113c7565b81146114be57600080fd5b5056fea26469706673582212204560466009d7492bb8679eb3421b0f3e094afd640be59e91ade4e14a9cd3811664736f6c63430007010033"

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
var ConfigContractBin = "0x60a06040523480156200001157600080fd5b5060405162003ba538038062003ba58339818101604052810190620000379190620006d1565b6000620000496200039760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506001620000f96200039f60201b60201c565b908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010190805190602001906200019e929190620004f2565b5060608201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101208201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101408201518160040160146101000a81548163ffffffff021916908360e01c02179055506101608201518160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050508067ffffffffffffffff1660808167ffffffffffffffff1660c01b81525050506200072b565b600033905090565b620003a962000581565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff81118015620003f057600080fd5b50604051908082528060200260200182016040528015620004205781602001602082028036833780820191505090505b508152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff16815250905090565b8280548282559060005260206000209081019282156200056e579160200282015b828111156200056d5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000513565b5b5090506200057d91906200067d565b5090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b5b80821115620006b657600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016200067e565b5090565b600081519050620006cb8162000711565b92915050565b600060208284031215620006e457600080fd5b6000620006f484828501620006ba565b91505092915050565b600067ffffffffffffffff82169050919050565b6200071c81620006fd565b81146200072857600080fd5b50565b60805160c01c61344d6200075860003980610a7352806114085280611f685280612203525061344d6000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c806373ed43db116100f9578063cd21aee711610097578063d9a58f2411610071578063d9a58f2414610453578063e008cb6214610483578063f2fde38b146104b3578063fa84ea02146104cf576101a8565b8063cd21aee7146103fd578063ce9919b81461041b578063d1ac2e5214610437576101a8565b80639d63753e116100d35780639d63753e1461038d578063bcf67268146103a9578063c7c6e9f4146103c5578063c9515c58146103e1576101a8565b806373ed43db1461033757806381e905a3146103535780638da5cb5b1461036f576101a8565b80635dc6fdb81161016657806364e9f6711161014057806364e9f671146102b9578063660744dc146102e1578063715018a614610311578063719f2e171461031b576101a8565b80635dc6fdb814610265578063606df5141461028157806362fced0e1461029d576101a8565b806298fa22146101ad5780630f0aae6f146101e757806318b5e83014610205578063287447c41461020f5780632b2cc6c41461022d578063564093fc14610249575b600080fd5b6101c760048036038101906101c29190612eaa565b6104ff565b6040516101de9b9a9998979695949392919061324a565b60405180910390f35b6101ef610653565b6040516101fc919061322f565b60405180910390f35b61020d610660565b005b6102176111c1565b604051610224919061322f565b60405180910390f35b61024760048036038101906102429190612e13565b6111d1565b005b610263600480360381019061025e9190612ed3565b6112ad565b005b61027f600480360381019061027a9190612ed3565b611371565b005b61029b60048036038101906102969190612ed3565b611474565b005b6102b760048036038101906102b29190612e3c565b611537565b005b6102c16116aa565b6040516102d89b9a9998979695949392919061324a565b60405180910390f35b6102fb60048036038101906102f69190612ed3565b6117df565b60405161030891906131b2565b60405180910390f35b61031961182d565b005b61033560048036038101906103309190612ed3565b611980565b005b610351600480360381019061034c9190612ed3565b611a44565b005b61036d60048036038101906103689190612ed3565b611b07565b005b610377611bcb565b60405161038491906131b2565b60405180910390f35b6103a760048036038101906103a29190612ed3565b611bf4565b005b6103c360048036038101906103be9190612e13565b611d32565b005b6103df60048036038101906103da9190612ed3565b611e0e565b005b6103fb60048036038101906103f69190612ed3565b611ed1565b005b610405612201565b604051610412919061322f565b60405180910390f35b61043560048036038101906104309190612ed3565b612225565b005b610451600480360381019061044c9190612e81565b6122e9565b005b61046d60048036038101906104689190612ed3565b6123a2565b60405161047a919061322f565b60405180910390f35b61049d60048036038101906104989190612ed3565b6123d7565b6040516104aa919061320d565b60405180910390f35b6104cd60048036038101906104c89190612e13565b612799565b005b6104e960048036038101906104e49190612efc565b61295b565b6040516104f691906131b2565b60405180910390f35b6001818154811061050c57fe5b90600052602060002090600502016000915090508060000160009054906101000a900467ffffffffffffffff16908060000160089054906101000a900467ffffffffffffffff16908060020160009054906101000a900467ffffffffffffffff16908060020160089054906101000a900467ffffffffffffffff16908060020160109054906101000a900467ffffffffffffffff16908060020160189054906101000a900467ffffffffffffffff16908060030160009054906101000a900467ffffffffffffffff16908060030160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460e01b908060040160189054906101000a900467ffffffffffffffff1690508b565b6000600180549050905090565b6106686129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ec906131ed565b60405180910390fd5b600167ffffffffffffffff0367ffffffffffffffff166001805490501061071b57600080fd5b610723612b24565b6001808080549050038154811061073657fe5b9060005260206000209060050201604051806101800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801561083957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116107ef575b505050505081526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020016004820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff164301600260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610acc57600080fd5b6000816080015167ffffffffffffffff161115610b8f57806000015167ffffffffffffffff16600260000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610b2057600080fd5b60008160000151600260000160009054906101000a900467ffffffffffffffff16039050600260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168183608001510283602001510167ffffffffffffffff1614610b8957600080fd5b50610bcd565b806000015167ffffffffffffffff16600260000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610bcc57600080fd5b5b6001600290806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900467ffffffffffffffff168160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160089054906101000a900467ffffffffffffffff168160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820181600101908054610c96929190612c20565b506002820160009054906101000a900467ffffffffffffffff168160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160089054906101000a900467ffffffffffffffff168160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160109054906101000a900467ffffffffffffffff168160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002820160189054906101000a900467ffffffffffffffff168160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506003820160009054906101000a900467ffffffffffffffff168160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004820160149054906101000a900460e01b8160040160146101000a81548163ffffffff021916908360e01c02179055506004820160189054906101000a900467ffffffffffffffff168160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050610f326129d5565b600260008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151816001019080519060200190610fae929190612c72565b5060608201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160020160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101208201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101408201518160040160146101000a81548163ffffffff021916908360e01c02179055506101608201518160040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050507f38889ef980014448a73b6e5dc5579ba1a4b7bd213a586b3f4832351448c48ae66001805490506040516111b6919061322f565b60405180910390a150565b6000600260010180549050905090565b6111d96129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125d906131ed565b60405180910390fd5b80600260030160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6112b56129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611342576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611339906131ed565b60405180910390fd5b80600260030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6113796129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fd906131ed565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168167ffffffffffffffff161061144657600080fd5b806002800160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b61147c6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611509576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611500906131ed565b60405180910390fd5b806002800160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b61153f6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146115cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c3906131ed565b60405180910390fd5b8181905067ffffffffffffffff80160360026001018054905011156115f057600080fd5b60005b828290508167ffffffffffffffff1610156116a557600260010183838367ffffffffffffffff1681811061162357fe5b90506020020160208101906116389190612e13565b9080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506115f3565b505050565b60028060000160009054906101000a900467ffffffffffffffff16908060000160089054906101000a900467ffffffffffffffff16908060020160009054906101000a900467ffffffffffffffff16908060020160089054906101000a900467ffffffffffffffff16908060020160109054906101000a900467ffffffffffffffff16908060020160189054906101000a900467ffffffffffffffff16908060030160009054906101000a900467ffffffffffffffff16908060030160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460e01b908060040160189054906101000a900467ffffffffffffffff1690508b565b600060026001018267ffffffffffffffff16815481106117fb57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6118356129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118b9906131ed565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6119886129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611a15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a0c906131ed565b60405180910390fd5b80600260040160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611a4c6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611ad9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad0906131ed565b60405180910390fd5b806002800160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611b0f6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611b9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b93906131ed565b60405180910390fd5b80600260000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611bfc6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611c89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c80906131ed565b60405180910390fd5b60006002600101805490509050808267ffffffffffffffff1611611d1c5760005b8267ffffffffffffffff168167ffffffffffffffff161015611d16576002600101805480611cd457fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590558080600101915050611caa565b50611d2e565b60026001016000611d2d9190612cfc565b5b5050565b611d3a6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dbe906131ed565b60405180910390fd5b80600260040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611e166129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611ea3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e9a906131ed565b60405180910390fd5b806002800160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b611ed96129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611f66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5d906131ed565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1643018167ffffffffffffffff1611611fa857600080fd5b600060018054905090506000600180805490500390505b60008111156121a557600060018281548110611fd757fe5b906000526020600020906005020190508367ffffffffffffffff168160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161061219057600180548061202557fe5b6001900381819060005260206000209060050201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556001820160006120809190612cfc565b6002820160006101000a81549067ffffffffffffffff02191690556002820160086101000a81549067ffffffffffffffff02191690556002820160106101000a81549067ffffffffffffffff02191690556002820160186101000a81549067ffffffffffffffff02191690556003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160146101000a81549063ffffffff02191690556004820160186101000a81549067ffffffffffffffff021916905550509055612196565b506121a5565b50808060019003915050611fbf565b508067ffffffffffffffff16600180549050106121c157600080fd5b7f202adac5e4f5fa65a6e6ec3afc99da8986c020c2799f4e0aee50552a05a0bfdf6001805490506040516121f5919061322f565b60405180910390a15050565b7f000000000000000000000000000000000000000000000000000000000000000081565b61222d6129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146122ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b1906131ed565b60405180910390fd5b80600260000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6122f16129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461237e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612375906131ed565b60405180910390fd5b80600260040160146101000a81548163ffffffff021916908360e01c021790555050565b600060018267ffffffffffffffff16815481106123bb57fe5b9060005260206000209060050201600101805490509050919050565b6123df612b24565b6000600180805490500390505b6000811061278a5760006001828154811061240357fe5b906000526020600020906005020190508367ffffffffffffffff168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff161161277b5780604051806101800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801561253c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116124f2575b505050505081526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020016004820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505092505050612794565b508080600190039150506123ec565b50600061279357fe5b5b919050565b6127a16129cd565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461282e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612825906131ed565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561289e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612895906131cd565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060018367ffffffffffffffff168154811061297457fe5b90600052602060002090600502016001018267ffffffffffffffff168154811061299a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b600033905090565b6129dd612b24565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff81118015612a2357600080fd5b50604051908082528060200260200182016040528015612a525781602001602082028036833780820191505090505b508152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff16815250905090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b828054828255906000526020600020908101928215612c615760005260206000209182015b82811115612c60578254825591600101919060010190612c45565b5b509050612c6e9190612d1d565b5090565b828054828255906000526020600020908101928215612ceb579160200282015b82811115612cea5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612c92565b5b509050612cf89190612d1d565b5090565b5080546000825590600052602060002090810190612d1a9190612d58565b50565b5b80821115612d5457600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612d1e565b5090565b5b80821115612d71576000816000905550600101612d59565b5090565b600081359050612d84816133bb565b92915050565b60008083601f840112612d9c57600080fd5b8235905067ffffffffffffffff811115612db557600080fd5b602083019150836020820283011115612dcd57600080fd5b9250929050565b600081359050612de3816133d2565b92915050565b600081359050612df8816133e9565b92915050565b600081359050612e0d81613400565b92915050565b600060208284031215612e2557600080fd5b6000612e3384828501612d75565b91505092915050565b60008060208385031215612e4f57600080fd5b600083013567ffffffffffffffff811115612e6957600080fd5b612e7585828601612d8a565b92509250509250929050565b600060208284031215612e9357600080fd5b6000612ea184828501612dd4565b91505092915050565b600060208284031215612ebc57600080fd5b6000612eca84828501612de9565b91505092915050565b600060208284031215612ee557600080fd5b6000612ef384828501612dfe565b91505092915050565b60008060408385031215612f0f57600080fd5b6000612f1d85828601612dfe565b9250506020612f2e85828601612dfe565b9150509250929050565b6000612f448383612f50565b60208301905092915050565b612f598161333f565b82525050565b612f688161333f565b82525050565b6000612f7982613305565b612f83818561331d565b9350612f8e836132f5565b8060005b83811015612fbf578151612fa68882612f38565b9750612fb183613310565b925050600181019050612f92565b5085935050505092915050565b612fd581613351565b82525050565b612fe481613351565b82525050565b6000612ff760268361332e565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061305d60208361332e565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000610180830160008301516130a96000860182613194565b5060208301516130bc6020860182613194565b50604083015184820360408601526130d48282612f6e565b91505060608301516130e96060860182613194565b5060808301516130fc6080860182613194565b5060a083015161310f60a0860182613194565b5060c083015161312260c0860182613194565b5060e083015161313560e0860182613194565b5061010083015161314a610100860182612f50565b5061012083015161315f610120860182612f50565b50610140830151613174610140860182612fcc565b50610160830151613189610160860182613194565b508091505092915050565b61319d816133a7565b82525050565b6131ac816133a7565b82525050565b60006020820190506131c76000830184612f5f565b92915050565b600060208201905081810360008301526131e681612fea565b9050919050565b6000602082019050818103600083015261320681613050565b9050919050565b600060208201905081810360008301526132278184613090565b905092915050565b600060208201905061324460008301846131a3565b92915050565b600061016082019050613260600083018e6131a3565b61326d602083018d6131a3565b61327a604083018c6131a3565b613287606083018b6131a3565b613294608083018a6131a3565b6132a160a08301896131a3565b6132ae60c08301886131a3565b6132bb60e0830187612f5f565b6132c9610100830186612f5f565b6132d7610120830185612fdb565b6132e56101408301846131a3565b9c9b505050505050505050505050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061334a8261337d565b9050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6133c48161333f565b81146133cf57600080fd5b50565b6133db81613351565b81146133e657600080fd5b50565b6133f28161339d565b81146133fd57600080fd5b50565b613409816133a7565b811461341457600080fd5b5056fea2646970667358221220d39a0c82d8f9749a28f3bf6c7b6f599d851804bb9f32c22138a8e43b5f516e7764736f6c63430007010033"

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
var DepositContractBin = "0x6080604052731820a4b7618bde71dce8cdc73aab6c95905fad246000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503480156200006557600080fd5b5060405162001d1d38038062001d1d83398181016040528101906200008b91906200019d565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166329965a1d307fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b306040518463ffffffff1660e01b81526004016200014b93929190620001eb565b600060405180830381600087803b1580156200016657600080fd5b505af11580156200017b573d6000803e3d6000fd5b505050505062000294565b60008151905062000197816200027a565b92915050565b600060208284031215620001b057600080fd5b6000620001c08482850162000186565b91505092915050565b620001d48162000228565b82525050565b620001e5816200023c565b82525050565b6000606082019050620002026000830186620001c9565b620002116020830185620001da565b620002206040830184620001c9565b949350505050565b600062000235826200025a565b9050919050565b6000819050919050565b6000620002538262000228565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620002858162000246565b81146200029157600080fd5b50565b611a7980620002a46000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c8063b611d8d911610066578063b611d8d91461011b578063b799036c1461014b578063b8ba16fd1461017b578063c96be4cb146101ab578063dbaf2145146101c757610092565b806223de291461009757806351cff8d9146100b35780639b4fed88146100cf578063aabc2496146100ff575b600080fd5b6100b160048036038101906100ac91906111c6565b6101d1565b005b6100cd60048036038101906100c8919061119d565b61030e565b005b6100e960048036038101906100e4919061119d565b61066b565b6040516100f69190611958565b60405180910390f35b6101196004803603810190610114919061119d565b6106cb565b005b6101356004803603810190610130919061119d565b6107a0565b6040516101429190611958565b60405180910390f35b6101656004803603810190610160919061119d565b610800565b60405161017291906116e9565b60405180910390f35b6101956004803603810190610190919061119d565b610859565b6040516101a29190611897565b60405180910390f35b6101c560048036038101906101c0919061119d565b6108a5565b005b6101cf610b62565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610261576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025890611877565b60405180910390fd5b600884849050146102a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029e906117f7565b60405180910390fd5b60006102f685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610dae565b9050610303888783610e03565b505050505050505050565b6103166110eb565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff16151515158152505090506000816000015111610432576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042990611797565b60405180910390fd5b80606001511561043e57fe5b6000816040015167ffffffffffffffff161161048f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048690611757565b60405180910390fd5b806020015181604001510167ffffffffffffffff164310156104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90611777565b60405180910390fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff02191690555050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bd9bbc68383600001516040518363ffffffff1660e01b81526004016105db9291906116ad565b600060405180830381600087803b1580156105f557600080fd5b505af1158015610609573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f60008060006001600060405161065f959493929190611704565b60405180910390a25050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff169050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461075c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075390611857565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff169050919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160109054906101000a900460ff169050919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ff57600080fd5b6109076110eb565b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff161515151581525050905060008160000181815250506000816020019067ffffffffffffffff16908167ffffffffffffffff16815250506000816040019067ffffffffffffffff16908167ffffffffffffffff1681525050600181606001901515908115158152505080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548160ff0219169083151502179055509050508173ffffffffffffffffffffffffffffffffffffffff167f04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f6000806000806001604051610b56959493929190611704565b60405180910390a25050565b610b6a6110eb565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff16151515158152505090506000816000015111610c86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7d90611797565b60405180910390fd5b806060015115610c9257fe5b6000816040015167ffffffffffffffff1614610ce3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cda90611837565b60405180910390fd5b43600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f8260000151836020015143600080604051610da3959493929190611905565b60405180910390a250565b60008060005b8351811015610df9576001810184510360080260020a848281518110610dd657fe5b602001015160f81c60f81b60f81c60ff1602820191508080600101915050610db4565b5080915050919050565b610e0b6110eb565b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1615151515815250509050806020015167ffffffffffffffff168267ffffffffffffffff161015610f3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f32906117b7565b60405180910390fd5b6000816040015167ffffffffffffffff1614610f8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f83906117d7565b60405180910390fd5b806060015115610fd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc890611817565b60405180910390fd5b82816000015101600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff167f04a1c8e18f4a4bf5e4fe7ea1e127365af43f3249cae762ca50d69a2257acc97f848360000151018460008060006040516110dd9594939291906118b2565b60405180910390a250505050565b604051806080016040528060008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000151581525090565b60008135905061113881611a15565b92915050565b60008083601f84011261115057600080fd5b8235905067ffffffffffffffff81111561116957600080fd5b60208301915083600182028301111561118157600080fd5b9250929050565b60008135905061119781611a2c565b92915050565b6000602082840312156111af57600080fd5b60006111bd84828501611129565b91505092915050565b60008060008060008060008060c0898b0312156111e257600080fd5b60006111f08b828c01611129565b98505060206112018b828c01611129565b97505060406112128b828c01611129565b96505060606112238b828c01611188565b955050608089013567ffffffffffffffff81111561124057600080fd5b61124c8b828c0161113e565b945094505060a089013567ffffffffffffffff81111561126b57600080fd5b6112778b828c0161113e565b92509250509295985092959890939650565b61129281611995565b82525050565b6112a1816119a7565b82525050565b6112b0816119f1565b82525050565b6112bf81611a03565b82525050565b60006112d2602d83611984565b91507f4465706f736974436f6e74726163743a207769746864726177616c206e6f742060008301527f72657175657374656420796574000000000000000000000000000000000000006020830152604082019050919050565b6000611338603083611984565b91507f4465706f736974436f6e74726163743a207769746864726177616c2064656c6160008301527f79206e6f742070617373656420796574000000000000000000000000000000006020830152604082019050919050565b600061139e601b83611984565b91507f4465706f736974436f6e74726163743a206e6f206465706f73697400000000006000830152602082019050919050565b60006113de603583611984565b91507f4465706f736974436f6e74726163743a207769746864726177616c2064656c6160008301527f792063616e6e6f742062652064656372656173656400000000000000000000006020830152604082019050919050565b6000611444602783611984565b91507f4465706f736974436f6e74726163743a207769746864726177616c20696e207060008301527f726f6772657373000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006114aa602783611984565b91507f4465706f736974436f6e74726163743a20696e76616c6964207573657220646160008301527f74612073697a65000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611510602083611984565b91507f4465706f736974436f6e74726163743a206163636f756e7420736c61736865646000830152602082019050919050565b6000611550602d83611984565b91507f4465706f736974436f6e74726163743a207769746864726177616c20616c726560008301527f61647920726571756573746564000000000000000000000000000000000000006020830152604082019050919050565b60006115b6600083611973565b9150600082019050919050565b60006115d0602c83611984565b91507f4465706f736974436f6e74726163743a20736c6173686572206164647265737360008301527f20616c72656164792073657400000000000000000000000000000000000000006020830152604082019050919050565b6000611636602783611984565b91507f4465706f736974436f6e74726163743a20726563656976656420696e76616c6960008301527f6420746f6b656e000000000000000000000000000000000000000000000000006020830152604082019050919050565b611698816119d3565b82525050565b6116a7816119dd565b82525050565b60006060820190506116c26000830185611289565b6116cf602083018461168f565b81810360408301526116e0816115a9565b90509392505050565b60006020820190506116fe6000830184611298565b92915050565b600060a08201905061171960008301886112a7565b61172660208301876112b6565b61173360408301866112b6565b6117406060830185611298565b61174d6080830184611298565b9695505050505050565b60006020820190508181036000830152611770816112c5565b9050919050565b600060208201905081810360008301526117908161132b565b9050919050565b600060208201905081810360008301526117b081611391565b9050919050565b600060208201905081810360008301526117d0816113d1565b9050919050565b600060208201905081810360008301526117f081611437565b9050919050565b600060208201905081810360008301526118108161149d565b9050919050565b6000602082019050818103600083015261183081611503565b9050919050565b6000602082019050818103600083015261185081611543565b9050919050565b60006020820190508181036000830152611870816115c3565b9050919050565b6000602082019050818103600083015261189081611629565b9050919050565b60006020820190506118ac600083018461168f565b92915050565b600060a0820190506118c7600083018861168f565b6118d4602083018761169e565b6118e160408301866112b6565b6118ee6060830185611298565b6118fb6080830184611298565b9695505050505050565b600060a08201905061191a600083018861168f565b611927602083018761169e565b611934604083018661169e565b6119416060830185611298565b61194e6080830184611298565b9695505050505050565b600060208201905061196d600083018461169e565b92915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006119a0826119b3565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b60006119fc826119d3565b9050919050565b6000611a0e826119dd565b9050919050565b611a1e81611995565b8114611a2957600080fd5b50565b611a35816119d3565b8114611a4057600080fd5b5056fea264697066735822122097b68f405309512834bc97299a8ce0ff2dca2c159b3fdef4ec6fc03573883d1064736f6c63430007010033"

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
var ECDSABin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207cf2158b720f9add322f5a2f21431976918eaa607b5d4b7b455136f3bfb243fc64736f6c63430007010033"

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
var ERC777Bin = "0x60806040523480156200001157600080fd5b506040516200331538038062003315833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084640100000000821115620001d257600080fd5b83820191506020820185811115620001e957600080fd5b82518660208202830111640100000000821117156200020757600080fd5b8083526020830192505050908051906020019060200280838360005b838110156200024057808201518184015260208101905062000223565b505050509050016040525050508260029080519060200190620002659291906200050b565b5081600390805190602001906200027e9291906200050b565b5080600490805190602001906200029792919062000592565b5060005b600480549050811015620003475760016005600060048481548110620002bd57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806001019150506200029b565b50731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff166329965a1d307fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce2177054306040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019350505050600060405180830381600087803b1580156200040c57600080fd5b505af115801562000421573d6000803e3d6000fd5b50505050731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff166329965a1d307faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a306040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019350505050600060405180830381600087803b158015620004e957600080fd5b505af1158015620004fe573d6000803e3d6000fd5b505050505050506200067d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200054e57805160ff19168380011785556200057f565b828001600101855582156200057f579182015b828111156200057e57825182559160200191906001019062000561565b5b5090506200058e919062000621565b5090565b8280548282559060005260206000209081019282156200060e579160200282015b828111156200060d5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190620005b3565b5b5090506200061d919062000640565b5090565b5b808211156200063c57600081600090555060010162000622565b5090565b5b808211156200067957600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010162000641565b5090565b612c88806200068d6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b637114610746578063dd62ed3e146107c0578063fad8b32a14610838578063fc673c4f1461087c578063fe9d9303146109f857610116565b8063959b8c3f1461053657806395d89b411461057a5780639bd9bbc6146105fd578063a9059cbb146106e257610116565b806323b872dd116100e957806323b872dd1461027f578063313ce56714610303578063556f0dc71461032457806362ad1b831461034257806370a08231146104de57610116565b806306e485381461011b57806306fdde031461017a578063095ea7b3146101fd57806318160ddd14610261575b600080fd5b610123610abd565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561016657808201518184015260208101905061014b565b505050509050019250505060405180910390f35b610182610b4b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c25780820151818401526020810190506101a7565b50505050905090810190601f1680156101ef5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102496004803603604081101561021357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bed565b60405180821515815260200191505060405180910390f35b610269610c10565b6040518082815260200191505060405180910390f35b6102eb6004803603606081101561029557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c1a565b60405180821515815260200191505060405180910390f35b61030b610e78565b604051808260ff16815260200191505060405180910390f35b61032c610e81565b6040518082815260200191505060405180910390f35b6104dc600480360360a081101561035857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156103bf57600080fd5b8201836020820111156103d157600080fd5b803590602001918460018302840111640100000000831117156103f357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561045657600080fd5b82018360208201111561046857600080fd5b8035906020019184600183028401116401000000008311171561048a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e8a565b005b610520600480360360208110156104f457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f06565b6040518082815260200191505060405180910390f35b6105786004803603602081101561054c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f4e565b005b6105826111c5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105c25780820151818401526020810190506105a7565b50505050905090810190601f1680156105ef5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106e06004803603606081101561061357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561065a57600080fd5b82018360208201111561066c57600080fd5b8035906020019184600183028401116401000000008311171561068e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611267565b005b61072e600480360360408110156106f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611291565b60405180821515815260200191505060405180910390f35b6107a86004803603604081101561075c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113b5565b60405180821515815260200191505060405180910390f35b610822600480360360408110156107d657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611566565b6040518082815260200191505060405180910390f35b61087a6004803603602081101561084e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115ed565b005b6109f66004803603608081101561089257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156108d957600080fd5b8201836020820111156108eb57600080fd5b8035906020019184600183028401116401000000008311171561090d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561097057600080fd5b82018360208201111561098257600080fd5b803590602001918460018302840111640100000000831117156109a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611864565b005b610abb60048036036040811015610a0e57600080fd5b810190808035906020019092919080359060200190640100000000811115610a3557600080fd5b820183602082011115610a4757600080fd5b80359060200191846001830284011164010000000083111715610a6957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506118dc565b005b60606004805480602002602001604051908101604052809291908181526020018280548015610b4157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610af7575b5050505050905090565b606060028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610be35780601f10610bb857610100808354040283529160200191610be3565b820191906000526020600020905b815481529060010190602001808311610bc657829003601f168201915b5050505050905090565b600080610bf8611902565b9050610c0581858561190a565b600191505092915050565b6000600154905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ca1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612b6e6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610d27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612be76026913960400191505060405180910390fd5b6000610d31611902565b9050610d5f818686866040518060200160405280600081525060405180602001604052806000815250611b01565b610d8b818686866040518060200160405280600081525060405180602001604052806000815250611dc3565b610e3e8582610e3986604051806060016040528060298152602001612bbe60299139600860008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546120d29092919063ffffffff16565b61190a565b610e6c8186868660405180602001604052806000815250604051806020016040528060008152506000612192565b60019150509392505050565b60006012905090565b60006001905090565b610e9b610e95611902565b866113b5565b610ef0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612b92602c913960400191505060405180910390fd5b610eff858585858560016124d5565b5050505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b8073ffffffffffffffffffffffffffffffffffffffff16610f6d611902565b73ffffffffffffffffffffffffffffffffffffffff161415610fda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612adc6024913960400191505060405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156110c45760076000611038611902565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055611161565b6001600660006110d2611902565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b611169611902565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f960405160405180910390a350565b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561125d5780601f106112325761010080835404028352916020019161125d565b820191906000526020600020905b81548152906001019060200180831161124057829003601f168201915b5050505050905090565b61128c611272611902565b8484846040518060200160405280600081525060016124d5565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611318576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612b6e6024913960400191505060405180910390fd5b6000611322611902565b9050611350818286866040518060200160405280600081525060405180602001604052806000815250611b01565b61137c818286866040518060200160405280600081525060405180602001604052806000815250611dc3565b6113aa8182868660405180602001604052806000815250604051806020016040528060008152506000612192565b600191505092915050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614806114cd5750600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680156114cc5750600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b5b8061155e5750600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b905092915050565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6115f5611902565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611679576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b006021913960400191505060405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561176c576001600760006116d9611902565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550611800565b60066000611778611902565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b611808611902565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa160405160405180910390a350565b61187561186f611902565b856113b5565b6118ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612b92602c913960400191505060405180910390fd5b6118d68484848461263e565b50505050565b6118fe6118e7611902565b83836040518060200160405280600081525061263e565b5050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611990576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612a4c6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180612c306023913960400191505060405180910390fd5b80600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b6000731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff1663aabbb8ca877f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe89560001b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060206040518083038186803b158015611ba957600080fd5b505afa158015611bbd573d6000803e3d6000fd5b505050506040513d6020811015611bd357600080fd5b81019080805190602001909291905050509050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611dba578073ffffffffffffffffffffffffffffffffffffffff166375ab97828888888888886040518763ffffffff1660e01b8152600401808773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611ce9578082015181840152602081019050611cce565b50505050905090810190601f168015611d165780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015611d4f578082015181840152602081019050611d34565b50505050905090810190601f168015611d7c5780820380516001836020036101000a031916815260200191505b5098505050505050505050600060405180830381600087803b158015611da157600080fd5b505af1158015611db5573d6000803e3d6000fd5b505050505b50505050505050565b611dcf86868686612960565b611e3a83604051806060016040528060278152602001612a93602791396000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546120d29092919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611ecd836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461296690919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987868686604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611fc2578082015181840152602081019050611fa7565b50505050905090810190601f168015611fef5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561202857808201518184015260208101905061200d565b50505050905090810190601f1680156120555780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a48373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3505050505050565b600083831115829061217f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612144578082015181840152602081019050612129565b50505050905090810190601f1680156121715780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6000731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff1663aabbb8ca877fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b60001b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060206040518083038186803b15801561223a57600080fd5b505afa15801561224e573d6000803e3d6000fd5b505050506040513d602081101561226457600080fd5b81019080805190602001909291905050509050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461244e578073ffffffffffffffffffffffffffffffffffffffff166223de298989898989896040518763ffffffff1660e01b8152600401808773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561237957808201518184015260208101905061235e565b50505050905090810190601f1680156123a65780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156123df5780820151818401526020810190506123c4565b50505050905090810190601f16801561240c5780820380516001836020036101000a031916815260200191505b5098505050505050505050600060405180830381600087803b15801561243157600080fd5b505af1158015612445573d6000803e3d6000fd5b505050506124cb565b81156124ca576124738673ffffffffffffffffffffffffffffffffffffffff166129ee565b156124c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604d815260200180612b21604d913960600191505060405180910390fd5b5b5b5050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561255b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612a716022913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614156125fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4552433737373a2073656e6420746f20746865207a65726f206164647265737381525060200191505060405180910390fd5b6000612608611902565b9050612618818888888888611b01565b612626818888888888611dc3565b61263581888888888888612192565b50505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156126c4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612aba6022913960400191505060405180910390fd5b60006126ce611902565b90506126dd8186600087612960565b6126ec81866000878787611b01565b61275784604051806060016040528060238152602001612c0d602391396000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546120d29092919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506127ae84600154612a0190919063ffffffff16565b6001819055508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098868686604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015612850578082015181840152602081019050612835565b50505050905090810190601f16801561287d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156128b657808201518184015260208101905061289b565b50505050905090810190601f1680156128e35780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a3600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040518082815260200191505060405180910390a35050505050565b50505050565b6000808284019050838110156129e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600080823b905060008111915050919050565b6000612a4383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506120d2565b90509291505056fe4552433737373a20617070726f76652066726f6d20746865207a65726f20616464726573734552433737373a2073656e642066726f6d20746865207a65726f20616464726573734552433737373a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433737373a206275726e2066726f6d20746865207a65726f20616464726573734552433737373a20617574686f72697a696e672073656c66206173206f70657261746f724552433737373a207265766f6b696e672073656c66206173206f70657261746f724552433737373a20746f6b656e20726563697069656e7420636f6e747261637420686173206e6f20696d706c656d656e74657220666f7220455243373737546f6b656e73526563697069656e744552433737373a207472616e7366657220746f20746865207a65726f20616464726573734552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f7220666f7220686f6c6465724552433737373a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654552433737373a207472616e736665722066726f6d20746865207a65726f20616464726573734552433737373a206275726e20616d6f756e7420657863656564732062616c616e63654552433737373a20617070726f766520746f20746865207a65726f2061646472657373a26469706673582212205ef4106aabeb53556f7c978b71b0ea4daa3a3d039b4838de5d0021838e086e5064736f6c63430007010033"

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
var ExecutorContractBin = "0x60806040523480156200001157600080fd5b5060405162002152380380620021528339818101604052810190620000379190620000ee565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001bf565b600081519050620000d1816200018b565b92915050565b600081519050620000e881620001a5565b92915050565b600080604083850312156200010257600080fd5b60006200011285828601620000d7565b92505060206200012585828601620000c0565b9150509250929050565b60006200013c826200016b565b9050919050565b600062000150826200012f565b9050919050565b600062000164826200012f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001968162000143565b8114620001a257600080fd5b50565b620001b08162000157565b8114620001bc57600080fd5b50565b611f8380620001cf6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c87190b31161005b578063c87190b314610103578063ce0c159614610137578063d57a29d014610153578063fa6385f41461016f57610088565b806325b36cbf1461008d5780638f6dccfb146100bd578063beb3b50e146100c7578063bf66a182146100e5575b600080fd5b6100a760048036038101906100a291906115be565b61018d565b6040516100b49190611c26565b60405180910390f35b6100c5610285565b005b6100cf61052f565b6040516100dc9190611b10565b60405180910390f35b6100ed610555565b6040516100fa9190611b2b565b60405180910390f35b61011d600480360381019061011891906115be565b610579565b60405161012e959493929190611a99565b60405180910390f35b610151600480360381019061014c9190611511565b6105f0565b005b61016d600480360381019061016891906114a3565b610bd2565b005b610177610f16565b6040516101849190611c41565b60405180910390f35b6101956110cf565b600260008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160159054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16816102b157fe5b0667ffffffffffffffff16146102fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f390611b86565b60405180910390fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161032857fe5b049050610333611126565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b815260040161038c9190611c41565b60006040518083038186803b1580156103a457600080fd5b505afa1580156103b8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103e1919061157d565b90506000816080015167ffffffffffffffff1611610434576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042b90611ba6565b60405180910390fd5b806101600151600183018260800151028260200151010167ffffffffffffffff16431015610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90611b66565b60405180910390fd5b6001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507fa43f7c79f47e7937048e7a80ce05ad2cf3da87b2b1bf1bd1122c22a6234d34be600160149054906101000a900467ffffffffffffffff166040516105239190611c41565b60405180910390a15050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160159054906101000a900467ffffffffffffffff16908060010154908060020154905085565b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161061c57fe5b0667ffffffffffffffff1614610667576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065e90611b86565b60405180910390fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168161069357fe5b04905061069e611126565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b81526004016106f79190611c41565b60006040518083038186803b15801561070f57600080fd5b505afa158015610723573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061074c919061157d565b90506000816080015167ffffffffffffffff161161079f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079690611ba6565b60405180910390fd5b6001820181608001510281602001510167ffffffffffffffff164310156107fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f290611b66565b60405180910390fd5b8060400151518367ffffffffffffffff161061084c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084390611c06565b60405180910390fd5b80604001518367ffffffffffffffff168151811061086657fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d290611b46565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c87afa8a8360006040518363ffffffff1660e01b8152600401610939929190611c85565b60206040518083038186803b15801561095157600080fd5b505afa158015610965573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098991906114e8565b86146109ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c190611bc6565b60405180910390fd5b60006109e78261012001518361014001518460e001518989610f30565b90506040518060a001604052806001151581526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020018881526020018281525060026000600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010155608082015181600201559050506001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a600160149054906101000a900467ffffffffffffffff1682604051610bc1929190611c5c565b60405180910390a150505050505050565b60016002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681610bfe57fe5b0667ffffffffffffffff1614610c49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4090611b86565b60405180910390fd5b60006002600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681610c7557fe5b049050610c80611126565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62836040518263ffffffff1660e01b8152600401610cd99190611c41565b60006040518083038186803b158015610cf157600080fd5b505afa158015610d05573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d2e919061157d565b90506000816080015167ffffffffffffffff1611610d4857fe5b6001820181608001510281602001510167ffffffffffffffff16431015610d6b57fe5b6000610d888261012001518361014001518460e001518888610f30565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c87afa8a8460016040518363ffffffff1660e01b8152600401610de8929190611c85565b60206040518083038186803b158015610e0057600080fd5b505afa158015610e14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3891906114e8565b8114610e79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7090611be6565b60405180910390fd5b6001601481819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f3d5ed901c456e33248250feaab56f76b579b96b1679d7866e5bed2b6e9c5619a600160149054906101000a900467ffffffffffffffff1682604051610f07929190611c5c565b60405180910390a15050505050565b600160149054906101000a900467ffffffffffffffff1681565b60008060005b848490508167ffffffffffffffff1610156110c15760608786868467ffffffffffffffff16818110610f6457fe5b9050602002810190610f769190611cae565b604051602401610f87929190611aec565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008973ffffffffffffffffffffffffffffffffffffffff168867ffffffffffffffff16836040516110189190611a82565b60006040518083038160008787f1925050503d8060008114611056576040519150601f19603f3d011682016040523d82523d6000602084013e61105b565b606091505b5050905086868467ffffffffffffffff1681811061107557fe5b90506020028101906110879190611cae565b8560405160200161109a93929190611a58565b60405160208183030381529060405280519060200120935050508080600101915050610f36565b508091505095945050505050565b6040518060a00160405280600015158152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff16815260200160008019168152602001600080191681525090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b60008151905061123181611ef1565b92915050565b600082601f83011261124857600080fd5b815161125b61125682611d32565b611d05565b9150818183526020840193506020810190508385602084028201111561128057600080fd5b60005b838110156112b057816112968882611222565b845260208401935060208301925050600181019050611283565b5050505092915050565b60008083601f8401126112cc57600080fd5b8235905067ffffffffffffffff8111156112e557600080fd5b6020830191508360208202830111156112fd57600080fd5b9250929050565b60008135905061131381611f08565b92915050565b60008151905061132881611f08565b92915050565b60008151905061133d81611f1f565b92915050565b6000610180828403121561135657600080fd5b611361610180611d05565b905060006113718482850161148e565b60008301525060206113858482850161148e565b602083015250604082015167ffffffffffffffff8111156113a557600080fd5b6113b184828501611237565b60408301525060606113c58482850161148e565b60608301525060806113d98482850161148e565b60808301525060a06113ed8482850161148e565b60a08301525060c06114018482850161148e565b60c08301525060e06114158482850161148e565b60e08301525061010061142a84828501611222565b6101008301525061012061144084828501611222565b610120830152506101406114568482850161132e565b6101408301525061016061146c8482850161148e565b6101608301525092915050565b60008135905061148881611f36565b92915050565b60008151905061149d81611f36565b92915050565b600080602083850312156114b657600080fd5b600083013567ffffffffffffffff8111156114d057600080fd5b6114dc858286016112ba565b92509250509250929050565b6000602082840312156114fa57600080fd5b600061150884828501611319565b91505092915050565b6000806000806060858703121561152757600080fd5b600061153587828801611304565b945050602085013567ffffffffffffffff81111561155257600080fd5b61155e878288016112ba565b9350935050604061157187828801611479565b91505092959194509250565b60006020828403121561158f57600080fd5b600082015167ffffffffffffffff8111156115a957600080fd5b6115b584828501611343565b91505092915050565b6000602082840312156115d057600080fd5b60006115de84828501611479565b91505092915050565b6115f081611d92565b82525050565b6115ff81611d92565b82525050565b61160e81611da4565b82525050565b61161d81611da4565b82525050565b61162c81611db0565b82525050565b61163b81611db0565b82525050565b61165261164d82611db0565b611ec9565b82525050565b60006116648385611d65565b9350611671838584611e87565b61167a83611ed3565b840190509392505050565b60006116918385611d76565b935061169e838584611e87565b82840190509392505050565b60006116b582611d5a565b6116bf8185611d76565b93506116cf818560208601611e96565b80840191505092915050565b6116e481611e2d565b82525050565b6116f381611e51565b82525050565b61170281611e75565b82525050565b6000611715603083611d81565b91507f4578656375746f72436f6e74726163743a2073656e646572206973206e6f742060008301527f737065636966696564206b6579706572000000000000000000000000000000006020830152604082019050919050565b600061177b602983611d81565b91507f4578656375746f72436f6e74726163743a206261746368206973206e6f74206360008301527f6c6f7365642079657400000000000000000000000000000000000000000000006020830152604082019050919050565b60006117e1602683611d81565b91507f4578656375746f72436f6e74726163743a20756e65787065637465642068616c60008301527f66207374657000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611847602483611d81565b91507f4578656375746f72436f6e74726163743a20636f6e66696720697320696e616360008301527f74697665000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118ad602d83611d81565b91507f4578656375746f72436f6e74726163743a20696e636f7272656374206369706860008301527f65722062617463682068617368000000000000000000000000000000000000006020830152604082019050919050565b6000611913602b83611d81565b91507f4578656375746f72436f6e74726163743a206261746368206861736820646f6560008301527f73206e6f74206d617463680000000000000000000000000000000000000000006020830152604082019050919050565b6000611979602c83611d81565b91507f4578656375746f72436f6e74726163743a206b657970657220696e646578206f60008301527f7574206f6620626f756e647300000000000000000000000000000000000000006020830152604082019050919050565b60a0820160008201516119e86000850182611605565b5060208201516119fb60208501826115e7565b506040820151611a0e6040850182611a3a565b506060820151611a216060850182611623565b506080820151611a346080850182611623565b50505050565b611a4381611e19565b82525050565b611a5281611e19565b82525050565b6000611a65828587611685565b9150611a718284611641565b602082019150819050949350505050565b6000611a8e82846116aa565b915081905092915050565b600060a082019050611aae6000830188611614565b611abb60208301876115f6565b611ac86040830186611a49565b611ad56060830185611632565b611ae26080830184611632565b9695505050505050565b60006020820190508181036000830152611b07818486611658565b90509392505050565b6000602082019050611b2560008301846116db565b92915050565b6000602082019050611b4060008301846116ea565b92915050565b60006020820190508181036000830152611b5f81611708565b9050919050565b60006020820190508181036000830152611b7f8161176e565b9050919050565b60006020820190508181036000830152611b9f816117d4565b9050919050565b60006020820190508181036000830152611bbf8161183a565b9050919050565b60006020820190508181036000830152611bdf816118a0565b9050919050565b60006020820190508181036000830152611bff81611906565b9050919050565b60006020820190508181036000830152611c1f8161196c565b9050919050565b600060a082019050611c3b60008301846119d2565b92915050565b6000602082019050611c566000830184611a49565b92915050565b6000604082019050611c716000830185611a49565b611c7e6020830184611632565b9392505050565b6000604082019050611c9a6000830185611a49565b611ca760208301846116f9565b9392505050565b60008083356001602003843603038112611cc757600080fd5b80840192508235915067ffffffffffffffff821115611ce557600080fd5b602083019250600182023603831315611cfd57600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff82111715611d2857600080fd5b8060405250919050565b600067ffffffffffffffff821115611d4957600080fd5b602082029050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611d9d82611df9565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050611df482611ee4565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000611e3882611e3f565b9050919050565b6000611e4a82611df9565b9050919050565b6000611e5c82611e63565b9050919050565b6000611e6e82611df9565b9050919050565b6000611e8082611de6565b9050919050565b82818337600083830152505050565b60005b83811015611eb4578082015181840152602081019050611e99565b83811115611ec3576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b60028110611eee57fe5b50565b611efa81611d92565b8114611f0557600080fd5b50565b611f1181611db0565b8114611f1c57600080fd5b50565b611f2881611dba565b8114611f3357600080fd5b50565b611f3f81611e19565b8114611f4a57600080fd5b5056fea2646970667358221220b5d7286842727584d85c7b71779c604605afd58bb89c006faeb2d6b93eb3f6d964736f6c63430007010033"

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
var FeeBankContractBin = "0x608060405234801561001057600080fd5b506106ec806100206000396000f3fe60806040526004361061003f5760003560e01c80633ccfd60b14610044578063d6dad0601461005b578063f340fa01146100c0578063fc7e286d14610104575b600080fd5b34801561005057600080fd5b50610059610173565b005b34801561006757600080fd5b506100be6004803603604081101561007e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803567ffffffffffffffff1690602001909291905050506101d1565b005b610102600480360360208110156100d657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101df565b005b34801561011057600080fd5b506101536004803603602081101561012757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103fb565b604051808267ffffffffffffffff16815260200191505060405180910390f35b6101cf336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16610422565b565b6101db8282610422565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561021957600080fd5b6000341161022657600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff0367ffffffffffffffff1634111561029957600080fd5b346000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff160192506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc8b0ade8d126aac77fd16ecf68538fc2dfcc7cf77e879421a5907c3dff4fc4d93382346000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16604051808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815260200194505050505060405180910390a150565b60006020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561045c57600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16905060008167ffffffffffffffff16116104ca57600080fd5b8067ffffffffffffffff168267ffffffffffffffff1611156104eb57600080fd5b8181036000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060008373ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff1660405180600001905060006040518083038185875af1925050503d80600081146105bc576040519150601f19603f3d011682016040523d82523d6000602084013e6105c1565b606091505b50509050806105cf57600080fd5b7f4b8a4210268358b51dbd708b44cd83ba67563b2fba3c695343cc3f3e160d796e3385856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16604051808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815260200194505050505060405180910390a15050505056fea2646970667358221220ea6b4e4a3019e912cac020d8e27c0a558c064c4d14f88ef73d4c463f1380eb2464736f6c63430007010033"

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
var KeyBroadcastContractBin = "0x608060405234801561001057600080fd5b50604051610c83380380610c838339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610b758061010e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632712860b1461003b578063bf66a18214610057575b600080fd5b61005560048036038101906100509190610605565b610075565b005b61005f6101ee565b60405161006c9190610893565b60405180910390f35b61007d610212565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb62886040518263ffffffff1660e01b81526004016100d691906108ae565b60006040518083038186803b1580156100ee57600080fd5b505afa158015610102573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061012b91906105c4565b90508060400151518867ffffffffffffffff161061014857600080fd5b80604001518867ffffffffffffffff168151811061016257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a157600080fd5b7f4ddc6902637f5238f6c476ab331709d0cacd9b9aadd36ab5661a6b3c943bb61a338888888888886040516101dc979695949392919061082e565b60405180910390a15050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b60008151905061031d81610ae3565b92915050565b600082601f83011261033457600080fd5b8151610347610342826108f6565b6108c9565b9150818183526020840193506020810190508385602084028201111561036c57600080fd5b60005b8381101561039c5781610382888261030e565b84526020840193506020830192505060018101905061036f565b5050505092915050565b60008083601f8401126103b857600080fd5b8235905067ffffffffffffffff8111156103d157600080fd5b6020830191508360208202830111156103e957600080fd5b9250929050565b60008083601f84011261040257600080fd5b8235905067ffffffffffffffff81111561041b57600080fd5b60208301915083602082028301111561043357600080fd5b9250929050565b60008135905061044981610afa565b92915050565b60008151905061045e81610b11565b92915050565b6000610180828403121561047757600080fd5b6104826101806108c9565b90506000610492848285016105af565b60008301525060206104a6848285016105af565b602083015250604082015167ffffffffffffffff8111156104c657600080fd5b6104d284828501610323565b60408301525060606104e6848285016105af565b60608301525060806104fa848285016105af565b60808301525060a061050e848285016105af565b60a08301525060c0610522848285016105af565b60c08301525060e0610536848285016105af565b60e08301525061010061054b8482850161030e565b610100830152506101206105618482850161030e565b610120830152506101406105778482850161044f565b6101408301525061016061058d848285016105af565b6101608301525092915050565b6000813590506105a981610b28565b92915050565b6000815190506105be81610b28565b92915050565b6000602082840312156105d657600080fd5b600082015167ffffffffffffffff8111156105f057600080fd5b6105fc84828501610464565b91505092915050565b600080600080600080600060a0888a03121561062057600080fd5b600061062e8a828b0161059a565b975050602061063f8a828b0161059a565b96505060406106508a828b0161043a565b955050606088013567ffffffffffffffff81111561066d57600080fd5b6106798a828b016103f0565b9450945050608088013567ffffffffffffffff81111561069857600080fd5b6106a48a828b016103a6565b925092505092959891949750929550565b60006106c28484846107d4565b90509392505050565b60006106d78383610810565b60208301905092915050565b6106ec81610a69565b82525050565b60006106fe838561094c565b9350836020840285016107108461091e565b8060005b8781101561075657848403895261072b828461097f565b6107368682846106b5565b955061074184610932565b935060208b019a505050600181019050610714565b50829750879450505050509392505050565b6000610774838561095d565b935061077f82610928565b8060005b858110156107b85761079582846109d6565b61079f88826106cb565b97506107aa8361093f565b925050600181019050610783565b5085925050509392505050565b6107ce816109ff565b82525050565b60006107e0838561096e565b93506107ed838584610ac3565b6107f683610ad2565b840190509392505050565b61080a81610a7b565b82525050565b61081981610a55565b82525050565b61082881610a55565b82525050565b600060a082019050610843600083018a6106e3565b610850602083018961081f565b61085d60408301886107c5565b8181036060830152610870818688610768565b905081810360808301526108858184866106f2565b905098975050505050505050565b60006020820190506108a86000830184610801565b92915050565b60006020820190506108c3600083018461081f565b92915050565b6000604051905081810181811067ffffffffffffffff821117156108ec57600080fd5b8060405250919050565b600067ffffffffffffffff82111561090d57600080fd5b602082029050602081019050919050565b6000819050919050565b6000819050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000808335600160200384360303811261099857600080fd5b83810192508235915060208301925067ffffffffffffffff8211156109bc57600080fd5b6001820236038413156109ce57600080fd5b509250929050565b60006109e5602084018461059a565b905092915050565b60006109f882610a35565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000610a7482610a9f565b9050919050565b6000610a8682610a8d565b9050919050565b6000610a9882610a35565b9050919050565b6000610aaa82610ab1565b9050919050565b6000610abc82610a35565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b610aec816109ed565b8114610af757600080fd5b50565b610b03816109ff565b8114610b0e57600080fd5b50565b610b1a81610a09565b8114610b2557600080fd5b50565b610b3181610a55565b8114610b3c57600080fd5b5056fea264697066735822122014da405bff93903057f643ae41e8afdf1bed7d1174459648e0657fb6d3db481d64736f6c63430007010033"

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
const KeyperSlasherABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_appealBlocks\",\"type\":\"uint256\"},{\"internalType\":\"contractConfigContract\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"contractExecutorContract\",\"name\":\"_executorContract\",\"type\":\"address\"},{\"internalType\":\"contractDepositContract\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"}],\"name\":\"Accused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Appealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"accusations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"accused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"appealed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_halfStep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_keyperIndex\",\"type\":\"uint64\"}],\"name\":\"accuse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"halfStep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"signerIndices\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structAuthorization\",\"name\":\"_authorization\",\"type\":\"tuple\"}],\"name\":\"appeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_halfStep\",\"type\":\"uint64\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeyperSlasherBin is the compiled bytecode used for deploying new contracts.
var KeyperSlasherBin = "0x60806040523480156200001157600080fd5b506040516200219f3803806200219f8339818101604052810190620000379190620001f7565b83600381905550826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663aabc2496306040518263ffffffff1660e01b81526004016200015d919062000274565b600060405180830381600087803b1580156200017857600080fd5b505af11580156200018d573d6000803e3d6000fd5b505050505050505062000373565b600081519050620001ac816200030b565b92915050565b600081519050620001c38162000325565b92915050565b600081519050620001da816200033f565b92915050565b600081519050620001f18162000359565b92915050565b600080600080608085870312156200020e57600080fd5b60006200021e87828801620001e0565b945050602062000231878288016200019b565b93505060406200024487828801620001c9565b92505060606200025787828801620001b2565b91505092959194509250565b6200026e8162000291565b82525050565b60006020820190506200028b600083018462000263565b92915050565b60006200029e82620002e1565b9050919050565b6000620002b28262000291565b9050919050565b6000620002c68262000291565b9050919050565b6000620002da8262000291565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200031681620002a5565b81146200032257600080fd5b50565b6200033081620002b9565b81146200033c57600080fd5b50565b6200034a81620002cd565b81146200035657600080fd5b50565b620003648162000301565b81146200037057600080fd5b50565b611e1c80620003836000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630e98ad4d1461005157806331217be11461006d5780636864e7ee14610089578063ab4dfa8a146100bd575b600080fd5b61006b60048036038101906100669190611729565b6100d9565b005b610087600480360381019061008291906117d4565b610469565b005b6100a3600480360381019061009e91906117d4565b6106bc565b6040516100b4959493929190611a56565b60405180910390f35b6100d760048036038101906100d291906117fd565b610754565b005b6100e1611078565b60046000836000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff161515151581526020016000820160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160169054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050806000015161021757600080fd5b80602001511561022657600080fd5b61022e6110d5565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166325b36cbf84600001516040518263ffffffff1660e01b815260040161028d9190611b6e565b60a06040518083038186803b1580156102a557600080fd5b505afa1580156102b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102dd91906117ab565b90506102e98382610b72565b60018260200190151590811515815250508160046000856000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160000160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160000160166101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550905050806020015173ffffffffffffffffffffffffffffffffffffffff16836000015167ffffffffffffffff167f8944310cd346a8f80f86856adefdb3198175e1aeaffba79cc48a59c5f5e833cf60405160405180910390a3505050565b610471611078565b600460008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff161515151581526020016000820160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160169054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050905080600001516105a357600080fd5b8060200151156105b257600080fd5b600354816080015167ffffffffffffffff16014310156105d157600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c96be4cb82604001516040518263ffffffff1660e01b81526004016106309190611a3b565b600060405180830381600087803b15801561064a57600080fd5b505af115801561065e573d6000803e3d6000fd5b50505050806040015173ffffffffffffffffffffffffffffffffffffffff16816060015167ffffffffffffffff167fa24f6ac4cdf4d4719e335f2105dd8dc53263629951b8eb8a4c9b478f348a9ac560405160405180910390a35050565b60046020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900460ff16908060000160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160169054906101000a900467ffffffffffffffff16908060010160009054906101000a900467ffffffffffffffff16905085565b600460008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff161561079657600080fd5b61079e61112c565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb6260028567ffffffffffffffff16816107ee57fe5b046040518263ffffffff1660e01b815260040161080b9190611b6e565b60006040518083038186803b15801561082357600080fd5b505afa158015610837573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610860919061176a565b90508060400151518267ffffffffffffffff161061087d57600080fd5b80604001518267ffffffffffffffff168151811061089757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d657600080fd5b6108de6110d5565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166325b36cbf856040518263ffffffff1660e01b81526004016109399190611b6e565b60a06040518083038186803b15801561095157600080fd5b505afa158015610965573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098991906117ab565b9050806000015161099957600080fd5b6040518060a00160405280600115158152602001600015158152602001826020015173ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff1681526020014367ffffffffffffffff16815250600460008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160000160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160000160166101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff168567ffffffffffffffff167f79772647abf0e802e4c10672afe8fe89a61f8fd54e247f62c24204b190639f1760405160405180910390a450505050565b610b7a61112c565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e008cb626002846040015167ffffffffffffffff1681610bce57fe5b046040518263ffffffff1660e01b8152600401610beb9190611b6e565b60006040518083038186803b158015610c0357600080fd5b505afa158015610c17573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610c40919061176a565b9050806060015167ffffffffffffffff168360600151511015610c6257600080fd5b82604001515183606001515114610c7857600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beb3b50e6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ce257600080fd5b505afa158015610cf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1a9190611700565b83606001518460800151604051602001610d36939291906119fe565b60405160208183030381529060405280519060200120905060005b8460600151518167ffffffffffffffff161015610e8457606085606001518267ffffffffffffffff1681518110610d8457fe5b60200260200101519050600086604001518367ffffffffffffffff1681518110610daa57fe5b6020026020010151905060008367ffffffffffffffff161480610e04575086604001516001840367ffffffffffffffff1681518110610de557fe5b602002602001015167ffffffffffffffff168167ffffffffffffffff16115b610e0d57600080fd5b6000610e198584610e8b565b905085604001518267ffffffffffffffff1681518110610e3557fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e7457600080fd5b5050508080600101915050610d51565b5050505050565b60006041825114610ed1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec890611b0e565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610f54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4b90611b2e565b60405180910390fd5b601b8160ff161480610f695750601c8160ff16145b610fa8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9f90611b4e565b60405180910390fd5b600060018783868660405160008152602001604052604051610fcd9493929190611aa9565b6020604051602081039080840390855afa158015610fef573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561106b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106290611aee565b60405180910390fd5b8094505050505092915050565b6040518060a00160405280600015158152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600015158152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff16815260200160008019168152602001600080191681525090565b604051806101800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600067ffffffffffffffff1681525090565b60008151905061123781611d5c565b92915050565b600082601f83011261124e57600080fd5b815161126161125c82611bb6565b611b89565b9150818183526020840193506020810190508385602084028201111561128657600080fd5b60005b838110156112b6578161129c8882611228565b845260208401935060208301925050600181019050611289565b5050505092915050565b600082601f8301126112d157600080fd5b81356112e46112df82611bde565b611b89565b9150818183526020840193506020810190508360005b8381101561132a5781358601611310888261140b565b8452602084019350602083019250506001810190506112fa565b5050505092915050565b600082601f83011261134557600080fd5b813561135861135382611c06565b611b89565b9150818183526020840193506020810190508385602084028201111561137d57600080fd5b60005b838110156113ad578161139388826116d6565b845260208401935060208301925050600181019050611380565b5050505092915050565b6000815190506113c681611d73565b92915050565b6000813590506113db81611d8a565b92915050565b6000815190506113f081611d8a565b92915050565b60008151905061140581611da1565b92915050565b600082601f83011261141c57600080fd5b813561142f61142a82611c2e565b611b89565b9150808252602083016020830185838301111561144b57600080fd5b611456838284611d12565b50505092915050565b60008151905061146e81611db8565b92915050565b60006080828403121561148657600080fd5b6114906080611b89565b905060006114a0848285016116d6565b60008301525060206114b4848285016113cc565b602083015250604082013567ffffffffffffffff8111156114d457600080fd5b6114e084828501611334565b604083015250606082013567ffffffffffffffff81111561150057600080fd5b61150c848285016112c0565b60608301525092915050565b6000610180828403121561152b57600080fd5b611536610180611b89565b90506000611546848285016116eb565b600083015250602061155a848285016116eb565b602083015250604082015167ffffffffffffffff81111561157a57600080fd5b6115868482850161123d565b604083015250606061159a848285016116eb565b60608301525060806115ae848285016116eb565b60808301525060a06115c2848285016116eb565b60a08301525060c06115d6848285016116eb565b60c08301525060e06115ea848285016116eb565b60e0830152506101006115ff84828501611228565b6101008301525061012061161584828501611228565b6101208301525061014061162b848285016113f6565b61014083015250610160611641848285016116eb565b6101608301525092915050565b600060a0828403121561166057600080fd5b61166a60a0611b89565b9050600061167a848285016113b7565b600083015250602061168e84828501611228565b60208301525060406116a2848285016116eb565b60408301525060606116b6848285016113e1565b60608301525060806116ca848285016113e1565b60808301525092915050565b6000813590506116e581611dcf565b92915050565b6000815190506116fa81611dcf565b92915050565b60006020828403121561171257600080fd5b60006117208482850161145f565b91505092915050565b60006020828403121561173b57600080fd5b600082013567ffffffffffffffff81111561175557600080fd5b61176184828501611474565b91505092915050565b60006020828403121561177c57600080fd5b600082015167ffffffffffffffff81111561179657600080fd5b6117a284828501611518565b91505092915050565b600060a082840312156117bd57600080fd5b60006117cb8482850161164e565b91505092915050565b6000602082840312156117e657600080fd5b60006117f4848285016116d6565b91505092915050565b6000806040838503121561181057600080fd5b600061181e858286016116d6565b925050602061182f858286016116d6565b9150509250929050565b61184281611c6b565b82525050565b61185961185482611c6b565b611d21565b82525050565b61186881611c7d565b82525050565b61187781611c89565b82525050565b61188e61188982611c89565b611d33565b82525050565b60006118a1601883611c5a565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006118e1601f83611c5a565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611921602283611c5a565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611987602283611c5a565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6119e981611cf1565b82525050565b6119f881611d05565b82525050565b6000611a0a8286611848565b601482019150611a1a828561187d565b602082019150611a2a828461187d565b602082019150819050949350505050565b6000602082019050611a506000830184611839565b92915050565b600060a082019050611a6b600083018861185f565b611a78602083018761185f565b611a856040830186611839565b611a9260608301856119e0565b611a9f60808301846119e0565b9695505050505050565b6000608082019050611abe600083018761186e565b611acb60208301866119ef565b611ad8604083018561186e565b611ae5606083018461186e565b95945050505050565b60006020820190508181036000830152611b0781611894565b9050919050565b60006020820190508181036000830152611b27816118d4565b9050919050565b60006020820190508181036000830152611b4781611914565b9050919050565b60006020820190508181036000830152611b678161197a565b9050919050565b6000602082019050611b8360008301846119e0565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611bac57600080fd5b8060405250919050565b600067ffffffffffffffff821115611bcd57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611bf557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611c1d57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611c4557600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000611c7682611cd1565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000611cca82611c6b565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b6000611d2c82611d3d565b9050919050565b6000819050919050565b6000611d4882611d4f565b9050919050565b60008160601b9050919050565b611d6581611c6b565b8114611d7057600080fd5b50565b611d7c81611c7d565b8114611d8757600080fd5b50565b611d9381611c89565b8114611d9e57600080fd5b50565b611daa81611c93565b8114611db557600080fd5b50565b611dc181611cbf565b8114611dcc57600080fd5b50565b611dd881611cf1565b8114611de357600080fd5b5056fea2646970667358221220da2767aa375a1f716e4e8950e9bf75bc728c4c850f0327d48b2193d5bf141fd564736f6c63430007010033"

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
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherCaller) Accusations(opts *bind.CallOpts, arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
	Executor    common.Address
	HalfStep    uint64
	BlockNumber uint64
}, error) {
	var out []interface{}
	err := _KeyperSlasher.contract.Call(opts, &out, "accusations", arg0)

	outstruct := new(struct {
		Accused     bool
		Appealed    bool
		Executor    common.Address
		HalfStep    uint64
		BlockNumber uint64
	})

	outstruct.Accused = out[0].(bool)
	outstruct.Appealed = out[1].(bool)
	outstruct.Executor = out[2].(common.Address)
	outstruct.HalfStep = out[3].(uint64)
	outstruct.BlockNumber = out[4].(uint64)

	return *outstruct, err

}

// Accusations is a free data retrieval call binding the contract method 0x6864e7ee.
//
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherSession) Accusations(arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
	Executor    common.Address
	HalfStep    uint64
	BlockNumber uint64
}, error) {
	return _KeyperSlasher.Contract.Accusations(&_KeyperSlasher.CallOpts, arg0)
}

// Accusations is a free data retrieval call binding the contract method 0x6864e7ee.
//
// Solidity: function accusations(uint64 ) view returns(bool accused, bool appealed, address executor, uint64 halfStep, uint64 blockNumber)
func (_KeyperSlasher *KeyperSlasherCallerSession) Accusations(arg0 uint64) (struct {
	Accused     bool
	Appealed    bool
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
var MockBatcherContractBin = "0x608060405234801561001057600080fd5b50610289806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ad15b6c51461003b578063c87afa8a14610057575b600080fd5b6100556004803603810190610050919061017e565b610087565b005b610071600480360381019061006c9190610142565b6100de565b60405161007e91906101dc565b60405180910390f35b806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008460018111156100bc57fe5b60018111156100c757fe5b815260200190815260200160002081905550505050565b6000602052816000526040600020602052806000526040600020600091509150505481565b60008135905061011281610215565b92915050565b6000813590506101278161022c565b92915050565b60008135905061013c8161023c565b92915050565b6000806040838503121561015557600080fd5b60006101638582860161012d565b925050602061017485828601610118565b9150509250929050565b60008060006060848603121561019357600080fd5b60006101a18682870161012d565b93505060206101b286828701610118565b92505060406101c386828701610103565b9150509250925092565b6101d6816101f7565b82525050565b60006020820190506101f160008301846101cd565b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b61021e816101f7565b811461022957600080fd5b50565b6002811061023957600080fd5b50565b61024581610201565b811461025057600080fd5b5056fea26469706673582212200a130182f2daa7a4a67d18cda50ab930886b77160958917febe24be77fb988b864736f6c63430007010033"

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
var MockTargetContractBin = "0x608060405234801561001057600080fd5b50610154806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635a6535fc14610030575b600080fd5b6100a76004803603602081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184600183028401116401000000008311171561009757600080fd5b90919293919293905050506100a9565b005b60005a90507fef861dcf69133c2f97e39df733a8a555f2ed9b49b745cc2dd29ae8e06186cf9283838360405180806020018381526020018281038252858582818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a150505056fea264697066735822122061eaf6493c006fac964d6e6d08b684f422e348545204698f8809acfbed21838164736f6c63430007010033"

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
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b1626c29176405d9dbc9e62e225e3bdfacb7c0c2d0290779484cb26a4cf373ab64736f6c63430007010033"

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
var TestDepositTokenContractBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600381526020017f53445400000000000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f5344540000000000000000000000000000000000000000000000000000000000815250600067ffffffffffffffff811180156200009757600080fd5b50604051908082528060200260200182016040528015620000c75781602001602082028036833780820191505090505b508260029080519060200190620000e092919062000814565b508160039080519060200190620000f992919062000814565b508060049080519060200190620001129291906200089b565b5060005b600480549050811015620001c257600160056000600484815481106200013857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808060010191505062000116565b50731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff166329965a1d307fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce2177054306040518463ffffffff1660e01b8152600401620002369392919062000bf7565b600060405180830381600087803b1580156200025157600080fd5b505af115801562000266573d6000803e3d6000fd5b50505050731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff166329965a1d307faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a306040518463ffffffff1660e01b8152600401620002dd9392919062000bf7565b600060405180830381600087803b158015620002f857600080fd5b505af11580156200030d573d6000803e3d6000fd5b505050505050506200034933620f424060405180602001604052806000815250604051806020016040528060008152506200034f60201b60201c565b62000dd8565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415620003c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003b99062000c56565b60405180910390fd5b6000620003d46200059a60201b60201c565b9050620003eb8160008787620005a260201b60201c565b6200040784600154620005a860201b620011421790919060201c565b60018190555062000465846000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620005a860201b620011421790919060201c565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550620004c08160008787878760016200060060201b60201c565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d868686604051620005239392919062000cb7565b60405180910390a38473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040516200058b919062000c9a565b60405180910390a35050505050565b600033905090565b50505050565b600080828401905083811015620005f6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620005ed9062000c34565b60405180910390fd5b8091505092915050565b6000731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff1663aabbb8ca877fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b60001b6040518363ffffffff1660e01b81526004016200067692919062000bca565b60206040518083038186803b1580156200068f57600080fd5b505afa158015620006a4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620006ca91906200099d565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146200077f578073ffffffffffffffffffffffffffffffffffffffff166223de298989898989896040518763ffffffff1660e01b8152600401620007459695949392919062000b4f565b600060405180830381600087803b1580156200076057600080fd5b505af115801562000775573d6000803e3d6000fd5b50505050620007f7565b8115620007f657620007b28673ffffffffffffffffffffffffffffffffffffffff166200080160201b620011971760201c565b15620007f5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007ec9062000c78565b60405180910390fd5b5b5b5050505050505050565b600080823b905060008111915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200085757805160ff191683800117855562000888565b8280016001018555821562000888579182015b82811115620008875782518255916020019190600101906200086a565b5b5090506200089791906200092a565b5090565b82805482825590600052602060002090810192821562000917579160200282015b82811115620009165782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190620008bc565b5b50905062000926919062000949565b5090565b5b80821115620009455760008160009055506001016200092b565b5090565b5b808211156200098257600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016200094a565b5090565b600081519050620009978162000dbe565b92915050565b600060208284031215620009b057600080fd5b6000620009c08482850162000986565b91505092915050565b620009d48162000d2f565b82525050565b620009e58162000d43565b82525050565b6000620009f88262000d02565b62000a04818562000d0d565b935062000a1681856020860162000d77565b62000a218162000dad565b840191505092915050565b600062000a3b601b8362000d1e565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b600062000a7d60208362000d1e565b91507f4552433737373a206d696e7420746f20746865207a65726f20616464726573736000830152602082019050919050565b600062000abf604d8362000d1e565b91507f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460008301527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60208301527f6b656e73526563697069656e74000000000000000000000000000000000000006040830152606082019050919050565b62000b498162000d6d565b82525050565b600060c08201905062000b666000830189620009c9565b62000b756020830188620009c9565b62000b846040830187620009c9565b62000b93606083018662000b3e565b818103608083015262000ba78185620009eb565b905081810360a083015262000bbd8184620009eb565b9050979650505050505050565b600060408201905062000be16000830185620009c9565b62000bf06020830184620009da565b9392505050565b600060608201905062000c0e6000830186620009c9565b62000c1d6020830185620009da565b62000c2c6040830184620009c9565b949350505050565b6000602082019050818103600083015262000c4f8162000a2c565b9050919050565b6000602082019050818103600083015262000c718162000a6e565b9050919050565b6000602082019050818103600083015262000c938162000ab0565b9050919050565b600060208201905062000cb1600083018462000b3e565b92915050565b600060608201905062000cce600083018662000b3e565b818103602083015262000ce28185620009eb565b9050818103604083015262000cf88184620009eb565b9050949350505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600062000d3c8262000d4d565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b8381101562000d9757808201518184015260208101905062000d7a565b8381111562000da7576000848401525b50505050565b6000601f19601f8301169050919050565b62000dc98162000d2f565b811462000dd557600080fd5b50565b612bde8062000de86000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b6371146102e3578063dd62ed3e14610313578063fad8b32a14610343578063fc673c4f1461035f578063fe9d93031461037b57610116565b8063959b8c3f1461025d57806395d89b41146102795780639bd9bbc614610297578063a9059cbb146102b357610116565b806323b872dd116100e957806323b872dd146101a5578063313ce567146101d5578063556f0dc7146101f357806362ad1b831461021157806370a082311461022d57610116565b806306e485381461011b57806306fdde0314610139578063095ea7b31461015757806318160ddd14610187575b600080fd5b610123610397565b6040516101309190612731565b60405180910390f35b610141610425565b60405161014e919061276e565b60405180910390f35b610171600480360381019061016c9190611f2b565b6104c7565b60405161017e9190612753565b60405180910390f35b61018f6104ea565b60405161019c9190612910565b60405180910390f35b6101bf60048036038101906101ba9190611e35565b6104f4565b6040516101cc9190612753565b60405180910390f35b6101dd610726565b6040516101ea9190612970565b60405180910390f35b6101fb61072f565b6040516102089190612910565b60405180910390f35b61022b60048036038101906102269190611e84565b610738565b005b61024760048036038101906102429190611da7565b61079e565b6040516102549190612910565b60405180910390f35b61027760048036038101906102729190611da7565b6107e6565b005b610281610a47565b60405161028e919061276e565b60405180910390f35b6102b160048036038101906102ac9190611f67565b610ae9565b005b6102cd60048036038101906102c89190611f2b565b610b13565b6040516102da9190612753565b60405180910390f35b6102fd60048036038101906102f89190611df9565b610c21565b60405161030a9190612753565b60405180910390f35b61032d60048036038101906103289190611df9565b610dd2565b60405161033a9190612910565b60405180910390f35b61035d60048036038101906103589190611da7565b610e59565b005b61037960048036038101906103749190611fce565b6110ba565b005b61039560048036038101906103909190612061565b61111c565b005b6060600480548060200260200160405190810160405280929190818152602001828054801561041b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103d1575b5050505050905090565b606060028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104bd5780601f10610492576101008083540402835291602001916104bd565b820191906000526020600020905b8154815290600101906020018083116104a057829003601f168201915b5050505050905090565b6000806104d26111aa565b90506104df8185856111b2565b600191505092915050565b6000600154905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055c90612870565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156105d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cc906128b0565b60405180910390fd5b60006105df6111aa565b905061060d81868686604051806020016040528060008152506040518060200160405280600081525061137d565b6106398186868660405180602001604052806000815250604051806020016040528060008152506114f6565b6106ec85826106e786604051806060016040528060298152602001612b5d60299139600860008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461172f9092919063ffffffff16565b6111b2565b61071a818686866040518060200160405280600081525060405180602001604052806000815250600061178a565b60019150509392505050565b60006012905090565b60006001905090565b6107496107436111aa565b86610c21565b610788576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077f90612890565b60405180910390fd5b6107978585858585600161196e565b5050505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b8073ffffffffffffffffffffffffffffffffffffffff166108056111aa565b73ffffffffffffffffffffffffffffffffffffffff16141561085c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085390612810565b60405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561094657600760006108ba6111aa565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556109e3565b6001600660006109546111aa565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6109eb6111aa565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f960405160405180910390a350565b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610adf5780601f10610ab457610100808354040283529160200191610adf565b820191906000526020600020905b815481529060010190602001808311610ac257829003601f168201915b5050505050905090565b610b0e610af46111aa565b84848460405180602001604052806000815250600161196e565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610b84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7b90612870565b60405180910390fd5b6000610b8e6111aa565b9050610bbc81828686604051806020016040528060008152506040518060200160405280600081525061137d565b610be88182868660405180602001604052806000815250604051806020016040528060008152506114f6565b610c16818286866040518060200160405280600081525060405180602001604052806000815250600061178a565b600191505092915050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480610d395750600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168015610d385750600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b5b80610dca5750600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b905092915050565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610e616111aa565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ecf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec690612830565b60405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610fc257600160076000610f2f6111aa565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550611056565b60066000610fce6111aa565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b61105e6111aa565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa160405160405180910390a350565b6110cb6110c56111aa565b85610c21565b61110a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110190612890565b60405180910390fd5b61111684848484611a8e565b50505050565b61113e6111276111aa565b838360405180602001604052806000815250611a8e565b5050565b60008082840190508381101561118d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611184906127d0565b60405180910390fd5b8091505092915050565b600080823b905060008111915050919050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611222576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121990612790565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611292576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611289906128f0565b60405180910390fd5b80600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516113709190612910565b60405180910390a3505050565b6000731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff1663aabbb8ca877f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe89560001b6040518363ffffffff1660e01b81526004016113f1929190612708565b60206040518083038186803b15801561140957600080fd5b505afa15801561141d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114419190611dd0565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146114ed578073ffffffffffffffffffffffffffffffffffffffff166375ab97828888888888886040518763ffffffff1660e01b81526004016114ba96959493929190612699565b600060405180830381600087803b1580156114d457600080fd5b505af11580156114e8573d6000803e3d6000fd5b505050505b50505050505050565b61150286868686611cc4565b61156d83604051806060016040528060278152602001612b36602791396000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461172f9092919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611600836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461114290919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc826146779878686866040516116ba9392919061292b565b60405180910390a48373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161171f9190612910565b60405180910390a3505050505050565b6000838311158290611777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176e919061276e565b60405180910390fd5b5060008385039050809150509392505050565b6000731820a4b7618bde71dce8cdc73aab6c95905fad2473ffffffffffffffffffffffffffffffffffffffff1663aabbb8ca877fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b60001b6040518363ffffffff1660e01b81526004016117fe929190612708565b60206040518083038186803b15801561181657600080fd5b505afa15801561182a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184e9190611dd0565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146118fd578073ffffffffffffffffffffffffffffffffffffffff166223de298989898989896040518763ffffffff1660e01b81526004016118c696959493929190612699565b600060405180830381600087803b1580156118e057600080fd5b505af11580156118f4573d6000803e3d6000fd5b50505050611964565b8115611963576119228673ffffffffffffffffffffffffffffffffffffffff16611197565b15611962576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195990612850565b60405180910390fd5b5b5b5050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156119de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d5906127b0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415611a4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a45906128d0565b60405180910390fd5b6000611a586111aa565b9050611a6881888888888861137d565b611a768188888888886114f6565b611a858188888888888861178a565b50505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611afe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611af5906127f0565b60405180910390fd5b6000611b086111aa565b9050611b178186600087611cc4565b611b268186600087878761137d565b611b9184604051806060016040528060238152602001612b86602391396000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461172f9092919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611be884600154611cca90919063ffffffff16565b6001819055508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098868686604051611c4f9392919061292b565b60405180910390a3600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef86604051611cb59190612910565b60405180910390a35050505050565b50505050565b6000611d0c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061172f565b905092915050565b600081359050611d2381612b07565b92915050565b600081519050611d3881612b07565b92915050565b600082601f830112611d4f57600080fd5b8135611d62611d5d826129b8565b61298b565b91508082526020830160208301858383011115611d7e57600080fd5b611d89838284612ab4565b50505092915050565b600081359050611da181612b1e565b92915050565b600060208284031215611db957600080fd5b6000611dc784828501611d14565b91505092915050565b600060208284031215611de257600080fd5b6000611df084828501611d29565b91505092915050565b60008060408385031215611e0c57600080fd5b6000611e1a85828601611d14565b9250506020611e2b85828601611d14565b9150509250929050565b600080600060608486031215611e4a57600080fd5b6000611e5886828701611d14565b9350506020611e6986828701611d14565b9250506040611e7a86828701611d92565b9150509250925092565b600080600080600060a08688031215611e9c57600080fd5b6000611eaa88828901611d14565b9550506020611ebb88828901611d14565b9450506040611ecc88828901611d92565b935050606086013567ffffffffffffffff811115611ee957600080fd5b611ef588828901611d3e565b925050608086013567ffffffffffffffff811115611f1257600080fd5b611f1e88828901611d3e565b9150509295509295909350565b60008060408385031215611f3e57600080fd5b6000611f4c85828601611d14565b9250506020611f5d85828601611d92565b9150509250929050565b600080600060608486031215611f7c57600080fd5b6000611f8a86828701611d14565b9350506020611f9b86828701611d92565b925050604084013567ffffffffffffffff811115611fb857600080fd5b611fc486828701611d3e565b9150509250925092565b60008060008060808587031215611fe457600080fd5b6000611ff287828801611d14565b945050602061200387828801611d92565b935050604085013567ffffffffffffffff81111561202057600080fd5b61202c87828801611d3e565b925050606085013567ffffffffffffffff81111561204957600080fd5b61205587828801611d3e565b91505092959194509250565b6000806040838503121561207457600080fd5b600061208285828601611d92565b925050602083013567ffffffffffffffff81111561209f57600080fd5b6120ab85828601611d3e565b9150509250929050565b60006120c183836120cd565b60208301905092915050565b6120d681612a55565b82525050565b6120e581612a55565b82525050565b60006120f6826129f4565b6121008185612a22565b935061210b836129e4565b8060005b8381101561213c57815161212388826120b5565b975061212e83612a15565b92505060018101905061210f565b5085935050505092915050565b61215281612a67565b82525050565b61216181612a73565b82525050565b6000612172826129ff565b61217c8185612a33565b935061218c818560208601612ac3565b61219581612af6565b840191505092915050565b60006121ab82612a0a565b6121b58185612a44565b93506121c5818560208601612ac3565b6121ce81612af6565b840191505092915050565b60006121e6602583612a44565b91507f4552433737373a20617070726f76652066726f6d20746865207a65726f20616460008301527f64726573730000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061224c602283612a44565b91507f4552433737373a2073656e642066726f6d20746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006122b2601b83612a44565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006122f2602283612a44565b91507f4552433737373a206275726e2066726f6d20746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000612358602483612a44565b91507f4552433737373a20617574686f72697a696e672073656c66206173206f70657260008301527f61746f72000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006123be602183612a44565b91507f4552433737373a207265766f6b696e672073656c66206173206f70657261746f60008301527f72000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000612424604d83612a44565b91507f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460008301527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60208301527f6b656e73526563697069656e74000000000000000000000000000000000000006040830152606082019050919050565b60006124b0602483612a44565b91507f4552433737373a207472616e7366657220746f20746865207a65726f2061646460008301527f72657373000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000612516602c83612a44565b91507f4552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f60008301527f7220666f7220686f6c64657200000000000000000000000000000000000000006020830152604082019050919050565b600061257c602683612a44565b91507f4552433737373a207472616e736665722066726f6d20746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006125e2602083612a44565b91507f4552433737373a2073656e6420746f20746865207a65726f20616464726573736000830152602082019050919050565b6000612622602383612a44565b91507f4552433737373a20617070726f766520746f20746865207a65726f206164647260008301527f65737300000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b61268481612a9d565b82525050565b61269381612aa7565b82525050565b600060c0820190506126ae60008301896120dc565b6126bb60208301886120dc565b6126c860408301876120dc565b6126d5606083018661267b565b81810360808301526126e78185612167565b905081810360a08301526126fb8184612167565b9050979650505050505050565b600060408201905061271d60008301856120dc565b61272a6020830184612158565b9392505050565b6000602082019050818103600083015261274b81846120eb565b905092915050565b60006020820190506127686000830184612149565b92915050565b6000602082019050818103600083015261278881846121a0565b905092915050565b600060208201905081810360008301526127a9816121d9565b9050919050565b600060208201905081810360008301526127c98161223f565b9050919050565b600060208201905081810360008301526127e9816122a5565b9050919050565b60006020820190508181036000830152612809816122e5565b9050919050565b600060208201905081810360008301526128298161234b565b9050919050565b60006020820190508181036000830152612849816123b1565b9050919050565b6000602082019050818103600083015261286981612417565b9050919050565b60006020820190508181036000830152612889816124a3565b9050919050565b600060208201905081810360008301526128a981612509565b9050919050565b600060208201905081810360008301526128c98161256f565b9050919050565b600060208201905081810360008301526128e9816125d5565b9050919050565b6000602082019050818103600083015261290981612615565b9050919050565b6000602082019050612925600083018461267b565b92915050565b6000606082019050612940600083018661267b565b81810360208301526129528185612167565b905081810360408301526129668184612167565b9050949350505050565b6000602082019050612985600083018461268a565b92915050565b6000604051905081810181811067ffffffffffffffff821117156129ae57600080fd5b8060405250919050565b600067ffffffffffffffff8211156129cf57600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000612a6082612a7d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015612ae1578082015181840152602081019050612ac6565b83811115612af0576000848401525b50505050565b6000601f19601f8301169050919050565b612b1081612a55565b8114612b1b57600080fd5b50565b612b2781612a9d565b8114612b3257600080fd5b5056fe4552433737373a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433737373a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654552433737373a206275726e20616d6f756e7420657863656564732062616c616e6365a26469706673582212205ee0bb58ef7bcb88003221b138a610226485efc750f8a78c0836cff25213ed7b64736f6c63430007010033"

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
