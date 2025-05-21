<template>
  <Modal :title="title" @close="$emit('cancel')">
    <div class="p-6">
      <p class="text-gray-700 mb-6">
        {{ message }}
      </p>
      <div class="flex justify-end space-x-3">
        <button 
          @click="$emit('cancel')" 
          class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          {{ cancelText }}
        </button>
        <button 
          @click="$emit('confirm')" 
          :class="[
            'px-4 py-2 text-white rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2',
            dangerMode 
              ? 'bg-red-600 hover:bg-red-700 focus:ring-red-500' 
              : 'bg-primary-600 hover:bg-primary-700 focus:ring-primary-500'
          ]"
        >
          {{ confirmText }}
        </button>
      </div>
    </div>
  </Modal>
</template>

<script setup lang="ts">
defineProps({
  title: {
    type: String,
    default: 'Confirmação'
  },
  message: {
    type: String,
    required: true
  },
  confirmText: {
    type: String,
    default: 'Confirmar'
  },
  cancelText: {
    type: String,
    default: 'Cancelar'
  },
  dangerMode: {
    type: Boolean,
    default: false
  }
})

defineEmits(['confirm', 'cancel'])
</script>
