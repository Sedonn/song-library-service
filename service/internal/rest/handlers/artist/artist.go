package artistrest

import (
	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/rest/handlers/artist/internal"
)

// ArtistService описывает поведение объекта, который обеспечивает бизнес-логику работы с исполнителями.
type ArtistService interface {
	internal.ArtistCreator
	internal.ArtistGetter
	internal.ArtistChanger
	internal.ArtistRemover
}

// Handler это корневой хендлер сервиса исполнителей.
type Handler struct {
	artistService ArtistService
}

// New создает новый корневой хендлер сервиса исполнителей.
func New(s ArtistService) *Handler {
	return &Handler{
		artistService: s,
	}
}

// BindTo привязывает хендлер к определенной группе маршрутов.
func (h *Handler) BindTo(router *gin.RouterGroup) {
	artistRouter := router.Group("/artists")
	{
		artistRouter.POST("/", internal.NewCreateArtistHandler(h.artistService))
		artistRouter.GET("/:artist-id", internal.NewGetArtistHandler(h.artistService))
		artistRouter.PATCH("/:artist-id", internal.NewChangeArtistHandler(h.artistService))
		artistRouter.DELETE("/:artist-id", internal.NewArtistRemoveHandler(h.artistService))
	}
}
