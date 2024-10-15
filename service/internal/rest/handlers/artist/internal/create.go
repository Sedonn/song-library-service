package internal

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
	"golang.org/x/net/context"
)

type ArtistCreator interface {
	CreateArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
}

func NewCreateArtistHandler(ac ArtistCreator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ArtistAttributesAPI
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		a, err := ac.CreateArtist(ctx, models.Artist{Name: req.Name})
		if err != nil {
			if errors.Is(err, services.ErrArtistExists) {
				_ = ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, a)
	}
}
