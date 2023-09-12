package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	Logger = logger

	Logger.Info("Logger initialized...")
}
