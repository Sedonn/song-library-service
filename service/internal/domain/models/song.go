package models

type Song struct {
	ID          uint64 `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name;index;size:130"`
	Group       string `gorm:"column:group;index;size:130"`
	ReleaseDate string `gorm:"column:release_date;size:10"`
	Text        string `gorm:"column:text;type:text"`
	Link        string `gorm:"column:link;size:150"`
}

// API трансформирует модель БД в модель API.
func (s *Song) API() SongAPI {
	return SongAPI{
		SongIDAPI: SongIDAPI{ID: s.ID},
		SongAttributesAPI: SongAttributesAPI{
			Name:        s.Name,
			Group:       s.Group,
			ReleaseDate: s.ReleaseDate,
			Text:        s.Text,
			Link:        s.Link,
		},
	}
}

type Songs []Song

// API трансформирует слайс моделей БД в слайс моделей API.
func (s *Songs) API() []SongAPI {
	songsAPI := make([]SongAPI, len(*s))
	for i, v := range *s {
		songsAPI[i] = v.API()
	}

	return songsAPI
}

type SongsAPI struct {
	Songs      []SongAPI             `json:"songs"`
	Pagination PaginationMetadataAPI `json:"pagination"`
}

type SongWithCoupletPaginationAPI struct {
	Song       SongAPI               `json:"song"`
	Pagination PaginationMetadataAPI `json:"pagination"`
}

type SongAPI struct {
	SongIDAPI
	SongAttributesAPI
}

type SongIDAPI struct {
	ID uint64 `uri:"id" json:"id" binding:"number,required"`
}

type SongAttributesAPI struct {
	Name        string `json:"name" binding:"required,lte=130"`
	Group       string `json:"group" binding:"required,lte=130"`
	ReleaseDate string `json:"releaseDate" binding:"required,songreleasedate"`
	Text        string `json:"text" binding:"required"`
	Link        string `json:"link" binding:"required,url"`
}

type SongOptionalAttributesAPI struct {
	Name        string `json:"name" binding:"omitempty,lte=130"`
	Group       string `json:"group" binding:"omitempty,lte=130"`
	ReleaseDate string `json:"releaseDate" binding:"omitempty,songreleasedate"`
	Text        string `json:"text" binding:"omitempty"`
	Link        string `json:"link" binding:"omitempty,url"`
}
