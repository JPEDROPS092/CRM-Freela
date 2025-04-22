&lt;template>
  &lt;form @submit.prevent="handleSubmit" class="space-y-4">
    &lt;div>
      &lt;label for="title" class="block text-sm font-medium text-gray-700">Título&lt;/label>
      &lt;input
        type="text"
        id="title"
        v-model="form.title"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="description" class="block text-sm font-medium text-gray-700">Descrição&lt;/label>
      &lt;textarea
        id="description"
        v-model="form.description"
        rows="3"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >&lt;/textarea>
    &lt;/div>

    &lt;div>
      &lt;label for="due_date" class="block text-sm font-medium text-gray-700">Data de Entrega&lt;/label>
      &lt;input
        type="date"
        id="due_date"
        v-model="form.due_date"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="status" class="block text-sm font-medium text-gray-700">Status&lt;/label>
      &lt;select
        id="status"
        v-model="form.status"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        &lt;option value="pending">Pendente&lt;/option>
        &lt;option value="in_progress">Em Andamento&lt;/option>
        &lt;option value="completed">Concluída&lt;/option>
        &lt;option value="cancelled">Cancelada&lt;/option>
      &lt;/select>
    &lt;/div>

    &lt;div>
      &lt;label for="priority" class="block text-sm font-medium text-gray-700">Prioridade&lt;/label>
      &lt;select
        id="priority"
        v-model="form.priority"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        &lt;option value="low">Baixa&lt;/option>
        &lt;option value="medium">Média&lt;/option>
        &lt;option value="high">Alta&lt;/option>
      &lt;/select>
    &lt;/div>

    &lt;div>
      &lt;label for="client_id" class="block text-sm font-medium text-gray-700">Cliente&lt;/label>
      &lt;select
        id="client_id"
        v-model="form.client_id"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        &lt;option v-for="client in clients" :key="client.id" :value="client.id">
          {{ client.name }}
        &lt;/option>
      &lt;/select>
    &lt;/div>

    &lt;div class="flex justify-end space-x-3">
      &lt;button
        type="button"
        @click="$emit('cancel')"
        class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50"
      >
        Cancelar
      &lt;/button>
      &lt;button
        type="submit"
        class="px-4 py-2 bg-primary text-white rounded-md text-sm font-medium hover:bg-primary-dark"
        :disabled="loading"
      >
        {{ loading ? 'Salvando...' : buttonText }}
      &lt;/button>
    &lt;/div>
  &lt;/form>
&lt;/template>

&lt;script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTasksStore } from '~/store/tasks'
import { useClientsStore } from '~/store/clients'

const props = defineProps<{
  task?: {
    id?: number
    title: string
    description: string
    due_date: string
    status: string
    priority: string
    client_id: number
  }
  buttonText?: string
}>()

const emit = defineEmits<{
  (e: 'submit', task: any): void
  (e: 'cancel'): void
}>()

const tasksStore = useTasksStore()
const clientsStore = useClientsStore()
const loading = ref(false)
const clients = ref([])

const form = ref({
  title: props.task?.title || '',
  description: props.task?.description || '',
  due_date: props.task?.due_date || '',
  status: props.task?.status || 'pending',
  priority: props.task?.priority || 'medium',
  client_id: props.task?.client_id || 0
})

const handleSubmit = async () => {
  loading.value = true
  try {
    if (props.task?.id) {
      await tasksStore.updateTask(props.task.id, form.value)
    } else {
      await tasksStore.createTask(form.value)
    }
    emit('submit', form.value)
  } catch (error) {
    console.error('Erro ao salvar tarefa:', error)
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  try {
    await clientsStore.fetchClients(1, 100)
    clients.value = clientsStore.getClients
  } catch (error) {
    console.error('Erro ao carregar clientes:', error)
  }
})
&lt;/script>
