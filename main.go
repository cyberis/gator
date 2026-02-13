package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
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
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", aggregateHandler)
	cmds.register("addfeed", addFeedHandler)
	cmds.register("feeds", listFeedsHandler)
	cmds.register("follow", followFeedHandler)
	cmds.register("following", listFollowingHandler)

	// Simulate command input
	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{name: cmdName, args: cmdArgs}

	// Run command
	if err := cmds.run(s, cmd); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}
