import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { Player } from "./room";

export const usePlayerStore = defineStore("player", () => {
  const players = ref(new Array<string>());
  const currentUser = ref<Player>({
    username: "Not connected",
    position: 0,
    isReady: false,
  });

  const orderedPlayers = computed((): Array<string> => {
    return players.value
      .sort((a, b) => (a > b ? 1 : -1))
      .filter((u) => u != currentUser.value.username);
  });

  return {
    players,
    currentUser,
    orderedPlayers,
  };
});
