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
<script setup>
    const route = useRoute()

    const response = await $fetch(`/api/find/incident?id=${route.params.id}`)
    const incident = await ref(response.data)

    onMounted(()=>{
        setTimeout(async ()=>{
            incident.value = await response.data;
        },100)
    })
</script>