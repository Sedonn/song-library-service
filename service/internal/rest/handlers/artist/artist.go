package artistrest

import (
	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/rest/handlers/artist/internal"
)

type ArtistService interface {
	internal.ArtistCreator
	internal.ArtistGetter
	internal.ArtistChanger
	internal.ArtistRemover
}

type Handler struct {
	artistService ArtistService
}

func New(s ArtistService) *Handler {
	return &Handler{
		artistService: s,
	}
}

func (h *Handler) BindTo(router *gin.RouterGroup) {
	artistRouter := router.Group("/artists")
	{
		artistRouter.POST("/", internal.NewCreateArtistHandler(h.artistService))
		artistRouter.GET("/:artist-id", internal.NewGetArtistHandler(h.artistService))
		artistRouter.PUT("/:artist-id", internal.NewChangeArtistHandler(h.artistService))
		artistRouter.DELETE("/:artist-id", internal.NewArtistRemoveHandler(h.artistService))
	}
}
