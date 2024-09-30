package repository

import "errors"

// ErrSongNotFound song_id не найден.
var ErrSongNotFound = errors.New("song not found")

// ErrPageNumberOutOfRange номер страницы выходит за границы допустимого диапазона страниц.
var ErrPageNumberOutOfRange = errors.New("page number out of range")
