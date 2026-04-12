import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      // /task, /done など Go のエンドポイントをすべて転送
      '/task': {
        target: 'http://go-app:8080',
        changeOrigin: true,
      },
      '/done': {
        target: 'http://go-app:8080',
        changeOrigin: true,
      },
      '/todo': {
        target: 'http://go-app:8080',
        changeOrigin: true,
      },
    },
  },
})
