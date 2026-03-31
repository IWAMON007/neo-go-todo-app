// src/composables/useTask.ts
// タスクの状態管理・API通信を行うComposable（複数コンポーネントで共有する）
import { ref } from 'vue'

// バックエンドから返ってくるタスクの型定義
interface Todo {
    ID: number,
    Task: string,
    IsDone: boolean,
}

// モジュールスコープで定義することで、全コンポーネント間で状態を共有する
const todoList = ref<Todo[]>([])

// 編集中のタスクID（null = 非編集状態）
const editingId = ref<number | null>(null)
// 編集中のテキスト
const editingText = ref<string>('')

// todoListをコンポーネントから参照するためのゲッター
export function getTodoList() {
    return todoList
}

export function useDoneTask() {
    async function doneTask(taskID: number) {
        try {
            const response = await fetch(`/task/done`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ ID: taskID, Task: editingText.value }),
            })

            if (!response.ok) {
                throw new Error('サーバーエラーが発生しました')
            }

            // 成功確認できればサーバーレスポンスを待たずにフロント側で直接更新
            const index = todoList.value.findIndex(t => t.ID === taskID)
            if (index !== -1) {
                todoList.value[index].IsDone = true
                console.log(todoList.value[index])
            }

        } catch (error: any) {
            console.error('更新に失敗しました:', error.message)
            alert('タスクの更新に失敗しました。通信状況を確認してください。')
        }
    }

    return { doneTask }
}

// 編集状態を管理するComposable（EditButton / SaveButton / CancelButton / TaskTable 間で共有）
export function useEditTask() {
    function editTask(taskID: number) {
        const todo = todoList.value.find(t => t.ID === taskID)
        editingId.value = taskID
        editingText.value = todo?.Task ?? ''
    }

    // 編集をキャンセルし、編集状態をリセットする
    function cancelEdit() {
        editingId.value = null
        editingText.value = ''
    }

    return { editingId, editingText, editTask, cancelEdit }
}

// タスク更新を管理するComposable（SaveButton から呼び出す）
export function useUpdateTask() {
    async function updateTask(taskID: number): Promise<void> {
        try {
            const response = await fetch(`/task/update`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ ID: taskID, Task: editingText.value }),
            })

            if (!response.ok) {
                throw new Error('サーバーエラーが発生しました')
            }

            // 成功確認できればサーバーレスポンスを待たずにフロント側で直接更新
            const index = todoList.value.findIndex(t => t.ID === taskID)
            if (index !== -1) {
                todoList.value[index].Task = editingText.value
            }

            editingId.value = null
            editingText.value = ''

        } catch (error: any) {
            console.error('更新に失敗しました:', error.message)
            alert('タスクの更新に失敗しました。通信状況を確認してください。')
        }
    }

    return { updateTask }
}

// タスク追加フォームに関するロジックをまとめたComposable
export function fromTask() {
    // フォームの入力値をリアクティブに管理
    const newTask = ref<string>("")

    // タスクをサーバーへPOSTし、成功時にtodoListへ追加する
    async function addTask(): Promise<void> {
        if (newTask.value.trim() === '') return

        try {
            const response = await fetch('/task', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ Task: newTask.value }),
            })

            if (!response.ok) {
                throw new Error('サーバーエラーが発生しました')
            }

            // レスポンスをTodo型としてパース（awaitが必要）
            const responseNewTask: Todo = await response.json()
            todoList.value.push(responseNewTask)

            newTask.value = ''

        } catch (error: any) {
            console.error('追加に失敗しました:', error.message)
            alert('タスクの追加に失敗しました。通信状況を確認してください。')
        }
    }

    return { newTask, addTask }
}