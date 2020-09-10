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
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Keypers                []common.Address
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}

// ConfigContractABI is the input ABI used to generate the binding from.
const ConfigContractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_configChangeHeadsUpBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numConfigs\",\"type\":\"uint256\"}],\"name\":\"ConfigScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numConfigs\",\"type\":\"uint256\"}],\"name\":\"ConfigUnscheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"configChangeHeadsUpBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_configIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keyperIndex\",\"type\":\"uint256\"}],\"name\":\"configKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_configIndex\",\"type\":\"uint256\"}],\"name\":\"configNumKeypers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBatchIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSpan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"executionTimeout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBatchIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"keypers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSpan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"executionTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structBatchConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBatchIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSpan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionSizeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"targetFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"executionTimeout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newKeypers\",\"type\":\"address[]\"}],\"name\":\"nextConfigAddKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"nextConfigKeypers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigNumKeypers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"nextConfigRemoveKeypers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"nextConfigSetActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchSizeLimit\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetBatchSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchSpan\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetBatchSpan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_executionTimeout\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetExecutionTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"nextConfigSetFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBatchIndex\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetStartBatchIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlockNumber\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetStartBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_targetAddress\",\"type\":\"address\"}],\"name\":\"nextConfigSetTargetAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_targetFunctionSelector\",\"type\":\"bytes4\"}],\"name\":\"nextConfigSetTargetFunctionSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionGasLimit\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetTransactionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionSizeLimit\",\"type\":\"uint256\"}],\"name\":\"nextConfigSetTransactionSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduleNextConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fromStartBlockNumber\",\"type\":\"uint256\"}],\"name\":\"unscheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigContractFuncSigs maps the 4-byte function signature to its string representation.
var ConfigContractFuncSigs = map[string]string{
	"cd21aee7": "configChangeHeadsUpBlocks()",
	"d1e27177": "configKeypers(uint256,uint256)",
	"875b341b": "configNumKeypers(uint256)",
	"0098fa22": "configs(uint256)",
	"a81b2f8d": "getConfig(uint256)",
	"64e9f671": "nextConfig()",
	"62fced0e": "nextConfigAddKeypers(address[])",
	"ef89319d": "nextConfigKeypers(uint256)",
	"287447c4": "nextConfigNumKeypers()",
	"763c538d": "nextConfigRemoveKeypers(uint256)",
	"ae6d2e21": "nextConfigSetActive(bool)",
	"c7b91e02": "nextConfigSetBatchSizeLimit(uint256)",
	"43c1b435": "nextConfigSetBatchSpan(uint256)",
	"ae76bf05": "nextConfigSetExecutionTimeout(uint256)",
	"2b2cc6c4": "nextConfigSetFeeReceiver(address)",
	"9d10e7ae": "nextConfigSetStartBatchIndex(uint256)",
	"9ee82110": "nextConfigSetStartBlockNumber(uint256)",
	"bcf67268": "nextConfigSetTargetAddress(address)",
	"d1ac2e52": "nextConfigSetTargetFunctionSelector(bytes4)",
	"cead20fd": "nextConfigSetThreshold(uint256)",
	"f95388c2": "nextConfigSetTransactionGasLimit(uint256)",
	"3c820436": "nextConfigSetTransactionSizeLimit(uint256)",
	"0f0aae6f": "numConfigs()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"18b5e830": "scheduleNextConfig()",
	"f2fde38b": "transferOwnership(address)",
	"ae8708f0": "unscheduleConfigs(uint256)",
}

