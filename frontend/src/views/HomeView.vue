<script setup>
import { ref, onMounted } from 'vue'
import FeatureCard from '@/components/FeatureCard.vue'
import { useCounterStore } from '@/stores/counter'
import apiService, { API_URL } from '@/services/api.service'

const counter = useCounterStore()
const people = ref([])
const planets = ref([])
const isLoading = ref(false)
const error = ref(null)

onMounted(async () => {
  try {
    isLoading.value = true
    const [peopleData, planetsData] = await Promise.all([
      apiService.getPeople(),
      apiService.getPlanets()
    ])
    
    // Create map of planet URLs to names
    const planetMap = {}
    for (const planet of planetsData) {
      planetMap[`${API_URL}/planets/${planet.id}`] = planet.name
    }

    // Fetch additional planet data for people's homeworlds not in initial planets list
    for (const person of peopleData) {
      if (!planetMap[person.homeworld]) {
        const planetId = person.homeworld.split('/').pop()
        const planet = await apiService.getPlanet(planetId)
        planetMap[person.homeworld] = planet.name
      }
    }

    people.value = peopleData.map(person => ({
      ...person,
      homeworldName: planetMap[person.homeworld] || 'Unknown'
    }))
    planets.value = planetsData
  } catch (err) {
    error.value = err.message
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <h1 class="text-4xl font-bold mb-6">Welcome to the Home</h1>
  <p class="text-lg mb-6">This is a responsive template app with dark/light theme support built with Tailwind CSS, Pinia store and Vue router.</p>
  
  <div class="mb-6">
    <button 
      @click="counter.increment()"
      class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 cursor-pointer"
    >
      Increment
    </button>
    <p class="mt-2">Count: {{ counter.count }}</p>
  </div>

  <div v-if="isLoading" class="text-center py-8">
    <p>Loading data...</p>
  </div>

  <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-6">
    Error loading data: {{ error }}
  </div>

  <div v-else>
    <h2 class="text-2xl font-bold mb-4">People</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
      <FeatureCard 
        v-for="person in people"
        :key="person.id"
        :title="person.name"
        :content="`From ${person.homeworldName}`"
      />
    </div>

    <h2 class="text-2xl font-bold mb-4">Planets</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <FeatureCard 
        v-for="planet in planets"
        :key="planet.id"
        :title="planet.name"
        :content="`Population: ${planet.population}`"
      />
    </div>
  </div>
</template>
