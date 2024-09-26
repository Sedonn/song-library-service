package get

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongGetter описывает поведение объекта, который извлекает данные библиотеки песен.
type SongGetter interface {
	GetSong(ctx context.Context, id uint64) (models.Song, error)
}

type request struct {
	ID uint64 `uri:"id" binding:"required"`
}

// New возвращает новый объект хендлера, который возвращает определенную песню.
func New(sg SongGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sg.GetSong(c, req.ID)
		if err != nil {
			if errors.Is(err, services.ErrSongNotFound) {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, s)
	}
}
