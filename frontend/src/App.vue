<script setup lang="ts">
import { RouterView } from "vue-router";
import Navbar from "./components/Navbar.vue";
import Alert from "./components/Alert.vue";
</script>

<template>
  <div>
    <Navbar />
    <Alert />
    <RouterView v-slot="{ Component }">
      <Transition name="slide" mode="out-in">
        <component :is="Component" />
      </Transition>
    </RouterView>
  </div>
</template>

<style>
.slide-enter-active {
  transition: 500ms ease all;
}

.slide-leave-active {
  transition: 700ms ease all;
}

.slide-enter-from,
.slide-leave-to {
  opacity: 0;
  transform: translateX(-105%);
}
</style>
