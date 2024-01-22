package main

import (
	"fmt"
	"os"
)

const (
	help = `These are SVCS commands:
config     Get and set a username.
add        Add a file to the index.
log        Show commit logs.
commit     Save changes.
checkout   Restore a file.`
	config   = "Get and set a username."
	add      = "Add a file to the index."
	log      = "Show commit logs."
	commit   = "Save changes."
	checkout = "Restore a file."
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(help)
		return
	}

	command := os.Args[1]
	switch command {
	case "--help":
		fmt.Println(help)
	case "config":
		fmt.Println(config)
	case "add":
		fmt.Println(add)
	case "log":
		fmt.Println(log)
	case "commit":
		fmt.Println(commit)
	case "checkout":
		fmt.Println(checkout)
	default:
		fmt.Printf("'%s' is not a SVCS command.", command)
	}
}
