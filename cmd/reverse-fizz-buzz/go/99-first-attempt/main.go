package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"go-simpler.org/env"
)

// A regex for checking if a string contains any numbers.
var numbersRgx = regexp.MustCompile(`[0-9]`)

func main() {

	// TODO: setup tracing.
	ctx := context.Background()

	// setup handler.
	var h handler
	if err := env.Load(&h, nil); err != nil {
		env.Usage(&h, os.Stdout, nil)
		var notSetErr *env.NotSetError
		if errors.As(err, &notSetErr) {
			fmt.Println(notSetErr)
		}
		fmt.Fprintf(os.Stderr, "Failed to setup the handler: %v; Exiting...\n", err)
		os.Exit(-1)
	}

	// setup logger.
	var logLevel slog.Level
	switch strings.ToLower(h.LogLevel) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		fmt.Fprintf(
			os.Stderr,
			"Failed to setup logger: %q is not a valid log level; Exiting...\n",
			h.LogLevel,
		)
		os.Exit(-1)
	}
	h.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05"))
			}
			return a
		},
	}))
	h.logger.Debug("Successfully setup logger")

	// ~run the program!
	h.run(ctx)
}
