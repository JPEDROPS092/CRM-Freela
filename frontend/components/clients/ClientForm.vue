&lt;template>
  &lt;form @submit.prevent="handleSubmit" class="space-y-4">
    &lt;div>
      &lt;label for="name" class="block text-sm font-medium text-gray-700">Nome&lt;/label>
      &lt;input
        type="text"
        id="name"
        v-model="form.name"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="email" class="block text-sm font-medium text-gray-700">Email&lt;/label>
      &lt;input
        type="email"
        id="email"
        v-model="form.email"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="phone" class="block text-sm font-medium text-gray-700">Telefone&lt;/label>
      &lt;input
        type="tel"
        id="phone"
        v-model="form.phone"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="company" class="block text-sm font-medium text-gray-700">Empresa&lt;/label>
      &lt;input
        type="text"
        id="company"
        v-model="form.company"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      />
    &lt;/div>

    &lt;div>
      &lt;label for="notes" class="block text-sm font-medium text-gray-700">Observações&lt;/label>
      &lt;textarea
        id="notes"
        v-model="form.notes"
        rows="3"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >&lt;/textarea>
    &lt;/div>

    &lt;div>
      &lt;label for="status" class="block text-sm font-medium text-gray-700">Status&lt;/label>
      &lt;select
        id="status"
        v-model="form.status"
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary"
      >
        &lt;option value="active">Ativo&lt;/option>
        &lt;option value="inactive">Inativo&lt;/option>
        &lt;option value="archived">Arquivado&lt;/option>
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
import { ref } from 'vue'
import { useClientsStore } from '~/store/clients'

const props = defineProps<{
  client?: {
    id?: number
    name: string
    email: string
    phone: string
    company: string
    notes: string
    status: string
  }
  buttonText?: string
}>()

const emit = defineEmits<{
  (e: 'submit', client: any): void
  (e: 'cancel'): void
}>()

const clientsStore = useClientsStore()
const loading = ref(false)

const form = ref({
  name: props.client?.name || '',
  email: props.client?.email || '',
  phone: props.client?.phone || '',
  company: props.client?.company || '',
  notes: props.client?.notes || '',
  status: props.client?.status || 'active'
})

const handleSubmit = async () => {
  loading.value = true
  try {
    if (props.client?.id) {
      await clientsStore.updateClient(props.client.id, form.value)
    } else {
      await clientsStore.createClient(form.value)
    }
    emit('submit', form.value)
  } catch (error) {
    console.error('Erro ao salvar cliente:', error)
  } finally {
    loading.value = false
  }
}
&lt;/script>
