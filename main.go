package main

import (
	"fmt"

	"github.com/frankheinz87/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	err = cfg.SetUser("frank")
	if err != nil {
		fmt.Println("error setting user:", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Printf("%+v\n", cfg)
}
