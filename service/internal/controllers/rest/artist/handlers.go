package artistrest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// getArtistHandler это хендлер, который возвращает определенного исполнителя.
//
//	@Summary		Получить данные определенного исполнителяяяя.
//	@Description	Получить данные определенного исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		GetArtistRequest	true	"ID исполнителя"
//	@Success		200			{object}	GetArtistResponse
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		404			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [get]
func (e *Endpoints) getArtistHandler(ctx *gin.Context) {
	var req GetArtistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := e.artistService.GetArtist(ctx, req.ID)
	if err != nil {
		if errors.Is(err, services.ErrArtistNotFound) {
			_ = ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, GetArtistResponse(a))
}

// createArtistHandler это хендлер, который добавляет новых исполнителей.
//
//	@Summary		Добавить нового исполнителя.
//	@Description	Добавить нового исполнителя. Название исполнителя должно быть уникальным.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist	body		CreateArtistRequest	true	"Данные нового исполнителя"
//	@Success		200		{object}	CreateArtistResponse
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/artists/ [post]
func (e *Endpoints) createArtistHandler(ctx *gin.Context) {
	var req CreateArtistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := e.artistService.CreateArtist(ctx, models.Artist{Name: req.Name})
	if err != nil {
		if errors.Is(err, services.ErrArtistExists) {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, CreateArtistResponse(a))
}

// changeArtistHandler это хендлер, который обновляет данные исполнителей.
//
//	@Summary		Изменить данные исполнителя.
//	@Description	Изменить данные исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		ChangeArtistRequestPath	true	"ID исполнителя"
//	@Param			artist		body		ChangeArtistRequestBody	true	"Новые данные исполнителя"
//	@Success		200			{object}	ChangeArtistResponse
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		404			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [patch]
func (e *Endpoints) changeArtistHandler(ctx *gin.Context) {
	var req ChangeArtistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := e.artistService.ChangeArtist(ctx, models.Artist{
		ID:   req.ID,
		Name: req.Name,
	})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrArtistNotFound):
			_ = ctx.AbortWithError(http.StatusNotFound, err)

		case errors.Is(err, services.ErrArtistExists):
			_ = ctx.AbortWithError(http.StatusBadRequest, err)

		default:
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, ChangeArtistResponse(a))
}

// removeArtistHandler это хендлер, который удаляет определенного исполнителя.
//
//	@Summary		Удалить данные исполнителя.
//	@Description	Удалить данные исполнителя.
//	@Tags			artist
//	@Accept			json
//	@Produce		json
//	@Param			artist-id	path		RemoveArtistRequest	true	"ID исполнителя"
//	@Success		200			{object}	RemoveArtistResponse
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		404			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/artists/{artist-id} [delete]
func (e *Endpoints) removeArtistHandler(ctx *gin.Context) {
	var req RemoveArtistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := e.artistService.RemoveArtist(ctx, req.ID)
	if err != nil {
		if errors.Is(err, services.ErrArtistNotFound) {
			_ = ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, RemoveArtistResponse(id))
}
