package services

import "errors"

// ErrSongNotFound song_id не найден.
var ErrSongNotFound = errors.New("song not found")

// ErrSongNotFound artist_id не найден.
var ErrArtistNotFound = errors.New("artist not found")

// ErrArtistExists artist_name уже существует.
var ErrArtistExists = errors.New("artist already exists")

// ErrPageNumberOutOfRange номер страницы выходит за границы допустимого диапазона страниц.
var ErrPageNumberOutOfRange = errors.New("page number out of range")
