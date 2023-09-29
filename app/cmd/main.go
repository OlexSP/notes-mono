package main

import (
	"github.com/OlexSP/notes-mono/internal/app"
	"github.com/OlexSP/notes-mono/internal/config"
	"github.com/OlexSP/notes-mono/pkg/logging"
	"log/slog"
	"os"
)

func main() {
	slog.Info("config initialization")
	cfg := config.GetConfig()

	logger := logging.SetupLogger(cfg.AppConfig.LogLevel)

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Error("cannot create app", logging.Err(err))
		os.Exit(1)
	}

	err = app.Run()
	if err != nil {
		logger.Error("cannot run app", logging.Err(err))
		os.Exit(1)
	}
}
