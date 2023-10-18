import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const usePlayerStore = defineStore("player", () => {
  const players = ref(new Array<string>());
  const currentUser = ref("Not Connected");

  const orderedPlayers = computed((): Array<string> => {
    return players.value
      .sort((a, b) => (a > b ? 1 : -1))
      .filter((u) => u != currentUser.value);
  });

  return {
    players,
    currentUser,
    orderedPlayers,
  };
});
