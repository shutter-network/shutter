// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

contract TestTargetContract {
    event ExecutedTransaction(address sender, bytes data, uint64 nonce);

    address private _executor;
    mapping(address => uint64) private _nonces;

    constructor(address executor) {
        _executor = executor;
    }

    function executeTransaction(bytes memory txData) external {
        require(
            msg.sender == _executor,
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

        require(nonce == _nonces[sender], "TestTargetContract: wrong nonce");
        _nonces[sender] = nonce + 1;

        emit ExecutedTransaction({sender: sender, data: data, nonce: nonce});
    }

    function getNonce(address account) public view returns (uint64) {
        return _nonces[account];
    }

    function getExecutor() public view returns (address) {
        return _executor;
    }
}
