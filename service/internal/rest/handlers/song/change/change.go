package change

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
	ChangeSong(ctx context.Context, s models.Song) (models.Song, error)
}

type changeSongRequest struct {
	models.SongIDAPI
	Data models.SongOptionalAttributesAPI
}

// New возвращает новый объект хендлера, который обновляет существующие песни.
func New(sc SongChanger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req changeSongRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindJSON(&req.Data); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sc.ChangeSong(ctx, models.Song{
			ID:          req.ID,
			Name:        req.Data.Name,
			Group:       req.Data.Group,
			ReleaseDate: req.Data.ReleaseDate,
			Text:        req.Data.Text,
			Link:        req.Data.Link,
		})
		if err != nil {
			if errors.Is(err, services.ErrSongNotFound) {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}
