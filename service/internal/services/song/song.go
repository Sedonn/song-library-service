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

// SongProvider описывает поведение объекта слоя данных, который обеспечивает предоставление данных о песнях.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=SongProvider
type SongProvider interface {
	// Song возвращает данные определенной песни.
	Song(ctx context.Context, id uint64) (models.Song, error)
	// Songs выполняет поиск песен по определенным параметрам.
	// Возвращает песни, общее количество найденных песен без учета пагинации, ошибку.
	Songs(ctx context.Context, attrs models.Song, p models.Pagination) (models.Songs, uint64, error)
}

// SongSaver описывает поведение объекта слоя данных, который обеспечивает сохранение данных песен.
type SongSaver interface {
	// SaveSong сохраняет данные новой песни.
	SaveSong(ctx context.Context, s models.Song) (models.Song, error)
}

// SongUpdater описывает поведение объекта слоя данных, который обеспечивает обновление данных песен.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=SongUpdater
type SongUpdater interface {
	// UpdateSong обновляет данные определенной песни.
	UpdateSong(ctx context.Context, s models.Song) (models.Song, error)
}

// SongDeleter описывает поведение объекта слоя данных, который обеспечивает удаление данных песен.
//
//go:generate go run github.com/vektra/mockery/v2@v2.46.1 --name=SongDeleter
type SongDeleter interface {
	// DeleteSong удаляет данные определенной песни.
	DeleteSong(ctx context.Context, s models.Song) (models.Song, error)
}

// Service предоставляет бизнес-логику работы с библиотекой песен.
type Service struct {
	log          *slog.Logger
	songProvider SongProvider
	songSaver    SongSaver
	songUpdater  SongUpdater
	songDeleter  SongDeleter
}

var _ songrest.SongService = (*Service)(nil)

// New создает новый объект библиотеки песен.
func New(log *slog.Logger, sp SongProvider, ss SongSaver, su SongUpdater, sd SongDeleter) *Service {
	return &Service{
		log:          log,
		songProvider: sp,
		songSaver:    ss,
		songUpdater:  su,
		songDeleter:  sd,
	}
}

// GetSongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
// Текст разбивается на куплеты по \n\n символам.
func (s *Service) GetSongWithCoupletPagination(ctx context.Context, id uint64, p models.Pagination) (models.SongWithCoupletPaginationAPI, error) {
	log := s.log.With(slog.Uint64("id", id))

	log.Info("attempt to get song")

	song, err := s.songProvider.Song(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to provide song", logger.ErrorString(err))

			return models.SongWithCoupletPaginationAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to get song", logger.ErrorString(err))

		return models.SongWithCoupletPaginationAPI{}, err
	}

	couplets := strings.Split(song.Text, "\n\n")
	if len(couplets) < int(p.PageNumber) {
		return models.SongWithCoupletPaginationAPI{}, services.ErrPageNumberOutOfRange
	}

	song.Text = couplets[p.PageNumber-1]

	return models.SongWithCoupletPaginationAPI{
		Song: song.API(),
		Pagination: models.PaginationMetadataAPI{
			CurrentPageNumber: p.PageNumber,
			PageCount:         uint64(len(couplets)),
			PageSize:          1,
			RecordCount:       uint64(len(couplets)),
		},
	}, nil
}

// SearchSongs выполняет поиск песен по определенным параметрам.
func (s *Service) SearchSongs(ctx context.Context, attrs models.Song, p models.Pagination) (models.SongsAPI, error) {
	s.log.Info("attempt to search songs")

	songs, total, err := s.songProvider.Songs(ctx, attrs, p)
	if err != nil {
		s.log.Error("failed to search songs", logger.ErrorString(err))

		return models.SongsAPI{}, err
	}

	s.log.Info("success to search songs", slog.Uint64("total", total))

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
func (s *Service) CreateSong(ctx context.Context, song models.Song) (models.SongAPI, error) {
	log := s.log.With(slog.String("name", song.Name))

	log.Info("attempt to create song")

	song, err := s.songSaver.SaveSong(ctx, song)
	if err != nil {
		log.Error("failed to create song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to create song", slog.Uint64("id", song.ID))

	return song.API(), nil
}

// ChangeSong обновляет данные определенной песни.
func (s *Service) ChangeSong(ctx context.Context, song models.Song) (models.SongAPI, error) {
	log := s.log.With(slog.Uint64("id", song.ID))

	log.Info("attempt to change song")

	song, err := s.songUpdater.UpdateSong(ctx, song)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to change song", logger.ErrorString(err))

			return models.SongAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to change song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to change song")

	return song.API(), nil
}

// RemoveSong удаляет определенную песню.
func (s *Service) RemoveSong(ctx context.Context, song models.Song) (models.SongAPI, error) {
	log := s.log.With(slog.Uint64("id", song.ID))

	log.Info("attempt to remove song")

	song, err := s.songDeleter.DeleteSong(ctx, song)
	if err != nil {
		if errors.Is(err, repository.ErrSongNotFound) {
			log.Warn("failed to remove song", logger.ErrorString(err))

			return models.SongAPI{}, services.ErrSongNotFound
		}

		log.Error("failed to remove song", logger.ErrorString(err))

		return models.SongAPI{}, err
	}

	log.Info("success to remove song")

	return song.API(), nil
}
