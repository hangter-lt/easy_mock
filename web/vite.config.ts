import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/manages': {
        target: 'http://127.0.0.1:7001',
        changeOrigin: true,
      },
      '/request': {
        target: 'http://127.0.0.1:7001',
        changeOrigin: true,
      }
    },
  },
})
