// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import {
    IERC777
} from "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/token/ERC777/IERC777.sol";
import {
    IERC777Recipient
} from "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/token/ERC777/IERC777Recipient.sol";
import {
    IERC1820Registry
} from "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/introspection/IERC1820Registry.sol";

struct Deposit {
    uint256 amount;
    uint64 withdrawalDelayBlocks;
    uint64 withdrawalRequestedBlock;
    bool slashed;
}

contract DepositContract is IERC777Recipient {
    event DepositChanged(
        address indexed account,
        uint256 amount,
        uint64 withdrawalDelayBlocks,
        uint64 withdrawalRequestedBlock,
        bool withdrawn,
        bool slashed
    );

    IERC1820Registry private _erc1820 = IERC1820Registry(
        0x1820a4B7618BdE71Dce8cdc73aAB6C95905faD24
    );
    bytes32 private constant TOKENS_RECIPIENT_INTERFACE_HASH = keccak256(
        "ERC777TokensRecipient"
    );

    IERC777 private _token;
    address private _slasher;
    mapping(address => Deposit) private _deposits;

    constructor(IERC777 token) {
        _token = token;

        _erc1820.setInterfaceImplementer(
            address(this),
            TOKENS_RECIPIENT_INTERFACE_HASH,
            address(this)
        );
    }

    function setSlasher(address slasherAddress) external {
        require(
            _slasher == address(0),
            "DepositContract: slasher address already set"
        );
        _slasher = slasherAddress;
    }

    function tokensReceived(
        address operator,
        address from,
        address to,
        uint256 amount,
        bytes calldata userData,
        bytes calldata operatorData
    ) external override {
        require(
            msg.sender == address(_token),
            "DepositContract: received invalid token"
        );
        require(
            userData.length == 8,
            "DepositContract: invalid user data size"
        );
        uint64 withdrawalInterval = _bytesToUint64(userData);
        _deposit(from, amount, withdrawalInterval);
    }

    function requestWithdrawal() external {
        Deposit memory deposit = _deposits[msg.sender];
        require(deposit.amount > 0, "DepositContract: no deposit");
        assert(!deposit.slashed);
        require(
            deposit.withdrawalRequestedBlock == 0,
            "DepositContract: withdrawal already requested"
        );

        _deposits[msg.sender].withdrawalRequestedBlock = uint64(block.number);

        emit DepositChanged({
            account: msg.sender,
            amount: deposit.amount,
            withdrawalDelayBlocks: deposit.withdrawalDelayBlocks,
            withdrawalRequestedBlock: uint64(block.number),
            withdrawn: false,
            slashed: false
        });
    }

    function withdraw(address recipient) external {
        Deposit memory deposit = _deposits[msg.sender];
        require(deposit.amount > 0, "DepositContract: no deposit");
        assert(!deposit.slashed);
        require(
            deposit.withdrawalRequestedBlock > 0,
            "DepositContract: withdrawal not requested yet"
        );
        require(
            block.number >=
                deposit.withdrawalRequestedBlock +
                    deposit.withdrawalDelayBlocks,
            "DepositContract: withdrawal delay not passed yet"
        );

        delete _deposits[msg.sender];
        _token.send(recipient, deposit.amount, "");

        emit DepositChanged({
            account: msg.sender,
            amount: 0,
            withdrawalDelayBlocks: 0,
            withdrawalRequestedBlock: 0,
            withdrawn: true,
            slashed: false
        });
    }

    function slash(address account) external {
        require(msg.sender == _slasher);

        Deposit memory deposit = _deposits[account];

        deposit.amount = 0;
        deposit.withdrawalDelayBlocks = 0;
        deposit.withdrawalRequestedBlock = 0;
        deposit.slashed = true;

        _deposits[account] = deposit;

        emit DepositChanged({
            account: account,
            amount: 0,
            withdrawalDelayBlocks: 0,
            withdrawalRequestedBlock: 0,
            withdrawn: false,
            slashed: true
        });
    }

    function getDepositAmount(address account) public view returns (uint256) {
        return _deposits[account].amount;
    }

    function getWithdrawalDelayBlocks(address account)
        public
        view
        returns (uint64)
    {
        return _deposits[account].withdrawalDelayBlocks;
    }

    function getWithdrawalRequestedBlock(address account)
        public
        view
        returns (uint64)
    {
        return _deposits[account].withdrawalRequestedBlock;
    }

    function isSlashed(address account) public view returns (bool) {
        return _deposits[account].slashed;
    }

    function _deposit(
        address depositor,
        uint256 amount,
        uint64 withdrawalDelayBlocks
    ) internal {
        Deposit memory deposit = _deposits[depositor];
        require(
            withdrawalDelayBlocks >= deposit.withdrawalDelayBlocks,
            "DepositContract: withdrawal delay cannot be decreased"
        );
        require(
            deposit.withdrawalRequestedBlock == 0,
            "DepositContract: withdrawal in progress"
        );
        require(!deposit.slashed, "DepositContract: account slashed");

        _deposits[depositor].amount = deposit.amount + amount;
        _deposits[depositor].withdrawalDelayBlocks = withdrawalDelayBlocks;

        emit DepositChanged({
            account: depositor,
            amount: deposit.amount + amount,
            withdrawalDelayBlocks: withdrawalDelayBlocks,
            withdrawalRequestedBlock: 0,
            withdrawn: false,
            slashed: false
        });
    }

    function _bytesToUint64(bytes memory b) internal pure returns (uint64) {
        uint64 number;
        for (uint256 i = 0; i < b.length; i++) {
            number =
                number +
                uint64(uint8(b[i]) * 2**(8 * (b.length - (i + 1))));
        }
        return number;
    }
}
