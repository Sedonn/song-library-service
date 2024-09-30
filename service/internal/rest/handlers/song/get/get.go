package get

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// SongGetter описывает поведение объекта слоя бизнес-логики, который извлекает данные библиотеки песен.
type SongGetter interface {
	// GetSongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
	// Текст разбивается на куплеты по \n\n символам.
	GetSongWithCoupletPagination(ctx context.Context, id uint64, p models.Pagination) (models.SongWithCoupletPaginationAPI, error)
	// SearchSongs выполняет поиск песен по определенным параметрам.
	// Поиск выполняется по подстроке каждого указанного поля.
	SearchSongs(ctx context.Context, attrs models.Song, p models.Pagination) (models.SongsAPI, error)
}

type getSongRequest struct {
	Song       models.SongIDAPI
	Pagination models.Pagination
}

// NewGetHandler возвращает новый объект хендлера, который возвращает определенную песню.
//
//	@Summary		Получить данные определенной песни.
//	@Description	Получить данные определенной песни с пагинацией по куплетами.
//	@Tags			song-library
//	@Accept			json
//	@Produce		json
//	@Param			song	path		models.SongIDAPI	true	"ID песни"
//	@Param			song	query		models.Pagination	true	"Настройки пагинации. pageNumber игнорируется."
//	@Success		200		{object}	models.SongWithCoupletPaginationAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{id} [get]
func NewGetHandler(sg SongGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req getSongRequest
		if err := ctx.ShouldBindUri(&req.Song); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindQuery(&req.Pagination); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sg.GetSongWithCoupletPagination(ctx, req.Song.ID, req.Pagination)
		if err != nil {
			switch {
			case errors.Is(err, services.ErrSongNotFound):
				ctx.AbortWithError(http.StatusBadRequest, err)

			case errors.Is(err, services.ErrPageNumberOutOfRange):
				ctx.AbortWithError(http.StatusBadRequest, err)

			default:
				ctx.AbortWithError(http.StatusInternalServerError, err)
			}

			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}
