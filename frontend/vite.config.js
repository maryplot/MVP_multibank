import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Для auth-service
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
      // Для accounts-service
      '/accounts': {
        target: 'http://localhost:8081',
        changeOrigin: true
      },
      '/balance': {
        target: 'http://localhost:8081',
        changeOrigin: true
      },
      // Для transfer-service
      '/transfer': {
        target: 'http://localhost:8082',
        changeOrigin: true
      }
    }
  }
})