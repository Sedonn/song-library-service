package logger

import (
	"log/slog"
	"os"
	"strings"

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

// cutFunctionWithPackage вырезает название функции и ее пакет из slog.Source.Function.
func cutFunctionWithPackage(f string) string {
	if i := strings.LastIndexByte(f, '/'); i != -1 {
		return f[i+1:]
	}

	return f
}
