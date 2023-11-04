<script lang="ts" setup>
import { useRoomStore } from "@/store/room";
import { storeToRefs } from "pinia";
import { watch, ref } from "vue";

const { currentRoom, playerReady } = storeToRefs(useRoomStore());

const waiting = ref(true);

const time = ref(5);

let countDown: NodeJS.Timeout;

watch(
  () => currentRoom.value.status,
  () => {
    countDown = setInterval(() => {
      time.value -= 1;
      if (time.value == 0) {
        waiting.value = false;
      }
    }, 1000);
  }
);

watch(waiting, () => {
  clearInterval(countDown);
});
</script>

<template>
  <div class="w-full border-2 border-solid border-secondary rounded-xl">
    <div v-if="waiting">
      <div class="animate-pulse flex space-x-4 p-5">
        <div class="flex-1 space-y-3 py-1">
          <div class="h-2 bg-slate-700 rounded"></div>
          <div class="space-y-2">
            <div class="grid grid-cols-3 gap-4">
              <div class="h-2 bg-slate-700 rounded col-span-2"></div>
              <div class="h-2 bg-slate-700 rounded col-span-1"></div>
            </div>
            <div class="h-2 bg-slate-700 rounded space-y-2"></div>
            <div class="h-2 bg-slate-700 rounded"></div>
            <div class="grid grid-cols-3 gap-4">
              <div class="h-2 bg-slate-700 rounded col-span-1"></div>
              <div class="h-2 bg-slate-700 rounded col-span-2"></div>
            </div>
            <div class="h-2 bg-slate-700 rounded"></div>
          </div>
        </div>
      </div>
      <div class="p-5">
        <p v-if="!currentRoom.status">
          Game should start waiting everyone to be ready,
          {{ playerReady.length }} on {{ currentRoom.users.length }} player are
          ready.
        </p>
        <p v-else>Game should start in {{ time }} secondes.</p>
      </div>
    </div>
    <div v-else>
      <div class="p-5">
        <p>{{ currentRoom.text }}</p>
      </div>
    </div>
  </div>
</template>
