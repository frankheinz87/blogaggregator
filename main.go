package main

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/frankheinz87/blogaggregator/internal/config"
	"github.com/frankheinz87/blogaggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("error opening database:", err)
		return
	}

	dbQueries := database.New(db)

	/*err = cfg.SetUser("frank")
	if err != nil {
		fmt.Println("error setting user:", err)
		return
	}*/

	st := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{m: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
