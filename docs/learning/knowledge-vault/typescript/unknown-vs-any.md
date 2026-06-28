# unknown vs any（エラーハンドリング）

## 違い

| | `any` | `unknown` |
|---|---|---|
| プロパティアクセス | 型チェックなしで可能（危険） | 型ガードするまで不可（安全） |
| 真偽値評価（`!error`） | 可能 | 可能 |
| `throw` | 可能 | 可能（`throw` は型チェック対象外） |

## catch ブロックでの正しい書き方

```ts
// NG: any はプロパティに型チェックなしでアクセスできてしまう
catch (error: any) {
    alert(error.message); // 実行時に undefined になる可能性がある
}

// OK: instanceof で型ガードしてからアクセスする
catch (error: unknown) {
    if (error instanceof Error) {
        throw error; // Error オブジェクトのままthrowする
    }
    throw new Error('予期せぬエラーが発生しました');
}
```

## よくあるミス

```ts
// NG: string を throw している
throw error.message; // error.message は string

// OK: Error オブジェクトをそのまま throw する
throw error;
```

`throw` に渡すのは `Error` オブジェクト。文字列を throw すると、受け取り側で `.message` にアクセスできなくなる。

## throw が unknown でも通る理由

`throw` は TypeScript の型チェック対象外。`throw 42` や `throw "文字列"` も合法なので、`error: unknown` のまま `throw error` と書いてもコンパイルエラーにならない。

## ネットワークエラー vs サーバーエラーの区別

`fetch` が投げるエラーの型は状況によって異なる。

| 状況 | エラーの型 |
|---|---|
| ネットワーク断・サーバー停止 | `TypeError: Failed to fetch` |
| サーバーが 4xx/5xx を返した（`!response.ok`） | `Error`（自分で throw したもの） |

```ts
catch (error: unknown) {
    if (error instanceof TypeError) {
        // fetch 自体が失敗（ネットワーク断）
        throw new Error('通信エラーが発生しました。接続を確認してください。')
    }
    // !response.ok で throw した Error → そのまま再 throw
    throw error
}
```

## dead code になりやすいパターン

```ts
// NG: catch ブロックでは常に何かしら値が入るため !error は常に false
catch (error: unknown) {
    if (!error) { // ← dead code
        throw new Error('エラー');
    }
    throw error;
}
```
