package app

import (
	"log/slog"

	restapp "github.com/sedonn/song-library-service/internal/app/rest"
	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/repository/postgresql"
	"github.com/sedonn/song-library-service/internal/services/song"
)

// App это микросервис библиотеки песен.
type App struct {
	RESTApp *restapp.App
}

// New создает новый микросервис библиотеки песен.
func New(log *slog.Logger, cfg *config.Config) *App {
	repository, err := postgresql.New(cfg)
	if err != nil {
		panic(err)
	}
	log.Info("database connected", slog.String("database", cfg.DB.Database))

	songLibraryService := song.New(log, repository, repository, repository)

	restApp := restapp.New(log, &cfg.REST, songLibraryService)

	return &App{
		RESTApp: restApp,
	}
}
