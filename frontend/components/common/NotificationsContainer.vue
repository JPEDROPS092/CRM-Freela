<template>
  <div class="fixed top-0 right-0 p-4 z-50">
    <div
      v-for="notification in notifications"
      :key="notification.id"
      :class="[
        'mb-2 p-4 rounded-md shadow-md transition-all duration-300 transform',
        'flex items-center',
        notification.type === 'success' ? 'bg-green-100 text-green-800 border-l-4 border-green-500' :
        notification.type === 'error' ? 'bg-red-100 text-red-800 border-l-4 border-red-500' :
        notification.type === 'warning' ? 'bg-yellow-100 text-yellow-800 border-l-4 border-yellow-500' :
        'bg-blue-100 text-blue-800 border-l-4 border-blue-500',
        { 'translate-x-0 opacity-100': true, 'translate-x-full opacity-0': false }
      ]"
    >
      <div class="mr-3">
        <svg
          v-if="notification.type === 'success'"
          class="h-5 w-5 text-green-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M5 13l4 4L19 7"
          />
        </svg>
        <svg
          v-else-if="notification.type === 'error'"
          class="h-5 w-5 text-red-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
        <svg
          v-else-if="notification.type === 'warning'"
          class="h-5 w-5 text-yellow-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
          />
        </svg>
        <svg
          v-else
          class="h-5 w-5 text-blue-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
      </div>
      <div class="flex-1">
        <div class="font-medium">{{ notification.title }}</div>
        <div class="text-sm">{{ notification.message }}</div>
      </div>
      <button
        @click="removeNotification(notification.id)"
        class="ml-4 text-gray-500 hover:text-gray-700 focus:outline-none"
      >
        <svg
          class="h-4 w-4"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useNotificationsStore } from '~/store/notifications'

const notificationsStore = useNotificationsStore()
const notifications = computed(() => notificationsStore.notifications)

const removeNotification = (id) => {
  notificationsStore.removeNotification(id)
}
</script>
