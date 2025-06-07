package global

import (
	"base_go_be/pkg/logger"
	"base_go_be/pkg/setting"

	"github.com/jackc/pgxpool/v5"
	"github.com/redis/go-redis/v9"
)

var (
	Config   setting.Config
	Logger   *logger.LogZap
	Redis    *redis.Client
	Postgres *pgxpool.Pool
)

/*
Config: Redis, Postgres, ...
*/
