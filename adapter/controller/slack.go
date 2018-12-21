package controller

import (
	"errors"

	"github.com/aptp/Kud/config"
	"github.com/nlopes/slack"
)

type SlackBot struct {
	cfg    *config.Config
	client *slack.Client
	// logger *log.Logger
	// repository *repository.Repository
}

func NewSlackBot(cfg *config.Config) *SlackBot {

	client := slack.New(cfg.Slack.APIToken)

	return &SlackBot{
		cfg:    cfg,
		client: client,
	}
}

func (s *SlackBot) Listen() error {

	rtm := s.client.NewRTM()
	go rtm.ManageConnection()

	// This Event loop watch 2 event.
	// First, slack incoming Event. For example, post message.
	// TODO: Second, cron job.
	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				err := s.handleMessageEvent(ev)
				if err != nil {
					return err
				}
			case *slack.InvalidAuthEvent:
				return errors.New("Invalid credentials")
			}
		}
	}
}

func (s *SlackBot) handleMessageEvent(ev *slack.MessageEvent) error {
	// TODO: impl
	return nil
}

func (s *SlackBot) HandleCronEvent(etype string) error {
	// TODO: impl
	return nil
}
