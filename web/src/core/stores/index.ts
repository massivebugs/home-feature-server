import { defineStore } from 'pinia'
import { ref, shallowRef, triggerRef } from 'vue'
import {
  type APIResponse,
  type AuthTokenDto,
  type AuthUserDto,
  type CreateAuthTokenDto,
  type CreateUserDto,
  type UserSystemPreferenceDto,
} from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import { AuthUser } from '../models/auth_user'
import { Process } from '../models/process'
import type { Program } from '../models/program'

export const useCoreStore = defineStore('core', () => {
  const authUser = ref<AuthUser | null>(null)
  const preferences = ref<UserSystemPreferenceDto>({ language: null })
  const processes = shallowRef<Map<string, Process>>(new Map())
  const programs = shallowRef<Map<string, Program>>(new Map())
  const topLevelProcessId = ref<string | null>(null)
  const processesByInsertOrder = shallowRef<Map<string, Process>>(new Map())

  const createUser = (data: CreateUserDto) => api.post(APIEndpoints.v1.auth.default, data)

  const getAuthUser = () => api.get<AuthUserDto>(APIEndpoints.v1.secure.user.default)

  const createAuthToken = (data: CreateAuthTokenDto) =>
    api.post<AuthTokenDto>(APIEndpoints.v1.auth.token, data)

  const createRefreshToken = () => api.post<AuthTokenDto>(APIEndpoints.v1.secure.auth.token)

  const getUserSystemPreferences = () =>
    api.get<APIResponse<UserSystemPreferenceDto>>(APIEndpoints.v1.secure.systemPreferences.default)

  const createUserSystemPreferences = () =>
    api.post<APIResponse<UserSystemPreferenceDto>>(APIEndpoints.v1.secure.systemPreferences.default)

  const updateUserSystemPreferences = () =>
    api.put<APIResponse<null>>(APIEndpoints.v1.secure.systemPreferences.default, preferences.value)

  const addProcess = (process: Process) => {
    processes.value.set(process.id, process)
    processesByInsertOrder.value.set(process.id, process)
    topLevelProcessId.value = process.id
    triggerRef(processes)
    triggerRef(processesByInsertOrder)
  }

  const removeProcess = (processId: string) => {
    processes.value.delete(processId)
    processesByInsertOrder.value.delete(processId)
    if (processId === topLevelProcessId.value) {
      topLevelProcessId.value = null
    }
    triggerRef(processes)
    triggerRef(processesByInsertOrder)
  }

  const setTopLevelProcess = (processId: string) => {
    const process = processes.value.get(processId)
    if (process) {
      processes.value.delete(processId)
      processes.value.set(process.id, process)
    }
    topLevelProcessId.value = processId
    triggerRef(processes)
  }

  const addProgram = (program: Program) => {
    programs.value.set(program.id, program)
    triggerRef(programs)
  }

  const findProgramProcesses = (programId: string): Process[] => {
    return [...processes.value.values()].filter((p) => {
      return p.program.id === programId
    })
  }

  return {
    authUser,
    preferences,
    processes,
    programs,
    topLevelProcessId,
    processesByInsertOrder,
    createUser,
    getAuthUser,
    createAuthToken,
    createRefreshToken,
    getUserSystemPreferences,
    createUserSystemPreferences,
    updateUserSystemPreferences,
    addProcess,
    removeProcess,
    setTopLevelProcess,
    addProgram,
    findProgramProcesses,
  }
})
