<template>
  <form @submit.prevent="handleSubmit" class="space-y-6">
    <div>
      <label for="email" class="form-label">Email</label>
      <div class="mt-1">
        <input 
          id="email" 
          v-model="email" 
          name="email" 
          type="email" 
          autocomplete="email" 
          required 
          class="form-input" 
          :class="{ 'border-danger-500 focus:ring-danger-500 focus:border-danger-500': errors.email }"
        />
        <p v-if="errors.email" class="form-error">{{ errors.email }}</p>
      </div>
    </div>

    <div>
      <label for="password" class="form-label">Senha</label>
      <div class="mt-1">
        <input 
          id="password" 
          v-model="password" 
          name="password" 
          type="password" 
          autocomplete="current-password" 
          required 
          class="form-input"
          :class="{ 'border-danger-500 focus:ring-danger-500 focus:border-danger-500': errors.password }"
        />
        <p v-if="errors.password" class="form-error">{{ errors.password }}</p>
      </div>
    </div>

    <div class="flex items-center justify-between">
      <div class="flex items-center">
        <input 
          id="remember-me" 
          v-model="rememberMe" 
          name="remember-me" 
          type="checkbox" 
          class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded" 
        />
        <label for="remember-me" class="ml-2 block text-sm text-gray-900">
          Lembrar-me
        </label>
      </div>

      <div class="text-sm">
        <NuxtLink to="/auth/forgot-password" class="font-medium text-primary-600 hover:text-primary-500">
          Esqueceu sua senha?
        </NuxtLink>
      </div>
    </div>

    <div>
      <button 
        type="submit" 
        class="btn-primary w-full"
        :disabled="isLoading"
      >
        <span v-if="isLoading" class="mr-2">
          <!-- Spinner simples -->
          <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
        </span>
        {{ isLoading ? 'Entrando...' : 'Entrar' }}
      </button>
    </div>

    <div v-if="authError" class="bg-danger-50 border border-danger-200 text-danger-700 px-4 py-3 rounded relative" role="alert">
      <strong class="font-bold">Erro!</strong>
      <span class="block sm:inline"> {{ authError }}</span>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '~/store/auth'

const authStore = useAuthStore()

// Estado do formulário
const email = ref('')
const password = ref('')
const rememberMe = ref(false)
const errors = ref({
  email: '',
  password: ''
})

// Obtém estado do store
const isLoading = computed(() => authStore.isLoading)
const authError = computed(() => authStore.getError)

// Validação do formulário
const validateForm = () => {
  let isValid = true
  errors.value = {
    email: '',
    password: ''
  }

  // Validação de email
  if (!email.value) {
    errors.value.email = 'O email é obrigatório'
    isValid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    errors.value.email = 'Por favor, insira um email válido'
    isValid = false
  }

  // Validação de senha
  if (!password.value) {
    errors.value.password = 'A senha é obrigatória'
    isValid = false
  } else if (password.value.length < 6) {
    errors.value.password = 'A senha deve ter pelo menos 6 caracteres'
    isValid = false
  }

  return isValid
}

// Manipulador de envio do formulário
const handleSubmit = async () => {
  if (!validateForm()) return

  const success = await authStore.login(email.value, password.value)
  
  if (success) {
    // Redireciona para o dashboard após login bem-sucedido
    navigateTo('/dashboard')
  }
}
</script>
