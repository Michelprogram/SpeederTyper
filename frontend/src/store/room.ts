import { defineStore } from "pinia";
import { computed, ref } from "vue";

export enum Status {
  Waiting = 0,
  Gaming = 1,
  Finish = 3,
}

export type Room = {
  id: string;
  text: string;
  created_by: string;
  status: Status;
  users: Array<Player>;
};

export type Player = {
  username: string;
  position: number;
  isReady: boolean;
};

export type Score = {
  wordCompleted: number;
  letterCompleted: number;
  characterPerSecond: number;
  classment: number;
  username: string;
};

export const useRoomStore = defineStore("room", () => {
  const rooms = ref(new Array<Room>());

  const scores = ref(new Array<Score>());

  const currentRoom = ref<Room>({
    id: "",
    text: "",
    created_by: "",
    status: Status.Waiting,
    users: new Array<Player>(),
  });

  const players = (id: string): Array<string> => {
    const room = rooms.value.find((r) => r.id === id);

    if (room === undefined) return [];

    return room.users.map((p) => p.username);
  };

  const firstPlayers = (id: string): string => {
    return players(id).slice(0, 3).join(",") + "...";
  };

  const orderedRooms = computed((): Array<Room> => {
    if (rooms.value === undefined) return [];
    return rooms.value.sort((a, b) => (a > b ? 1 : -1));
  });

  const playerReady = computed((): Array<Player> => {
    if (currentRoom.value.users === undefined) return [];
    return currentRoom.value.users.filter((u) => u.isReady);
  });

  return {
    players,
    firstPlayers,
    scores,
    playerReady,
    orderedRooms,
    currentRoom,
    rooms,
  };
});
