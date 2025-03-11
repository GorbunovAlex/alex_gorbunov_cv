package main

import (
	"alex_gorbunov_cv/internal/config"
	"alex_gorbunov_cv/internal/lib/logger/sl"
	"alex_gorbunov_cv/internal/server/router"
	"alex_gorbunov_cv/internal/storage"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// TODO: Move to global? 3 invokations in app
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("env", cfg.Env))

	storage, err := storage.NewStorage()
	if err != nil {
		log.Error("failed to init storage", sl.Error(err))
		os.Exit(1)
	}

	log.Info("strating server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router.Router(log, *storage),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("server stopped", sl.Error(err))
	}
}
