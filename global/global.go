package global

import (
	"base_go_be/pkg/logger"
	"base_go_be/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LogZap
	Mysql  *gorm.DB
)

/*
Config: Redis, Mysql, ...
*/
