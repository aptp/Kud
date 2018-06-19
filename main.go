package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
)

// Env has some token and key.
type Env struct {
	SlackToken string `required:"true"`
}

var phrase = []string{
	"WELCOME!!!!!",
	"THANK YOU!!!!",
}

func main() {
	var e Env
	var err error

	// set config
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
				if strings.Contains(ev.Text, "遅れ") {
					rand.Seed(time.Now().UnixNano())
					rtm.SendMessage(rtm.NewOutgoingMessage(phrase[rand.Intn(2)], ev.Channel))
				}
			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}
