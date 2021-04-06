<template>
  <div>
    <p>
      Enter the message you want to send and a private key to sign it with. Do
      not use a key that corresponds to a real account. The default is just a
      random string which works just fine.
    </p>
    <p>
      To send the message, press one of the buttons below and confirm the
      transaction in your wallet. Only encrypted messages will be frontrunning
      protected.
    </p>
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
      <p v-else class="help is-success">Valid</p>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button
          class="button is-primary"
          :class="{ 'is-loading': waitingForTx }"
          v-on:click="onSend(0)"
          :disabled="sendDisabled"
        >
          Send Encrypted
        </button>
      </div>
      <div class="control">
        <button
          class="button is-primary"
          :class="{ 'is-loading': waitingForTx }"
          v-on:click="onSend(1)"
          :disabled="sendDisabled"
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
  getBatchIndexAtBlock,
  getRandomNonce,
  encodeMessage,
  encryptMessage,
} from "../utils.js";
import { getBlockNumber } from "../blocknumber.js";

// number of blocks between sending the tx and it being included in the chain
const expectedInclusionDelay = 3;

export default {
  name: "SubmitForm",
  props: ["config", "eonKey"],

  data() {
    const keyBytes = window.crypto.getRandomValues(new Uint8Array(32));
    return {
      batcherContract: null,

      message: "",
      privateKey: ethers.utils.hexlify(keyBytes),
      waitingForTx: false,
    };
  },

  computed: {
    sendDisabled() {
      return (
        this.waitingForTx ||
        this.eonKey === null ||
        this.config === null ||
        !this.privateKeyValid ||
        !this.messageValid
      );
    },
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
    messageValid() {
      return this.message.length > 0;
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
      if (this.sendDisabled) {
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

      let blockNumber = await getBlockNumber(this.$provider);
      let batchIndex = getBatchIndexAtBlock(
        blockNumber + expectedInclusionDelay,
        this.config
      );

      if (type == 0) {
        encodedMessage = await encryptMessage(
          encodedMessage,
          this.eonKey,
          batchIndex
        );
      }

      let tx = await this.batcherContract.addTransaction(
        batchIndex,
        type,
        encodedMessage
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
  },
};
</script>

<style></style>
