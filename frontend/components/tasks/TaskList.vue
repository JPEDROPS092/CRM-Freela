&lt;template>
  &lt;div class="space-y-4">
    &lt;div class="flex justify-between items-center">
      &lt;h2 class="text-lg font-semibold text-gray-900">{{ title }}&lt;/h2>
      &lt;button
        v-if="showAddButton"
        @click="$emit('add')"
        class="px-4 py-2 bg-primary text-white rounded-md text-sm font-medium hover:bg-primary-dark"
      >
        Adicionar Tarefa
      &lt;/button>
    &lt;/div>

    &lt;div class="bg-white shadow overflow-hidden sm:rounded-md">
      &lt;ul v-if="tasks.length > 0" class="divide-y divide-gray-200">
        &lt;li v-for="task in tasks" :key="task.id" class="hover:bg-gray-50">
          &lt;div class="px-4 py-4 sm:px-6">
            &lt;div class="flex items-center justify-between">
              &lt;div class="flex-1 min-w-0">
                &lt;div class="flex items-center">
                  &lt;p class="text-sm font-medium text-primary truncate">
                    {{ task.title }}
                  &lt;/p>
                  &lt;span
                    :class="{
                      'ml-2 px-2 py-1 text-xs font-medium rounded-full': true,
                      'bg-yellow-100 text-yellow-800': task.priority === 'high',
                      'bg-blue-100 text-blue-800': task.priority === 'medium',
                      'bg-gray-100 text-gray-800': task.priority === 'low'
                    }"
                  >
                    {{ getPriorityLabel(task.priority) }}
                  &lt;/span>
                &lt;/div>
                &lt;p v-if="task.description" class="mt-1 text-sm text-gray-500 line-clamp-2">
                  {{ task.description }}
                &lt;/p>
              &lt;/div>
              &lt;div class="ml-4 flex items-center space-x-3">
                &lt;span
                  :class="{
                    'px-2 py-1 text-xs font-medium rounded-full': true,
                    'bg-blue-100 text-blue-800': task.status === 'pending',
                    'bg-yellow-100 text-yellow-800': task.status === 'in_progress',
                    'bg-green-100 text-green-800': task.status === 'completed',
                    'bg-red-100 text-red-800': task.status === 'cancelled'
                  }"
                >
                  {{ getStatusLabel(task.status) }}
                &lt;/span>
                &lt;div class="flex items-center space-x-2">
                  &lt;button
                    @click="$emit('edit', task)"
                    class="text-primary hover:text-primary-dark"
                  >
                    &lt;span class="sr-only">Editar&lt;/span>
                    &lt;svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      &lt;path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                      />
                    &lt;/svg>
                  &lt;/button>
                  &lt;button
                    @click="$emit('delete', task)"
                    class="text-red-600 hover:text-red-800"
                  >
                    &lt;span class="sr-only">Excluir&lt;/span>
                    &lt;svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      &lt;path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                      />
                    &lt;/svg>
                  &lt;/button>
                &lt;/div>
              &lt;/div>
            &lt;/div>
            &lt;div class="mt-2 sm:flex sm:justify-between">
              &lt;div class="sm:flex">
                &lt;p class="flex items-center text-sm text-gray-500">
                  &lt;svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    &lt;path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                    />
                  &lt;/svg>
                  {{ task.client?.name || 'Cliente não especificado' }}
                &lt;/p>
              &lt;/div>
              &lt;div class="mt-2 flex items-center text-sm text-gray-500 sm:mt-0">
                &lt;svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  &lt;path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                  />
                &lt;/svg>
                Entrega: {{ formatDate(task.due_date) }}
              &lt;/div>
            &lt;/div>
          &lt;/div>
        &lt;/li>
      &lt;/ul>
      &lt;div v-else class="p-4 text-center text-gray-500">
        Nenhuma tarefa encontrada
      &lt;/div>
    &lt;/div>

    &lt;div v-if="showPagination && totalPages > 1" class="flex justify-center mt-4">
      &lt;nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
        &lt;button
          @click="$emit('page-change', currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
          :class="{ 'opacity-50 cursor-not-allowed': currentPage === 1 }"
        >
          Anterior
        &lt;/button>
        &lt;button
          v-for="page in totalPages"
          :key="page"
          @click="$emit('page-change', page)"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium"
          :class="page === currentPage ? 'text-primary bg-primary-50' : 'text-gray-700 hover:bg-gray-50'"
        >
          {{ page }}
        &lt;/button>
        &lt;button
          @click="$emit('page-change', currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
          :class="{ 'opacity-50 cursor-not-allowed': currentPage === totalPages }"
        >
          Próxima
        &lt;/button>
      &lt;/nav>
    &lt;/div>
  &lt;/div>
&lt;/template>

&lt;script setup lang="ts">
defineProps<{
  title?: string
  tasks: Array<{
    id: number
    title: string
    description: string
    due_date: string
    status: string
    priority: string
    client?: {
      id: number
      name: string
    }
  }>
  showAddButton?: boolean
  showPagination?: boolean
  currentPage?: number
  totalPages?: number
}>()

defineEmits<{
  (e: 'add'): void
  (e: 'edit', task: any): void
  (e: 'delete', task: any): void
  (e: 'page-change', page: number): void
}>()

const getStatusLabel = (status: string) => {
  const labels = {
    pending: 'Pendente',
    in_progress: 'Em Andamento',
    completed: 'Concluída',
    cancelled: 'Cancelada'
  }
  return labels[status as keyof typeof labels] || status
}

const getPriorityLabel = (priority: string) => {
  const labels = {
    low: 'Baixa',
    medium: 'Média',
    high: 'Alta'
  }
  return labels[priority as keyof typeof labels] || priority
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('pt-BR')
}
&lt;/script>
