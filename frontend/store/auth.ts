import { defineStore } from 'pinia'
import { jwtDecode } from 'jwt-decode'
import { useRuntimeConfig } from '#app'

interface User {
  id: number
  name: string
  email: string
  plan: string
}

interface JwtPayload {
  exp: number
  sub: string
  [key: string]: any
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
    getError: (state) => state.error,
    
    // Verifica se o token está expirado
    isTokenExpired: (state) => {
      if (!state.accessToken) return true
      
      try {
        const decoded = jwtDecode<JwtPayload>(state.accessToken)
        const currentTime = Date.now() / 1000
        
        // Considera expirado se faltar menos de 5 minutos para expirar
        return decoded.exp < currentTime + 300
      } catch (error) {
        console.error('Erro ao decodificar token:', error)
        return true
      }
    }
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
          body: JSON.stringify({ email, password }),
          credentials: 'include' // Importante para cookies HttpOnly
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha na autenticação')
        }
        
        // Ajustado para corresponder à estrutura da resposta do backend
        this.accessToken = data.token
        this.refreshToken = data.token // Usando o mesmo token como refresh por enquanto
        this.isAuthenticated = true
        
        // Armazena tokens no localStorage com segurança
        this.securelyStoreTokens(this.accessToken, this.refreshToken)
        
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
          body: JSON.stringify({ name, email, password }),
          credentials: 'include' // Importante para cookies HttpOnly
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha no registro')
        }
        
        // Após o registro bem-sucedido, faça login automaticamente
        return await this.login(email, password)
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
        
        // Ajustado para corresponder à estrutura da resposta do backend
        this.user = {
          id: data.id || data.user?.id,
          name: data.name || data.user?.name,
          email: data.email || data.user?.email,
          plan: data.plan || data.user?.plan || 'free'
        }
        
        return true
      } catch (error: any) {
        this.error = error.message || 'Erro ao buscar perfil'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async refreshToken() {
      if (!this.refreshToken) return false
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/auth/refresh`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ refresh_token: this.refreshToken }),
          credentials: 'include' // Importante para cookies HttpOnly
        })
        
        const data = await response.json()
        
        if (!response.ok) {
          throw new Error(data.error || 'Falha ao renovar token')
        }
        
        this.accessToken = data.tokens.access_token
        
        // Se o refresh token também for rotacionado
        if (data.tokens.refresh_token) {
          this.refreshToken = data.tokens.refresh_token
        }
        
        // Atualiza os tokens armazenados
        this.securelyStoreTokens(this.accessToken, this.refreshToken)
        
        return true
      } catch (error) {
        // Se falhar ao renovar o token, faz logout
        this.logout()
        throw new Error('Sessão expirada. Por favor, faça login novamente.')
      }
    },
    
    logout() {
      // Tenta fazer logout no servidor para invalidar o token
      this.serverLogout().catch(console.error)
      
      // Limpa o estado local independentemente da resposta do servidor
      this.clearAuthState()
    },
    
    async serverLogout() {
      if (!this.accessToken) return
      
      try {
        const config = useRuntimeConfig()
        await fetch(`${config.public.apiBase}/auth/logout`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${this.accessToken}`,
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ refresh_token: this.refreshToken }),
          credentials: 'include'
        })
      } catch (error) {
        console.error('Erro ao fazer logout no servidor:', error)
      }
    },
    
    clearAuthState() {
      this.user = null
      this.accessToken = null
      this.refreshToken = null
      this.isAuthenticated = false
      
      // Remove tokens do localStorage
      localStorage.removeItem('accessToken')
      localStorage.removeItem('refreshToken')
      localStorage.removeItem('tokenTimestamp')
    },
    
    // Armazena tokens com segurança
    securelyStoreTokens(accessToken: string, refreshToken: string) {
      localStorage.setItem('accessToken', accessToken)
      localStorage.setItem('refreshToken', refreshToken)
      
      // Define quando o token foi armazenado para verificar validade
      localStorage.setItem('tokenTimestamp', Date.now().toString())
    },
    
    // Inicializa o estado de autenticação a partir do localStorage
    initAuth() {
      const accessToken = localStorage.getItem('accessToken')
      const refreshToken = localStorage.getItem('refreshToken')
      const tokenTimestamp = localStorage.getItem('tokenTimestamp')
      
      // Verifica se os tokens existem e não estão muito antigos (7 dias)
      const isTokenValid = tokenTimestamp && (Date.now() - parseInt(tokenTimestamp)) < 7 * 24 * 60 * 60 * 1000
      
      if (accessToken && refreshToken && isTokenValid) {
        this.accessToken = accessToken
        this.refreshToken = refreshToken
        this.isAuthenticated = true
        
        // Busca os dados do usuário
        this.fetchUserProfile().catch(() => {
          // Se falhar ao buscar o perfil, tenta renovar o token
          this.refreshToken().catch(() => {
            // Se falhar ao renovar o token, faz logout
            this.clearAuthState()
          })
        })
      } else {
        // Se os tokens não existem ou estão muito antigos, limpa o estado
        this.clearAuthState()
      }
    }
  }
})
