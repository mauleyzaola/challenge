package utils

import (
	"log"
	"os"
	"strconv"
)

func ReadPort() int {
	if portNumber := os.Getenv("PORT"); len(portNumber) != 0 {
		if val, err := strconv.Atoi(portNumber); err != nil {
			log.Fatal(err)
		} else {
			return val
		}
	}
	return DefaultPortNumber()
}

func DefaultPortNumber() int { return 8000 }
