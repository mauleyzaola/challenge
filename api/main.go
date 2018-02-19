package main

import (
	"fmt"
	"net/http"
)

var ctx *context

func init() {
	ctx = newContext()
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx.storage.Create()
	fmt.Fprintln(w, "listado:", ctx.storage.List())
}
