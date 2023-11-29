<script setup lang="ts">
import { state, webSocketService as websocket } from "@/services/socket";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { Plus, Signal } from "lucide-vue-next";
import { storeToRefs } from "pinia";
import { usePlayerStore } from "@/store/player";
import { onMounted } from "vue";

const { players, gamePlayed, orderedPlayers } = storeToRefs(usePlayerStore());

const joinRoomByUser = (user: string) => {
  websocket.sendMessage({
    name: "join-by-username",
    data: {
      username: user,
    },
  });
};

onMounted(() => {
  websocket.sendMessage({
    name: "stats-app",
  });
});
</script>

<template>
  <div class="flex gap-10 pt-4 m-20 font-mono">
    <div class="flex flex-col gap-8 w-2/5">
      <h1 class="font-bold text-5xl break-words leading-snug">
        Who's the fastest typer ?
      </h1>
      <p class="text-gray-500">
        This platforme is a game to challenge your friend in a race to know who
        of you is the fastest typer
      </p>
      <div class="flex">
        <div class="flex flex-col gap-2">
          <span class="font-bold text-3xl text-orange-600">{{
            players === undefined ? 0 : players.length
          }}</span>
          <p class="w-2/3 font-bold">Number of player connected</p>
        </div>
        <div class="flex flex-col gap-2">
          <span class="font-bold text-3xl text-orange-600">{{
            gamePlayed
          }}</span>
          <p class="w-2/3 font-bold">Game played since last deploy</p>
        </div>
      </div>
      <div class="w-2/3">
        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Find a player" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem
                v-for="user in orderedPlayers"
                :value="user"
                @click="joinRoomByUser(user)"
              >
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
      <div class="w-full col-span-2 row-span-1 flex">
        <img
          src="/undraw_pair_programming.svg"
          alt="Pair programming"
          class="h-44 w-44"
        />
        <div
          class="bg-white rounded-b p-4 flex flex-col justify-between leading-normal"
        >
          <div>
            <p
              class="text-sm text-gray-500 flex items-center gap-3 justify-start"
            >
              <Signal :size="20" v-if="state.connected" />
              <Plus v-else :size="20" class="rotate-45" />
              Online content
            </p>
            <div class="text-gray-900 font-bold text-xl mb-2">
              You think your're the best at typing ?
            </div>
            <p class="text-gray-700 text-base">
              Challenge your friends or random through rooms. Check your
              capabilities to typing in real world application context.
            </p>
            <RouterLink
              to="/join"
              :class="
                state.connected
                  ? ''
                  : 'cursor-not-allowed opacity-25 pointer-events-none'
              "
            >
              <Button class="flex items-center w-fit mt-4"> Join room </Button>
            </RouterLink>
          </div>
        </div>
      </div>
      <div class="max-w-sm rounded overflow-hidden col-span-1 row-span-1">
        <img
          class="w-28 h-28"
          src="/undraw_programmer.svg"
          alt="Sunset in the mountains"
        />
        <div class="py-4">
          <div class="font-bold text-xl mb-2">Train alone</div>
          <p class="text-gray-700 text-base">
            Challenge yourself in infinity mode, and deal with yourself.
          </p>
        </div>
        <!--         <RouterLink to="/offline">
 -->
        <Button disabled>Start a game offline</Button>
        <!--         </RouterLink>
 -->
      </div>
      <div class="max-w-sm rounded overflow-hidden col-span-1 row-span-1">
        <img
          class="w-28 h-28"
          src="/undraw_chart.svg"
          alt="Sunset in the mountains"
        />
        <div class="py-4">
          <div class="font-bold text-xl mb-2">Check the leaderboard</div>
          <p class="text-gray-700 text-base">
            You'd like to see who are the top players ?
          </p>
        </div>
        <!--         <RouterLink to="/">
 -->
        <Button disabled>Scoreboard</Button>
        <!--         </RouterLink>
 -->
      </div>
    </div>
  </div>
</template>
