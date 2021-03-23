<template>
  <div id="app">
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title has-text-centered">Shutter Example Dapp</h1>
          <div class="columns">
            <div class="column is-offset-2 is-two-thirds">
              <p>
                This is a minimal example of an application using Shutter. Users
                can send messages which the application will emit as events.
              </p>
              <p>
                To send a message, press the button below to connect with your
                browser wallet and make sure you are connected to Goerli. Next,
                enter an arbitrary message and a private key to sign it with
                (this can be any key, the account does not need to have any
                funds).
              </p>
              <p>
                Next, press either "Send Encrypted" or "Send Unencrypted" to add
                a nonce and signature and wrap. It will be put in an envelope
                transaction which your wallet will prompt you to sign.
              </p>
              <p>
                Once the envelope transaction is included in a block, it will
                appear in the "Recently Submitted Transactions" table.
              </p>
              <p>
                Transactions are processed in batches. Once the batch in which
                your transaction has been included is closed, it will be
                decrypted and executed. Then, your message will appear under
                "Recently Executed Transactions."
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
    <div class="container">
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
  </div>
</template>

<script>
import StatusPanel from "./components/StatusPanel.vue";
import SubmitPanel from "./components/SubmitPanel.vue";
import BatchSection from "./components/BatchSection.vue";
import TargetSection from "./components/TargetSection.vue";
import { getBlockNumber } from "./blocknumber.js";
import { getConfigAtBlock } from "./utils.js";

export default {
  name: "App",
  components: {
    BatchSection,
    StatusPanel,
    SubmitPanel,
    TargetSection,
  },

  data() {
    return {
      config: null,
      eonKey: null,
    };
  },

  mounted() {
    this.getKeyAndConfig();
  },

  methods: {
    async getKeyAndConfig() {
      const blockNumber = await getBlockNumber(this.$provider);
      const config = await getConfigAtBlock(blockNumber, this.$configContract);

      const bestKey = await this.$keyBroadcastContract.getBestKey(
        config.startBatchIndex
      );
      const bestKeyNumVotes = await this.$keyBroadcastContract.getBestKeyNumVotes(
        config.startBatchIndex
      );

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
