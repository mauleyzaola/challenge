package main

import (
	"fmt"
	"strings"

	"github.com/mauleyzaola/challenge/operations"
)

func parseCommand(input string) error {
	values := strings.Split(input, ":")
	cmd, ok := commands[values[0]]
	if !ok {
		return fmt.Errorf("invalid command:%s", values[0])
	}
	result, err := cmd(values[1:]...)
	if err != nil {
		return err
	}
	if printer, ok := result.(operations.Printer); ok {
		printer.Print()
	} else {
		fmt.Println(strings.Repeat("=", 40))
	}

	return nil
}
