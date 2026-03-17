package route

import (
	"html/template"
	"net/http"
)

type Todo struct {
	ID   int
	Task string
}

var TodoList []Todo
var nextID int = 1

// htmlを事前にパース
// hmtlにエラーがあれば、ここで止まる
var homeHTML = template.Must(template.ParseFiles("views/home.html"))

// 初期表示時
func home(w http.ResponseWriter, _ *http.Request) {
	homeHTML.Execute(w, TodoList)
}

// タスクの追加
func addTask(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("task")

	TodoList = append(TodoList, Todo{ID: nextID, Task: t})
	nextID++

	homeHTML.Execute(w, TodoList)
}

// ルーティング設定
func SetRoute() {
	route := map[string]func(w http.ResponseWriter, r *http.Request){
		"/":     home,
		"/task": addTask,
	}

	for r, h := range route {
		http.HandleFunc(r, h)
	}
}
