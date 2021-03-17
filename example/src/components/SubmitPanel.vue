<template>
  <div>
    <h1 class="title">Submit Transaction</h1>
    <Enabler v-if="!enabled" v-on:enabled="enabled = true" />
    <SubmitForm v-if="enabled" />
  </div>
</template>

<script>
import SubmitForm from "./SubmitForm.vue";
import Enabler from "./Enabler.vue";

export default {
  name: "SubmitPanel",
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
