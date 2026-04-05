<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute } from 'vue-router'
import { useGetTodoList } from '../composables/useTask';
import FromTask from '../components/FormTask.vue'
import TaskTable from '../components/TaskTable.vue'

// getTodoList も一緒に受け取る
const { todoList, getTodoList } = useGetTodoList()

const route = useRoute()

// マウント時にデータを取得する処理を実行
onMounted(() => {
    getTodoList()
})
</script>

<template>
<div class="container">

    <div class="header-row">
        <h1>TODOアプリ</h1>
        <router-link to="/Done">完了したタスク</router-link>
    </div>

    <FromTask />

    <TaskTable :todoList="todoList" :pathName="route.name"/>
</div>
</template>