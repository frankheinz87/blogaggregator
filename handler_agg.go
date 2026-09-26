package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		return errors.New("time between reqs is required")
	}

	dur, err := time.ParseDuration(cmd.args[0])

	if err != nil {
		return err
	}

	ticker := time.NewTicker(dur)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {

	next, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), next.ID)

	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), next.Url)

	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("%v\n", item.Title)
	}

	return nil
}
