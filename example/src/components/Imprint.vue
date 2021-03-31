<template>
  <footer class="footer">
    <div class="content has-text-centered">
      <a v-if="hidden" v-on:click="hidden = false">Imprint</a>
      <div v-else>
        <h1 class="title is-size-6">
          Imprint (information according to § 5 TMG)
        </h1>

        <p class="mb-0">Brainbot Technologies AG</p>
        <p class="mt-0">
          represented by the board members Heiko Hees and Heiko Franßen
        </p>

        <p class="mb-0">Taunusstrasse 61</p>
        <p class="my-0">55120 Mainz</p>
        <p class="mt-0">Germany</p>

        <p class="mb-0">Phone: +49 (0) 6131 2116391</p>
        <p class="my-0">Fax: +49 (0) 6131 2116392</p>
        <p class="mt-0">Email: contact@brainbot.com</p>

        <p class="mb-0">District Court: Mainz</p>
        <p class="my-0">Register number: HRB 8375</p>
        <p class="mt-0">VAT ID: DE231725642</p>

        <p>Responsible for content according to § 6 MDStV: Heiko Hees</p>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "Imprint",
  data() {
    return {
      hidden: true,
    };
  },
};
</script>
