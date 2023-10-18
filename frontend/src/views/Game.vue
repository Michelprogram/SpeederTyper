<script lang="ts" setup>
import { onMounted, onUnmounted } from "vue";
import { webSocketService as socket } from "@/services/socket";
import { useRoute } from "vue-router";
import { computed } from "vue";
import { Button } from "@/components/ui/button";
import { useRoomStore } from "@/store/room";

const idRoom = useRoute().params.id as string;

const store = useRoomStore();

const copy = computed(() => {
  navigator.clipboard.writeText(idRoom.toString());
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
    <div class="text-center">
      <h1 class="text-6xl font-bold tracking-wider">Speeder typer</h1>
      <p class="mt-5">
        <span class="m-3 text-xs">{{ $route.params.id }}</span>
        <font-awesome-icon
          icon="fa-copy"
          class="cursor-pointer"
          @click="copy"
        />
      </p>
    </div>
    <div class="flex w-5/6 m-auto gap-5 h-96">
      <div class="border-2 border-solid border-secondary rounded-xl w-1/4">
        <h4 class="mb-4 text-lg font-medium leading-none text-center">
          Players - {{ store.players(idRoom).length }}
        </h4>
        <div v-for="user in store.players(idRoom)" :key="user">
          <div class="text-sm mb-3">
            <p>
              {{ user }}
            </p>
          </div>
        </div>
      </div>
      <div
        class="w-full border-2 border-solid border-secondary rounded-xl"
      ></div>
    </div>
    <div class="w-full flex justify-center items-center">
      <Button class="rounded-xl m-auto">Start game</Button>
    </div>
  </div>
</template>
