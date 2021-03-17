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
          <th>Executed Batches</th>
          <td>{{ executedBatches !== null ? executedBatches : "Unknown" }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { getBlockNumber } from "../blocknumber.js";
import { getConfigAtBlock, getBatchIndexAtBlock } from "../utils.js";

export default {
  name: "StatusPanel",

  data() {
    return {
      blockNumber: null,
      batchIndex: null,
      halfSteps: null,
    };
  },

  computed: {
    executedBatches() {
      if (this.halfSteps == null) {
        return null;
      }
      return Math.floor((this.halfSteps.toNumber() + 1) / 2);
    },
  },

  mounted() {
    this.poll();
  },

  methods: {
    async update() {
      const blockNumber = await getBlockNumber(this.$provider);
      const config = await getConfigAtBlock(blockNumber, this.$configContract);
      const batchIndex = getBatchIndexAtBlock(blockNumber, config);
      const halfSteps = await this.$executorContract.numExecutionHalfSteps();

      this.blockNumber = blockNumber;
      this.batchIndex = batchIndex;
      this.halfSteps = halfSteps;
    },
    async poll() {
      for (;;) {
        await this.update();
        await new Promise((resolve) => setTimeout(resolve, 2000));
      }
    },
  },
};
</script>
