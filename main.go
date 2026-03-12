package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello</h1>")
}

func calSquareArea(w http.ResponseWriter, r *http.Request) {
	l := r.URL.Query().Get("length")
	oneSide, _ := strconv.Atoi(l)

	area := oneSide * oneSide
	fmt.Fprintf(w, "<p>One side length : %s Squar = %d</p> ", l, area)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cal", calSquareArea)
	http.ListenAndServe(":8080", nil)
}
