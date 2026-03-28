<script setup lang="ts">
import FromTask from '../components/FromTask.vue'
</script>

<template>
<div class="container">

    <div class="header-row">
        <h1>TODOアプリ</h1>
        <a href="/done/list" class="done-link">完了したタスク</a>
    </div>

    <FromTask />
    
    <table>
        <thead>
            <tr>
                <th class="id-cell">ID</th>
                <th>タスク</th>
                <th></th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <th>1</th>
                <th>選択</th>
                <th>編集</th>
                <th>完了</th>
            </tr>
        </tbody>
    </table>
</div>
</template>