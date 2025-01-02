package main

import (
	"base_go_be/internal/routers"
)

func main() {
	r := routers.InitRouter()
	err := r.Run()
	if err != nil {
		return
	}
}
