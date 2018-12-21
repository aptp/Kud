package github

import (
	"context"

	"github.com/aptp/Kud/entity"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubRepository struct {
	AccessToken string
}

func (r *GitHubRepository) GetContoributions(ctx context.Context, userName string, to string, from string) (*entity.WeekContributionsCount, error) {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: r.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	var q struct {
		user struct {
			contributionsCollection struct {
				contributionCalendar struct {
					weeks struct {
						contributionDays struct {
							contributionCount githubv4.Int
						}
					}
				}
			} `graphql:"contributionsCollection(from:$from,to:$to)"`
		} `graphql:"user(login:$userName)"`
	}

	variables := map[string]interface{}{
		"from":     githubv4.String(from),
		"to":       githubv4.String(to),
		"userName": githubv4.String(userName),
	}

	if err := client.Query(ctx, &q, variables); err != nil {
		return nil, err
	}

	return nil, nil
}
