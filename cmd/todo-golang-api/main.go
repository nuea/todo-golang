package main

import (
	"fmt"

	"github.com/nuea/todo-golang/cmd/todo-golang-api/di"
	"github.com/nuea/todo-golang/internal/config"
)

func main() {
	cfg := config.LoadAppConfig()
	fmt.Println("cfg >>>", cfg)

	ctn, err := di.InitializeContainer(cfg)
	if err != nil {
		panic(err)
	}
	ctn.Run()
}
