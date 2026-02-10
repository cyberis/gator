package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/cyberis/gator/internal/database"
)

func addFeedHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feed Name is required")
	}
	feedName := cmd.args[0]
	if len(cmd.args) < 2 {
		return fmt.Errorf("feed URL is required")
	}
	feedURL := cmd.args[1]

	// Get current user
	currentUserName := s.cfg.CurrentUserName
	if currentUserName == "" {
		return fmt.Errorf("no current user set, please login first")
	}

	// Check if user exists
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	// Add feed to database
	var createFeedParams database.CreateFeedParams
	createFeedParams.ID = uuid.New()
	createFeedParams.UserID = user.ID
	createFeedParams.CreatedAt = time.Now()
	createFeedParams.UpdatedAt = time.Now()
	createFeedParams.Name = feedName
	createFeedParams.Url = feedURL

	feed, err := s.db.CreateFeed(ctx, createFeedParams)
	if err != nil {
		return fmt.Errorf("failed to add feed: %v", err)
	}
	fmt.Printf("Feed created.\n%#v\n", feed)
	return nil
}
