package global

import (
	"base_go_be/pkg/logger"
	"base_go_be/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LogZap
	Redis  *redis.Client
	Mysql  *gorm.DB
)

/*
Config: Redis, Mysql, ...
*/
