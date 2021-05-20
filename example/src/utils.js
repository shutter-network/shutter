import { BigNumber, ethers } from "ethers";

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

async function encodeMessage(message, nonce, privateKey) {
  let messageBytes = ethers.utils.toUtf8Bytes(message);
  let payload = ethers.utils.defaultAbiCoder.encode(
    ["uint64", "bytes"],
    [nonce, messageBytes]
  );
  let payloadHash = ethers.utils.keccak256(payload);
  let wallet = new ethers.Wallet(privateKey);
  let flatSig = await wallet.signMessage(ethers.utils.arrayify(payloadHash));
  let sig = ethers.utils.splitSignature(flatSig);
  let encoded = ethers.utils.defaultAbiCoder.encode(
    ["bytes", "uint8", "bytes32", "bytes32"],
    [payload, sig.v, sig.r, sig.s]
  );
  return encoded;
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
