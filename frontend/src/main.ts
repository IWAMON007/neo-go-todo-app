// Vueアプリケーションのエントリーポイント
import { createApp } from 'vue'
import './assets/css/style.css' // グローバルCSSを適用
import App from './App.vue' // ルートコンポーネント
import router from './router/route'

// Homeをルートコンポーネントとしてアプリを生成し、#appにマウント
const app = createApp(App)
app.use(router)
app.mount('#app')
