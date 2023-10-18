import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";

import Home from "@/views/Home.vue";
import Rooms from "@/views/Rooms.vue";
import Offline from "@/views/Offline.vue";
import Game from "@/views/Game.vue";

//const PREFIX = "/postman-like";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/join",
    name: "Rooms",
    component: Rooms,
  },
  {
    path: "/offline",
    name: "Offline",
    component: Offline,
  },
  {
    path: "/game/:id",
    name: "Game",
    component: Game,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
