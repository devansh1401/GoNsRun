package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run main.go [run|child] <cmd> <args>")
	}

	switch os.Args[1] {
	case "run":
		fmt.Println("'run' command not yet implemented")
	case "child":
		fmt.Println("'child' command not yet implemented")
	default:
		panic("Invalid command. Available commands:\n" +
			"\t'run'    : Creates a new process in a containerized environment.\n" +
			"\t'child'  : Runs the specified command in the isolated environment.")
	}
}