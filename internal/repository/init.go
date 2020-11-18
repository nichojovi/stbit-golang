package repository

import (
	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/utils/database"
)

func NewUserRepository(db *database.Store, config *config.MainConfig) UserRepository {
	return &userRepo{
		db:  db,
		cfg: config,
	}
}

func NewMovieRepository(db *database.Store, config *config.MainConfig) MovieRepository {
	return &movieRepo{
		db:  db,
		cfg: config,
	}
}
