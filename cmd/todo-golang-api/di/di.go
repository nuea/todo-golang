//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal/handler"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal/router"
	"github.com/nuea/todo-golang/internal/client"
	"github.com/nuea/todo-golang/internal/config"
	"github.com/nuea/todo-golang/internal/repository"
)

type Container struct {
	cfg    *config.AppConfig
	server *internal.Server
	routes *router.Route
}

func (c *Container) Run() {
	c.server.Run()
}

var BaseSet = wire.NewSet(
	internal.ProvideServer,
	client.ClientSet,

	wire.Struct(new(Container), "*"),
	wire.Struct(new(client.Client), "*"),
	wire.Struct(new(handler.Handler), "*"),
	wire.Struct(new(repository.Repository), "*"),
)

var MainSet = wire.NewSet(
	BaseSet,
	repository.RepositorySet,
	router.ProvideRoute,
	handler.HandlerSet,
)

func InitializeContainer(cfg *config.AppConfig) (*Container, error) {
	wire.Build(MainSet)
	return &Container{}, nil
}
