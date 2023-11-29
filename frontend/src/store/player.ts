import { defineStore } from "pinia";
import { computed, ref } from "vue";

export type Player = {
  username: string;
  position: number;
  isReady: boolean;
};

export const usePlayerStore = defineStore("player", () => {
  const players = ref<Array<string>>([""]);
  const gamePlayed = ref<number>(0);

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
    gamePlayed,
    players,
    currentUser,
    orderedPlayers,
  };
});
