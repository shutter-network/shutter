// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./ConfigContract.sol";
import "./ExecutorContract.sol";
import "./BatcherContract.sol";

struct Accusation {
    bool accused;
    bool appealed;
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
    // DepositContract depositContract;
    ExecutorContract executorContract;

    uint256 appealBlocks;

    mapping(uint64 => Accusation) public accusations;

    constructor(
        uint256 _appealBlocks,
        ConfigContract _configContract,
        ExecutorContract _executorContract
    ) {
        appealBlocks = _appealBlocks;

        configContract = _configContract;
        executorContract = _executorContract;
    }

    function accuse(uint64 _halfStep, uint64 _keyperIndex) external {
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
        require(block.number >= _accusation.blockNumber + appealBlocks);

        // depositContract.slash(_accusation.executor);

        emit Slashed({
            halfStep: _accusation.halfStep,
            executor: _accusation.executor
        });
    }
}

// struct Deposit {
//     bool deposited;
//     address depositor;
//     uint256 amount;
//     uint256 withdrawalDelayBlocks;
//     uint256 withdrawalInitiatedBlock;
// }

// contract DepositContract {

//     event Deposited(
//         address depositor,
//         uint256 amount,
//         uint256 withdrawalDelayBlocks
//     );
//     event WithdrawalInitiated(
//         address depositor,
//         uint256 withdrawalInitiatedBlock,
//         uint256 withdrawalDelayBlocks
//     );
//     event Withdrawn(
//         address depositor,
//         address withdrawalAddress,
//         uint256 amount
//     );
//     event Slashed(
//         address depositor
//     );

//     KeyperSlasher slasher;

//     mapping(address => Deposit) deposits;

//     function deposit(uint256 _withdrawalDelayBlocks) external payable {
//         require(!deposits[msg.sender].deposited);
//         require(msg.value > 0);
//         deposits[msg.sender] = Deposit({
//             deposited: true,
//             depositor: msg.sender,
//             amount: msg.value,
//             withdrawalDelayBlocks: _withdrawalDelayBlocks,
//             withdrawalInitiatedBlock: 0
//         });

//         emit Deposited({
//             depositor: msg.sender,
//             amount: msg.value,
//             withdrawalDelayBlocks: _withdrawalDelayBlocks
//         });
//     }

//     function initiateWithdrawal() external {
//         require(deposits[msg.sender].deposited);
//         require(deposits[msg.sender].withdrawalInitiatedBlock == 0);
//         deposits[msg.sender].withdrawalInitiatedBlock = block.number;

//         emit WithdrawalInitiated({
//             depositor: msg.sender,
//             withdrawalInitiatedBlock: block.number,
//             withdrawalDelayBlocks: deposits[msg.sender].withdrawalDelayBlocks
//         });
//     }

//     function withdraw(address _address) external {
//         require(deposits[msg.sender].deposited);
//         require(block.number >= deposits[msg.sender].withdrawalInitiatedBlock + deposits[msg.sender].withdrawalDelayBlocks);
//         uint256 _amount = deposits[msg.sender].amount;
//         delete deposits[msg.sender];
//         (bool success, ) = _address.call{value: _amount}("");
//         require(success);

//         emit Withdrawn({
//             depositor: msg.sender,
//             withdrawalAddress: _address,
//             amount: _amount
//         });
//     }

//     function slash(address _address) external {
//         require(!deposits[_address].deposited);
//         require(msg.sender == address(slasher));
//         delete deposits[_address];

//         emit Slashed({
//             depositor: _address
//         });
//     }
// }

// ,
//         bytes32 _decryptionKey,
//         uint64[] calldata _signerIndices,
//         bytes[] calldata _signatures

//         // Check the signatures (can only be done after execution as we need the batch hash)
//         require(_signatures.length >= _config.threshold);
//         require(_signatures.length == _signerIndices.length);
//         bytes32 _decryptionSignaturePreimage = keccak256(
//             abi.encodePacked(
//                 address(batcherContract),
//                 _cipherBatchHash,
//                 _decryptionKey,
//                 _batchHash
//             )
//         );
//         for (uint64 _i = 0; _i < _signatures.length; _i++) {
//             bytes calldata _signature = _signatures[_i];
//             uint64 _signerIndex = _signerIndices[_i];

//             // Check order to easily prevent duplicates
//             require(_i == 0 || _signerIndex > _signerIndices[_i - 1]);

//             address _signer = ECDSA.recover(
//                 _decryptionSignaturePreimage,
//                 _signature
//             );
//             require(_signer == _config.keypers[_signerIndex]);
//         }
