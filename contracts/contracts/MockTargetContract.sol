// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

contract MockTargetContract {
    event Called(bytes transaction, uint256 gas);

    function call(bytes calldata transaction) external {
        uint256 gas = gasleft();
        emit Called(transaction, gas);
    }
}
