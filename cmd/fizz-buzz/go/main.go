package main

import (
	"context"
	"fmt"
)

// divisor defines a divisor-word pair used to generate a FizzBuzz output.
type divisor struct {
	value int    // the value assigned to the divisor (eg. `2`)
	word  string // the word assigned to the value of the divisor (eg. `rock`)
}

// fizzBuzzInput represents the input to the fizzBuzz function.
type fizzBuzzInput struct {
	n        int       // the number to generate numbers to.
	divisors []divisor // the divisor-word used when generating the output string.
}

// fizzBuzz generates a slice of strings representing a FizzBuzz output, from
// '1' up to the given 'n' integer. The function itself loops over each given
// divisor-word pair and builds a string, per line, based on if the number is
// divisible by any knowm divisor value.
func fizzBuzz(ctx context.Context, input fizzBuzzInput) (out []string) {
	for i := 1; i <= input.n; i++ {

		// setup output line.
		var line string

		// determine if the current index is divisible by any known divisor.
		// if divisible, append the corresponding word to the line.
		for _, d := range input.divisors {
			if i%d.value == 0 {
				line += d.word
			}
		}

		// if no words were appended to the line, add the number itself.
		if line == "" {
			line = fmt.Sprint(i)
		} else {
			line = fmt.Sprintf("%v - %s", i, line)
		}

		// capture line to output.
		out = append(out, line)
	}
	return out
}

func main() {

	// setup tracing.
	ctx := context.Background()

	// setup inputs.
	inputs := []fizzBuzzInput{
		{
			n: 20,
			divisors: []divisor{
				{2, "hello"},
				{7, "world"},
			},
		},
	}

	// ~run!
	for _, i := range inputs {
		for _, line := range fizzBuzz(ctx, i) {
			fmt.Println(line)
		}
	}
}
