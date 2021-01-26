<template>
  <div>
    <div class="field">
      <label class="label">Message</label>
      <div class="control">
        <input class="input" type="text" placeholder="Message" />
      </div>
    </div>

    <div class="field">
      <label class="label">Batch Index</label>
      <div class="control">
        <input class="input" type="text" placeholder="Batch index" />
      </div>
    </div>

    <div class="field">
      <label class="label">Nonce</label>
      <div class="control">
        <input class="input" type="text" placeholder="Nonce" />
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button class="button is-primary" v-on:click="sendEncrypted">
          Send Encrypted
        </button>
      </div>
      <div class="control">
        <button class="button is-primary" v-on:click="sendPlain">
          Send Unencrypted
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {getConfigAtBlock, getBatchIndexAtBlock} from "../utils.js";

export default {
  name: "SubmitForm",

  data() {
    return {
      batcherContract: null,
    };
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
      let blockNumber = (await this.$provider.getBlockNumber()) + 1;
      let config = await getConfigAtBlock(blockNumber, this.configContract);
      let batchIndex = getBatchIndexAtBlock(blockNumber, config);
      console.log(
        "start block",
        config.startBlockNumber.toString(),
        "batch span",
        config.batchSpan.toString(),
        "start batch",
        config.startBatchIndex.toString()
      );
      console.log(
        "sending tx for batch",
        batchIndex.toString(),
        "at block",
        blockNumber.toString()
      );
      let tx = await this.batcherContract.addTransaction(
        batchIndex,
        type,
        "0xaabbcc"
      );
      await tx.wait();
    },
  },
};
</script>

<style></style>
