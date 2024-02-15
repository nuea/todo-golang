package repository

import (
	"github.com/google/wire"
	"github.com/nuea/todo-golang/internal/repository/todolist"
)

// Repository
// @@provide-struct@@
// @@no-locator-generation@@
type Repository struct {
	Todolist todolist.TodolistRepository
}

var RepositorySet = wire.NewSet(
	todolist.ProvideTodolistRepository,
)
