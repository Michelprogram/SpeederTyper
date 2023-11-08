import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/join",
    name: "Rooms",
    component: () => import("@/views/Rooms.vue"),
  },
  {
    path: "/offline",
    name: "Offline",
    component: () => import("@/views/Offline.vue"),
  },
  {
    path: "/game/:id",
    name: "Game",
    component: () => import("@/views/Game.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
