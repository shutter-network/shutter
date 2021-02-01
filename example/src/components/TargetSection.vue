<template>
  <section class="section content">
    <h1 class="title">Executed Transactions</h1>
    <table class="table">
      <thead>
        <tr>
          <th>Sender</th>
          <th>Nonce</th>
          <th>Data</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(tx, index) in txs" v-bind:key="index">
          <td>{{ tx.sender }}</td>
          <td>{{ tx.nonce }}</td>
          <td>{{ tx.data }}</td>
        </tr>
      </tbody>
    </table>
  </section>
</template>

<script>
export default {
  name: "TargetSection",

  data() {
    return {
      txs: [],
    };
  },

  async mounted() {
    // get events from recent blocks
    let blockNumber = await this.$provider.getBlockNumber();
    let events = await this.$targetContract.queryFilter(
      "ExecutedTransaction",
      blockNumber - 100
    );
    for (let event of events) {
      this.handleExecutedTransaction(...event.args);
    }

    this.$targetContract.on(
      "ExecutedTransaction",
      this.handleExecutedTransaction
    );
  },

  methods: {
    handleExecutedTransaction(sender, data, nonce) {
      this.txs.push({
        sender: sender,
        nonce: nonce,
        data: data,
      });
    },
  },
};
</script>
