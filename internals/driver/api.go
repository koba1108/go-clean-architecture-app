package driver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/controller"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/gateway"
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/presenter"
	"github.com/koba1108/go-clean-architecture-app/internals/config"
	"github.com/koba1108/go-clean-architecture-app/internals/usecase/interactor"
	"log"
	"time"
)

func Serve(addr string) {
	db, err := config.NewGorm()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		DB:            db,
	}

	r := newGinApp()
	if err = r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}
	apiVi := r.Group("v1")
	{
		userGroup := apiVi.Group("/user")
		{
			userGroup.GET("", user.GetUserAll)
			userGroup.GET("/:userId", user.GetUserByID)
			userGroup.POST("", user.CreateUser)
			userGroup.PUT("/:userId", user.UpdateUser)
			userGroup.DELETE("/:userId", user.DeleteUser)
		}
	}
	if err = r.Run(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func newGinApp() *gin.Engine {
	r := gin.New()
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	conf.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	conf.AllowCredentials = true
	conf.MaxAge = 12 * time.Hour
	r.Use(cors.New(conf))
	return r
}
