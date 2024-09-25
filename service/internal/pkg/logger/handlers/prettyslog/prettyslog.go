package prettyslog

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Handler это специальный обработчик логгера для форматирования JSON логов.
// Он окрашивает и добавляет отступы логам.
type Handler struct {
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

// New создает новый Handler.
func New(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return &Handler{
		Handler: slog.NewJSONHandler(w, opts),
		l:       log.New(w, "", 0),
	}
}

// Реализует slog.Handler.WithAttrs.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   attrs,
	}
}

// Реализует slog.Handler.WithGroup.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

// Handle форматирует, окрашивает и добавляет отступы в JSON.
//
// Реализует slog.Handler.Handle.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String()

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	attrs := make(map[string]any, r.NumAttrs()+2)
	attrs[slog.MessageKey] = r.Message
	attrs[slog.SourceKey] = recordSource(r)
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()

		return true
	})
	for _, a := range h.attrs {
		attrs[a.Key] = a.Value.Any()
	}

	b, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return err
	}

	h.l.Println(r.Time.Format(time.RFC3339Nano), level, color.WhiteString(string(b)))

	return nil
}

// recordSource возвращает метаданные кода, который вызывает логгер.
func recordSource(r slog.Record) *slog.Source {
	frames := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := frames.Next()

	return &slog.Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}
