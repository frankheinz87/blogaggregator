package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("no arguments required")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return err
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
