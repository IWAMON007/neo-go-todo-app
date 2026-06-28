# タスク一覧

## 進行中

<!-- 現在取り組んでいるタスクをここに書く -->

## やりたいこと（未着手）

### 1. フロントエンドリファクタリング

#### ~~1-1. API通信の共通化（useTask.ts）~~ → 完了（`refactor/api-fetch-helper`）

#### ~~1-2. ボタンコンポーネントのスタイル共通化~~ → 完了（`refactor/base-icon-button`）

#### ~~1-3. TaskTable の責務分離~~ → 完了（`refactor/task-table-split`）

#### ~~1-4. Done.vue のデータ取得バグ修正~~ → 完了（`fix/done-view-onmounted`）

#### 1-5. 型安全性の改善

- [ ] 各composableの `catch (error: any)` を `catch (error: unknown)` に変更する
- [ ] `error` を `unknown` で受けた場合のエラーメッセージ取り出し方を調べて実装する
- **学べること**: TypeScript の `unknown` vs `any`・型ガード

#### 1-6. 細かいバグ・コードの整理

- [ ] `TaskTable.vue` に `import CanselButton`（Cansel = typo、Cancel が正しい）がある
- [ ] `useDeleteTask` の catch ブロックのエラーメッセージが「追加に失敗しました」になっている（削除なのに）
- [ ] 各操作のエラーメッセージを操作内容に合わせて統一する
- **学べること**: コードレビューの視点・バグを見つける力

#### 1-7. エラーハンドリングのリファクタ

- [ ] `useFetch.ts` の `if (!error)` dead code を `instanceof TypeError` による分岐に修正
- [ ] ネットワークエラーとサーバーエラーを区別してメッセージを出し分ける
- [ ] `alert` ではなく UI 上にエラーメッセージを表示する（`ref` で管理）
- [ ] エラーメッセージの表示・非表示のライフサイクルを設計する（成功時に消す・手動で閉じるなど）
- **学べること**: エラーの種類の分類・TypeScript の型ガード・Vue の UI 状態管理

---

### 2. バックエンドリファクタリング

> 今後追記予定

- [ ] DI（依存性の注入）を意識した構造へ修正
- DB実装前に構造を整えておくと、リポジトリパターン導入がスムーズになる

### 3. DB設計ドキュメントの作成

> 今後追記予定

- [ ] テーブル構成の定義
- [ ] リレーションの定義
- [ ] `docs/` 以下に配置

### 4. DB永続化の実装

> 今後追記予定

- [ ] PostgreSQL を Docker コンテナで起動する設定（`docker-compose.yaml` 追記）
- [ ] Go バックエンドから PostgreSQL に接続
- [ ] 既存のインメモリ実装をDB実装に置き換え

### 5. テストの作成

> 今後追記予定

- [ ] バックエンドのユニットテスト・統合テスト
- [ ] フロントエンドのテスト
- Claude にテストを書いてもらい、それが通る実装にする方針

## 完了

### ブランチ: `fix/done-view-onmounted`

> ブランチ名は `onMounted` 修正を想定していたが、調査の結果 `Done.vue` には既に `onMounted` が実装済みだった。
> 代わりに同ブランチで以下のバグ修正・改善を実施。

- [x] `TaskTable.vue` / `DoneTaskTable.vue` に `IsDone` によるフィルタリングがなく、DoneButton 押下後もタスクが消えないバグを修正
- [x] `v-for` + `v-if` の組み合わせを `computed` によるフィルタ済みリストに変更（`activeTodoList` / `doneTodoList`）
- [x] `Done.vue` の `onMounted` から不要な `async/await` を削除（`apiFetch` 内で `try/catch` が完結しているため）
- [x] `Done.vue` に `onErrorCaptured` を追加（`Home.vue` と同じエラーハンドリングパターンに統一）

### ブランチ: `refactor/task-table-split`

- [x] `BaseTable.vue`（`layouts/`）を新規作成し、`:slotted()` でテーブルスタイルを共通化
- [x] `DoneTaskTable.vue` を新規作成（Done 用・`DeleteButton` のみ）
- [x] `TaskTable.vue` を Home 専用に整理（`pathName` prop 廃止・重複 CSS 修正）
- [x] `Done.vue` が `DoneTaskTable` を使うよう差し替え
- [x] `Home.vue` の不要な `useRoute` / `route` / `:pathName` を除去

### ブランチ: `refactor/base-icon-button`

- [x] `ButtonIcon.vue` を `components/layouts/` に作成（`<slot>` + CSS 変数ベースの汎用ボタン）
- [x] CSS 変数（`--color`・`--border`・`--padding`・`--border-radius`・hover 系）を props 経由で注入する設計を採用
- [x] 全ボタン（EditButton / SaveButton / DoneButton / DeleteButton / CancelButton）を `ButtonIcon` ベースに移行
- [x] 各ボタンから重複していた `.icon-btn` スタイルを排除
- [x] DeleteButton の `::before` アニメーションを廃止し、他ボタンと統一感のある hover に変更

### ブランチ: `refactor/api-fetch-helper`

- [x] `useFetch.ts` を新規作成し、ジェネリクスを用いた `apiFetch<T>` を実装
- [x] `Params` 型に省略可能な `body?: object` を追加
- [x] 全 API 関数（getTodoList / addTask / doneTask / updateTask / deleteTask）を `apiFetch` に置き換え
- [x] `fetch` + `try/catch` + `alert` の重複パターンを排除
- [x] `todoList` / `editingId` / `editingText` を直接 `export` する形に整理
- [x] `onErrorCaptured` による Vue のエラーハンドリング導入

### ブランチ: `docs/task-tracking-setup`

- [x] `docs/TASKS.md` の作成・タスク整理
- [x] `CLAUDE.md` へのタスク管理・メンタリング方針の追記
- [x] `docs/learning/` ディレクトリ構成の構築（daily / inbox / knowledge-vault）
- [x] `docs/learning/CLAUDE.md` の作成（学習ログ運用ルール）
- [x] daily テンプレート・初回 daily の作成
