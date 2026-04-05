<script setup lang="ts">
// getTodoList でモジュールスコープのtodoListを参照（全コンポーネント共通の状態）
import { useEditTask } from '../composables/useTask'
import type { Todo } from '../types/todo'
import EditButton from './EditButton.vue'
import SaveButton from './SaveButton.vue'
import CanselButton from './CancelButton.vue'
import DoneButton from './DoneButton.vue'

const props = defineProps<{
    todoList: Todo[];
    pathName?: string | symbol | null;
}>()

// 編集状態管理
const { editingId, editingText } = useEditTask()
</script>

<template>
    <table>
        <thead>
            <tr>
                <th class="task-cell">タスク</th>
                <th class="edit-cell"></th>
                <th class="done-cell"></th>
            </tr>
        </thead>
        <tbody>
            <!-- v-for でtodoListを順に描画。todoList更新時に自動で再レンダリングされる -->
            <template v-for="todo in todoList" :key="todo.ID">
                <tr v-if="pathName === 'Home' && !todo.IsDone">
                    <td class="task-cell">
                        <!-- 編集中の行だけ input を表示、それ以外はテキスト -->
                        <input
                            v-if="editingId === todo.ID"
                            v-model="editingText"
                            type="text"
                            class="task-input"
                            autofocus
                        />

                        <!-- 通常時はタスクを表示 -->
                        <span v-else>{{ todo.Task }}</span>
                    </td>

                    <!-- 編集ボタン -->
                    <td v-if="editingId !== todo.ID" class="edit-cell">
                        <div class="cell-inner">
                            <EditButton :taskId="todo.ID" />
                        </div>
                    </td>

                    <!-- 編集中の保存とキャンセルボタン -->
                    <td v-else class="edit-cell">
                        <div class="cell-inner">
                            <SaveButton :taskId="todo.ID" />
                            <CanselButton />
                        </div>
                    </td>

                    <!-- タスク完了ボタン -->
                    <td class="done-cell">
                        <DoneButton :taskId="todo.ID" />
                    </td>
                </tr>

                <!-- 完了タスク表示時 -->
                <tr v-if="pathName === 'Done' && todo.IsDone">
                    <td>{{ todo.Task }}</td>
                    <td></td>
                    <td>削除</td>
                </tr>
            </template>
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

thead tr th.task-cell {
    width: auto;
}

thead tr th.edit-cell {
    width: 120px;
    text-align: right;
}

thead tr th.done-cell {
    width: 80px;
    text-align: right;
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

tbody tr td.edit-cell,
tbody tr th.edit-cell {
    padding: var(--space-2) var(--space-2);
}

tbody tr td.edit-cell :deep(.cell-inner),
tbody tr th.edit-cell :deep(.cell-inner) {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: var(--space-4);
}

tbody tr td.done-cell,
tbody tr th.done-cell {
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
