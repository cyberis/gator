package main

import (
	"context"
	"fmt"
)

func usersHandler(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}
	// Get Current user
	currentUser := s.cfg.CurrentUserName

	// Print users
	fmt.Println("Users:")
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
