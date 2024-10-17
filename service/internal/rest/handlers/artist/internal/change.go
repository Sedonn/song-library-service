package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// ArtistChanger описывает поведение объекта слоя бизнес-логики, который обновляет данные исполнителей.
type ArtistChanger interface {
	// ChangeSong обновляет данные определенного исполнителя.
	ChangeArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
}

type changeArtistRequest struct {
	models.ArtistIDAPI
	models.ArtistOptionalAttributesAPI
}

// NewChangeArtistHandler возвращает новый объект хендлера, который обновляет данные исполнителей.
//
//	@Summary		Изменить данные исполнителя.
//	@Description	Изменить данные исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		models.ArtistIDAPI	true	"ID исполнителя"
//	@Param			artist		body		changeArtistRequest	true	"Новые данные исполнителя"
//	@Success		200			{object}	models.ArtistAPI
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [put]
func NewChangeArtistHandler(ac ArtistChanger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req changeArtistRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ac.ChangeArtist(ctx, models.Artist{
			ID:   req.ID,
			Name: req.Name,
		})
		if err != nil {
			switch {
			case errors.Is(err, services.ErrArtistNotFound):
				fallthrough

			case errors.Is(err, services.ErrArtistExists):
				_ = ctx.AbortWithError(http.StatusBadRequest, err)

			default:
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
