package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/port"
	"net/http"
)

type User struct {
	ctx *gin.Context
}

func NewUserOutputPort(ctx *gin.Context) port.UserOutputPort {
	return &User{
		ctx: ctx,
	}
}

func (u *User) Render(user *model.User) {
	u.ctx.JSON(http.StatusOK, user)
}

func (u *User) RenderError(status int, err error) {
	_ = u.ctx.AbortWithError(status, err)
}
