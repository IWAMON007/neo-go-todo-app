package route

import (
	"encoding/json"
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
var doneListHTML = template.Must(template.ParseFiles("views/done_list.html"))

// タスクの一覧を取得
func getTodoList(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TodoList)
}

func doneList(w http.ResponseWriter, _ *http.Request) {
	var DoneList []Todo

	for i, t := range TodoList {
		if t.IsDone {
			DoneList = append(DoneList, TodoList[i])
		}
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(DoneList)
}

// タスクの追加
func addTask(w http.ResponseWriter, r *http.Request) {
	var request struct{ Task string }

	fmt.Println("タスク追加")

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "読み取り失敗", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	newTodo := Todo{
		ID:     nextID,
		Task:   request.Task,
		IsDone: false,
	}

	TodoList = append(TodoList, newTodo)
	nextID++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

// タスクの完了
func doneTask(w http.ResponseWriter, r *http.Request) {
	var request Todo

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "読み取り失敗", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Printf("[タスク完了リクエスト] 完了タスク：ID = %d\n", request.ID)

	for i, t := range TodoList {
		if t.ID == request.ID {
			TodoList[i].IsDone = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
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

// タスクの更新
func updateTask(w http.ResponseWriter, r *http.Request) {
	var request Todo

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "読み取り失敗", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	for i, t := range TodoList {
		if t.ID == request.ID {
			TodoList[i].Task = request.Task
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}

// ルーティング設定
func SetRoute() {
	route := map[string]func(w http.ResponseWriter, r *http.Request){
		"/todo/list":   getTodoList,
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
