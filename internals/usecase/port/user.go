package port

import (
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
	"time"
)

type UserInputPort interface {
	ValidationError(error)
	GetUserAll()
	GetUserByID(userID int)
	CreateUser(displayName, firstName, lastName string, birthday time.Time)
	UpdateUser(userID int, displayName, firstName, lastName *string, birthday *time.Time)
	DeleteUser(userID int)
}

type UserOutputPort interface {
	Render(*model.User)
	RenderList([]*model.User)
	RenderNoContent()
	RenderError(int, error)
}
