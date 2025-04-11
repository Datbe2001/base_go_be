# name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

mysql:
	docker restart mysql

swag:
	swag init -g ./cmd/${APP_NAME}/main.go -o docs
