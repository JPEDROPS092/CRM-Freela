import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore()
  const publicPages = ['/auth/login', '/auth/register', '/auth/forgot-password', '/auth/reset-password']
  const authRequired = !publicPages.includes(to.path)

  // Verifica se o usuário está autenticado
  if (authRequired && !authStore.isAuthenticated) {
    return navigateTo({
      path: '/auth/login',
      query: { redirect: to.fullPath }
    })
  }

  // Redireciona usuário autenticado para dashboard se tentar acessar páginas públicas
  if (!authRequired && authStore.isAuthenticated) {
    return navigateTo('/')
  }
})
