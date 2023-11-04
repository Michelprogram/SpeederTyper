import router from "@/routes";

import { useAlertStore } from "@/store/alert";
import { storeToRefs } from "pinia";
import { usePlayerStore } from "@/store/player";
import { useRoomStore } from "@/store/room";

export enum EventTypes {
  "pong" = "pong",
  "username" = "username",
  "rooms-info" = "rooms-info",
  "room-info" = "room-info",
  "room-created" = "room-created",
  "room-join-by-username" = "room-join-by-username",
  "room-start" = "room-start",
}

type WebSocketEvent = any;

export type EventListener = (data: WebSocketEvent) => void;

export const eventListeners: { [key in EventTypes]: EventListener } = {
  [EventTypes.pong]: (data: WebSocketEvent) => {
    console.log("Message from ping", data);
  },
  [EventTypes["rooms-info"]]: (data: WebSocketEvent) => {
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
  [EventTypes["room-info"]]: function (data: WebSocketEvent): void {
    const { currentRoom } = storeToRefs(useRoomStore());

    currentRoom.value = data;
  },
  [EventTypes["room-start"]]: function (everyoneReady: WebSocketEvent): void {
    if (everyoneReady) {
      //TODO start game
    } else {
      const store = useAlertStore();

      const { title, isVisible, description } = storeToRefs(store);

      isVisible.value = true;

      title.value = "Everyone should be ready to start the game";

      description.value =
        "You can't run the game because everyone should be ready.";
    }
  },
};
