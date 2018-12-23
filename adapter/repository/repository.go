package repository

import "github.com/aptp/Kud/entity"

type Repository struct {
	GitHub entity.GitHubRepository
	Slack  entity.SlackRepository
}
