# CLAUDE.md

## Claudeへの作業指示

### タスク管理
- **作業前に必ず [`docs/TASKS.md`](docs/TASKS.md) を確認すること** — 現在進行中・次にやることを把握してから作業に入る
- ユーザーがやりたいことをラフに伝えたら、Claude がタスク形式に整理して `docs/TASKS.md` に反映する

### 開発スタンス（重要）
このプロジェクトはユーザーのスキルアップが目的。**Claude が全部作るのではなく、ユーザーが自分で実装することを支援する。**

- **課題を出す** — タスクに取り組む前に「何を調べるべきか」「どう設計するか」をユーザーに考えさせる
- **ヒントを出す** — 詰まったときはコード全体を渡すのではなく、考え方・キーワード・参考になる概念を示す
- **コードレビューをする** — ユーザーが書いたコードを確認して、改善点や別のアプローチを提案する
- **全部書くのは最終手段** — ユーザーが「書いて」と明示した場合や、学習目的で完成形を見たい場合のみ

### 学習ログの記録
詳細なルールは [`docs/learning/CLAUDE.md`](docs/learning/CLAUDE.md) に記載。

- **会話中に重要なやりとりが発生したら都度** `docs/learning/daily/YYYY-MM-DD.md` に追記する
- 日付をまたいだ場合は新しいファイルを作成する
- タスク完了時はレビューをまとめて daily の「タスクレビュー」セクションに追記する
- 「トピックに整理して」と指示があったら daily を読んで `knowledge-vault/` 以下に整理し、使用済み daily は `inbox/` に移動する

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
    │   ├── composables/useTask.ts   # 全API通信・状態管理ロジック
    │   ├── views/
    │   │   ├── Home.vue             # 未完了タスク一覧
    │   │   └── Done.vue             # 完了タスク一覧
    │   └── components/
    │       ├── FormTask.vue         # タスク追加フォーム（テキスト入力・送信）
    │       ├── TaskTable.vue        # タスク一覧テーブル（Home/Done共用・編集状態の表示切替）
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

## コミットルール


| Prefix | 意味・用途 | 具体例 |
| :--- | :--- | :--- |
| `feat` | 新機能の追加 | 検索フィルター機能の実装 |
| `fix` | バグの修正 | ログイン時のバリデーションエラー修正 |
| `docs` | ドキュメントのみの変更 | READMEの更新、コメントの追記 |
| `style` | コードの意味に影響しない変更 | インデントの修正、セミコロンの追加 |
| `refactor` | リファクタリング | 変数名の変更、関数の共通化 |
| `perf` | パフォーマンス改善 | ループ処理の最適化、メモリ使用量削減 |
| `test` | テスト関連 | ユニットテストの追加、既存テストの修正 |
| `chore` | 雑務・設定変更 | ライブラリのアップデート、ビルド設定の変更 |
| `build` | ビルドシステムの影響 | npmやcmakeなどの依存関係変更 |
| `ci` | CI設定の影響 | GitHub ActionsやCircleCIの更新 |
| `revert` | 変更の打ち消し | 以前のコミットを取り消す |
