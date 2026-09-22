package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("no arguments required")
	}

	feeds, err := s.db.GetFeedsWithUser(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("feed name:	 %+v\n", feed.Name)
		fmt.Printf("feed URL: 	 %+v\n", feed.Url)
		fmt.Printf("user name: 	%+v\n", feed.UserName)
		fmt.Println()
	}

	fmt.Printf("feeds have been listed\n")

	return nil
}
