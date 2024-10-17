package services

import "errors"

var ( // ErrSongNotFound song_id не найден.
	ErrSongNotFound = errors.New("song not found")

	// ErrSongNotFound artist_id не найден.
	ErrArtistNotFound = errors.New("artist not found")

	// ErrArtistExists artist_name уже существует.
	ErrArtistExists = errors.New("artist already exists")

	// ErrPageNumberOutOfRange номер страницы выходит за границы допустимого диапазона страниц.
	ErrPageNumberOutOfRange = errors.New("page number out of range")
)
