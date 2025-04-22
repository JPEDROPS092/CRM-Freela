import { defineStore } from 'pinia'

interface User {
  id: number
  name: string
  email: string
  plan: string
}

interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
  isAuthenticated: boolean
  loading: boolean
  error: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    accessToken: null,
    refreshToken: null,
    isAuthenticated: false,
    loading: false,
    error: null
  }),

  getters: {
    getUser: (state) => state.user,
    getUserPlan: (state) => state.user?.plan || 'free',
    isLoggedIn: (state) => state.isAuthenticated,
    isLoading: (state) => state.loading,
    getError: (state) => state.error
  },

  actions: {
    async login(email: string, password: string) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/auth/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ email, password })
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha na autenticação')
        }
        
        this.accessToken = data.tokens.access_token
        this.refreshToken = data.tokens.refresh_token
        this.isAuthenticated = true
        
        // Armazena tokens no localStorage
        localStorage.setItem('accessToken', this.accessToken as string)
        localStorage.setItem('refreshToken', this.refreshToken as string)
        
        // Busca os dados do usuário
        await this.fetchUserProfile()
        
        return true
      } catch (error: any) {
        this.error = error.message || 'Erro ao fazer login'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async register(name: string, email: string, password: string) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/auth/register`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ name, email, password })
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha no registro')
        }
        
        this.user = data.user
        this.accessToken = data.tokens.access_token
        this.refreshToken = data.tokens.refresh_token
        this.isAuthenticated = true
        
        // Armazena tokens no localStorage
        localStorage.setItem('accessToken', this.accessToken as string)
        localStorage.setItem('refreshToken', this.refreshToken as string)
        
        return true
      } catch (error: any) {
        this.error = error.message || 'Erro ao registrar'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async fetchUserProfile() {
      if (!this.accessToken) return false
      
      this.loading = true
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/user/profile`, {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${this.accessToken}`
          }
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha ao buscar perfil')
        }
        
        this.user = data.user
        return true
      } catch (error: any) {
        this.error = error.message || 'Erro ao buscar perfil'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async refreshAccessToken() {
      if (!this.refreshToken) return false
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/auth/refresh`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ refresh_token: this.refreshToken })
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha ao renovar token')
        }
        
        this.accessToken = data.tokens.access_token
        localStorage.setItem('accessToken', this.accessToken as string)
        
        return true
      } catch (error) {
        // Se falhar ao renovar o token, faz logout
        this.logout()
        return false
      }
    },
    
    logout() {
      this.user = null
      this.accessToken = null
      this.refreshToken = null
      this.isAuthenticated = false
      
      // Remove tokens do localStorage
      localStorage.removeItem('accessToken')
      localStorage.removeItem('refreshToken')
      
      // Redireciona para a página de login
      navigateTo('/auth/login')
    },
    
    // Inicializa o estado de autenticação a partir do localStorage
    initAuth() {
      const accessToken = localStorage.getItem('accessToken')
      const refreshToken = localStorage.getItem('refreshToken')
      
      if (accessToken && refreshToken) {
        this.accessToken = accessToken
        this.refreshToken = refreshToken
        this.isAuthenticated = true
        
        // Busca os dados do usuário
        this.fetchUserProfile()
      }
    }
  }
})
