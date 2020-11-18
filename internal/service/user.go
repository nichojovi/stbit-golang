package service

import (
	"context"

	"github.com/nichojovi/stbit-test/internal/entity"
	"github.com/opentracing/opentracing-go"
)

func (us *userService) GetUserAuth(ctx context.Context, username, password string) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUserAuth")
	defer span.Finish()

	var user *entity.User
	user, err := us.userRepo.GetUserAuth(ctx, username, password)
	if err != nil {
		return user, err
	}

	return user, nil
}
