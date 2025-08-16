import axios from 'axios'

export const API_URL = import.meta.env.VITE_API_URL

export default {
  async getPeople() {
    const response = await axios.get(`${API_URL}/people`)
    return response.data.results
  },

  async getPerson(id) {
    const response = await axios.get(`${API_URL}/people/${id}`)
    return response.data
  },

  async getPlanets() {
    const response = await axios.get(`${API_URL}/planets`)
    return response.data.results
  },

  async getPlanet(id) {
    const response = await axios.get(`${API_URL}/planets/${id}`)
    return response.data
  }
}
