package initialize

import (
	_ "base_go_be/docs"
	"base_go_be/global"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func handleErr(err error) {
	if err != nil {
		global.Logger.Error("Run app fail!", zap.Error(err))
		panic(err)
	}
}

func Run() {
	LoadConfig()
	InitLogger()
	//global.Logger.Info("check logger", zap.String("key", "value"))
	Mysql()
	Redis()

	r := InitRouter()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8386")
	handleErr(err)
}
