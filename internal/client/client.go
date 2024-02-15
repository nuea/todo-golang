package client

import (
	"github.com/google/wire"
	"github.com/nuea/todo-golang/internal/client/mongo"
)

// Client
// @@provide-struct@@
// @@no-locator-generation@@
type Client struct {
	MongoClient *mongo.MongoClient
}

var ClientSet = wire.NewSet(
	mongo.ProvideMongoClient,
)
