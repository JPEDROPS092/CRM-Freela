import { useAuthStore } from '~/store/auth'
import { navigateTo } from '#app'

export default defineNuxtRouteMiddleware(async (to) => {
  const authStore = useAuthStore()
  const publicPages = ['/auth/login', '/auth/register', '/auth/forgot-password', '/auth/reset-password', '/auth/forgot-password-sent']
  const authRequired = !publicPages.includes(to.path)

  // Verifica se o token está expirado e tenta renovar se necessário
  if (authStore.isAuthenticated && authStore.isTokenExpired) {
    try {
      await authStore.refreshToken()
    } catch (error) {
      // Se não conseguir renovar o token, faz logout e redireciona para login
      authStore.logout()
      return navigateTo({
        path: '/auth/login',
        query: { redirect: to.fullPath, expired: 'true' }
      })
    }
  }

  // Verifica se o usuário está autenticado
  if (authRequired && !authStore.isAuthenticated) {
    return navigateTo({
      path: '/auth/login',
      query: { redirect: to.fullPath }
    })
  }

  // Redireciona usuário autenticado para dashboard se tentar acessar páginas públicas
  if (!authRequired && authStore.isAuthenticated) {
    return navigateTo('/dashboard')
  }
})
