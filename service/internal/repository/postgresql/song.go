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
	if err := r.db.WithContext(ctx).Take(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Song{}, repository.ErrSongNotFound
		}
	}

	return s, nil
}

// SearchSongs выполняет поиск песен по определенным параметрам.
func (r *Repository) Songs(ctx context.Context, attrs models.Song, p models.Pagination) (models.Songs, uint64, error) {
	var (
		songs models.Songs
		total int64
	)

	err := r.db.
		WithContext(ctx).
		Model(models.Song{}).
		Scopes(withSearchByStringAttributes(attrs)).
		Count(&total).
		Scopes(withPagination(p)).
		Find(&songs).
		Error
	if err != nil {
		return models.Songs{}, 0, err
	}

	return songs, uint64(total), nil
}

// SaveSong сохраняет данные новой песни.
func (r *Repository) SaveSong(ctx context.Context, s models.Song) (models.Song, error) {
	if err := r.db.WithContext(ctx).Create(&s).Error; err != nil {
		return models.Song{}, err
	}

	return s, nil
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

// DeleteSong удаляет данные определенной песни.
func (r *Repository) DeleteSong(ctx context.Context, s models.Song) (models.Song, error) {
	tx := r.db.WithContext(ctx).Delete(&s)
	if tx.Error != nil {
		return models.Song{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return models.Song{}, repository.ErrSongNotFound
	}

	return s, nil
}
