package main

import (
	"go-gin-todo/pkg/api"
)

func main() {
	api := api.NewRouter()
	api.Run()
}
