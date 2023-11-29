<script setup lang="ts">
import { Button } from "@/components/ui/button";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import { webSocketService as socket } from "@/services/socket";
import { useToast } from "@/components/ui/toast/use-toast";
import { useRoomStore, Room } from "@/store/room";
import { useRouter } from "vue-router";

const { toast } = useToast();

const store = useRoomStore();

const router = useRouter();

const createRoom = () => {
  socket.sendMessage({ name: "create-room" });
};

const joinRoom = (room: Room) => {
  if (room.status) {
    toast({
      title: "Can't join the room",
      description: "You can't join the room the game already started.",
    });
  } else {
    router.push({ name: "Game", params: { id: room.id } });
  }
};
</script>

<template>
  <div class="flex justify-center items-center flex-col gap-5 font-mono">
    <h1 class="font-bold text-6xl leading-snug tracking-wider">Rooms</h1>
    <div class="m-auto w-3/4 flex flex-col gap-5">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead class="w-[100px]">Room id</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Username</TableHead>
            <TableHead class="text-right">Players</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow
            v-for="room in store.orderedRooms"
            class="cursor-pointer"
            @click="joinRoom(room)"
          >
            <TableCell class="font-medium"
              >#{{ room.id.slice(0, 4) }}</TableCell
            >
            <TableCell>{{ room.status ? "In game..." : "In lobby" }}</TableCell>
            <TableCell>{{ store.firstPlayers(room.id) }}</TableCell>
            <TableCell class="text-right">{{ room.users.length }}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
      <Button @click="createRoom" class="rounded-xl w-fit">Create room</Button>
    </div>
  </div>
</template>
