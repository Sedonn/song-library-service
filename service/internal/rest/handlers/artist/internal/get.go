package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongGetter описывает поведение объекта слоя бизнес-логики, который извлекает данные исполнителей.
type ArtistGetter interface {
	// GetSongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
	GetArtist(ctx context.Context, id uint64) (models.ArtistAPI, error)
}

// NewGetArtistHandler возвращает новый объект хендлера, который возвращает определенного исполнителя.
//
//	@Summary		Получить данные определенного исполнителя.
//	@Description	Получить данные определенного исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		models.ArtistIDAPI	true	"ID исполнителя"
//	@Success		200			{object}	models.ArtistAPI
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [get]
func NewGetArtistHandler(ag ArtistGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ArtistIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ag.GetArtist(ctx, req.ID)
		if err != nil {
			if errors.Is(err, services.ErrArtistNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
