package entity

import (
	"time"
)

// WeekContributionsCount represents the number of daily contributions in a week.
type WeekContributionsCount [7]int

// ContributionPeriod return current date and one week ago in ISO-8601 encoded UTC date format.
func ContributionPeriod() (from time.Time, to time.Time) {

	utc, _ := time.LoadLocation("UTC")

	from = time.Now().AddDate(0, 0, -7).Add(-(12 + 9) * time.Hour).In(utc)
	// TODO: 現在の日付になるとなぜか1年分取れてしまう。
	to = time.Now().In(utc)

	return
}
