<script setup lang="ts">
import { onMounted } from 'vue';
import { useGetTodoList } from '../composables/useTask';
import FromTask from '../components/FromTask.vue'
import TaskTable from '../components/TaskTable.vue'

// getTodoList も一緒に受け取る
const { todoList, getTodoList } = useGetTodoList()

// マウント時にデータを取得する処理を実行
onMounted(() => {
    getTodoList()
})
</script>

<template>
<div class="container">

    <div class="header-row">
        <h1>TODOアプリ</h1>
        <a href="/done/list" class="done-link">完了したタスク</a>
    </div>

    <FromTask />

    <TaskTable :todoList="todoList"/>
</div>
</template>