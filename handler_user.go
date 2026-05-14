package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/frankheinz87/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "user doesn't exist: %v\n", err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %v\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("user has been created\n")
	fmt.Printf("%+v\n", user)

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %v\n", cmd.args[0])
	return nil
}

func handlerList(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentusername := s.cfg.CurrentUserName

	for _, user := range users {
		if user.Name == currentusername {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}
