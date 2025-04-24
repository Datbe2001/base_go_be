package user

import (
	"base_go_be/internal/middlewares"
	"base_go_be/internal/wire"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productController, _ := wire.InitProductRouterHandler()
	productRouterPublic := Router.Group("/product")
	productRouterPublic.Use(middlewares.AuthMiddleware())
	{
		productRouterPublic.GET("/detail/:id", productController.GetProductByID)
		productRouterPublic.GET("/list", productController.GetListProduct)
		productRouterPublic.POST("/create", productController.CreateProduct)
	}

	//private router
}
