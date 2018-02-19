package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createResponse(w http.ResponseWriter, data interface{}, err error) {
	var (
		status  int
		message []byte
	)
	if err != nil {
		message, _ = json.Marshal(&struct {
			Error   bool   `json:"error"`
			Message string `json:"message"`
		}{
			true,
			err.Error(),
		})
		status = http.StatusBadRequest
	} else {
		if message, err = json.Marshal(data); err != nil {
			log.Printf("cannot marshal response message:%s", err)
			status = http.StatusInternalServerError
		}
		status = http.StatusOK
	}
	w.WriteHeader(status)
	if _, err = w.Write(message); err != nil {
		http.Error(w, fmt.Sprintf("cannot write http response message:%s", err), http.StatusInternalServerError)
	}
}

func createHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
