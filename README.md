# Shutter

This repository contains Shutter, a threshold encryption based frontrunning protection system for
Ethereum smart contracts.

The code is split into three directories:

- `contracts/`: the contracts to be deployed on mainnet, written in Solidity and tested in Python
  with eth-brownie
- `shuttermint/`: the main application code written in Go, in particular the Shuttermint and Keyper
  implementation
- `example/`: a little example dapp user interface

To see it in action, check out the example application hosted
[here](https://brainbot-com.github.io/shutter-example-dapp/).

## Setup a Full Testing Environment

The following steps will setup a local testing environment consisting of an Ethereum dev chain, a
single Shuttermint node, and an arbitrary number of keyper nodes. The commands referenced can be
found in `shuttermint/bin` after `make build` has been run in the `shuttermint/` directory.

### 1) Start the Ethereum Dev Chain

There are many options, but we suggest using either Geth or Ganache with a block time of 3s. Make
sure that the JSON RPC interface is enabled and accessible via websockets and that you know the
private key of at least one funded account. In the rest of the guide, we'll assume the JSON RPC URL
is `ws://localhost:8545` and the private key is
`b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773`. If Ganache is used, the
following command replicates these settings:

```
ganache-cli -d -b 3
```

### 2) Deploy the Shutter Contracts:

Next, the Shutter contract suite has to be deployed on the dev chain:

```
shuttermint deploy -e ws://localhost:8545 -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773 -o contracts.json
```

This will output the addresses into a JSON file for later reference.

### 3) Initialize the Keypers

Now, we can initialize the keypers. Running

```
prepare configs -c contracts.json -e ws://localhost:8545 --fixed-shuttermint-port -n 3
```

creates a new directory called `testrun` (customizable via the `-d` flag) with one subdirectory for
each of three keypers.

The newly created config files found in `testrun/keyper<n>/config.toml` contain an Ethereum private
key under `SigningKey`. This key is randomly generated, so doesn't yet have access to ETH needed
to send transactions. To change this, run

```
prepare fund -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
```

which will send 1 (dev chain) ETH from the deployer account to each keyper.

### 4) Schedule a Batch Config

Next, a batch config has to be defined and scheduled on the main chain, defining some important
system parameters. Write the following contents to a new file called `config.json`:

```
{
    "StartBatchIndex": 0,
    "StartBlockNumber": <start block number>,
    "Keypers": [
        <keyper addresses>
    ],
    "Threshold": <threshold>,
    "BatchSpan": 10,
    "BatchSizeLimit": 100000,
    "TransactionSizeLimit": 1000,
    "TransactionGasLimit": 10000,
    "FeeReceiver": "0x1111111111111111111111111111111111111111",
    "TargetAddress": <target contract address>,
    "TargetFunctionSelector": "0x943d7209",
    "ExecutionTimeout": 15
}
```

Make sure to fill in the following fields:

- `StartBlockNumber`: Use a block number of the Ethereum dev chain that's in the near future (e.g.,
  current block number plus 100 blocks)
- `TargetAddress`: The hex encoded, checksummed address of the test target contract to be found in
  `contracts.json` under `TargetContract`.
- `Threshold`: The threshold parameter to use. We suggest two thirds of the number of keypers,
  rounded up.
- `Keypers`: The list of keyper addresses. The addresses must correspond to the private keys
  defined in the keyper config files. To convert a private key to an address, you can use the
  following command, provided you have installed the right packages:

  ```
  python -c "import sys; from eth_utils import *; from eth_keys import keys; print(keys.PrivateKey(decode_hex(sys.argv[1])).public_key.to_checksum_address())" <private key>
  ```

Now, send the config to the config contract:

```
shuttermint config -c contracts.json -e ws://localhost:8545 set-next --config config.json -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
```

After all transactions have been confirmed, double check that the contract has accepted everything:

```
shuttermint config -c contracts.json -e ws://localhost:8545 query -i next | diff config.json -
```

If so (i.e., the output is empty), schedule the config with

