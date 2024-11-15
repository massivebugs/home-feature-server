import type { UserDto } from './dto'

export class User {
  id: number
  name: string
  loggedInAt: Date
  createdAt: Date
  constructor(userDto: UserDto) {
    this.id = userDto.id
    this.name = userDto.name
    this.loggedInAt = new Date(userDto.logged_in_at)
    this.createdAt = new Date(userDto.created_at)
  }
}
