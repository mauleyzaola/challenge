package main

import (
	"fmt"
	"strings"
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
	fmt.Printf("result:%#v\n", result)
	return nil
}
