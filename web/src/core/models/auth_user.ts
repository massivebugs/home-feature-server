import type { AuthUserDto } from './dto'

export class AuthUser {
  id: number
  name: string
  loggedInAt: Date
  createdAt: Date
  constructor(authUserDto: AuthUserDto) {
    this.id = authUserDto.id
    this.name = authUserDto.name
    this.loggedInAt = new Date(authUserDto.logged_in_at)
    this.createdAt = new Date(authUserDto.created_at)
  }
}
