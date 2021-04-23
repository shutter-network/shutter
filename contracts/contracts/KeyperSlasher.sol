// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;
pragma experimental ABIEncoderV2;

import {ECDSA} from "openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {ConfigContract, BatchConfig} from "./ConfigContract.sol";
import {DepositContract} from "./DepositContract.sol";
import {ExecutorContract, CipherExecutionReceipt} from "./ExecutorContract.sol";
import {BatcherContract} from "./BatcherContract.sol";

struct Accusation {
    bool accused;
    bool appealed;
    bool slashed;
    address executor;
    uint64 halfStep;
    uint64 blockNumber;
}

struct Authorization {
    uint64 halfStep;
    bytes32 batchHash;
    uint64[] signerIndices;
    bytes[] signatures;
}

contract KeyperSlasher {
    event Accused(
        uint64 indexed halfStep,
        address indexed executor,
        address indexed accuser
    );
    event Appealed(uint64 indexed halfStep, address indexed executor);
    event Slashed(uint64 indexed halfStep, address indexed executor);

    ConfigContract public configContract;
    ExecutorContract public executorContract;
    DepositContract public depositContract;

    uint256 public appealBlocks;

    mapping(uint64 => Accusation) public accusations;

    constructor(
        uint256 appealPeriod,
        ConfigContract configContractAddress,
        ExecutorContract executorContractAddress,
        DepositContract depositContractAddress
    ) {
        appealBlocks = appealPeriod;

        configContract = configContractAddress;
        executorContract = executorContractAddress;
        depositContract = depositContractAddress;

        depositContract.setSlasher(address(this));
    }

    function accuse(uint64 halfStep, uint64 keyperIndex) external {
        require(halfStep % 2 == 0, "KeyperSlasher: not a cipher half step");

        require(
            !accusations[halfStep].accused,
            "KeyperSlasher: already accused"
        );

        BatchConfig memory config = configContract.getConfig(halfStep / 2);
        require(
            keyperIndex < config.keypers.length,
            "KeyperSlasher: keyper index out of range"
        );
        require(
            msg.sender == config.keypers[keyperIndex],
            "KeyperSlasher: sender does not match keyper"
        );

        CipherExecutionReceipt memory receipt =
            executorContract.getReceipt(halfStep);
        require(receipt.executed, "KeyperSlasher: half step not yet executed");
        require(
            receipt.cipherBatchHash != bytes32(0),
            "KeyperSlasher: cannot accuse empty batch"
        );
        accusations[halfStep] = Accusation({
            accused: true,
            appealed: false,
            slashed: false,
            executor: receipt.executor,
            halfStep: halfStep,
            blockNumber: uint64(block.number)
        });

        emit Accused({
            halfStep: halfStep,
            executor: receipt.executor,
            accuser: msg.sender
        });
    }

    function appeal(Authorization memory authorization) external {
        Accusation memory accusation = accusations[authorization.halfStep];
        require(accusation.accused, "KeyperSlasher: no accusation");
        require(!accusation.appealed, "KeyperSlasher: already appealed");
        CipherExecutionReceipt memory receipt =
            executorContract.getReceipt(authorization.halfStep);

        _verifyAuthorization(authorization, receipt);

        accusation.appealed = true;
        accusations[authorization.halfStep] = accusation;

        emit Appealed({
            halfStep: authorization.halfStep,
            executor: receipt.executor
        });
    }

    function slash(uint64 halfStep) external {
        Accusation memory accusation = accusations[halfStep];
        require(accusation.accused, "KeyperSlasher: no accusation");
        require(!accusation.appealed, "KeyperSlasher: successfully appealed");
        require(!accusation.slashed, "KeyperSlasher: already slashed");
        require(
            block.number >= accusation.blockNumber + appealBlocks,
            "KeyperSlasher: appeal period not over yet"
        );

        depositContract.slash(accusation.executor);
        accusations[halfStep].slashed = true;

        emit Slashed({
            halfStep: accusation.halfStep,
            executor: accusation.executor
        });
    }

    function _verifyAuthorization(
        Authorization memory authorization,
        CipherExecutionReceipt memory receipt
    ) internal view {
        BatchConfig memory config =
            configContract.getConfig(receipt.halfStep / 2);

        require(
            authorization.signatures.length >= config.threshold,
            "KeyperSlasher: not enough signatures"
        );
        require(
            authorization.signatures.length ==
                authorization.signerIndices.length,
            "KeyperSlasher: number of signatures and indices does not match"
        );
        bytes32 decryptionSignatureHash =
            keccak256(
                _decryptionSignaturePreimage(
                    address(executorContract.batcherContract()),
                    receipt
                )
            );
        for (uint64 i = 0; i < authorization.signatures.length; i++) {
            bytes memory signature = authorization.signatures[i];
            uint64 signerIndex = authorization.signerIndices[i];
            require(
                signerIndex < config.keypers.length,
                "KeyperSlasher: signer index out of range"
            );

            // Check order as a simple way to check for duplicates
            require(
                i == 0 || signerIndex > authorization.signerIndices[i - 1],
                "KeyperSlasher: signer indices not ordered"
            );

            address signer = ECDSA.recover(decryptionSignatureHash, signature);
            require(
                signer == config.keypers[signerIndex],
                "KeyperSlasher: wrong signer"
            );
        }
    }

    function _decryptionSignaturePreimage(
        address batcherContract,
        CipherExecutionReceipt memory receipt
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(
                "\x19dectx",
                batcherContract,
                receipt.halfStep / 2, // uint64
                receipt.cipherBatchHash,
                receipt.batchHash
            );
    }
}
