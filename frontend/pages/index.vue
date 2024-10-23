<template>
  <UHorizontalNavigation :links="links">
    <template #default="{ link }">
      <span class="group-hover:text-primary relative">{{ link.label }}</span>
    </template>
  </UHorizontalNavigation>
  <div style="display: flex; align-items: center; justify-content: center; gap: 5px; flex-wrap: wrap; margin-bottom: 5vh; margin-top: 5vh;">
    <div style="width: 100%;text-align: center;">
      <UChip color="red" :text="`${incidents.length}`" size="2xl">
        <h3 class="text-5xl mb-1">{{ selectedYear }}</h3>
      </UChip>

      <div style="width:100%;display: flex;align-items: center; justify-content: center;margin-top: 20px; margin-bottom: 20px;gap:5px">
      <UInput
        style="width:100%"
        icon="i-heroicons-magnifying-glass-20-solid"
        size="sm"
        color="white"
        v-model="search"
        placeholder="Arama yap..."
      />

      <UButton
      @click="searchIncident()"
      >
        Ara
    </UButton>
    </div>
    </div>
    <UButton
      v-for="year in years"
      :key="year"
      inline
      @click="listIncidents(year)"
    >
      {{ year }}
    </UButton>
  </div>
  <div style="display: flex; align-items: center; gap: 5px; flex-wrap: wrap; width: max(100% - 40px); margin-inline: auto;margin-top: 5vh;">
    <UButton
      v-for="incident in incidents"
      :key="incident.id"
      inline
      :to="`/incident/${incident.id}`"
      color="blue"
    >
      {{ incident.name }}
    </UButton>
  </div>
</template>
<script setup>
import { ref } from 'vue';

const links = [{
  label: 'Açıklama',
  to: '/explanation',
}]

const startYear = 2008;
const currentYear = new Date().getFullYear();
const years = ref([]);
const search = ref([]);
const selectedYear = ref([]);

for (let year = startYear; year <= currentYear; year++) {
  years.value.push(year);
}

const incidents = ref([]);

async function listIncidents(year) {
  const response = await $fetch(`/api/list/incidents?year=${year}`);
  incidents.value = response.data;
  selectedYear.value = year;
}

async function searchIncident() {
  const response = await $fetch(`/api/search/incidents?name=${search.value}`);
  incidents.value = response.data;
}

await listIncidents(startYear);
</script>
