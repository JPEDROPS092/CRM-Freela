// Declarações de módulos para resolver problemas de importação
declare module '#app' {
  import type { Ref } from 'vue'
  
  export const defineNuxtRouteMiddleware: any
  export const navigateTo: any
  export const useRuntimeConfig: () => {
    public: {
      apiBase: string
      [key: string]: any
    }
    [key: string]: any
  }
  export const useHead: (head: any) => void
}

// Declaração para o módulo jwt-decode
declare module 'jwt-decode' {
  export function jwtDecode<T = any>(token: string): T
}

// Declaração para os stores
declare module '~/store/auth' {
  export const useAuthStore: any
}

declare module '~/store/clients' {
  export const useClientsStore: any
}

declare module '~/store/tasks' {
  export const useTasksStore: any
}

declare module '~/store/payments' {
  export const usePaymentsStore: any
}

declare module '~/store/notifications' {
  export const useNotificationsStore: any
}

// Declaração para componentes
declare module 'vue' {
  interface GlobalComponents {
    NotificationsContainer: any
    Modal: any
    ConfirmationModal: any
    Notification: any
  }
}

// Extensão do namespace NodeJS
declare namespace NodeJS {
  interface Timeout {}
}
