package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/cyberis/gator/internal/database"
)

func addFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feed Name is required")
	}
	feedName := cmd.args[0]
	if len(cmd.args) < 2 {
		return fmt.Errorf("feed URL is required")
	}
	feedURL := cmd.args[1]

	ctx := context.Background()

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
	printFeedEntry(&feed, user.Name)
	fmt.Println()
	fmt.Println("=======================================================")

	// Follow the feed for the user
	var followFeedParams database.CreateFeedFollowParams
	followFeedParams.ID = uuid.New()
	followFeedParams.CreatedAt = time.Now()
	followFeedParams.UpdatedAt = time.Now()
	followFeedParams.UserID = user.ID
	followFeedParams.FeedID = feed.ID

	_, err = s.db.CreateFeedFollow(ctx, followFeedParams)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}
	fmt.Printf("Successfully followed feed %s for user %s\n", feed.Name, user.Name)
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