// ConfigContractBin is the compiled bytecode used for deploying new contracts.
var ConfigContractBin = "0x60a06040523480156200001157600080fd5b5060405162001bd438038062001bd483398101604081905262000034916200032f565b60006200004062000196565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001620000966200019a565b81546001808201845560009384526020938490208351600c909302019182558284015190820155604082015160028201805460ff19169115159190911790556060820151805192939192620000f292600385019201906200021e565b50608082810151600483015560a0830151600583015560c0830151600683015560e080840151600784015561010084015160088401556101208401516009840180546001600160a01b03199081166001600160a01b0393841617909155610140860151600a860180546101608901519316919093161763ffffffff60a01b1916600160a01b9190931c0291909117905561018090920151600b909101555262000348565b3390565b620001a462000288565b50604080516101a08101825260008082526020808301829052828401829052835182815290810190935260608201929092526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915290565b82805482825590600052602060002090810192821562000276579160200282015b828111156200027657825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200023f565b50620002849291506200030e565b5090565b604051806101a00160405280600081526020016000815260200160001515815260200160608152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160e01b0319168152602001600081525090565b5b80821115620002845780546001600160a01b03191681556001016200030f565b60006020828403121562000341578081fd5b5051919050565b60805161186262000372600039806105de52806109e05280610ec952806110b852506118626000f3fe608060405234801561001057600080fd5b50600436106101c35760003560e01c80639ee82110116100f9578063cd21aee711610097578063d1e2717711610071578063d1e271771461038d578063ef89319d146103a0578063f2fde38b146103b3578063f95388c2146103c6576101c3565b8063cd21aee71461035f578063cead20fd14610367578063d1ac2e521461037a576101c3565b8063ae76bf05116100d3578063ae76bf0514610313578063ae8708f014610326578063bcf6726814610339578063c7b91e021461034c576101c3565b80639ee82110146102cd578063a81b2f8d146102e0578063ae6d2e2114610300576101c3565b806362fced0e11610166578063763c538d11610140578063763c538d1461027f578063875b341b146102925780638da5cb5b146102a55780639d10e7ae146102ba576101c3565b806362fced0e1461025c57806364e9f6711461026f578063715018a614610277576101c3565b8063287447c4116101a2578063287447c41461021b5780632b2cc6c4146102235780633c8204361461023657806343c1b43514610249576101c3565b806298fa22146101c85780630f0aae6f146101fc57806318b5e83014610211575b600080fd5b6101db6101d63660046115bc565b6103d9565b6040516101f39c9b9a999897969594939291906117c2565b60405180910390f35b61020461045f565b6040516101f391906117b9565b610219610465565b005b610204610912565b6102196102313660046114d7565b610918565b6102196102443660046115bc565b61096f565b6102196102573660046115bc565b6109a9565b61021961026a366004611505565b610a0f565b6101db610aaf565b610219610af3565b61021961028d3660046115bc565b610b72565b6102046102a03660046115bc565b610c07565b6102ad610c33565b6040516101f39190611659565b6102196102c83660046115bc565b610c42565b6102196102db3660046115bc565b610c7c565b6102f36102ee3660046115bc565b610cb6565b6040516101f391906116e8565b61021961030e366004611574565b610e10565b6102196103213660046115bc565b610e58565b6102196103343660046115bc565b610e92565b6102196103473660046114d7565b611025565b61021961035a3660046115bc565b61107c565b6102046110b6565b6102196103753660046115bc565b6110da565b610219610388366004611594565b611114565b6102ad61039b3660046115d4565b61116d565b6102ad6103ae3660046115bc565b6111b4565b6102196103c13660046114d7565b6111e1565b6102196103d43660046115bc565b611297565b600181815481106103e657fe5b60009182526020909120600c90910201805460018201546002830154600484015460058501546006860154600787015460088801546009890154600a8a0154600b909a0154989a50969860ff90961697949693959294919390926001600160a01b039182169291811691600160a01b90910460e01b908c565b60015490565b61046d6112d1565b6000546001600160a01b039081169116146104a35760405162461bcd60e51b815260040161049a906116b3565b60405180910390fd5b6104ab611357565b6001805460001981019081106104bd57fe5b6000918252602091829020604080516101a081018252600c90930290910180548352600181015483850152600281015460ff161515838301526003810180548351818702810187019094528084529394919360608601939283018282801561054e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610530575b505050918352505060048201546020820152600582015460408201526006820154606082015260078201546080820152600882015460a082015260098201546001600160a01b0390811660c0830152600a83015490811660e0808401919091526001600160e01b0319600160a01b909204901b16610100820152600b9091015461012090910152600354909150437f0000000000000000000000000000000000000000000000000000000000000000011061060857600080fd5b80604001511561064e5780516002541161062157600080fd5b805160025460035460a0840151602085015193909203929183029091011461064857600080fd5b5061065d565b80516002541461065d57600080fd5b60045460ff16156106795760075461067457600080fd5b610686565b6007541561068657600080fd5b60018054808201825560009190915260028054600c9092027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf681019283556003547fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf78201556004547fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf88201805460ff909216151560ff199092169190911790556005805492939261075a927fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf90191906113dd565b50600482810154908201556005808301549082015560068083015490820155600780830154908201556008808301549082015560098083015490820180546001600160a01b03199081166001600160a01b0393841617909155600a808501805491850180549093169190931617808255915463ffffffff600160a01b91829004160263ffffffff60a01b19909216919091179055600b918201549101556107ff6112d5565b8051600290815560208083015160035560408301516004805460ff19169115159190911790556060830151805161083a92600592019061142d565b506080820151600482015560a0820151600582015560c0820151600682015560e080830151600783015561010083015160088301556101208301516009830180546001600160a01b03199081166001600160a01b0393841617909155610140850151600a850180546101608801519316919093161763ffffffff60a01b1916600160a01b9190931c0291909117905561018090910151600b909101556001546040517f70314c9d999fae58774cd9cfa5d4b3bb270ce8d953911e74bbae0b0ad7f4089091610907916117b9565b60405180910390a150565b60055490565b6109206112d1565b6000546001600160a01b0390811691161461094d5760405162461bcd60e51b815260040161049a906116b3565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6109776112d1565b6000546001600160a01b039081169116146109a45760405162461bcd60e51b815260040161049a906116b3565b600955565b6109b16112d1565b6000546001600160a01b039081169116146109de5760405162461bcd60e51b815260040161049a906116b3565b7f00000000000000000000000000000000000000000000000000000000000000008110610a0a57600080fd5b600755565b610a176112d1565b6000546001600160a01b03908116911614610a445760405162461bcd60e51b815260040161049a906116b3565b60005b81811015610aaa576005838383818110610a5d57fe5b9050602002016020810190610a7291906114d7565b815460018082018455600093845260209093200180546001600160a01b0319166001600160a01b039290921691909117905501610a47565b505050565b600254600354600454600654600754600854600954600a54600b54600c54600d5460ff909816976001600160a01b0392831692821691600160a01b900460e01b908c565b610afb6112d1565b6000546001600160a01b03908116911614610b285760405162461bcd60e51b815260040161049a906116b3565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b610b7a6112d1565b6000546001600160a01b03908116911614610ba75760405162461bcd60e51b815260040161049a906116b3565b600554808211610bf75760005b82811015610bf1576005805480610bc757fe5b600082815260209020810160001990810180546001600160a01b0319169055019055600101610bb4565b50610c03565b610c0360056000611482565b5050565b600060018281548110610c1657fe5b90600052602060002090600c02016003018054905090505b919050565b6000546001600160a01b031690565b610c4a6112d1565b6000546001600160a01b03908116911614610c775760405162461bcd60e51b815260040161049a906116b3565b600255565b610c846112d1565b6000546001600160a01b03908116911614610cb15760405162461bcd60e51b815260040161049a906116b3565b600355565b610cbe611357565b600154600019015b600060018281548110610cd557fe5b90600052602060002090600c0201905083816000015411610e0657604080516101a081018252825481526001830154602080830191909152600284015460ff16151582840152600384018054845181840281018401909552808552929385936060860193919291830182828015610d7557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d57575b505050918352505060048201546020820152600582015460408201526006820154606082015260078201546080820152600882015460a082015260098201546001600160a01b0390811660c0830152600a83015490811660e0808401919091526001600160e01b0319600160a01b909204901b16610100820152600b90910154610120909101529250610c2e915050565b5060001901610cc6565b610e186112d1565b6000546001600160a01b03908116911614610e455760405162461bcd60e51b815260040161049a906116b3565b6004805460ff1916911515919091179055565b610e606112d1565b6000546001600160a01b03908116911614610e8d5760405162461bcd60e51b815260040161049a906116b3565b600d55565b610e9a6112d1565b6000546001600160a01b03908116911614610ec75760405162461bcd60e51b815260040161049a906116b3565b7f000000000000000000000000000000000000000000000000000000000000000043018111610ef557600080fd5b60015460001981015b8015610fd957600060018281548110610f1357fe5b90600052602060002090600c0201905083816001015410610fc9576001805480610f3957fe5b600082815260208120600c600019909301928302018181556001810182905560028101805460ff1916905590610f726003830182611482565b50600060048201819055600582018190556006820181905560078201819055600882018190556009820180546001600160a01b0319169055600a820180546001600160c01b0319169055600b909101559055610fcf565b50610fd9565b5060001901610efe565b506001548111610fe857600080fd5b6001546040517f41c1b23350f511849a3106d2288e79ceb1bb67dfa351403abe1917144b2db9ee91611019916117b9565b60405180910390a15050565b61102d6112d1565b6000546001600160a01b0390811691161461105a5760405162461bcd60e51b815260040161049a906116b3565b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b6110846112d1565b6000546001600160a01b039081169116146110b15760405162461bcd60e51b815260040161049a906116b3565b600855565b7f000000000000000000000000000000000000000000000000000000000000000081565b6110e26112d1565b6000546001600160a01b0390811691161461110f5760405162461bcd60e51b815260040161049a906116b3565b600655565b61111c6112d1565b6000546001600160a01b039081169116146111495760405162461bcd60e51b815260040161049a906116b3565b600c805460e09290921c600160a01b0263ffffffff60a01b19909216919091179055565b60006001838154811061117c57fe5b90600052602060002090600c0201600301828154811061119857fe5b6000918252602090912001546001600160a01b03169392505050565b6000600260030182815481106111c657fe5b6000918252602090912001546001600160a01b031692915050565b6111e96112d1565b6000546001600160a01b039081169116146112165760405162461bcd60e51b815260040161049a906116b3565b6001600160a01b03811661123c5760405162461bcd60e51b815260040161049a9061166d565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b61129f6112d1565b6000546001600160a01b039081169116146112cc5760405162461bcd60e51b815260040161049a906116b3565b600a55565b3390565b6112dd611357565b50604080516101a08101825260008082526020808301829052828401829052835182815290810190935260608201929092526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915290565b604051806101a00160405280600081526020016000815260200160001515815260200160608152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160e01b0319168152602001600081525090565b82805482825590600052602060002090810192821561141d5760005260206000209182015b8281111561141d578254825591600101919060010190611402565b506114299291506114a3565b5090565b82805482825590600052602060002090810192821561141d579160200282015b8281111561141d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061144d565b50805460008255906000526020600020908101906114a091906114c2565b50565b5b808211156114295780546001600160a01b03191681556001016114a4565b5b8082111561142957600081556001016114c3565b6000602082840312156114e8578081fd5b81356001600160a01b03811681146114fe578182fd5b9392505050565b60008060208385031215611517578081fd5b823567ffffffffffffffff8082111561152e578283fd5b818501915085601f830112611541578283fd5b81358181111561154f578384fd5b8660208083028501011115611562578384fd5b60209290920196919550909350505050565b600060208284031215611585578081fd5b813580151581146114fe578182fd5b6000602082840312156115a5578081fd5b81356001600160e01b0319811681146114fe578182fd5b6000602082840312156115cd578081fd5b5035919050565b600080604083850312156115e6578182fd5b50508035926020909101359150565b6001600160a01b03169052565b6000815180845260208085019450808401835b8381101561163a5781516001600160a01b031687529582019590820190600101611615565b509495945050505050565b15159052565b6001600160e01b0319169052565b6001600160a01b0391909116815260200190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060208252825160208301526020830151604083015260408301516117116060840182611645565b5060608301516101a080608085015261172e6101c0850183611602565b9150608085015160a085015260a085015160c085015260c085015160e085015260e0850151610100818187015280870151915050610120818187015280870151915050610140611780818701836115f5565b8601519050610160611794868201836115f5565b86015190506101806117a88682018361164b565b959095015193019290925250919050565b90815260200190565b9b8c5260208c019a909a5297151560408b015260608a0196909652608089019490945260a088019290925260c087015260e08601526001600160a01b03908116610100860152166101208401526001600160e01b031916610140830152610160820152610180019056fea2646970667358221220bc729fd402c866430ca9c65d1ab721d21e941cc1b077bdb7f21f2d647bfaf3d264736f6c63430007010033"

