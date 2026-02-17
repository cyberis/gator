package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cyberis/gator/internal/database"
	"github.com/google/uuid"
)

func insertPost(s *state, feedID uuid.UUID, title, link, description, pubDate string) (*database.Post, error) {
	ctx := context.Background()
	// Parse publication date
	var pubTime sql.NullTime
	pubTimeParsed, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		log.Printf("Failed to parse publication date '%s': %v", pubDate, err)
		pubTime = sql.NullTime{Time: time.Time{}, Valid: false} // Fallback to sql.NullTime if parsing fails
	} else {
		pubTime = sql.NullTime{Time: pubTimeParsed, Valid: true}
	}

	// Set Description to empty string if it's nil
	var descriptionNull sql.NullString
	if description == "" {
		descriptionNull = sql.NullString{String: "", Valid: false}
	} else {
		descriptionNull = sql.NullString{String: description, Valid: true}
	}

	// Check if post already exists based on URL
	existingPost, err := s.db.GetPostByURL(ctx, link)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Failed to check for existing post with URL '%s': %v", link, err)
		return nil, fmt.Errorf("failed to check for existing post: %v", err)
	}
	if existingPost.ID != uuid.Nil {
		log.Printf("Post with URL '%s' already exists, skipping insert", link)
		return &existingPost, errors.New("post already exists") // Return existing post if it already exists
	}

	var createPostParams database.CreatePostParams
	createPostParams.ID = uuid.New()
	createPostParams.CreatedAt = time.Now()
	createPostParams.UpdatedAt = time.Now()
	createPostParams.Title = title
	createPostParams.Url = link
	createPostParams.Description = descriptionNull
	createPostParams.PublishedAt = pubTime
	createPostParams.FeedID = feedID

	post, err := s.db.CreatePost(ctx, createPostParams)
	if err != nil {
		log.Printf("Failed to insert post with URL '%s': %v", link, err)
		return nil, fmt.Errorf("failed to insert post: %v", err)
	}
	return &post, nil
}
