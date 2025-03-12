package storage

import (
	models "alex_gorbunov_cv/internal/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetExperices(storage *Storage) ([]*models.Experience, error) {
	const fn = "storage.GetExperices"

	var experiences []*models.Experience

	cursor, err := storage.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var experience models.Experience
		err := cursor.Decode(&experience)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		experiences = append(experiences, &experience)
	}

	return experiences, nil
}
