package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/controller"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/gateway"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/presenter"
	"github.com/koba1108/go-clean-architecture-app/internals/config"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/interactor"
	"log"
)

func Serve(addr string) {
	db, err := config.NewMySQL()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		DB:            db,
	}

	server := gin.New()
	userGroup := server.Group("/user")
	{
		userGroup.GET("/:id", user.GetUserByID)
	}
	if err = server.Run(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
