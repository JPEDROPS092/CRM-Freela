<template>
  <div>
    <div class="flex flex-col items-center justify-center min-h-screen px-4 py-12 sm:px-6 lg:px-8">
      <div class="w-full max-w-md space-y-8">
        <div class="text-center">
          <h1 class="text-3xl font-extrabold text-gray-900">CRM Freelancer</h1>
          <p class="mt-2 text-sm text-gray-600">
            Gerencie seus clientes, projetos e pagamentos em um só lugar
          </p>
        </div>
        
        <div class="mt-8 space-y-4">
          <NuxtLink 
            to="/auth/login" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
          >
            Entrar
          </NuxtLink>
          
          <NuxtLink 
            to="/auth/register" 
            class="w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
          >
            Criar conta
          </NuxtLink>
        </div>
        
        <div class="mt-8">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white text-gray-500">Ou veja como funciona</span>
            </div>
          </div>
          
          <div class="mt-6">
            <button 
              class="w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
              @click="showDemo = true"
            >
              Ver demonstração
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Modal de demonstração -->
    <Modal v-if="showDemo" @close="showDemo = false" title="Demonstração">
      <div class="space-y-4">
        <p>Veja como o CRM Freelancer pode ajudar você a gerenciar seus projetos:</p>
        
        <ul class="list-disc pl-5 space-y-2">
          <li>Gerencie seus clientes em um só lugar</li>
          <li>Acompanhe o status de seus projetos</li>
          <li>Controle seus pagamentos</li>
          <li>Visualize relatórios e estatísticas</li>
        </ul>
        
        <div class="mt-4">
          <NuxtLink 
            to="/auth/register" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
          >
            Começar agora
          </NuxtLink>
        </div>
      </div>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Modal from '~/components/ui/Modal.vue'

// Define o título da página
useHead({
  title: 'CRM Freelancer - Gerencie seus projetos'
})

const showDemo = ref(false)
</script>
