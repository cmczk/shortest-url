package main

import (
	"log/slog"
	"os"

	"github.com/cmczk/shortest-url/internal/config"
	"github.com/cmczk/shortest-url/internal/lib/logger/sl"
	"github.com/cmczk/shortest-url/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("cannot init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	log.Info("starting shortest url app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
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
