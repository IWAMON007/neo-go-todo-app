<script setup lang="ts">
import { useEditTask } from '../composables/useTask';
import { X } from 'lucide-vue-next';
import ButtonIcon from './layouts/ButtonIcon.vue';

const { cancelEdit } = useEditTask();
</script>

<template>
    <ButtonIcon
        @click="cancelEdit()"
        color="var(--color-danger)"
        border="2px solid var(--color-danger)"
        hover-background-color="var(--color-danger)"
        hover-color="white"
    >
        <X :size="16" :stroke-width="3" />
    </ButtonIcon>
</template>
