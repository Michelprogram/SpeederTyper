import { Player } from "@/store/player";
import { Room, Score } from "@/store/room";

export type StatsAppEvent = {
  players: Array<string>;
  game: number;
};

export type TextInfoEvent = {
  repoUrl: string;
  repoNametext: string;
  fullUrl: string;
  language: string;
  fileName: string;
};

export type RoomsInfoEvent = {
  stats: Array<Room>;
};

export type RoomInfoEvent = {
  created_by: string;
  id: string;
  status: Status;
  text: string;
  users: Array<Player>;
};

export enum Status {
  Waiting = 0,
  Gaming = 1,
  GenerateTexting = 2,
  Finish = 3,
}
