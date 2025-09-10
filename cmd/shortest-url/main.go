package main

import (
	"log/slog"
	"os"

	"github.com/cmczk/shortest-url/internal/config"
	customLogger "github.com/cmczk/shortest-url/internal/http-server/middleware/logger"
	"github.com/cmczk/shortest-url/internal/lib/logger/sl"
	"github.com/cmczk/shortest-url/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting shortest url app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	_, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("cannot init storage", sl.Err(err))
		os.Exit(1)
	}

	log.Info("storage connected")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(customLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
