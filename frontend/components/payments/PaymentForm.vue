<template>
  <form @submit.prevent="handleSubmit" class="space-y-4">
    <div>
      <label for="amount" class="block text-sm font-medium text-gray-700">Valor</label>
      <div class="mt-1 relative rounded-md shadow-sm">
        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <span class="text-gray-500 sm:text-sm">R$</span>
        </div>
        <input
          type="number"
          id="amount"
          v-model="form.amount"
          required
          step="0.01"
          min="0"
          class="pl-10 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
        />
      </div>
    </div>

    <div>
      <label for="description" class="block text-sm font-medium text-gray-700">Descrição</label>
      <textarea
        id="description"
        v-model="form.description"
        required
        rows="3"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      ></textarea>
    </div>

    <div>
      <label for="due_date" class="block text-sm font-medium text-gray-700">Data de Vencimento</label>
      <input
        type="date"
        id="due_date"
        v-model="form.due_date"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    </div>

    <div>
      <label for="status" class="block text-sm font-medium text-gray-700">Status</label>
      <select
        id="status"
        v-model="form.status"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        <option value="pending">Pendente</option>
        <option value="paid">Pago</option>
        <option value="cancelled">Cancelado</option>
        <option value="overdue">Vencido</option>
      </select>
    </div>

    <div>
      <label for="client_id" class="block text-sm font-medium text-gray-700">Cliente</label>
      <select
        id="client_id"
        v-model="form.client_id"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
        @change="loadClientTasks"
      >
        <option value="">Selecione um cliente</option>
        <option v-for="client in clients" :key="client.id" :value="client.id">
          {{ client.name }}
        </option>
      </select>
    </div>

    <div v-if="form.client_id">
      <label for="task_id" class="block text-sm font-medium text-gray-700">Tarefa (opcional)</label>
      <select
        id="task_id"
        v-model="form.task_id"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        <option value="">Nenhuma tarefa</option>
        <option v-for="task in clientTasks" :key="task.id" :value="task.id">
          {{ task.title }}
        </option>
      </select>
    </div>

    <div class="flex justify-end space-x-3">
      <button
        type="button"
        @click="$emit('cancel')"
        class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50"
      >
        Cancelar
      </button>
      <button
        type="submit"
        class="px-4 py-2 bg-primary text-white rounded-md text-sm font-medium hover:bg-primary-dark"
        :disabled="loading"
      >
        {{ loading ? 'Salvando...' : buttonText }}
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePaymentsStore } from '~/store/payments'
import { useClientsStore } from '~/store/clients'
import { useTasksStore } from '~/store/tasks'

const props = defineProps<{
  payment?: {
    id?: number
    amount: number
    description: string
    due_date: string
    status: string
    client_id: number
    task_id?: number
  }
  buttonText?: string
}>()

const emit = defineEmits<{
  (e: 'submit', payment: any): void
  (e: 'cancel'): void
}>()

const paymentsStore = usePaymentsStore()
const clientsStore = useClientsStore()
const tasksStore = useTasksStore()
const loading = ref(false)
const clients = ref([])
const clientTasks = ref([])

const form = ref({
  amount: props.payment?.amount || 0,
  description: props.payment?.description || '',
  due_date: props.payment?.due_date || '',
  status: props.payment?.status || 'pending',
  client_id: props.payment?.client_id || '',
  task_id: props.payment?.task_id || ''
})

const handleSubmit = async () => {
  loading.value = true
  try {
    const paymentData = {
      ...form.value,
      amount: Number(form.value.amount),
      task_id: form.value.task_id || undefined
    }

    if (props.payment?.id) {
      await paymentsStore.updatePayment(props.payment.id, paymentData)
    } else {
      await paymentsStore.createPayment(paymentData)
    }
    emit('submit', paymentData)
  } catch (error) {
    console.error('Erro ao salvar pagamento:', error)
  } finally {
    loading.value = false
  }
}

const loadClientTasks = async () => {
  if (!form.value.client_id) {
    clientTasks.value = []
    return
  }

  try {
    await tasksStore.fetchTasks(1, 100, undefined, form.value.client_id)
    clientTasks.value = tasksStore.getTasks
  } catch (error) {
    console.error('Erro ao carregar tarefas do cliente:', error)
    clientTasks.value = []
  }
}

onMounted(async () => {
  try {
    await clientsStore.fetchClients(1, 100)
    clients.value = clientsStore.getClients

    if (form.value.client_id) {
      await loadClientTasks()
    }
  } catch (error) {
    console.error('Erro ao carregar dados iniciais:', error)
  }
})
</script>
