<template>
  <div id="app">
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title has-text-centered">Shutter Example Dapp</h1>
          <Intro />
        </div>
      </div>
    </section>
    <div v-if="providerFound" class="container">
      <section class="section content">
        <div class="columns">
          <div class="column">
            <SubmitPanel :config="config" :eonKey="eonKey" />
          </div>
          <div class="column">
            <StatusPanel :config="config" />
          </div>
        </div>
      </section>
      <BatchSection />
      <TargetSection />
    </div>
    <div v-else class="container">
      <section class="section content">
        <div class="columns">
          <div class="column is-offset-2 is-two-thirds">
            <p>
              No web3 provider found. Please install Metamask or a similar
              wallet.
            </p>
          </div>
        </div>
      </section>
    </div>
    <Imprint />
  </div>
</template>

<script>
import StatusPanel from "./components/StatusPanel.vue";
import SubmitPanel from "./components/SubmitPanel.vue";
import BatchSection from "./components/BatchSection.vue";
import TargetSection from "./components/TargetSection.vue";
import Intro from "./components/Intro.vue";
import Imprint from "./components/Imprint.vue";
import { getBlockNumber } from "./blocknumber.js";
import { getConfigAtBlock } from "./utils.js";

export default {
  name: "App",
  components: {
    BatchSection,
    StatusPanel,
    SubmitPanel,
    TargetSection,
    Intro,
    Imprint,
  },

  data() {
    return {
      config: null,
      eonKey: null,
      providerFound: this.$providerFound,
    };
  },

  mounted() {
    if (this.$providerFound) {
      this.getKeyAndConfig();
    }
  },

  methods: {
    async getKeyAndConfig() {
      const blockNumber = await getBlockNumber(this.$provider);
      const config = await getConfigAtBlock(blockNumber, this.$configContract);

      const bestKey = await this.$keyBroadcastContract.getBestKey(0);
      const bestKeyNumVotes =
        await this.$keyBroadcastContract.getBestKeyNumVotes(0);

      this.config = config;
      if (bestKeyNumVotes >= config.threshold) {
        this.eonKey = bestKey;
      } else {
        this.eonKey = null;
        console.log("not enough votes for eon public key");
      }
    },
  },
};
</script>

<style></style>
