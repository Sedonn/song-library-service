package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

type searchSongRequest struct {
	Name        string `form:"name"`
	Group       string `form:"group"`
	ReleaseDate string `form:"releaseDate"`
	Text        string `form:"text"`
	Link        string `form:"link"`
}

// NewGetHandler возвращает новый объект хендлера, который выполняет поиск песен по определенным параметрам.
func NewSearchHandler(sg SongGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req searchSongRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		songs, err := sg.SearchSongs(ctx, models.Song{
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

		ctx.JSON(http.StatusOK, songs)
	}
}