// DeployConfigContract deploys a new Ethereum contract, binding an instance of ConfigContract to it.
func DeployConfigContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configChangeHeadsUpBlocks *big.Int) (common.Address, *types.Transaction, *ConfigContract, error) {
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
// Solidity: function configChangeHeadsUpBlocks() view returns(uint256)
func (_ConfigContract *ConfigContractCaller) ConfigChangeHeadsUpBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configChangeHeadsUpBlocks")
	return *ret0, err
}

// ConfigChangeHeadsUpBlocks is a free data retrieval call binding the contract method 0xcd21aee7.
//
// Solidity: function configChangeHeadsUpBlocks() view returns(uint256)
func (_ConfigContract *ConfigContractSession) ConfigChangeHeadsUpBlocks() (*big.Int, error) {
	return _ConfigContract.Contract.ConfigChangeHeadsUpBlocks(&_ConfigContract.CallOpts)
}

// ConfigChangeHeadsUpBlocks is a free data retrieval call binding the contract method 0xcd21aee7.
//
// Solidity: function configChangeHeadsUpBlocks() view returns(uint256)
func (_ConfigContract *ConfigContractCallerSession) ConfigChangeHeadsUpBlocks() (*big.Int, error) {
	return _ConfigContract.Contract.ConfigChangeHeadsUpBlocks(&_ConfigContract.CallOpts)
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xd1e27177.
//
// Solidity: function configKeypers(uint256 _configIndex, uint256 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCaller) ConfigKeypers(opts *bind.CallOpts, _configIndex *big.Int, _keyperIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configKeypers", _configIndex, _keyperIndex)
	return *ret0, err
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xd1e27177.
//
// Solidity: function configKeypers(uint256 _configIndex, uint256 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractSession) ConfigKeypers(_configIndex *big.Int, _keyperIndex *big.Int) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, _configIndex, _keyperIndex)
}

