package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is application config
type Config struct {
	Mode  *Mode
	Slack *Slack
	Repo  *Repository
}

// Mode is application mode.
type Mode struct {
	// If Debug is true, output application debug logging.
	Debug bool `default:"false"`
}

// Slack has some token and key.
type Slack struct {
	APIToken           string `required:"true"`
	DisplayName        string `dafault:"クド"`
	WorkingCronChannel []string
}

// Repository is config to data are stored.
type Repository struct {
	Datastore *Datastore
	GitHub    *GitHub
}

// Datastore is google cloud datastore config
type Datastore struct {
	ProjectID string `required:"true"`
}

// GitHub has access token.
type GitHub struct {
	/*
		request the following scopes:
			user
			public_repo
			repo
			repo_deployment
			repo:status
			read:repo_hook
			read:org
			read:public_key
			read:gpg_key
	*/

	AccessToken string `required:"true"`
}

// Load read Config from .env file and environment variable.
func Load() (*Config, error) {

	var cfg Config
	var err error

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
