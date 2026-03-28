<script setup lang="ts">
import { fromTask } from '../composables/fromTask'

const { newTask, addTask } = fromTask()
</script>

<template>
    <form @submit.prevent="addTask" class="form-row">
        <input 
        v-model="newTask" 
        type="text" 
        placeholder="タスクを入力しよう！！"
        >
        <input type="submit" value="追加">
    </form>
</template>