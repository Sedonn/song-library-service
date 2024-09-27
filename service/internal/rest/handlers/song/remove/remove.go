package remove

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongRemover описывает поведение объекта слоя бизнес-логики, который удаляет песни.
type SongRemover interface {
	// RemoveSong удаляет определенную песню.
	RemoveSong(ctx context.Context, s models.Song) (uint64, error)
}

type removeSongResponse struct {
	ID uint64 `json:"id"`
}

// New возвращает новый объект хендлера, который удаляет определенные песни.
func New(sr SongRemover) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.SongIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id, err := sr.RemoveSong(ctx, models.Song{ID: req.ID})
		if err != nil {
			if errors.Is(err, services.ErrSongNotFound) {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, removeSongResponse{ID: id})
	}
}
