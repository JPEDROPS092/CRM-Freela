<template>
  <div>
    <div class="mb-6 sm:flex sm:items-center sm:justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900">Clientes</h1>
        <p class="mt-2 text-sm text-gray-700">
          Gerencie seus clientes e projetos relacionados
        </p>
      </div>
      <div class="mt-4 sm:mt-0">
        <button
          @click="showNewClientModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Novo Cliente
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
              placeholder="Buscar clientes..."
            />
          </div>
        </div>

        <div class="mt-4 sm:mt-0">
          <select
            v-model="filters.status"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm rounded-md"
          >
            <option value="">Todos os status</option>
            <option value="active">Ativos</option>
            <option value="inactive">Inativos</option>
            <option value="archived">Arquivados</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Lista de clientes -->
    <ClientList
      title="Todos os Clientes"
      :clients="clients"
      :show-add-button="true"
      :show-pagination="true"
      :current-page="currentPage"
      :total-pages="totalPages"
      @add="showNewClientModal = true"
      @edit="handleEdit"
      @delete="handleDelete"
      @page-change="handlePageChange"
    />

    <!-- Modal de novo/editar cliente -->
    <Modal
      v-if="showNewClientModal || showEditClientModal"
      :title="editingClient ? 'Editar Cliente' : 'Novo Cliente'"
      @close="closeModal"
    >
      <ClientForm
        :client="editingClient"
        :button-text="editingClient ? 'Salvar Alterações' : 'Criar Cliente'"
        @submit="handleFormSubmit"
        @cancel="closeModal"
      />
    </Modal>

    <!-- Modal de confirmação de exclusão -->
    <ConfirmationModal
      v-if="showDeleteModal"
      title="Excluir Cliente"
      message="Tem certeza que deseja excluir este cliente? Esta ação não pode ser desfeita."
      @confirm="confirmDelete"
      @cancel="showDeleteModal = false"
    />
  </div>
</template>

<script setup lang="ts">
// Importações
import { ref, watch, onMounted } from '#imports'
import { useHead } from '#imports'
import { useClientsStore } from '~/store/clients'
import { useNotificationsStore } from '~/store/notifications'
import { Client, ClientsResponse } from '~/types/client'

// Define o título da página
useHead({
  title: 'Clientes - CRM Freelancer'
})

const clientsStore = useClientsStore()
const notificationsStore = useNotificationsStore()

const clients = ref([])
const totalPages = ref(1)
const currentPage = ref(1)
const filters = ref({
  search: '',
  status: ''
})

const showNewClientModal = ref(false)
const showEditClientModal = ref(false)
const showDeleteModal = ref(false)
const editingClient = ref(null)
const clientToDelete = ref(null)

// Carrega os clientes
const loadClients = async () => {
  try {
    const response = await clientsStore.fetchClients(
      currentPage.value,
      10,
      filters.value.search,
      filters.value.status
    )
    
    if (response) {
      clients.value = response.clients || []
      totalPages.value = response.totalPages || Math.ceil((response.total || 0) / 10) || 1
    }
  } catch (error) {
    notificationsStore.showError('Erro ao carregar clientes')
    console.error('Erro ao carregar clientes:', error)
  }
}

// Manipuladores de eventos
const handleEdit = (client) => {
  editingClient.value = { ...client }
  showEditClientModal.value = true
}

const handleDelete = (client) => {
  clientToDelete.value = client
  showDeleteModal.value = true
}

const handlePageChange = (page) => {
  currentPage.value = page
  loadClients()
}

const handleFormSubmit = async (clientData) => {
  try {
    if (editingClient.value) {
      // Atualiza cliente existente
      await clientsStore.updateClient(editingClient.value.id, clientData)
      notificationsStore.showSuccess('Cliente atualizado com sucesso')
    } else {
      // Cria novo cliente
      await clientsStore.createClient(clientData)
      notificationsStore.showSuccess('Cliente criado com sucesso')
    }
    
    // Fecha o modal e recarrega a lista
    closeModal()
    loadClients()
  } catch (error) {
    notificationsStore.showError(editingClient.value ? 'Erro ao atualizar cliente' : 'Erro ao criar cliente')
    console.error('Erro ao salvar cliente:', error)
  }
}

const confirmDelete = async () => {
  if (!clientToDelete.value) return
  
  try {
    await clientsStore.deleteClient(clientToDelete.value.id)
    notificationsStore.showSuccess('Cliente excluído com sucesso')
    showDeleteModal.value = false
    loadClients()
  } catch (error) {
    notificationsStore.showError('Erro ao excluir cliente')
    console.error('Erro ao excluir cliente:', error)
  }
}

const closeModal = () => {
  showNewClientModal.value = false
  showEditClientModal.value = false
  editingClient.value = null
}

// Observa mudanças nos filtros
watch(filters, () => {
  currentPage.value = 1
  loadClients()
}, { deep: true })

// Carrega os clientes quando o componente é montado
onMounted(() => {
  loadClients()
})
</script>
