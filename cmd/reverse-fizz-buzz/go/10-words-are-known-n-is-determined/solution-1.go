package main

import (
	"fmt"
	"strings"
)

// func SolutionV1(A []string) (int, int) {
// 	var n, m int
//
// 	// Loop through the list of strings
// 	for i := 0; i < len(A); i++ {
// 		line := A[i]
//
// 		// Print the line and index for debugging purposes
// 		// fmt.Printf("Line: %s, Index (1-based): %d\n", line, i+1)
//
// 		// Adjust the index i to be 1-based for comparison
// 		index := i + 1 // Convert 0-based to 1-based indexing
//
// 		// Check for the conditions that set n and m
// 		if line == "FizzBuzz" && n == 0 && m == 0 {
// 			// First occurrence of FizzBuzz, set both n and m
// 			n = index
// 			m = index
// 		} else if (line == "Fizz" || line == "FizzBuzz") && n == 0 {
// 			// First occurrence of Fizz (or FizzBuzz), set n
// 			n = index
// 		} else if (line == "Buzz" || line == "FizzBuzz") && m == 0 {
// 			// First occurrence of Buzz (or FizzBuzz), set m
// 			m = index
// 		}
//
// 		// If both n and m have been set, return them
// 		if n != 0 && m != 0 {
// 			return n, m
// 		}
// 	}
//
// 	// Return -1, -1 if no solution was found
// 	return -1, -1
// }
//

// findWords is used to filter out numbers from FizzBuzz input & determine the
// sequence of words that appear in-order.
func findWordSequence(A []string) (out []string) {
	for i := 0; i < len(A); i++ {
		line := A[i]
		// ignore numbers.
		if !wordRgx.MatchString(line) {
			continue
		}
		out = append(out, line)
	}
	return out
}

// contains check if a given string value is found in the given array.
func contains(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func allScoresEqualOne(m map[string]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}

func allScoresGreaterThanOne(m map[string]int) bool {
	for _, v := range m {
		if v <= 1 {
			return false
		}
	}
	return true
}

func SolutionV1(A []string) (out []int) {

	// determine sequence of words.
	sequence := findWordSequence(A)

	// find words in sequence, in-order.
	var words []string
	for i := 0; i < len(sequence); i++ {
		line := sequence[i]
		split := splitPascalCase(line)
		for j := 0; j < len(split); j++ {

			// check if the word is already in the word list.
			if contains(words, split[j]) {
				continue
			}
			words = append(words, split[j])
		}
	}

	// determine when words intersect.
	intersectWord := strings.Join(words, "")

	fmt.Println("sequence")
	fmt.Println(sequence)
	fmt.Println("words")
	fmt.Println(words)
	fmt.Println(intersectWord)

	// setup scores.
	scores := make(map[string]int)
	for _, w := range words {
		scores[w] = 1
	}

	// score words.
	for i := 1; i < len(A); i++ {
		line := A[i-1]

		if line == intersectWord && allScoresEqualOne(scores) {
			for word := range scores {
				scores[word] = i
			}
		} else {

			for j := 0; j < len(words); j++ {
				word := words[j]
				if (line == word || line == intersectWord) && scores[word] == 0 {
					scores[word] = i
				}
			}
		}

		if allScoresGreaterThanOne(scores) {
			break
		}
	}

	// build output.
	for i := 0; i < len(words); i++ {
		out = append(out, scores[words[i]])
	}
	return out

	//
	// // Loop through the list of strings
	// for i := 1; i < len(A); i++ {
	//
	// 	line := A[i-1]
	//
	// 	// Print the line and index for debugging purposes
	// 	// fmt.Printf("Line: %s, Index (1-based): %d\n", line, i+1)
	//
	// 	// Adjust the index i to be 1-based for comparison
	// 	index := i + 1 // Convert 0-based to 1-based indexing
	//
	// 	// Check for the conditions that set n and m
	// 	if line == "FizzBuzz" && n == 0 && m == 0 {
	// 		// First occurrence of FizzBuzz, set both n and m
	// 		n = index
	// 		m = index
	// 	} else if (line == "Fizz" || line == "FizzBuzz") && n == 0 {
	// 		// First occurrence of Fizz (or FizzBuzz), set n
	// 		n = index
	// 	} else if (line == "Buzz" || line == "FizzBuzz") && m == 0 {
	// 		// First occurrence of Buzz (or FizzBuzz), set m
	// 		m = index
	// 	}
	//
	// 	// If both n and m have been set, return them
	// 	if n != 0 && m != 0 {
	// 		return n, m
	// 	}
	// }
	//
	// // Return -1, -1 if no solution was found
	// return -1, -1
}
