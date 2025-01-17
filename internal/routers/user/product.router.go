package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/detail/:id")
		productRouterPublic.GET("/list")
	}

	//private router
}
