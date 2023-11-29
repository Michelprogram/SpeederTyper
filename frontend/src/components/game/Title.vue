<script setup lang="ts">
import { useRoomStore } from "@/store/room";
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { Copy } from "lucide-vue-next";

const { currentRoom } = storeToRefs(useRoomStore());

const copy = computed(() => {
  try {
    navigator.clipboard.writeText(currentRoom.value.id.toString());
  } catch {
    console.warn("Clipboard not supported");
  }
});
</script>
<template>
  <div class="text-center">
    <h1 class="text-6xl font-bold tracking-wider">Speeder typer</h1>
    <p class="mt-5 flex justify-center items-center">
      <span class="m-3 text-xs">{{ $route.params.id }}</span>
      <Copy class="cursor-pointer" :click="copy" :size="20" />
    </p>
  </div>
</template>
