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
          <td>{{ tx.transaction }}</td>
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

  mounted() {
    this.$batcherContract.on(
      "TransactionAdded",
      (batchIndex, type, transaction) => {
        this.txs.push({
          batchIndex: batchIndex,
          type: type,
          transaction: transaction,
        });
      }
    );
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
  },
};
</script>
