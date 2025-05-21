<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Barra de navegação superior -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <!-- Logo -->
            <div class="flex-shrink-0 flex items-center">
              <NuxtLink to="/" class="text-primary-600 font-bold text-xl">CRM Freelancer</NuxtLink>
            </div>
            
            <!-- Links de navegação principal -->
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <NuxtLink to="/dashboard" class="border-transparent text-gray-500 hover:border-primary-500 hover:text-primary-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                Dashboard
              </NuxtLink>
              <NuxtLink to="/clients" class="border-transparent text-gray-500 hover:border-primary-500 hover:text-primary-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                Clientes
              </NuxtLink>
              <NuxtLink to="/tasks" class="border-transparent text-gray-500 hover:border-primary-500 hover:text-primary-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                Tarefas
              </NuxtLink>
              <NuxtLink to="/payments" class="border-transparent text-gray-500 hover:border-primary-500 hover:text-primary-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                Pagamentos
              </NuxtLink>
            </div>
          </div>
          
          <!-- Menu do usuário -->
          <div class="hidden sm:ml-6 sm:flex sm:items-center">
            <div class="ml-3 relative" ref="userMenuRef">
              <div>
                <button type="button" class="bg-white rounded-full flex text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500" id="user-menu-button" aria-expanded="false" aria-haspopup="true" @click="isUserMenuOpen = !isUserMenuOpen">
                  <span class="sr-only">Abrir menu do usuário</span>
                  <div v-if="user && user.avatar" class="h-8 w-8 rounded-full overflow-hidden">
                    <img :src="user.avatar" alt="Avatar do usuário" class="h-full w-full object-cover">
                  </div>
                  <div v-else class="h-8 w-8 rounded-full bg-primary-600 flex items-center justify-center text-white">
                    {{ user && user.name ? user.name.charAt(0).toUpperCase() : 'U' }}
                  </div>
                </button>
              </div>
              
              <!-- Menu dropdown do usuário -->
              <div v-if="isUserMenuOpen" class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                <div class="px-4 py-2 border-b border-gray-100">
                  <p class="text-sm font-medium text-gray-700">{{ user && user.name ? user.name : 'Usuário' }}</p>
                  <p class="text-xs text-gray-500 truncate">{{ user && user.email ? user.email : 'usuario@exemplo.com' }}</p>
                </div>
                <NuxtLink to="/profile" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1">Seu Perfil</NuxtLink>
                <NuxtLink to="/settings" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1">Configurações</NuxtLink>
                <button @click="logout" class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1">Sair</button>
              </div>
            </div>
          </div>
          
          <!-- Botão do menu móvel -->
          <div class="-mr-2 flex items-center sm:hidden">
            <button type="button" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500" aria-controls="mobile-menu" aria-expanded="false" @click="isMobileMenuOpen = !isMobileMenuOpen">
              <span class="sr-only">Abrir menu principal</span>
              <!-- Ícone de menu (pode ser substituído por um componente de ícone) -->
              <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>
      
      <!-- Menu móvel -->
      <div v-if="isMobileMenuOpen" class="sm:hidden" ref="mobileMenuRef" id="mobile-menu">
        <div class="pt-2 pb-3 space-y-1">
          <NuxtLink to="/dashboard" class="bg-primary-50 border-primary-500 text-primary-700 block pl-3 pr-4 py-2 border-l-4 text-base font-medium">Dashboard</NuxtLink>
          <NuxtLink to="/clients" class="border-transparent text-gray-500 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-700 block pl-3 pr-4 py-2 border-l-4 text-base font-medium">Clientes</NuxtLink>
          <NuxtLink to="/tasks" class="border-transparent text-gray-500 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-700 block pl-3 pr-4 py-2 border-l-4 text-base font-medium">Tarefas</NuxtLink>
          <NuxtLink to="/payments" class="border-transparent text-gray-500 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-700 block pl-3 pr-4 py-2 border-l-4 text-base font-medium">Pagamentos</NuxtLink>
        </div>
        <div class="pt-4 pb-3 border-t border-gray-200">
          <div class="flex items-center px-4">
            <div class="flex-shrink-0">
              <div v-if="user && user.avatar" class="h-10 w-10 rounded-full overflow-hidden">
                <img :src="user.avatar" alt="Avatar do usuário" class="h-full w-full object-cover">
              </div>
              <div v-else class="h-10 w-10 rounded-full bg-primary-600 flex items-center justify-center text-white">
                {{ user && user.name ? user.name.charAt(0).toUpperCase() : 'U' }}
              </div>
            </div>
            <div class="ml-3">
              <div class="text-base font-medium text-gray-800">{{ user && user.name ? user.name : 'Usuário' }}</div>
              <div class="text-sm font-medium text-gray-500">{{ user && user.email ? user.email : 'usuario@exemplo.com' }}</div>
            </div>
          </div>
          <div class="mt-3 space-y-1">
            <NuxtLink to="/profile" class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100">Seu Perfil</NuxtLink>
            <NuxtLink to="/settings" class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100">Configurações</NuxtLink>
            <button @click="logout" class="block w-full text-left px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100">Sair</button>
          </div>
        </div>
      </div>
    </nav>
    
    <!-- Conteúdo principal -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <slot />
    </main>
    
    <!-- Rodapé -->
    <footer class="bg-white border-t border-gray-200 mt-auto">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <p class="text-center text-sm text-gray-500">
          &copy; {{ new Date().getFullYear() }} CRM Freelancer. Todos os direitos reservados.
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { useAuthStore } from '~/store/auth'
import { navigateTo } from '#app'

const authStore = useAuthStore()
const isAuthenticated = computed(() => authStore.isAuthenticated)
const user = computed(() => authStore.user)

// Estado para controlar a exibição dos menus
const isMobileMenuOpen = ref(false)
const isUserMenuOpen = ref(false)
const userMenuRef = ref(null)

// Alternar a exibição do menu móvel
const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

// Alternar a exibição do menu do usuário
const toggleUserMenu = () => {
  isUserMenuOpen.value = !isUserMenuOpen.value
}

// Função de logout
const logout = async () => {
  await authStore.logout()
  // Redirecionar para a página de login
  navigateTo('/auth/login')
}

// Implementação manual para substituir onClickOutside
const handleClickOutside = (event) => {
  // Verificar se o clique foi fora do menu do usuário
  if (userMenuRef.value && !userMenuRef.value.contains(event.target)) {
    isUserMenuOpen.value = false
  }
  
  // Verificar se o clique foi fora do menu mobile
  const mobileMenu = document.getElementById('mobile-menu')
  if (mobileMenu && !mobileMenu.contains(event.target)) {
    isMobileMenuOpen.value = false
  }
}

// Adicionar e remover o event listener
onMounted(() => {
  if (process.client) {
    document.addEventListener('click', handleClickOutside)
  }
})

onUnmounted(() => {
  if (process.client) {
    document.removeEventListener('click', handleClickOutside)
  }
})
</script>
