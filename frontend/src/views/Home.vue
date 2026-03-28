<script setup lang="ts">
// タスク追加フォームコンポーネント
import FromTask from '../components/FromTask.vue'
// タスク一覧テーブルコンポーネント
import TaskTable from '../components/TaskTable.vue'
</script>

<template>
<div class="container">

    <div class="header-row">
        <h1>TODOアプリ</h1>
        <a href="/done/list" class="done-link">完了したタスク</a>
    </div>

    <FromTask />

    <TaskTable />
</div>
</template>