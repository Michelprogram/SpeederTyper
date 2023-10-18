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
import { useRoomStore } from "@/store/room";

const store = useRoomStore();

const createRoom = () => {
  socket.sendMessage({ name: "join-room" });
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
            @click="$router.push({ name: 'Game', params: { id: room.id } })"
          >
            <TableCell class="font-medium"
              >#{{ room.id.slice(0, 4) }}</TableCell
            >
            <TableCell>{{ room.status ? "In game..." : "In lobby" }}</TableCell>
            <TableCell>{{ store.firstPlayers(room.id) }}</TableCell>
            <TableCell class="text-right">{{ room.players.length }}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
      <Button @click="createRoom" class="rounded-xl w-fit">Create room</Button>
    </div>
  </div>
</template>
