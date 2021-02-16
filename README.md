# Shutter

This repository contains Shutter, a threshold encryption based frontrunning protection guard for
Ethereum.

The code is split into three directories:

- `contracts/`: the contracts to be deployed on mainnet, written in Solidity and tested in Python
  with eth-brownie
- `shuttermint/`: the main application code written in Go, in particular the Shuttermint and Keyper
  implementation
- `example/`: a little example dapp user interface

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
deploy deploy -e ws://localhost:8545 -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773 -o contracts.json
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
config -c contracts.json -e ws://localhost:8545 set-next --config config.json -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
```

After all transactions have been confirmed, double check that the contract has accepted everything:

```
config -c contracts.json -e ws://localhost:8545 query -i next | diff config.json -
```

If so (i.e., the output is empty), schedule the config with

```
config -c contracts.json -e ws://localhost:8545 schedule -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
```

and, once more,

```
config -c contracts.json -e ws://localhost:8545 query -i last | diff config.json -
```

double check that the changes are live.

### 5) Start and Bootstrap Shuttermint

Now that the main chain is ready, we have to tend to the Shuttermint chain. Initialize its directory with

```
shuttermint init --dev --root testchain
```

and launch it:

```
shuttermint run --config testchain/config/config.toml
```

Lastly, the Shuttermint chain has to be told about the initial keyper set. To do so, run the
following command, substituting the config contract address from `contracts.json`:

```
shuttermint bootstrap -c <config contract address> -k b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773 -e ws://localhost:8545
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
