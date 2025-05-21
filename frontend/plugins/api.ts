import { defineNuxtPlugin, useRuntimeConfig } from 'nuxt/app'
import { useAuthStore } from '~/store/auth'
import { useRouter } from 'nuxt/app'

export default defineNuxtPlugin((nuxtApp: any) => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()

  // Interceptor para adicionar token de autenticação
  // Verifica se estamos no navegador antes de tentar acessar window
  if (process.client) {
    nuxtApp.hook('app:created', () => {
      const originalFetch = window.fetch
      window.fetch = async (input: RequestInfo | URL, init?: RequestInit) => {
        const url = input instanceof Request ? input.url : input.toString()
        
        // Não intercepta chamadas para URLs externas
        if (typeof config.public.apiBase === 'string' && !url.startsWith(config.public.apiBase)) {
          return originalFetch(input, init)
        }

        // Configura headers padrão
        const headers = new Headers(init?.headers || {})
        headers.set('Content-Type', 'application/json')

        // Adiciona token de autenticação se disponível
        if (authStore.accessToken) {
          headers.set('Authorization', `Bearer ${authStore.accessToken}`)
        }

        // Configurações adicionais para CORS
        const fetchOptions: RequestInit = {
          ...init,
          headers,
          credentials: 'include', // Importante para cookies HttpOnly
          mode: 'cors' // Explicitamente define o modo CORS
        }

        // Faz a requisição
        try {
          const response = await originalFetch(input, fetchOptions)

          // Trata erros de autenticação
          if (response.status === 401) {
            try {
              // Tenta renovar o token
              await authStore.refreshAccessToken()

              // Refaz a requisição original com o novo token
              headers.set('Authorization', `Bearer ${authStore.accessToken}`)
              return await originalFetch(input, {
                ...fetchOptions,
                headers
              })
            } catch (error) {
              // Se não conseguir renovar o token, faz logout
              authStore.logout()
              // Usar navegação segura para SSR
              const router = useRouter()
              router.push('/auth/login')
              return response
            }
          }

          return response
        } catch (error) {
          // Trata erros gerais
          console.error('Erro ao fazer requisição:', error)
          throw error
        }
      }
    })
  }
})
