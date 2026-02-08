package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to reset users: %v", err)
	}
	fmt.Println("All users have been reset.")
	return nil
}
