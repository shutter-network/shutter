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
          v-on:click="sendEncrypted"
          :disabled="!privateKeyValid"
        >
          Send Encrypted
        </button>
      </div>
      <div class="control">
        <button
          class="button is-primary"
          v-on:click="sendPlain"
          :disabled="!privateKeyValid"
        >
          Send Unencrypted
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {ethers} from "ethers";
import {
  getConfigAtBlock,
  getBatchIndexAtBlock,
  encodeMessage,
} from "../utils.js";

export default {
  name: "SubmitForm",

  data() {
    return {
      batcherContract: null,

      message: "",
      privateKey:
        "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",

      encodedMessage: null,
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
    async sendEncrypted() {
      await this.sendTransaction(0);
    },

    async sendPlain() {
      await this.sendTransaction(1);
    },

    async sendTransaction(type) {
      let nonce = await this.$targetContract.getNonce(this.address);
      let encodedMessage = await encodeMessage(
        this.message,
        nonce,
        this.privateKey
      );

      let blockNumber = (await this.$provider.getBlockNumber()) + 1;
      let config = await getConfigAtBlock(blockNumber, this.configContract);
      let batchIndex = getBatchIndexAtBlock(blockNumber, config);
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
      await tx.wait();
    },
  },
};
</script>

<style></style>
