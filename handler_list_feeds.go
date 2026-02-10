package main

import (
	"context"
	"fmt"
)

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
		fmt.Printf("* %s (%s) by %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}
