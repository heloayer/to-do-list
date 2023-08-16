package main

import (
	"log"

	"github.com/heloayer/todo-list/config"
	"github.com/heloayer/todo-list/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("сonfig error: %s", err)
	}
	app.Run(cfg)
}
