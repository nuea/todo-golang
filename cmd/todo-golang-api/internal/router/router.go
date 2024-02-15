package router

import (
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal/handler"
)

type Route struct{}

// ProvideRoute
func ProvideRoute(server *internal.Server, h *handler.Handler) *Route {
	route := server.Route()
	todolist := route.Group("/todolist")
	{
		todolist.POST("/create", h.Todolist.Create)
	}
	// admin := route.Group("/admin")
	// {
	// 	admin.POST("/create", h.Todolist.Create)
	// }
	return &Route{}
}
