package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongChanger описывает поведение объекта слоя бизнес-логики, который обновляет данные песен.
type SongChanger interface {
	// ChangeSong обновляет данные определенной песни.
	ChangeSong(ctx context.Context, s models.Song) (models.SongAPI, error)
}

type changeSongRequest struct {
	models.SongIDAPI
	models.SongOptionalAttributesAPI
	Artist models.ArtistIDAPI `json:"artist" binding:"omitempty"`
}

// NewChangeSongHandler возвращает новый объект хендлера, который обновляет песни.
//
//	@Summary		Изменить данные песни.
//	@Description	Изменить данные песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song-id	path		models.SongIDAPI	true	"ID песни"
//	@Param			song	body		changeSongRequest	true	"Новые данные песни"
//	@Success		200		{object}	models.SongAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{song-id} [patch]
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
			ArtistID:    req.Artist.ID,
			Name:        req.Name,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
		if err != nil {
			switch {
			case errors.Is(err, services.ErrArtistNotFound):
				fallthrough

			case errors.Is(err, services.ErrSongNotFound):
				_ = ctx.AbortWithError(http.StatusBadRequest, err)

			default:
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}
