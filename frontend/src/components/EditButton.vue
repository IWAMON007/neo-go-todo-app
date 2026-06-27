<script setup lang="ts">
import { useEditTask } from '../composables/useTask';
import ButtonIcon from './layouts/ButtonIcon.vue';
import { Pencil } from 'lucide-vue-next';

const props = defineProps<{
    taskId: number;
}>();

const { editTask } = useEditTask();
</script>

<template>
    <ButtonIcon
        @click="editTask(props.taskId)"
        :transition="`
            background-color var(--transition-fast),
            color var(--transition-fast)
        `"
        hover-transform="scale(1.3)"
    >
        <Pencil :size="20" />
    </ButtonIcon>
</template>
