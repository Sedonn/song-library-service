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

// Artist implements artist.ArtistProvider.
func (r *Repository) Artist(ctx context.Context, id uint64) (models.Artist, error) {
	var s models.Artist
	if err := r.db.WithContext(ctx).Take(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Artist{}, repository.ErrArtistNotFound
		}

		return models.Artist{}, err
	}

	return s, nil
}

// SaveArtist implements artist.ArtistSaver.
func (r *Repository) SaveArtist(ctx context.Context, a models.Artist) (models.Artist, error) {
	if tx := r.db.WithContext(ctx).Clauses(clause.Returning{}).Create(&a); tx.Error != nil {
		pgErr, ok := tx.Error.(*pgconn.PgError)
		if ok && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return models.Artist{}, repository.ErrArtistExists
		}

		return models.Artist{}, tx.Error
	}

	return a, nil
}

// UpdateArtist implements artist.ArtistUpdater.
func (r *Repository) UpdateArtist(ctx context.Context, a models.Artist) (models.Artist, error) {
	tx := r.db.WithContext(ctx).Clauses(clause.Returning{}).Updates(&a)
	if tx.Error != nil {
		pgErr, ok := tx.Error.(*pgconn.PgError)
		if ok && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return models.Artist{}, repository.ErrArtistExists
		}

		return models.Artist{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return models.Artist{}, repository.ErrArtistNotFound
	}

	return a, nil
}

// DeleteArtist implements artist.ArtistDeleter.
func (r *Repository) DeleteArtist(ctx context.Context, id uint64) (uint64, error) {
	tx := r.db.WithContext(ctx).Delete(models.Artist{ID: id})
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, repository.ErrArtistNotFound
	}

	return id, nil
}
