<template>
  <NuxtLayout name="auth">
    <template #title>
      Crie sua conta
    </template>
    
    <RegisterForm />
    
    <template #footer>
      <p class="text-sm text-gray-600">
        Já tem uma conta?
        <NuxtLink to="/auth/login" class="font-medium text-primary-600 hover:text-primary-500">
          Faça login
        </NuxtLink>
      </p>
    </template>
  </NuxtLayout>
</template>

<script setup lang="ts">
import RegisterForm from '~/components/auth/RegisterForm.vue'

// Define o título da página
useHead({
  title: 'Registro - CRM Freelancer'
})
</script>
