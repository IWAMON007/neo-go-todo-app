<script setup lang="ts">
import { onErrorCaptured, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { todoList, getTodoList } from '../composables/useTask';
import FromTask from '../components/FormTask.vue';
import TaskTable from '../components/TaskTable.vue';

const route = useRoute();

// マウント時にデータを取得する処理を実行
onMounted(() => {
    getTodoList();
});

onErrorCaptured((error) => {
    alert(error.message);
    return false;
});
</script>

<template>
    <div class="container">
        <div class="header-row">
            <h1>TODOアプリ</h1>
            <router-link to="/Done">完了したタスク</router-link>
        </div>

        <FromTask />

        <TaskTable :todoList="todoList" :pathName="route.name" />
    </div>
</template>
