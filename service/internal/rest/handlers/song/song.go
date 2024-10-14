package songrest

import (
	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/rest/handlers/song/internal"
)

// SongLibraryManager описывает поведение объекта, который обеспечивает бизнес-логику работы с библиотекой песен.
type SongLibraryManager interface {
	internal.SongCreator
	internal.SongGetter
	internal.SongChanger
	internal.SongRemover
}

// Handler это корневой хендлер библиотеки песен.
type Handler struct {
	songLibraryManager SongLibraryManager
}

// New создает новый корневой хендлер библиотеки песен.
func New(m SongLibraryManager) *Handler {
	return &Handler{
		songLibraryManager: m,
	}
}

// BindTo привязывает хендлер к определенной группе маршрутов.
func (h *Handler) BindTo(router *gin.RouterGroup) {
	songLibrary := router.Group("/songs")
	{
		songLibrary.GET("/:id", internal.NewGetSongHandler(h.songLibraryManager))
		songLibrary.GET("/", internal.NewSearchSongsHandler(h.songLibraryManager))
		songLibrary.POST("/", internal.NewCreateSongHandler(h.songLibraryManager))
		songLibrary.PUT("/:id", internal.NewChangeSongHandler(h.songLibraryManager))
		songLibrary.DELETE("/:id", internal.NewRemoveSongHandler(h.songLibraryManager))
	}
}
