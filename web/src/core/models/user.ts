import type { UserResponse } from '../composables/useAPI'

export class User {
  id: number
  name: string
  loggedInAt: Date
  createdAt: Date
  constructor(user: UserResponse) {
    this.id = user.id
    this.name = user.name
    this.loggedInAt = new Date(user.loggedInAt)
    this.createdAt = new Date(user.createdAt)
  }
}
