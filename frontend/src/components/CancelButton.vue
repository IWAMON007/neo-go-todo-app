<script setup lang="ts">
import { useEditTask } from '../composables/useTask'
import { X } from 'lucide-vue-next'

const { cancelEdit } = useEditTask()
</script>

<template>
    <button @click="cancelEdit()" class="icon-btn">
        <X :size="16" :stroke-width="3" />
    </button>
</template>

<style scoped>
.icon-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2px;
    background: transparent;
    border: 2px solid var(--color-danger);
    border-radius: var(--radius-full);
    color:  var(--color-danger);
    cursor: pointer;
    transition: background-color var(--transition-fast), color var(--transition-fast);
}

.icon-btn:hover {
    background-color: var(--color-danger);
    color: white;
    transform: scale(1.2);
}
</style>