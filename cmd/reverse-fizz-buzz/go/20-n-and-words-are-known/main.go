package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	files, err := filepath.Glob("input-*")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution V1\n")
	for i, f := range files {

		// Read input file.
		b, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}

		// Determine lines.
		lines := strings.Split(strings.TrimSpace(string(b)), "\n")

		// Determine n from first line.
		n, err := strconv.Atoi(lines[0])
		if err != nil {
			panic(err)
		}
		lines = lines[1:]

		// Run solution.
		fmt.Printf("%v. ", i+1)
		fmt.Println(SolutionV1(n, lines))
	}

}
