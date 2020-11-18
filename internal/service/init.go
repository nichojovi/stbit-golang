package service

import (
	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/repository"
)

func NewUserService(user repository.UserRepository, cfg *config.MainConfig) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: user,
	}
}

func NewMovieService(movie repository.MovieRepository, cfg *config.MainConfig) MovieService {
	return &movieService{
		cfg:       cfg,
		movieRepo: movie,
	}
}
