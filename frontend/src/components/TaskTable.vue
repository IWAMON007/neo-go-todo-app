<script setup lang="ts">
// getTodoList でモジュールスコープのtodoListを参照（全コンポーネント共通の状態）
import { getTodoList } from '../composables/useTask'

const todoList = getTodoList()
</script>

<template>
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
            <!-- v-for でtodoListを順に描画。todoList更新時に自動で再レンダリングされる -->
            <tr v-for="todo in todoList" :key="todo.ID">
                <th>{{ todo.ID }}</th>
                <th>{{ todo.Task }}</th>
                <th>編集</th>
                <th>完了</th>
            </tr>
        </tbody>
    </table>
</template>

<style scoped>
table {
    width: 100%;
    border-collapse: collapse;
    background-color: var(--color-surface);
    border-radius: var(--radius-lg);
    overflow: hidden;
    box-shadow: var(--shadow-md);
}

thead {
    background-color: var(--color-surface-alt);
    border-bottom: 1px solid var(--color-border);
}

thead tr th {
    padding: var(--space-3) var(--space-4);
    font-size: var(--font-size-xs);
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--color-text-muted);
    text-align: left;
}

thead tr th.id-cell {
    width: 60px;
    text-align: center;
}

tbody tr {
    border-bottom: 1px solid var(--color-border);
    transition: background-color var(--transition-fast);
}

tbody tr:last-child {
    border-bottom: none;
}

tbody tr td,
tbody tr th {
    padding: var(--space-4);
    font-size: var(--font-size-sm);
    font-weight: 400;
    color: var(--color-text);
    text-align: left;
    vertical-align: middle;
}
</style>