// ConfigKeypers is a free data retrieval call binding the contract method 0xd1e27177.
//
// Solidity: function configKeypers(uint256 _configIndex, uint256 _keyperIndex) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) ConfigKeypers(_configIndex *big.Int, _keyperIndex *big.Int) (common.Address, error) {
	return _ConfigContract.Contract.ConfigKeypers(&_ConfigContract.CallOpts, _configIndex, _keyperIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0x875b341b.
//
// Solidity: function configNumKeypers(uint256 _configIndex) view returns(uint256)
func (_ConfigContract *ConfigContractCaller) ConfigNumKeypers(opts *bind.CallOpts, _configIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "configNumKeypers", _configIndex)
	return *ret0, err
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0x875b341b.
//
// Solidity: function configNumKeypers(uint256 _configIndex) view returns(uint256)
func (_ConfigContract *ConfigContractSession) ConfigNumKeypers(_configIndex *big.Int) (*big.Int, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, _configIndex)
}

// ConfigNumKeypers is a free data retrieval call binding the contract method 0x875b341b.
//
// Solidity: function configNumKeypers(uint256 _configIndex) view returns(uint256)
func (_ConfigContract *ConfigContractCallerSession) ConfigNumKeypers(_configIndex *big.Int) (*big.Int, error) {
	return _ConfigContract.Contract.ConfigNumKeypers(&_ConfigContract.CallOpts, _configIndex)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractCaller) Configs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	ret := new(struct {
		StartBatchIndex        *big.Int
		StartBlockNumber       *big.Int
		Active                 bool
		Threshold              *big.Int
		BatchSpan              *big.Int
		BatchSizeLimit         *big.Int
		TransactionSizeLimit   *big.Int
		TransactionGasLimit    *big.Int
		FeeReceiver            common.Address
		TargetAddress          common.Address
		TargetFunctionSelector [4]byte
		ExecutionTimeout       *big.Int
	})
	out := ret
	err := _ConfigContract.contract.Call(opts, out, "configs", arg0)
	return *ret, err
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractSession) Configs(arg0 *big.Int) (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	return _ConfigContract.Contract.Configs(&_ConfigContract.CallOpts, arg0)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractCallerSession) Configs(arg0 *big.Int) (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	return _ConfigContract.Contract.Configs(&_ConfigContract.CallOpts, arg0)
}

// GetConfig is a free data retrieval call binding the contract method 0xa81b2f8d.
//
// Solidity: function getConfig(uint256 _batchIndex) view returns((uint256,uint256,bool,address[],uint256,uint256,uint256,uint256,uint256,address,address,bytes4,uint256))
func (_ConfigContract *ConfigContractCaller) GetConfig(opts *bind.CallOpts, _batchIndex *big.Int) (BatchConfig, error) {
	var (
		ret0 = new(BatchConfig)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "getConfig", _batchIndex)
	return *ret0, err
}

// GetConfig is a free data retrieval call binding the contract method 0xa81b2f8d.
//
// Solidity: function getConfig(uint256 _batchIndex) view returns((uint256,uint256,bool,address[],uint256,uint256,uint256,uint256,uint256,address,address,bytes4,uint256))
func (_ConfigContract *ConfigContractSession) GetConfig(_batchIndex *big.Int) (BatchConfig, error) {
	return _ConfigContract.Contract.GetConfig(&_ConfigContract.CallOpts, _batchIndex)
}

