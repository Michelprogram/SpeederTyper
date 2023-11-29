<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { webSocketService } from "@/services/socket";
import { usePlayerStore } from "@/store/player";
import { computed, onBeforeUnmount } from "vue";

const store = usePlayerStore();

onBeforeUnmount(() => {
  webSocketService.disconnect();
});

const wikiwand = computed(() => {
  if (store.currentUser.username === "Not connected")
    return "https://www.wikiwand.com";
  const pseudo = store.currentUser.username
    .split(/ /g)
    .slice(0, -1)
    .join("_")
    .toLowerCase();
  return `https://www.wikiwand.com/en/${pseudo}`;
});
</script>

<template>
  <div class="flex justify-around align-middle items-center p-5">
    <RouterLink to="/" class="text-5xl">
      <img
        src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Keyboard.png"
        alt="Ring Buoy"
        class="w-1/4"
      />
    </RouterLink>
    <div class="flex justify-between items-center w-2/4 font-mono">
      <div class="relative">
        <a
          class="cursor-pointer underline underline-offset-4"
          :href="wikiwand"
          target="_blank"
        >
          {{ store.currentUser.username }}
        </a>
        <img
          src="https://wikiwandv2-19431.kxcdn.com/icons/icon-180x180.png"
          alt="wikiwand logo"
          class="absolute w-8 h-8 -z-10 -left-6 top-0 -rotate-45"
        />
      </div>
      <RouterLink to="/">
        <p class="cursor-pointer">HOME</p>
      </RouterLink>
      <RouterLink to="/about">
        <p class="cursor-pointer">HOW DOES IT WORK</p>
      </RouterLink>
      <a href="https://dorian-gauron.com" target="_blank">
        <Button>Contact</Button>
      </a>
    </div>
  </div>
</template>
