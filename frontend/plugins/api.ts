import { useAuthStore } from '~/store/auth'

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()

  // Interceptor para adicionar token de autenticação
  nuxtApp.hook('app:created', () => {
    const originalFetch = window.fetch
    window.fetch = async (input: RequestInfo | URL, init?: RequestInit) => {
      const url = input instanceof Request ? input.url : input.toString()
      
      // Não intercepta chamadas para URLs externas
      if (!url.startsWith(config.public.apiBase)) {
        return originalFetch(input, init)
      }

      // Configura headers padrão
      const headers = new Headers(init?.headers || {})
      headers.set('Content-Type', 'application/json')

      // Adiciona token de autenticação se disponível
      if (authStore.accessToken) {
        headers.set('Authorization', `Bearer ${authStore.accessToken}`)
      }

      // Faz a requisição
      const response = await originalFetch(input, {
        ...init,
        headers
      })

      // Trata erros de autenticação
      if (response.status === 401) {
        try {
          // Tenta renovar o token
          await authStore.refreshAccessToken()

          // Refaz a requisição original com o novo token
          headers.set('Authorization', `Bearer ${authStore.accessToken}`)
          return await originalFetch(input, {
            ...init,
            headers
          })
        } catch (error) {
          // Se não conseguir renovar o token, faz logout
          authStore.logout()
          navigateTo('/auth/login')
          return response
        }
      }

      return response
    }
  })
})
