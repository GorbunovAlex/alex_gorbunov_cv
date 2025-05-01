package storage

import (
	models "alex_gorbunov_cv/internal/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

type Storage struct {
	Experiences []models.Experience
	Projects    []models.Project
}

func NewStorage(log *slog.Logger) (*Storage, error) {
	const fn = "storage.local.NewStorage"
	type jsonData struct {
		Experiences []models.Experience `json:"experiences"`
		Projects    []models.Project    `json:"projects"`
	}

	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)

	// Build full path to the JSON file
	fullPath := filepath.Join(currentDir, "cv.json")

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var data jsonData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	storage := &Storage{
		Experiences: data.Experiences,
		Projects:    data.Projects,
	}

	return storage, nil
}
