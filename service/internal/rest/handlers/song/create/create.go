package create

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

// SongCreator описывает поведение объекта, который добавляет новые песни.
type SongCreator interface {
	CreateSong(ctx context.Context, s models.Song) (uint64, error)
}

type request struct {
	Name        string `json:"name" binding:"required,lte=130"`
	Group       string `json:"group" binding:"required,lte=130"`
	ReleaseDate string `json:"releaseDate" binding:"required,songreleasedate"`
	Text        string `json:"text" binding:"required"`
	Link        string `json:"link" binding:"required,url"`
}

type response struct {
	ID uint64 `json:"id"`
}

// New возвращает новый объект хендлера, который добавляет новые песни.
func New(m SongCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id, err := m.CreateSong(c, models.Song{
			Name:        req.Name,
			Group:       req.Group,
			ReleaseDate: req.ReleaseDate,
			Text:        req.Text,
			Link:        req.Link,
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, response{ID: id})
	}
}
