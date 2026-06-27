<script setup lang="ts">
const props = defineProps<{
    padding?: string;
    border?: string;
    color?: string;
    transition?: string;
    borderRadius?: string;
    hoverColor?: string;
    hoverBackgroundColor?: string;
    hoverTransform?: string;
}>();
</script>

<template>
    <button
        class="icon-btn"
        :style="{
            '--padding': padding ?? '2px',
            '--border': border ?? 'none',
            '--color': color ?? 'var(--color-text-muted)',
            '--transition': transition ?? 'none',
            '--border-radius': borderRadius ?? 'var(--radius-full)',
            '--hover-color': hoverColor ?? 'var(--color-text)',
            '--hover-background-color': hoverBackgroundColor ?? 'transparent',
            '--hover-transform': hoverTransform ?? 'scale(1.2)',
        }"
    >
        <slot> </slot>
    </button>
</template>

<style scoped>
.icon-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    padding: var(--padding);
    background-color: transparent;
    border-radius: var(--border-radius);
    border: var(--border);
    color: var(--color);
    transition: var(--transition);
}

.icon-btn:hover {
    color: var(--hover-color);
    background-color: var(--hover-background-color);
    transform: var(--hover-transform);
}
</style>
