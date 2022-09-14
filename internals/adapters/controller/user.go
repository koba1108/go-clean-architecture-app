package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/port"
)

type User struct {
	OutputFactory func(ctx *gin.Context) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, ur repository.UserRepository) port.UserInputPort
	RepoFactory   func(c *sql.DB) repository.UserRepository
	DB            *sql.DB
}

// GetUserByID は，httpを受け取り，portを組み立てて，inputPort.GetUserByIDを呼び出します．
func (u *User) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")
	outputPort := u.OutputFactory(ctx)
	repo := u.RepoFactory(u.DB)
	inputPort := u.InputFactory(outputPort, repo)
	inputPort.GetUserByID(ctx, userID)
}
