package github

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Repository struct {
	AccessToken string
}

func (r *Repository) GetContributions(ctx context.Context, userName string, from time.Time, to time.Time) ([]int, error) {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: r.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	var q struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Weeks []struct {
						ContributionDays []struct{ ContributionCount githubv4.Int }
					}
				}
			} `graphql:"contributionsCollection(from:$from,to:$to)"`
		} `graphql:"user(login:$userName)"`
	}

	variables := map[string]interface{}{
		"from":     githubv4.DateTime{from},
		"to":       githubv4.DateTime{to},
		"userName": githubv4.String(userName),
	}

	if err := client.Query(ctx, &q, variables); err != nil {
		return []int{}, err
	}

	// github end of slice maybe has int value 7 or less. So get week before last.
	// TODO: move this logic
	// TODO: リスクエスト内容によってコントリビューションのmetaレベルから変化するので仕様を確認する
	end := len(q.User.ContributionsCollection.ContributionCalendar.Weeks) - 1
	w := q.User.ContributionsCollection.ContributionCalendar.Weeks[end]
	wbl := q.User.ContributionsCollection.ContributionCalendar.Weeks[end-1]
	fmt.Printf("w: %v\n", w)
	fmt.Printf("wbl: %v\n", wbl)

	c := make([]int, 7)
	if 7 > len(w.ContributionDays) {

		for i := 0; i < 7-len(w.ContributionDays); i++ {
			c[i] = int(wbl.ContributionDays[len(w.ContributionDays)+i].ContributionCount)
		}

		for i := range w.ContributionDays {
			c[len(wbl.ContributionDays)-len(w.ContributionDays)+i] = int(wbl.ContributionDays[i].ContributionCount)
		}

	} else {

		for i := range w.ContributionDays {
			c[i] = int(w.ContributionDays[i].ContributionCount)
		}

	}

	return c, nil
}
