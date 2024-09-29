package songrest

import (
	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/rest/handlers/song/change"
	"github.com/sedonn/song-library-service/internal/rest/handlers/song/create"
	"github.com/sedonn/song-library-service/internal/rest/handlers/song/get"
	"github.com/sedonn/song-library-service/internal/rest/handlers/song/remove"
)

// SongLibraryManager описывает поведение объекта, который обеспечивает бизнес-логику работы с библиотекой песен.
type SongLibraryManager interface {
	create.SongCreator
	get.SongGetter
	change.SongChanger
	remove.SongRemover
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
		songLibrary.GET("/:id", get.NewGetHandler(h.songLibraryManager))
		songLibrary.GET("/", get.NewSearchHandler(h.songLibraryManager))
		songLibrary.POST("/", create.New(h.songLibraryManager))
		songLibrary.PUT("/:id", change.New(h.songLibraryManager))
		songLibrary.DELETE("/:id", remove.New(h.songLibraryManager))
	}
}
