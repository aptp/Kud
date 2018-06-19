package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
)

type Env struct {
	SlackToken string `required:"true"`
}

func main() {
	// set config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var e Env
	err = envconfig.Process("klyuchan", &e)
	if err != nil {
		log.Fatal(err)
	}

	// application run
	api := slack.New(e.SlackToken)
	os.Exit(run(api))
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}
