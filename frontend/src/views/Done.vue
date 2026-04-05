<script setup lang="ts">
// import { watch } from 'vue';
import { useRoute } from 'vue-router'
import { useGetTodoList } from '../composables/useTask';
import TaskTable from '../components/TaskTable.vue'

const route = useRoute()
const { todoList } = useGetTodoList()

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