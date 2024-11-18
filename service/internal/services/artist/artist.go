package artist

import (
	"context"
	"errors"
	"log/slog"

	artistrest "github.com/sedonn/song-library-service/internal/controllers/rest/artist"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/repositories"
	"github.com/sedonn/song-library-service/internal/services"
)

// ArtistSaver описывает поведение объекта слоя данных, который обеспечивает сохранение данных исполнителей.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=ArtistSaver
type ArtistSaver interface {
	// SaveArtist сохраняет данные нового исполнителя.
	SaveArtist(ctx context.Context, a models.Artist) (models.Artist, error)
}

// ArtistProvider описывает поведение объекта слоя данных, который обеспечивает предоставление данных о исполнителях.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=ArtistProvider
type ArtistProvider interface {
	// Artist возвращает данные определенного исполнителя.
	Artist(ctx context.Context, id uint64) (models.Artist, error)
}

// ArtistUpdater описывает поведение объекта слоя данных, который обеспечивает обновление данных о исполнителях.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=ArtistUpdater
type ArtistUpdater interface {
	// UpdateArtist обновляет данные определенного исполнителя.
	UpdateArtist(ctx context.Context, a models.Artist) (models.Artist, error)
}

// ArtistDeleter описывает поведение объекта слоя данных, который обеспечивает удаление данных о исполнителях.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=ArtistDeleter
type ArtistDeleter interface {
	// DeleteArtist удаляет данные определенного исполнителя.
	DeleteArtist(ctx context.Context, id uint64) (uint64, error)
}

type Service struct {
	log            *slog.Logger
	artistProvider ArtistProvider
	artistSaver    ArtistSaver
	artistUpdater  ArtistUpdater
	artistDeleter  ArtistDeleter
}

var _ artistrest.ArtistService = (*Service)(nil)

// New создает новый объект сервиса исполнителей.
func New(log *slog.Logger, as ArtistSaver, ap ArtistProvider, au ArtistUpdater, ad ArtistDeleter) *Service {
	return &Service{
		log:            log,
		artistProvider: ap,
		artistSaver:    as,
		artistUpdater:  au,
		artistDeleter:  ad,
	}
}

// CreateArtist создает нового исполнителя.
func (s *Service) CreateArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error) {
	log := s.log.With(slog.String("name", a.Name))

	log.Info("attempt to create artist")

	a, err := s.artistSaver.SaveArtist(ctx, a)
	if err != nil {
		if errors.Is(err, repositories.ErrArtistExists) {
			log.Warn("failed to create artist", logger.ErrorString(err))

			return models.ArtistAPI{}, services.ErrArtistExists
		}

		log.Error("failed to create artist", logger.ErrorString(err))

		return models.ArtistAPI{}, err
	}

	log.Info("success to create artist", slog.Uint64("id", a.ID))

	return a.API(), nil
}

// GetArtist возвращает нового исполнителя.
func (s *Service) GetArtist(ctx context.Context, id uint64) (models.ArtistAPI, error) {
	log := s.log.With(slog.Uint64("id", id))

	log.Info("attempt to get artist")

	a, err := s.artistProvider.Artist(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrArtistNotFound) {
			log.Warn("failed to provide artist", logger.ErrorString(err))

			return models.ArtistAPI{}, services.ErrArtistNotFound
		}

		log.Error("failed to get artist", logger.ErrorString(err))

		return models.ArtistAPI{}, err
	}

	return a.API(), nil
}

// ChangeArtist изменяет данные определенного исполнителя.
func (s *Service) ChangeArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error) {
	log := s.log.With(slog.Uint64("id", a.ID))

	log.Info("attempt to change artist")

	a, err := s.artistUpdater.UpdateArtist(ctx, a)
	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrArtistNotFound):
			log.Warn("failed to change artist", logger.ErrorString(err))
			return models.ArtistAPI{}, services.ErrArtistNotFound

		case errors.Is(err, repositories.ErrArtistExists):
			log.Warn("failed to change artist", logger.ErrorString(err))
			return models.ArtistAPI{}, services.ErrArtistExists

		default:
			log.Error("failed to change artist", logger.ErrorString(err))
			return models.ArtistAPI{}, err
		}
	}

	log.Info("success to change artist")

	return a.API(), nil
}

// RemoveArtist удаляет данные определенного исполнителя.
func (s *Service) RemoveArtist(ctx context.Context, id uint64) (models.ArtistIDAPI, error) {
	log := s.log.With(slog.Uint64("id", id))

	log.Info("attempt to remove artist")

	id, err := s.artistDeleter.DeleteArtist(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrArtistNotFound) {
			log.Warn("failed to remove artist", logger.ErrorString(err))

			return models.ArtistIDAPI{}, services.ErrArtistNotFound
		}

		log.Error("failed to remove artist", logger.ErrorString(err))

		return models.ArtistIDAPI{}, err
	}

	log.Info("success to remove artist")

	return models.ArtistIDAPI{ID: id}, nil
}
