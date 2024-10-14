package internal

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
	RemoveSong(ctx context.Context, s models.Song) (models.SongAPI, error)
}

// NewRemoveSongHandler возвращает новый объект хендлера, который удаляет определенные песни.
//
//	@Summary		Удалить данные существующей песни.
//	@Description	Удалить данные существующей песни.
//	@Tags			song-library
//	@Accept			json
//	@Produce		json
//	@Param			song	path		models.SongIDAPI	true	"ID песни"
//	@Success		200		{object}	models.SongAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{id} [delete]
func NewRemoveSongHandler(sr SongRemover) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.SongIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sr.RemoveSong(ctx, models.Song{ID: req.ID})
		if err != nil {
			if errors.Is(err, services.ErrSongNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}
