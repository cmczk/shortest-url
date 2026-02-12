package main

import (
	"log/slog"
	"os"

	"github.com/cmczk/shortest-url/internal/config"
	"github.com/cmczk/shortest-url/internal/lib/logger"
	"github.com/cmczk/shortest-url/internal/lib/logger/sl"
	"github.com/cmczk/shortest-url/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info("starting shortest-url", slog.String("env", cfg.Env))
	log.Debug("debug messages enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("cannot init storage", sl.Err(err))
		os.Exit(1)
	}

	log.Info("connection to db has been established")

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
}
