<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import TaskTable from '../components/TaskTable.vue'
import { getTodoList } from '../composables/useTask';
import type { Todo } from '../types/todo'

const route = useRoute()
getTodoList()
const todoList = ref<Todo[]>([])

// watch(
//     () => route.fullPath, 
//     () => {
//         getDoneList()
//     }
// )

</script>

<template>
<div class="container">

    <div class="header-row">
        <h1>完了したタスク</h1>
        <router-link to="/">戻る</router-link>
    </div>

    <TaskTable :todoList="todoList" :pathName="route.name"/>
</div>
</template>