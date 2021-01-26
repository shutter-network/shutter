<template>
  <button class="button is-primary" v-on:click="enable">Enable Ethereum</button>
</template>

<script>
export default {
  name: "Enabler",
  methods: {
    enable() {
      window.ethereum.request({method: "eth_requestAccounts"}).then(() => {
        this.$emit("enabled");
      });
    },
  },
};
</script>
