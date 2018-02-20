package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/mauleyzaola/challenge/utils"
)

var (
	ctx  *context
	port int
)

func init() {
	ctx = newContext()
	port = utils.ReadPort()
	flag.IntVar(&port, "port", port, "specify the port number where the api will listen to requests")
	flag.Parse()
}

func main() {
	router := setupRouter()
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
	}
	server.ListenAndServe()
}
