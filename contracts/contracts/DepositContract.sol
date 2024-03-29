// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

import {IERC777} from "openzeppelin/contracts/token/ERC777/IERC777.sol";
import {IERC777Recipient} from "openzeppelin/contracts/token/ERC777/IERC777Recipient.sol";
import {IERC1820Registry} from "openzeppelin/contracts/utils/introspection/IERC1820Registry.sol";
import {Ownable} from "openzeppelin/contracts/access/Ownable.sol";

struct Deposit {
    uint256 amount;
    uint64 withdrawalDelayBlocks;
    uint64 withdrawalRequestedBlock;
    bool slashed;
}

contract DepositContract is IERC777Recipient, Ownable {
    event DepositChanged(
        address indexed account,
        uint256 amount,
        uint64 withdrawalDelayBlocks,
        uint64 withdrawalRequestedBlock,
        bool withdrawn,
        bool slashed
    );
    event SlashingReceiverSet(address newSlashingReceiver);

    IERC1820Registry private _erc1820 =
        IERC1820Registry(0x1820a4B7618BdE71Dce8cdc73aAB6C95905faD24);
    bytes32 private constant TOKENS_RECIPIENT_INTERFACE_HASH =
        keccak256("ERC777TokensRecipient");

    IERC777 public token;
    address public slasher;
    mapping(address => Deposit) private _deposits;
    address public slashingReceiver;

    constructor(IERC777 tokenContract) Ownable() {
        token = tokenContract;

        _erc1820.setInterfaceImplementer(
            address(this),
            TOKENS_RECIPIENT_INTERFACE_HASH,
            address(this)
        );
    }

    function setSlasher(address slasherAddress) public {
        require(
            slasher == address(0),
            "DepositContract: slasher address already set"
        );
        slasher = slasherAddress;
    }

    function setSlashingReceiver(address newSlashingReceiver) public onlyOwner {
        slashingReceiver = newSlashingReceiver;
        emit SlashingReceiverSet({newSlashingReceiver: newSlashingReceiver});
    }

    function tokensReceived(
        address, // operator
        address from,
        address, // to
        uint256 amount,
        bytes calldata userData,
        bytes calldata // operatorData
    ) public override {
        require(
            msg.sender == address(token),
            "DepositContract: received invalid token"
        );
        uint64 withdrawalInterval = abi.decode(userData, (uint64));
        _deposit(from, amount, withdrawalInterval);
    }

    function requestWithdrawal() public {
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

    function withdraw(address recipient) public {
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
        token.send(recipient, deposit.amount, "");

        emit DepositChanged({
            account: msg.sender,
            amount: 0,
            withdrawalDelayBlocks: 0,
            withdrawalRequestedBlock: 0,
            withdrawn: true,
            slashed: false
        });
    }

    function slash(address account) public {
        require(msg.sender == slasher);

        Deposit memory deposit = _deposits[account];

        uint256 amount = deposit.amount;

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

        if (slashingReceiver == address(0)) {
            token.burn(amount, "");
        } else {
            token.send(slashingReceiver, amount, "");
        }
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
}
