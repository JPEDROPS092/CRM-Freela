import { defineStore } from 'pinia'
import { useRuntimeConfig } from '#app'

interface Payment {
  id: number
  amount: number
  description: string
  due_date: string
  status: string
  client_id: number
  task_id?: number
  client?: {
    id: number
    name: string
  }
  task?: {
    id: number
    title: string
  }
  created_at: string
  updated_at: string
}

interface PaymentSummary {
  start_date: string
  end_date: string
  total: number
}

interface PaymentStats {
  total: number
  pending: number
  paid: number
  overdue: number
  cancelled: number
  monthlyRevenue: number
}

interface PaymentsState {
  payments: Payment[]
  currentPayment: Payment | null
  totalPayments: number
  overduePayments: Payment[]
  summary: PaymentSummary | null
  stats: PaymentStats
  loading: boolean
  error: string | null
}

export const usePaymentsStore = defineStore('payments', {
  state: (): PaymentsState => ({
    payments: [],
    currentPayment: null,
    totalPayments: 0,
    overduePayments: [],
    summary: null,
    stats: {
      total: 0,
      pending: 0,
      paid: 0,
      overdue: 0,
      cancelled: 0,
      monthlyRevenue: 0
    },
    loading: false,
    error: null
  }),

  getters: {
    getPayments: (state) => state.payments,
    getCurrentPayment: (state) => state.currentPayment,
    getTotalPayments: (state) => state.totalPayments,
    getOverduePayments: (state) => state.overduePayments,
    getSummary: (state) => state.summary,
    getStats: (state) => state.stats,
    isLoading: (state) => state.loading,
    getError: (state) => state.error
  },

  actions: {
    async fetchPayments(page: number = 1, pageSize: number = 10, status?: string, clientId?: number, taskId?: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        let url = `${config.public.apiBase}/payments?page=${page}&page_size=${pageSize}`
        if (status) url += `&status=${status}`
        if (clientId) url += `&client_id=${clientId}`
        if (taskId) url += `&task_id=${taskId}`

        const response = await fetch(url, {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
          }
        })
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar pagamentos')
        }

        const data = await response.json()
        this.payments = data.data || []
        this.totalPayments = data.meta?.total || 0
        
        return {
          payments: this.payments,
          totalPages: Math.ceil(this.totalPayments / pageSize)
        }
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchPaymentById(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments/${id}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar pagamento')
        }

        const data = await response.json()
        this.currentPayment = data
        
        return this.currentPayment
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async createPayment(paymentData: Omit<Payment, 'id' | 'created_at' | 'updated_at' | 'client' | 'task'>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments`,
          {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(paymentData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao criar pagamento')
        }

        const data = await response.json()
        
        // Atualiza as estatísticas
        this.stats.total++
        if (paymentData.status === 'pending') this.stats.pending++
        else if (paymentData.status === 'paid') {
          this.stats.paid++
          this.stats.monthlyRevenue += paymentData.amount
        }
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async updatePayment(id: number, paymentData: Partial<Payment>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments/${id}`,
          {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(paymentData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao atualizar pagamento')
        }

        const data = await response.json()
        const index = this.payments.findIndex((p: Payment) => p.id === id)
        if (index !== -1) {
          this.payments[index] = data
        }
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async deletePayment(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments/${id}`,
          {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao excluir pagamento')
        }

        // Remove o pagamento da lista
        const paymentToRemove = this.payments.find((p: Payment) => p.id === id)
        this.payments = this.payments.filter((p: Payment) => p.id !== id)
        
        // Atualiza as estatísticas
        if (paymentToRemove) {
          this.stats.total--
          if (paymentToRemove.status === 'pending') this.stats.pending--
          else if (paymentToRemove.status === 'paid') {
            this.stats.paid--
            this.stats.monthlyRevenue -= paymentToRemove.amount
          }
          else if (paymentToRemove.status === 'overdue') this.stats.overdue--
          else if (paymentToRemove.status === 'cancelled') this.stats.cancelled--
        }
        
        return true
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchOverduePayments() {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments/overdue`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar pagamentos vencidos')
        }

        const data = await response.json()
        this.overduePayments = data.payments || []
        
        return this.overduePayments
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchPaymentSummary(startDate: string, endDate: string) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/payments/summary?start_date=${startDate}&end_date=${endDate}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar resumo de pagamentos')
        }

        const data = await response.json()
        this.summary = data
        
        return this.summary
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
          `${config.public.apiBase}/payments/stats`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar estatísticas de pagamentos')
        }

        const data = await response.json()
        this.stats = data
        
        return this.stats
      } catch (error: any) {
        this.error = error.message
        return {
          total: 0,
          pending: 0,
          paid: 0,
          overdue: 0,
          cancelled: 0,
          monthlyRevenue: 0
        }
      } finally {
        this.loading = false
      }
    }
  }
})
