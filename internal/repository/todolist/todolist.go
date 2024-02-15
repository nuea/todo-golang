package todolist

import (
	"github.com/nuea/todo-golang/internal/client"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "todolist"

type TodolistRepository interface {
}

type repository struct {
	coll *mongo.Collection
}

// ProvideTodolistRepository
func ProvideTodolistRepository(client *client.Client) TodolistRepository {
	return &repository{coll: client.MongoClient.GetCollection(collectionName)}
}
