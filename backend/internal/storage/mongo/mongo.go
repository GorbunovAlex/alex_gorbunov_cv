package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

type Storage struct {
	db *mongo.Database
}

func NewStorage() (*Storage, error) {
	const fn = "storage.mongo.NewStorage"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	db := client.Database("cv")

	return &Storage{db: db}, nil
}
