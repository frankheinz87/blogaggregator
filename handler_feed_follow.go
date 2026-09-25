package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/frankheinz87/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("feed url is required")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])

	if err != nil {
		return err
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("feed follow has been created\n")
	fmt.Printf("feed name:	 %+v\n", feedfollow.FeedName)
	fmt.Printf("user name: 	%+v\n", feedfollow.UserName)
	fmt.Println()

	return nil
}
