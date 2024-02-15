package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuea/todo-golang/internal/repository"
	todolistRepo "github.com/nuea/todo-golang/internal/repository/todolist"
)

type TodolistHandler interface {
	Create(ctx *gin.Context)
}

type handler struct {
	todolistRepo todolistRepo.TodolistRepository
}

// ProvideTodolist
func ProvideTodolist(repo *repository.Repository) TodolistHandler {
	return &handler{
		todolistRepo: repo.Todolist,
	}
}

func (h *handler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
