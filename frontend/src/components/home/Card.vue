<script setup lang="ts">
const props = defineProps({
  pictureURL: String,
  icon: String,
  text: String,
  URL: String,
});
</script>

<template>
  <div class="col-span-1 row-span-1 relative">
    <img :src="props.pictureURL" alt="" class="w-[500px] rounded-3xl" />
    <RouterLink :to="props.URL" class="text-5xl">
      <font-awesome-icon
        :icon="props.icon"
        class="absolute p-5 bg-white rounded-full left-5 top-40 text-lg cursor-pointer"
      />
    </RouterLink>
    <p class="absolute text-white top-10 left-5 text-3xl w-2/3">
      {{ props.text }}
    </p>
  </div>
</template>
