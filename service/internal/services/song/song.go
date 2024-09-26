package song

import (
	"log/slog"

	"context"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	songrest "github.com/sedonn/song-library-service/internal/rest/handlers/song"
)

// SongSaver описывает поведение объекта, который обеспечивает сохранение данных песен.
type SongSaver interface {
	// SaveSong сохраняет данные нового сообщения.
	SaveSong(ctx context.Context, s models.Song) (uint64, error)
}

// Message предоставляет бизнес-логику работы с библиотекой песен.
type SongLibrary struct {
	log       *slog.Logger
	songSaver SongSaver
}

var _ songrest.SongLibraryManager = (*SongLibrary)(nil)

// New создает новый сервис для работы с сообщениями.
func New(log *slog.Logger, ss SongSaver) *SongLibrary {
	return &SongLibrary{
		log:       log,
		songSaver: ss,
	}
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
