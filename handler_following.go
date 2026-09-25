package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/frankheinz87/blogaggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 0 {
		return errors.New("no arguments required")
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feed_follows {
		fmt.Printf("feed name:	 %+v\n", feed.FeedName)
	}

	fmt.Printf("feeds have been listed\n")

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 1 {
		return errors.New("feed url required")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])

	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	return nil
}
