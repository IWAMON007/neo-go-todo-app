# CLAUDE.md

## Claudeへの作業指示

### ブランチ運用
- **worktree は自動作成しない** — 勝手に `.claude/worktrees/` 以下に worktree を作らないこと
- **main ブランチで作業しない** — 現在のブランチが `main` の場合、作業前に必ずブランチ作成を提案すること
  - 提案例: `feature/add-db-persistence` などの適切な名称を候補として示す
  - ユーザーが承認してからブランチを作成し、そこで作業する
- **既に作業ブランチにいる場合** — そのまま作業して構わない

---

## プロジェクト概要

Go製バックエンド + Vue 3フロントエンドのTODOアプリ。個人学習目的。

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

## 技術スタック

| レイヤー | 技術 |
|---|---|
| バックエンド | Go 1.25.6、標準ライブラリのみ（`net/http`） |
| フロントエンド | Vue 3.5 + TypeScript + Vite 8 + vue-router 4 |
| インフラ | Docker Compose |

## ディレクトリ構成

```
Go_webapp/
├── main.go                          # サーバー起動 + ロギングミドルウェア
├── route/route.go                   # 全APIルーティング・ハンドラー・Todo型定義
├── dockerfile                       # Go開発用（Air込み）
├── docker-compose.yaml
└── frontend/
    ├── vite.config.ts               # APIプロキシ設定（/task, /done, /todo → :8080）
    ├── src/
    │   ├── main.ts                  # アプリ初期化・router登録
    │   ├── App.vue
    │   ├── types/todo.ts            # Todo型定義（バックエンドと対応）
    │   ├── router/route.ts          # ページルーティング
    │   ├── composables/useTask.ts   # 全API通信・状態管理ロジック
    │   ├── views/
    │   │   ├── Home.vue             # 未完了タスク一覧
    │   │   └── Done.vue             # 完了タスク一覧
    │   └── components/
    │       ├── FormTask.vue         # タスク追加フォーム
    │       ├── TaskTable.vue        # タスク一覧テーブル（Home/Done共用）
    │       ├── EditButton.vue
    │       ├── SaveButton.vue
    │       ├── CancelButton.vue
    │       ├── DoneButton.vue
    │       └── DeleteButton.vue
    └── dockerfile
```

## APIエンドポイント

| メソッド | パス | 機能 | リクエストBody |
|---|---|---|---|
| GET | `/todo/list` | 全タスク取得 | なし |
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

## 現状の制約

- **データ永続化なし** — タスクはメモリ内のみ。サーバー再起動で消える。
- **HTTPメソッド検証なし** — バックエンドはメソッドに関係なくルートが応答する。
