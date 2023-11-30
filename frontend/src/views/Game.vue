<script lang="ts" setup>
import { onMounted, onUnmounted, ref, watch } from "vue";
import { webSocketService as socket } from "@/services/socket";
import { useRoute } from "vue-router";
import Bottom from "@/components/game/Bottom.vue";
import Title from "@/components/game/Title.vue";
import Editor from "@/components/game/Editor.vue";
import Users from "@/components/game/Users.vue";
import Score from "@/components/game/Score.vue";
import TextInfo from "@/components/game/TextInfo.vue";
import { storeToRefs } from "pinia";
import { useRoomStore } from "@/store/room";
import { Status } from "@/services/types";

const { currentRoom, playerReady } = storeToRefs(useRoomStore());

const idRoom = useRoute().params.id as string;

const time = ref<number>(5);

const waiting = ref<boolean>(true);

let countDown: NodeJS.Timeout;

watch(
  () => currentRoom.value.status,
  () => {
    countDown = setInterval(() => {
      time.value -= 1;
      if (time.value === 0) {
        waiting.value = false;
      }
    }, 1000);
  }
);

watch(waiting, () => {
  clearInterval(countDown);
});

onMounted(() => {
  socket.sendMessage({ name: "join-room", data: { id: idRoom } });
});

onUnmounted(() => {
  socket.sendMessage({ name: "leave-room", data: { id: idRoom } });
});
</script>

<template>
  <div class="font-mono flex flex-col gap-5">
    <Score />
    <Title />
    <div class="flex w-10/12 m-auto gap-5 h-fit">
      <Users />
      <div
        class="w-full h-fit border-2 border-solid border-secondary rounded-xl"
      >
        <TextInfo />
        <Transition name="fade">
          <div class="flex justify-center" v-if="waiting">
            <p v-if="currentRoom.status == Status.Waiting">
              Game should start soon, waiting everyone to be ready,
              {{ playerReady.length }} on {{ currentRoom.users.length }} player
              are ready.
            </p>
            <p v-else>Game should start in {{ time }} secondes.</p>
          </div>
          <div class="flex" v-else>
            <Editor :code="currentRoom.text" />
          </div>
        </Transition>
      </div>
    </div>
    <Bottom />
  </div>
</template>

<style>
/* we will explain what these classes do next! */
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
