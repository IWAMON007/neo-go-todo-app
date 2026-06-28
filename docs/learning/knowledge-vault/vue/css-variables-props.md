# CSS 変数を props 経由で渡す

## scoped CSS の詳細度問題

scoped CSS は内部的に `[data-v-xxx]` 属性セレクタが付くため、詳細度が上がる。
グローバル CSS（scoped なし）よりも常に優先されてしまう。

```vue
<!-- ButtonIcon.vue（scoped なし） -->
<style>
.icon-btn { color: blue; }
</style>

<!-- EditButton.vue（scoped あり） -->
<style scoped>
.icon-btn { color: red; }  /* [data-v-xxx].icon-btn → 詳細度が高く常に勝つ -->
</style>
```

この問題を回避するために CSS 変数（カスタムプロパティ）を props 経由で渡す設計を採用。

## CSS 変数を props で渡す実装

```vue
<!-- ButtonIcon.vue -->
<script setup lang="ts">
const props = defineProps<{
    color?: string;
    border?: string;
}>();
</script>

<template>
    <button
        class="icon-btn"
        :style="{
            '--color': props.color ?? 'var(--color-default)',
            '--border': props.border ?? 'none',
        }"
    >
        <slot />
    </button>
</template>

<style scoped>
.icon-btn {
    color: var(--color);
    border: var(--border);
}
</style>
```

```vue
<!-- 各ボタンから使う -->
<ButtonIcon color="var(--color-success)" border="1px solid green">
    ✓
</ButtonIcon>
```

## ポイント

- 既存の CSS 変数（`var(--color-success)` など）は文字列としてそのまま渡せる
- デフォルト値は `??` 演算子で設定し、props をオプショナルにできる
- `transition` などの複雑な値も CSS 文字列として渡す方がシンプル（オブジェクトにしない）

## 静的 vs 動的バインディング

```vue
<!-- ❌ 静的な文字列に : バインディングは不要 -->
<ButtonIcon :color="'red'">

<!-- ✅ 静的なら : なしで渡す -->
<ButtonIcon color="red">

<!-- ✅ 動的な値（変数・式）には : を使う -->
<ButtonIcon :color="buttonColor">
```
