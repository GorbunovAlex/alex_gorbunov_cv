package handlers

import (
	models "alex_gorbunov_cv/internal/models"
	"alex_gorbunov_cv/web/templates"
	"fmt"
	"log/slog"

	"github.com/a-h/templ"
)

type MainPageHadler interface {
	GetExperices() ([]*models.Experience, error)
	GetProjects() ([]*models.Project, error)
}

func MainPageHandler(log *slog.Logger, mainPageHandler MainPageHadler) (*templ.ComponentHandler, error) {
	const fn = "handlers.MainPageHandler"

	experiences, err := mainPageHandler.GetExperices()
	if err != nil {
		log.Error("%s: %v", fn, err)
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	projects, err := mainPageHandler.GetProjects()
	if err != nil {
		log.Error("%s: %v", fn, err)
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	main_page_component := templates.MainPageComponent(experiences, projects)
	return templ.Handler(main_page_component), nil
}
