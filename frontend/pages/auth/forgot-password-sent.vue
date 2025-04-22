&lt;template>
  &lt;NuxtLayout name="auth">
    &lt;template #title>
      Email Enviado
    &lt;/template>
    
    &lt;div class="space-y-6">
      &lt;div class="flex justify-center">
        &lt;div class="rounded-full bg-success-100 p-3">
          &lt;svg class="h-6 w-6 text-success-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            &lt;path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          &lt;/svg>
        &lt;/div>
      &lt;/div>

      &lt;div class="text-center">
        &lt;h3 class="text-lg font-medium text-gray-900">Instruções enviadas&lt;/h3>
        &lt;p class="mt-2 text-sm text-gray-600">
          Enviamos um email com instruções para redefinir sua senha.
          Por favor, verifique sua caixa de entrada e siga as instruções.
        &lt;/p>
      &lt;/div>

      &lt;div class="mt-6">
        &lt;NuxtLink
          to="/auth/login"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          Voltar para login
        &lt;/NuxtLink>
      &lt;/div>
    &lt;/div>
  &lt;/NuxtLayout>
&lt;/template>

&lt;script setup lang="ts">
// Define o título da página
useHead({
  title: 'Email Enviado - CRM Freelancer'
})
&lt;/script>
