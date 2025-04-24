package user

import (
	"base_go_be/internal/middlewares"
	"base_go_be/internal/wire"

	"github.com/gin-gonic/gin"
)

type UsersRouter struct{}

func (pr *UsersRouter) InitUserRouter(Router *gin.RouterGroup) {
	// WIRE go - get user controller with dependency injection
	userController, _ := wire.InitUserRouterHandler()

	// public router - no authentication required
	usersRouterPublic := Router.Group("/user")
	{
		usersRouterPublic.POST("/login", userController.Login)
		usersRouterPublic.POST("/register", userController.Register)
		usersRouterPublic.GET("/get_user/:id", userController.GetUserByID)
	}

	// private router - authentication required
	usersRouterPrivate := Router.Group("/user")
	usersRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		usersRouterPrivate.GET("/me", userController.GetCurrentUser)
		usersRouterPrivate.POST("/create_user", userController.CreateUser)
		usersRouterPrivate.GET("/list_user", userController.GetListUser)
	}

	// admin router - authentication and admin role required
	usersRouterAdmin := Router.Group("/admin")
	usersRouterAdmin.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("ADMIN", "SUPER_ADMIN"))
	{
		// Admin-only endpoints can be added here
	}
}
