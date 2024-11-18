package artistrest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sedonn/song-library-service/internal/domain/models"
)

// ArtistService описывает поведение объекта, который обеспечивает бизнес-логику работы с исполнителями.
type ArtistService interface {
	// GetArtist получает данные определенного исполнителя.
	GetArtist(ctx context.Context, id uint64) (models.ArtistAPI, error)
	// CreateArtist добавляет нового исполнителя.
	CreateArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
	// ChangeArtist обновляет данные определенного исполнителя.
	ChangeArtist(ctx context.Context, a models.Artist) (models.ArtistAPI, error)
	// RemoveArtist удаляет определенного исполнителя.
	RemoveArtist(ctx context.Context, id uint64) (models.ArtistIDAPI, error)
}

// Endpoints это конечные точки сервиса исполнителей.
type Endpoints struct {
	artistService ArtistService
}

// New создает новый объект конечных точек сервиса исполнителей.
func New(s ArtistService) *Endpoints {
	return &Endpoints{
		artistService: s,
	}
}

// BindTo привязывает конечные точки к определенной группе маршрутов.
func (e *Endpoints) BindTo(router *gin.RouterGroup) {
	artistRouter := router.Group("/artists")
	{
		artistRouter.GET("/:artist-id", e.getArtistHandler)
		artistRouter.POST("/", e.createArtistHandler)
		artistRouter.PATCH("/:artist-id", e.changeArtistHandler)
		artistRouter.DELETE("/:artist-id", e.removeArtistHandler)
	}
}
