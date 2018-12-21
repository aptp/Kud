package main

import (
	"log"
	"os"

	"github.com/aptp/Kud/adapter/controller"
	"github.com/aptp/Kud/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed on config.Load(): %s\n", err.Error())
	}

	slackBot := controller.NewSlackBot(cfg)

	os.Exit(
		run(slackBot),
	)
}

func run(slackBot controller.SlackBot) int {
	if err := slackBot.Listen(); err != nil {
		log.Printf("Error :%s", err.Error())
		return 1
	}

	return 0
}
