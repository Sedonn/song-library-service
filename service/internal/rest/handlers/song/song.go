package songrest

import (
	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/rest/handlers/song/create"
)

// SongLibraryManager описывает поведение объекта, который обеспечивает бизнес-логику работы с библиотекой песен.
type SongLibraryManager interface {
	create.SongCreator
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
	message := router.Group("/song")
	{
		message.POST("/", create.New(h.songLibraryManager))
	}
}
