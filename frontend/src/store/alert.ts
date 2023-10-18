import { defineStore } from "pinia";
import { ref } from "vue";

export const useAlertStore = defineStore("alert", () => {
  const isVisible = ref(false);
  const title = ref("");
  const description = ref("");

  return {
    isVisible,
    title,
    description,
  };
});
