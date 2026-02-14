package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cyberis/gator/internal/database"
	"github.com/google/uuid"
)

func followFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feed URL is required")
	}
	feedURL := cmd.args[0]

	ctx := context.Background()

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

func listFollowingHandler(s *state, cmd command, user database.User) error {

	ctx := context.Background()

	// Get followed feeds from database
	followedFeeds, err := s.db.GetFeedFollowForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to get followed feeds: %v", err)
	}
	if len(followedFeeds) == 0 {
		fmt.Println("No followed feeds found.")
		return nil
	}

	// Print followed feeds
	fmt.Printf("Feeds followed by %s:\n\n", user.Name)
	for _, feed := range followedFeeds {
		fmt.Printf("* %s (URL: %s)\n", feed.FeedName, feed.FeedUrl)
	}
	return nil
}
