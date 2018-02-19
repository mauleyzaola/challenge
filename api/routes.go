package main

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/basket", createBasket).Methods("POST")
	r.HandleFunc("/scan", scanProduct).Methods("POST")
	r.HandleFunc("/total/{id}", totalBasket).Methods("GET")
	return r
}

func createBasket(w http.ResponseWriter, r *http.Request) {
	createHeaders(w)
	id, err := ctx.createBasket()
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
		createResponse(w, nil, fmt.Errorf("missing id in path, please use:/total/{id}"))
		return
	}
	basket, amount, err := ctx.totalAmount(id)
	result := &struct {
		Basket interface{} `json:"basket"`
		Amount float64     `json:"amount"`
	}{
		basket,
		amount,
	}
	createResponse(w, result, err)
}
