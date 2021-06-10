// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

contract TargetProxyContract {
    event TransactionForwarded(address receiver, bytes data);

    address public executor;

    constructor(address executorAddress) {
        executor = executorAddress;
    }

    function executeTransaction(bytes memory txData) external {
        require(
            msg.sender == executor,
            "TargetProxyContract: only executor can execute"
        );

        (address receiver, bytes memory data) = abi.decode(
            txData,
            (address, bytes)
        );
        (bool success, ) = receiver.call(data);
        require(success, "TargetProxyContract: call reverted");

        emit TransactionForwarded({receiver: receiver, data: data});
    }
}
