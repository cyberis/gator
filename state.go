package main

import (
	"database/sql"

	"github.com/cyberis/gator/internal/config"
	"github.com/cyberis/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func newState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		return nil, err
	}
	dbQueries := database.New(db)
	return &state{
		db:  dbQueries,
		cfg: cfg,
	}, nil
}
