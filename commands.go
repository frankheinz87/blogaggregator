package main

import (
	"errors"

	"github.com/frankheinz87/blogaggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	m map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.m[cmd.name]

	if exists {
		return handler(s, cmd)
	}
	return errors.New("unknown command")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.m[name] = f
}
