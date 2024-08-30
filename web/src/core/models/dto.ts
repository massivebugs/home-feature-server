export type APIResponse<T> = {
  error: null | {
    code: string
    message: string
    validation_errors: { [key: string]: string }
  }
  data: T
}

export type CreateAuthTokenDto = {
  username: string
  password: string
}
export type AuthTokenDto = string
export type AuthTokenResponse = APIResponse<AuthTokenDto>

export type AuthUserDto = {
  id: number
  name: string
  logged_in_at: string // time
  created_at: string // time
}
export type AuthUserResponse = APIResponse<AuthUserDto>
