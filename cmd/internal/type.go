package internal

import (
	"github.com/nichojovi/stbit-test/internal/service"
)

type (
	Service struct {
		User  service.UserService
		Movie service.MovieService
	}
)
