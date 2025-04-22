<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
// Configuração global do aplicativo
useHead({
  title: 'CRM Freelancer',
  meta: [
    { name: 'description', content: 'Sistema CRM para freelancers gerenciarem clientes, tarefas e pagamentos' }
  ],
  link: [
    { rel: 'icon', type: 'image/png', href: '/favicon.png' }
  ]
})
</script>

<style>
/* Estilos globais adicionais podem ser adicionados aqui */
</style>
