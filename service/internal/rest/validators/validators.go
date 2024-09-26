package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var songReleaseDateRegex = regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)

// SongReleaseDate проверяет, соответствует ли строка формату 10.10.2000.
var SongReleaseDate validator.Func = func(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(string); ok {
		return songReleaseDateRegex.Match([]byte(date))
	}

	return false
}
