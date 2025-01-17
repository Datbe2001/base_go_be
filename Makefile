# name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

mysql:
	docker restart mysql

sudo_docker:
	sudo usermod -aG docker $USER
