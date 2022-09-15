package repository

import (
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
)

type UserRepository interface {
	GetAll() ([]*model.User, error)
	GetByID(userID int) (*model.User, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	DeleteByID(userID int) error
}
