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

// todoListをコンポーネントから参照するためのゲッター
export function getTodoList () {
    return todoList
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