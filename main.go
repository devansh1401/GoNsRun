package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run main.go [run|child] <cmd> <args>")
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		fmt.Println("'child' command not yet implemented")
	default:
		panic("Invalid command. Available commands:\n" +
			"\t'run'    : Creates a new process in a containerized environment.\n" +
			"\t'child'  : Runs the specified command in the isolated environment.")
	}
}

func run() {
	fmt.Printf("Running %v as pid %d (run)\n", os.Args[2:], os.Getpid())

	// This means the same program (the current executable) is restarted with the
	// child argument, simulating a new "containerized" process
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// This attaches the child process's input/output to the same terminal as the parent,
	// so the user can interact with it as if it were a normal command
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to start child process: %v", err))
	}
}