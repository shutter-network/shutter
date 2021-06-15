import { BigNumber, ethers } from "ethers";

const targetContractSelector = "0x943d7209";

async function getConfigIndexAtBlock(blockNumber, configContract) {
  let numConfigs = await configContract.numConfigs();
  for (let i = numConfigs - 1; i >= 0; i--) {
    let config = await configContract.configForConfigIndex(i);
    if (config.startBlockNumber <= blockNumber) {
      return i;
    }
  }
  return null;
}

async function getConfigAtBlock(blockNumber, configContract) {
  let index = await getConfigIndexAtBlock(blockNumber, configContract);
  let configArray = await configContract.configForConfigIndex(index);
  return configArrayToObject(configArray);
}

function configArrayToObject(configArray) {
  return {
    startBatchIndex: configArray[0],
    startBlockNumber: configArray[1],
    keypers: configArray[2],
    threshold: configArray[3],
    batchSpan: configArray[4],
    batchSizeLimit: configArray[5],
    transactionSizeLimit: configArray[6],
    transactionGasLimit: configArray[7],
    feeReceiver: configArray[8],
    targetAddress: configArray[9],
    targetFunctionSelector: configArray[10],
    executionTimeout: configArray[11],
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

function getRandomNonce() {
  // nonces range from 0 to 2^64 - 1, but we don't have to use the whole range and smaller
  // numbers are nicer to display.
  return Math.floor(Math.random() * 100000);
}

async function encodeTargetContractData(message, nonce, privateKey) {
  const messageBytes = ethers.utils.toUtf8Bytes(message);
  const data = ethers.utils.defaultAbiCoder.encode(
    ["uint64", "bytes"],
    [nonce, messageBytes]
  );
  const dataHash = ethers.utils.keccak256(data);

  const wallet = new ethers.Wallet(privateKey);
  const flatSig = await wallet.signMessage(ethers.utils.arrayify(dataHash));
  const sig = ethers.utils.splitSignature(flatSig);

  let dataWithSignature = ethers.utils.defaultAbiCoder.encode(
    ["bytes", "uint8", "bytes32", "bytes32"],
    [data, sig.v, sig.r, sig.s]
  );

  return dataWithSignature;
}

function encodeTargetContractCall(data) {
  return ethers.utils.concat([
    targetContractSelector,
    ethers.utils.defaultAbiCoder.encode(["bytes"], [data]),
  ]);
}

function encodeProxyContractData(receiver, data) {
  return ethers.utils.defaultAbiCoder.encode(
    ["address", "bytes"],
    [receiver, data]
  );
}

async function encodeMessage(receiver, message, nonce, privateKey) {
  const targetContractData = await encodeTargetContractData(
    message,
    nonce,
    privateKey
  );
  const targetContractCall = encodeTargetContractCall(targetContractData);
  const proxyContractData = encodeProxyContractData(
    receiver,
    targetContractCall
  );
  return proxyContractData;
}

async function encryptMessage(message, eonPublicKey, batchIndex) {
  var sigma = new Uint8Array(32);
  window.crypto.getRandomValues(sigma);
  const messageArray = ethers.utils.arrayify(message);
  const publicKeyArray = ethers.utils.arrayify(eonPublicKey);
  const result = window.shcrypto.encrypt(
    messageArray,
    publicKeyArray,
    batchIndex.toNumber(),
    sigma
  );
  if (result.error !== null) {
    throw result.error;
  }
  return result.encryptedMessage;
}

export {
  getConfigAtBlock,
  getBatchIndexAtBlock,
  getRandomNonce,
  encodeMessage,
  encryptMessage,
};
