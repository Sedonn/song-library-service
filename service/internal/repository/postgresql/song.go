package postgresql

import (
	"context"
	"errors"
	"math"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/repository"
)

// SongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
func (r *Repository) SongWithCoupletPagination(ctx context.Context, id uint64, p models.PaginationAPI) (models.SongWithCoupletPaginationAPI, error) {
	var s models.SongAPI
	if err := r.db.WithContext(ctx).Model(models.Song{}).Take(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SongWithCoupletPaginationAPI{}, repository.ErrSongNotFound
		}
	}

	couplets := strings.Split(s.Text, "\n\n")
	if len(couplets) < int(p.PageNumber) {
		return models.SongWithCoupletPaginationAPI{}, repository.ErrPageNumberOutOfRange
	}

	s.Text = couplets[p.PageNumber-1]

	return models.SongWithCoupletPaginationAPI{
		Song: s,
		Pagination: models.PaginationMetadataAPI{
			CurrentPageNumber: p.PageNumber,
			PageCount:         uint64(len(couplets)),
			PageSize:          1,
			RecordCount:       uint64(len(couplets)),
		},
	}, nil
}

// SearchSongs выполняет поиск песен по определенным параметрам.
func (r *Repository) Songs(ctx context.Context, attrs models.Song, p models.PaginationAPI) (models.SongsAPI, error) {
	var (
		songs []models.SongAPI
		count int64
	)

	err := r.db.
		WithContext(ctx).
		Model(models.Song{}).
		Scopes(withSearchByStringAttributes(attrs)).
		Count(&count).
		Scopes(withPagination(p)).
		Find(&songs).
		Error
	if err != nil {
		return models.SongsAPI{}, err
	}

	return models.SongsAPI{
		Songs: songs,
		Pagination: models.PaginationMetadataAPI{
			CurrentPageNumber: p.PageNumber,
			PageCount:         uint64(math.Ceil(float64(count) / float64(p.PageSize))),
			RecordCount:       uint64(count),
			PageSize:          p.PageSize,
		},
	}, nil
}

// SaveSong сохраняет данные новой песни.
func (r *Repository) SaveSong(ctx context.Context, s models.Song) (uint64, error) {
	if err := r.db.WithContext(ctx).Create(&s).Error; err != nil {
		return 0, err
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

// DeleteSong удаляет данные определенной песни.
func (r *Repository) DeleteSong(ctx context.Context, s models.Song) (uint64, error) {
	tx := r.db.WithContext(ctx).Delete(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, repository.ErrSongNotFound
	}

	return s.ID, nil
}
