/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_PORT: number
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

declare const APP_VERSION: string
