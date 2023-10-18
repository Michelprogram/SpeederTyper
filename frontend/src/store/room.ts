import { defineStore } from "pinia";
import { computed, ref } from "vue";

type Room = {
  id: string;
  status: boolean; //true in game false not running
  players: Array<Player>;
};

type Player = {
  Ws: {
    MaxPayloadBytes: number;
    PayloadType: number;
  };
  username: string;
};

export const useRoomStore = defineStore("room", () => {
  const rooms = ref(new Array<Room>());

  const players = (id: string): Array<string> => {
    const room = rooms.value.find((r) => r.id == id);

    if (room == undefined) return [];

    return room.players.map((p) => p.username);
  };

  const firstPlayers = (id: string): string => {
    return players(id).slice(0, 3).join(",") + "...";
  };

  const orderedRooms = computed((): Array<Room> => {
    if (rooms.value == undefined) return [];
    return rooms.value.sort((a, b) => (a > b ? 1 : -1));
  });

  return {
    players,
    firstPlayers,
    orderedRooms,
    rooms,
  };
});
