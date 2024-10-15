package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

type ArtistGetter interface {
	GetArtist(ctx context.Context, id uint64) (models.ArtistAPI, error)
}

func NewGetArtistHandler(ag ArtistGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ArtistIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ag.GetArtist(ctx, req.ID)
		if err != nil {
			if errors.Is(err, services.ErrArtistNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
