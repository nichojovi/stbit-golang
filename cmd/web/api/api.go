package api

import (
	"net/http"

	"github.com/nichojovi/stbit-test/cmd/internal"
	"github.com/nichojovi/stbit-test/internal/service"
	"github.com/nichojovi/stbit-test/internal/utils/auth"
	"github.com/nichojovi/stbit-test/internal/utils/response"
	"github.com/nichojovi/stbit-test/internal/utils/router"
)

type Options struct {
	Prefix         string
	DefaultTimeout int
	AuthService    auth.AuthService
	Service        *internal.Service
}

type API struct {
	options      *Options
	authService  auth.AuthService
	userService  service.UserService
	movieService service.MovieService
}

func New(o *Options) *API {
	return &API{
		options:      o,
		authService:  o.AuthService,
		userService:  o.Service.User,
		movieService: o.Service.Movie,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Timeout: a.options.DefaultTimeout, Prefix: a.options.Prefix})

	// Testing
	r.GET("/ping", a.Ping)

	// Movies
	r.GET("/movies", a.authService.Authorize(a.GetAllMovies))
}

func (a *API) Ping(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	return response.NewJSONResponse().SetMessage("pong")
}