```
shuttermint config -c contracts.json -e ws://localhost:8545 schedule -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
```

and, once more,

```
shuttermint config -c contracts.json -e ws://localhost:8545 query -i last | diff config.json -
```

double check that the changes are live.

### 5) Start and Bootstrap Shuttermint

Now that the main chain is ready, we have to tend to the Shuttermint chain. Initialize its directory with

```
shuttermint init --dev --root testchain
```

and launch it:

```
shuttermint chain --config testchain/config/config.toml
```

Lastly, the Shuttermint chain has to be told about the initial keyper set. To do so, run the
following command:

```
shuttermint bootstrap -c contracts.json -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773 -e ws://localhost:8545
```

### 6) Run the Keypers

Finally, the keyper nodes can be started:

```
shuttermint keyper --config testrun/keyper<n>/config.toml
```

Run this command for each of the keypers initialized.

Now, the keypers should start generating keys as well as decrypting and executing batches on the
main chain.

### 7) Shut It Down

To shut everything down in the end, kill the following processes:

- the keyper nodes started in step 6
- the Shuttermint node started in step 5
- the Ethereum chain started in step 1

Delete the `testchain` and `testrun` directories as well as the `contracts.json` and `config.json`
files to clean up.

## Contracts

At its core, Shutter's contract suite consists of three main components:

- the config contract
- the batcher contract
- the executor contract

The following contracts perform auxiliary tasks:

- the deposit contract
- the fee bank contract
- the key broadcast contract
- the keyper slasher contract

All contracts can be found in the `contracts/` subdirectory.

### Config Contract

The config contract defines the system parameters. Since they must be able to evolve over time,
they are given as a list of batch config objects.

New batch configs can be scheduled for the future, as long as the `configChangeHeadsUp` interval
set during construction is abided by. This is intended to give keypers enough time to react to
changes. Once scheduled, batch configs cannot be removed and their order cannot be changed.

Scheduling a config is a two step process: First, the config object is drafted using the various
`setNext...` functions and, second, it is finalized using `schedule`.

Only the contract owner is allowed to schedule configs. This role is inteded to be played by a DAO.

The batch config objects divide time into a sequence of batches and each config is applicable to a
range of them. The following fields specify this:

- `startBatchIndex`: the index of the first batch for which the config is applicable to
- `startBlockNumber`: the Ethereum block number at which the first batch starts
- `batchSpan`: the length of each batch in Ethereum blocks

The first config has a `startBatchIndex` of 0. It will end with the `startBatchIndex` of the second
config, and so on. The config contract enforces that transitions between batches are seamless,
i.e., the start block number of a config fits to the batch span and start block number of the
previous config (e.g., if config 1 starts at batch index `100` and block `500` and has a span of
`10` and config 2 at batch index `200`, it's start block number must be `600`).

The batch span can be zero denoting that the system is disabled. It can be enabled again by
scheduling a new config with a batch span greater than zero.

The key generation parameters are set by the `keypers` and `threshold` fields. Note that `keypers`
must not contain duplicates and that the threshold must less than or equal to the number of
keypers, but greater than half of it.

The transactions that are allowed to enter a batch are constrained by the following fields:

- `transactionSizeLimit`: the maximum size of a transaction in bytes
- `batchSizeLimit`: the maximum size of all transactions in a single batch in bytes
- `feeReceiver`: address that will receive transaction fees

For transaction execution after decryption, the following fields are relevant:

- `transactionGasLimit`: the maximum amount of gas a transaction is allowed to consume during
  execution once it is decrypted
- `targetAddress`: the address of the contract to which the derypted transactions will be passed
  to
- `targetFunctionSelector`: the 4 byte function selector that specifies the function in the target
  contract that will be called with each decrypted transaction
- `executionTimeout`: the number of blocks to pass between the end of a batch and the time at which
  it is assumed that decryption has failed and can be skipped (e.g., because too many keyper were
  offline)

### Batcher Contract

### Executor Contract

### Auxiliary Contracts
