package storage

import (
	"alex_gorbunov_cv/internal/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

type Storage struct {
	collection *mongo.Collection
}

func NewStorage() (*Storage, error) {
	const fn = "storage.mongodb.NewStorage"

	cfg := config.MustLoad()

	client_options := options.Client().ApplyURI(cfg.Database.ClientURI)
	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	collection = client.Database(cfg.Database.Database).Collection(cfg.Database.Collection)

	return &Storage{collection: collection}, nil
}
