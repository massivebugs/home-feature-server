import vue from '@vitejs/plugin-vue'
import fs from 'fs'
import { URL, fileURLToPath } from 'node:url'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    https: {
      cert: fs.readFileSync('/devcerts/localhost.pem'),
      key: fs.readFileSync('/devcerts/localhost-key.pem'),
    },
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
