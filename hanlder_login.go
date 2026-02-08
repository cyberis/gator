package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username argument is required")
	}
	userName := cmd.args[0]

	// Check if user exists
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, userName)
	log.Printf("GetUser error: %v", err)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("failed to get user: %v", err)
	}
	if err != nil && err.Error() == "sql: no rows in result set" {
		return fmt.Errorf("user %s does not exist", userName)
	}

	// Set current user in config
	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set current user: %v", err)
	}
	fmt.Println("User set to:", userName)
	return nil
}
