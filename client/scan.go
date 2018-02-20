package main

import "fmt"

func readText() {
	var input string
	fmt.Scanln(&input)
	if err := parseCommand(input); err != nil {
		fmt.Printf("error:[%s]\n", err)
	}
}
