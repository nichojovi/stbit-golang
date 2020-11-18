package internal

import (
	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/repository"
	"github.com/nichojovi/stbit-test/internal/service"
	"github.com/nichojovi/stbit-test/internal/utils/database"
)

func GetService(db *database.Store, config *config.MainConfig) *Service {
	//REPO
	userRepository := repository.NewUserRepository(db, config)
	movieRepository := repository.NewMovieRepository(db, config)

	//SERVICE
	userService := service.NewUserService(userRepository, config)
	movieService := service.NewMovieService(movieRepository, config)

	return &Service{
		User:  userService,
		Movie: movieService,
	}
}
