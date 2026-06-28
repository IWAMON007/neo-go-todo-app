# モジュールスコープによる状態共有

## Pinia/Vuex なしでグローバル状態を共有する方法

`ref` をモジュールスコープ（関数の外）に置くと、全コンポーネントで同じインスタンスを共有できる。

```ts
// ✅ 関数の外 → 全コンポーネントで共有される同一インスタンス
export const todoList = ref<Todo[]>([])

export async function getTodoList() {
    const list = await apiFetch<Todo[]>({ ... })
    todoList.value = list
}
```

```ts
// ❌ 関数の中 → 呼ぶたびに新しい ref が生まれる（共有されない）
export function useXxx() {
    const todoList = ref<Todo[]>([])  // 呼び出しごとに別インスタンス
}
```

## 旧設計 vs 新設計

| | 旧（composable 関数を返す） | 新（直接 export） |
|---|---|---|
| `todoList` の取り出し方 | `useGetTodoList()` の戻り値から分割代入 | 直接 `import` |
| 関数の形 | composable 関数の内部に定義・返す | 単独の `async` 関数として `export` |
| コンポーネント側 | `const { todoList, getTodoList } = useGetTodoList()` | `import { todoList, getTodoList } from '...'` |

## 注意点

- シンプルで強力だが、モジュール全体がシングルトンになる
- テスト時にリセットしにくい点はトレードオフ
- 小〜中規模アプリでは Pinia なしでも十分に機能する
