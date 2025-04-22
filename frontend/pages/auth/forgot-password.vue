&lt;template>
  &lt;NuxtLayout name="auth">
    &lt;template #title>
      Recuperar Senha
    &lt;/template>
    
    &lt;div class="space-y-6">
      &lt;p class="text-sm text-gray-600">
        Digite seu e-mail e enviaremos instruções para redefinir sua senha.
      &lt;/p>

      &lt;form @submit.prevent="handleSubmit" class="space-y-6">
        &lt;div>
          &lt;label for="email" class="form-label">Email&lt;/label>
          &lt;div class="mt-1">
            &lt;input 
              id="email" 
              v-model="email" 
              name="email" 
              type="email" 
              autocomplete="email" 
              required 
              class="form-input" 
              :class="{ 'border-danger-500 focus:ring-danger-500 focus:border-danger-500': error }"
            />
            &lt;p v-if="error" class="form-error">{{ error }}&lt;/p>
          &lt;/div>
        &lt;/div>

        &lt;div>
          &lt;button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            :disabled="loading"
          >
            {{ loading ? 'Enviando...' : 'Enviar instruções' }}
          &lt;/button>
        &lt;/div>
      &lt;/form>
    &lt;/div>
    
    &lt;template #footer>
      &lt;p class="text-sm text-gray-600">
        Lembrou sua senha?
        &lt;NuxtLink to="/auth/login" class="font-medium text-primary-600 hover:text-primary-500">
          Voltar para login
        &lt;/NuxtLink>
      &lt;/p>
    &lt;/template>
  &lt;/NuxtLayout>
&lt;/template>

&lt;script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '~/store/auth'

// Define o título da página
useHead({
  title: 'Recuperar Senha - CRM Freelancer'
})

const authStore = useAuthStore()
const email = ref('')
const error = ref('')
const loading = ref(false)

const handleSubmit = async () => {
  try {
    loading.value = true
    error.value = ''
    
    await authStore.forgotPassword(email.value)
    
    // Redireciona para uma página de confirmação ou mostra uma mensagem de sucesso
    navigateTo('/auth/forgot-password-sent')
  } catch (err: any) {
    error.value = err.message || 'Ocorreu um erro ao processar sua solicitação'
  } finally {
    loading.value = false
  }
}
&lt;/script>
