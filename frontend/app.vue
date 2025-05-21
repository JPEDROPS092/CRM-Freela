<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
      <ClientOnly>
        <CommonNotificationsContainer />
      </ClientOnly>
    </NuxtLayout>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from '#imports'
import { useHead } from '#imports'
import { useAuthStore } from '~/store/auth'
import CommonNotificationsContainer from '~/components/common/NotificationsContainer.vue'

// Inicializa o estado de autenticação
const authStore = useAuthStore()

onMounted(() => {
  // Inicializa a autenticação ao carregar a aplicação
  authStore.initAuth()
})

// Configuração global do aplicativo
useHead({
  title: 'CRM Freelancer',
  meta: [
    { name: 'description', content: 'Sistema CRM para freelancers gerenciarem clientes, tarefas e pagamentos' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { charset: 'utf-8' },
    // Cabeçalhos de segurança
    { 'http-equiv': 'X-Content-Type-Options', content: 'nosniff' },
    { 'http-equiv': 'X-Frame-Options', content: 'DENY' }
  ],
  link: [
    { rel: 'icon', type: 'image/png', href: '/favicon.png' }
  ]
})
</script>

<style>
/* Estilos globais adicionais podem ser adicionados aqui */
:root {
  --color-primary: #4f46e5;
  --color-primary-dark: #4338ca;
  --color-primary-light: #6366f1;
  --color-danger: #ef4444;
  --color-success: #10b981;
  --color-warning: #f59e0b;
  --color-info: #3b82f6;
}

/* Classes utilitárias consistentes */
.btn-primary {
  @apply py-2 px-4 bg-primary-600 text-white rounded-md shadow-sm hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed;
}

.btn-secondary {
  @apply py-2 px-4 bg-white text-gray-700 border border-gray-300 rounded-md shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed;
}

.btn-danger {
  @apply py-2 px-4 bg-red-600 text-white rounded-md shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed;
}

.form-label {
  @apply block text-sm font-medium text-gray-700;
}

.form-input {
  @apply block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm;
}

.form-error {
  @apply mt-1 text-sm text-red-600;
}

/* Otimização para dispositivos móveis */
@media (max-width: 640px) {
  .container {
    @apply px-4;
  }
}
</style>
