// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/cryptography/ECDSA.sol";
import "./ConfigContract.sol";
import "./DepositContract.sol";
import "./ExecutorContract.sol";
import "./BatcherContract.sol";

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

    ConfigContract configContract;
    ExecutorContract executorContract;
    DepositContract depositContract;

    uint256 appealBlocks;

    mapping(uint64 => Accusation) public accusations;

    constructor(
        uint256 _appealBlocks,
        ConfigContract _configContract,
        ExecutorContract _executorContract,
        DepositContract _depositContract
    ) {
        appealBlocks = _appealBlocks;

        configContract = _configContract;
        executorContract = _executorContract;
        depositContract = _depositContract;

        depositContract.setSlasher(address(this));
    }

    function accuse(uint64 _halfStep, uint64 _keyperIndex) external {
        require(_halfStep % 2 == 0);

        require(!accusations[_halfStep].accused);

        BatchConfig memory _config = configContract.getConfig(_halfStep / 2);
        require(_keyperIndex < _config.keypers.length);
        require(msg.sender == _config.keypers[_keyperIndex]);

        CipherExecutionReceipt memory _receipt = executorContract.getReceipt(
            _halfStep
        );
        require(_receipt.executed);

        accusations[_halfStep] = Accusation({
            accused: true,
            appealed: false,
            slashed: false,
            executor: _receipt.executor,
            halfStep: _halfStep,
            blockNumber: uint64(block.number)
        });

        emit Accused({
            halfStep: _halfStep,
            executor: _receipt.executor,
            accuser: msg.sender
        });
    }

    function appeal(Authorization memory _authorization) external {
        Accusation memory _accusation = accusations[_authorization.halfStep];
        require(_accusation.accused);
        require(!_accusation.appealed);
        CipherExecutionReceipt memory _receipt = executorContract.getReceipt(
            _authorization.halfStep
        );

        verifyAuthorization(_authorization, _receipt);

        _accusation.appealed = true;
        accusations[_authorization.halfStep] = _accusation;

        emit Appealed({
            halfStep: _authorization.halfStep,
            executor: _receipt.executor
        });
    }

    function verifyAuthorization(
        Authorization memory _authorization,
        CipherExecutionReceipt memory _receipt
    ) internal view {
        BatchConfig memory _config = configContract.getConfig(
            _receipt.halfStep / 2
        );

        require(_authorization.signatures.length >= _config.threshold);
        require(
            _authorization.signatures.length ==
                _authorization.signerIndices.length
        );
        bytes32 _decryptionSignatureHash = keccak256(
            abi.encodePacked(
                address(executorContract.batcherContract()),
                _receipt.cipherBatchHash,
                _receipt.batchHash
            )
        );
        for (uint64 _i = 0; _i < _authorization.signatures.length; _i++) {
            bytes memory _signature = _authorization.signatures[_i];
            uint64 _signerIndex = _authorization.signerIndices[_i];

            // Check order as a simple way to check for duplicates
            require(
                _i == 0 || _signerIndex > _authorization.signerIndices[_i - 1]
            );

            address _signer = ECDSA.recover(
                _decryptionSignatureHash,
                _signature
            );
            require(_signer == _config.keypers[_signerIndex]);
        }
    }

    function slash(uint64 _halfStep) external {
        Accusation memory _accusation = accusations[_halfStep];
        require(_accusation.accused);
        require(!_accusation.appealed);
        require(!_accusation.slashed);
        require(block.number >= _accusation.blockNumber + appealBlocks);

        depositContract.slash(_accusation.executor);
        accusations[_halfStep].slashed = true;

        emit Slashed({
            halfStep: _accusation.halfStep,
            executor: _accusation.executor
        });
    }
}
