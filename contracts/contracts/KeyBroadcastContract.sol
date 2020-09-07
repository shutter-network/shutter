// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./ConfigContract.sol";

/// @title A contract that keypers can use to publish generated keys.
contract KeyBroadcastContract {
    /// @notice The event emitted when an encryption key is broadcasted.
    /// @notice Note that authenticity of the key is not checked.
    /// @param sender The address of the account that broadcasted the key.
    /// @param batchIndex The index of the batch the key is supposed to encrypt.
    /// @param encryptionKey The encryption key.
    /// @param signerIndices An array of indices corresponding to the keypers who signed the key.
    /// @param signatures An array of signatures by keypers attesting to the key.
    event EncryptionKeyBroadcasted(
        address sender,
        uint256 batchIndex,
        bytes32 encryptionKey,
        uint256[] signerIndices,
        bytes[] signatures
    );

    ConfigContract public configContract;

    constructor(address _configContractAddress) public {
        configContract = ConfigContract(_configContractAddress);
    }

    /// @notice Broadcast an encryption key to the world via an event.
    /// @notice This function only verifies that the caller is a keyper. It does not verify the
    ///     signer indices, the signatures, or the key itself. It also does not prevent
    ///     broadcasting the same or different keys for the same batch multiple times.
    /// @param _keyperIndex The index of the caller in the keyper set.
    /// @param _batchIndex The index of the batch the key is supposed to encrypt.
    /// @param _encryptionKey The encryption key.
    /// @param _signerIndices An array of indices corresponding to the keypers who signed the key.
    /// @param _signatures An array of signatures by keypers attesting to the key.
    function broadcastEncryptionKey(
        uint256 _keyperIndex,
        uint256 _batchIndex,
        bytes32 _encryptionKey,
        uint256[] calldata _signerIndices,
        bytes[] calldata _signatures
    ) public {
        BatchConfig memory _config = configContract.getConfig(_batchIndex);
        require(_keyperIndex < _config.keypers.length);
        require(msg.sender == _config.keypers[_keyperIndex]);

        emit EncryptionKeyBroadcasted(
            msg.sender,
            _batchIndex,
            _encryptionKey,
            _signerIndices,
            _signatures
        );
    }
}
