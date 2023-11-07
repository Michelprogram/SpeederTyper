<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { webSocketService } from "@/services/socket";
import { usePlayerStore } from "@/store/player";
import { onBeforeUnmount } from "vue";

const store = usePlayerStore();

onBeforeUnmount(() => {
  webSocketService.disconnect();
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
      <p class="cursor-pointer">{{ store.currentUser.username }}</p>
      <RouterLink to="/">
        <p class="cursor-pointer">HOME</p>
      </RouterLink>
      <p class="cursor-pointer">OUR TEAM</p>
      <Button>
        <a href="https://dorian-gauron.com" target="_blank">Contact</a>
      </Button>
    </div>
  </div>
</template>
