// SPDX-License-Identifier: MIT

// This file is the sole input file we use when running abigen to generate the bindings in shuttermint/contract/binding.go

// Please keep the version in sync with ../brownie-config.yaml
pragma solidity =0.7.6;

import "./BatcherContract.sol";
import "./ConfigContract.sol";
import "./DepositContract.sol";
import "./ExecutorContract.sol";
import "./FeeBankContract.sol";
import "./KeyBroadcastContract.sol";
import "./KeyperSlasher.sol";

import "./TestDepositTokenContract.sol";
import "./TestTargetContract.sol";

import "./MockBatcherContract.sol";
import "./MockTargetContract.sol";
