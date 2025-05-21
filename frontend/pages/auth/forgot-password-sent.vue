<template>
  <NuxtLayout name="auth">
    <template #title>
      Email Enviado
    </template>
    
    <div class="space-y-6">
      <div class="flex justify-center">
        <div class="rounded-full bg-success-100 p-3">
          <svg class="h-6 w-6 text-success-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
      </div>

      <div class="text-center">
        <h3 class="text-lg font-medium text-gray-900">Instruções enviadas</h3>
        <p class="mt-2 text-sm text-gray-600">
          Enviamos um email com instruções para redefinir sua senha.
          Por favor, verifique sua caixa de entrada e siga as instruções.
        </p>
      </div>

      <div class="mt-6">
        <NuxtLink
          to="/auth/login"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          Voltar para login
        </NuxtLink>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup lang="ts">
// Define o título da página
useHead({
  title: 'Email Enviado - CRM Freelancer'
})
</script>
