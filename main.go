package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	h := `<h1>計算できる</h1>`
	form := `<form action="/cal" method="GET"><input name="length"><input type="submit"></form>`

	fmt.Fprintln(w, h+form)
}

func calSquareArea(w http.ResponseWriter, r *http.Request) {
	l := r.URL.Query().Get("length")
	oneSide, err := strconv.Atoi(l)

	if err != nil {
		fmt.Printf("入力された文字列:%s", l)
		fmt.Println("数値ではない、文字列が入力されました。")
	}

	area := oneSide * oneSide
	p := "<p>One side length : " + l + " Squar = " + strconv.Itoa(area) + "</p> "
	a := `<a href="/">ここをクリック</a>`
	fmt.Fprint(w, p+a)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/cal", calSquareArea)
	http.ListenAndServe(":8080", nil)
}
