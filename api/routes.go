package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/basket", createBasket).Methods("POST")
	return r
}

// TODO add headers
func createBasket(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	id, err := ctx.storage.Create()
	createResponse(w, &struct {
		Id string `json:"id"`
	}{
		id,
	}, err)
}
