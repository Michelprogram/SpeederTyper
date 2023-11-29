<script lang="ts" setup>
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";

import { useRoomStore } from "@/store/room";
import { Player, usePlayerStore } from "@/store/player";
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { useRouter } from "vue-router";
import { Status } from "@/services/types";

const { scores, currentRoom } = storeToRefs(useRoomStore());
const { currentUser } = storeToRefs(usePlayerStore());

const router = useRouter();

const open = computed(() => currentRoom.value.status === Status.Finish);

const isWinner = computed(
  () => scores.value.at(0)?.username === currentUser.value.username
);

const resetCurrentRoom = () => {
  currentRoom.value = {
    id: "",
    text: "",
    created_by: "",
    status: Status.Waiting,
    users: new Array<Player>(),
  };

  //Reset user with websocket

  router.push({ name: "Rooms" });
};
</script>

<template>
  <AlertDialog :open="open">
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle v-if="isWinner" class="flex items-center p-2">
          <img
            src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Travel%20and%20places/Fire.png"
            alt=""
            class="h-12"
          />
          <p>It's lit you have won the game.</p>
        </AlertDialogTitle>
        <AlertDialogTitle v-else class="flex items-center p-2">
          <img
            src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Smilies/Zzz.png"
            alt=""
            class="h-12"
          />
          <p>Arff you have lost the game.</p>
        </AlertDialogTitle>
        <AlertDialogDescription class="p-2">
          <table class="table-auto">
            <thead>
              <tr>
                <th class="w-1/3">Classment</th>
                <th class="w-1/3">Username</th>
                <th class="w-1/6">CPS</th>
                <th class="w-2/3">Letters</th>
                <th class="w-2/3">Words</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="score in scores" :key="score.classment">
                <td>
                  <span class="font-bold">{{ score.classment }}</span>
                </td>
                <td>{{ score.username }}</td>
                <td>{{ score.characterPerSecond.toFixed(2) }}</td>
                <td>{{ score.letterCompleted }}</td>
                <td>{{ score.wordCompleted }}</td>
              </tr>
            </tbody>
          </table>
          <p class="pt-2">I hope you enjoyed the game :)</p>
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogAction @click="resetCurrentRoom"
          >Continue</AlertDialogAction
        >
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
