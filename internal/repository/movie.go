package repository

import (
	"context"

	"github.com/nichojovi/stbit-test/internal/entity"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	insertMovieLogQuery = "INSERT INTO movie_log(user_id, search_word, pagination) VALUES (?, ?, ?)"
)

func (mr *movieRepo) InsertMovieLog(ctx context.Context, request entity.MovieDB) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "roomRepo.InsertMovieLog")
	defer span.Finish()

	result, err := mr.db.GetMaster().ExecContext(ctx, insertMovieLogQuery,
		request.UserID,
		request.SearchWord,
		request.Pagination,
	)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
