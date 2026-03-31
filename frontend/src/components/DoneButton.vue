<script setup lang="ts">
import { useDoneTask } from '../composables/useTask'
import { CircleCheck } from 'lucide-vue-next'

const props = defineProps<{
    taskId: number
}>()

const { doneTask } = useDoneTask()
</script>

<template>
    <button @click="doneTask(props.taskId)" class="icon-btn">
        <CircleCheck :size="30" />
    </button>
</template>

<style scoped>
.icon-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2px;
    background: transparent;
    border: none;
    border-radius: var(--radius-full);
    color: var(--color-success);
    cursor: pointer;
    transition: 
        scale var(--transition-slow), 
        color var(--transition-slow),
        background-color var(--transition-slow);
}

.icon-btn:hover {
    color: white;
    scale: 1.3;
    background-color: var(--color-success);
}
</style>