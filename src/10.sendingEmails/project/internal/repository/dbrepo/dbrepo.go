package dbrepo

import (
	"database/sql"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(a *config.AppConfig, db *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  db,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
