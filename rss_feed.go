package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cyberis/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// SSet a User-Agent header to avoid being blocked by some servers
	req.Header.Set("User-Agent", "Gator/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}

	// Unescape HTML entities in titles and descriptions
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	// Get the next feed to fetch based on last_fetched_at
	feed, err := s.db.GetNextFeedToFetch(ctx, int32(s.cfg.RefreshAtMins))
	if err != nil {
		// 1. Check if the error is "No Rows Found"
		if errors.Is(err, sql.ErrNoRows) {
			// This is often "expected" behavior (e.g., no feeds need refreshing)
			fmt.Println("No feeds are currently due for fetching")
			return
		}

		// 2. Otherwise, it's a real database error (connection, syntax, etc.)
		fmt.Printf("Failed to fetch next feed: %v\n", err)
		return
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Printf("Failed to fetch feed %s: %v\n", feed.Name, err)
		return
	}
	fmt.Printf("Fetched feed %s successfully\n", feed.Name)

	// Print feed Titles for demonstration purposes
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}
	log.Printf("Finished processing feed %s, %v posts found\n", feed.Name, len(rssFeed.Channel.Item))

	// Update last_fetched_at in database
	var markFeedFetchedParams database.MarkFeedFetchedParams
	markFeedFetchedParams.LastFetchedAt = sql.NullTime{Time: time.Now(), Valid: true}
	markFeedFetchedParams.ID = feed.ID

	err = s.db.MarkFeedFetched(ctx, markFeedFetchedParams)
	if err != nil {
		fmt.Printf("Failed to update last_fetched_at for feed %s: %v\n", feed.Name, err)
		return
	}
}
