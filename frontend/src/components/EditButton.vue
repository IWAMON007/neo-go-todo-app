<script setup lang="ts">
import { useEditTask } from '../composables/useTask'
import { Pencil } from 'lucide-vue-next'

const props = defineProps<{
    taskId: number
}>()

const { editTask } = useEditTask()
</script>

<template>
    <button @click="editTask(props.taskId)" class="icon-btn">
        <Pencil :size="20" />
    </button>
</template>

<style scoped>
.icon-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 6px;
    background: transparent;
    border: none;
    border-radius: var(--radius-sm);
    color: var(--color-text-muted);
    cursor: pointer;
    transition: background-color var(--transition-fast), color var(--transition-fast);
}

.icon-btn:hover {
    color: var(--color-text);
    transform: scale(1.3);
}
</style>