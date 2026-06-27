<script setup lang="ts">
import { updateTask } from '../composables/useTask';
import { Save } from 'lucide-vue-next';
import ButtonIcon from './layouts/ButtonIcon.vue';

const props = defineProps<{
    taskId: number;
}>();
</script>

<template>
    <ButtonIcon @click="updateTask(props.taskId)">
        <Save :size="24" />
    </ButtonIcon>
</template>
