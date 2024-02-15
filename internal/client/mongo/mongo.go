package mongo

import (
	"context"

	"github.com/nuea/todo-golang/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	cfg      *config.AppConfig
	client   *mongo.Client
	database *mongo.Database
}

func (m *MongoClient) GetCollection(name string) (collection *mongo.Collection) {
	collection = m.database.Collection(name)
	if collection == nil {
		// Specify the options for creating the collection
		m.database.CreateCollection(context.Background(), name, options.CreateCollection().SetCapped(false))
	}
	return collection
}

func (m *MongoClient) GetCollectionWithDBName(dbName, name string) (collection *mongo.Collection) {
	collection = m.client.Database(dbName).Collection(name)
	if collection == nil {
		// Specify the options for creating the collection
		m.client.Database(dbName).CreateCollection(context.Background(), name, options.CreateCollection().SetCapped(false))
	}
	return collection
}

// ProvideMongoClient
func ProvideMongoClient(cfg *config.AppConfig) (*MongoClient, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opt := options.Client().ApplyURI(cfg.Mongo.URI).SetServerAPIOptions(serverAPI)

	// if tp != nil {
	// 	// Mongo OpenTelemetry instrumentation
	// 	clientOpt.Monitor = otelmongo.NewMonitor(otelmongo.WithCommandAttributeDisabled(false), otelmongo.WithTracerProvider(tp))
	// }
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		cfg:      cfg,
		client:   client,
		database: client.Database(cfg.Mongo.DBName),
	}, nil
}
