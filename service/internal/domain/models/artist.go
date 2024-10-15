package models

type Artist struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;uniqueIndex;size:130"`
}

func (a Artist) API() ArtistAPI {
	return ArtistAPI{
		ArtistIDAPI:         ArtistIDAPI{ID: a.ID},
		ArtistAttributesAPI: ArtistAttributesAPI{Name: a.Name},
	}
}

type ArtistAPI struct {
	ArtistIDAPI
	ArtistAttributesAPI
}

type ArtistIDAPI struct {
	ID uint64 `uri:"artist-id" json:"id" binding:"required,number"`
}

type ArtistAttributesAPI struct {
	Name string `json:"name" binding:"required,lte=130"`
}

type ArtistOptionalAttributesAPI struct {
	Name string `json:"name" binding:"omitempty,lte=130"`
}
