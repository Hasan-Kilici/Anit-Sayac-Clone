<script setup>
    const route = useRoute()
    const response = await $fetch(`/api/find/incident?id=${route.params.id}`)
    const incident = ref(response.data)
</script>
<template>
<ul>
    <li>{{ incident.name }}</li>
    <li>{{ incident.date }}</li>
    <li>{{ incident.reason }}</li>
    <li>{{ incident.by }}</li>
    <li>{{ incident.method }}</li>
    <li>{{ incident.protection }}</li>
</ul>
</template>