import { createApp } from "vue";
import "@/assets/index.css";
import App from "@/App.vue";
import router from "@/routes/index";
import { createPinia } from "pinia";

import icons from "@/components/icons/fontAwesome";

const pinia = createPinia();

createApp(App)
  .use(router)
  .use(pinia)
  .component("font-awesome-icon", icons)
  .mount("#app");
