package slack

import "context"

type Repository struct {
	AccessToken string
	ProjectID   string
}

func (r *Repository) GetWorkingCronChannel(ctx context.Context) (string, error) {
	// TODO: impl

	return "testestest", nil
}
