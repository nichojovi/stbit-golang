package api

import (
	"encoding/json"
	"net/http"

	"github.com/nichojovi/stbit-test/internal/entity"
	"github.com/nichojovi/stbit-test/internal/utils/auth"
	"github.com/nichojovi/stbit-test/internal/utils/response"
	opentracing "github.com/opentracing/opentracing-go"
)

func (a *API) GetAllMovies(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "api.GetAllMovies")
	defer span.Finish()

	var request entity.MovieDB
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}
	user := auth.GetAuthDetailFromContext(ctx)
	request.UserID = user.ID

	movies, err := a.movieService.GetAllMovies(ctx, request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}

	return response.NewJSONResponse().SetData(movies)
}
