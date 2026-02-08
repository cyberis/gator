package main

import (
	"context"
	"fmt"

	"github.com/cyberis/gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	// Validate arguments
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}
	userName := cmd.args[0]

	ctx := context.Background()

	// Check if user already exists
	user, err := s.db.GetUser(ctx, userName)
	if err == nil && user.Name == userName {
		return fmt.Errorf("user %s already exists", userName)
	}

	// Create New user object
	newUser := database.CreateUserParams{
		ID:   uuid.New(),
		Name: userName,
	}

	// Register new user
	_, err = s.db.CreateUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}
	fmt.Println("User registered:", userName)

	// Set as current user
	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set current user: %v", err)
	}
	fmt.Println("Current user set to:", userName)
	return nil
}
