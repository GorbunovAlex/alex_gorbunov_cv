package storage

import (
	"alex_gorbunov_cv/internal/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (storage *Storage) GetProjects() ([]*models.Project, error) {
	const fn = "storage.GetProjects"

	ctx := context.Background()

	var projects []*models.Project

	cursor, err := storage.collection.Find(ctx, bson.M{"projects": bson.M{"$exists": true}})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var document struct {
			Projects []*models.Project `bson:"projects"`
		}
		err := cursor.Decode(&document)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		projects = append(projects, document.Projects...)
	}

	return projects, nil
}
