package songrest

import "github.com/sedonn/song-library-service/internal/domain/models"

type GetSongRequest struct {
	Song       GetSongRequestPath
	Pagination GetSongRequestQuery
}

type GetSongRequestPath models.SongIDAPI

type GetSongRequestQuery models.Pagination

type GetSongResponse models.SongWithCoupletPaginationAPI

type SearchSongsRequest struct {
	Name       string `form:"name"`
	ArtistName string `form:"artistName"`
	Link       string `form:"link"`
	Pagination models.Pagination
}

type SearchSongsResponse models.SongsAPI

type CreateSongRequest struct {
	models.SongAttributesAPI
	Artist models.ArtistIDAPI `json:"artist"`
}

type CreateSongResponse models.SongAPI

type ChangeSongRequest struct {
	ChangeSongRequestPath
	ChangeSongRequestBody
}

type ChangeSongRequestPath models.SongIDAPI

type ChangeSongRequestBody struct {
	models.SongOptionalAttributesAPI
	Artist models.ArtistIDAPI `json:"artist" binding:"omitempty"`
}

type ChangeSongResponse models.SongAPI

type RemoveSongRequest models.SongIDAPI

type RemoveSongResponse models.SongIDAPI
