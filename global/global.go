package global

import (
	"base_go_be/pkg/logger"
	"base_go_be/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LogZap
)

/*
Config: Redis, Mysql, ...
*/
