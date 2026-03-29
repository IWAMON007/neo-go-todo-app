<script setup lang="ts">
import { useUpdateTask } from '../composables/useTask'
import { Save } from 'lucide-vue-next'

const props = defineProps<{
    taskId: number
}>()

const { updateTask } = useUpdateTask()
</script>

<template>
    <button @click="updateTask(props.taskId)" class="icon-btn">
        <Save :size="24" />
    </button>
</template>

<style scoped>
.icon-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0px;
    background: transparent;
    border: none;
    border-radius: var(--radius-sm);
    color: var(--color-text-muted);
    cursor: pointer;
    transition: background-color var(--transition-fast), color var(--transition-fast);
}

.icon-btn:hover {
    color: var(--color-text);
    transform: scale(1.2);
}
</style>