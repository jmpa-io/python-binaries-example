package solution

import (
	"strconv"
)

// NOTES:

// N is a positive integer.

// A binary gap is the maximal sequence of consecutive zeros that is
// surrounded by ones at both ends in the binary representation of N.

// return 0 if N doesn't contain a binary gap.

func Solution(N int) int {

	// convert int to binary string.
	binary := strconv.FormatInt(int64(N), 2)

	// loop over each character in the string, and determine the longest gap.
	currentGap, largestGap := 0, 0
	for _, num := range binary {

		// when we find a 1, it could either be when the gap begins or gap closes.
		if num == '1' {

			// only update the largest gap if the current is larger.
			if currentGap > largestGap {
				largestGap = currentGap
			}

			currentGap = 0 // reset current gap.
			continue
		}

		// increment gap, since we're hitting a 0.
		currentGap++
	}

	return largestGap
}
