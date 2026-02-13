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
	fmt.Println("Feed added successfully:")
	fmt.Println()
	printFeedEntry(&feed, currentUserName)
	fmt.Println()
	fmt.Println("=======================================================")

	return nil
}

func listFeedsHandler(s *state, cmd command) error {

	// Get Feeds from database
	ctx := context.Background()
	feeds, err := s.db.GetAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	// Print feeds
	fmt.Println("Feeds:")
	for _, feed := range feeds {
		println()
		var feedData database.Feed
		feedData.ID = feed.ID
		feedData.CreatedAt = feed.CreatedAt
		feedData.UpdatedAt = feed.UpdatedAt
		feedData.Name = feed.Name
		feedData.Url = feed.Url
		printFeedEntry(&feedData, feed.UserName)
		println()
	}
	fmt.Println("=======================================================")
	return nil
}

func printFeedEntry(feed *database.Feed, userName string) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", userName)
}
