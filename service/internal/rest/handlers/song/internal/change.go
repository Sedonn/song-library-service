package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongChanger описывает поведение объекта слоя бизнес-логики, который обновляет данные существующих песен.
type SongChanger interface {
	// ChangeSong обновляет данные определенной песни.
	ChangeSong(ctx context.Context, s models.Song) (models.SongAPI, error)
}

type changeSongRequest struct {
	models.SongIDAPI
	models.SongOptionalAttributesAPI
}

// NewChangeSongHandler возвращает новый объект хендлера, который обновляет существующие песни.
//
//	@Summary		Изменить данные существующей песни.
//	@Description	Изменить данные существующей песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song-library
//	@Accept			json
//	@Produce		json
//	@Param			song	path		models.SongIDAPI					true	"ID песни"
//	@Param			song	body		models.SongOptionalAttributesAPI	true	"Данные новой песни"
//	@Success		200		{object}	models.SongAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{id} [put]
func NewChangeSongHandler(sc SongChanger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req changeSongRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sc.ChangeSong(ctx, models.Song{
			ID:          req.ID,
			Name:        req.Name,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
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
