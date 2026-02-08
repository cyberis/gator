package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if commandFunc, exists := c.commands[cmd.name]; exists {
		return commandFunc(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func newCommands() *commands {
	return &commands{
		commands: make(map[string]func(*state, command) error),
	}
}
