import { BigNumber, ethers } from "ethers";

async function getConfigIndexAtBlock(blockNumber, configContract) {
  let numConfigs = await configContract.numConfigs();
  for (let i = numConfigs - 1; i >= 0; i--) {
    let config = await configContract.configs(i);
    if (config.startBlockNumber <= blockNumber) {
      return i;
    }
  }
  return null;
}

async function getKeypers(configIndex, configContract) {
  let numKeypers = await configContract.configNumKeypers(configIndex);
  let keypers = [];
  for (let i = 0; i < numKeypers; i++) {
    let keyper = await configContract.configKeypers(configIndex, i);
    keypers.push(keyper);
  }
  return keypers;
}

async function getConfigAtBlock(blockNumber, configContract) {
  let index = await getConfigIndexAtBlock(blockNumber, configContract);
  let configArray = await configContract.configs(index);
  let keypers = await getKeypers(index, configContract);
  return configArrayToObject(configArray, keypers);
}

function configArrayToObject(configArray, keypers) {
  return {
    startBatchIndex: configArray[0],
    startBlockNumber: configArray[1],
    keypers: keypers,
    threshold: configArray[2],
    batchSpan: configArray[3],
    batchSizeLimit: configArray[4],
    transactionSizeLimit: configArray[5],
    transactionGasLimit: configArray[6],
    feeReceiver: configArray[7],
    targetAddress: configArray[8],
    targetFunctionSelector: configArray[9],
    executionTimeout: configArray[10],
  };
}

function getBatchIndexAtBlock(blockNumber, config) {
  let blockNumberBig = BigNumber.from(blockNumber);
  console.assert(config.batchSpan.gt(0), "config is inactive");
  let blocksSinceStart = blockNumberBig.sub(config.startBlockNumber);
  let batchesSinceStart = blocksSinceStart.div(config.batchSpan);
  let batchIndex = config.startBatchIndex.add(batchesSinceStart);
  return batchIndex;
}

async function encodeMessage(message, nonce, privateKey) {
  let messageBytes = ethers.utils.toUtf8Bytes(message);
  let payload = ethers.utils.defaultAbiCoder.encode(
    ["uint64", "bytes"],
    [nonce, messageBytes]
  );
  let wallet = new ethers.Wallet(privateKey);
  let flatSig = await wallet.signMessage(payload);
  let sig = ethers.utils.splitSignature(flatSig);
  let encoded = ethers.utils.defaultAbiCoder.encode(
    ["bytes", "uint8", "bytes32", "bytes32"],
    [payload, sig.v, sig.r, sig.s]
  );
  return encoded;
}

export { getConfigAtBlock, getBatchIndexAtBlock, encodeMessage };
