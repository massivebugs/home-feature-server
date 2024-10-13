export const DateTimeFormat = 'YYYY-MM-DDTHH:mm'
export const DateFormat = 'YYYY-MM-DD'

export const sleep = (ms: number): Promise<void> => {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
