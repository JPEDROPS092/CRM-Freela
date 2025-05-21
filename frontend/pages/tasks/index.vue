<template>
  <div>
    <div class="mb-6 sm:flex sm:items-center sm:justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900">Tarefas</h1>
        <p class="mt-2 text-sm text-gray-700">
          Gerencie suas tarefas e projetos
        </p>
      </div>
      <div class="mt-4 sm:mt-0">
        <button
          @click="showNewTaskModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Nova Tarefa
        </button>
      </div>
    </div>

    <!-- Filtros -->
    <div class="mb-6">
      <div class="sm:flex sm:items-center sm:space-x-4">
        <div class="flex-1">
          <label for="search" class="sr-only">Buscar</label>
          <div class="relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <input
              type="search"
              id="search"
              v-model="filters.search"
              class="focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md"
              placeholder="Buscar tarefas..."
            />
          </div>
        </div>

        <div class="mt-4 sm:mt-0 sm:ml-4">
          <select
            v-model="filters.status"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm rounded-md"
          >
            <option value="">Todos os status</option>
            <option value="pending">Pendentes</option>
            <option value="in_progress">Em Andamento</option>
            <option value="completed">Concluídas</option>
            <option value="cancelled">Canceladas</option>
          </select>
        </div>

        <div class="mt-4 sm:mt-0">
          <select
            v-model="filters.priority"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm rounded-md"
          >
            <option value="">Todas as prioridades</option>
            <option value="high">Alta</option>
            <option value="medium">Média</option>
            <option value="low">Baixa</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Lista de tarefas -->
    <TaskList
      title="Todas as Tarefas"
      :tasks="tasks"
      :show-pagination="true"
      :current-page="currentPage"
      :total-pages="totalPages"
      @edit="handleEdit"
      @delete="handleDelete"
      @page-change="handlePageChange"
      @add="showNewTaskModal = true"
    />

    <!-- Modal de nova/editar tarefa -->
    <Modal
      v-if="showNewTaskModal || showEditTaskModal"
      :title="editingTask ? 'Editar Tarefa' : 'Nova Tarefa'"
      @close="closeModal"
    >
      <TaskForm
        :task="editingTask"
        :button-text="editingTask ? 'Salvar Alterações' : 'Criar Tarefa'"
        @submit="handleFormSubmit"
        @cancel="closeModal"
      />
    </Modal>

    <!-- Modal de confirmação de exclusão -->
    <Modal
      v-if="showDeleteModal"
      title="Excluir Tarefa"
      @close="showDeleteModal = false"
    >
      <div class="p-6">
        <p class="text-gray-700 mb-6">
          Tem certeza que deseja excluir esta tarefa? Esta ação não pode ser desfeita.
        </p>
        <div class="flex justify-end space-x-3">
          <button 
            @click="showDeleteModal = false" 
            class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
          >
            Cancelar
          </button>
          <button 
            @click="confirmDelete" 
            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
          >
            Excluir
          </button>
        </div>
      </div>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useTasksStore } from '~/store/tasks'
import { useNotificationsStore } from '~/store/notifications'

// Define o título da página
useHead({
  title: 'Tarefas - CRM Freelancer'
})

const tasksStore = useTasksStore()
const notificationsStore = useNotificationsStore()

const tasks = ref([])
const totalPages = ref(1)
const currentPage = ref(1)
const filters = ref({
  search: '',
  status: '',
  priority: '',
  clientId: ''
})

const showNewTaskModal = ref(false)
const showEditTaskModal = ref(false)
const showDeleteModal = ref(false)
const editingTask = ref(null)
const taskToDelete = ref(null)

// Carrega as tarefas
const loadTasks = async () => {
  try {
    const response = await tasksStore.fetchTasks(
      currentPage.value,
      10,
      filters.value.status,
      filters.value.clientId ? parseInt(filters.value.clientId) : undefined
    )
    
    tasks.value = response?.tasks || []
    totalPages.value = response?.totalPages || 1
  } catch (error) {
    notificationsStore.showError('Erro ao carregar tarefas')
  }
}

// Manipuladores de eventos
const handleEdit = (task: any) => {
  editingTask.value = task
  showEditTaskModal.value = true
}

const handleDelete = (task: any) => {
  taskToDelete.value = task
  showDeleteModal.value = true
}

const handlePageChange = (page: number) => {
  currentPage.value = page
}

const handleFormSubmit = async (taskData: any) => {
  try {
    if (editingTask.value) {
      await tasksStore.updateTask(editingTask.value.id, taskData)
      notificationsStore.showSuccess('Tarefa atualizada com sucesso')
    } else {
      await tasksStore.createTask(taskData)
      notificationsStore.showSuccess('Tarefa criada com sucesso')
    }
    
    closeModal()
    loadTasks()
  } catch (error: any) {
    notificationsStore.showError(error.message || 'Erro ao salvar tarefa')
  }
}

const confirmDelete = async () => {
  try {
    if (taskToDelete.value) {
      await tasksStore.deleteTask(taskToDelete.value.id)
      notificationsStore.showSuccess('Tarefa excluída com sucesso')
      showDeleteModal.value = false
      loadTasks()
    }
  } catch (error: any) {
    notificationsStore.showError(error.message || 'Erro ao excluir tarefa')
  }
}

const closeModal = () => {
  showNewTaskModal.value = false
  showEditTaskModal.value = false
  editingTask.value = null
}

// Observa mudanças nos filtros
watch(filters, () => {
  currentPage.value = 1
  loadTasks()
}, { deep: true })

// Carrega as tarefas ao montar o componente
onMounted(() => {
  loadTasks()
})
</script>

<style scoped>
/* Estilos específicos da página, se necessário */
</style>
