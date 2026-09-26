package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/frankheinz87/blogaggregator/internal/database"
	"github.com/google/uuid"
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

		var publishedAt sql.NullTime
		for _, layout := range []string{time.RFC1123Z, time.RFC1123} {
			if t, err := time.Parse(layout, item.PubDate); err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
				break
			}
		}

		var description sql.NullString
		description.String = item.Description
		if description.String != "" {
			description.Valid = true
		} else {
			description.Valid = false
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      next.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
				continue
			} else {
				log.Println("Error occured:", err)
			}
		}
	}

	return nil
}
