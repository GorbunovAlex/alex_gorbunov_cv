package storage

import (
	"alex_gorbunov_cv/internal/config"
	"context"
	"fmt"
	"log/slog"

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

func NewStorage(log *slog.Logger) (*Storage, error) {
	const fn = "storage.mongodb.NewStorage"

	cfg := config.MustLoad()

	credentianls := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}

	client_options := options.Client().ApplyURI(cfg.Database.ClientURI).SetAuth(credentianls)
	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	collection = client.Database(cfg.Database.Database).Collection(cfg.Collection)

	return &Storage{collection: collection}, nil
}