// GetConfig is a free data retrieval call binding the contract method 0xa81b2f8d.
//
// Solidity: function getConfig(uint256 _batchIndex) view returns((uint256,uint256,bool,address[],uint256,uint256,uint256,uint256,uint256,address,address,bytes4,uint256))
func (_ConfigContract *ConfigContractCallerSession) GetConfig(_batchIndex *big.Int) (BatchConfig, error) {
	return _ConfigContract.Contract.GetConfig(&_ConfigContract.CallOpts, _batchIndex)
}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractCaller) NextConfig(opts *bind.CallOpts) (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	ret := new(struct {
		StartBatchIndex        *big.Int
		StartBlockNumber       *big.Int
		Active                 bool
		Threshold              *big.Int
		BatchSpan              *big.Int
		BatchSizeLimit         *big.Int
		TransactionSizeLimit   *big.Int
		TransactionGasLimit    *big.Int
		FeeReceiver            common.Address
		TargetAddress          common.Address
		TargetFunctionSelector [4]byte
		ExecutionTimeout       *big.Int
	})
	out := ret
	err := _ConfigContract.contract.Call(opts, out, "nextConfig")
	return *ret, err
}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractSession) NextConfig() (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	return _ConfigContract.Contract.NextConfig(&_ConfigContract.CallOpts)
}

// NextConfig is a free data retrieval call binding the contract method 0x64e9f671.
//
// Solidity: function nextConfig() view returns(uint256 startBatchIndex, uint256 startBlockNumber, bool active, uint256 threshold, uint256 batchSpan, uint256 batchSizeLimit, uint256 transactionSizeLimit, uint256 transactionGasLimit, address feeReceiver, address targetAddress, bytes4 targetFunctionSelector, uint256 executionTimeout)
func (_ConfigContract *ConfigContractCallerSession) NextConfig() (struct {
	StartBatchIndex        *big.Int
	StartBlockNumber       *big.Int
	Active                 bool
	Threshold              *big.Int
	BatchSpan              *big.Int
	BatchSizeLimit         *big.Int
	TransactionSizeLimit   *big.Int
	TransactionGasLimit    *big.Int
	FeeReceiver            common.Address
	TargetAddress          common.Address
	TargetFunctionSelector [4]byte
	ExecutionTimeout       *big.Int
}, error) {
	return _ConfigContract.Contract.NextConfig(&_ConfigContract.CallOpts)
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0xef89319d.
//
// Solidity: function nextConfigKeypers(uint256 _index) view returns(address)
func (_ConfigContract *ConfigContractCaller) NextConfigKeypers(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "nextConfigKeypers", _index)
	return *ret0, err
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0xef89319d.
//
// Solidity: function nextConfigKeypers(uint256 _index) view returns(address)
func (_ConfigContract *ConfigContractSession) NextConfigKeypers(_index *big.Int) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, _index)
}

// NextConfigKeypers is a free data retrieval call binding the contract method 0xef89319d.
//
// Solidity: function nextConfigKeypers(uint256 _index) view returns(address)
func (_ConfigContract *ConfigContractCallerSession) NextConfigKeypers(_index *big.Int) (common.Address, error) {
	return _ConfigContract.Contract.NextConfigKeypers(&_ConfigContract.CallOpts, _index)
}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint256)
func (_ConfigContract *ConfigContractCaller) NextConfigNumKeypers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "nextConfigNumKeypers")
	return *ret0, err
}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint256)
func (_ConfigContract *ConfigContractSession) NextConfigNumKeypers() (*big.Int, error) {
	return _ConfigContract.Contract.NextConfigNumKeypers(&_ConfigContract.CallOpts)
}

// NextConfigNumKeypers is a free data retrieval call binding the contract method 0x287447c4.
//
// Solidity: function nextConfigNumKeypers() view returns(uint256)
func (_ConfigContract *ConfigContractCallerSession) NextConfigNumKeypers() (*big.Int, error) {
	return _ConfigContract.Contract.NextConfigNumKeypers(&_ConfigContract.CallOpts)
}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint256)
func (_ConfigContract *ConfigContractCaller) NumConfigs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ConfigContract.contract.Call(opts, out, "numConfigs")
	return *ret0, err
}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint256)
func (_ConfigContract *ConfigContractSession) NumConfigs() (*big.Int, error) {
	return _ConfigContract.Contract.NumConfigs(&_ConfigContract.CallOpts)
}

// NumConfigs is a free data retrieval call binding the contract method 0x0f0aae6f.
//
// Solidity: function numConfigs() view returns(uint256)
func (_ConfigContract *ConfigContractCallerSession) NumConfigs() (*big.Int, error) {
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

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x763c538d.
//
// Solidity: function nextConfigRemoveKeypers(uint256 n) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigRemoveKeypers(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigRemoveKeypers", n)
}

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x763c538d.
//
// Solidity: function nextConfigRemoveKeypers(uint256 n) returns()
func (_ConfigContract *ConfigContractSession) NextConfigRemoveKeypers(n *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigRemoveKeypers(&_ConfigContract.TransactOpts, n)
}

// NextConfigRemoveKeypers is a paid mutator transaction binding the contract method 0x763c538d.
//
// Solidity: function nextConfigRemoveKeypers(uint256 n) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigRemoveKeypers(n *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigRemoveKeypers(&_ConfigContract.TransactOpts, n)
}

// NextConfigSetActive is a paid mutator transaction binding the contract method 0xae6d2e21.
//
// Solidity: function nextConfigSetActive(bool _active) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetActive(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetActive", _active)
}

// NextConfigSetActive is a paid mutator transaction binding the contract method 0xae6d2e21.
//
// Solidity: function nextConfigSetActive(bool _active) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetActive(_active bool) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetActive(&_ConfigContract.TransactOpts, _active)
}

