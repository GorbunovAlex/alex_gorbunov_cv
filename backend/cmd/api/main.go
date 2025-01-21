package main

import (
	"gorbunov_alex.space/cv/internal/config"
	"gorbunov_alex.space/cv/internal/lib/logger/sl"
	"gorbunov_alex.space/cv/internal/server/router"
	"gorbunov_alex.space/cv/internal/storage/mongo"

	"log/slog"
	"net/http"
	"os"
)

// @title           Swagger Alex CV API
// @version         1.0
// @description     API for my personal website.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alexander Gorbunov
// @contact.email  alex.gorbunov.de@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// TODO: Move to global? 3 invocations in the app
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("env", cfg.Env))

	storage, err := mongo.NewStorage()
	if err != nil {
		log.Error("failed to init storage", sl.Error(err))
		os.Exit(1)
	}

	log.Info("strating server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router.Router(log, storage),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("server stopped", sl.Error(err))
	}
}
