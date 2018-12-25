package slack

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/aptp/Kud/adapter/repository"
	"github.com/aptp/Kud/config"
	"github.com/aptp/Kud/usecase"
	"github.com/carlescere/scheduler"
	"github.com/nlopes/slack"
)

type Bot struct {
	cfg        *config.Config
	client     *slack.Client
	logger     *log.Logger
	repository *repository.Repository
}

func NewBot(cfg *config.Config, repo *repository.Repository) *Bot {

	client := slack.New(cfg.Slack.APIToken)

	return &Bot{
		cfg:        cfg,
		client:     client,
		repository: repo,
	}
}

func (s *Bot) Listen() error {

	rtm := s.client.NewRTM()
	go rtm.ManageConnection()

	notifycontributionsCron := func() {
		s.handleNotifyContributionsEvent()
	}

	scheduler.Every().Day().At("22:00").Run(notifycontributionsCron)

	// This Event loop watch 2 event.
	// First, slack incoming Event. For example, post message.
	// Second, cron job.
	// TODO: implement cron library. And put in this loop?

	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				err := s.handleMessageEvent(rtm, ev)
				if err != nil {
					return err
				}
			case *slack.InvalidAuthEvent:
				return errors.New("Invalid credentials")
			}
		}
	}
}

func (s *Bot) handleMessageEvent(rtm *slack.RTM, ev *slack.MessageEvent) error {
	if strings.Contains(ev.Text, "クド、ヘルスチェック") {
		rtm.SendMessage(rtm.NewOutgoingMessage("はい！", ev.Channel))
	}

	return nil
}

func (s *Bot) handleNotifyContributionsEvent() {

	ctx := context.Background()

	// TODO: deal with multi user
	message, channelID, err := usecase.NotifyTodaysContributions(ctx, s.repository.GitHub, s.repository.Slack, "natumn")
	if err != nil {
		// TODO: s.logger.Errorf("Error on NotifyContributionsEvent: %s\n", err.Error())
		log.Printf("Error on handleNotifyContributionsEvent: %s", err.Error())
	}

	s.postCronMessage(channelID, message)
}

func (s *Bot) postCronMessage(channelID string, msg string) {

	s.client.PostMessage(channelID, msg, slack.PostMessageParameters{
		Username:    s.cfg.Slack.DisplayName,
		AsUser:      true,
		UnfurlMedia: true,
		UnfurlLinks: true,
	})

}
