from brownie import accounts
from brownie import ConfigContract
from brownie import KeyBroadcastContract


def main():
    acc = accounts.load("ganache9")
    print("Starting deployment")
    cc = ConfigContract.deploy(5, {"from": acc})
    KeyBroadcastContract.deploy(cc.address, {"from": acc})
