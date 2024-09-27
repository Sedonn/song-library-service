package create

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

// SongCreator описывает поведение объекта слоя бизнес-логики, который добавляет новые песни.
type SongCreator interface {
	// CreateSong добавляют новую песню.
	CreateSong(ctx context.Context, s models.Song) (uint64, error)
}

type createSongResponse struct {
	ID uint64 `json:"id"`
}

// New возвращает новый объект хендлера, который добавляет новые песни.
func New(sc SongCreator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.SongAttributesAPI
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id, err := sc.CreateSong(ctx, models.Song{
			Name:        req.Name,
			Group:       req.Group,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, createSongResponse{ID: id})
	}
}
