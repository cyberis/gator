package main

import (
	"log"
	"os"
)

func main() {
	// Initialize state
	s, err := newState()
	if err != nil {
		log.Fatalf("Failed to initialize state: %v", err)
	}

	// Initialize commands
	cmds := newCommands()
	cmds.register("login", handlerLogin)

	// Simulate command input
	if len(os.Args) < 2 {
		log.Fatalf("No command provided")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{name: cmdName, args: cmdArgs}

	// Run command
	if err := cmds.run(s, cmd); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}
