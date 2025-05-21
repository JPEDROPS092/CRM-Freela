<template>
  <div>
    <form @submit.prevent="submitForm" class="space-y-6">
      <div>
        <label for="title" class="block text-sm font-medium text-gray-700">Título</label>
        <input
          type="text"
          id="title"
          v-model="form.title"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
          required
        />
      </div>

      <div>
        <label for="description" class="block text-sm font-medium text-gray-700">Descrição</label>
        <textarea
          id="description"
          v-model="form.description"
          rows="3"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
        ></textarea>
      </div>

      <div class="grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-2">
        <div>
          <label for="client" class="block text-sm font-medium text-gray-700">Cliente</label>
          <select
            id="client"
            v-model="form.client_id"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
          >
            <option value="">Selecione um cliente</option>
            <option v-for="client in clients" :key="client.id" :value="client.id">
              {{ client.name }}
            </option>
          </select>
        </div>

        <div>
          <label for="due_date" class="block text-sm font-medium text-gray-700">Data de Entrega</label>
          <input
            type="date"
            id="due_date"
            v-model="form.due_date"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-2">
        <div>
          <label for="status" class="block text-sm font-medium text-gray-700">Status</label>
          <select
            id="status"
            v-model="form.status"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
          >
            <option value="pending">Pendente</option>
            <option value="in_progress">Em Andamento</option>
            <option value="completed">Concluída</option>
            <option value="cancelled">Cancelada</option>
          </select>
        </div>

        <div>
          <label for="priority" class="block text-sm font-medium text-gray-700">Prioridade</label>
          <select
            id="priority"
            v-model="form.priority"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
          >
            <option value="low">Baixa</option>
            <option value="medium">Média</option>
            <option value="high">Alta</option>
          </select>
        </div>
      </div>

      <div>
        <label for="estimated_hours" class="block text-sm font-medium text-gray-700">Horas Estimadas</label>
        <input
          type="number"
          id="estimated_hours"
          v-model="form.estimated_hours"
          min="0"
          step="0.5"
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary sm:text-sm"
        />
      </div>

      <div class="flex justify-end space-x-3">
        <button
          type="button"
          @click="$emit('cancel')"
          class="inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary"
        >
          Cancelar
        </button>
        <button
          type="submit"
          class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary"
        >
          {{ task ? 'Atualizar' : 'Criar' }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useTasksStore } from '~/store/tasks'
import { useClientsStore } from '~/store/clients'

const props = defineProps({
  task: {
    type: Object,
    default: null
  },
  buttonText: {
    type: String,
    default: 'Criar'
  }
})

const emit = defineEmits(['submit', 'cancel'])

const tasksStore = useTasksStore()
const clientsStore = useClientsStore()
const loading = ref(false)
const clients = ref([])

// Formulário
const form = ref({
  title: '',
  description: '',
  client_id: '',
  due_date: '',
  status: 'pending',
  priority: 'medium',
  estimated_hours: 0
})

// Preencher o formulário se estiver editando uma tarefa existente
onMounted(async () => {
  if (props.task) {
    form.value = {
      title: props.task.title || '',
      description: props.task.description || '',
      client_id: props.task.client_id || '',
      due_date: props.task.due_date || '',
      status: props.task.status || 'pending',
      priority: props.task.priority || 'medium',
      estimated_hours: props.task.estimated_hours || 0
    }
  }

  try {
    await clientsStore.fetchClients(1, 100)
    clients.value = clientsStore.getClients
  } catch (error) {
    console.error('Erro ao carregar clientes:', error)
  }
})

// Enviar o formulário
const submitForm = async () => {
  loading.value = true
  try {
    if (props.task) {
      await tasksStore.updateTask(props.task.id, form.value)
    } else {
      await tasksStore.createTask(form.value)
    }
    emit('submit', { ...form.value })
  } catch (error) {
    console.error('Erro ao salvar tarefa:', error)
  } finally {
    loading.value = false
  }
}
</script>
