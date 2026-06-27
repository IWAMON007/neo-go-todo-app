<script setup lang="ts">
import { CircleCheck } from 'lucide-vue-next';
import ButtonIcon from './layouts/ButtonIcon.vue';
import { doneTask } from '../composables/useTask';

const props = defineProps<{
    taskId: number;
}>();
</script>

<template>
    <ButtonIcon
        @click="doneTask(props.taskId)"
        :transition="`
            scale var(--transition-slow),
            color var(--transition-slow),
            background-color var(--transition-slow)
            `"
        color="var(--color-success)"
        hover-background-color="var(--color-success)"
        hover-color="white"
    >
        <CircleCheck :size="30" />
    </ButtonIcon>
</template>
