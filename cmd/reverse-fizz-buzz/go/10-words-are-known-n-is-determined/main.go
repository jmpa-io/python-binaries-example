package main

import (
	"fmt"
	"os"
	"path/filepath"
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

		// Run solution.
		fmt.Printf("%v. ", i+1)
		fmt.Println(SolutionV1(lines))
	}

	fmt.Printf("Solution V2\n")
	for i, f := range files {

		// Read input file.
		b, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}

		// Determine lines.
		lines := strings.Split(strings.TrimSpace(string(b)), "\n")

		// Run solution.
		fmt.Printf("%v. ", i+1)
		fmt.Println(SolutionV2(lines))
	}

}
