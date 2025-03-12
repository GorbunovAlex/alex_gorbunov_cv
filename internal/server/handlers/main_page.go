package handlers

import (
	"alex_gorbunov_cv/web/templates"

	"github.com/a-h/templ"
)

func MainPageHandler() *templ.ComponentHandler {
	main_page_component := templates.MainPageComponent()
	return templ.Handler(main_page_component)
}
