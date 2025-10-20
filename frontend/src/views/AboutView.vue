<template>
  <section class="flex flex-col md:flex-row items-center justify-center min-h-screen bg-gradient-to-br from-gray-50 to-emerald-50 text-center p-8">
    <div class="max-w-xl">
      <h1 class="text-4xl font-bold text-emerald-700 mb-4">
        About This Project
      </h1>
      <p class="text-lg text-gray-700 leading-relaxed">
        This project was built with <span class="font-semibold text-emerald-600">Vue 3</span>,
        <span class="font-semibold text-blue-600">Vite</span>, and
        <span class="font-semibold text-cyan-600">TailwindCSS</span>.
        It’s designed to be clean, fast, and ready for scaling.
      </p>

      <div class="mt-8">
        <RouterLink
          to="/"
          class="inline-block px-5 py-2 bg-emerald-600 text-white rounded-lg shadow hover:bg-emerald-700 transition"
        >
          Back Home
        </RouterLink>
      </div>
    </div>
  </section>
</template>
