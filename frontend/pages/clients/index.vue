&lt;template>
  &lt;div>
    &lt;div class="mb-6 sm:flex sm:items-center sm:justify-between">
      &lt;div>
        &lt;h1 class="text-2xl font-semibold text-gray-900">Clientes&lt;/h1>
        &lt;p class="mt-2 text-sm text-gray-700">
          Gerencie seus clientes e projetos relacionados
        &lt;/p>
      &lt;/div>
      &lt;div class="mt-4 sm:mt-0">
        &lt;button
          @click="showNewClientModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          &lt;svg class="-ml-1 mr-2 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            &lt;path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          &lt;/svg>
          Novo Cliente
        &lt;/button>
      &lt;/div>
    &lt;/div>

    &lt;!-- Filtros -->
    &lt;div class="mb-6">
      &lt;div class="sm:flex sm:items-center sm:space-x-4">
        &lt;div class="flex-1">
          &lt;label for="search" class="sr-only">Buscar&lt;/label>
          &lt;div class="relative rounded-md shadow-sm">
            &lt;div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              &lt;svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                &lt;path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              &lt;/svg>
            &lt;/div>
            &lt;input
              type="search"
              id="search"
              v-model="filters.search"
              class="focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md"
              placeholder="Buscar clientes..."
            />
          &lt;/div>
        &lt;/div>

        &lt;div class="mt-4 sm:mt-0">
          &lt;select
            v-model="filters.status"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm rounded-md"
          >
            &lt;option value="">Todos os status&lt;/option>
            &lt;option value="active">Ativos&lt;/option>
            &lt;option value="inactive">Inativos&lt;/option>
            &lt;option value="archived">Arquivados&lt;/option>
          &lt;/select>
        &lt;/div>
      &lt;/div>
    &lt;/div>

    &lt;!-- Lista de clientes -->
    &lt;ClientList
      title="Todos os Clientes"
      :clients="clients"
      :show-pagination="true"
      :current-page="currentPage"
      :total-pages="totalPages"
      @edit="handleEdit"
      @delete="handleDelete"
      @page-change="handlePageChange"
    />

    &lt;!-- Modal de novo/editar cliente -->
    &lt;Modal
      v-if="showNewClientModal || showEditClientModal"
      :title="editingClient ? 'Editar Cliente' : 'Novo Cliente'"
      @close="closeModal"
    >
      &lt;ClientForm
        :client="editingClient"
        :button-text="editingClient ? 'Salvar Alterações' : 'Criar Cliente'"
        @submit="handleFormSubmit"
        @cancel="closeModal"
      />
    &lt;/Modal>

    &lt;!-- Modal de confirmação de exclusão -->
    &lt;ConfirmationModal
      v-if="showDeleteModal"
      title="Excluir Cliente"
      message="Tem certeza que deseja excluir este cliente? Esta ação não pode ser desfeita."
      @confirm="confirmDelete"
      @cancel="showDeleteModal = false"
    />
  &lt;/div>
&lt;/template>

&lt;script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useClientsStore } from '~/store/clients'
import { useNotificationsStore } from '~/store/notifications'

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
    
    clients.value = response.clients
    totalPages.value = response.totalPages
  } catch (error) {
    notificationsStore.showError('Erro ao carregar clientes')
  }
}

// Manipuladores de eventos
const handleEdit = (client: any) => {
  editingClient.value = client
  showEditClientModal.value = true
}

const handleDelete = (client: any) => {
  clientToDelete.value = client
  showDeleteModal.value = true
}

const handlePageChange = (page: number) => {
  currentPage.value = page
}

const handleFormSubmit = async (clientData: any) => {
  try {
    if (editingClient.value) {
      await clientsStore.updateClient(editingClient.value.id, clientData)
      notificationsStore.showSuccess('Cliente atualizado com sucesso')
    } else {
      await clientsStore.createClient(clientData)
      notificationsStore.showSuccess('Cliente criado com sucesso')
    }
    
    closeModal()
    loadClients()
  } catch (error: any) {
    notificationsStore.showError(error.message || 'Erro ao salvar cliente')
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

// Carrega dados iniciais
onMounted(() => {
  loadClients()
})
&lt;/script>
