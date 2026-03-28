// src/composables/fromTask.ts
import { ref } from 'vue'

export function fromTask() {
    const newTask = ref<string>("")

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

            newTask.value = ''

        } catch (error: any) {
            console.error('追加に失敗しました:', error.message)
            alert('タスクの追加に失敗しました。通信状況を確認してください。')
        }
    }

    return { newTask, addTask }
}