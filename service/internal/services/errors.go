package services

import "errors"

// ErrSongNotFound song_id не найден.
var ErrSongNotFound = errors.New("song not found")
