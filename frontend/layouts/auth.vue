<template>
  <div class="min-h-screen flex flex-col justify-center py-12 sm:px-6 lg:px-8 bg-gray-50">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        <NuxtLink to="/" class="text-primary-600 hover:text-primary-500">
          CRM Freelancer
        </NuxtLink>
      </h2>
      <h3 class="mt-2 text-center text-xl text-gray-600">
        <slot name="title">Bem-vindo</slot>
      </h3>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <slot />
      </div>
      
      <div class="mt-6 text-center">
        <slot name="footer" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Layout específico para páginas de autenticação
</script>
