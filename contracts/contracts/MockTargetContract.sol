// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;

import "./ConfigContract.sol";

contract MockTargetContract {
    event Called(bytes transaction);

    function call(bytes memory transaction) external {
        emit Called(transaction);
    }
}
