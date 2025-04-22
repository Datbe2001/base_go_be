# base_go_be
- install dependency: go mod tidy
- run container: make mysql
- run project: make run
- call api: curl  http://localhost:8386/v1/user/test
- install wire: go install github.com/google/wire/cmd/wire@latest
- cd to folder contain wire: wire
- run swagger: swag init
- visit swagger: http://localhost:8386/docs/index.html