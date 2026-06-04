// src/composables/useTask.ts
// タスクの状態管理・API通信を行うComposable（複数コンポーネントで共有する）
import { ref } from 'vue';
import type { Todo } from '../types/todo';
import { apiFetch } from './useFetch';

// モジュールスコープで定義することで、全コンポーネント間で状態を共有する
export const todoList = ref<Todo[]>([]);

export const newTask = ref<string>('');

// 編集中のタスクID（null = 非編集状態）
export const editingId = ref<number | null>(null);
// 編集中のテキスト
export const editingText = ref<string>('');

// todoListの取得
export async function getTodoList(): Promise<void> {
    const list = await apiFetch<Todo[]>({
        route: '/todo/list',
        method: 'GET',
        error: {
            message: 'タスクの取得に失敗しました。通信状況を確認してください。',
        },
    });

    if (list) todoList.value = list;
}

// タスクの追加
export async function addTask(): Promise<void> {
    if (newTask.value.trim() === '') {
        throw new Error('タスクを入力してください。');
    }

    const addedTask = await apiFetch<Todo>({
        route: '/task',
        method: 'POST',
        body: { Task: newTask.value },
        error: {
            message: 'タスクの追加に失敗しました。通信状況を確認してください。',
        },
    });

    newTask.value = '';

    todoList.value.push(addedTask);
}

// タスクの完了
export async function doneTask(taskID: number): Promise<void> {
    await apiFetch<void>({
        route: '/task/done',
        method: 'PUT',
        body: { ID: taskID },
        error: {
            message: 'タスクの完了に失敗しました。通信状況を確認してください。',
        },
    });

    const index = todoList.value.findIndex((t) => t.ID === taskID);
    if (index !== -1) {
        todoList.value[index].IsDone = true;
    }
}

// タスクの更新
export async function updateTask(taskID: number): Promise<void> {
    if (editingText.value.trim() === '') {
        throw new Error('タスクを入力してください。');
    }
    await apiFetch<void>({
        route: '/task/update',
        method: 'PUT',
        body: {
            ID: taskID,
            Task: editingText.value,
        },
        error: {
            message: 'タスクの更新に失敗しました。通信状況を確認してください。',
        },
    });

    const index = todoList.value.findIndex((t) => t.ID === taskID);
    if (index !== -1) {
        todoList.value[index].Task = editingText.value;
    }

    editingId.value = null;
    editingText.value = '';
}

// タスクの削除
export async function deleteTask(taskID: number): Promise<void> {
    await apiFetch<void>({
        route: '/task/delete',
        method: 'DELETE',
        body: { ID: taskID },
        error: {
            message: 'タスクの削除に失敗しました。通信状況を確認してください。',
        },
    });

    todoList.value = todoList.value.filter((todo) => todo.ID !== taskID);
}

// 編集状態を管理するComposable（EditButton / SaveButton / CancelButton / TaskTable 間で共有）
export function useEditTask() {
    function editTask(taskID: number) {
        const todo = todoList.value.find((t) => t.ID === taskID);
        editingId.value = taskID;
        editingText.value = todo?.Task ?? '';
    }

    // 編集をキャンセルし、編集状態をリセットする
    function cancelEdit() {
        editingId.value = null;
        editingText.value = '';
    }

    return { editingId, editingText, editTask, cancelEdit };
}