// NextConfigSetActive is a paid mutator transaction binding the contract method 0xae6d2e21.
//
// Solidity: function nextConfigSetActive(bool _active) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetActive(_active bool) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetActive(&_ConfigContract.TransactOpts, _active)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7b91e02.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint256 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSizeLimit(opts *bind.TransactOpts, _batchSizeLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSizeLimit", _batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7b91e02.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint256 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSizeLimit(_batchSizeLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, _batchSizeLimit)
}

// NextConfigSetBatchSizeLimit is a paid mutator transaction binding the contract method 0xc7b91e02.
//
// Solidity: function nextConfigSetBatchSizeLimit(uint256 _batchSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSizeLimit(_batchSizeLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSizeLimit(&_ConfigContract.TransactOpts, _batchSizeLimit)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x43c1b435.
//
// Solidity: function nextConfigSetBatchSpan(uint256 _batchSpan) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetBatchSpan(opts *bind.TransactOpts, _batchSpan *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetBatchSpan", _batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x43c1b435.
//
// Solidity: function nextConfigSetBatchSpan(uint256 _batchSpan) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetBatchSpan(_batchSpan *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, _batchSpan)
}

// NextConfigSetBatchSpan is a paid mutator transaction binding the contract method 0x43c1b435.
//
// Solidity: function nextConfigSetBatchSpan(uint256 _batchSpan) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetBatchSpan(_batchSpan *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetBatchSpan(&_ConfigContract.TransactOpts, _batchSpan)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0xae76bf05.
//
// Solidity: function nextConfigSetExecutionTimeout(uint256 _executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetExecutionTimeout(opts *bind.TransactOpts, _executionTimeout *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetExecutionTimeout", _executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0xae76bf05.
//
// Solidity: function nextConfigSetExecutionTimeout(uint256 _executionTimeout) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetExecutionTimeout(_executionTimeout *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetExecutionTimeout(&_ConfigContract.TransactOpts, _executionTimeout)
}

// NextConfigSetExecutionTimeout is a paid mutator transaction binding the contract method 0xae76bf05.
//
// Solidity: function nextConfigSetExecutionTimeout(uint256 _executionTimeout) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetExecutionTimeout(_executionTimeout *big.Int) (*types.Transaction, error) {
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

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0x9d10e7ae.
//
// Solidity: function nextConfigSetStartBatchIndex(uint256 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBatchIndex(opts *bind.TransactOpts, _startBatchIndex *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBatchIndex", _startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0x9d10e7ae.
//
// Solidity: function nextConfigSetStartBatchIndex(uint256 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBatchIndex(_startBatchIndex *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, _startBatchIndex)
}

// NextConfigSetStartBatchIndex is a paid mutator transaction binding the contract method 0x9d10e7ae.
//
// Solidity: function nextConfigSetStartBatchIndex(uint256 _startBatchIndex) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBatchIndex(_startBatchIndex *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBatchIndex(&_ConfigContract.TransactOpts, _startBatchIndex)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x9ee82110.
//
// Solidity: function nextConfigSetStartBlockNumber(uint256 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetStartBlockNumber(opts *bind.TransactOpts, _startBlockNumber *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetStartBlockNumber", _startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x9ee82110.
//
// Solidity: function nextConfigSetStartBlockNumber(uint256 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetStartBlockNumber(_startBlockNumber *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetStartBlockNumber(&_ConfigContract.TransactOpts, _startBlockNumber)
}

// NextConfigSetStartBlockNumber is a paid mutator transaction binding the contract method 0x9ee82110.
//
// Solidity: function nextConfigSetStartBlockNumber(uint256 _startBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetStartBlockNumber(_startBlockNumber *big.Int) (*types.Transaction, error) {
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

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0xcead20fd.
//
// Solidity: function nextConfigSetThreshold(uint256 _threshold) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetThreshold", _threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0xcead20fd.
//
// Solidity: function nextConfigSetThreshold(uint256 _threshold) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, _threshold)
}

// NextConfigSetThreshold is a paid mutator transaction binding the contract method 0xcead20fd.
//
// Solidity: function nextConfigSetThreshold(uint256 _threshold) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetThreshold(&_ConfigContract.TransactOpts, _threshold)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0xf95388c2.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint256 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionGasLimit(opts *bind.TransactOpts, _transactionGasLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionGasLimit", _transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0xf95388c2.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint256 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionGasLimit(_transactionGasLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, _transactionGasLimit)
}

// NextConfigSetTransactionGasLimit is a paid mutator transaction binding the contract method 0xf95388c2.
//
// Solidity: function nextConfigSetTransactionGasLimit(uint256 _transactionGasLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionGasLimit(_transactionGasLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionGasLimit(&_ConfigContract.TransactOpts, _transactionGasLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x3c820436.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint256 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactor) NextConfigSetTransactionSizeLimit(opts *bind.TransactOpts, _transactionSizeLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "nextConfigSetTransactionSizeLimit", _transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x3c820436.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint256 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractSession) NextConfigSetTransactionSizeLimit(_transactionSizeLimit *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.NextConfigSetTransactionSizeLimit(&_ConfigContract.TransactOpts, _transactionSizeLimit)
}

