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
      <SubmitSection />
      <BatchSection />
      <TargetSection />
    </div>
  </div>
</template>

<script>
import SubmitSection from "./components/SubmitSection.vue";
import TargetSection from "./components/TargetSection.vue";
import BatchSection from "./components/BatchSection.vue";

export default {
  name: "App",
  components: {
    SubmitSection,
    TargetSection,
    BatchSection,
  },
};
</script>

<style></style>
