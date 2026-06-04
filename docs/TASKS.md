# タスク一覧

## 進行中

<!-- 現在取り組んでいるタスクをここに書く -->

## やりたいこと（未着手）

### 1. フロントエンドリファクタリング

#### ~~1-1. API通信の共通化（useTask.ts）~~ → 完了（`refactor/api-fetch-helper`）

#### 1-2. ボタンコンポーネントのスタイル共通化
- [ ] `EditButton` / `SaveButton` / `DoneButton` / `DeleteButton` それぞれに `.icon-btn` スタイルが重複している
- [ ] 共通の `BaseIconButton` コンポーネントを作り、各ボタンはそれを利用する形にする
- **学べること**: Vue のコンポーネント設計・スロット（`<slot>`）の使い方・スタイルの継承

#### 1-3. TaskTable の責務分離
- [ ] 現在 `pathName` prop で `'Home'` / `'Done'` を判定してレイアウトを切り替えている
- [ ] 文字列比較で分岐するのは壊れやすい設計。`variant: 'active' | 'done'` のような型付きpropに変更するか、コンポーネントを分割する
- **学べること**: Vueのprops設計・Union型・コンポーネント分割の判断軸

#### 1-4. Done.vue のデータ取得バグ修正
- [ ] `Done.vue` は `getTodoList()` を呼んでいない（`onMounted` なし）
- [ ] `/done` に直接アクセスすると一覧が空になるバグがある
- [ ] コメントアウトされた `watch` コードも整理する
- **学べること**: Vue のライフサイクル・ルート直接アクセスを考慮した設計

#### 1-5. 型安全性の改善
- [ ] 各composableの `catch (error: any)` を `catch (error: unknown)` に変更する
- [ ] `error` を `unknown` で受けた場合のエラーメッセージ取り出し方を調べて実装する
- **学べること**: TypeScript の `unknown` vs `any`・型ガード

#### 1-6. 細かいバグ・コードの整理
- [ ] `TaskTable.vue` に `import CanselButton`（Cansel = typo、Cancel が正しい）がある
- [ ] `useDeleteTask` の catch ブロックのエラーメッセージが「追加に失敗しました」になっている（削除なのに）
- [ ] 各操作のエラーメッセージを操作内容に合わせて統一する
- **学べること**: コードレビューの視点・バグを見つける力

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
