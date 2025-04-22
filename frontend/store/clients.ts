import { defineStore } from 'pinia'

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

interface ClientsState {
  clients: Client[]
  currentClient: Client | null
  totalClients: number
  activeClients: number
  loading: boolean
  error: string | null
}

export const useClientsStore = defineStore('clients', {
  state: (): ClientsState => ({
    clients: [],
    currentClient: null,
    totalClients: 0,
    activeClients: 0,
    loading: false,
    error: null
  }),

  getters: {
    getClients: (state) => state.clients,
    getCurrentClient: (state) => state.currentClient,
    getTotalClients: (state) => state.totalClients,
    getActiveClients: (state) => state.activeClients,
    isLoading: (state) => state.loading,
    getError: (state) => state.error
  },

  actions: {
    async fetchClients(page: number = 1, pageSize: number = 10) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients?page=${page}&page_size=${pageSize}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          throw new Error('Falha ao buscar clientes')
        }

        const data = await response.json()
        this.clients = data.clients
        this.totalClients = data.total
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar cliente')
        }

        const data = await response.json()
        this.currentClient = data
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao criar cliente')
        }

        const data = await response.json()
        this.clients.push(data.client)
        this.totalClients++
        if (data.client.status === 'active') {
          this.activeClients++
        }
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao atualizar cliente')
        }

        const data = await response.json()
        const index = this.clients.findIndex(c => c.id === id)
        if (index !== -1) {
          this.clients[index] = data.client
        }
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao excluir cliente')
        }

        this.clients = this.clients.filter(c => c.id !== id)
        this.totalClients--
        
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    },

    async fetchActiveClientsCount() {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/clients/count/active`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          throw new Error('Falha ao buscar contagem de clientes ativos')
        }

        const data = await response.json()
        this.activeClients = data.count
        
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    }
  }
})
