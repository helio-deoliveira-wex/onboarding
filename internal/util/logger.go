package util

import (
	"log/slog"
	"os"
)

var AppLogger *slog.Logger

func LoggerConfig() {
	AppLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
