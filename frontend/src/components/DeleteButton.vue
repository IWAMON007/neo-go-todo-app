<script setup lang="ts">
import { useDeleteTask } from '../composables/useTask'
import { Trash2 } from 'lucide-vue-next'

const props = defineProps<{
    taskId: number
}>()

const { deleteTask } = useDeleteTask()
</script>

<template>
    <button @click="deleteTask(props.taskId)" class="icon-btn">
        <Trash2 :size="26" />
    </button>
</template>

<style scoped>
.icon-btn {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px;
    background: transparent;
    border-radius: var(--radius-full);
    color: var(--color-danger);
    cursor: pointer;
    overflow: hidden;
    transition: color 0.3s ease;
    border: none;
}

.icon-btn::before {
    content: '';
    position: absolute;
    inset: 0;
    background-color: var(--color-danger);
    border-radius: 50%;
    transform: scale(0);
    transition: transform 0.3s ease;
}

.icon-btn:hover::before {
    transform: scale(1);
}

.icon-btn:hover {
    color: white;
}

.icon-btn > * {
    position: relative;
    z-index: 1;
}
</style>