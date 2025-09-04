import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    // host: '192.168.0.102', // ← ЗАКОММЕНТИРОВАТЬ для локальной разработки 
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
      '/accounts': {
        target: 'http://localhost:8081',
        changeOrigin: true
      },
      '/balance': {
        target: 'http://localhost:8081',
        changeOrigin: true
      },
      '/transfer': {
        target: 'http://localhost:8082',
        changeOrigin: true
      }
    }
  }
})