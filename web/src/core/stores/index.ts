import { defineStore } from 'pinia'
import { ref, shallowRef, triggerRef } from 'vue'
import {
  type AuthTokenResponse,
  type AuthUserResponse,
  type CreateAuthTokenDto,
} from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import { AuthUser } from '../models/auth_user'
import { Process } from '../models/process'
import type { Program } from '../models/program'

export const useCoreStore = defineStore('core', () => {
  const authUser = ref<AuthUser | null>(null)
  const processes = shallowRef<Map<string, Process>>(new Map())
  const programs = shallowRef<Map<string, Program>>(new Map())
  const topLevelProcessId = ref<string | null>(null)
  const processesByInsertOrder = shallowRef<Map<string, Process>>(new Map())

  const getAuthUser = () => api.get<AuthUserResponse>(APIEndpoints.v1.secure.authUser)

  const getAuthToken = (data: CreateAuthTokenDto) =>
    api.post<AuthTokenResponse>(APIEndpoints.v1.authToken, data)

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
    processes,
    programs,
    topLevelProcessId,
    processesByInsertOrder,
    getAuthUser,
    getAuthToken,
    addProcess,
    removeProcess,
    setTopLevelProcess,
    addProgram,
    findProgramProcesses,
  }
})
