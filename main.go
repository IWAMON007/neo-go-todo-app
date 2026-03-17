package main

import (
	"fmt"
	"myapp/route"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	route.SetRoute()

	fmt.Println("Server Start !")
	http.ListenAndServe(":8080", nil)
}
