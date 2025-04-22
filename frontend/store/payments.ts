import { defineStore } from 'pinia'

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

interface PaymentsState {
  payments: Payment[]
  currentPayment: Payment | null
  totalPayments: number
  overduePayments: Payment[]
  summary: PaymentSummary | null
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
    loading: false,
    error: null
  }),

  getters: {
    getPayments: (state) => state.payments,
    getCurrentPayment: (state) => state.currentPayment,
    getTotalPayments: (state) => state.totalPayments,
    getOverduePayments: (state) => state.overduePayments,
    getSummary: (state) => state.summary,
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
          throw new Error('Falha ao buscar pagamentos')
        }

        const data = await response.json()
        this.payments = data.payments
        this.totalPayments = data.total
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar pagamento')
        }

        const data = await response.json()
        this.currentPayment = data
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao criar pagamento')
        }

        const data = await response.json()
        this.payments.push(data.payment)
        this.totalPayments++
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao atualizar pagamento')
        }

        const data = await response.json()
        const index = this.payments.findIndex(p => p.id === id)
        if (index !== -1) {
          this.payments[index] = data.payment
        }
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao excluir pagamento')
        }

        this.payments = this.payments.filter(p => p.id !== id)
        this.totalPayments--
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar pagamentos vencidos')
        }

        const data = await response.json()
        this.overduePayments = data.payments
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar resumo de pagamentos')
        }

        const data = await response.json()
        this.summary = data
        
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    }
  }
})
