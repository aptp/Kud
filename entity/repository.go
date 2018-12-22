package entity

import (
	"context"
	"time"
)

type GitHubRepository interface {
	GetContributions(ctx context.Context, userName string, from time.Time, to time.Time) ([]int, error)
}
