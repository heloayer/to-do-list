package main

import (
	"log"

	"github.com/heloayer/to-do-list/config"
	"github.com/heloayer/to-do-list/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("сonfig error: %s", err)
	}
	app.Run(cfg)
}
