package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// ArtistRemover описывает поведение объекта слоя бизнес-логики, который удаляет исполнителей.
type ArtistRemover interface {
	// RemoveSong удаляет определенного исполнителя.
	RemoveArtist(ctx context.Context, id uint64) (models.ArtistIDAPI, error)
}

// NewArtistRemoveHandler возвращает новый объект хендлера, который удаляет определенного исполнителя.
//
//	@Summary		Удалить данные исполнителя.
//	@Description	Удалить данные исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		models.ArtistIDAPI	true	"ID исполнителя"
//	@Success		200			{object}	models.ArtistIDAPI
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [delete]
func NewArtistRemoveHandler(ar ArtistRemover) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ArtistIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id, err := ar.RemoveArtist(ctx, req.ID)
		if err != nil {
			if errors.Is(err, services.ErrArtistNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, id)
	}
}
