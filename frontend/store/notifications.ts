import { defineStore } from 'pinia'

interface Notification {
  id: string
  type: 'info' | 'success' | 'error' | 'warning'
  title?: string
  message: string
  duration?: number
  position?: 'justify-end' | 'justify-center' | 'justify-start'
}

interface NotificationsState {
  notifications: Notification[]
}

export const useNotificationsStore = defineStore('notifications', {
  state: (): NotificationsState => ({
    notifications: []
  }),

  actions: {
    showNotification(notification: Omit<Notification, 'id'>) {
      const id = Date.now().toString()
      this.notifications.push({
        id,
        ...notification
      })

      // Remove a notificação após o tempo definido
      if (notification.duration !== 0) {
        setTimeout(() => {
          this.removeNotification(id)
        }, notification.duration || 5000)
      }

      return id
    },

    showSuccess(message: string, title?: string, options = {}) {
      return this.showNotification({
        type: 'success',
        message,
        title,
        ...options
      })
    },

    showError(message: string, title?: string, options = {}) {
      return this.showNotification({
        type: 'error',
        message,
        title: title || 'Erro',
        ...options
      })
    },

    showWarning(message: string, title?: string, options = {}) {
      return this.showNotification({
        type: 'warning',
        message,
        title: title || 'Atenção',
        ...options
      })
    },

    showInfo(message: string, title?: string, options = {}) {
      return this.showNotification({
        type: 'info',
        message,
        title,
        ...options
      })
    },

    removeNotification(id: string) {
      const index = this.notifications.findIndex(n => n.id === id)
      if (index !== -1) {
        this.notifications.splice(index, 1)
      }
    },

    clearAll() {
      this.notifications = []
    }
  }
})
