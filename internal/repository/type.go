package repository

import (
	"context"

	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/entity"
	"github.com/nichojovi/stbit-test/internal/utils/database"
)

type (
	userRepo struct {
		db  *database.Store
		cfg *config.MainConfig
	}
	movieRepo struct {
		db  *database.Store
		cfg *config.MainConfig
	}
)

type (
	UserRepository interface {
		GetUserAuth(ctx context.Context, username, password string) (*entity.User, error)
	}
	MovieRepository interface {
		InsertMovieLog(ctx context.Context, request entity.MovieDB) (int64, error)
	}
)
