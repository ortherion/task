package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Database struct {
	DB *mongo.Database
}

func New(ctx context.Context, connectionString string) (*Database, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("create mongo client failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //nolint:govet
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Database{DB: client.Database("mts")}, nil
}
