<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <div class="text-center md:text-left">
    <h1 class="text-4xl font-semibold text-emerald-600 mb-2">{{ msg }}</h1>

    <h3 class="text-lg text-gray-700">
      You’ve successfully created a project with
      <a
        href="https://vite.dev/"
        target="_blank"
        rel="noopener"
        class="text-blue-600 hover:underline hover:text-blue-700"
      >
        Vite
      </a>
      +
      <a
        href="https://vuejs.org/"
        target="_blank"
        rel="noopener"
        class="text-green-600 hover:underline hover:text-green-700"
      >
        Vue 3
      </a>.
      What's next?
    </h3>
  </div>
</template>
