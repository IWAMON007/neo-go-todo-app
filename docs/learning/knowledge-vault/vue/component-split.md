# コンポーネント分割の判断軸

## 文字列分岐 vs コンポーネント分割

```ts
// NG: pathName による文字列分岐
// パス名変更時に if 文の修正が必要になる（ハードコーディング）
if (pathName === 'home') { ... }
```

```html
<!-- OK: ルーティングで制御し、用途別にコンポーネントを分割する -->
<!-- Home.vue → TaskTable.vue -->
<!-- Done.vue → DoneTaskTable.vue -->
```

## 分割の判断基準

- **カラム数・表示内容が異なる** — 共通化するより分けた方が素直
- **片方だけ処理が複雑** — 共通コンポーネントに条件分岐が増えると読みにくくなる
- **ルーティングで制御できる** — `pathName` を渡す迂回をやめ、ルーティングの構造に任せる

## BaseTable のようなレイアウト共通化

テーブルのスタイルだけ共通化して、中身は slot で差し込む設計。

```html
<!-- BaseTable.vue -->
<template>
    <div class="table-wrapper">
        <slot />
    </div>
</template>

<style scoped>
/* slot 内の要素には :slotted() が必要 */
:slotted(table) { ... }
</style>
```

scoped CSS の制約で、slot 内の要素には通常のセレクタが効かないため `:slotted()` を使う。
