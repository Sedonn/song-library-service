package internal

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

// SongCreator описывает поведение объекта слоя бизнес-логики, который добавляет новые песни.
type SongCreator interface {
	// CreateSong добавляют новую песню.
	CreateSong(ctx context.Context, s models.Song) (models.SongAPI, error)
}

// NewCreateSongHandler возвращает новый объект хендлера, который добавляет новые песни.
//
//	@Summary		Добавить новую песню.
//	@Description	Добавление новой песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song-library
//	@Accept			json
//	@Produce		json
//	@Param			song	body		models.SongAttributesAPI	true	"Данные новой песни"
//	@Success		200		{object}	models.SongAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/ [post]
func NewCreateSongHandler(sc SongCreator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.SongAttributesAPI
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sc.CreateSong(ctx, models.Song{
			Name:        req.Name,
			Group:       req.Group,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, s)
	}
}
