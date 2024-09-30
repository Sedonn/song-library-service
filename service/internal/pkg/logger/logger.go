package logger

import (
	"io"
	"log/slog"
	"os"

	gormlog "gorm.io/gorm/logger"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/pkg/logger/handlers/prettyslog"
)

// New создает и настраивает объект логгера на основе типа текущего окружения.
func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = slog.New(prettyslog.New(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	}

	return log
}

// NewDiscardLogger создает заглушку-логгер, в которой отсутствует вывод логов.
func NewDiscardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
}

// NewGORMLogger создает и настраивает логгер GORM.
func NewGORMLogger(env string) gormlog.Interface {
	var level gormlog.LogLevel

	switch env {
	case config.EnvLocal:
		level = gormlog.Info
	}

	return gormlog.Default.LogMode(level)
}
