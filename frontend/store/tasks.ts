import { defineStore } from 'pinia'

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

interface TasksState {
  tasks: Task[]
  currentTask: Task | null
  totalTasks: number
  upcomingTasks: Task[]
  loading: boolean
  error: string | null
}

export const useTasksStore = defineStore('tasks', {
  state: (): TasksState => ({
    tasks: [],
    currentTask: null,
    totalTasks: 0,
    upcomingTasks: [],
    loading: false,
    error: null
  }),

  getters: {
    getTasks: (state) => state.tasks,
    getCurrentTask: (state) => state.currentTask,
    getTotalTasks: (state) => state.totalTasks,
    getUpcomingTasks: (state) => state.upcomingTasks,
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
          throw new Error('Falha ao buscar tarefas')
        }

        const data = await response.json()
        this.tasks = data.tasks
        this.totalTasks = data.total
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar tarefa')
        }

        const data = await response.json()
        this.currentTask = data
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao criar tarefa')
        }

        const data = await response.json()
        this.tasks.push(data.task)
        this.totalTasks++
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao atualizar tarefa')
        }

        const data = await response.json()
        const index = this.tasks.findIndex(t => t.id === id)
        if (index !== -1) {
          this.tasks[index] = data.task
        }
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao excluir tarefa')
        }

        this.tasks = this.tasks.filter(t => t.id !== id)
        this.totalTasks--
        
      } catch (error: any) {
        this.error = error.message
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
          throw new Error('Falha ao buscar tarefas próximas')
        }

        const data = await response.json()
        this.upcomingTasks = data.tasks
        
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    }
  }
})
