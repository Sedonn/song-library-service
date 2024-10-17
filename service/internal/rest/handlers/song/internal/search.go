package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

type searchSongRequest struct {
	Name       string `form:"name"`
	ArtistName string `form:"artistName"`
	Link       string `form:"link"`
	Pagination models.Pagination
}

// NewGetHandler возвращает новый объект хендлера, который выполняет поиск песен по определенным параметрам.
//
//	@Summary		Поиск определенной песни.
//	@Description	Поиск определенной песни по всем атрибутам.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	query		searchSongRequest	true	"Настройки поиска."
//	@Success		200		{object}	models.SongsAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/ [get]
func NewSearchSongsHandler(sg SongGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req searchSongRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		songs, err := sg.SearchSongs(ctx,
			models.Song{
				Name: req.Name,
				Artist: models.Artist{
					Name: req.ArtistName,
				},
				Link: req.Link,
			},
			req.Pagination,
		)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, songs)
	}
}
