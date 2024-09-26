package postgresql

import (
	"context"

	"github.com/sedonn/song-library-service/internal/domain/models"
)

// SaveSong сохраняет данные новой песни.
func (r *Repository) SaveSong(ctx context.Context, s models.Song) (uint64, error) {
	if tx := r.db.WithContext(ctx).Create(&s); tx.Error != nil {
		return 0, tx.Error
	}

	return s.ID, nil
}
