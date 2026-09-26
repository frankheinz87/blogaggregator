package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/frankheinz87/blogaggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2 // default value
	if len(cmd.args) == 1 {
		// an argument was provided, try to use it
		specifiedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = specifiedLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Found post: %s\n", post.Title)
		fmt.Printf("Description: %v\n", post.Description)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("published: %v\n", post.PublishedAt)
	}

	fmt.Printf("posts have been listed\n")

	return nil
}
