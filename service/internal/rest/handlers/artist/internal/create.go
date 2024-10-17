package internal

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
	"golang.org/x/net/context"
)

// ArtistCreator описывает поведение объекта слоя бизнес-логики, который добавляет новых исполнителей.
type ArtistCreator interface {
	// CreateArtist добавляет нового исполнителя.
	CreateArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
}

// NewCreateArtistHandler возвращает новый объект хендлера, который добавляет новых исполнителей.
//
//	@Summary		Добавить нового исполнителя.
//	@Description	Добавить нового исполнителя. Название исполнителя должно быть уникальным.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist	body		models.ArtistAttributesAPI	true	"Данные нового исполнителя"
//	@Success		200		{object}	models.ArtistAPI
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/artists/ [post]
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
