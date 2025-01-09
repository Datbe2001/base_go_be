package initialize

import (
	"base_go_be/global"
	"fmt"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	sql := global.Config.Mysql
	logger := global.Config.Logger
	fmt.Println("Load config mysql:", sql.UserName, sql.Password, sql.DBName)
	fmt.Println("Load config logger:", logger.File_log_name, logger.Log_level)
	InitLogger()
	global.Logger.Info("hello hello", zap.String("ok", "success"))
	Mysql()
	Redis()

	r := InitRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
