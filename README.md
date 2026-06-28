# Go_webapp

Go製バックエンド + Vue 3フロントエンドのTODOアプリ。

## このリポジトリについて

構文から学ぶ教科書的なアプローチではなく、**実際のアプリを作りながら体系的に学ぶ**ことを目的とした学習環境の実験的な取り組み。

### 対象読者

- ある程度プログラミング経験があり、今さら構文から学ぶのは遠回りだと感じている人
- フレームワークのアーキテクチャはざっくり理解しているので、まず動くものを作って理解を深めたい人
- 自分で実装してみて、AIにレビューしてもらいながら知識を整理していきたい人

### アプローチ

1. **実装する** — まず自分でコードを書く
2. **レビューを受ける** — Claude に実装をレビューしてもらい、改善点・設計の意図・より良い書き方のフィードバックをもらう
3. **振り返る** — やりとりを `docs/learning/` に記録し、後から知識として参照できる形に整理する

## 技術スタック

| レイヤー | 技術 |
|---|---|
| バックエンド | Go 1.25.6、標準ライブラリのみ（`net/http`） |
| フロントエンド | Vue 3.5 + TypeScript + Vite 8 + vue-router 4 |
| インフラ | Docker Compose |

## 起動方法

```bash
# 開発環境の起動（ルートディレクトリで実行）
docker compose up

# バックエンドのみ
docker compose up go-app

# フロントエンドのみ
docker compose up vue-app
```

| サービス | URL |
|---|---|
| フロントエンド | http://localhost:5173 |
| バックエンドAPI | http://localhost:8080 |

バックエンドは **Air** によるホットリロードが有効。フロントエンドへのAPIリクエストは Vite のプロキシ（`vite.config.ts`）経由で `:8080` に転送される。

## ディレクトリ構成

```
Go_webapp/
├── main.go                          # サーバー起動 + ロギングミドルウェア
├── route/route.go                   # 全APIルーティング・ハンドラー・Todo型定義
├── dockerfile                       # Go開発用（Air込み）
├── docker-compose.yaml
├── docs/                            # プロジェクトドキュメント
│   ├── TASKS.md                     # タスク一覧・進捗管理
│   └── learning/                    # 学習ログ（daily / inbox / knowledge-vault）
└── frontend/
    ├── vite.config.ts               # APIプロキシ設定（/task, /done, /todo → :8080）
    ├── src/
    │   ├── main.ts                  # アプリ初期化・router登録
    │   ├── App.vue
    │   ├── types/todo.ts            # Todo型定義（バックエンドと対応）
    │   ├── router/route.ts          # ページルーティング
    │   ├── composables/
    │   │   ├── useFetch.ts          # 汎用API通信ヘルパー（apiFetch）
    │   │   └── useTask.ts           # 全API関数・状態管理ロジック
    │   ├── views/
    │   │   ├── Home.vue             # 未完了タスク一覧
    │   │   └── Done.vue             # 完了タスク一覧
    │   └── components/
    │       ├── layouts/
    │       │   ├── BaseTable.vue    # テーブル共通レイアウト（:slotted()でスタイル共通化）
    │       │   └── ButtonIcon.vue   # ボタン共通コンポーネント（CSS変数で各ボタンにスタイル注入）
    │       ├── FormTask.vue         # タスク追加フォーム（テキスト入力・送信）
    │       ├── TaskTable.vue        # タスク一覧テーブル（Home用・編集状態の表示切替）
    │       ├── DoneTaskTable.vue    # 完了タスク一覧テーブル（Done用・削除のみ）
    │       ├── EditButton.vue       # 編集モードを開始するボタン
    │       ├── SaveButton.vue       # 編集内容を保存するボタン
    │       ├── CancelButton.vue     # 編集をキャンセルするボタン
    │       ├── DoneButton.vue       # タスクを完了済みにするボタン
    │       └── DeleteButton.vue     # タスクを削除するボタン
    └── dockerfile
```

## APIエンドポイント

| メソッド | パス | 機能 | リクエストBody |
|---|---|---|---|
| GET | `/todo/list` | 未完了タスク取得 | なし |
| GET | `/done/list` | 完了タスク取得 | なし |
| POST | `/task` | タスク追加 | `{"Task": string}` |
| PUT | `/task/done` | タスク完了 | `{"ID": number}` |
| PUT | `/task/update` | タスク編集 | `{"ID": number, "Task": string}` |
| DELETE | `/task/delete` | タスク削除 | `{"ID": number}` |

## データ構造

```go
// route/route.go
type Todo struct {
    ID     int
    Task   string
    IsDone bool
}
```

```typescript
// frontend/src/types/todo.ts
interface Todo {
    ID: number;
    Task: string;
    IsDone: boolean;
}
```

## フロントエンドの状態管理

`useTask.ts` をモジュールスコープで定義することで Pinia/Vuex なしにグローバル状態を共有している。

- `todoList` — 全タスクのリアクティブな配列（全Composable共通）
- `editingId` / `editingText` — 編集中タスクのID・テキスト

API成功後はサーバーレスポンスを待たず、フロント側で直接 `todoList` を更新する（楽観的更新）。
