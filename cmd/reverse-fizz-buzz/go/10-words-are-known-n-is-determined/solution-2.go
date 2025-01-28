package main

import (
	"fmt"
	"regexp"
	"unicode"
)

// wordRgx is used to determine if a string is a word or not.
var wordRgx *regexp.Regexp

func init() {
	wordRgx = regexp.MustCompile(`^[a-zA-Z]+$`)
}

// splitPascalCase splits a string into PascalCase words.
func splitPascalCase(s string) (out []string) {

	// Loop over each character and determine the starting position of any
	// uppercase characters to split on.
	var position int
	for i, character := range s {

		// Check if the current character is uppercase, and append the substring
		// before the current uppercase letter as a new word.
		if unicode.IsUpper(character) && i > 0 {
			out = append(out, s[position:i])
			position = i
		}
	}

	// Append any characters after the last uppercase letter in the string.
	out = append(out, s[position:])
	return out
}

func SolutionV2(A []string) []int {

	// count words.
	counter := make(map[string]int)
	for _, a := range A {

		// ignore numbers.
		if !wordRgx.MatchString(a) {
			continue
		}

		for _, word := range splitPascalCase(a) {
			counter[word]++
		}
	}

	fmt.Println(len(A))

	// determine divisors.
	divisors := make(map[string]int)
	for word, count := range counter {
		divisors[word] = len(A) / count
		fmt.Printf("-- %s %v\n", word, count)
	}

	// build output.
	var out []int
	for _, d := range divisors {
		out = append(out, d)
	}
	return out

}
