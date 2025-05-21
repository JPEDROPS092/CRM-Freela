<template>
  <div>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">Dashboard</h1>
    </div>

    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <!-- Card de clientes -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Total de Clientes</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-gray-900">{{ stats.totalClients }}</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <NuxtLink to="/clients" class="font-medium text-primary-700 hover:text-primary-900">Ver todos</NuxtLink>
          </div>
        </div>
      </div>

      <!-- Card de tarefas -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Tarefas Pendentes</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-gray-900">{{ stats.pendingTasks }}</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <NuxtLink to="/tasks" class="font-medium text-primary-700 hover:text-primary-900">Ver todas</NuxtLink>
          </div>
        </div>
      </div>

      <!-- Card de pagamentos -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Pagamentos Pendentes</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-gray-900">R$ {{ formatCurrency(stats.pendingPayments) }}</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <NuxtLink to="/payments" class="font-medium text-primary-700 hover:text-primary-900">Ver todos</NuxtLink>
          </div>
        </div>
      </div>

      <!-- Card de receita -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Receita do Mês</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-gray-900">R$ {{ formatCurrency(stats.monthlyRevenue) }}</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <NuxtLink to="/reports" class="font-medium text-primary-700 hover:text-primary-900">Ver relatório</NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-8 grid grid-cols-1 gap-6 lg:grid-cols-2">
      <!-- Lista de tarefas recentes -->
      <div class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg leading-6 font-medium text-gray-900">Tarefas Recentes</h3>
          <div class="mt-5">
            <TaskList 
              :tasks="recentTasks" 
              :show-pagination="false" 
              :show-add-button="false" 
            />
          </div>
        </div>
      </div>

      <!-- Lista de pagamentos recentes -->
      <div class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg leading-6 font-medium text-gray-900">Pagamentos Recentes</h3>
          <div class="mt-5">
            <PaymentList 
              :payments="recentPayments" 
              :show-pagination="false" 
              :show-add-button="false" 
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTasksStore } from '~/store/tasks'
import { usePaymentsStore } from '~/store/payments'
import { useClientsStore } from '~/store/clients'

// Define o título da página
useHead({
  title: 'Dashboard - CRM Freelancer'
})

const tasksStore = useTasksStore()
const paymentsStore = usePaymentsStore()
const clientsStore = useClientsStore()

const stats = ref({
  totalClients: 0,
  pendingTasks: 0,
  pendingPayments: 0,
  monthlyRevenue: 0
})

const recentTasks = ref([])
const recentPayments = ref([])

const formatCurrency = (value: number) => {
  return value.toLocaleString('pt-BR', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  })
}

onMounted(async () => {
  try {
    // Carrega estatísticas
    const [clientStats, taskStats, paymentStats] = await Promise.all([
      clientsStore.fetchStats(),
      tasksStore.fetchStats(),
      paymentsStore.fetchStats()
    ])

    stats.value = {
      totalClients: clientStats.total,
      pendingTasks: taskStats.pending,
      pendingPayments: paymentStats.pending,
      monthlyRevenue: paymentStats.monthlyRevenue
    }

    // Carrega tarefas e pagamentos recentes
    const [tasks, payments] = await Promise.all([
      tasksStore.fetchTasks(1, 5),
      paymentsStore.fetchPayments(1, 5)
    ])

    recentTasks.value = tasks
    recentPayments.value = payments
  } catch (error) {
    console.error('Erro ao carregar dados do dashboard:', error)
  }
})
</script>
