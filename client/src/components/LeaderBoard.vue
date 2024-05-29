<script setup lang="ts">
import {ref, onMounted, watch} from 'vue';
import axios from 'axios';

interface Player {
  place: number;
  nickname: string;
  score: number;
}
const baseUrl = 'http://77.221.139.206:3000/api';
const players = ref<Player[]>([]);
const error = ref<string | null>(null);
const loading = ref(true);
const currentNickname = ref(localStorage.getItem('currentNickname') || '');

const fetchPlayers = async () => {
  try {
    const response = await axios.get(`${baseUrl}/players`);
    players.value = response.data;
  } catch (err) {
    error.value = 'Failed to fetch leaderboard data';
  } finally {
    loading.value = false;
  }
};

const addPlayer = async () => {
  try {
    await axios.post(`${baseUrl}/players`, {
      nickname: currentNickname.value,
      score: localStorage.getItem('opencraft/resources') ? JSON.parse(localStorage.getItem('opencraft/resources') || '[]').length : 0
    });
    await fetchPlayers();
  } catch (err) {
    console.error('Failed to add player:', err);
  }
};

watch(currentNickname, (newVal) => {
  localStorage.setItem('currentNickname', newVal);
});


onMounted(() => {
  fetchPlayers();
});
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="mt-4">
      <div class="flex flex-row w-full">
        <input v-model="currentNickname" type="text" class="rounded-lg border-gray-300 px-3 py-2 w-full focus:outline-none" placeholder="Введите никнейм...">
        <button @click="addPlayer" class="bg-[#6050bc] text-white px-2 py-2 rounded-lg flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-1" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M13.293 5.293a1 1 0 0 1 1.414 1.414l-6 6a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L8 10.586l5.293-5.293a1 1 0 0 1 1.414 0z" clip-rule="evenodd" />
          </svg>
        </button>
      </div>
    </div>
    <div class="mt-4 flex-grow">
      <div v-if="loading" class="text-center">Loading...</div>
      <div v-else-if="error" class="text-center text-red-500">{{ error }}</div>
      <div v-else>
        <div v-for="player in players" :key="player.place" class="mt-4 mb-4">
          <div :class="{ 'bg-[#6050bc]': player.nickname === currentNickname }"  class="bg-gray-200 rounded-lg p-3 flex items-center justify-between">
            <div :class="{ 'text-white': player.nickname === currentNickname }"  class="text-lg font-semibold">{{ player.nickname }}</div>
            <div :class="{ 'text-white': player.nickname === currentNickname }" class="text-lg">{{ player.score }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>
