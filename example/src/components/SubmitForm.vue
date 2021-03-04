<template>
  <div>
    <div class="field">
      <label class="label">Private Key</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="0x..."
          v-model="privateKey"
        />
      </div>
      <p v-if="!privateKeyValid" class="help is-danger">Invalid private key</p>
      <p v-else class="help is-success">&check;</p>
    </div>

    <div class="field">
      <label class="label">Message</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Message"
          v-model="message"
        />
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button
          class="button is-primary"
          :class="{ 'is-loading': waitingForTx }"
          v-on:click="onSend(0)"
          :disabled="!privateKeyValid"
        >
          Send Encrypted
        </button>
      </div>
      <div class="control">
        <button
          class="button is-primary"
          :class="{ 'is-loading': waitingForTx }"
          v-on:click="onSend(1)"
          :disabled="!privateKeyValid"
        >
          Send Unencrypted
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ethers } from "ethers";
import {
  getConfigAtBlock,
  getBatchIndexAtBlock,
  getRandomNonce,
  encodeMessage,
  encryptMessage,
} from "../utils.js";
import { getBlockNumber } from "../blocknumber.js";

export default {
  name: "SubmitForm",

  data() {
    return {
      batcherContract: null,

      message: "",
      privateKey:
        "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
      waitingForTx: false,
    };
  },

  computed: {
    privateKeyValid() {
      if (!ethers.utils.isHexString(this.privateKey, 32)) {
        return false;
      }
      try {
        new ethers.Wallet(this.privateKey);
      } catch {
        return false;
      }
      return true;
    },
    address() {
      if (!this.privateKeyValid) {
        return null;
      }
      return ethers.utils.computeAddress(this.privateKey);
    },
  },

  mounted() {
    let signer = this.$provider.getSigner();
    this.configContract = this.$configContract.connect(signer);
    this.batcherContract = this.$batcherContract.connect(signer);
  },

  methods: {
    async onSend(type) {
      if (this.waitingForTx) {
        return;
      }
      this.waitingForTx = true;
      try {
        await this.sendTransaction(type);
      } finally {
        this.waitingForTx = false;
      }
    },

    async sendTransaction(type) {
      let nonce = getRandomNonce();
      let encodedMessage = await encodeMessage(
        this.message,
        nonce,
        this.privateKey
      );

      let blockNumber = await this.waitForGoodBlock();
      let config = await getConfigAtBlock(blockNumber + 1, this.configContract);
      let batchIndex = getBatchIndexAtBlock(blockNumber + 1, config);

      if (type == 0) {
        let bestKey = await this.$keyBroadcastContract.getBestKey(
          config.startBatchIndex
        );
        let bestKeyNumVotes = await this.$keyBroadcastContract.getBestKeyNumVotes(
          config.startBatchIndex
        );
        if (bestKeyNumVotes < config.threshold) {
          console.log("not enough votes for eon public key");
          return;
        }
        encodedMessage = await encryptMessage(
          encodedMessage,
          bestKey,
          batchIndex
        );
      }

      let tx = await this.batcherContract.addTransaction(
        batchIndex,
        type,
        encodedMessage,
        {
          gasLimit: 200000,
        }
      );
      console.log(
        "tx",
        tx.hash,
        "sent for batch",
        batchIndex.toString(),
        "in block",
        blockNumber.toString()
      );
    },

    async waitForGoodBlock() {
      for (;;) {
        const blockNumber = await getBlockNumber(this.$provider);
        const config = await getConfigAtBlock(
          blockNumber + 1,
          this.configContract
        );
        const batchIndexNow = getBatchIndexAtBlock(blockNumber + 1, config);
        const batchIndexSoon = getBatchIndexAtBlock(blockNumber + 2, config);
        if (batchIndexNow.eq(batchIndexSoon)) {
          return blockNumber;
        }
        await new Promise((resolve) => setTimeout(resolve, 1000));
      }
    },
  },
};
</script>

<style></style>
