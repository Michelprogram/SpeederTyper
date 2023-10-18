<script setup lang="ts">
import router from "@/routes";
import { state, webSocketService } from "@/services/socket";
import { computed } from "vue";
import Card from "@/components/home/Card.vue";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { usePlayerStore } from "@/store/player";

const styleJoinRoom = computed(() => {
  if (!state.connected)
    return "col-span-2 row-span-1 overflow-y-hidden relative cursor-not-allowed opacity-25";
  return "col-span-2 row-span-1 overflow-y-hidden relative cursor-pointer";
});

const joinRoomByUser = (user: string) => {
  webSocketService.sendMessage({
    name: "join-by-username",
    data: {
      username: user,
    },
  });
};

const store = usePlayerStore();
</script>

<template>
  <div class="flex gap-10 pt-6 m-20 font-mono">
    <div class="flex flex-col gap-8 w-2/5">
      <h1 class="font-bold text-6xl break-words leading-snug">
        Who's the fatest typer ?
      </h1>
      <p class="text-gray-500">
        This platforme is a game to challenge your friend in a race to know who
        of you is the fatest typer
      </p>
      <div class="flex">
        <div class="flex flex-col gap-2">
          <span class="font-bold text-3xl text-orange-600">21</span>
          <p class="w-2/3 font-bold">Number of player connected</p>
        </div>
        <div class="flex flex-col gap-2">
          <span class="font-bold text-3xl text-orange-600">200</span>
          <p class="w-2/3 font-bold">Game played since</p>
        </div>
      </div>
      <div class="w-2/3">
        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Find a player" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup v-for="user in store.orderedPlayers">
              <SelectItem :value="user" @click="joinRoomByUser(user)">
                <p class="flex items-center">
                  {{ user }}
                  <span
                    class="w-3 h-3 ml-3 bg-green-400 rounded-full inline-block"
                  ></span>
                </p>
              </SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
    </div>
    <div class="grid grid-cols-2 grid-rows-2 gap-x-5 gap-y-5 w-3/5">
      <div :class="styleJoinRoom" @click="router.push('/join')">
        <img
          src="https://images.unsplash.com/photo-1511512578047-dfb367046420?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2942&q=80"
          alt=""
          srcset=""
          class="rounded-3xl object-cover w-full h-[350px]"
        />
        <p class="absolute text-white top-5 text-center text-3xl w-full">
          Join room
        </p>
      </div>
      <div class="col-span-1 row-span-1 relative">
        <Card
          pictureURL="https://images.unsplash.com/photo-1531857097499-b97a9ecf225d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2940&q=80"
          icon="plus"
          text="Start a game"
          URL="/offline"
        />
      </div>
      <div class="col-span-1 row-span-1 relative">
        <Card
          pictureURL="https://images.unsplash.com/photo-1533237264985-ee62f6d342bb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2938&q=80"
          icon="bars"
          text="Scoreboard"
          URL="/"
        />
      </div>
    </div>
  </div>
</template>
