package repository

import (
	"context"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}
