<template>
  <div>
    <div class="container">
      <p>
        <button
          class="button is-primary"
          v-bind:disabled="pressed"
          v-on:click="enable"
        >
          Connect Wallet
        </button>
      </p>
      <p v-if="wrongNetwork">
        Your wallet seems to be connected to the wrong network. Please select
        Goerli and refresh the page.
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "Enabler",
  data() {
    return {
      pressed: false,
      wrongNetwork: false,
    };
  },
  methods: {
    async enable() {
      this.pressed = true;
      await window.ethereum.request({ method: "eth_requestAccounts" });
      const network = await this.$provider.getNetwork();
      this.wrongNetwork = network.name != "goerli";
      if (!this.wrongNetwork) {
        this.$emit("enabled");
      }
      this.pressed = false;
    },
  },
};
</script>
