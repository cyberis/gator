package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cyberis/gator/internal/database"
)

func browsePostsHandler(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	rowLimit := 2
	rowOffset := 0
	if len(cmd.args) > 0 && cmd.args[0] == "--usage" {
		fmt.Printf("Usage: %s [--usage] [limit] [offset]\n", cmd.name)
		fmt.Println("--usage: Show this usage message")
		return nil
	}
	if len(cmd.args) > 0 {
		if limit, err := strconv.Atoi(cmd.args[0]); err == nil {
			rowLimit = limit
		}
	}
	if len(cmd.args) > 1 {
		if offset, err := strconv.Atoi(cmd.args[1]); err == nil {
			rowOffset = offset
		}
	}
	var getPostsForUserParams database.GetPostsForUserParams
	getPostsForUserParams.Name = s.cfg.CurrentUserName
	getPostsForUserParams.Limit = int32(rowLimit)
	getPostsForUserParams.Offset = int32(rowOffset)

	posts, err := s.db.GetPostsForUser(ctx, getPostsForUserParams)
	if err != nil {
		return fmt.Errorf("failed to retrieve posts: %v", err)
	}
	for _, post := range posts {
		fmt.Printf("Title:      %s\n", post.Title)
		fmt.Printf("URL:        %s\n", post.Url)
		description := post.Description
		if description.Valid {
			fmt.Printf("Description: %s\n", description.String)
		} else {
			fmt.Printf("Description: N/A\n")
		}
		pubTime := post.PublishedAt
		if pubTime.Valid {
			fmt.Printf("Published:  %s\n", pubTime.Time.Format(time.RFC1123))
		} else {
			fmt.Printf("Published:  N/A\n")
		}
		fmt.Printf("Feed:       %s\n", post.FeedName)
		fmt.Println("--------------------------------------------------")
	}
	return nil
}
