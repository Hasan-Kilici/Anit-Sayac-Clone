<template>
  <div>
    <NuxtLoadingIndicator :throttle="0" :duration="200" />
    <NuxtRouteAnnouncer/>
    <NuxtLayout>
    <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<style>
.page-enter-active,
.page-leave-active {
    transition: all 0.15s cubic-bezier(0, 0, 0.58, 1);
}

.page-enter-from,
.page-leave-to {
    opacity: 0;
    scale: 0.95;
    transform: translateY(10px);
}
</style>