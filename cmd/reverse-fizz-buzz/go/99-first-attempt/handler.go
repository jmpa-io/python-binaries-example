package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// handler handles the state across this program.
type handler struct {
	TestdataDir             string `env:"TESTDATA_DIR"               usage:"The path to the testdata directory."                                    default:"testdata"`
	LogLevel                string `env:"LOG_LEVEL"                  usage:"The log level used for this service."                                   default:"debug"`
	ExpectedColumnsInHeader int    `env:"EXPECTED_COLUMNS_IN_HEADER" usage:"The number of expected columns in the header of a FizzBuzz input file." default:"2"`

	// misc.
	logger *slog.Logger
}

// inputFile refers to the input of a FizzBuzz input file.
type inputFile struct {
	wordCount int      // The number of different words found in the wordList.
	totalSize int      // The overall size of the FizzBuzz.
	wordList  []string // A slice of word sequences used to determine the divisors in the FizzBuzz.
}

// run is essentially the 'main' function of this file.
func (h *handler) run(ctx context.Context) {

	// check if testdata directory exists.
	if _, err := os.Stat(h.TestdataDir); err != nil {
		h.logger.Error("The testdata directory doesn't exist", "directory", h.TestdataDir)
		os.Exit(-1)
	}

	// recursively discover testdata files.
	var files []string
	walker := func(path string, info os.FileInfo, err error) error {

		// stop discovering files, if there are any errors.
		if err != nil {
			return err
		}

		// skip any directories, since we only care for files.
		if info.IsDir() {
			return nil
		}

		// skip files intended for testing or answers.
		switch {
		case strings.Contains(path, "errors"):
			fallthrough
		case strings.Contains(info.Name(), "answer"):
			return nil
		}

		// capture any files.
		files = append(files, path)
		return nil
	}
	if err := filepath.Walk(h.TestdataDir, walker); err != nil {
		h.logger.Error(
			"Failed to discover files in testdata directory",
			"directory",
			h.TestdataDir,
			"error",
			err,
		)
		os.Exit(-1)
	}
	h.logger.Debug("Successfully discovered input files", "count", len(files))

	// read contents of each file.
	fileContents := make(map[string]inputFile)
	var errs []error
	for _, f := range files {

		h.logger.Debug("Reading input file", "file", f)

		infile, err := readInputFile(ctx, h.ExpectedColumnsInHeader, f)
		if err != nil {
			h.logger.Error("Failed to read input file", "file", f, "error", err)
			errs = append(errs, err)
			continue
		}

		// ASSUMPTION: since file names are unique, there is no possibility of this
		// having overwritten values.
		fileContents[f] = infile
	}

	// stop if any errors are found.
	if len(errs) > 0 {
		os.Exit(-1)
	}

	// do the goods!
	for file, input := range fileContents {

		fmt.Printf("----- %s\n", file)
		// fmt.Printf("%v %v\n", input.wordCount, input.totalSize)
		// fmt.Printf("%s\n", input.wordList)

		for i := 1; i <= input.totalSize; i++ {

			switch {
			case i%3 == 0:
				fmt.Println("fizz")

			case i%5 == 0:
				fmt.Println("buzz")
			default:
				fmt.Println(i)
			}
		}

	}
}

// parseHeader parses the header of a FizzBuzz input file. An error is returned
// if there is a mismatch between the EXPECTED_COLUMNS_IN_HEADER and the found
// number of columns in the header.
func parseInputFileHeader(
	ctx context.Context,
	expectedColumnsInHeader int,
	header string,
) (uniqueWords int, n int, err error) {

	// check if the header given contains numbers only.
	if !numbersRgx.MatchString(header) {
		return 0, 0, fmt.Errorf("header %q doesn't contain any numbers", header)
	}

	split := strings.Split(header, " ")

	// validate the expected headers match what is given.
	if len(split) != expectedColumnsInHeader {
		return 0, 0, fmt.Errorf(
			"expected headers mismatch; got=%v, want=%v",
			len(split),
			expectedColumnsInHeader,
		)
	}

	// determine output.
	for i := range expectedColumnsInHeader {
		val, err := strconv.Atoi(split[i])
		if err != nil {
			return 0, 0, fmt.Errorf("Failed to convert %q to int: %v", split[i], err)
		}
		switch i {
		case 0:
			uniqueWords = val
		case 1:
			n = val
		default:
			return 0, 0, fmt.Errorf(
				"unexpected number of columns in header found; please ask the developer to implement %q column implementation in header",
				val,
			)
		}
	}
	return uniqueWords, n, nil
}

// parseInputWordList parses the wordlist of a FizzBuzz input file, where uniqueWords
// is the number of unique words expected in the wordList. An error is returned
// if the number of words in the wordList doesn't match the number of
// uniqueWords given to the function.
func parseInputWordList(
	ctx context.Context,
	uniqueWords int,
	wordlist []string,
) (out []string, err error) {

	wordMonitor := make(map[string]int)
	for _, w := range wordlist {
		out = append(out, w)

		// monitor the words coming through.
		if !strings.Contains(w, " ") {
			wordMonitor[w]++
		} else {
			split := strings.Split(w, " ")
			for _, s := range split {
				wordMonitor[s]++
			}
		}
	}

	// validate if there is a mismatch between the number of words and the word
	// count given.
	if len(wordMonitor) != uniqueWords {
		return nil, fmt.Errorf(
			"found %v words in the wordList, when expecting %v; words=%v",
			len(wordMonitor),
			uniqueWords,
			wordMonitor,
		)
	}

	return out, err
}

// readInputFile reads the given path to a file (pathToFile) & parses it's contents.
func readInputFile(
	ctx context.Context,
	expectedColumnsInHeader int,
	file string,
) (infile inputFile, err error) {

	// read file contents.
	b, err := os.ReadFile(file)
	if err != nil {
		return inputFile{}, fmt.Errorf("failed to read file: %v", err)
	}

	// break file contents into lines + remove any empty lines.
	lines := strings.FieldsFunc(string(b), func(r rune) bool {
		return r == '\n'
	})

	fmt.Println(len(lines))

	infile.wordCount, infile.totalSize, err = parseInputFileHeader(
		ctx,
		expectedColumnsInHeader,
		lines[0],
	)
	if err != nil {
		return inputFile{}, fmt.Errorf("failed to parse header of input file: %v", err)
	}

	infile.wordList, err = parseInputWordList(ctx, infile.wordCount, lines[1:])
	if err != nil {
		return inputFile{}, fmt.Errorf("failed to parse wordlist in input file: %v", err)
	}

	return infile, err
}
