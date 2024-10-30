package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/services/artist"
	"github.com/sedonn/song-library-service/internal/services/song"
)

// Repository содержит методы взаимодействия с базой данных PostgreSQL.
type Repository struct {
	db *gorm.DB
}

var (
	_ song.SongProvider = (*Repository)(nil)
	_ song.SongSaver    = (*Repository)(nil)
	_ song.SongUpdater  = (*Repository)(nil)
	_ song.SongDeleter  = (*Repository)(nil)

	_ artist.ArtistProvider = (*Repository)(nil)
	_ artist.ArtistSaver    = (*Repository)(nil)
	_ artist.ArtistUpdater  = (*Repository)(nil)
	_ artist.ArtistDeleter  = (*Repository)(nil)
)

// New создает новый объект репозитория.
func New(cfg *config.Config) (*Repository, error) {
	dsn := makeDSN(&cfg.DB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.NewGORMLogger(cfg.Env),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Repository{db: db}, nil
}

// withPagination обеспечивает постраничную навигацию в результатах запроса.
func withPagination(p models.Pagination) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(p.PageSize)).Offset(int(p.PageNumber-1) * int(p.PageSize))
	}
}

// withSearchByStringColumn добавляет поиск по подстроке для определенного столбца определенной таблицы.
func withSearchByStringColumn(table, column, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}

		return db.Where(fmt.Sprintf("%q.%q ILIKE '%%%s%%'", table, column, value))
	}
}

// makeDSN создает строку подключения к базе данных на основе текущей конфигурации.
func makeDSN(cfg *config.DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
}
