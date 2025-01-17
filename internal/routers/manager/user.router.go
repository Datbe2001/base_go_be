package manager

import "github.com/gin-gonic/gin"

type UsersRouter struct{}

func (pr *UsersRouter) InitUserRouter(Router *gin.RouterGroup) {
	//private router
	usersRouterPrivate := Router.Group("/admin/user")
	{
		usersRouterPrivate.GET("/list_user")
	}
}
