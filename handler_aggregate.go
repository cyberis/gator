package main

import (
	"context"
	"fmt"
)

func aggregateHandler(s *state, cmd command) error {
	ctx := context.Background()
	rssFeed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %v", err)
	}

	fmt.Println("Latest Posts from Wagslane:")
	fmt.Printf("Fetch Feed\n%#v\n", rssFeed)
	return nil
}
