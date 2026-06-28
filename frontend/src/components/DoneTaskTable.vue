<script setup lang="ts">
import { computed } from 'vue';
import type { Todo } from '../types/todo';
import DeleteButton from './DeleteButton.vue';
import BaseTable from './layouts/BaseTable.vue';

const { todoList } = defineProps<{
    todoList: Todo[];
}>();

const doneTodoList = computed(() => todoList.filter((todo) => todo.IsDone));
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
                <tr v-for="todo in doneTodoList" :key="todo.ID">
                    <td>{{ todo.Task }}</td>
                    <td class="delete-cell">
                        <DeleteButton :taskId="todo.ID" />
                    </td>
                </tr>
            </tbody>
        </table>
    </BaseTable>
</template>

<style scoped>
.delete-cell {
    width: 80px;
}
</style>
