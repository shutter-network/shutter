// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

contract TestProxyReceiver {
    event Called(bytes data);

    fallback() external {
        if (data.length > 0 && data[0] == 0) {
            require(false, "TestTargetContract: failing");
        } else {
            emit Called(data);
        }
    }
}
