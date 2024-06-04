package mongodb

import (
	"context"
	"os"

	"github.com/natansa/go-leilao/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_BD  = "MONGODB_BD"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		logger.Error("error connecting to mongodb", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("error pinging to mongodb", err)
		return nil, err
	}

	var mongoDatabase string
	mongoDatabase = os.Getenv(MONGODB_BD)
	return client.Database(mongoDatabase), nil
}
