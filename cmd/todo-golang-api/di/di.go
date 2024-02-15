//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal"
	"github.com/nuea/todo-golang/internal/config"
)

type Container struct {
	cfg    *config.AppConfig
	server *internal.Server
}

func (c *Container) Run() {
	c.server.Run()
}

var ContainerSet = wire.NewSet(
	internal.ProvideServer,
)

var MainSet = wire.NewSet(
	ContainerSet,
	wire.Struct(new(Container), "*"),
)

func InitializeContainer(cfg *config.AppConfig) (*Container, error) {
	wire.Build(MainSet)
	return &Container{}, nil
}
