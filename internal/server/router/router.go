package router

import (
	handlers "alex_gorbunov_cv/internal/server/handlers"
	store "alex_gorbunov_cv/internal/storage"
	"log/slog"
	"net/http"

	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(log *slog.Logger, storage store.Storage) http.Handler {
	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(gorilla_handlers.RecoveryHandler())

	mainComponentHandler, err := handlers.MainPageHandler(log, &storage)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	router.Handle("/", mainComponentHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return router
}
