# onMounted と async/await

## async/await が必要なケース

```ts
// 必要: 呼び出し側でエラーをハンドリングしたい場合
onMounted(async () => {
    try {
        await getTodoList();
    } catch (e) {
        // ここでエラー処理
    }
});

// 必要: 完了を待って次の処理をしたい場合
onMounted(async () => {
    await getTodoList();
    doSomethingAfter(); // getTodoList 完了後に実行したい処理
});
```

## async/await が不要なケース

```ts
// 不要: apiFetch 内で try/catch が完結していてエラーは onErrorCaptured に委譲する場合
// todoList は ref でリアクティブに共有されているため、完了次第 UI に反映される
onMounted(() => {
    getTodoList();
});
```

## ポイント

`await` が必要なのは「呼び出し側でエラーをハンドリングしたい」か「完了を待って次の処理をしたい」場合。どちらも不要なら省略して問題ない。
