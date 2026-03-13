package route

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	ID   int
	Task string
}

var TodoList []Todo
var nextID int = 1

// 初期表示時
func home(w http.ResponseWriter, _ *http.Request) {
	html, err := template.ParseFiles("home.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	html.Execute(w, nil)
}

// タスクの追加
func addTask(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("task")

	TodoList = append(TodoList, Todo{ID: nextID, Task: t})
	nextID++

	html, err := template.ParseFiles("home.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	html.Execute(w, TodoList)
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
