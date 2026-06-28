# slot vs props の使い分け

## 基本の判断軸

| | slot | props |
|---|---|---|
| 向いているもの | HTML・コンポーネントなどの「見た目の塊」 | 文字列・数値・真偽値などのデータ |
| 親側の記述 | テンプレートに直接書ける・可読性が高い | 型定義が必要・複雑な値は煩雑になりやすい |

## ボタンコンポーネントでの適用例

```vue
<!-- ButtonIcon.vue: アイコンは slot で受け取る -->
<template>
    <button class="icon-btn">
        <slot />  <!-- 親がアイコンを直接書ける -->
    </button>
</template>
```

```vue
<!-- 親コンポーネント: アイコンを slot に渡す -->
<ButtonIcon>
    <span>✏️</span>
</ButtonIcon>
```

アイコンを props で渡そうとすると型処理が煩雑になるため、slot の方がシンプル。

## フォールスルー属性（@click の扱い）

`@click` などのイベントリスナーは props として明示的に定義しなくても、コンポーネントのルート要素に自動で引き継がれる（フォールスルー属性）。

```vue
<!-- 親: @click を渡す -->
<ButtonIcon @click="handleClick">...</ButtonIcon>

<!-- ButtonIcon.vue: defineProps に @click を書かなくてもルート要素に渡る -->
<template>
    <button>  <!-- ← @click が自動で付く -->
        <slot />
    </button>
</template>
```

Base コンポーネントに全ボタンの処理を持たせるのは責務が広すぎる。`@click` の処理は親で持ち、フォールスルーで渡すのが正しい設計。
