import { defineNuxtPlugin } from 'nuxt/app'

export default defineNuxtPlugin((nuxtApp: any) => {
  // Adiciona cabeçalhos de segurança para todas as requisições
  if (process.client) {
    nuxtApp.hook('app:created', () => {
      // Intercepta todas as requisições para adicionar cabeçalhos de segurança
      const originalFetch = globalThis.fetch
      globalThis.fetch = async function(input, init) {
        if (!init) init = {}
        if (!init.headers) init.headers = {}

        // Previne CSRF adicionando token - só executa no cliente
        try {
          const csrfToken = localStorage.getItem('csrf-token')
          if (csrfToken) {
            Object.assign(init.headers, {
              'X-CSRF-Token': csrfToken
            })
          }
        } catch (error) {
          console.error('Erro ao acessar localStorage:', error)
        }

        // Cabeçalhos de segurança simplificados para desenvolvimento
        // Removidos cabeçalhos que podem interferir na comunicação

        return originalFetch(input, init)
      }
    })

    // Proteção contra clickjacking removida para desenvolvimento
  }

  // Sanitiza inputs para prevenir XSS
  nuxtApp.provide('sanitize', (input: string): string => {
    if (!input) return ''
    
    // Remove tags HTML e caracteres perigosos
    return input
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;')
      .replace(/'/g, '&#039;')
  })

  // Validação de dados
  nuxtApp.provide('validate', {
    // Validação de email
    email(email: string): boolean {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return re.test(email)
    },

    // Validação de senha forte
    password(password: string): { valid: boolean, message: string } {
      if (!password) {
        return { valid: false, message: 'Senha é obrigatória' }
      }

      if (password.length < 8) {
        return { valid: false, message: 'Senha deve ter pelo menos 8 caracteres' }
      }

      // Verifica se tem pelo menos uma letra maiúscula
      if (!/[A-Z]/.test(password)) {
        return { valid: false, message: 'Senha deve conter pelo menos uma letra maiúscula' }
      }

      // Verifica se tem pelo menos uma letra minúscula
      if (!/[a-z]/.test(password)) {
        return { valid: false, message: 'Senha deve conter pelo menos uma letra minúscula' }
      }

      // Verifica se tem pelo menos um número
      if (!/[0-9]/.test(password)) {
        return { valid: false, message: 'Senha deve conter pelo menos um número' }
      }

      return { valid: true, message: 'Senha válida' }
    },

    // Validação de dados numéricos
    number(value: any): boolean {
      return !isNaN(parseFloat(value)) && isFinite(value)
    }
  })
})
