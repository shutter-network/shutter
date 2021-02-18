import Vue from "vue";
import { ethers } from "ethers";

import App from "./App.vue";
import {
  configAddress,
  batcherAddress,
  targetAddress,
  keyBroadcastAddress,
} from "./assets/addresses.js";
import configContractMetadata from "./assets/abis/ConfigContract.json";
import batcherContractMetadata from "./assets/abis/BatcherContract.json";
import targetContractMetadata from "./assets/abis/TestTargetContract.json";
import keyBroadcastContractMetadata from "./assets/abis/KeyBroadcastContract.json";

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
let keyBroadcastContract = new ethers.Contract(
  keyBroadcastAddress,
  keyBroadcastContractMetadata.abi,
  provider
);

Vue.prototype.$provider = provider;
Vue.prototype.$configContract = configContract;
Vue.prototype.$batcherContract = batcherContract;
Vue.prototype.$targetContract = targetContract;
Vue.prototype.$keyBroadcastContract = keyBroadcastContract;

new Vue({
  render: (h) => h(App),
}).$mount("#app");
