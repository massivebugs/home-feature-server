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

  const getAuthUser = () => api.get<AuthUserResponse>(APIEndpoints.v1.secure.authUser)

  const getAuthToken = (data: CreateAuthTokenDto) =>
    api.post<AuthTokenResponse>(APIEndpoints.v1.authToken, data)

  const addProcess = (process: Process) => {
    processes.value.set(process.id, process)
    triggerRef(processes)
  }

  const removeProcess = (processId: string) => {
    processes.value.delete(processId)
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
    getAuthUser,
    getAuthToken,
    addProcess,
    removeProcess,
    addProgram,
    findProgramProcesses,
  }
})
