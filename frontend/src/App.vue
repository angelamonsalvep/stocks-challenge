<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <header class="flex flex-col md:flex-row items-center justify-between p-6 bg-gray-50 shadow-sm">
    <div class="flex flex-col items-center md:flex-row md:space-x-6">
      <img
        alt="Vue logo"
        class="w-28 h-28 mb-4 md:mb-0"
        src="@/assets/logo.svg"
      />

      <HelloWorld msg="You did it!" />
    </div>

    <nav class="mt-4 md:mt-0 flex space-x-4 text-gray-600 text-sm md:text-base">
      <RouterLink
        to="/"
        class="hover:text-blue-600 transition-colors"
        active-class="text-blue-700 font-semibold"
      >
        Home
      </RouterLink>
      <RouterLink
        to="/about"
        class="hover:text-blue-600 transition-colors"
        active-class="text-blue-700 font-semibold"
      >
        About
      </RouterLink>
    </nav>
  </header>

  <main class="p-6">
    <RouterView />
  </main>
</template>

<style scoped>
/* Puedes eliminar completamente el bloque CSS anterior */
</style>
