package main

import (
	"github.com/cyberis/gator/internal/config"
)

type state struct {
	Config *config.Config
}

func newState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}
	return &state{
		Config: cfg,
	}, nil
}
