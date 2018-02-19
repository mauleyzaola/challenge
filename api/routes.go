package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/basket", createBasket).Methods("POST")
	r.HandleFunc("/scan", scanProduct).Methods("POST")
	return r
}

func createBasket(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	id, err := ctx.CreateBasket()
	createResponse(w, &struct {
		Id string `json:"Id"`
	}{
		id,
	}, err)
}

func scanProduct(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	input := &struct {
		Codes []string
		Id    string
	}{}
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		createResponse(w, nil, err)
		return
	}
	items, err := ctx.ScanProduct(input.Id, input.Codes)
	if err != nil {
		createResponse(w, nil, err)
		return
	}
	createResponse(w, &items, err)
}
