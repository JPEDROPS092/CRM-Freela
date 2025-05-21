import { defineStore } from 'pinia'
import { useRuntimeConfig } from '#app'

interface Task {
  id: number
  title: string
  description: string
  due_date: string
  status: string
  priority: string
  client_id: number
  client?: {
    id: number
    name: string
  }
  created_at: string
  updated_at: string
}

interface TaskStats {
  total: number
  pending: number
  in_progress: number
  completed: number
  cancelled: number
}

interface TasksState {
  tasks: Task[]
  currentTask: Task | null
  totalTasks: number
  upcomingTasks: Task[]
  stats: TaskStats
  loading: boolean
  error: string | null
}

export const useTasksStore = defineStore('tasks', {
  state: (): TasksState => ({
    tasks: [],
    currentTask: null,
    totalTasks: 0,
    upcomingTasks: [],
    stats: {
      total: 0,
      pending: 0,
      in_progress: 0,
      completed: 0,
      cancelled: 0
    },
    loading: false,
    error: null
  }),

  getters: {
    getTasks: (state) => state.tasks,
    getCurrentTask: (state) => state.currentTask,
    getTotalTasks: (state) => state.totalTasks,
    getUpcomingTasks: (state) => state.upcomingTasks,
    getStats: (state) => state.stats,
    isLoading: (state) => state.loading,
    getError: (state) => state.error
  },

  actions: {
    async fetchTasks(page: number = 1, pageSize: number = 10, status?: string, clientId?: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        let url = `${config.public.apiBase}/tasks?page=${page}&page_size=${pageSize}`
        if (status) url += `&status=${status}`
        if (clientId) url += `&client_id=${clientId}`

        const response = await fetch(url, {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
          }
        })
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar tarefas')
        }

        const data = await response.json()
        this.tasks = data.data || []
        this.totalTasks = data.meta?.total || 0
        
        return {
          tasks: this.tasks,
          totalPages: Math.ceil(this.totalTasks / pageSize)
        }
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchTaskById(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/tasks/${id}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar tarefa')
        }

        const data = await response.json()
        this.currentTask = data
        
        return this.currentTask
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async createTask(taskData: Omit<Task, 'id' | 'created_at' | 'updated_at' | 'client'>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/tasks`,
          {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(taskData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao criar tarefa')
        }

        const data = await response.json()
        
        // Atualiza as estatísticas
        this.stats.total++
        if (taskData.status === 'pending') this.stats.pending++
        else if (taskData.status === 'in_progress') this.stats.in_progress++
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async updateTask(id: number, taskData: Partial<Task>) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/tasks/${id}`,
          {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: JSON.stringify(taskData)
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao atualizar tarefa')
        }

        const data = await response.json()
        const index = this.tasks.findIndex((t: Task) => t.id === id)
        if (index !== -1) {
          this.tasks[index] = data
        }
        
        return data
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async deleteTask(id: number) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/tasks/${id}`,
          {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao excluir tarefa')
        }

        // Remove a tarefa da lista
        const taskToRemove = this.tasks.find((t: Task) => t.id === id)
        this.tasks = this.tasks.filter((t: Task) => t.id !== id)
        
        // Atualiza as estatísticas
        if (taskToRemove) {
          this.stats.total--
          if (taskToRemove.status === 'pending') this.stats.pending--
          else if (taskToRemove.status === 'in_progress') this.stats.in_progress--
          else if (taskToRemove.status === 'completed') this.stats.completed--
          else if (taskToRemove.status === 'cancelled') this.stats.cancelled--
        }
        
        return true
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchUpcomingTasks(days: number = 7) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(
          `${config.public.apiBase}/tasks/upcoming?days=${days}`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar tarefas próximas')
        }

        const data = await response.json()
        this.upcomingTasks = data.tasks || []
        
        return this.upcomingTasks
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
          `${config.public.apiBase}/tasks/stats`,
          {
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
          }
        )
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Falha ao buscar estatísticas de tarefas')
        }

        const data = await response.json()
        this.stats = data
        
        return this.stats
      } catch (error: any) {
        this.error = error.message
        return {
          total: 0,
          pending: 0,
          in_progress: 0,
          completed: 0,
          cancelled: 0
        }
      } finally {
        this.loading = false
      }
    }
  }
})
