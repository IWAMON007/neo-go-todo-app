# computed によるリストフィルタリング

## v-for + v-if を同じ要素に書くのは避ける

```html
<!-- NG: 全件ループしてから表示を切り替えるだけ。非表示の <tr> もDOMに残る -->
<tr v-for="todo in todoList" :key="todo.ID" v-if="!todo.IsDone">
```

```html
<!-- OK: computed でフィルタ済みリストを作り、v-for の対象を絞る -->
<tr v-for="todo in activeTodoList" :key="todo.ID">
```

```ts
const activeTodoList = computed(() => todoList.filter((todo) => !todo.IsDone));
```

## なぜ computed が良いか

- ループ対象を最初から絞るため、不要な DOM ノードが生まれない
- フィルタロジックがテンプレートから分離され、読みやすくなる
- `computed` はリアクティブなので、`todoList` が変わると自動で再計算される

## 命名の慣習

フィルタ済みリストの命名は内容が伝わる形にする。

```ts
const activeTodoList = computed(...)  // 未完了タスク
const doneTodoList = computed(...)    // 完了タスク
```
