// router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Done from '../views/Done.vue'

const routes = [
    { 
        path: '/',
        name: 'Home',
        component: Home 
    },
    { 
        path: '/Done', 
        name: 'Done', 
        component: Done 
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
