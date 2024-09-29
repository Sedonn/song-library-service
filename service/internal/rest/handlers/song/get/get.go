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
	// GetSong возвращает определенную песню.
	GetSong(ctx context.Context, id uint64) (models.Song, error)
	// SearchSongs выполняет поиск песен по определенным параметрам.
	// Поиск выполняется по подстроке каждого указанного поля.
	SearchSongs(ctx context.Context, attrs models.Song, p models.PaginationAPI) (models.SongsAPI, error)
}

// NewGetHandler возвращает новый объект хендлера, который возвращает определенную песню.
func NewGetHandler(sg SongGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.SongIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s, err := sg.GetSong(ctx, req.ID)
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
