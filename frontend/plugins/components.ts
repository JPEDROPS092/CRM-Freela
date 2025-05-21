// Plugin para registrar componentes globais
import { defineNuxtPlugin } from '#app'

// Importação dos componentes comuns
import NotificationsContainer from '~/components/common/NotificationsContainer.vue'
import ConfirmationModal from '~/components/common/ConfirmationModal.vue'

// Importação dos componentes de pagamentos
import PaymentList from '~/components/payments/PaymentList.vue'
import PaymentForm from '~/components/payments/PaymentForm.vue'

// Importação dos componentes de tarefas
import TaskList from '~/components/tasks/TaskList.vue'
import TaskForm from '~/components/tasks/TaskForm.vue'

// Importação dos componentes de clientes
import ClientList from '~/components/clients/ClientList.vue'
import ClientForm from '~/components/clients/ClientForm.vue'

export default defineNuxtPlugin((nuxtApp) => {
  // No Nuxt 3, não é necessário registrar componentes manualmente
  // Os componentes são auto-importados desde que estejam no diretório components
  // Este plugin está aqui apenas para referência e pode ser removido
  
  // Se você precisar registrar componentes manualmente, pode usar:
  // nuxtApp.vueApp.component('ComponentName', ComponentDefinition)
  
  // Mas o recomendado é deixar o Nuxt fazer isso automaticamente
})
