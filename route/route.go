package route

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Todo struct {
	ID     int
	Task   string
	IsDone bool
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

	TodoList = append(TodoList, Todo{ID: nextID, Task: t, IsDone: false})
	nextID++

	homeHTML.Execute(w, TodoList)
}

func doneTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskId, _ := strconv.Atoi(id)

	for i, t := range TodoList {
		if t.ID == taskId {
			TodoList[i].IsDone = true
		}
	}

	fmt.Println(TodoList)

	homeHTML.Execute(w, TodoList)
}

// タスクの削除
func deteleTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskId, _ := strconv.Atoi(id)

	for i, t := range TodoList {
		if t.ID == taskId {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
		}
	}

	homeHTML.Execute(w, TodoList)
}

// ルーティング設定
func SetRoute() {
	route := map[string]func(w http.ResponseWriter, r *http.Request){
		"/":            home,
		"/task":        addTask,
		"/task/delete": deteleTask,
		"/task/done":   doneTask,
	}

	for r, h := range route {
		http.HandleFunc(r, h)
	}
}
