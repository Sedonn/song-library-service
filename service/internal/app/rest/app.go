package restapp

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	songrest "github.com/sedonn/song-library-service/internal/rest/handlers/song"
	mwerror "github.com/sedonn/song-library-service/internal/rest/middleware/error"
	"github.com/sedonn/song-library-service/internal/rest/validators"
)

// App это REST-сервер.
type App struct {
	log        *slog.Logger
	httpServer *http.Server
	cfg        *config.RESTConfig
}

// New создает новый REST-сервер.
func New(log *slog.Logger, cfg *config.RESTConfig, s songrest.SongLibraryManager) *App {
	router := gin.Default()

	mustRegisterValidators()

	router.Use(mwerror.New())

	api := router.Group("api")
	{
		v1 := api.Group("/v1")
		{
			songrest.New(s).BindTo(v1)
		}
	}

	srv := &http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(cfg.Port)),
		Handler: router.Handler(),
	}

	return &App{
		log:        log,
		httpServer: srv,
		cfg:        cfg,
	}
}

// mustRegisterValidators регистрирует кастомные методы валидации.
func mustRegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("songreleasedate", validators.SongReleaseDate)
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
