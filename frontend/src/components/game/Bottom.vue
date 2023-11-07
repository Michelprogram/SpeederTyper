<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { useRoute } from "vue-router";
import { ref, watch, computed } from "vue";
import { useRoomStore } from "@/store/room";
import { usePlayerStore } from "@/store/player";
import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { webSocketService as socket } from "@/services/socket";

const readyMode = ref(false);
const idRoom = useRoute().params.id as string;

const { currentUser } = storeToRefs(usePlayerStore());
const { currentRoom } = storeToRefs(useRoomStore());

const isOwner = computed((): boolean => {
  return currentRoom.value.created_by == currentUser.value.username;
});

const startGame = () => {
  socket.sendMessage({
    name: "start-game",
    data: {
      id: idRoom,
    },
  });
};

watch(readyMode, (newVal: boolean) => {
  socket.sendMessage({
    name: "set-ready",
    data: { id: idRoom, username: currentUser.value.username, ready: newVal + "" },
  });
});
</script>

<template>
  <div
    class="w-full flex justify-center items-center"
    v-if="!currentRoom.status"
  >
    <Button class="rounded-xl m-auto" v-if="isOwner" @click="startGame"
      >Start game</Button
    >
    <div class="flex justify-center space-x-5" v-else>
      <Switch id="ready-mode" v-model:checked="readyMode" />
      <Label for="ready-mode">Ready</Label>
    </div>
  </div>
</template>
