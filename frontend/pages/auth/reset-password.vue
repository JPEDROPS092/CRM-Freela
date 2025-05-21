<template>
  <NuxtLayout name="auth">
    <template #title>
      Redefinir Senha
    </template>
    
    <div class="space-y-6">
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div>
          <label for="password" class="form-label">Nova Senha</label>
          <div class="mt-1">
            <input 
              id="password" 
              v-model="password" 
              name="password" 
              type="password" 
              autocomplete="new-password" 
              required 
              class="form-input" 
              :class="{ 'border-danger-500 focus:ring-danger-500 focus:border-danger-500': errors.password }"
            />
            <p v-if="errors.password" class="form-error">{{ errors.password }}</p>
          </div>
        </div>

        <div>
          <label for="confirmPassword" class="form-label">Confirmar Nova Senha</label>
          <div class="mt-1">
            <input 
              id="confirmPassword" 
              v-model="confirmPassword" 
              name="confirmPassword" 
              type="password" 
              autocomplete="new-password" 
              required 
              class="form-input"
              :class="{ 'border-danger-500 focus:ring-danger-500 focus:border-danger-500': errors.confirmPassword }"
            />
            <p v-if="errors.confirmPassword" class="form-error">{{ errors.confirmPassword }}</p>
          </div>
        </div>

        <div>
          <button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            :disabled="loading"
          >
            {{ loading ? 'Redefinindo...' : 'Redefinir Senha' }}
          </button>
        </div>
      </form>
    </div>
  </NuxtLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '~/store/auth'
import { useRoute, useRouter } from 'vue-router'

// Define o título da página
useHead({
  title: 'Redefinir Senha - CRM Freelancer'
})

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const password = ref('')
const confirmPassword = ref('')
const errors = ref({
  password: '',
  confirmPassword: ''
})
const loading = ref(false)

const validateForm = () => {
  const newErrors = {
    password: '',
    confirmPassword: ''
  }

  if (password.value.length < 6) {
    newErrors.password = 'A senha deve ter pelo menos 6 caracteres'
  }

  if (password.value !== confirmPassword.value) {
    newErrors.confirmPassword = 'As senhas não coincidem'
  }

  errors.value = newErrors
  return !newErrors.password && !newErrors.confirmPassword
}

const handleSubmit = async () => {
  if (!validateForm()) return

  try {
    loading.value = true
    const token = route.query.token as string

    if (!token) {
      throw new Error('Token de redefinição não encontrado')
    }

    await authStore.resetPassword(token, password.value)
    
    // Redireciona para login com mensagem de sucesso
    router.push({
      path: '/auth/login',
      query: { message: 'Senha redefinida com sucesso' }
    })
  } catch (err: any) {
    errors.value.password = err.message || 'Ocorreu um erro ao redefinir sua senha'
  } finally {
    loading.value = false
  }
}
</script>
