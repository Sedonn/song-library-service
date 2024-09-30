package postgresql

import (
	"fmt"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/sedonn/song-library-service/internal/config"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/services/song"
)

// Repository содержит методы взаимодействия с базой данных PostgreSQL.
type Repository struct {
	db *gorm.DB
}

var _ song.SongProvider = (*Repository)(nil)
var _ song.SongSaver = (*Repository)(nil)

var _ song.SongUpdater = (*Repository)(nil)
var _ song.SongDeleter = (*Repository)(nil)

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

// withSearchByStringAttributes добавляет поиск по подстроке для всех указанных строковых атрибутов модели.
func withSearchByStringAttributes(model any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		mType := reflect.TypeOf(model)
		mValue := reflect.ValueOf(model)

		for i := range mType.NumField() {
			field := mType.Field(i)
			if field.Type.Kind() != reflect.String {
				continue
			}

			fieldValue := mValue.Field(i).String()
			if fieldValue == "" {
				continue
			}

			column := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")["COLUMN"]
			db.Where(fmt.Sprintf("%q ILIKE '%%%s%%'", column, fieldValue))
		}

		return db
	}
}

// makeDSN создает строку подключения к базе данных на основе текущей конфигурации.
func makeDSN(cfg *config.DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
}
