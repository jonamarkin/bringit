<script setup lang="ts">
const { message } = useToast()
</script>

<template>
  <div v-if="message" class="toast">
    {{ message }}
  </div>
</template>
