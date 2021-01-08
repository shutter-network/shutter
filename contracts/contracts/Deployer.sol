// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;

import "./ConfigContract.sol";
import "./FeeBankContract.sol";
import "./BatcherContract.sol";
import "./ExecutorContract.sol";
import "./TestDepositTokenContract.sol";
import "./DepositContract.sol";
import "./KeyperSlasher.sol";

contract Deployer {
    event Deployed(
        ConfigContract configContract,
        BatcherContract batcherContract,
        ExecutorContract executorContract,
        FeeBankContract feeBankContract,
        DepositContract depositContract,
        KeyperSlasher keyperSlasher
    );

    constructor(uint64 configChangeHeadsUpBlocks, uint64 appealBlocks) {
        ConfigContract configContract = new ConfigContract(
            configChangeHeadsUpBlocks
        );
        FeeBankContract feeBankContract = new FeeBankContract();
        BatcherContract batcherContract = new BatcherContract(
            configContract,
            feeBankContract
        );
        ExecutorContract executorContract = new ExecutorContract(
            configContract,
            batcherContract
        );


            TestDepositTokenContract depositTokenContract
         = new TestDepositTokenContract();
        DepositContract depositContract = new DepositContract(
            depositTokenContract
        );
        KeyperSlasher keyperSlasher = new KeyperSlasher(
            appealBlocks,
            configContract,
            executorContract,
            depositContract
        );

        configContract.transferOwnership(msg.sender);
        batcherContract.transferOwnership(msg.sender);

        emit Deployed({
            configContract: configContract,
            batcherContract: batcherContract,
            executorContract: executorContract,
            feeBankContract: feeBankContract,
            depositContract: depositContract,
            keyperSlasher: keyperSlasher
        });

        selfdestruct(address(0));
    }
}
