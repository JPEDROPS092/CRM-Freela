<template>
  <div>
    <div class="mb-6 sm:flex sm:items-center sm:justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900">Pagamentos</h1>
        <p class="mt-2 text-sm text-gray-700">
          Gerencie seus pagamentos e recebimentos
        </p>
      </div>
      <div class="mt-4 sm:mt-0">
        <button
          @click="showNewPaymentModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Novo Pagamento
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
              placeholder="Buscar pagamentos..."
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
            <option value="paid">Pagos</option>
            <option value="overdue">Vencidos</option>
            <option value="cancelled">Cancelados</option>
          </select>
        </div>

        <div class="mt-4 sm:mt-0">
          <select
            v-model="filters.clientId"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm rounded-md"
          >
            <option value="">Todos os clientes</option>
            <option v-for="client in clients" :key="client.id" :value="client.id">
              {{ client.name }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- Lista de pagamentos -->
    <PaymentList
      title="Todos os Pagamentos"
      :payments="payments"
      :show-pagination="true"
      :current-page="currentPage"
      :total-pages="totalPages"
      @edit="handleEdit"
      @delete="handleDelete"
      @page-change="handlePageChange"
      @add="showNewPaymentModal = true"
    />

    <!-- Modal de novo/editar pagamento -->
    <Modal
      v-if="showNewPaymentModal || showEditPaymentModal"
      :title="editingPayment ? 'Editar Pagamento' : 'Novo Pagamento'"
      @close="closeModal"
    >
      <PaymentForm
        :payment="editingPayment"
        :button-text="editingPayment ? 'Salvar Alterações' : 'Criar Pagamento'"
        @submit="handleFormSubmit"
        @cancel="closeModal"
      />
    </Modal>

    <!-- Modal de confirmação de exclusão -->
    <Modal
      v-if="showDeleteModal"
      title="Excluir Pagamento"
      @close="showDeleteModal = false"
    >
      <div class="p-6">
        <p class="text-gray-700 mb-6">
          Tem certeza que deseja excluir este pagamento? Esta ação não pode ser desfeita.
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
import { usePaymentsStore } from '~/store/payments'
import { useClientsStore } from '~/store/clients'
import { useNotificationsStore } from '~/store/notifications'

// Define o título da página
useHead({
  title: 'Pagamentos - CRM Freelancer'
})

const paymentsStore = usePaymentsStore()
const clientsStore = useClientsStore()
const notificationsStore = useNotificationsStore()

const payments = ref([])
const clients = ref([])
const totalPages = ref(1)
const currentPage = ref(1)
const filters = ref({
  search: '',
  status: '',
  clientId: ''
})

const showNewPaymentModal = ref(false)
const showEditPaymentModal = ref(false)
const showDeleteModal = ref(false)
const editingPayment = ref(null)
const paymentToDelete = ref(null)

// Carrega os pagamentos
const loadPayments = async () => {
  try {
    const response = await paymentsStore.fetchPayments(
      currentPage.value,
      10,
      filters.value.status,
      filters.value.clientId ? parseInt(filters.value.clientId) : undefined
    )
    
    payments.value = response?.payments || []
    totalPages.value = response?.totalPages || 1
  } catch (error) {
    notificationsStore.showError('Erro ao carregar pagamentos')
  }
}

// Carrega os clientes para o filtro
const loadClients = async () => {
  try {
    await clientsStore.fetchClients(1, 100)
    clients.value = clientsStore.getClients
  } catch (error) {
    notificationsStore.showError('Erro ao carregar clientes')
  }
}

// Manipuladores de eventos
const handleEdit = (payment: any) => {
  editingPayment.value = payment
  showEditPaymentModal.value = true
}

const handleDelete = (payment: any) => {
  paymentToDelete.value = payment
  showDeleteModal.value = true
}

const handlePageChange = (page: number) => {
  currentPage.value = page
}

const handleFormSubmit = async (paymentData: any) => {
  try {
    if (editingPayment.value) {
      await paymentsStore.updatePayment(editingPayment.value.id, paymentData)
      notificationsStore.showSuccess('Pagamento atualizado com sucesso')
    } else {
      await paymentsStore.createPayment(paymentData)
      notificationsStore.showSuccess('Pagamento criado com sucesso')
    }
    
    closeModal()
    loadPayments()
  } catch (error: any) {
    notificationsStore.showError(error.message || 'Erro ao salvar pagamento')
  }
}

const confirmDelete = async () => {
  try {
    if (paymentToDelete.value) {
      await paymentsStore.deletePayment(paymentToDelete.value.id)
      notificationsStore.showSuccess('Pagamento excluído com sucesso')
      showDeleteModal.value = false
      loadPayments()
    }
  } catch (error: any) {
    notificationsStore.showError(error.message || 'Erro ao excluir pagamento')
  }
}

const closeModal = () => {
  showNewPaymentModal.value = false
  showEditPaymentModal.value = false
  editingPayment.value = null
}

// Observa mudanças nos filtros
watch(filters, () => {
  currentPage.value = 1
  loadPayments()
}, { deep: true })

// Carrega dados iniciais
onMounted(() => {
  loadPayments()
  loadClients()
})
</script>

<style scoped>
/* Estilos específicos da página, se necessário */
</style>
