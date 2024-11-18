package artistrest

import "github.com/sedonn/song-library-service/internal/domain/models"

type GetArtistRequest models.ArtistIDAPI

type GetArtistResponse models.ArtistAPI

type CreateArtistRequest models.ArtistAttributesAPI

type CreateArtistResponse models.ArtistAPI

type ChangeArtistRequest struct {
	ChangeArtistRequestPath
	ChangeArtistRequestBody
}

type ChangeArtistRequestPath models.ArtistIDAPI

type ChangeArtistRequestBody models.ArtistOptionalAttributesAPI

type ChangeArtistResponse models.ArtistAPI

type RemoveArtistRequest models.ArtistIDAPI

type RemoveArtistResponse models.ArtistIDAPI
