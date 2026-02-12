package main

import (
	"log/slog"

	"github.com/cmczk/shortest-url/internal/config"
	"github.com/cmczk/shortest-url/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info("starting shortest-url", slog.String("env", cfg.Env))
	log.Debug("debug messages enabled")

	// TODO: init logger
	// TODO: init storage
	// TODO: init router
}
