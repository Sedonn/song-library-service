package songrest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/services"
)

// getSongCoupletsHandler это хендлер, который возвращает определенную песню с пагинацией по куплетам.
//
//	@Summary		Получить данные определенной песни.
//	@Description	Получить данные определенной песни с пагинацией по куплетами.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song-id		path		GetSongRequestPath	true	"ID песни"
//	@Param			pagination	query		GetSongRequestQuery	true	"Настройки пагинации. pageSize игнорируется."
//	@Success		200			{object}	GetSongResponse
//	@Failure		400			{object}	mwerror.ErrorResponse
//	@Failure		404			{object}	mwerror.ErrorResponse
//	@Failure		500			{object}	mwerror.ErrorResponse
//	@Router			/songs/{song-id}/couplets [get]
func (e *Endpoints) getSongCoupletsHandler(ctx *gin.Context) {
	var req GetSongRequest
	if err := ctx.ShouldBindUri(&req.Song); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctx.ShouldBindQuery(&req.Pagination); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s, err := e.songService.GetSongWithCoupletPagination(ctx, req.Song.ID, models.Pagination(req.Pagination))
	if err != nil {
		switch {
		case errors.Is(err, services.ErrSongNotFound):
			_ = ctx.AbortWithError(http.StatusNotFound, err)

		case errors.Is(err, services.ErrPageNumberOutOfRange):
			_ = ctx.AbortWithError(http.StatusBadRequest, err)

		default:
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	ctx.JSON(http.StatusOK, GetSongResponse(s))
}

// searchSongsHandler это хендлер, который выполняет поиск песен по определенным параметрам.
//
//	@Summary		Поиск определенной песни.
//	@Description	Поиск определенной песни по всем атрибутам.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	query		SearchSongsRequest	true	"Настройки поиска."
//	@Success		200		{object}	SearchSongsResponse
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/ [get]
func (e *Endpoints) searchSongsHandler(ctx *gin.Context) {
	var req SearchSongsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	songs, err := e.songService.SearchSongs(ctx,
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

	ctx.JSON(http.StatusOK, SearchSongsResponse(songs))
}

// createSongHandler это хендлер, который добавляет новые песни.
//
//	@Summary		Добавить новую песню.
//	@Description	Добавление новой песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	body		CreateSongRequest	true	"Данные новой песни"
//	@Success		200		{object}	CreateSongResponse
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		404		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/ [post]
func (e *Endpoints) createSongHandler(ctx *gin.Context) {
	var req CreateSongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s, err := e.songService.CreateSong(ctx, models.Song{
		Name:        req.Name,
		ArtistID:    req.Artist.ID,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	})
	if err != nil {
		if errors.Is(err, services.ErrArtistNotFound) {
			_ = ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, CreateSongResponse(s))
}

// changeSongHandler это хендлер, который обновляет песни.
//
//	@Summary		Изменить данные песни.
//	@Description	Изменить данные песни. Для разделения куплетов необходимо использовать '\n\n'.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song-id	path		ChangeSongRequestPath true	"ID песни"
//	@Param			song	body		ChangeSongRequestBody	true	"Новые данные песни"
//	@Success		200		{object}	ChangeSongResponse
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		404		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{song-id} [patch]
func (e *Endpoints) changeSongHandler(ctx *gin.Context) {
	var req ChangeSongRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s, err := e.songService.ChangeSong(ctx, models.Song{
		ID:          req.ID,
		ArtistID:    req.Artist.ID,
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrArtistNotFound):
			_ = ctx.AbortWithError(http.StatusNotFound, err)

		case errors.Is(err, services.ErrSongNotFound):
			_ = ctx.AbortWithError(http.StatusBadRequest, err)

		default:
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, ChangeSongResponse(s))
}

// removeSongHandler это хендлер, который удаляет определенные песни.
//
//	@Summary		Удалить данные песни.
//	@Description	Удалить данные песни.
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song-id	path		RemoveSongRequest	true	"ID песни"
//	@Success		200		{object}	RemoveSongResponse
//	@Failure		400		{object}	mwerror.ErrorResponse
//	@Failure		404		{object}	mwerror.ErrorResponse
//	@Failure		500		{object}	mwerror.ErrorResponse
//	@Router			/songs/{song-id} [delete]
func (e *Endpoints) removeSongHandler(ctx *gin.Context) {
	var req RemoveSongRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s, err := e.songService.RemoveSong(ctx, req.ID)
	if err != nil {
		if errors.Is(err, services.ErrSongNotFound) {
			_ = ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, RemoveSongResponse(s))
}
