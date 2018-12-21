package entity

import (
	"fmt"
	"time"
)

// WeekContributionsCount represents the number of daily contributions in a week.
type WeekContributionsCount [7]int

// ContributionPeriod return current date and one week ago in ISO-8601 encoded UTC date format.
func ContributionPeriod() (to string, from string) {

	to = fmt.Sprintf(time.Now().Format(time.RFC3339))
	from = fmt.Sprintf(time.Now().AddDate(0, 0, -7).Add(-(12 + 9) * time.Hour).Format(time.RFC3339))

	return
}
