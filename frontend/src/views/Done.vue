<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute } from 'vue-router';
import TaskTable from '../components/TaskTable.vue';
import { todoList, getTodoList } from '../composables/useTask';

const route = useRoute();

onMounted(async () => {
    await getTodoList();
});
</script>

<template>
    <div class="container">
        <div class="header-row">
            <h1>完了したタスク</h1>
            <router-link to="/">戻る</router-link>
        </div>

        <TaskTable :todoList="todoList" :pathName="route.name" />
    </div>
</template>
