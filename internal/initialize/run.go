package initialize

import (
	"base_go_be/global"
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
	err := r.Run(":8386")
	handleErr(err)
}
