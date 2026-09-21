package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/frankheinz87/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("feed name and url is required")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("feed has been created\n")
	fmt.Printf("%+v\n", feed)

	return nil
}
