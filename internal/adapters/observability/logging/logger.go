package logging

import (
	"io"
	"log/slog"
)

func New(writer io.Writer, json bool) *slog.Logger {
	options := &slog.HandlerOptions{Level: slog.LevelInfo}
	if json {
		return slog.New(slog.NewJSONHandler(writer, options))
	}

	return slog.New(slog.NewTextHandler(writer, options))
}
