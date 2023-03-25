package slogger

import (
	"context"
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

func TestFromContext(t *testing.T) {
	var (
		want = slog.New(slog.NewTextHandler(io.Discard))
		have = FromContext(NewContext(context.Background(), want))
	)

	assert.Same(t, want, have)
}

func TestFromContextReturnsDefault(t *testing.T) {
	var (
		want = slog.Default()
		have = FromContext(context.Background())
	)

	assert.Same(t, want, have)
}

func TestFromContextReturnsNil(t *testing.T) {
	have := FromContext(NewContext(context.Background(), nil))

	assert.Nil(t, have)
}

func TestFromEnv(t *testing.T) {
	cases := []struct {
		env      map[string]string // environment variables
		json     bool              // whether to output json or not
		minLevel slog.Level        // minimum verbosity level
	}{
		0: {
			env: map[string]string{
				"LOG_FORMAT": "",
				"LOG_LEVEL":  "",
			},
			minLevel: slog.LevelInfo,
		},
		1: {
			env: map[string]string{
				"LOG_FORMAT": "nonsense",
				"LOG_LEVEL":  "nonsense",
			},
			minLevel: slog.LevelInfo,
		},
		2: {
			env: map[string]string{
				"LOG_FORMAT": "text",
				"LOG_LEVEL":  "debug",
			},
			minLevel: slog.LevelDebug,
		},
		3: {
			env: map[string]string{
				"LOG_FORMAT": "json",
				"LOG_LEVEL":  "warn",
			},
			json:     true,
			minLevel: slog.LevelWarn,
		},
	}

	for caseIndex := range cases {
		kase := cases[caseIndex]

		t.Run(strconv.Itoa(caseIndex), func(t *testing.T) {
			for envVar, varVal := range kase.env {
				t.Setenv(envVar, varVal)
			}

			logger := FromEnv()
			if have := logger.Handler(); kase.json {
				assert.IsType(t, new(slog.JSONHandler), have)
			} else {
				assert.IsType(t, new(slog.TextHandler), have)
			}

			for level := slog.LevelDebug - 1; level <= slog.LevelError+1; level++ {
				have := logger.Enabled(context.Background(), level)

				if want := level >= kase.minLevel; want {
					assert.True(t, have, "level: %s", level)
				} else {
					assert.False(t, have, "level: %s", level)
				}
			}
		})
	}
}
