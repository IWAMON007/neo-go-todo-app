<script setup lang="ts">
import type { Todo } from '../types/todo';
import DeleteButton from './DeleteButton.vue';
import BaseTable from './layouts/BaseTable.vue';

defineProps<{
    todoList: Todo[];
}>();
</script>

<template>
    <BaseTable>
        <table>
            <thead>
                <tr>
                    <th class="task-cell">タスク</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <template v-for="todo in todoList" :key="todo.ID">
                    <tr>
                        <td>{{ todo.Task }}</td>
                        <td class="delete-cell">
                            <DeleteButton :taskId="todo.ID" />
                        </td>
                    </tr>
                </template>
            </tbody>
        </table>
    </BaseTable>
</template>

<style scoped>
.delete-cell {
    width: 80px;
}
</style>
