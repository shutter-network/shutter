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
    <p>
      Here you can see the execution status. The current batch is the batch that
      will close next, so transactions usually go into this one. It advances
      every ten blocks.
    </p>
    <p>
      Once a batch is closed, keypers will execute it, so the last executed
      batch usually trails behind the current batch by one or two blocks as long
      as the keypers are running and there is no backlog.
    </p>
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
