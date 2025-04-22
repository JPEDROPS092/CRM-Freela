<template>
  <NuxtLayout name="auth">
    <template #title>
      Acesse sua conta
    </template>
    
    <LoginForm />
    
    <template #footer>
      <p class="text-sm text-gray-600">
        Não tem uma conta?
        <NuxtLink to="/auth/register" class="font-medium text-primary-600 hover:text-primary-500">
          Registre-se agora
        </NuxtLink>
      </p>
    </template>
  </NuxtLayout>
</template>

<script setup lang="ts">
import LoginForm from '~/components/auth/LoginForm.vue'

// Define o título da página
useHead({
  title: 'Login - CRM Freelancer'
})
</script>
