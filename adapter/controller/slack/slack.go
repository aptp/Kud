package slack

import (
	"errors"

	"github.com/aptp/Kud/adapter/repository"
	"github.com/aptp/Kud/config"
	"github.com/nlopes/slack"
)

type Bot struct {
	cfg    *config.Config
	client *slack.Client
	// logger *log.Logger
	repository *repository.Repository
}

func NewBot(cfg *config.Config) *Bot {

	client := slack.New(cfg.Slack.APIToken)

	return &Bot{
		cfg:    cfg,
		client: client,
	}
}

func (s *Bot) Listen() error {

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

func (s *Bot) handleMessageEvent(ev *slack.MessageEvent) error {
	// TODO: impl
	return nil
}

func (s *Bot) HandleNotifyContributionsEvent(etype string) error {

	return nil
}
