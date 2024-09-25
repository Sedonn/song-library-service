package app

import (
	"log/slog"

	restapp "github.com/sedonn/song-library-service/internal/app/rest"
	"github.com/sedonn/song-library-service/internal/config"
)

// App это микросервис библиотеки песен.
type App struct {
	RESTApp *restapp.App
}

// New создает новый микросервис библиотеки песен.
func New(log *slog.Logger, cfg *config.Config) *App {
	restApp := restapp.New(log, cfg.REST.Port)

	return &App{
		RESTApp: restApp,
	}
}
