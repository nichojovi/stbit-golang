package service

import (
	"context"

	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/entity"
	"github.com/nichojovi/stbit-test/internal/repository"
)

type (
	userService struct {
		cfg      *config.MainConfig
		userRepo repository.UserRepository
	}
	movieService struct {
		cfg       *config.MainConfig
		movieRepo repository.MovieRepository
	}
)

type (
	UserService interface {
		GetUserAuth(ctx context.Context, username, password string) (*entity.User, error)
	}
	MovieService interface {
		GetAllMovies(ctx context.Context, request entity.MovieDB) ([]entity.Movie, error)
	}
)
