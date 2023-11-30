<script setup lang="ts">
import { webSocketService as socket } from "@/services/socket";
import { useRoomStore } from "@/store/room";
import { usePlayerStore } from "@/store/player";
import { onMounted, computed, ref, CSSProperties } from "vue";
import { storeToRefs } from "pinia";

const { currentUser } = storeToRefs(usePlayerStore());
const { currentRoom } = storeToRefs(useRoomStore());

const props = defineProps({
  code: {
    required: true,
    type: String,
  },
});

const lettersRefs = ref<Array<HTMLSpanElement>>([]);
const cursor = ref<HTMLDivElement>();

let textSize = 0;

const splitedLine = computed(() => {
  return props.code.split(/\n/gm);
});

const cleanedLines = computed(() => {
  return splitedLine.value.map((line) => {
    line = line.replace(/\t/g, "").trim();

    textSize += line.length;

    return line;
  });
});

const indentationCls = (line: string): CSSProperties => {
  const factor = 1.25;

  const tabulations = line.match(/^\t+/g);
  const startsWithSpace = line.match(/^ +/g);

  //Match tabulation
  if (tabulations) {
    return { marginLeft: `${factor * tabulations[0].length - 1}rem` };
  }

  if (startsWithSpace) {
    return { marginLeft: `${factor * startsWithSpace[0].length - 1}rem` };
  }

  return {};
};

const getSpanAtPosition = (): HTMLSpanElement => {
  const position = currentUser.value.position;
  return lettersRefs.value[position];
};

const typer = (event: KeyboardEvent) => {
  const letter = event.key;
  console.log(letter);
  const spanAtPosition = getSpanAtPosition();

  if (letter === spanAtPosition.textContent) {
    forwardLetter(spanAtPosition);
    positionCursor();
  } else {
    wrongLetter(spanAtPosition);
  }
};

const forwardLetter = (span: HTMLSpanElement) => {
  socket.sendMessage({
    name: "typing-game",
    data: {
      id: currentRoom.value.id,
      username: currentUser.value.username,
    },
  });

  span.classList.remove("letter-unvalid");
  span.classList.add("letter-valid");

  currentUser.value.position++;
  console.log(currentUser.value.position, textSize);
  if (currentUser.value.position === textSize) {
    socket.sendMessage({
      name: "end-game",
      data: {
        id: currentRoom.value.id,
      },
    });
  }
};

const wrongLetter = (span: HTMLSpanElement) => {
  span.classList.add("letter-unvalid");
};

const positionCursor = () => {
  if (cursor.value == undefined) return;

  const span = getSpanAtPosition();

  cursor.value.style.top = `${span.offsetTop + 20}px`;
  cursor.value.style.left = `${span.offsetLeft}px`;

  cursor.value.classList.replace("opacity-0", "opacity-100");
};

const myOwnPercentage = computed((): string => {
  return ((currentUser.value.position / textSize) * 100).toFixed(2);
});

onMounted(() => {
  document.addEventListener("keypress", typer);
});
</script>

<template>
  <div class="w-full relative">
    <div
      class="absolute w-2 h-[3px] bg-black rounded-sm opacity-0 transition-all"
      ref="cursor"
    ></div>
    <p class="text-center">{{ myOwnPercentage }}%</p>
    <div v-for="(line, i) in cleanedLines" class="flex mb-1">
      <p class="opacity-80 mx-3 text-primary">{{ i + 1 }}</p>
      <p :style="indentationCls(splitedLine[i])">
        <span
          v-for="(letter, j) in line"
          :key="i - j"
          ref="lettersRefs"
          class="letter"
        >
          {{ letter }}
        </span>
      </p>
    </div>
  </div>
</template>

<style scoped>
.letter {
  opacity: 0.35;

  transition: all 0.15s;
}

.letter-valid {
  opacity: 1;
  color: rgb(21 128 61);
}

.letter-unvalid {
  opacity: 1;
  color: rgb(180 83 9);
}
</style>
