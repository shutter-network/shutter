<template>
  <div>
    <p>
      Enter the message you want to send, press one of the buttons below, and
      confirm the transaction in your wallet. Only encrypted messages will be
      frontrunning protected.
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
      <a v-if="!showAdvancedOptions" v-on:click="showAdvancedOptions = true"
        >Show advanced options</a
      >
      <a v-else v-on:click="showAdvancedOptions = false"
        >Hide advanced options</a
      >
    </div>

    <div v-if="showAdvancedOptions" class="field">
      <label class="label">Signing Private Key</label>
      <p>
        This is the private key used to sign the message. Anything works here
        and for convenience a random key has been generated for you. As general
        security practice dictates, do not enter a key used in other contexts,
        especially not one that controls any funds.
      </p>
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
          Send encrypted
        </button>
      </div>
      <div class="control">
        <button
          class="button is-primary"
          :class="{ 'is-loading': waitingForTx }"
          v-on:click="onSend(1)"
          :disabled="sendDisabled"
        >
          Send unencrypted
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
      showAdvancedOptions: false,
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
