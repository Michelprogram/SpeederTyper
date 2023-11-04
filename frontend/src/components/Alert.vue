<script setup lang="ts">
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { useAlertStore } from "@/store/alert";

import { storeToRefs } from "pinia";

const { isVisible, title, description } = storeToRefs(useAlertStore());
</script>

<template>
  <AlertDialog :open="isVisible">
    <AlertDialogContent class="bg-black border-black">
      <AlertDialogHeader>
        <AlertDialogTitle class="text-white">{{ title }}</AlertDialogTitle>
        <AlertDialogDescription class="text-gray-50">
          {{ description }}
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogAction @click="() => isVisible = !isVisible">Continue</AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
