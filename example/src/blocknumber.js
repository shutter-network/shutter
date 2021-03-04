const blockTime = 3;
const pollInterval = 25;

let latestBlockNumber = null;
let latestBlockNumberTime = null;

async function getBlockNumber(provider) {
  await maybeUpdateLatestBlockNumber(provider);

  const dt = new Date() - latestBlockNumberTime;
  const dtBlocks = Math.round(dt / (blockTime * 1000));
  return latestBlockNumber + dtBlocks;
}

async function maybeUpdateLatestBlockNumber(provider) {
  if (
    latestBlockNumberTime === null ||
    new Date() - latestBlockNumberTime >= pollInterval * 1000
  ) {
    await updateLatestBlockNumber(provider);
  }
}

async function updateLatestBlockNumber(provider) {
  const t0 = new Date();
  const blockNumber = await provider.getBlockNumber();
  const t1 = new Date();
  const t = new Date((t1.getTime() + t0.getTime()) / 2);

  latestBlockNumber = blockNumber;
  latestBlockNumberTime = t;
}

export { getBlockNumber };
