package storage

import (
	"alex_gorbunov_cv/internal/models"
)

func (storage *Storage) GetProjects() ([]*models.Project, error) {
	const fn = "storage.GetProjects"

	var projects []*models.Project

	for _, project := range storage.Projects {
		projects = append(projects, &project)
	}

	return projects, nil
}
