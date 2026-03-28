// Vueアプリケーションのエントリーポイント
import { createApp } from 'vue'
import './assets/css/style.css' // グローバルCSSを適用
import Home from './views/Home.vue' // ルートコンポーネント

// Homeをルートコンポーネントとしてアプリを生成し、#appにマウント
createApp(Home).mount('#app')
