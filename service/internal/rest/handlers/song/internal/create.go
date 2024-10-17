package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongCreator описывает поведение объекта слоя бизнес-логики, который добавляет новые песни.
type SongCreator interface {
	// CreateSong добавляют новую песню.
	CreateSong(ctx context.Context, s models.Song) (models.SongAPI, error)
}

type createSongRequest struct {
	models.SongAttributesAPI
	Artist models.ArtistIDAPI `json:"artist"`
}

// NewCreateSongHandler возвращает новый объект хендлера, который добавляет новые песни.
//
//	@Summary		Добавить новую песню.
//	@Description	Добавление новой песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	body		createSongRequest	true	"Данные новой песни"
//	@Success		200		{object}	models.SongAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/ [post]
func NewCreateSongHandler(sc SongCreator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createSongRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sc.CreateSong(ctx, models.Song{
			Name:        req.Name,
			ArtistID:    req.Artist.ID,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
		if err != nil {
			if errors.Is(err, services.ErrArtistNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, s)
	}
}
