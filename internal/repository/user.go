package repository

import (
	"context"
	"database/sql"

	"github.com/nichojovi/stbit-test/internal/entity"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	getAllUserInfoQuery = "SELECT id, user_name, password, full_name, email, phone FROM user"
)

func (ur *userRepo) GetUserAuth(ctx context.Context, username, password string) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetUserAuth")
	defer span.Finish()

	query := getAllUserInfoQuery + " where user_name = ? and password = ?"

	result := new(entity.User)
	err := ur.db.GetSlave().GetContext(ctx, result, query, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
