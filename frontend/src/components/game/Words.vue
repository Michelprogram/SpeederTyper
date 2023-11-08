<script lang="ts" setup>
import { webSocketService as socket } from "@/services/socket";
import { Status, useRoomStore } from "@/store/room";
import { usePlayerStore } from "@/store/player";
import { storeToRefs } from "pinia";
import { watch, ref, onMounted, computed } from "vue";

const { currentRoom, playerReady } = storeToRefs(useRoomStore());

const { currentUser } = storeToRefs(usePlayerStore());

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

const typer = (event: KeyboardEvent) => {
  if (currentRoom.value.status !== Status.Gaming && waiting.value) return;

  const letter = event.key;
  const position = currentUser.value.position;
  const characterAtPosition = currentRoom.value.text.at(position);
  const letterElement = document.querySelector(
    `#letter-${position}`
  ) as HTMLSpanElement;

  if (letter == characterAtPosition) {
    forwardLetter(letterElement);
  } else {
    wrongLetter(letterElement);
  }
};

const forwardLetter = (element: HTMLSpanElement) => {
  socket.sendMessage({
    name: "typing-game",
    data: {
      id: currentRoom.value.id,
      username: currentUser.value.username,
    },
  });

  element.classList.remove("text-red-400");
  element.classList.add("text-green-400");

  currentUser.value.position++;

  if (currentUser.value.position == currentRoom.value.text.length) {
    socket.sendMessage({
      name: "end-game",
      data: {
        id: currentRoom.value.id,
      },
    });
  }
};

const wrongLetter = (element: HTMLSpanElement) => {
  element.classList.add("text-red-400", "decoration-red-400");
};

const myOwnPercentage = computed((): string => {
  const sizeText = useRoomStore().currentRoom.text.length;

  return ((currentUser.value.position / sizeText) * 100).toFixed(2);
});

onMounted(() => {
  document.addEventListener("keypress", typer);
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
        <p class="">
          <span
            v-for="(letter, index) in currentRoom.text"
            :key="index"
            :id="'letter-' + index"
            >{{ letter }}</span
          >
        </p>
        <p>{{ myOwnPercentage }}%</p>
      </div>
    </div>
  </div>
</template>
