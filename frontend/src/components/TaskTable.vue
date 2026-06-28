<script setup lang="ts">
// getTodoList でモジュールスコープのtodoListを参照（全コンポーネント共通の状態）
import { useEditTask } from '../composables/useTask';
import type { Todo } from '../types/todo';
import EditButton from './EditButton.vue';
import SaveButton from './SaveButton.vue';
import CanselButton from './CancelButton.vue';
import DoneButton from './DoneButton.vue';
import BaseTable from './layouts/BaseTable.vue';
import { computed } from 'vue';

const { todoList } = defineProps<{
    todoList: Todo[];
}>();

// 編集状態管理
const { editingId, editingText } = useEditTask();

const activeTodoList = computed(() => todoList.filter((todo) => !todo.IsDone));
</script>

<template>
    <BaseTable>
        <table>
            <thead>
                <tr>
                    <th class="task-cell">タスク</th>
                    <th class="edit-cell"></th>
                    <th class="done-cell"></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="todo in activeTodoList" :key="todo.ID">
                    <td class="task-cell">
                        <input
                            v-if="editingId === todo.ID"
                            v-model="editingText"
                            type="text"
                            class="task-input"
                            autofocus
                        />

                        <span v-else>{{ todo.Task }}</span>
                    </td>

                    <td v-if="editingId !== todo.ID" class="edit-cell">
                        <div class="cell-inner">
                            <EditButton :taskId="todo.ID" />
                        </div>
                    </td>

                    <td v-else class="edit-cell">
                        <div class="cell-inner">
                            <SaveButton :taskId="todo.ID" />
                            <CanselButton />
                        </div>
                    </td>

                    <td class="done-cell">
                        <DoneButton :taskId="todo.ID" />
                    </td>
                </tr>
            </tbody>
        </table>
    </BaseTable>
</template>

<style scoped>
.task-cell {
    width: auto;
}

.edit-cell {
    width: 120px;
    text-align: right;
}

.done-cell {
    width: 80px;
    text-align: right;
}

.edit-cell {
    padding: var(--space-2) var(--space-2);
}

.edit-cell :deep(.cell-inner) {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: var(--space-4);
}

.done-cell {
    text-align: center;
}

.task-input {
    width: 100%;
    background: transparent;
    border: none;
    border-bottom: 1px solid var(--color-border);
    outline: none;
    font-size: var(--font-size-sm);
    color: var(--color-text);
    font-family: inherit;
    padding: 2px 0;
}
</style>
