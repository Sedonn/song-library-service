package logger

import "log/slog"

// ErrorString создает аттрибут slog для вывода текста ошибки.
func ErrorString(err error) slog.Attr {
	return slog.String("error", err.Error())
}
