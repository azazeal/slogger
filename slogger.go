// Package slogger implements functionality around the [slog] package.
package slogger

import (
	"context"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

type contextKey struct{}

// NewContext derives a [context.Context] from ctx that carries logger. You may retrieve logger by
// calling [FromContext] on the returned [context.Context].
func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext returns the [slog.Logger] ctx carries. In case ctx carries no [slog.Logger] calling
// FromContext is a passthrough call to [slog.Default].
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(contextKey{}).(*slog.Logger); ok {
		return logger
	}

	return slog.Default()
}

// FromEnv returns a reference to a [slog.Logger] that'll write records to [os.Stderr] with
// verbosity and format configured by the LOG_LEVEL & LOG_FORMAT environment variables,
// respectively. It's essentially a shortcut to calling [FromEnvWithWriter] with [os.Stderr] as the
// argument.
func FromEnv() *slog.Logger {
	return FromEnvWithWriter(os.Stderr)
}

// FromEnvWithWriter returns a reference to a [slog.Logger] that'll write records to the provided
// [io.Writer] with verbosity and format configured by the LOG_LEVEL & LOG_FORMAT environment
// variables, respectively.
func FromEnvWithWriter(w io.Writer) *slog.Logger {
	opt := slog.HandlerOptions{
		Level: logLevelFromEnv(),
	}

	var handler slog.Handler
	if logJSONFromEnv() {
		handler = opt.NewJSONHandler(w)
	} else {
		handler = opt.NewTextHandler(w)
	}

	return slog.New(handler)
}

func logLevelFromEnv() (level slog.Level) {
	if err := level.UnmarshalText([]byte(os.Getenv("LOG_LEVEL"))); err != nil {
		level = slog.LevelInfo
	}

	return
}

func logJSONFromEnv() bool {
	v := os.Getenv("LOG_FORMAT")
	return strings.EqualFold(v, "json")
}
