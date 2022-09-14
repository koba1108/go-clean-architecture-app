package port

import (
	"context"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*model.User)
	RenderError(int, error)
}
