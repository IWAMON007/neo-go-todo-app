<script setup lang="ts">
import { deleteTask } from '../composables/useTask';
import { Trash2 } from 'lucide-vue-next';
import ButtonIcon from './layouts/ButtonIcon.vue';

const props = defineProps<{
    taskId: number;
}>();
</script>

<template>
    <ButtonIcon
        @click="deleteTask(props.taskId)"
        :transition="`
            scale var(--transition-slow),
            color var(--transition-slow),
            background-color var(--transition-slow)
            `"
        color="var(--color-danger)"
        hover-background-color="var(--color-danger)"
        padding="5px"
        hover-color="white"
    >
        <Trash2 :size="24" />
    </ButtonIcon>
</template>
