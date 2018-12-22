package main

import (
	"log"
	"os"

	"github.com/aptp/Kud/adapter/controller/slack"
	"github.com/aptp/Kud/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed on loading configure: %s\n", err.Error())
	}

	os.Exit(
		run(slack.NewBot(cfg)),
	)
}

func run(sb *slack.Bot) int {
	if err := sb.Listen(); err != nil {
		log.Printf("Error :%s", err.Error())
		return 1
	}

	return 0
}
