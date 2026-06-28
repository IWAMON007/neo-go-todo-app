# エラーの伝播と onErrorCaptured

## throw するとエラーは呼び出し元に伝わる

`apiFetch` 内で `throw` すると、`await` している呼び出し元にエラーが伝播する。
`catch` していない関数は素通りして、さらに上へ伝わり続ける。

```
apiFetch → throw
  ↓
addTask → catch なし → 素通り
  ↓
FormTask.vue → catch なし → 素通り
  ↓
Home.vue → onErrorCaptured でキャッチ ✓
```

## onErrorCaptured の基本

- 子コンポーネントから伝播してきたエラーをキャッチするライフサイクルフック
- `return false` を返すと、エラーがグローバルエラーハンドラーへ伝播するのを止める
- **同じコンポーネント自身のエラーは自分の `onErrorCaptured` では拾えない**（親が必要）
- `<script setup>` の中で呼ばないと "no active component instance" の警告が出る

```ts
// Home.vue
onErrorCaptured((error) => {
    alert(error.message);
    return false; // 上位への伝播を止める
});
```

## onMounted での Promise の渡し方

```ts
// NG: Promise を返していない → Vue がエラーを追跡できない → unhandled rejection になる
onMounted(async () => {
    await getTodoList()
})

// OK: Promise を返す → Vue が callWithAsyncErrorHandling でラップする
onMounted(() => getTodoList())
```

## 「通知」と「処理を止める」は別の関心事

| 場所 | 役割の例 |
|---|---|
| `apiFetch` | 通信エラーの検知・throw |
| `addTask` など | ビジネスロジックの失敗処理 |
| コンポーネント | UI の状態をエラー表示に切り替える |

「`apiFetch` が全部やる」設計はシンプルだが、コンポーネントごとに異なるエラーUIを出したい場合に対応しにくい。学習段階では十分。
