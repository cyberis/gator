package main

import (
	"context"
	"fmt"

	"github.com/cyberis/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		currentUserName := s.cfg.CurrentUserName
		if currentUserName == "" {
			return fmt.Errorf("no user is currently logged in")
		}
		user, err := s.db.GetUser(ctx, currentUserName)
		if err != nil {
			return fmt.Errorf("failed to retrieve user: %v", err)
		}
		return handler(s, cmd, user)
	}
}
