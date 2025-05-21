// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  devtools: { enabled: true },

  // Application metadata
  app: {
    head: {
      title: 'CRM Freelancer',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'CRM para freelancers gerenciarem clientes, tarefas e pagamentos' },
        // Previne MIME sniffing
        { 'http-equiv': 'X-Content-Type-Options', content: 'nosniff' },
        // Política de referenciamento
        { name: 'referrer', content: 'strict-origin-when-cross-origin' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap' }
      ]
    }
  },

  // CSS
  css: [
    '~/assets/css/main.css'
  ],

  // Modules
  modules: [
    '@pinia/nuxt',
    '@nuxtjs/tailwindcss',
  ],

  // Runtime config
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE || 'http://localhost:8080/api'
    }
  },

  // TypeScript
  typescript: {
    strict: true
  },

  // Configurações de segurança
  nitro: {
    routeRules: {
      '/**': {
        headers: {
          // CSP removido para facilitar o desenvolvimento
        },
        cors: true
      }
    }
  },

  compatibilityDate: '2025-04-22'
})