// NextConfigSetTransactionSizeLimit is a paid mutator transaction binding the contract method 0x3c820436.
//
// Solidity: function nextConfigSetTransactionSizeLimit(uint256 _transactionSizeLimit) returns()
func (_ConfigContract *ConfigContractTransactorSession) NextConfigSetTransactionSizeLimit(_transactionSizeLimit *big.Int) (*types.Transaction, error) {
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

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xae8708f0.
//
// Solidity: function unscheduleConfigs(uint256 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactor) UnscheduleConfigs(opts *bind.TransactOpts, _fromStartBlockNumber *big.Int) (*types.Transaction, error) {
	return _ConfigContract.contract.Transact(opts, "unscheduleConfigs", _fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xae8708f0.
//
// Solidity: function unscheduleConfigs(uint256 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractSession) UnscheduleConfigs(_fromStartBlockNumber *big.Int) (*types.Transaction, error) {
	return _ConfigContract.Contract.UnscheduleConfigs(&_ConfigContract.TransactOpts, _fromStartBlockNumber)
}

// UnscheduleConfigs is a paid mutator transaction binding the contract method 0xae8708f0.
//
// Solidity: function unscheduleConfigs(uint256 _fromStartBlockNumber) returns()
func (_ConfigContract *ConfigContractTransactorSession) UnscheduleConfigs(_fromStartBlockNumber *big.Int) (*types.Transaction, error) {
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
	NumConfigs *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigScheduled is a free log retrieval operation binding the contract event 0x70314c9d999fae58774cd9cfa5d4b3bb270ce8d953911e74bbae0b0ad7f40890.
//
// Solidity: event ConfigScheduled(uint256 numConfigs)
func (_ConfigContract *ConfigContractFilterer) FilterConfigScheduled(opts *bind.FilterOpts) (*ConfigContractConfigScheduledIterator, error) {

	logs, sub, err := _ConfigContract.contract.FilterLogs(opts, "ConfigScheduled")
	if err != nil {
		return nil, err
	}
	return &ConfigContractConfigScheduledIterator{contract: _ConfigContract.contract, event: "ConfigScheduled", logs: logs, sub: sub}, nil
}

// WatchConfigScheduled is a free log subscription operation binding the contract event 0x70314c9d999fae58774cd9cfa5d4b3bb270ce8d953911e74bbae0b0ad7f40890.
//
// Solidity: event ConfigScheduled(uint256 numConfigs)
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

// ParseConfigScheduled is a log parse operation binding the contract event 0x70314c9d999fae58774cd9cfa5d4b3bb270ce8d953911e74bbae0b0ad7f40890.
//
// Solidity: event ConfigScheduled(uint256 numConfigs)
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
	NumConfigs *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUnscheduled is a free log retrieval operation binding the contract event 0x41c1b23350f511849a3106d2288e79ceb1bb67dfa351403abe1917144b2db9ee.
//
// Solidity: event ConfigUnscheduled(uint256 numConfigs)
func (_ConfigContract *ConfigContractFilterer) FilterConfigUnscheduled(opts *bind.FilterOpts) (*ConfigContractConfigUnscheduledIterator, error) {

	logs, sub, err := _ConfigContract.contract.FilterLogs(opts, "ConfigUnscheduled")
	if err != nil {
		return nil, err
	}
	return &ConfigContractConfigUnscheduledIterator{contract: _ConfigContract.contract, event: "ConfigUnscheduled", logs: logs, sub: sub}, nil
}

// WatchConfigUnscheduled is a free log subscription operation binding the contract event 0x41c1b23350f511849a3106d2288e79ceb1bb67dfa351403abe1917144b2db9ee.
//
// Solidity: event ConfigUnscheduled(uint256 numConfigs)
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

// ParseConfigUnscheduled is a log parse operation binding the contract event 0x41c1b23350f511849a3106d2288e79ceb1bb67dfa351403abe1917144b2db9ee.
//
// Solidity: event ConfigUnscheduled(uint256 numConfigs)
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

// KeyBroadcastContractABI is the input ABI used to generate the binding from.
const KeyBroadcastContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_configContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encryptionKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"signerIndices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"EncryptionKeyBroadcasted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keyperIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_encryptionKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_signerIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"broadcastEncryptionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractConfigContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KeyBroadcastContractFuncSigs maps the 4-byte function signature to its string representation.
var KeyBroadcastContractFuncSigs = map[string]string{
	"4e48f976": "broadcastEncryptionKey(uint256,uint256,bytes32,uint256[],bytes[])",
	"bf66a182": "configContract()",
}

// KeyBroadcastContractBin is the compiled bytecode used for deploying new contracts.
var KeyBroadcastContractBin = "0x608060405234801561001057600080fd5b506040516106fd3803806106fd83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b61066c806100916000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634e48f9761461003b578063bf66a18214610050575b600080fd5b61004e610049366004610452565b61006e565b005b61005861018f565b60405161006591906105f2565b60405180910390f35b61007661019e565b60005460405163a81b2f8d60e01b81526001600160a01b039091169063a81b2f8d906100a6908a90600401610606565b60006040518083038186803b1580156100be57600080fd5b505afa1580156100d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100fa919081019061033a565b9050806060015151881061010d57600080fd5b8060600151888151811061011d57fe5b60200260200101516001600160a01b0316336001600160a01b03161461014257600080fd5b7f49cb4b2fc5d10bf803d7f52c3e7c4c6bc3e7c6aa813792f7eec1638fdf8465c73388888888888860405161017d9796959493929190610506565b60405180910390a15050505050505050565b6000546001600160a01b031681565b604051806101a00160405280600081526020016000815260200160001515815260200160608152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160e01b0319168152602001600081525090565b80516001600160a01b038116811461023b57600080fd5b92915050565b600082601f830112610251578081fd5b815167ffffffffffffffff811115610267578182fd5b602080820261027782820161060f565b8381529350818401858301828701840188101561029357600080fd5b600092505b848310156102be576102aa8882610224565b825260019290920191908301908301610298565b505050505092915050565b60008083601f8401126102da578081fd5b50813567ffffffffffffffff8111156102f1578182fd5b602083019150836020808302850101111561030b57600080fd5b9250929050565b8051801515811461023b57600080fd5b80516001600160e01b03198116811461023b57600080fd5b60006020828403121561034b578081fd5b815167ffffffffffffffff80821115610362578283fd5b81840191506101a0808387031215610378578384fd5b6103818161060f565b9050825181526020830151602082015261039e8660408501610312565b60408201526060830151828111156103b4578485fd5b6103c087828601610241565b6060830152506080830151608082015260a083015160a082015260c083015160c082015260e083015160e082015261010091508183015182820152610120915061040c86838501610224565b82820152610140915061042186838501610224565b82820152610160915061043686838501610322565b9181019190915261018091820151918101919091529392505050565b600080600080600080600060a0888a03121561046c578283fd5b873596506020880135955060408801359450606088013567ffffffffffffffff80821115610498578485fd5b6104a48b838c016102c9565b909650945060808a01359150808211156104bc578384fd5b506104c98a828b016102c9565b989b979a50959850939692959293505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b6001600160a01b038816815260208082018890526040820187905260a06060830181905282018590526000906001600160fb1b03861115610545578182fd5b808602808860c0860137830183810360c090810160808601528101859052818502810160e090810190820187855b888110156105df5784840360df190183528135368b9003601e19018112610598578788fd5b8a01803567ffffffffffffffff8111156105b0578889fd5b8036038c13156105be578889fd5b6105cb86828a85016104dc565b955050509185019190850190600101610573565b50919d9c50505050505050505050505050565b6001600160a01b0391909116815260200190565b90815260200190565b60405181810167ffffffffffffffff8111828210171561062e57600080fd5b60405291905056fea26469706673582212203dfcd478b98af3dd87845f5be234cf44f7ac51d4f65bd88d2816b017b0237be964736f6c63430007010033"

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

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x4e48f976.
//
// Solidity: function broadcastEncryptionKey(uint256 _keyperIndex, uint256 _batchIndex, bytes32 _encryptionKey, uint256[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactor) BroadcastEncryptionKey(opts *bind.TransactOpts, _keyperIndex *big.Int, _batchIndex *big.Int, _encryptionKey [32]byte, _signerIndices []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.contract.Transact(opts, "broadcastEncryptionKey", _keyperIndex, _batchIndex, _encryptionKey, _signerIndices, _signatures)
}

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x4e48f976.
//
// Solidity: function broadcastEncryptionKey(uint256 _keyperIndex, uint256 _batchIndex, bytes32 _encryptionKey, uint256[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractSession) BroadcastEncryptionKey(_keyperIndex *big.Int, _batchIndex *big.Int, _encryptionKey [32]byte, _signerIndices []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.BroadcastEncryptionKey(&_KeyBroadcastContract.TransactOpts, _keyperIndex, _batchIndex, _encryptionKey, _signerIndices, _signatures)
}

