<script lang="ts" setup>
import { useRoomStore } from "@/store/room";
import { storeToRefs } from "pinia";
import { computed } from "vue";
const { currentRoom } = storeToRefs(useRoomStore());

const users = computed(() => {
  return currentRoom.value.users;
});
</script>

<template>
  <div class="border-2 border-solid border-secondary rounded-xl w-1/4">
    <h4 class="mb-4 mt-4 text-xl font-medium leading-none text-center">
      Players - {{ users.length }}
    </h4>
    <div class="flex flex-col items-center">
      <div v-for="{ username, isReady } in users" :key="username">
        <div class="text-sm mb-3">
          <p class="flex items-center gap-3">
            <span
              :class="`w-3 h-3 rounded-full block ${
                isReady ? 'bg-green-400' : 'bg-red-400'
              }`"
            ></span>
            <span>{{ username }}</span>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
