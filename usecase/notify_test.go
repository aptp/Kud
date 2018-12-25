package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/aptp/Kud/adapter/repository/badger/slack"
	"github.com/aptp/Kud/adapter/repository/github"
	"github.com/aptp/Kud/usecase"
)

func TestNotifyTodaysContributions(t *testing.T) {

	ctx := context.Background()

	type args struct {
		ctx      context.Context
		userName string
	}

	type mockReturn struct {
		getContributionsErr        error
		contributions              []int
		makeToDaysContributionsErr error
		getWorkingCronChannelErr   error
		channelID                  string
	}

	var cases = []struct {
		err        error
		msg        string
		cid        string
		args       args
		mockReturn mockReturn
	}{
		{
			err: nil,
			msg: "natumn さんの今日のコントリビューションは 6 ! \n ここ１週間では 49 contributions。微妙ですね... 頑張ってください！\n",
			cid: "testestest",
			args: args{
				ctx:      ctx,
				userName: "natumn",
			},
			mockReturn: mockReturn{
				getContributionsErr:        nil,
				contributions:              []int{7, 7, 7, 7, 7, 8, 6},
				makeToDaysContributionsErr: nil,
				getWorkingCronChannelErr:   nil,
				channelID:                  "testestest",
			},
		},
	}

	for _, tt := range cases {

		githubMock := github.GitHubRepositoryMock{
			GetContributionsFunc: func(ctx context.Context, userName string, from time.Time, to time.Time) ([]int, error) {
				return tt.mockReturn.contributions, tt.mockReturn.getContributionsErr
			},
		}

		slackMock := slack.SlackRepositoryMock{
			GetWorkingCronChannelFunc: func(ctx context.Context) (string, error) {
				return tt.mockReturn.channelID, tt.mockReturn.getWorkingCronChannelErr
			},
		}

		msg, cid, err := usecase.NotifyTodaysContributions(tt.args.ctx, &githubMock, &slackMock, tt.args.userName)

		if tt.msg != msg {
			t.Fatalf("Test Failed expect: %s, but result: %s\n", tt.msg, msg)
		}

		if tt.cid != cid {
			t.Fatalf("test failed expect: %s, but result: %s\n", tt.cid, cid)
		}

		if tt.err != err {
			t.Fatalf("test failed expect: %s, but result: %s\n", tt.err, err)
		}
	}

}
