package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testdataFile represents FizzBuzz input & answer files found under the
// testdata directory.
type testdataFile struct {
	input  inputFile
	answer map[string]int
}

// The default logger used in these tests.
var logger *slog.Logger

func init() {

	// setup logger.
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05"))
			}
			return a
		},
	}))
	logger.Debug("Successfully setup testing logger")

	// discover testdata files.
	var inputFiles, answerFiles, errorFiles []string
	walker := func(path string, info os.FileInfo, err error) error {

		// stop discovering files, if there are any errors.
		if err != nil {
			return err
		}

		// skip any directories, since we only care for files.
		if info.IsDir() {
			return nil
		}

		// filter discovered files into their appropriate slice.
		switch {
		case strings.Contains(info.Name(), "input"):
			inputFiles = append(inputFiles, path)
		case strings.Contains(info.Name(), "answer"):
			answerFiles = append(answerFiles, path)
		case strings.Contains(path, "errors"):
			errorFiles = append(errorFiles, path)
		}
		return nil
	}
	if err := filepath.Walk("testdata", walker); err != nil {
		logger.Error("Failed to discover testdata files", "error", err)
		os.Exit(-1)
	}
}

func Test_run(t *testing.T) {

	// setup test-cases.
	tests := map[string]struct {
		want map[string]int
	}{}

	// run tests.
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Error("oh")
			fmt.Println(tt.want)
			fmt.Println("getting here!")
		})
	}
}
