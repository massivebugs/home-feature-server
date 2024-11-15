import { defineStore } from 'pinia'
import { ref, shallowRef, triggerRef } from 'vue'
import {
  type AuthTokenDto,
  type CreateAuthTokenDto,
  type CreateUserDto,
  type GetUserDto,
  type GetUserSystemPreferenceDto,
  type UserSystemPreferenceDto,
} from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import { Process } from '../models/process'
import type { Program } from '../models/program'
import { User } from '../models/user'

export const useCoreStore = defineStore('core', () => {
  const user = ref<User | null>(null)
  const systemPreference = ref<UserSystemPreferenceDto>({ language: null })
  const processes = shallowRef<Map<string, Process>>(new Map())
  const programs = shallowRef<Map<string, Program>>(new Map())
  const topLevelProcessId = ref<string | null>(null)
  const processesByInsertOrder = shallowRef<Map<string, Process>>(new Map())

  const createUser = (data: CreateUserDto) => api.post(APIEndpoints.v1.auth.default, data)

  const getUser = () => api.get<GetUserDto>(APIEndpoints.v1.secure.user.default)

  const createAuthToken = (data: CreateAuthTokenDto) =>
    api.post<AuthTokenDto>(APIEndpoints.v1.auth.token, data)

  const createRefreshToken = () => api.post<AuthTokenDto>(APIEndpoints.v1.secure.auth.token)

  const getUserSystemPreference = () =>
    api.get<GetUserSystemPreferenceDto>(APIEndpoints.v1.secure.systemPreferences.default)

  const createUserSystemPreference = () =>
    api.post<GetUserSystemPreferenceDto>(APIEndpoints.v1.secure.systemPreferences.default)

  const updateUserSystemPreference = () =>
    api.put(APIEndpoints.v1.secure.systemPreferences.default, systemPreference.value)

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
    user,
    systemPreference,
    processes,
    programs,
    topLevelProcessId,
    processesByInsertOrder,
    createUser,
    getUser,
    createAuthToken,
    createRefreshToken,
    getUserSystemPreference,
    createUserSystemPreference,
    updateUserSystemPreference,
    addProcess,
    removeProcess,
    setTopLevelProcess,
    addProgram,
    findProgramProcesses,
  }
})
