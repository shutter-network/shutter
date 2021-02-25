// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import {ConfigContract, BatchConfig} from "./ConfigContract.sol";

/// @title A contract that keypers can use to vote on eon public keys. For each eon public key
///     generated, the keypers are expected to submit one vote. The contract logs the number of
///     votes so that users can only pick keys once they have reached votes from enough keypers
///     and thus have confidence that the key is actually correct.
contract KeyBroadcastContract {
    /// @notice The event emitted when a keyper voted on an eon key.
    /// @param keyper The address of the keyper who sent the vote.
    /// @param startBatchIndex The index of the first batch for which the key should be used.
    /// @param key The eon public key for which the keyper voted.
    /// @param numVotes The number of keypers (including this one) who have voted for the key so
    ///     far.
    event Voted(
        address indexed keyper,
        uint64 startBatchIndex,
        bytes key,
        uint64 numVotes
    );

    ConfigContract private _configContract;
    mapping(uint64 => mapping(address => bool)) private _voted; // start batch index => keyper => voted or not
    mapping(uint64 => mapping(bytes32 => uint64)) private _numVotes; // start batch index => key hash => number of votes
    mapping(bytes32 => bytes) private _keys; // key hash => key

    mapping(uint64 => bytes32) private _bestKeyHashes;
    mapping(uint64 => uint64) private _bestKeyNumVotes;

    constructor(address configContractAddress) {
        _configContract = ConfigContract(configContractAddress);
    }

    /// @notice Submit a vote.
    /// @notice Can only be called by keypers defined in the config responsible for
    ///     `startBatchIndex`, and only once per `startBatchIndex`.
    /// @param keyperIndex The index of the calling keyper in the batch config.
    /// @param startBatchIndex The index of the first batch for which the key should be used.
    /// @param key The eon public key to vote for.
    function vote(
        uint64 keyperIndex,
        uint64 startBatchIndex,
        bytes memory key
    ) public {
        BatchConfig memory config = _configContract.getConfig(startBatchIndex);
        require(
            config.batchSpan > 0,
            "KeyBroadcastContract: config is inactive"
        );

        require(
            keyperIndex < config.keypers.length,
            "KeyBroadcastContract: keyper index out of range"
        );
        require(
            msg.sender == config.keypers[keyperIndex],
            "KeyBroadcastContract: sender is not keyper"
        );

        require(
            !_voted[startBatchIndex][msg.sender],
            "KeyBroadcastContract: keyper has already voted"
        );

        bytes32 keyHash = keccak256(key);
        // store the key if it hasn't already
        if (_keys[keyHash].length == 0 && key.length >= 0) {
            _keys[keyHash] = key;
        }

        // count vote
        uint64 numVotes = _numVotes[startBatchIndex][keyHash] + 1;
        _voted[startBatchIndex][msg.sender] = true;
        _numVotes[startBatchIndex][keyHash] = numVotes;

        if (numVotes > _bestKeyNumVotes[startBatchIndex]) {
            _bestKeyNumVotes[startBatchIndex] = numVotes;
            _bestKeyHashes[startBatchIndex] = keyHash;
        }

        emit Voted({
            keyper: msg.sender,
            startBatchIndex: startBatchIndex,
            key: key,
            numVotes: numVotes
        });
    }

    function getConfigContract() public view returns (ConfigContract) {
        return _configContract;
    }

    function hasVoted(address keyper, uint64 startBatchIndex)
        public
        view
        returns (bool)
    {
        return _voted[startBatchIndex][keyper];
    }

    function getNumVotes(uint64 startBatchIndex, bytes memory key)
        public
        view
        returns (uint64)
    {
        return _numVotes[startBatchIndex][keccak256(key)];
    }

    function getBestKeyHash(uint64 startBatchIndex)
        public
        view
        returns (bytes32)
    {
        return _bestKeyHashes[startBatchIndex];
    }

    function getBestKey(uint64 startBatchIndex)
        public
        view
        returns (bytes memory)
    {
        return _keys[_bestKeyHashes[startBatchIndex]];
    }

    function getBestKeyNumVotes(uint64 startBatchIndex)
        public
        view
        returns (uint256)
    {
        return _bestKeyNumVotes[startBatchIndex];
    }
}
