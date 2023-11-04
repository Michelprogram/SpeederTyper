<script lang="ts" setup>
import { onMounted, onUnmounted } from "vue";
import { webSocketService as socket } from "@/services/socket";
import { useRoute } from "vue-router";
import { computed } from "vue";
import Bottom from "@/components/game/Bottom.vue";
import ListUser from "@/components/game/ListUser.vue";
import Words from "@/components/game/Words.vue";

const idRoom = useRoute().params.id as string;

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
      <ListUser />
      <Words />
    </div>
    <Bottom />
  </div>
</template>
