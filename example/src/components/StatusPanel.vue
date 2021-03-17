<template>
  <div>
    <h1 class="title">Status</h1>
    <table class="table is-fullwidth">
      <tbody>
        <tr>
          <th>Block Number</th>
          <td>{{ blockNumber !== null ? blockNumber : "Unknown" }}</td>
        </tr>
        <tr>
          <th>Current Batch</th>
          <td>{{ batchIndex !== null ? batchIndex : "Unknown" }}</td>
        </tr>
        <tr>
          <th>Last Executed Batch</th>
          <td>
            {{ lastExecutedBatch !== null ? lastExecutedBatch : "Unknown" }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { getBlockNumber } from "../blocknumber.js";
import { getBatchIndexAtBlock } from "../utils.js";

export default {
  name: "StatusPanel",
  props: ["config"],

  data() {
    return {
      blockNumber: null,
      batchIndex: null,
      halfSteps: null,
    };
  },

  computed: {
    lastExecutedBatch() {
      if (this.halfSteps == null) {
        return null;
      }
      const fullyExecutedBatches = Math.floor(this.halfSteps.toNumber() / 2);
      return fullyExecutedBatches - 1;
    },
  },

  watch: {
    config: {
      immediate: true,
      handler() {
        if (this.config !== null) {
          this.poll();
        }
      },
    },
  },

  methods: {
    async update() {
      if (this.config === null) {
        return;
      }

      const blockNumber = await getBlockNumber(this.$provider);
      const batchIndex = getBatchIndexAtBlock(blockNumber, this.config);
      const halfSteps = await this.$executorContract.numExecutionHalfSteps();

      this.blockNumber = blockNumber;
      this.batchIndex = batchIndex;
      this.halfSteps = halfSteps;
    },
    async poll() {
      for (;;) {
        await this.update();
        await new Promise((resolve) => setTimeout(resolve, 1000));
      }
    },
  },
};
</script>
