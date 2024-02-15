package main

import (
	"github.com/nuea/todo-golang/cmd/todo-golang-api/di"
	"github.com/nuea/todo-golang/internal/config"
)

func main() {
	cfg := config.LoadAppConfig()

	ctn, err := di.InitializeContainer(cfg)
	if err != nil {
		panic(err)
	}
	ctn.Run()
}
