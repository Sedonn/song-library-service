package artist

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/repositories"
	"github.com/sedonn/song-library-service/internal/services"
	"github.com/sedonn/song-library-service/internal/services/artist/mocks"
)

var (
	discardLogger           = logger.NewDiscardLogger()
	expectedArtistID uint64 = 1
	expectedArtist          = models.Artist{ID: expectedArtistID}
)

func TestService_CreateArtist(t *testing.T) {
	t.Parallel()

	type fields struct {
		artistSaver ArtistSaver
	}
	type args struct {
		ctx context.Context
		a   models.Artist
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArtistAPI
		wantErr error
	}{
		{
			name: "CreateArtist happy path",
			fields: fields{
				artistSaver: func() ArtistSaver {
					as := mocks.NewArtistSaver(t)
					as.
						On("SaveArtist", mock.Anything, expectedArtist).
						Once().
						Return(expectedArtist, nil)

					return as
				}(),
			},
			args: args{
				a: expectedArtist,
			},
			want: expectedArtist.API(),
		},
		{
			name: "CreateArtist error artist exists",
			fields: fields{
				artistSaver: func() ArtistSaver {
					as := mocks.NewArtistSaver(t)
					as.
						On("SaveArtist", mock.Anything, expectedArtist).
						Once().
						Return(models.Artist{}, repositories.ErrArtistExists)
					return as
				}(),
			},
			args: args{
				a: expectedArtist,
			},
			want:    models.ArtistAPI{},
			wantErr: services.ErrArtistExists,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				log:         discardLogger,
				artistSaver: tt.fields.artistSaver,
			}
			got, err := s.CreateArtist(tt.args.ctx, tt.args.a)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "Service.CreateArtist() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestService_GetArtist(t *testing.T) {
	t.Parallel()

	type fields struct {
		artistProvider ArtistProvider
	}
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArtistAPI
		wantErr error
	}{
		{
			name: "GetArtist happy path",
			fields: fields{
				artistProvider: func() ArtistProvider {
					ap := mocks.NewArtistProvider(t)
					ap.
						On("Artist", mock.Anything, expectedArtistID).
						Once().
						Return(expectedArtist, nil)

					return ap
				}(),
			},
			args: args{
				id: expectedArtistID,
			},
			want:    expectedArtist.API(),
			wantErr: nil,
		},
		{
			name: "GetArtist error artist not found",
			fields: fields{
				artistProvider: func() ArtistProvider {
					ap := mocks.NewArtistProvider(t)
					ap.
						On("Artist", mock.Anything, expectedArtistID).
						Once().
						Return(models.Artist{}, repositories.ErrArtistNotFound)

					return ap
				}(),
			},
			args: args{
				id: expectedArtistID,
			},
			want:    models.ArtistAPI{},
			wantErr: services.ErrArtistNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				log:            discardLogger,
				artistProvider: tt.fields.artistProvider,
			}
			got, err := s.GetArtist(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "Service.GetArtist() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestService_ChangeArtist(t *testing.T) {
	t.Parallel()

	type fields struct {
		artistUpdater ArtistUpdater
	}
	type args struct {
		ctx context.Context
		a   models.Artist
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArtistAPI
		wantErr error
	}{
		{
			name: "ChangeArtist happy path",
			fields: fields{
				artistUpdater: func() ArtistUpdater {
					au := mocks.NewArtistUpdater(t)
					au.
						On("UpdateArtist", mock.Anything, expectedArtist).
						Once().
						Return(expectedArtist, nil)

					return au
				}(),
			},
			args: args{
				a: expectedArtist,
			},
			want: expectedArtist.API(),
		},
		{
			name: "ChangeArtist error artist not found",
			fields: fields{
				artistUpdater: func() ArtistUpdater {
					au := mocks.NewArtistUpdater(t)
					au.
						On("UpdateArtist", mock.Anything, expectedArtist).
						Once().
						Return(models.Artist{}, repositories.ErrArtistNotFound)

					return au
				}(),
			},
			args: args{
				a: expectedArtist,
			},
			want:    models.ArtistAPI{},
			wantErr: services.ErrArtistNotFound,
		},
		{
			name: "ChangeArtist error artist exists",
			fields: fields{
				artistUpdater: func() ArtistUpdater {
					au := mocks.NewArtistUpdater(t)
					au.
						On("UpdateArtist", mock.Anything, expectedArtist).
						Once().
						Return(models.Artist{}, repositories.ErrArtistExists)

					return au
				}(),
			},
			args: args{
				a: expectedArtist,
			},
			want:    models.ArtistAPI{},
			wantErr: services.ErrArtistExists,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				log:           discardLogger,
				artistUpdater: tt.fields.artistUpdater,
			}
			got, err := s.ChangeArtist(tt.args.ctx, tt.args.a)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "Service.ChangeArtist() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestService_RemoveArtist(t *testing.T) {
	t.Parallel()

	type fields struct {
		artistDeleter ArtistDeleter
	}
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArtistIDAPI
		wantErr error
	}{
		{
			name: "DeleteArtist happy path",
			fields: fields{
				artistDeleter: func() ArtistDeleter {
					ad := mocks.NewArtistDeleter(t)
					ad.
						On("DeleteArtist", mock.Anything, expectedArtistID).
						Once().
						Return(expectedArtistID, nil)

					return ad
				}(),
			},
			args: args{
				id: expectedArtistID,
			},
			want:    models.ArtistIDAPI{ID: expectedArtistID},
			wantErr: nil,
		},
		{
			name: "DeleteArtist error artist not found",
			fields: fields{
				artistDeleter: func() ArtistDeleter {
					ad := mocks.NewArtistDeleter(t)
					ad.
						On("DeleteArtist", mock.Anything, expectedArtistID).
						Once().
						Return(uint64(0), repositories.ErrArtistNotFound)

					return ad
				}(),
			},
			args: args{
				id: expectedArtistID,
			},
			want:    models.ArtistIDAPI{},
			wantErr: services.ErrArtistNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				log:           discardLogger,
				artistDeleter: tt.fields.artistDeleter,
			}
			got, err := s.RemoveArtist(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "Service.RemoveArtist() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
