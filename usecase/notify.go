package usecase

import (
	"context"

	"github.com/aptp/Kud/entity"
)

// NotifyTodaysContributions is usecase for notify user contributions as cron job.
func NotifyTodaysContributions(ctx context.Context, gRepo entity.GitHubRepository, sRepo entity.SlackRepository, userName string) (string, string, error) {

	// TODO: pass location data as argment or not.
	from, to := entity.ContributionPeriod()
	contributions, err := gRepo.GetContributions(ctx, userName, from, to)
	if err != nil {
		return "", "", err
	}

	reply := entity.NewReply(userName, contributions)
	message, err := reply.MakeTodaysContributions()
	if err != nil {
		return "", "", err
	}

	channelID, err := sRepo.GetWorkingCronChannel(ctx)
	if err != nil {
		return "", "", err
	}

	return message, channelID, nil
}
