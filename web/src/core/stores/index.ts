import { defineStore } from 'pinia'
import { ref, shallowRef, triggerRef } from 'vue'
import { Process } from '../models/process'
import { api, APIEndpoints } from '@/utils/api'
import { type AuthUserResponse, type AuthTokenResponse } from '@/modules/budget_planner/models/dto'
import { AuthUser } from '../models/auth_user'

export const useStore = defineStore('core', () => {
  const authUser = ref<AuthUser | null>(null)
  const processes = shallowRef<Map<string, Process>>(new Map())

  const removeProcess = (processId: string) => {
    processes.value.delete(processId)
    triggerRef(processes)
  }

  const getAuthUser = () => api.get<AuthUserResponse>(APIEndpoints.v1.authUser)

  const getAuthToken = (username: string, password: string) =>
    api.post<AuthTokenResponse>(APIEndpoints.v1.authToken, { username, password })

  return { authUser, processes, removeProcess, getAuthUser, getAuthToken }
})
