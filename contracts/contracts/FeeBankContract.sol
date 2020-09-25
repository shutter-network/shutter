// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;

/// @title A contract that stores fees for later withdrawal
contract FeeBankContract {
    /// @notice The event emitted whenever ETH is deposited.
    /// @param depositor The address of the account making the deposit.
    /// @param receiver The address of the account eligible for withdrawal.
    /// @param amount The newly deposited amount.
    /// @param totalAmount The total amount the receiver can withdraw, including the new deposit.
    event DepositEvent(
        address depositor,
        address receiver,
        uint64 amount,
        uint64 totalAmount
    );

    /// @notice The event emitted whenever ETH is withdrawn.
    /// @param sender The address of the account that triggered the withdrawal.
    /// @param receiver The address of the account to which the ETH is sent.
    /// @param amount The withdrawn amount.
    /// @param totalAmount The remaining deposit.
    event WithdrawEvent(
        address sender,
        address receiver,
        uint64 amount,
        uint64 totalAmount
    );

    mapping(address => uint64) public deposits;

    /// @notice Deposit ETH for later withdrawal
    /// @param _receiver Address of the account that is eligible for withdrawal.
    function deposit(address _receiver) external payable {
        require(_receiver != address(0));
        require(msg.value > 0);
        require(msg.value <= type(uint64).max - deposits[_receiver]);
        deposits[_receiver] += uint64(msg.value);

        emit DepositEvent(
            msg.sender,
            _receiver,
            uint64(msg.value),
            deposits[_receiver]
        );
    }

    /// @notice Withdraw ETH previously deposited in favor of the caller.
    /// @param _receiver The address to which the ETH will be sent.
    /// @param _amount The amount to withdraw (must not be greater than the deposited amount)
    function withdraw(address _receiver, uint64 _amount) external {
        withdrawInternal(_receiver, _amount);
    }

    /// @notice Withdraw all ETH previously deposited in favor of the caller and send it to them.
    function withdraw() external {
        withdrawInternal(msg.sender, deposits[msg.sender]);
    }

    function withdrawInternal(address _receiver, uint64 _amount) internal {
        require(_receiver != address(0));
        uint64 _deposit = deposits[msg.sender];
        require(_deposit > 0);
        require(_amount <= _deposit);
        deposits[msg.sender] = _deposit - _amount;
        (bool _success, ) = _receiver.call{value: _amount}("");
        require(_success);
        emit WithdrawEvent(
            msg.sender,
            _receiver,
            _amount,
            deposits[msg.sender]
        );
    }
}
