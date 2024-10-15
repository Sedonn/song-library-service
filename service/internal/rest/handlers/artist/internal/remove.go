package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

type ArtistRemover interface {
	RemoveArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
}

func NewArtistRemoveHandler(ar ArtistRemover) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ArtistIDAPI
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ar.RemoveArtist(ctx, models.Artist{
			ID: req.ID,
		})
		if err != nil {
			if errors.Is(err, services.ErrArtistNotFound) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
