package songrest

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sedonn/song-library-service/internal/domain/models"
)

// SongService описывает поведение объекта, который обеспечивает бизнес-логику работы с песнями.
type SongService interface {
	// GetSongWithCoupletPagination возвращает определенную песню с пагинацией по куплетам.
	// Текст разбивается на куплеты по \n\n символам.
	GetSongWithCoupletPagination(ctx context.Context, id uint64, p models.Pagination) (models.SongWithCoupletPaginationAPI, error)
	// SearchSongs выполняет поиск песен по определенным параметрам.
	// Поиск выполняется по подстроке каждого указанного поля.
	SearchSongs(ctx context.Context, attrs models.Song, p models.Pagination) (models.SongsAPI, error)
	// CreateSong добавляют новую песню.
	CreateSong(ctx context.Context, s models.Song) (models.SongAPI, error)
	// ChangeSong обновляет данные определенной песни.
	ChangeSong(ctx context.Context, s models.Song) (models.SongAPI, error)
	// RemoveSong удаляет определенную песню.
	RemoveSong(ctx context.Context, id uint64) (models.SongIDAPI, error)
}

// Endpoints это конечные точки сервиса песен.
type Endpoints struct {
	songService SongService
}

// New создает новый объект конечных точек сервиса песен.
func New(m SongService) *Endpoints {
	return &Endpoints{
		songService: m,
	}
}

// BindTo привязывает конечные точки к определенной группе маршрутов.
func (e *Endpoints) BindTo(router *gin.RouterGroup) {
	songRouter := router.Group("/songs")
	{
		songRouter.GET("/:song-id/couplets", e.getSongCoupletsHandler)
		songRouter.GET("/", e.searchSongsHandler)
		songRouter.POST("/", e.createSongHandler)
		songRouter.PATCH("/:song-id", e.changeSongHandler)
		songRouter.DELETE("/:song-id", e.removeSongHandler)
	}
}