// BroadcastEncryptionKey is a paid mutator transaction binding the contract method 0x4e48f976.
//
// Solidity: function broadcastEncryptionKey(uint256 _keyperIndex, uint256 _batchIndex, bytes32 _encryptionKey, uint256[] _signerIndices, bytes[] _signatures) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactorSession) BroadcastEncryptionKey(_keyperIndex *big.Int, _batchIndex *big.Int, _encryptionKey [32]byte, _signerIndices []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
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
	BatchIndex    *big.Int
	EncryptionKey [32]byte
	SignerIndices []*big.Int
	Signatures    [][]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEncryptionKeyBroadcasted is a free log retrieval operation binding the contract event 0x49cb4b2fc5d10bf803d7f52c3e7c4c6bc3e7c6aa813792f7eec1638fdf8465c7.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint256 batchIndex, bytes32 encryptionKey, uint256[] signerIndices, bytes[] signatures)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) FilterEncryptionKeyBroadcasted(opts *bind.FilterOpts) (*KeyBroadcastContractEncryptionKeyBroadcastedIterator, error) {

	logs, sub, err := _KeyBroadcastContract.contract.FilterLogs(opts, "EncryptionKeyBroadcasted")
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractEncryptionKeyBroadcastedIterator{contract: _KeyBroadcastContract.contract, event: "EncryptionKeyBroadcasted", logs: logs, sub: sub}, nil
}

// WatchEncryptionKeyBroadcasted is a free log subscription operation binding the contract event 0x49cb4b2fc5d10bf803d7f52c3e7c4c6bc3e7c6aa813792f7eec1638fdf8465c7.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint256 batchIndex, bytes32 encryptionKey, uint256[] signerIndices, bytes[] signatures)
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

// ParseEncryptionKeyBroadcasted is a log parse operation binding the contract event 0x49cb4b2fc5d10bf803d7f52c3e7c4c6bc3e7c6aa813792f7eec1638fdf8465c7.
//
// Solidity: event EncryptionKeyBroadcasted(address sender, uint256 batchIndex, bytes32 encryptionKey, uint256[] signerIndices, bytes[] signatures)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) ParseEncryptionKeyBroadcasted(log types.Log) (*KeyBroadcastContractEncryptionKeyBroadcasted, error) {
	event := new(KeyBroadcastContractEncryptionKeyBroadcasted)
	if err := _KeyBroadcastContract.contract.UnpackLog(event, "EncryptionKeyBroadcasted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

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
