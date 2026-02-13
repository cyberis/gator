package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cyberis/gator/internal/database"
	"github.com/google/uuid"
)

func followFeedHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feed URL is required")
	}
	feedURL := cmd.args[0]

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

	// Get feed by URL
	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed by URL: %v", err)
	}

	// Check if feed exists
	if feed.ID == uuid.Nil {
		return fmt.Errorf("feed with URL %s does not exist", feedURL)
	}

	// Follow feed in database
	var followFeedParams database.CreateFeedFollowParams
	followFeedParams.ID = uuid.New()
	followFeedParams.CreatedAt = time.Now()
	followFeedParams.UpdatedAt = time.Now()
	followFeedParams.UserID = user.ID
	followFeedParams.FeedID = feed.ID

	follow, err := s.db.CreateFeedFollow(ctx, followFeedParams)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}
	fmt.Printf("Successfully followed feed %s for user %s\n", follow.FeedName, follow.UserName)
	return nil
}
