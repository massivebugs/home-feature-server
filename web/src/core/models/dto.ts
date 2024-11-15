export type APIResponse<T> = {
  error: null | {
    code: string
    message: string
    validation_errors: { [key: string]: string }
  }
  data: T
}

export type APIError = {
  error: string
  validation_messages: { [key: string]: string }
}

export type CreateUserDto = {
  email: string
  username: string
  password: string
}

export type CreateAuthTokenDto = {
  username: string
  password: string
}

export type AuthTokenDto = string

export type UserDto = {
  id: number
  name: string
  logged_in_at: string // time
  created_at: string // time
}

export type GetUserDto = {
  user: UserDto
}

export type UserSystemPreferenceDto = {
  language: string | null
}

export type GetUserSystemPreferenceDto = {
  user_system_preference: UserSystemPreferenceDto
}
