package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nichojovi/stbit-test/internal/entity"
	opentracing "github.com/opentracing/opentracing-go"
)

func (ms *movieService) GetAllMovies(ctx context.Context, request entity.MovieDB) ([]entity.Movie, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "roomService.UpdateHost")
	defer span.Finish()

	resp, err := http.Get(ms.cfg.Omdb.OmdbUrl + fmt.Sprintf("?apikey=%s&s=%s&page=%d", ms.cfg.Omdb.OmdbKey, request.SearchWord, request.Pagination))
	if err != nil {
		log.Println("[movieService][GetAllMovies] Error when hit client omdb, err: ", err.Error())
		return []entity.Movie{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[movieService][GetAllMovies] Error when read body, err: ", err.Error())
		return []entity.Movie{}, err
	}

	var result entity.Movies
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("[movieService][GetAllMovies] Error when unmarshal, err: ", err.Error())
		return []entity.Movie{}, err
	}

	go func(request entity.MovieDB) {
		_, err = ms.movieRepo.InsertMovieLog(context.Background(), request)
		if err != nil {
			log.Println("[movieService][GetAllMovies] Error when InsertMovieLog, err: ", err.Error())
		}
	}(request)

	return result.Search, nil
}
