<script setup lang="ts">
import { onMounted } from 'vue';
import { todoList, getTodoList } from '../composables/useTask';
import DoneTaskTable from '../components/DoneTaskTable.vue';

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

        <DoneTaskTable :todoList="todoList" />
    </div>
</template>
