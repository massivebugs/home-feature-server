import type { defineStore } from 'pinia'

// Get the type of a store's instance
// source: https://github.com/vuejs/pinia/discussions/1054
export type PiniaStore<T extends (...args: any) => any> = Omit<
  ReturnType<T>,
  keyof ReturnType<typeof defineStore>
>
