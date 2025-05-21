<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h2 class="text-xl font-semibold text-gray-900">{{ title }}</h2>
      <button
        v-if="showAddButton"
        @click="$emit('add')"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary"
      >
        Adicionar Tarefa
      </button>
    </div>

    <div v-if="tasks.length === 0" class="text-center py-10">
      <p class="text-gray-500">Nenhuma tarefa encontrada.</p>
    </div>

    <div v-else class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Título</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Cliente</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Data de Entrega</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Prioridade</th>
            <th scope="col" class="relative px-6 py-3">
              <span class="sr-only">Ações</span>
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="task in tasks" :key="task.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              {{ task.title }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ task.client_name }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(task.due_date) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                :class="{
                  'bg-green-100 text-green-800': task.status === 'completed',
                  'bg-yellow-100 text-yellow-800': task.status === 'in_progress',
                  'bg-gray-100 text-gray-800': task.status === 'pending',
                  'bg-red-100 text-red-800': task.status === 'cancelled'
                }"
              >
                {{ getStatusLabel(task.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                :class="{
                  'bg-red-100 text-red-800': task.priority === 'high',
                  'bg-yellow-100 text-yellow-800': task.priority === 'medium',
                  'bg-green-100 text-green-800': task.priority === 'low'
                }"
              >
                {{ getPriorityLabel(task.priority) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button
                @click="$emit('edit', task)"
                class="text-primary hover:text-primary-dark mr-3"
              >
                Editar
              </button>
              <button
                @click="$emit('delete', task)"
                class="text-red-600 hover:text-red-900"
              >
                Excluir
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Paginação -->
    <div v-if="showPagination && totalPages > 1" class="flex items-center justify-between border-t border-gray-200 px-4 py-3 sm:px-6">
      <div class="flex flex-1 justify-between sm:hidden">
        <button
          @click="$emit('page-change', currentPage - 1)"
          :disabled="currentPage === 1"
          :class="[
            currentPage === 1 ? 'cursor-not-allowed opacity-50' : 'hover:bg-gray-50',
            'relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700'
          ]"
        >
          Anterior
        </button>
        <button
          @click="$emit('page-change', currentPage + 1)"
          :disabled="currentPage === totalPages"
          :class="[
            currentPage === totalPages ? 'cursor-not-allowed opacity-50' : 'hover:bg-gray-50',
            'relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700'
          ]"
        >
          Próximo
        </button>
      </div>
      <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Mostrando <span class="font-medium">{{ (currentPage - 1) * 10 + 1 }}</span> a
            <span class="font-medium">{{ Math.min(currentPage * 10, totalItems) }}</span> de
            <span class="font-medium">{{ totalItems }}</span> resultados
          </p>
        </div>
        <div>
          <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
            <button
              @click="$emit('page-change', currentPage - 1)"
              :disabled="currentPage === 1"
              class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              :class="{ 'cursor-not-allowed opacity-50': currentPage === 1 }"
            >
              <span class="sr-only">Anterior</span>
              <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z" clip-rule="evenodd" />
              </svg>
            </button>
            <button
              @click="$emit('page-change', currentPage + 1)"
              :disabled="currentPage === totalPages"
              class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              :class="{ 'cursor-not-allowed opacity-50': currentPage === totalPages }"
            >
              <span class="sr-only">Próximo</span>
              <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

interface Task {
  id: number;
  title: string;
  client_name: string;
  due_date: string;
  status: string;
  priority: string;
}

const props = defineProps<{
  title?: string;
  tasks: Task[];
  showAddButton?: boolean;
  currentPage: number;
  totalPages: number;
  totalItems: number;
  showPagination?: boolean;
}>()

const emit = defineEmits<{
  (e: 'add'): void;
  (e: 'edit', task: Task): void;
  (e: 'delete', task: Task): void;
  (e: 'page-change', page: number): void;
}>()

// Formatar data para o formato brasileiro
const formatDate = (dateString: string): string => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('pt-BR').format(date)
}

// Obter rótulo de status
const getStatusLabel = (status: string): string => {
  const statusMap: Record<string, string> = {
    pending: 'Pendente',
    in_progress: 'Em Andamento',
    completed: 'Concluída',
    cancelled: 'Cancelada'
  }
  return statusMap[status] || status
}

// Obter rótulo de prioridade
const getPriorityLabel = (priority: string): string => {
  const priorityMap: Record<string, string> = {
    low: 'Baixa',
    medium: 'Média',
    high: 'Alta'
  }
  return priorityMap[priority] || priority
}
</script>
