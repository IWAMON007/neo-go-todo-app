<script setup lang="ts">
import { onErrorCaptured } from 'vue';

// App.vue: getTodoList のエラーを拾う
onErrorCaptured((error) => {
    alert(error.message);
    return false;
});
</script>
<template>
    <RouterView />
</template>
