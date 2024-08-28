export const DateTimeFormat = 'YYYY-MM-DDTHH:mm'

export const sleep = (ms: number): Promise<void> => {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
