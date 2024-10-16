package songrest

import (
	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/rest/handlers/song/internal"
)

// SongService описывает поведение объекта, который обеспечивает бизнес-логику работы с библиотекой песен.
type SongService interface {
	internal.SongCreator
	internal.SongGetter
	internal.SongChanger
	internal.SongRemover
}

// Handler это корневой хендлер библиотеки песен.
type Handler struct {
	songService SongService
}

// New создает новый корневой хендлер библиотеки песен.
func New(m SongService) *Handler {
	return &Handler{
		songService: m,
	}
}

// BindTo привязывает хендлер к определенной группе маршрутов.
func (h *Handler) BindTo(router *gin.RouterGroup) {
	songRouter := router.Group("/songs")
	{
		songRouter.POST("/", internal.NewCreateSongHandler(h.songService))
		songRouter.GET("/:song-id/couplets", internal.NewGetSongCoupletsHandler(h.songService))
		songRouter.GET("/", internal.NewSearchSongsHandler(h.songService))
		songRouter.PUT("/:song-id", internal.NewChangeSongHandler(h.songService))
		songRouter.DELETE("/:song-id", internal.NewRemoveSongHandler(h.songService))
	}
}
