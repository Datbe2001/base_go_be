package initialize

import (
	"base_go_be/global"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgxpool/v5"
	"go.uber.org/zap"
)

func checkErrPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		println(err)
		panic(err)
	}
}

func Postgres() {
	global.Logger.Info("Start connecting to PostgreSQL")
	p := global.Config.Postgres
	dsn := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	var s = fmt.Sprintf(dsn, p.Host, p.Port, p.UserName, p.Password, p.DBName, p.SSLMode)

	config, err := pgxpool.ParseConfig(s)
	checkErrPanic(err, "Parse PostgreSQL config failed")

	// Configure connection pool
	config.MaxConns = int32(p.MaxOpenConn)
	config.MinConns = int32(p.MaxIdleConn)
	config.MaxConnLifetime = time.Duration(p.ConnMaxLifeTime) * time.Second

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	checkErrPanic(err, "Initialize PostgreSQL database failed")

	// Test connection
	err = db.Ping(context.Background())
	checkErrPanic(err, "PostgreSQL ping failed")

	global.Postgres = db
	global.Logger.Info("PostgreSQL connect successfully!")

	createTables()
}

func createTables() {
	ctx := context.Background()

	// Create users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		role VARCHAR(20) NOT NULL CHECK (role IN ('ADMIN', 'USER')),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := global.Postgres.Exec(ctx, createUsersTable)
	checkErrPanic(err, "Create users table failed")

	// Create products table
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = global.Postgres.Exec(ctx, createProductsTable)
	checkErrPanic(err, "Create products table failed")

	global.Logger.Info("Tables created successfully")
}
