package main

import (
	"base_go_be/internal/routers"
)

func main() {
	r := routers.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
