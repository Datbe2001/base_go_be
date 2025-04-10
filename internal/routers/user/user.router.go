package user

import (
	"base_go_be/internal/wire"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct{}

func (pr *UsersRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// this is non-dependency
	//ur := repo.NewUserRepository()
	//us := service.NewUserService(ur)
	//userHandlerNonDependency := controller.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()

	// WIRE go
	// Dependency Injection (DI) DI java
	usersRouterPublic := Router.Group("/user")
	{
		usersRouterPublic.GET("/register")
		usersRouterPublic.GET("/test/:id", userController.GetUserByID)
		usersRouterPublic.GET("/list_user", userController.GetListUser)
		usersRouterPublic.POST("/create_user", userController.CreateUser)
		usersRouterPublic.POST("/login")
	}

	//private router
	usersRouterPrivate := Router.Group("/user")
	{
		usersRouterPrivate.GET("/me")
	}
}
