package initialize

import (
	"base_go_be/global"
	"base_go_be/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.DisableConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	//middleware
	//r.Use() //logging
	//r.Use() // cross
	//r.Use() // limiter global

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User
	MainGroup := r.Group("/v1")
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)

	}
	return r
}
