package routers

//import (
//	"base_go_be/internal/controller"
//	"base_go_be/internal/middlewares"
//	"github.com/gin-gonic/gin"
//)
//
//func InitRouter() *gin.Engine {
//	r := gin.Default()
//	r.Use(middlewares.AuthMiddleware())
//	v1 := r.Group("/api/v1")
//	{
//		v1.GET("/get_me/:name", controller.NewUserController().GetMe)
//	}
//	return r
//}
