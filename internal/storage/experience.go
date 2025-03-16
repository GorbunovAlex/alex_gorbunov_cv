package storage

import (
	models "alex_gorbunov_cv/internal/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (storage *Storage) GetExperices() ([]*models.Experience, error) {
	const fn = "storage.GetExperices"
	ctx := context.Background()

	var experiences []*models.Experience

	cursor, err := storage.collection.Find(ctx, bson.M{"experiences": bson.M{"$exists": true}})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var document struct {
			Experiences []*models.Experience `bson:"experiences"`
		}
		err := cursor.Decode(&document)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		experiences = append(experiences, document.Experiences...)
	}

	return experiences, nil
}
