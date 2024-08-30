export type ValidationErrors<K> = { [k in keyof K]: string }
