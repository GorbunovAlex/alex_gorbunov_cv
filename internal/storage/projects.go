package storage

import (
	"alex_gorbunov_cv/internal/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (storage *Storage) GetProjects() ([]*models.Project, error) {
	const fn = "storage.GetExperices"

	var projects []*models.Project

	cursor, err := storage.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var project models.Project
		err := cursor.Decode(&project)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		projects = append(projects, &project)
	}

	return projects, nil
}
