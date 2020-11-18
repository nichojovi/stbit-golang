package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/entity"
	test_mock "github.com/nichojovi/stbit-test/internal/utils/test_mock"
	. "github.com/onsi/gomega"
)

// InsertMovieLog
func TestInsertMovieLog_Success(t *testing.T) {
	g := NewGomegaWithT(t)
	err := setupInsertMovieLog(g, func(eq *sqlmock.ExpectedExec) {
		eq.WillReturnResult(sqlmock.NewResult(1, 1))
	})
	g.Expect(err).ShouldNot(HaveOccurred())
}

func setupInsertMovieLog(g *GomegaWithT, fn func(ee *sqlmock.ExpectedExec)) error {
	testDB := test_mock.NewTestDBStore()
	defer testDB.DBSet.Close()
	fn(testDB.DBSet.Master.Mock.ExpectExec(expectedInsertMovieLogQuery()))

	module := NewMovieRepository(testDB.Store, &config.MainConfig{})
	_, err := module.InsertMovieLog(context.Background(), entity.MovieDB{})
	notMetErr := testDB.DBSet.Master.Mock.ExpectationsWereMet()
	g.Expect(notMetErr).ShouldNot(HaveOccurred())

	return err
}

func expectedInsertMovieLogQuery() string {
	return test_mock.NormalizeRegexpQuery("INSERT INTO movie_log(user_id, search_word, pagination) VALUES (?, ?, ?)")
}
