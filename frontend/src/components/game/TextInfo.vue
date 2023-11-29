<script setup lang="ts">
import { useTextInfoStore } from "@/store/textinfo";
import { storeToRefs } from "pinia";
import { computed } from "vue";

const { textinfo, iconLanguage } = storeToRefs(useTextInfoStore());

const cssLogo = computed(() => {
  return textinfo.value.language === "" ? "animate-pulse bg-slate-700" : "";
});
</script>

<template>
  <div class="flex items-center justify-center gap-5 w-full p-5">
    <img
      id="logo"
      :class="`w-8 h-8 rounded-sm ${cssLogo}`"
      :src="iconLanguage"
      alt=""
    />
    <div
      v-if="textinfo.repoUrl === ''"
      class="animate-pulse h-2 bg-slate-700 rounded w-1/4"
    ></div>
    <a v-else class="cursor-pointer" :href="textinfo.fullUrl" target="_blank">
      {{ textinfo.repoUrl }}
    </a>

    <div
      v-if="textinfo.fileName === ''"
      class="animate-pulse h-2 bg-slate-700 rounded w-1/5"
    ></div>
    <p v-else>{{ textinfo.fileName }}</p>
  </div>
</template>
