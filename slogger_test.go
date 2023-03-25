package slogger

import (
	"context"
	"io"
	"strconv"
	"testing"

	"golang.org/x/exp/slog"
)

func TestFromContext(t *testing.T) {
	var (
		want = slog.New(slog.NewTextHandler(io.Discard))
		have = FromContext(NewContext(context.Background(), want))
	)

	if want != have {
		t.Errorf("expected %v, got %v", want, have)
	}
}

func TestFromContextReturnsDefault(t *testing.T) {
	var (
		want = slog.Default()
		have = FromContext(context.Background())
	)

	if want != have {
		t.Errorf("expected %v, got %v", want, have)
	}
}

func TestFromContextReturnsNil(t *testing.T) {
	have := FromContext(NewContext(context.Background(), nil))

	if have != nil {
		t.Errorf("expected nil, got %#v", have)
	}
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
			handler := logger.Handler()

			if _, ok := handler.(*slog.JSONHandler); kase.json && !ok {
				t.Fatalf("expected *slog.JSONHandler, got %T", handler)
			} else if _, ok := handler.(*slog.TextHandler); !kase.json && !ok {
				t.Fatalf("expected *slog.TextHandler, got %T", handler)
			}

			for level := slog.LevelDebug - 1; level <= slog.LevelError+1; level++ {
				have := logger.Enabled(context.Background(), level)

				if want := level >= kase.minLevel; want != have {
					t.Errorf("expected %t for %s, got %t", want, level, have)
				}
			}
		})
	}
}
