package song

import (
	"context"
	"errors"
	"log/slog"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/repository"
	songrest "github.com/sedonn/song-library-service/internal/rest/handlers/song"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongSaver описывает поведение объекта слоя данных, который обеспечивает предоставление данных о песнях.
type SongProvider interface {
	// Song возвращает данные определенной песни.
	Song(ctx context.Context, id uint64) (models.Song, error)
	// Songs выполняет поиск песен по определенным параметрам.
	Songs(ctx context.Context, attrs models.Song) ([]models.Song, error)
}

// SongSaver описывает поведение объекта слоя данных, который обеспечивает сохранение данных песен.
type SongSaver interface {
	// SaveSong сохраняет данные нового сообщения.
	SaveSong(ctx context.Context, s models.Song) (uint64, error)
}

// SongUpdater описывает поведение объекта слоя данных, который обеспечивает обновление данных песен.
type SongUpdater interface {
	// UpdateSong обновляет данные определенной песни.
	UpdateSong(ctx context.Context, s models.Song) (models.Song, error)
}

// SongDeleter описывает поведение объекта слоя данных, который обеспечивает удаление данных песен.
type SongDeleter interface {
	// DeleteSong удаляет данные определенной песни.
	DeleteSong(ctx context.Context, s models.Song) (uint64, error)
}

// SongLibrary предоставляет бизнес-логику работы с библиотекой песен.
type SongLibrary struct {
	log          *slog.Logger
	songProvider SongProvider
	songSaver    SongSaver
	songUpdater  SongUpdater
	songDeleter  SongDeleter
}

var _ songrest.SongLibraryManager = (*SongLibrary)(nil)

// New создает новый сервис для работы с сообщениями.
func New(log *slog.Logger, sp SongProvider, ss SongSaver, su SongUpdater, sd SongDeleter) *SongLibrary {
	return &SongLibrary{
		log:          log,
		songProvider: sp,
		songSaver:    ss,
		songUpdater:  su,
		songDeleter:  sd,
	}
}

// GetSong возвращает данные определенной песни.
func (sl SongLibrary) GetSong(ctx context.Context, id uint64) (models.Song, error) {
	log := sl.log.With(slog.Uint64("id", id))

	log.Info("attempt to get song")

	s, err := sl.songProvider.Song(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to provide song", logger.ErrorString(err))

			return models.Song{}, services.ErrSongNotFound
		}

		log.Error("failed to get song", logger.ErrorString(err))

		return models.Song{}, err
	}

	log.Info("success to get song", slog.String("name", s.Name), slog.String("group", s.Group))

	return s, nil
}

// SearchSongs выполняет поиск песен по определенным параметрам.
func (sl *SongLibrary) SearchSongs(ctx context.Context, attrs models.Song) ([]models.Song, error) {
	sl.log.Info("attempt to search songs")

	songs, err := sl.songProvider.Songs(ctx, attrs)
	if err != nil {
		sl.log.Error("failed to search songs", logger.ErrorString(err))

		return nil, err
	}

	sl.log.Info("success to search songs", slog.Int("count", len(songs)))

	return songs, nil
}

// CreateSong создает новую песню.
func (sl *SongLibrary) CreateSong(ctx context.Context, s models.Song) (uint64, error) {
	log := sl.log.With(slog.String("name", s.Name), slog.String("group", s.Group))

	log.Info("attempt to create song")

	id, err := sl.songSaver.SaveSong(ctx, s)
	if err != nil {
		log.Error("failed to create song", logger.ErrorString(err))

		return 0, err
	}

	log.Info("success to create song", slog.Uint64("id", id))

	return id, nil
}

// ChangeSong обновляет данные определенной песни.
func (sl *SongLibrary) ChangeSong(ctx context.Context, s models.Song) (models.Song, error) {
	log := sl.log.With(slog.Uint64("id", s.ID))

	log.Info("attempt to change song")

	s, err := sl.songUpdater.UpdateSong(ctx, s)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to change song", logger.ErrorString(err))

			return models.Song{}, services.ErrSongNotFound
		}

		log.Error("failed to change song", logger.ErrorString(err))

		return models.Song{}, err
	}

	log.Info("success to change song")

	return s, nil
}

// RemoveSong удаляет определенную песню.
func (sl *SongLibrary) RemoveSong(ctx context.Context, s models.Song) (uint64, error) {
	log := sl.log.With(slog.Uint64("id", s.ID))

	log.Info("attempt to remove song")

	id, err := sl.songDeleter.DeleteSong(ctx, s)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to remove song", logger.ErrorString(err))

			return 0, services.ErrSongNotFound
		}

		log.Error("failed to remove song", logger.ErrorString(err))

		return 0, err
	}

	log.Info("success to remove song")

	return id, nil
}
