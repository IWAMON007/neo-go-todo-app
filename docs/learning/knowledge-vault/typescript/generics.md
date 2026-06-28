# ジェネリクス `<T>`

## 基本の考え方

`<T>` は「型の引数」。関数の引数が値を受け取るように、**型自体を外から受け取る**仕組み。

- 定義側は `<T>` をプレースホルダーとして宣言するだけで、中身は呼び出し側が決める
- `T` は慣習的な名前（Type の略）。`Data` や `Response` でも動くが `T` が一般的

```ts
// 定義側: T が何かは知らない
async function apiFetch<T>(params: Params): Promise<T> { ... }

// 呼び出し側: T を指定して型を確定させる
const list = await apiFetch<Todo[]>({ ... })
```

## Promise との組み合わせ

```ts
// Promise<void>: 値を返さず副作用だけ行う非同期関数
export async function getTodoList(): Promise<void> {
    const list = await apiFetch<Todo[]>({ ... })
    if (list) todoList.value = list
}
```

- `async` 関数を定義すると、戻り値は自動的に `Promise` でラップされる
- `await` は `Promise` が解決されるまで待ち、中身の値を取り出す
- `await` を使うには、その関数自体も `async` である必要がある

## 無名関数の注意点

```ts
// NG: 関数を定義しているだけで実行していない
async () => { await fn() }

// OK: 即時実行（IIFE）
(async () => { await fn() })()

// 無名関数が輝くのは「引数として渡すとき」
onMounted(async () => { await fn() })
array.map(item => item.id)
```
