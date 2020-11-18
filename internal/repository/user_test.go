package repository

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nichojovi/stbit-test/cmd/config"
	"github.com/nichojovi/stbit-test/internal/entity"
	test_mock "github.com/nichojovi/stbit-test/internal/utils/test_mock"
	. "github.com/onsi/gomega"
)

var mockString string

// GetUserAuth
func TestGetUserAuth_success(t *testing.T) {
	row := []driver.Value{
		1,
		mockString,
		mockString,
		mockString,
		mockString,
		mockString,
	}
	g := NewGomegaWithT(t)
	found, err := setupGetUserAuth(g, func(exec *sqlmock.ExpectedQuery) {
		rows := test_mock.MockRowsFromQuery(expectedGetUserAuthQuery()).AddRow(row...)
		exec.WillReturnRows(rows)
	})
	result := &entity.User{
		ID:       1,
		UserName: mockString,
		Password: mockString,
		FullName: mockString,
		Email:    mockString,
		Phone:    mockString,
	}

	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(found).Should(Equal(result))

}

func TestGetUserAuth_returnErrWhenFetchDataError(t *testing.T) {
	g := NewGomegaWithT(t)
	found, err := setupGetUserAuth(g, func(exec *sqlmock.ExpectedQuery) {
		exec.WillReturnError(errors.New("error"))
	})
	var result *entity.User
	g.Expect(err).Should(HaveOccurred())
	g.Expect(found).Should(Equal(result))
}

func setupGetUserAuth(g *GomegaWithT, fn func(eq *sqlmock.ExpectedQuery)) (*entity.User, error) {
	testDB := test_mock.NewTestDBStore()
	defer testDB.DBSet.Close()
	fn(testDB.DBSet.Slave.Mock.ExpectQuery(expectedGetUserAuthQuery()))
	module := NewUserRepository(testDB.Store, &config.MainConfig{})
	result, err := module.GetUserAuth(context.Background(), mockString, mockString)
	notMetErr := testDB.DBSet.Slave.Mock.ExpectationsWereMet()
	g.Expect(notMetErr).ShouldNot(HaveOccurred())
	return result, err
}

func expectedGetUserAuthQuery() string {
	return test_mock.NormalizeRegexpQuery("SELECT id, user_name, password, full_name, email, phone FROM user where user_name = ? and password = ?")
}
