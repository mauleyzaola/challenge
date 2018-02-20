package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/basket", createBasket).Methods("POST")
	r.HandleFunc("/basket/{id}/scan", scanProduct).Methods("POST")
	r.HandleFunc("/basket/{id}", totalBasket).Methods("GET")
	r.HandleFunc("/basket/{id}", removeBasket).Methods("DELETE")
	r.HandleFunc("/ping", ping).Methods("GET")
	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	createResponse(w, &struct {
		Message string `json:"message"`
	}{
		"pong",
	}, nil)
}

func createBasket(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	basket, err := ctx.createBasket()
	createResponse(w, basket, err)
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
	if len(input.Id) == 0 {
		input.Id = mux.Vars(r)["id"]
	}

	items, err := ctx.scanProduct(input.Id, input.Codes)
	if err != nil {
		createResponse(w, nil, err)
		return
	}
	createResponse(w, &items, err)
}

func totalBasket(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		createResponse(w, nil, fmt.Errorf("missing id in path, please use:/basket/{id}"))
		return
	}
	basket, err := ctx.totalAmount(id)
	createResponse(w, basket, err)
}

func removeBasket(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		createResponse(w, nil, fmt.Errorf("missing id in path, please use:/basket/{id}"))
		return
	}
	err := ctx.removeBasket(id)
	createResponse(w, nil, err)
}
