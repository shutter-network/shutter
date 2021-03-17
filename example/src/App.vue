<template>
  <div id="app">
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Shutter Example Dapp</h1>
        </div>
      </div>
    </section>
    <div class="container">
      <section class="section content">
        <div class="columns">
          <div class="column">
            <SubmitPanel />
          </div>
          <div class="column">
            <StatusPanel />
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

export default {
  name: "App",
  components: {
    BatchSection,
    StatusPanel,
    SubmitPanel,
    TargetSection,
  },
};
</script>

<style></style>
