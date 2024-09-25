package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/sedonn/song-library-service/internal/app"
	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)
	log.Info("logger initialized", slog.String("env", cfg.Env))

	application := app.New(log, cfg)
	go application.RESTApp.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	application.RESTApp.Stop()
}
