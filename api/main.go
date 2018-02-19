package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	ctx  *context
	port int
)

func init() {
	ctx = newContext()
	if portNumber := os.Getenv("PORT"); len(portNumber) != 0 {
		if val, err := strconv.Atoi(portNumber); err != nil {
			log.Fatal(err)
		} else {
			port = val
		}
	} else {
		port = 8000 // default value
	}
	flag.IntVar(&port, "port", port, "specify the port number where the api will listen to requests")
}

func main() {
	flag.Parse()
	router := setupRouter()
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
	}
	server.ListenAndServe()
}
