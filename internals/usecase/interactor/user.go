package interactor

import (
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/port"
	"net/http"
	"time"
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

func (u *User) ValidationError(err error) {
	u.OutputPort.RenderError(http.StatusBadRequest, err)
}

func (u *User) GetUserAll() {
	users, err := u.UserRepo.GetAll()
	if err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.RenderList(users)
}

func (u *User) GetUserByID(userID int) {
	user, err := u.UserRepo.GetByID(userID)
	if err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.Render(user)
}

func (u *User) CreateUser(displayName, firstName, lastName string, birthday time.Time) {
	newUser, err := model.NewUser(displayName, firstName, lastName, birthday)
	if err != nil {
		u.OutputPort.RenderError(http.StatusBadRequest, err)
		return
	}
	user, err := u.UserRepo.Create(newUser)
	if err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.Render(user)
}

func (u *User) UpdateUser(userID int, displayName, firstName, lastName *string, birthday *time.Time) {
	user, err := u.UserRepo.GetByID(userID)
	if err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	if displayName != nil {
		user.DisplayName = *displayName
	}
	if firstName != nil {
		user.FirstName = *firstName
	}
	if lastName != nil {
		user.LastName = *lastName
	}
	if birthday != nil {
		user.Birthday = *birthday
	}
	user, err = u.UserRepo.Update(user)
	if err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.Render(user)
}

func (u *User) DeleteUser(userID int) {
	if err := u.UserRepo.DeleteByID(userID); err != nil {
		u.OutputPort.RenderError(http.StatusInternalServerError, err)
		return
	}
	u.OutputPort.RenderNoContent()
}
