package song

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"strings"

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
	// Возвращает песни, общее количество найденных песен без учета пагинации, ошибку.
	Songs(ctx context.Context, attrs models.Song, p models.Pagination) (models.Songs, uint64, error)
}

// SongSaver описывает поведение объекта слоя данных, который обеспечивает сохранение данных песен.
type SongSaver interface {
	// SaveSong сохраняет данные нового сообщения.
	SaveSong(ctx context.Context, s models.Song) (models.Song, error)
}

// SongUpdater описывает поведение объекта слоя данных, который обеспечивает обновление данных песен.
type SongUpdater interface {
	// UpdateSong обновляет данные определенной песни.
	UpdateSong(ctx context.Context, s models.Song) (models.Song, error)
}

// SongDeleter описывает поведение объекта слоя данных, который обеспечивает удаление данных песен.
type SongDeleter interface {
	// DeleteSong удаляет данные определенной песни.
	DeleteSong(ctx context.Context, s models.Song) (models.Song, error)
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

// GetSongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
// Текст разбивается на куплеты по \n\n символам.
func (sl *SongLibrary) GetSongWithCoupletPagination(ctx context.Context, id uint64, p models.Pagination) (models.SongWithCoupletPaginationAPI, error) {
	log := sl.log.With(slog.Uint64("id", id))

	log.Info("attempt to get song")

	s, err := sl.songProvider.Song(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to provide song", logger.ErrorString(err))

			return models.SongWithCoupletPaginationAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to get song", logger.ErrorString(err))

		return models.SongWithCoupletPaginationAPI{}, err
	}

	couplets := strings.Split(s.Text, "\n\n")
	if len(couplets) < int(p.PageNumber) {
		return models.SongWithCoupletPaginationAPI{}, services.ErrPageNumberOutOfRange
	}

	s.Text = couplets[p.PageNumber-1]

	return models.SongWithCoupletPaginationAPI{
		Song: s.API(),
		Pagination: models.PaginationMetadataAPI{
			CurrentPageNumber: p.PageNumber,
			PageCount:         uint64(len(couplets)),
			PageSize:          1,
			RecordCount:       uint64(len(couplets)),
		},
	}, nil
}

// SearchSongs выполняет поиск песен по определенным параметрам.
func (sl *SongLibrary) SearchSongs(ctx context.Context, attrs models.Song, p models.Pagination) (models.SongsAPI, error) {
	sl.log.Info("attempt to search songs")

	songs, total, err := sl.songProvider.Songs(ctx, attrs, p)
	if err != nil {
		sl.log.Error("failed to search songs", logger.ErrorString(err))

		return models.SongsAPI{}, err
	}

	sl.log.Info("success to search songs", slog.Uint64("total", total))

	return models.SongsAPI{
		Songs: songs.API(),
		Pagination: models.PaginationMetadataAPI{
			CurrentPageNumber: p.PageNumber,
			PageCount:         uint64(math.Ceil(float64(total) / float64(p.PageSize))),
			RecordCount:       total,
			PageSize:          p.PageSize,
		},
	}, nil
}

// CreateSong создает новую песню.
func (sl *SongLibrary) CreateSong(ctx context.Context, s models.Song) (models.SongAPI, error) {
	log := sl.log.With(slog.String("name", s.Name), slog.String("group", s.Group))

	log.Info("attempt to create song")

	s, err := sl.songSaver.SaveSong(ctx, s)
	if err != nil {
		log.Error("failed to create song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to create song", slog.Uint64("id", s.ID))

	return s.API(), nil
}

// ChangeSong обновляет данные определенной песни.
func (sl *SongLibrary) ChangeSong(ctx context.Context, s models.Song) (models.SongAPI, error) {
	log := sl.log.With(slog.Uint64("id", s.ID))

	log.Info("attempt to change song")

	s, err := sl.songUpdater.UpdateSong(ctx, s)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to change song", logger.ErrorString(err))

			return models.SongAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to change song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to change song")

	return s.API(), nil
}

// RemoveSong удаляет определенную песню.
func (sl *SongLibrary) RemoveSong(ctx context.Context, s models.Song) (models.SongAPI, error) {
	log := sl.log.With(slog.Uint64("id", s.ID))

	log.Info("attempt to remove song")

	s, err := sl.songDeleter.DeleteSong(ctx, s)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to remove song", logger.ErrorString(err))

			return models.SongAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to remove song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to remove song")

	return s.API(), nil
}
