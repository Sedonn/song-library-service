package postgresql

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/repository"
)

// Song возвращает данные определенной песни.
func (r *Repository) Song(ctx context.Context, id uint64) (models.Song, error) {
	var s models.Song
	if tx := r.db.WithContext(ctx).Take(&s, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return models.Song{}, repository.ErrSongNotFound
		}
	}

	return s, nil
}

// SaveSong сохраняет данные новой песни.
func (r *Repository) SaveSong(ctx context.Context, s models.Song) (uint64, error) {
	if tx := r.db.WithContext(ctx).Create(&s); tx.Error != nil {
		return 0, tx.Error
	}

	return s.ID, nil
}

// UpdateSong обновляет данные определенной песни.
func (r *Repository) UpdateSong(ctx context.Context, s models.Song) (models.Song, error) {
	tx := r.db.WithContext(ctx).Clauses(clause.Returning{}).Updates(&s)
	if tx.Error != nil {
		return models.Song{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return models.Song{}, repository.ErrSongNotFound
	}

	return s, nil
}
