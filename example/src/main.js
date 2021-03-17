import Vue from "vue";
import { ethers } from "ethers";

import App from "./App.vue";
import contracts from "./assets/contracts.json";
import configContractMetadata from "./assets/abis/ConfigContract.json";
import batcherContractMetadata from "./assets/abis/BatcherContract.json";
import targetContractMetadata from "./assets/abis/TestTargetContract.json";
import keyBroadcastContractMetadata from "./assets/abis/KeyBroadcastContract.json";
import executorContractMetadata from "./assets/abis/ExecutorContract.json";

require("@/assets/main.scss");

Vue.config.productionTip = false;

let provider = new ethers.providers.Web3Provider(window.ethereum);
let configContract = new ethers.Contract(
  contracts.ConfigContract,
  configContractMetadata.abi,
  provider
);
let batcherContract = new ethers.Contract(
  contracts.BatcherContract,
  batcherContractMetadata.abi,
  provider
);
let targetContract = new ethers.Contract(
  contracts.TargetContract,
  targetContractMetadata.abi,
  provider
);
let keyBroadcastContract = new ethers.Contract(
  contracts.KeyBroadcastContract,
  keyBroadcastContractMetadata.abi,
  provider
);
let executorContract = new ethers.Contract(
  contracts.ExecutorContract,
  executorContractMetadata.abi,
  provider
);

Vue.prototype.$provider = provider;
Vue.prototype.$configContract = configContract;
Vue.prototype.$batcherContract = batcherContract;
Vue.prototype.$targetContract = targetContract;
Vue.prototype.$keyBroadcastContract = keyBroadcastContract;
Vue.prototype.$executorContract = executorContract;

new Vue({
  render: (h) => h(App),
}).$mount("#app");
