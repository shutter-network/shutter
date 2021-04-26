// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

contract TestTargetContract {
    event ExecutedTransaction(address sender, bytes data, uint64 nonce);

    address public executor;
    mapping(address => mapping(uint64 => bool)) private _nonces;

    constructor(address executorAddress) {
        executor = executorAddress;
    }

    function executeTransaction(bytes memory txData) external {
        require(
            msg.sender == executor,
            "TestTargetContract: only executor can execute"
        );

        (bytes memory payload, uint8 v, bytes32 r, bytes32 s) =
            abi.decode(txData, (bytes, uint8, bytes32, bytes32));
        bytes32 payloadHash = keccak256(payload);
        bytes32 signedHash =
            keccak256(
                abi.encodePacked(
                    "\x19Ethereum Signed Message:\n32",
                    payloadHash
                )
            );
        address sender = ecrecover(signedHash, v, r, s);
        (uint64 nonce, bytes memory data) =
            abi.decode(payload, (uint64, bytes));

        require(
            !_nonces[sender][nonce],
            "TestTargetContract: nonce already used"
        );
        _nonces[sender][nonce] = true;

        emit ExecutedTransaction({sender: sender, data: data, nonce: nonce});
    }

    function isNonceUsed(address account, uint64 nonce)
        public
        view
        returns (bool)
    {
        return _nonces[account][nonce];
    }
}
