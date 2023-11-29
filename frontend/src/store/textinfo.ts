import { TextInfoEvent } from "@/services/types";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useTextInfoStore = defineStore("textinfo", () => {
  const textinfo = ref<TextInfoEvent>({
    repoUrl: "",
    repoNametext: "",
    fullUrl: "",
    language: "",
    fileName: "",
  });

  const iconLanguage = computed((): string => {
    const lower =
      textinfo.value.language === ""
        ? "ie10"
        : textinfo.value.language.toLowerCase();
    return `https://cdn.jsdelivr.net/gh/devicons/devicon/icons/${lower}/${lower}-original.svg`;
  });

  return {
    textinfo,
    iconLanguage,
  };
});
