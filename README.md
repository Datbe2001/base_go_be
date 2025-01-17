# base_go_be
- run container: make mysql
- run project: make run
- call api: curl  http://localhost:8386/v1/user/test
- install wire: run go install github.com/google/wire/cmd/wire@latest
- cd to folder contain wire: run wire