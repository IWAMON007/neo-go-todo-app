package route

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
var doneListHTML = template.Must(template.ParseFiles("views/done_list.html"))

// 初期表示時
func home(w http.ResponseWriter, _ *http.Request) {
	homeHTML.Execute(w, TodoList)
}

func doneList(w http.ResponseWriter, _ *http.Request) {
	doneListHTML.Execute(w, TodoList)
}

// タスクの追加
func addTask(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Task string
	}

	t, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "読み取り失敗", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println(string(t))

	json.Unmarshal(t, &request)

	TodoList = append(TodoList, Todo{ID: nextID, Task: request.Task, IsDone: false})
	nextID++

	fmt.Println(TodoList)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}

// タスクの完了
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

	doneListHTML.Execute(w, TodoList)
}

// タスクの編集
func updateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Todo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "読み取り失敗", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Printf("受信データ: %s\n", string(body))

	json.Unmarshal(body, &newTask)

	for i, t := range TodoList {
		if t.ID == newTask.ID {
			TodoList[i].Task = newTask.Task
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}

// ルーティング設定
func SetRoute() {
	route := map[string]func(w http.ResponseWriter, r *http.Request){
		"/":            home,
		"/done/list":   doneList,
		"/task":        addTask,
		"/task/delete": deteleTask,
		"/task/done":   doneTask,
		"/task/update": updateTask,
	}

	for r, h := range route {
		http.HandleFunc(r, h)
	}
}
