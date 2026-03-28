<script setup lang="ts">
import { fromTask } from '../composables/fromTask'

const { newTask, addTask } = fromTask()
</script>

<template>
    <form @submit.prevent="addTask" class="form-row">
        <input 
        v-model="newTask" 
        type="text" 
        placeholder="タスクを入力しよう！！"
        >
        <input type="submit" value="追加">
    </form>
</template>

<style scoped>
.form-row {
    display: flex;
    gap: var(--space-3);
    margin-bottom: var(--space-8);
}

.form-row input[type="text"] {
    flex: 1;
    padding: var(--space-3) var(--space-5);
    font-family: var(--font-family);
    font-size: var(--font-size-base);
    color: var(--color-text);
    background-color: var(--color-surface);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-md);
    outline: none;
    transition:
    border-color var(--transition-fast),
    box-shadow var(--transition-fast);
}

.form-row input[type="text"]::placeholder {
    color: var(--color-text-dim);
}

.form-row input[type="text"]:focus {
    border-color: var(--color-primary);
    box-shadow: 0 0 0 3px var(--color-primary-glow);
}

.form-row input[type="submit"] {
    padding: var(--space-3) var(--space-6);
    font-family: var(--font-family);
    font-size: var(--font-size-sm);
    font-weight: 600;
    color: #fff;
    background: linear-gradient(135deg, var(--color-primary-light) 0%, var(--color-primary) 100%);
    border: none;
    border-radius: var(--radius-md);
    cursor: pointer;
    transition:
    opacity var(--transition-fast),
    transform var(--transition-fast),
    box-shadow var(--transition-fast);
}

.form-row input[type="submit"]:hover {
    opacity: 0.9;
    transform: translateY(-1px);
    box-shadow: var(--shadow-primary);
}

.form-row input[type="submit"]:active {
    transform: translateY(0);
}
</style>