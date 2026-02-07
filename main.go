package main

import (
	"fmt"
	"os"

	"github.com/frankheinz87/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	/*err = cfg.SetUser("frank")
	if err != nil {
		fmt.Println("error setting user:", err)
		return
	}*/

	st := &state{cfg: &cfg}

	cmds := commands{m: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	name := args[1]
	cmdArgs := []string{}
	if len(args) > 2 {
		cmdArgs = args[2:]
	}

	cmd := command{
		name: name,
		args: cmdArgs,
	}

	if err := cmds.run(st, cmd); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

}
