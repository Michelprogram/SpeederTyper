<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { RandomWords } from "@/util/util";
import { onMounted, ref, watch } from "vue";

const ids = new Array<string>();

const index = ref(0);
const words = ref(Array<string>());

watch(index, () => {
  const nextElement = document.querySelector(`#${ids[index.value]}`);
  const lighting = document.querySelector(".ligthing-letter");

  if (nextElement != null) nextElement.classList.add("ligthing-letter");
  if (lighting != null) lighting?.classList.remove("ligthing-letter");
});

const computedId = (parentIndex: number, index: number) => {
  const id = `letter-${parentIndex}-${index}`;

  ids.push(id);

  return id;
};

const backLetter = () => {
  index.value--;

  const element = document.querySelector(`#${ids[index.value]}`);

  if (element === null) return;

  element.classList.remove("text-lime-400");
};

const forwardLetter = () => {
  const element = document.querySelector(`#${ids[index.value]}`);
  const wrongLetter = document.querySelector(".text-red-400");

  if (element === null) return;

  if (wrongLetter != null) wrongLetter.classList.remove("text-red-400");

  element.classList.add("text-lime-400");
  index.value++;
};

const wrongLetter = () => {
  const element = document.querySelector(`#${ids[index.value]}`);

  if (element === null) return;

  element.classList.add("text-red-400");
};

const lightingMode = () => {};

const keyDownHandler = (ev: KeyboardEvent) => {
  const element = document.querySelector(`#${ids[index.value]}`);
  const keyPressed = ev.key;

  if (element === null) return;

  if (keyPressed === "Backspace" && index.value > 0) {
    backLetter();
  } else if (element.textContent === keyPressed) {
    forwardLetter();
  } else {
    wrongLetter();
  }
};

onMounted(() => {
  window.addEventListener("keydown", keyDownHandler);
  RandomWords(10).then((w) => {
    words.value = w;
    lightingMode();
  });
});
</script>

<template>
  <div class="flex flex-col items-center">
    <div class="m-auto w-2/3 border-4">
      <h1 class="text-2xl text-center">Game in offline mode</h1>
      <div class="flex m-auto w-2/3">
        <p v-for="(word, i) in words" class="flex">
          <span
            v-for="(letter, j) in word"
            :id="computedId(i, j)"
            class="relative block text-lg"
            >{{ letter }}</span
          >
          <span>&nbsp;</span>
        </p>
      </div>
    </div>
    <Button class="w-fit mt-10">Start game</Button>
  </div>
</template>

<style>
.ligthing-letter::after {
  position: absolute;

  content: "";

  background-color: rgba(103, 100, 100, 0.413);

  top: 0;
  left: 0;

  width: 100%;
  height: 100%;

  animation: lighting infinite 1s ease-in-out;
}

@keyframes lighting {
  0% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>
