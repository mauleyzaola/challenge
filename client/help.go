package main

import "fmt"

func help(params ...string) (interface{}, error) {
	text := `
Command: quit
Syntax: quit
Exits the program
Example: quit

Command: help
Syntax: help
Displays this help
Example: help

Command: use
Syntax: use:[id]
Points to a different basket. We need to pass the id
Example: use:2

Command: create
Syntax: create
Creates a new basket and displays the generated id for it
Example: create

Command: scan
Syntax: scan:[code1]:[code2]:..:[code n]
Adds products to the basket we are pointing to
Example: scan:MUG:VOUCHER:VOUCHER:TSHIRT

Command: remove
Syntax: remove
Removes the basket
Example: remove

Command: total
Syntax: total
Displays the total amount for the basket we are pointing at and also, displays its items
Example: total
`
	fmt.Println(text)
	return nil, nil
}
