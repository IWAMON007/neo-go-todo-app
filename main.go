package main

import (
	"myapp/route"
	"net/http"
)

func main() {
	route.SetRoute()
	http.ListenAndServe(":8080", nil)
}
