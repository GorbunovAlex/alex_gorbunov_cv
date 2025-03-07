package router

import (
	templates "alex_gorbunov_cv/templates"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(log *slog.Logger) http.Handler {
	router := mux.NewRouter()

	// router.Use(mLogger.New(log))
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(handlers.RecoveryHandler())

	main_page_component := templates.MainPageComponent()

	router.Handle("/", templ.Handler(main_page_component))

	return router
}
