package main

import (
	"log/slog"
	"os"
)

var DefaultLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

func init() {
	slog.SetDefault(DefaultLogger)
	slog.Info("DefaultLogger initialized")
}
