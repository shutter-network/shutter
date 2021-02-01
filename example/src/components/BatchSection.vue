<template>
  <section class="section content">
    <h1 class="title">Submitted Transactions</h1>
    <table class="table">
      <thead>
        <tr>
          <th>Batch</th>
          <th>Type</th>
          <th>Data</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(tx, index) in txs" v-bind:key="index">
          <td>{{ tx.batchIndex }}</td>
          <td>{{ formatType(tx.type) }}</td>
          <td style="overflow-wrap: break-word; max-width: 200px">
            {{ tx.transaction }}
          </td>
        </tr>
      </tbody>
    </table>
  </section>
</template>

<script>
export default {
  name: "BatchSection",
  data() {
    return {
      txs: [],
    };
  },

  async mounted() {
    // get events from recent blocks
    let blockNumber = await this.$provider.getBlockNumber();
    let events = await this.$batcherContract.queryFilter(
      "TransactionAdded",
      blockNumber - 100
    );
    for (let event of events) {
      this.handleTransactionAdded(...event.args);
    }

    this.$batcherContract.on("TransactionAdded", this.handleTransactionAdded);
  },

  methods: {
    formatType(n) {
      if (n == 0) {
        return "Cipher";
      } else if (n == 1) {
        return "Plain";
      } else {
        return "unknown";
      }
    },
    handleTransactionAdded(batchIndex, type, transaction) {
      this.txs.push({
        batchIndex: batchIndex,
        type: type,
        transaction: transaction,
      });
    },
  },
};
</script>
