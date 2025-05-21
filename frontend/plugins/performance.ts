import { defineNuxtPlugin } from 'nuxt/app'

export default defineNuxtPlugin((nuxtApp: any) => {
  // Implementa lazy loading para componentes grandes
  if (process.client) {
    nuxtApp.vueApp.directive('lazy-load', {
      mounted(el: HTMLElement, binding: any) {
        const options = {
          root: null,
          rootMargin: '0px',
          threshold: 0.1
        }

        const observer = new IntersectionObserver((entries) => {
          entries.forEach(entry => {
            if (entry.isIntersecting) {
              if (typeof binding.value === 'function') {
                binding.value()
              }
              observer.unobserve(el)
            }
          })
        }, options)

        observer.observe(el)
      }
    })
  }

  // Implementa debounce para otimizar chamadas de API
  nuxtApp.provide('debounce', (fn: Function, delay: number) => {
    let timeout: NodeJS.Timeout
    return (...args: any[]) => {
      clearTimeout(timeout)
      timeout = setTimeout(() => fn(...args), delay)
    }
  })

  // Implementa throttle para limitar chamadas de API
  nuxtApp.provide('throttle', (fn: Function, limit: number) => {
    let inThrottle: boolean
    return (...args: any[]) => {
      if (!inThrottle) {
        fn(...args)
        inThrottle = true
        setTimeout(() => (inThrottle = false), limit)
      }
    }
  })
})
