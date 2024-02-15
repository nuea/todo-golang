package handler

import (
	"github.com/google/wire"
	"github.com/nuea/todo-golang/cmd/todo-golang-api/internal/handler/todolist"
)

// Handler
// @@provide-struct@@
// @@no-locator-generation@@
type Handler struct {
	Todolist todolist.TodolistHandler
}

var HandlerSet = wire.NewSet(
	todolist.ProvideTodolist,
)
