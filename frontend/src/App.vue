<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'
import logo from '@/assets/logo.svg'

const showMobileMenu = ref(false)
</script>

<template>
  <!-- Navigation -->
  <nav class="bg-white dark:bg-gray-800 shadow-md">
    <div class="max-w-6xl mx-auto px-4">
        <div class="flex justify-between items-center h-16">
            <div class="flex items-center space-x-4">
                <RouterLink to="/" class="flex items-center text-xl font-bold text-primary-600 dark:text-primary-400">
                    <img :src="logo" class="h-8 mr-2" />
                    Template App
                </RouterLink>
            </div>
            <div class="hidden md:flex items-center space-x-8">
                <RouterLink to="/" class="hover:text-primary-600 dark:hover:text-primary-400">Home</RouterLink>
                <RouterLink to="/about" class="hover:text-primary-600 dark:hover:text-primary-400">About</RouterLink>
            </div>

            <div class="md:hidden flex items-center space-x-4">                
                <button class="md:hidden p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer" @click="showMobileMenu = !showMobileMenu">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                    </svg>
                </button>
            </div>  
        </div>
    </div>
  </nav>

    <!-- Mobile Menu (hidden by default) -->
    <div id="mobile-menu" class="md:hidden bg-white dark:bg-gray-800 shadow-lg" v-show="showMobileMenu">
        <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                    <RouterLink to="/" class="block px-3 py-2 rounded-md text-base font-medium hover:text-primary-600 dark:hover:text-primary-400">Home</RouterLink>
                    <RouterLink to="/about" class="block px-3 py-2 rounded-md text-base font-medium hover:text-primary-600 dark:hover:text-primary-400">About</RouterLink>
        </div>
    </div>

  <!-- Main Content -->
  <main class="max-w-6xl mx-auto px-4 py-8">
    <section class="mb-12">  
      <RouterView />
    </section>
  </main>
</template>

<style scoped>

</style>
