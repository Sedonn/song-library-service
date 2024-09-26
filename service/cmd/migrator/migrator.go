// Migrator это инструмент для выполнения и отката миграций в PostgreSQL базу данных.

package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// configPath должен содержать путь к файлу конфигурации (.yaml).
var configPath = flag.String("config_path", "", "Path to the .yaml config file.")

// migrationsPath должен содержать путь к папке с файлами миграций (.sql).
//
// Обязательный параметр.
//
// Файлы миграций должны быть в следующем формате:
// см. https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md
var migrationsPath = flag.String("migrations_path", "", "Path to a folder with migrations.")

// migrationsPath содержит название таблицы со с метаданными миграций.
//
// По умолчанию: schema_migrations.
var migrationsTable = flag.String("migrations_table", "schema_migrations", "Name of table with migrations data")

// migrateMode содержит режим мигратора (выполнить или откатить).
//
// Возможные значения: 'up' или 'down'.
// По умолчанию: 'up'.
var migrateMode = flag.String("mode", "up", "Mode of migrator, can be 'up' or 'down'")

// verboseLogger это флаг включения для verbose-режима логгера.
//
// По умолчанию: false.
var verboseLogger = flag.Bool("verbose", false, "Enable verbose mode in migrator logger.")

const (
	migrateUp   = "up"
	migrateDown = "down"
)

func main() {
	flag.Parse()
	if *configPath == "" {
		panic("config_path is empty: " + *configPath)
	}
	if *migrationsPath == "" {
		panic("migrations_path is empty: " + *migrationsPath)
	}

	cfg := config.MustLoadByPath(*configPath)
	log := logger.New(cfg.Env)

	sourceURL := fmt.Sprintf("file://%s", *migrationsPath)
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&x-migrations-table=%v",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database,
		*migrationsTable,
	)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		panic("failed to init golang-migrate: " + err.Error())
	}
	log.Info("golang-migrate initialized", slog.String("database", cfg.DB.Database))

	m.Log = &MigratorLogger{log: log, verbose: *verboseLogger}

	var migrateErr error
	switch *migrateMode {
	case migrateUp:
		migrateErr = m.Up()
	case migrateDown:
		migrateErr = m.Down()
	default:
		panic("unknown migrate mode: " + *migrateMode)
	}

	if migrateErr != nil {
		if errors.Is(migrateErr, migrate.ErrNoChange) {
			m.Log.Printf("%s", "no migrations to apply")

			return
		}

		panic("migrate error: " + migrateErr.Error())
	}

	m.Log.Printf("%s", "new migrations applied")
}

// MigratorLogger это логгер для golang-migrate.
type MigratorLogger struct {
	log     *slog.Logger
	verbose bool
}

var _ migrate.Logger = (*MigratorLogger)(nil)

// Printf выводит результаты логов.
func (l *MigratorLogger) Printf(format string, v ...any) {
	l.log.Info(fmt.Sprintf(format, v...))
}

// Verbose показывает включен ли verbose-режим логгера.
func (l *MigratorLogger) Verbose() bool {
	return l.verbose
}
