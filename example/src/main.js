import Vue from "vue";
import {ethers} from "ethers";

import App from "./App.vue";
import {
  configAddress,
  batcherAddress,
  targetAddress,
} from "./assets/addresses.js";
import configContractMetadata from "./assets/abis/ConfigContract.json";
import batcherContractMetadata from "./assets/abis/BatcherContract.json";
import targetContractMetadata from "./assets/abis/TestTargetContract.json";

require("@/assets/main.scss");

Vue.config.productionTip = false;

let provider = new ethers.providers.Web3Provider(window.ethereum);
let configContract = new ethers.Contract(
  configAddress,
  configContractMetadata.abi,
  provider
);
let batcherContract = new ethers.Contract(
  batcherAddress,
  batcherContractMetadata.abi,
  provider
);
let targetContract = new ethers.Contract(
  targetAddress,
  targetContractMetadata.abi,
  provider
);

Vue.prototype.$provider = provider;
Vue.prototype.$configContract = configContract;
Vue.prototype.$batcherContract = batcherContract;
Vue.prototype.$targetContract = targetContract;

new Vue({
  render: (h) => h(App),
}).$mount("#app");
