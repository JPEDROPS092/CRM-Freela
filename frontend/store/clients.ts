import { defineStore } from 'pinia'
import { useRuntimeConfig } from '#app'

interface Client {
  id: number
  name: string
  email: string
  phone: string
  company: string
  notes: string
  status: string
  created_at: string
  updated_at: string
}

interface ClientStats {
  total: number
  active: number
  inactive: number
  archived: number
}

interface ClientsState {
  clients: Client[]
  currentClient: Client | null
  totalClients: number
  activeClients: number
  stats: ClientStats
  loading: boolean
  error: string | null
}

export const useClientsStore = defineStore('clients', {
  state: (): ClientsState => ({
    clients: [],
    currentClient: null,
    totalClients: 0,
    activeClients: 0,
    stats: {
      total: 0,
      active: 0,
      inactive: 0,
      archived: 0
    },
    loading: false,
    error: null
  }),

  getters: {
    getClients: (state) => state.clients,
    getCurrentClient: (state) => state.currentClient,
    getTotalClients: (state) => state.totalClients,
    getActiveClients: (state) => state.activeClients,
    getStats: (state) => state.stats,
    isLoading: (state) => state.loading,
    getError: (state) => state.error
  },

  actions: {
    async fetchClients(page: number = 1, pageSize: number = 10, search?: string, status?: string) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        let url = `${config.public.apiBase}/clients?page=${page}&page_size=${pageSize}`
        if (search) url += `&search=${encodeURIComponent(search)}`
        if (status) url += `&status=${status}`
        
        const response = await fetch(
          url,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar clientes')
        }

        const data = await response.json()
        this.clients = data.data || []
        this.totalClients = data.meta?.total || 0
        
        return {
          clients: this.clients,
          totalPages: Math.ceil(this.totalClients / pageSize)
        }
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchClientById(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients/${id}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar cliente')
        }

        const data = await response.json()
        this.currentClient = data
        
        return this.currentClient
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async createClient(clientData: Omit<Client, 'id' | 'created_at' | 'updated_at'>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients`,
          {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(clientData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao criar cliente')
        }

        const data = await response.json()
        
        // Atualiza as estatísticas
        this.stats.total++
        if (clientData.status === 'active') this.stats.active++
        else if (clientData.status === 'inactive') this.stats.inactive++
        else if (clientData.status === 'archived') this.stats.archived++
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async updateClient(id: number, clientData: Partial<Client>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients/${id}`,
          {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(clientData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao atualizar cliente')
        }

        const data = await response.json()
        const index = this.clients.findIndex((c: Client) => c.id === id)
        if (index !== -1) {
          // Se o status mudou, atualiza as estatísticas
          if (clientData.status && this.clients[index].status !== clientData.status) {
            // Decrementa o contador do status antigo
            if (this.clients[index].status === 'active') this.stats.active--
            else if (this.clients[index].status === 'inactive') this.stats.inactive--
            else if (this.clients[index].status === 'archived') this.stats.archived--
            
            // Incrementa o contador do novo status
            if (clientData.status === 'active') this.stats.active++
            else if (clientData.status === 'inactive') this.stats.inactive++
            else if (clientData.status === 'archived') this.stats.archived++
          }
          
          this.clients[index] = data
        }
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async deleteClient(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients/${id}`,
          {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao excluir cliente')
        }

        // Remove o cliente da lista
        const clientToRemove = this.clients.find((c: Client) => c.id === id)
        this.clients = this.clients.filter((c: Client) => c.id !== id)
        
        // Atualiza as estatísticas
        if (clientToRemove) {
          this.stats.total--
          if (clientToRemove.status === 'active') this.stats.active--
          else if (clientToRemove.status === 'inactive') this.stats.inactive--
          else if (clientToRemove.status === 'archived') this.stats.archived--
        }
        
        return true
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async fetchStats() {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients/stats`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar estatísticas de clientes')
        }

        const data = await response.json()
        this.stats = data
        
        return this.stats
      } catch (error: any) {
        this.error = error.message
        return {
          total: 0,
          active: 0,
          inactive: 0,
          archived: 0
        }
      } finally {
        this.loading = false
      }
    }
  }
})
