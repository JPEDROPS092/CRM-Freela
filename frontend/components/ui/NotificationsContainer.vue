<template>
  <div>
    <Notification
      v-for="notification in notifications"
      :key="notification.id"
      :type="notification.type"
      :title="notification.title"
      :message="notification.message"
      :duration="notification.duration"
      :position="notification.position"
      @close="removeNotification(notification.id)"
    />
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useNotificationsStore } from '~/store/notifications'

const notificationsStore = useNotificationsStore()
const { notifications } = storeToRefs(notificationsStore)
const { removeNotification } = notificationsStore
</script>
