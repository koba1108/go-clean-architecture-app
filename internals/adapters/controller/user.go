package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/port"
	"time"
)

type User struct {
	OutputFactory func(ctx *gin.Context) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, ur repository.UserRepository) port.UserInputPort
	RepoFactory   func(c *gorm.DB) repository.UserRepository
	DB            *gorm.DB
}

type GetUserByIDRequest struct {
	UserID int `uri:"userId"`
}

type CreateUserRequest struct {
	DisplayName string    `json:"displayName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Birthday    time.Time `json:"birthday" time_format:"2006-01-02"`
}

func (cur *CreateUserRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateUserRequest
	aux := &struct {
		Birthday string `json:"birthday"`
		*Alias
	}{
		Alias: (*Alias)(cur),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	birthday, err := time.Parse("2006-01-02", aux.Birthday)
	if err != nil {
		return errors.New("birthday format is invalid")
	}
	cur.Birthday = birthday
	return nil
}

type UpdateUserRequest struct {
	UserID      int        `uri:"userId"`
	DisplayName *string    `json:"displayName"`
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	Birthday    *time.Time `json:"birthday" time_format:"2006-01-02"`
}

func (uur *UpdateUserRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateUserRequest
	aux := &struct {
		Birthday *string `json:"birthday"`
		*Alias
	}{
		Alias: (*Alias)(uur),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Birthday != nil {
		birthday, err := time.Parse("2006-01-02", *aux.Birthday)
		if err != nil {
			return errors.New("birthday format is invalid")
		}
		uur.Birthday = &birthday
	}
	return nil
}

type DeleteUserRequest struct {
	UserID int `uri:"userId"`
}

func (u *User) inputPortFactory(ctx *gin.Context) port.UserInputPort {
	outputPort := u.OutputFactory(ctx)
	repo := u.RepoFactory(u.DB)
	inputPort := u.InputFactory(outputPort, repo)
	return inputPort
}

func (u *User) GetUserAll(ctx *gin.Context) {
	inputPort := u.inputPortFactory(ctx)
	inputPort.GetUserAll()
}

func (u *User) GetUserByID(ctx *gin.Context) {
	var req GetUserByIDRequest
	inputPort := u.inputPortFactory(ctx)
	if err := ctx.BindUri(&req); err != nil {
		inputPort.ValidationError(err)
		return
	}
	inputPort.GetUserByID(req.UserID)
}

func (u *User) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	inputPort := u.inputPortFactory(ctx)
	if err := ctx.BindJSON(&req); err != nil {
		inputPort.ValidationError(err)
		return
	}
	inputPort.CreateUser(req.DisplayName, req.FirstName, req.LastName, req.Birthday)
}

func (u *User) UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest
	inputPort := u.inputPortFactory(ctx)
	if err := ctx.BindUri(&req); err != nil {
		inputPort.ValidationError(err)
		return
	}
	if err := ctx.BindJSON(&req); err != nil {
		inputPort.ValidationError(err)
		return
	}
	inputPort.UpdateUser(req.UserID, req.DisplayName, req.FirstName, req.LastName, req.Birthday)
}

func (u *User) DeleteUser(ctx *gin.Context) {
	var req DeleteUserRequest
	inputPort := u.inputPortFactory(ctx)
	if err := ctx.BindUri(&req); err != nil {
		inputPort.ValidationError(err)
		return
	}
	inputPort.DeleteUser(req.UserID)
}
