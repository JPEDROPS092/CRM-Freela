&lt;template>
  &lt;div class="space-y-4">
    &lt;div class="flex justify-between items-center">
      &lt;h2 class="text-lg font-semibold text-gray-900">{{ title }}&lt;/h2>
      &lt;button
        v-if="showAddButton"
        @click="$emit('add')"
        class="px-4 py-2 bg-primary text-white rounded-md text-sm font-medium hover:bg-primary-dark"
      >
        Adicionar Cliente
      &lt;/button>
    &lt;/div>

    &lt;div class="bg-white shadow overflow-hidden sm:rounded-md">
      &lt;ul v-if="clients.length > 0" class="divide-y divide-gray-200">
        &lt;li v-for="client in clients" :key="client.id" class="hover:bg-gray-50">
          &lt;div class="px-4 py-4 sm:px-6">
            &lt;div class="flex items-center justify-between">
              &lt;div class="flex-1 min-w-0">
                &lt;p class="text-sm font-medium text-primary truncate">
                  {{ client.name }}
                &lt;/p>
                &lt;p class="mt-1 text-sm text-gray-500">
                  {{ client.email || 'Sem email' }}
                &lt;/p>
              &lt;/div>
              &lt;div class="ml-4 flex items-center space-x-3">
                &lt;span
                  :class="{
                    'px-2 py-1 text-xs font-medium rounded-full': true,
                    'bg-green-100 text-green-800': client.status === 'active',
                    'bg-yellow-100 text-yellow-800': client.status === 'inactive',
                    'bg-gray-100 text-gray-800': client.status === 'archived'
                  }"
                >
                  {{ getStatusLabel(client.status) }}
                &lt;/span>
                &lt;div class="flex items-center space-x-2">
                  &lt;button
                    @click="$emit('edit', client)"
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
                    @click="$emit('delete', client)"
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
                &lt;p v-if="client.phone" class="flex items-center text-sm text-gray-500">
                  &lt;svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    &lt;path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
                    />
                  &lt;/svg>
                  {{ client.phone }}
                &lt;/p>
                &lt;p v-if="client.company" class="mt-2 flex items-center text-sm text-gray-500 sm:mt-0 sm:ml-6">
                  &lt;svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    &lt;path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                    />
                  &lt;/svg>
                  {{ client.company }}
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
                Cliente desde {{ formatDate(client.created_at) }}
              &lt;/div>
            &lt;/div>
          &lt;/div>
        &lt;/li>
      &lt;/ul>
      &lt;div v-else class="p-4 text-center text-gray-500">
        Nenhum cliente encontrado
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
  clients: Array<{
    id: number
    name: string
    email: string
    phone: string
    company: string
    status: string
    created_at: string
  }>
  showAddButton?: boolean
  showPagination?: boolean
  currentPage?: number
  totalPages?: number
}>()

defineEmits<{
  (e: 'add'): void
  (e: 'edit', client: any): void
  (e: 'delete', client: any): void
  (e: 'page-change', page: number): void
}>()

const getStatusLabel = (status: string) => {
  const labels = {
    active: 'Ativo',
    inactive: 'Inativo',
    archived: 'Arquivado'
  }
  return labels[status as keyof typeof labels] || status
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('pt-BR')
}
&lt;/script>
