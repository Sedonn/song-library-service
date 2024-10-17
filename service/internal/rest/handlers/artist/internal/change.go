package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

type ArtistChanger interface {
	ChangeArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
}

type changeArtistRequest struct {
	models.ArtistIDAPI
	models.ArtistOptionalAttributesAPI
}

func NewChangeArtistHandler(ac ArtistChanger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req changeArtistRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ac.ChangeArtist(ctx, models.Artist{
			ID:   req.ID,
			Name: req.Name,
		})
		if err != nil {
			switch {
			case errors.Is(err, services.ErrArtistNotFound):
				fallthrough

			case errors.Is(err, services.ErrArtistExists):
				_ = ctx.AbortWithError(http.StatusBadRequest, err)

			default:
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
