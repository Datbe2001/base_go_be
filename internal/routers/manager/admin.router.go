package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.GET("/login")
	}

	//private router
	adminRouterPrivate := Router.Group("/admin")
	{
		adminRouterPrivate.PUT("/active_user")
	}
}
