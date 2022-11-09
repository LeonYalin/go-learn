package dbrepo

import (
	"database/sql"
	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(a *config.AppConfig, db *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  db,
	}
}
