package restapp

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	artistrest "github.com/sedonn/song-library-service/internal/rest/handlers/artist"
	songrest "github.com/sedonn/song-library-service/internal/rest/handlers/song"
	"github.com/sedonn/song-library-service/internal/rest/handlers/swagdocs"
	mwerror "github.com/sedonn/song-library-service/internal/rest/middleware/error"
)

// App это REST-сервер.
type App struct {
	log        *slog.Logger
	httpServer *http.Server
}

// New создает новый REST-сервер.
func New(log *slog.Logger, cfg *config.RESTConfig, as artistrest.ArtistService, ss songrest.SongService) *App {
	router := gin.Default()

	router.Use(mwerror.New())

	api := router.Group("api")
	{
		v1 := api.Group("/v1")
		{
			artistrest.New(as).BindTo(v1)
			songrest.New(ss).BindTo(v1)
		}
	}

	swagdocs.BindTo(router)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(cfg.Port)),
		Handler: router.Handler(),
	}

	return &App{
		log:        log,
		httpServer: srv,
	}
}

// MustRun запускает REST-API сервер. Паникует при ошибке.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run запускает REST-API сервер.
func (a *App) Run() error {
	a.log.Info("starting REST-API server", slog.String("address", a.httpServer.Addr))

	if err := a.httpServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

// Stop останавливает REST-API сервер.
func (a *App) Stop() {
	a.log.Info("shutting down REST-API server")

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		a.log.Error("failed to shut down REST-API server", logger.ErrorString(err))
	}

	a.log.Info("REST-API server is shut down")
}
