package initialize

import (
	"base_go_be/global"
	"base_go_be/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
