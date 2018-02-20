package main

import (
	"flag"

	"github.com/mauleyzaola/challenge/utils"
)

var activeBasket string

func init() {
	client = newClient()
	port = utils.ReadPort()
	flag.IntVar(&port, "port", port, "specify the port number where the api is listening to requests")
	flag.Parse()
}

func main() {
	for {
		readText()
	}
}
