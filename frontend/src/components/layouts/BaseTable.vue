<script setup lang="ts"></script>

<template>
    <slot> </slot>
</template>

<style scoped>
:slotted(table) {
    width: 100%;
    border-collapse: collapse;
    background-color: var(--color-surface);
    border-radius: var(--radius-lg);
    overflow: hidden;
    box-shadow: var(--shadow-md);
}

:slotted(thead) {
    background-color: var(--color-surface-alt);
    border-bottom: 1px solid var(--color-border);
}

:slotted(th) {
    padding: var(--space-3) var(--space-4);
    font-size: var(--font-size-xs);
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--color-text-muted);
    text-align: left;
}

:slotted(tr) {
    border-bottom: 1px solid var(--color-border);
    transition: background-color var(--transition-fast);
}

:slotted(tr:last-child) {
    border-bottom: none;
}
:slotted(td),
:slotted(th) {
    padding: var(--space-4);
    font-size: var(--font-size-sm);
    font-weight: 400;
    color: var(--color-text);
    text-align: left;
    vertical-align: middle;
}
</style>
