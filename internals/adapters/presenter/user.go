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
	return &User{ctx: ctx}
}

func (u *User) Render(user *model.User) {
	u.ctx.JSON(http.StatusOK, user)
}

func (u *User) RenderList(users []*model.User) {
	if len(users) == 0 {
		u.ctx.JSON(http.StatusOK, make([]*model.User, 0))
	} else {
		u.ctx.JSON(http.StatusOK, users)
	}
}

func (u *User) RenderNoContent() {
	u.ctx.Status(http.StatusNoContent)
}

func (u *User) RenderError(status int, err error) {
	u.ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
