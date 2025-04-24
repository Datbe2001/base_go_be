package main

import "base_go_be/internal/initialize"

// @title Go API
// @version 1.0
// @description This is a sample server Go API server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @host localhost:8386
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description JWT Authorization header using Bearer scheme. Example: "Bearer {token}"
func main() {
	initialize.Run()
}
