package song

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sedonn/song-library-service/internal/domain/models"
	"github.com/sedonn/song-library-service/internal/pkg/logger"
	"github.com/sedonn/song-library-service/internal/repository"
	"github.com/sedonn/song-library-service/internal/services"
	"github.com/sedonn/song-library-service/internal/services/song/mocks"
)

var (
	discardLogger                           = logger.NewDiscardLogger()
	defaultPageNumber                uint64 = 1
	expectedSongOutOfRangePageNumber uint64 = 2
	coupletCountPerPage              uint32 = 1
	expectedSongID                   uint64 = 1
	expectedSongCoupletCount         uint64 = 1
	expectedSong                            = models.Song{
		ID:   expectedSongID,
		Text: "one couplet",
	}
)

func TestSongLibrary_GetSongWithCoupletPagination(t *testing.T) {
	t.Parallel()

	type fields struct {
		songProvider SongProvider
	}
	type args struct {
		ctx context.Context
		id  uint64
		p   models.Pagination
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SongWithCoupletPaginationAPI
		wantErr error
	}{
		{
			name: "GetSongWithCoupletPagination happy path",
			fields: fields{
				songProvider: func() SongProvider {
					sp := mocks.NewSongProvider(t)
					sp.
						On("Song", mock.Anything, expectedSongID).
						Once().
						Return(expectedSong, nil)

					return sp
				}(),
			},
			args: args{
				id: expectedSongID,
				p: models.Pagination{
					PageNumber: defaultPageNumber,
				},
			},
			want: models.SongWithCoupletPaginationAPI{
				Song: expectedSong.API(),
				Pagination: models.PaginationMetadataAPI{
					CurrentPageNumber: defaultPageNumber,
					PageCount:         expectedSongCoupletCount,
					PageSize:          coupletCountPerPage,
					RecordCount:       expectedSongCoupletCount,
				},
			},
		},
		{
			name: "GetSongWithCoupletPagination error song not found",
			fields: fields{
				songProvider: func() SongProvider {
					sp := mocks.NewSongProvider(t)
					sp.
						On("Song", mock.Anything, expectedSongID).
						Once().
						Return(models.Song{}, repository.ErrSongNotFound)

					return sp
				}(),
			},
			args: args{
				id: expectedSongID,
				p: models.Pagination{
					PageNumber: defaultPageNumber,
				},
			},
			want:    models.SongWithCoupletPaginationAPI{},
			wantErr: services.ErrSongNotFound,
		},
		{
			name: "GetSongWithCoupletPagination error page number out of range",
			fields: fields{
				songProvider: func() SongProvider {
					sp := mocks.NewSongProvider(t)
					sp.
						On("Song", mock.Anything, expectedSongID).
						Once().
						Return(expectedSong, nil)

					return sp
				}(),
			},
			args: args{
				id: expectedSongID,
				p: models.Pagination{
					PageNumber: expectedSongOutOfRangePageNumber,
				},
			},
			want:    models.SongWithCoupletPaginationAPI{},
			wantErr: services.ErrPageNumberOutOfRange,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sl := &Service{
				log:          discardLogger,
				songProvider: tt.fields.songProvider,
			}
			got, err := sl.GetSongWithCoupletPagination(tt.args.ctx, tt.args.id, tt.args.p)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "SongLibrary.GetSongWithCoupletPagination() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestSongLibrary_ChangeSong(t *testing.T) {
	t.Parallel()

	type fields struct {
		songUpdater SongUpdater
	}
	type args struct {
		ctx context.Context
		s   models.Song
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SongAPI
		wantErr error
	}{
		{
			name: "ChangeSong happy path",
			fields: fields{
				songUpdater: func() SongUpdater {
					sp := mocks.NewSongUpdater(t)
					sp.
						On("UpdateSong", mock.Anything, expectedSong).
						Once().
						Return(expectedSong, nil)

					return sp
				}(),
			},
			args: args{
				s: expectedSong,
			},
			want: expectedSong.API(),
		},
		{
			name: "ChangeSong error song not found",
			fields: fields{
				songUpdater: func() SongUpdater {
					sp := mocks.NewSongUpdater(t)
					sp.
						On("UpdateSong", mock.Anything, expectedSong).
						Once().
						Return(models.Song{}, repository.ErrSongNotFound)

					return sp
				}(),
			},
			args: args{
				s: expectedSong,
			},
			wantErr: services.ErrSongNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sl := &Service{
				log:         discardLogger,
				songUpdater: tt.fields.songUpdater,
			}
			got, err := sl.ChangeSong(tt.args.ctx, tt.args.s)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "SongLibrary.ChangeSong() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestSongLibrary_RemoveSong(t *testing.T) {
	type fields struct {
		songDeleter SongDeleter
	}
	type args struct {
		ctx context.Context
		s   models.Song
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SongAPI
		wantErr error
	}{
		{
			name: "RemoveSong happy path",
			fields: fields{
				songDeleter: func() SongDeleter {
					sp := mocks.NewSongDeleter(t)
					sp.
						On("DeleteSong", mock.Anything, expectedSong).
						Once().
						Return(expectedSong, nil)

					return sp
				}(),
			},
			args: args{
				s: expectedSong,
			},
			want: expectedSong.API(),
		},
		{
			name: "RemoveSong error song not found",
			fields: fields{
				songDeleter: func() SongDeleter {
					sp := mocks.NewSongDeleter(t)
					sp.
						On("DeleteSong", mock.Anything, expectedSong).
						Once().
						Return(models.Song{}, repository.ErrSongNotFound)

					return sp
				}(),
			},
			args: args{
				s: expectedSong,
			},
			wantErr: services.ErrSongNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &Service{
				log:         discardLogger,
				songDeleter: tt.fields.songDeleter,
			}
			got, err := sl.RemoveSong(tt.args.ctx, tt.args.s)
			assert.Equal(t, tt.want, got)
			assert.ErrorIsf(t, err, tt.wantErr, "SongLibrary.RemoveSong() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
