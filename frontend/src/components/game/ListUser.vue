<script lang="ts" setup>
import { Progress } from "@/components/ui/progress";
import { useRoomStore } from "@/store/room";
import { storeToRefs } from "pinia";
import { computed, ref } from "vue";
const { currentRoom } = storeToRefs(useRoomStore());

const percentage = (position: number): number => {
  return Number.parseFloat(
    ((position / currentRoom.value.text.length) * 100).toFixed(2)
  );
};

const users = ref(currentRoom.value.users);

const usersSortedByPercentage = computed(() => {
  const users = currentRoom.value.users;
  return users.sort((a, b) => {
    return a.position < b.position ? 1 : -1;
  });
});
</script>

<template>
  <div class="border-2 border-solid border-secondary rounded-xl w-1/4">
    <h4 class="mb-4 mt-4 text-xl font-medium leading-none text-center">
      Players - {{ users.length }}
    </h4>
    <div class="flex flex-col items-center">
      <TransitionGroup name="player-list">
        <div
          v-for="{ username, isReady, position } in usersSortedByPercentage"
          :key="username"
        >
          <div class="text-sm mb-3">
            <div class="flex flex-col gap-1">
              <p class="flex items-center gap-3">
                <span
                  :class="`w-3 h-3 rounded-full block ${
                    isReady ? 'bg-green-400' : 'bg-red-400'
                  }`"
                ></span>
                <span>{{ username }}</span>
              </p>
              <Progress
                class="h-2"
                :title="percentage(position)"
                :model-value="percentage(position)"
              />
            </div>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </div>
</template>

<style>
.player-list-move,
.player-list-enter-active,
.player-list-leave-active {
  transition: all 0.5s ease;
}

.player-list-enter-from,
.player-list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.player-list-leave-active {
  position: absolute;
}
</style>
