package storage

import (
	models "alex_gorbunov_cv/internal/models"
)

func (storage *Storage) GetExperices() ([]*models.Experience, error) {
	const fn = "storage.GetExperices"

	var experiences []*models.Experience

	for _, experience := range storage.Experiences {
		experiences = append(experiences, &experience)
	}

	return experiences, nil
}
