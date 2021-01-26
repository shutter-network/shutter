<template>
  <section class="section">
    <h1 class="title">Submit Transaction</h1>
    <Enabler v-if="!enabled" v-on:enabled="enabled = true" />
    <SubmitForm v-if="enabled" />
  </section>
</template>

<script>
import SubmitForm from "./SubmitForm.vue";
import Enabler from "./Enabler.vue";

export default {
  name: "SubmitSection",
  components: {
    SubmitForm,
    Enabler,
  },

  data() {
    return {
      enabled: false,
    };
  },
};
</script>
