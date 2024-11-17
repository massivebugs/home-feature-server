export type APIResponse<T> = {
  error: null | {
    code: string
    message: string
    validation_errors: { [key: string]: string }
  }
  data: T
}
