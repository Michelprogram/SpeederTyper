import router from "@/routes";
import { useToast } from "@/components/ui/toast/use-toast";
import { storeToRefs } from "pinia";
import { usePlayerStore } from "@/store/player";
import { Score, useRoomStore } from "@/store/room";
import { useTextInfoStore } from "@/store/textinfo";
import {
  Status,
  StatsAppEvent,
  TextInfoEvent,
  RoomsInfoEvent,
  RoomInfoEvent,
} from "./types";

export enum EventTypes {
  "pong" = "pong",
  "username" = "username",
  "rooms-info" = "rooms-info",
  "room-info" = "room-info",
  "room-created" = "room-created",
  "room-join-by-username" = "room-join-by-username",
  "game-cant-start" = "game-cant-start",
  "game-can-start" = "game-can-start",
  "game-end" = "game-end",
  "text-info" = "text-info",
  "stats-app" = "stats-app",
}

export const eventListeners: {
  [key in EventTypes]: (data: any) => void;
} = {
  [EventTypes.pong]: (pong: string) => {
    console.log("Message from ping", pong);
  },
  [EventTypes["rooms-info"]]: ({ stats }: RoomsInfoEvent) => {
    const { rooms } = storeToRefs(useRoomStore());
    rooms.value = stats;
  },
  [EventTypes["room-created"]]: function (id: string): void {
    router.push({ name: "Game", params: { id: id } });
  },
  [EventTypes["username"]]: function (username: string): void {
    const { currentUser } = storeToRefs(usePlayerStore());
    currentUser.value.username = username;
  },
  [EventTypes["room-join-by-username"]]: function (id: string): void {
    if (id === "-1") {
      const { toast } = useToast();

      toast({
        title: "Not in game.",
        description: "The user you trying to join is not in game.",
      });

      return;
    }

    router.push({ name: "Game", params: { id: id } });
  },
  [EventTypes["room-info"]]: function (room: RoomInfoEvent): void {
    const { currentRoom } = storeToRefs(useRoomStore());
    currentRoom.value = room;
  },
  [EventTypes["game-cant-start"]]: function (_: any): void {
    const { toast } = useToast();

    toast({
      title: "Everyone should be ready to start the game",
      description: "You can't run the game because everyone should be ready.",
    });
  },
  [EventTypes["game-can-start"]]: function (_: any): void {
    const { currentRoom } = storeToRefs(useRoomStore());

    currentRoom.value.status = Status.Gaming;
  },
  [EventTypes["game-end"]]: function (data: Array<Score>): void {
    const { scores, currentRoom } = storeToRefs(useRoomStore());

    scores.value = data;
    currentRoom.value.status = Status.Finish;
  },
  [EventTypes["text-info"]]: function (data: TextInfoEvent): void {
    const { textinfo } = storeToRefs(useTextInfoStore());
    textinfo.value = data;
  },
  [EventTypes["stats-app"]]: (stats: StatsAppEvent) => {
    const { players, gamePlayed } = storeToRefs(usePlayerStore());

    players.value = stats.players;
    gamePlayed.value = stats.game;
  },
};
