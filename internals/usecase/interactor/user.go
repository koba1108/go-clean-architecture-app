package interactor

import (
	"context"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/port"
	"log"
	"net/http"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   repository.UserRepository
}

func NewUserInputPort(outputPort port.UserOutputPort, userRepo repository.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepo,
	}
}

func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("err", err)
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.Render(user)
}
