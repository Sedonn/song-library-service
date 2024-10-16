package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/repository"
)

// Song возвращает данные определенной песни.
func (r *Repository) Song(ctx context.Context, id uint64) (models.Song, error) {
	var s models.Song
	if err := r.db.WithContext(ctx).Preload(clause.Associations).Take(&s, id).Error; err != nil {
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
		Joins("Artist").
		Find(&songs).
		Error
	if err != nil {
		return models.Songs{}, 0, err
	}

	return songs, uint64(total), nil
}

// SaveSong сохраняет данные новой песни.
func (r *Repository) SaveSong(ctx context.Context, s models.Song) (models.Song, error) {
	err := r.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(&s).
		Joins("Artist").
		Take(&s).
		Error
	if err != nil {
		if isSongArtistNotFoundError(err) {
			return models.Song{}, repository.ErrArtistNotFound
		}

		return models.Song{}, err
	}

	return s, nil
}

// UpdateSong обновляет данные определенной песни.
func (r *Repository) UpdateSong(ctx context.Context, s models.Song) (models.Song, error) {
	tx := r.db.WithContext(ctx).Updates(&s)
	if tx.Error != nil {
		if isSongArtistNotFoundError(tx.Error) {
			return models.Song{}, repository.ErrArtistNotFound
		}

		return models.Song{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return models.Song{}, repository.ErrSongNotFound
	}

	if err := r.db.WithContext(ctx).Joins("Artist").Take(&s).Error; err != nil {
		return models.Song{}, nil
	}

	return s, nil
}

// DeleteSong удаляет данные определенной песни.
func (r *Repository) DeleteSong(ctx context.Context, id uint64) (uint64, error) {
	tx := r.db.WithContext(ctx).Delete(models.Song{ID: id})
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, repository.ErrSongNotFound
	}

	return id, nil
}

func isSongArtistNotFoundError(err error) bool {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return false
	}

	return pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) && pgErr.ConstraintName == "fk_songs_artist"
}
