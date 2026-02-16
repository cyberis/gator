package main

import (
	"fmt"
	"time"
)

func aggregateHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <time_between_requests>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid time format: %v", err)
	}

	// Start our Feed Fetching Loop
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	fmt.Printf("Starting feed aggregation with interval: %s\n", timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
