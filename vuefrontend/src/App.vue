<template>
  <main class="w-full h-full flex flex-col items-center justify-center">
    <RouterLink to="/" class="hover:text-gray-300">Home</RouterLink>
    <RouterLink to="/table" class="hover:text-gray-300">Table</RouterLink>
    <RouterView />
  </main>
  <footer class="flex-none w-full bg-gray-100 text-gray-800">
    <div class="mx-auto py-4 px-4 sm:px-6 lg:px-8">
      <p class="text-center text-sm text-gray-500 hover:text-gray-700">
        <a 
          href="https://janharkonen.fi" 
          target="_blank" 
          rel="noopener noreferrer"
          class="hover:bg-gray-200"
          >
          © Jan Härkönen 2025
        </a>
      </p>
    </div>
  </footer>
</template>