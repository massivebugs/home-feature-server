import { defineStore } from 'pinia'
import { shallowRef, triggerRef } from 'vue'
import { Process } from '../models/process'

export const useStore = defineStore('core', () => {
  const processes = shallowRef<Map<string, Process>>(new Map())

  const removeProcess = (processId: string) => {
    processes.value.delete(processId)
    triggerRef(processes)
  }

  return { processes, removeProcess }
})
