# fetch のエラーハンドリング

## fetch はHTTPエラーで例外を投げない

```ts
// fetch() はサーバーが 404/500 を返しても例外にならない
const response = await fetch('/api/todos');

// !response.ok で意図的に例外に変換する必要がある
if (!response.ok) {
    throw new Error('サーバーエラーが発生しました');
}
```

| 状況 | `response.ok` | catch に入るか |
|---|---|---|
| 正常（2xx） | `true` | 入らない |
| サーバーエラー（4xx/5xx） | `false` | **入らない**（`throw` しないと） |
| ネットワーク障害 | — | **入る** |

## エラーの種類と処理の分担

```ts
try {
    const response = await fetch(url);

    if (!response.ok) {
        // HTTPエラー（4xx/5xx）→ 呼び出し元固有のメッセージを使う
        throw new Error(params.error.message);
    }

    return await response.json();
} catch (error: unknown) {
    if (error instanceof Error) {
        throw error; // そのまま上位に伝播させる
    }
    // 非 Error な例外（文字列など）は Error にラップして投げる
    throw new Error('予期せぬエラーが発生しました');
}
```

## 将来の発展形

より厳密にするには HTTP ステータスコードごとにカスタム `Error` 型を作る設計が有効。

```ts
class HttpError extends Error {
    constructor(public status: number, message: string) {
        super(message);
    }
}
```

小規模なうちは過剰設計になるため、アプリが大きくなったタイミングで導入を検討する。
