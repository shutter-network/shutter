# Brownie compiles each solidity source file only when it has changed.
# If contract A depends on contract B, and we change B in an
# incompatible way, A will not be recompiled when we run 'brownie
# test' or 'brownie compile' and it will even run the tests with the
# cached compilation output of A.
#
# This Makefile makes sure we always recompile all contracts.

.POSIX:

SRCS = ./contracts/BatcherContract.sol ./contracts/binding.sol ./contracts/ConfigContract.sol ./contracts/DepositContract.sol ./contracts/ExecutorContract.sol ./contracts/FeeBankContract.sol ./contracts/KeyBroadcastContract.sol ./contracts/KeyperSlasher.sol ./contracts/MockBatcherContract.sol ./contracts/MockTargetContract.sol ./contracts/TestDepositTokenContract.sol ./contracts/TestTargetContract.sol ./openzeppelin/contracts/access/Ownable.sol ./openzeppelin/contracts/utils/cryptography/ECDSA.sol ./openzeppelin/contracts/GSN/Context.sol ./openzeppelin/contracts/introspection/IERC1820Registry.sol ./openzeppelin/contracts/math/SafeMath.sol ./openzeppelin/contracts/token/ERC20/IERC20.sol ./openzeppelin/contracts/token/ERC777/ERC777.sol ./openzeppelin/contracts/token/ERC777/IERC777Recipient.sol ./openzeppelin/contracts/token/ERC777/IERC777Sender.sol ./openzeppelin/contracts/token/ERC777/IERC777.sol ./openzeppelin/contracts/utils/Address.sol

compile: build/compile

build/compile: ${SRCS}
	brownie compile --all
	@touch build/compile

test: compile
	brownie test

clean:
	rm -rf build

.PHONY: compile test clean
