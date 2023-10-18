import router from "@/routes";

import { useAlertStore } from "@/store/alert";
import { storeToRefs } from "pinia";
import { usePlayerStore } from "@/store/player";
import { useRoomStore } from "@/store/room";

export enum EventTypes {
  "pong" = "pong",
  "username" = "username",
  "room-info" = "room-info",
  "room-created" = "room-created",
  "room-join-by-username" = "room-join-by-username",
}

type WebSocketEvent = any;

export type EventListener = (data: WebSocketEvent) => void;

export const eventListeners: { [key in EventTypes]: EventListener } = {
  [EventTypes.pong]: (data: WebSocketEvent) => {
    console.log("Message from ping", data);
  },
  [EventTypes["room-info"]]: (data: WebSocketEvent) => {
    const { players } = storeToRefs(usePlayerStore());
    const { rooms } = storeToRefs(useRoomStore());

    data = JSON.parse(JSON.stringify(data));

    rooms.value = data.stats;

    players.value = data.users;
  },
  [EventTypes["room-created"]]: function (data: WebSocketEvent): void {
    const idRoom = JSON.parse(JSON.stringify(data));
    router.push({ name: "Game", params: { id: idRoom } });
  },
  [EventTypes["username"]]: function (data: WebSocketEvent): void {
    const { currentUser } = storeToRefs(usePlayerStore());

    currentUser.value = data;
  },
  [EventTypes["room-join-by-username"]]: function (data: WebSocketEvent): void {
    const roomID = data;

    if (roomID == "-1") {
      const store = useAlertStore();

      const { title, isVisible, description } = storeToRefs(store);

      isVisible.value = true;

      title.value = "Not in game.";

      description.value = "The user you trying to join is not in game.";

      return;
    }

    router.push({ name: "Game", params: { id: roomID } });
  },
};
