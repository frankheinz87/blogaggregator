package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %v\n", cmd.args[0])
	return nil
}
